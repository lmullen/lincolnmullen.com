---
layout: post
title: "Modeling Historical Events and Lives in YAML"
date: 2012-05-04 09:16
comments: true
categories:
- blog
- digital humanities
- dissertation
- code
tags:
- Ruby
- YAML
---

{% img right http://upload.wikimedia.org/wikipedia/commons/thumb/8/81/Orestes_Brownson_by_GPA_Healy%2C_1863.jpg/375px-Orestes_Brownson_by_GPA_Healy%2C_1863.jpg 375 475 "Orestes Brownson" "Orestes Brownson" %}

For my [dissertation][], I am researching the lives of converts from the
nineteenth century. Some people who converted left behind an enormous
source base. [Orestes Brownson][] converted from Congregationalism to
Presbyterianism to Universalism to Unitarianism to Transcendentalism to
Catholicism, publishing voluminously all along the way. For other
converts, I can find the barest of mentions in a newspaper or collection
of papers. The dissertation needs to get both at the experience of
well-known, articulate converts like Brownson, and lesser- or unknown
converts. To retrieve that second kind of experience, I want to try
analyzing all the conversions as data.

As I compile my research, I want to use it for two purposes. First, I
need regular research notes to use when writing the dissertation.
Second, I'd like to use the research as data, which I'll analyze from
some unknown tool (maybe Ruby). I have an idea of some of the questions
that I'll ask: How many people converted from X to Y? How likely were
converts who were clergy in one religion likely to become clergy in
another? How were conversions distributed over time? over space? But I
won't know which questions can be investigated programmatically or what
the data to answer them will look like until I've done substantially
more research.

#### The idea: use YAML to model lives and events

With that research problem in mind, I've drawn up a list of
specifications for what my data model should look like.

1.  The data must be human-readable and -writable as research notes.
2.  The data model must be able to grow organically as I do the
    research.
3.  The data model must be able to hold large amounts of undigested
    text as notes.
4.  The data must be portable to other formats, possible JSON or
    XML/TEI.

My idea is to use [YAML][] as the format for the data. YAML is a "human
friendly data serialization standard for all programming languages."
YAML's two top priorities are "YAML is easily readable by humans" and
"YAML data is portable between programming languages," which match my
own priorities. I'm familiar with YAML from using [Jekyll][] for this
blog and [another web project][]. YAML also fits well into the
principles I learned from [*Linux and the Unix Philosophy*][],
especially "store data in flat text files."

#### Example YAML model and Ruby script

I've created a working example with two YAML files and a Ruby script to
output some of the data. I've shared the example as a [Gist on GitHub][].

The YAML file for Orestes Brownson is below, and there is another sample
file for Charles Wharton in the [Gist][Gist on GitHub]. You'll notice
that at the outermost level of indentation, there are keys and values
for basic biographical information, such as `born: 1803-09-16`. The most
important part of the model is the list of conversions, which is a YAML
array as signaled by the `-` character and indentation. The markup for the notes
field (`notes: >`) lets that field contain as many paragraphs as
necessary. Finally, the `source` array has one value
(`@carey_orestes_2004`) which is the key to an entry in [my BibTeX database][], 
which I've added with [Vim's][] autocomplete function.

{% highlight yaml %}
# A model of a convert's life
---
name-last       : Brownson
name-first      : Orestes Augustus
born            : 1803-09-16
died            : 1876-04-17
birth-religion  : Congregationalism

conversions     :

-   origin-religion         : Congregationalism
    destination-religion    : Presbyterianism
    date                    : 1822
    ritual                  : church membership
    citation                : ANB
    notes                   : >
      Brownson's change to congregationalism was more denominational 
      switching than a change in conscience.

-   origin-religion         : Presbyterianism
    destination-religion    : Universalism
    date                    : 1826
    ritual                  : ordination
    location                : "Jaffrey, New Hampshire"
    citation                : ANB
    notes                   : >
      "He would later refer to his years in this fold as 'the most 
      anti-Christian period of my life'" (ANB).

      Brownson was editor of _The Gospel Advocate and Impartial 
      Investigator_, a Universalist publication.

-   origin-religion         : Universalism
    destination-religion    : Unitarianism
    ritual                  : further research
    location                : "Walpole, New Hampshire"
    citation                : ANB
    notes                   : >
      Brownson spent some time at Brook Farm, which prepared him for 
      Transcendentalism

-   origin-religion         : Unitarianism and Transcendentalism
    destination-religion    : Catholicism
    date                    : 1844-10-19
    ritual                  : baptism
    citation                : ANB
    notes                   : >
      Brownson studied after his conversion with a Sulpician priest.

source          :
-   @carey_orestes_2004
-   American National Biography

comments        : >
  This is a minimal example of what a model of a convert might look 
  like. The historical data is hastily gathered, so only the model is 
  of interest here.

  N.B. I would like to replace the citations with BibTeX keys.
...
{% endhighlight %}

I had to prove to myself that I could get at the data programmatically, so I
wrote the Ruby script below. It's just a proof-of-concept, and it's the
first Ruby script I've written, so there are ugly parts. The script
creates a class `Converts`, which loads an array of YAML files into a
hash. The class has a few methods to display the names of the converts
and a list of all the conversions. Doubtless there are more interesting
things that can be done. 

{% highlight ruby %}
#!/usr/bin/env ruby
# A proof-of-concept script that outputs some simple data from YAML 
# files modeling conversions
#
# Author:: Lincoln Mullen (lincoln@lincolnmullen.com)

require 'yaml' 

# This class loads data from YAML files, and outputs some values

class Converts

  attr_accessor :files, :data

  def initialize (files = nil, data = nil)
    @files = files
    @data = Hash.new

    if @files.nil?
      puts "You didn't pass me any files."
    elsif @files.respond_to?("each")
      # walk through the array of files, creating a hash with the 
      # file name as the key and the file data as the value
      @files.each do |file|
        @data[file] = YAML.load_file( file )
      end
    end
  end

  # output the hash we can see what we're working with
  def display_raw
    puts "\nThis is the raw data we have loaded:"
    p( @data )
  end

  # walk through the hash, outputting the names of each person
  def display_names
    puts "\nThese people converted:"
    @data.each_key do |key|
      puts " - #{@data[key]["name-first"]} #{@data[key]["name-last"]}"
    end
  end

  # walk through the hash, outputting the names and conversions of 
  # each person
  def display_conversions
    puts "\nWe know about these conversions:"
    @data.each_key do |key|
      puts " - #{@data[key]["name-first"]} #{@data[key]["name-last"]}:"
      # each person has an array of conversions (even if there is 
      # only one conversion)
      @data[key]["conversions"].each { |conversion|
        puts "     + From #{conversion["origin-religion"]} to #{conversion["destination-religion"]} by #{conversion["ritual"]} in #{conversion["date"]}."
      }
    end
  end

end

# get sample data by loading every YAML file in the directory
puts "Let's load all the YAML files in this directory:"
puts Dir.glob( '*.yml').join(', ')
c = Converts.new(Dir.glob('*.yml'))

# call the methods to display the names and conversions
c.display_names
c.display_conversions
{% endhighlight %}

If you run the script on the sample YAML files, you get the output
below. (Yes---the script does output in [Markdown][].
I only know one trick.)

{% highlight bash %}
Let's load all the YAML files in this directory:
brownson-orestes.yml, wharton-charles.yml

These people converted:
 - Charles Wharton
 - Orestes Augustus Brownson

We know about these conversions:
 - Charles Wharton:
     + From Catholicism to Church of England by conformity in .
 - Orestes Augustus Brownson:
     + From Congregationalism to Presbyterianism by church membership in 1822.
     + From Presbyterianism to Universalism by ordination in 1826.
     + From Universalism to Unitarianism by further research in .
     + From Unitarianism and Transcendentalism to Catholicism by baptism in 1844-10-19.
{% endhighlight %}

#### What's next?

If this model works for modeling conversions, it should also work for
modeling other kinds of historical events. For example, suppose a labor
historian is researching strikes and kept a YAML file for each strike
...

{% highlight yaml  %}
id:	Pullman strike
location: Pullman, Illinois
date: 1894-05-11
corporations:
-	Pullman Palace Car Company
unions:
-	American Railway Union
accounts:

-	name: John A. Doe
	source: Chicago Tribune
	description: >
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris
	malesuada, purus vel posuere aliquam, enim orci tempor quam, ac
	rutrum arcu arcu nec leo."

-	name: Jane B. Doe
	source: New York Times
	description: >
	"Maecenas in velit nulla, pretium vestibulum lacus. Morbi dui purus,
	imperdiet ac aliquam sodales, gravida ut diam. Vestibulum nec erat a
	ligula tincidunt dignissim in et diam. Quisque tincidunt
	pellentesque lorem, a scelerisque quam lacinia vitae."
{% endhighlight %}

and another for each union ...

{% highlight yaml  %}
union: American Railway Union
leaders:	
	-	name: Eugene V. Debs
		start: 1893-06-20
		end: ~
founded:
	date: 1893-06-20
	place: Chicago, Illinois
{% endhighlight %}


I asked about this idea at [Digital Humanities Questions & Answers][]
and on Twitter. [Chad Black][], [Ben Brumfield][], Ethan Gruber, [Caleb McDaniel][], 
and Conal Tuohy offered valuable advice about how to think
about this problem and what tools might be helpful later in the project.
The [TEI][] markup for an [event][] and [person][] (recommended by
Conal) seems promising because it can accommodate types of data that I
know I'll need, such as uncertain dates and name changes.

For now, though, I'm going to work with YAML, since I can get started on
it right away and since I'm completely sure it will work as research
notes and reasonably sure it can be munged into another format later.

I'll be glad for any advice about how to improve the data model or
script and about what considerations I should think about to make sure the
data is useful. If you have any ideas about what to do with the data
once I've gathered it, I'll be glad for those too.

  [Orestes Brownson]: http://en.wikipedia.org/wiki/Orestes_Brownson
  [YAML]: http://yaml.org/
  [Jekyll]: http://jekyllrb.com
  [another web project]: http://jsr.fsu.edu
  [*Linux and the Unix Philosophy*]: http://openlibrary.org/works/OL3494423W/Linux_and_the_Unix_philosophy
  [Digital Humanities Questions & Answers]: http://digitalhumanities.org/answers/topic/using-yaml-to-model-historical-lives-or-events
  [Chad Black]: http://parezcoydigo.wordpress.com/
  [Ben Brumfield]: http://manuscripttranscription.blogspot.com/
  [Caleb McDaniel]: http://www.owlnet.rice.edu/~wcm1/
  [TEI]: http://www.tei-c.org/
  [event]: http://www.tei-c.org/release/doc/tei-p5-doc/en/html/ref-event.html
  [person]: http://www.tei-c.org/release/doc/tei-p5-doc/en/html/ref-person.html
  [Gist on GitHub]: https://gist.github.com/2580742
  [my BibTeX database]: https://github.com/lmullen/historybib
  [Vim's]: http://www.vim.org/
  [Markdown]: http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295
  [dissertation]: http://lincolnmullen.com/blog/categories/dissertation/
