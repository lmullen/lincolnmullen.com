---
date: '2012-05-04T13:16:00+00:00'
slug: 'modeling-historical-events-and-lives-in-yaml'
title: Modeling Historical Events and Lives in YAML
---

For my [dissertation](http://lincolnmullen.com/blog/categories/dissertation/), I am researching the lives of converts from the nineteenth century. Some people who converted left behind an enormous source base. [Orestes Brownson](http://en.wikipedia.org/wiki/Orestes_Brownson) converted from Congregationalism to Presbyterianism to Universalism to Unitarianism to Transcendentalism to Catholicism, publishing voluminously all along the way. For other converts, I can find the barest of mentions in a newspaper or collection of papers. The dissertation needs to get both at the experience of well-known, articulate converts like Brownson, and lesser- or unknown converts. To retrieve that second kind of experience, I want to try analyzing all the conversions as data.

As I compile my research, I want to use it for two purposes. First, I need regular research notes to use when writing the dissertation. Second, I'd like to use the research as data, which I'll analyze from some unknown tool (maybe Ruby). I have an idea of some of the questions that I'll ask: How many people converted from X to Y? How likely were converts who were clergy in one religion likely to become clergy in another? How were conversions distributed over time? over space? But I won't know which questions can be investigated programmatically or what the data to answer them will look like until I've done substantially more research.

#### The idea: use YAML to model lives and events

With that research problem in mind, I've drawn up a list of specifications for what my data model should look like.

<ol type="1">
<li>
The data must be human-readable and -writable as research notes.
</li>
<li>
The data model must be able to grow organically as I do the research.
</li>
<li>
The data model must be able to hold large amounts of undigested text as notes.
</li>
<li>
The data must be portable to other formats, possible JSON or XML/TEI.
</li>
</ol>
My idea is to use [YAML](http://yaml.org/) as the format for the data. YAML is a "human friendly data serialization standard for all programming languages." YAML's two top priorities are "YAML is easily readable by humans" and "YAML data is portable between programming languages," which match my own priorities. I'm familiar with YAML from using [Jekyll](http://jekyllrb.com) for this blog and [another web project](http://jsr.fsu.edu). YAML also fits well into the principles I learned from [*Linux and the Unix Philosophy*](http://openlibrary.org/works/OL3494423W/Linux_and_the_Unix_philosophy), especially "store data in flat text files."

#### Example YAML model and Ruby script

I've created a working example with two YAML files and a Ruby script to output some of the data. I've shared the example as a [Gist on GitHub](https://gist.github.com/2580742).

The YAML file for Orestes Brownson is below, and there is another sample file for Charles Wharton in the [Gist](https://gist.github.com/2580742). You'll notice that at the outermost level of indentation, there are keys and values for basic biographical information, such as `born: 1803-09-16`. The most important part of the model is the list of conversions, which is a YAML array as signaled by the `-` character and indentation. The markup for the notes field (`notes: >`) lets that field contain as many paragraphs as necessary. Finally, the `source` array has one value (`@carey_orestes_2004`) which is the key to an entry in [my BibTeX database](https://github.com/lmullen/historybib), which I've added with [Vim's](http://www.vim.org/) autocomplete function.
<figure class="highlight">
<pre><code class="language-yaml" data-lang="yaml">&lt;span class="c1"># A model of a convert&#39;s life&lt;/span>
&lt;span class="nn">---&lt;/span>
&lt;span class="s">name-last&lt;/span>       &lt;span class="pi">:&lt;/span> &lt;span class="s">Brownson&lt;/span>
&lt;span class="s">name-first&lt;/span>      &lt;span class="pi">:&lt;/span> &lt;span class="s">Orestes Augustus&lt;/span>
&lt;span class="s">born&lt;/span>            &lt;span class="pi">:&lt;/span> &lt;span class="s">1803-09-16&lt;/span>
&lt;span class="s">died&lt;/span>            &lt;span class="pi">:&lt;/span> &lt;span class="s">1876-04-17&lt;/span>
&lt;span class="s">birth-religion&lt;/span>  &lt;span class="pi">:&lt;/span> &lt;span class="s">Congregationalism&lt;/span>

&lt;span class="s">conversions&lt;/span>     &lt;span class="pi">:&lt;/span>

&lt;span class="pi">-&lt;/span>   &lt;span class="s">origin-religion&lt;/span>         &lt;span class="pi">:&lt;/span> &lt;span class="s">Congregationalism&lt;/span>
    &lt;span class="s">destination-religion&lt;/span>    &lt;span class="pi">:&lt;/span> &lt;span class="s">Presbyterianism&lt;/span>
    &lt;span class="s">date&lt;/span>                    &lt;span class="pi">:&lt;/span> &lt;span class="s">1822&lt;/span>
    &lt;span class="s">ritual&lt;/span>                  &lt;span class="pi">:&lt;/span> &lt;span class="s">church membership&lt;/span>
    &lt;span class="s">citation&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s">ANB&lt;/span>
    &lt;span class="s">notes&lt;/span>                   &lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
      &lt;span class="no">Brownson&#39;s change to congregationalism was more denominational &lt;/span>
      &lt;span class="no">switching than a change in conscience.&lt;/span>

&lt;span class="pi">-&lt;/span>   &lt;span class="s">origin-religion&lt;/span>         &lt;span class="pi">:&lt;/span> &lt;span class="s">Presbyterianism&lt;/span>
    &lt;span class="s">destination-religion&lt;/span>    &lt;span class="pi">:&lt;/span> &lt;span class="s">Universalism&lt;/span>
    &lt;span class="s">date&lt;/span>                    &lt;span class="pi">:&lt;/span> &lt;span class="s">1826&lt;/span>
    &lt;span class="s">ritual&lt;/span>                  &lt;span class="pi">:&lt;/span> &lt;span class="s">ordination&lt;/span>
    &lt;span class="s">location&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="s">Jaffrey,&lt;/span>&lt;span class="nv"> &lt;/span>&lt;span class="s">New&lt;/span>&lt;span class="nv"> &lt;/span>&lt;span class="s">Hampshire"&lt;/span>
    &lt;span class="s">citation&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s">ANB&lt;/span>
    &lt;span class="s">notes&lt;/span>                   &lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
      &lt;span class="no">"He would later refer to his years in this fold as &#39;the most &lt;/span>
      &lt;span class="no">anti-Christian period of my life&#39;" (ANB).&lt;/span>

      &lt;span class="no">Brownson was editor of _The Gospel Advocate and Impartial &lt;/span>
      &lt;span class="no">Investigator_, a Universalist publication.&lt;/span>

&lt;span class="pi">-&lt;/span>   &lt;span class="s">origin-religion&lt;/span>         &lt;span class="pi">:&lt;/span> &lt;span class="s">Universalism&lt;/span>
    &lt;span class="s">destination-religion&lt;/span>    &lt;span class="pi">:&lt;/span> &lt;span class="s">Unitarianism&lt;/span>
    &lt;span class="s">ritual&lt;/span>                  &lt;span class="pi">:&lt;/span> &lt;span class="s">further research&lt;/span>
    &lt;span class="s">location&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="s">Walpole,&lt;/span>&lt;span class="nv"> &lt;/span>&lt;span class="s">New&lt;/span>&lt;span class="nv"> &lt;/span>&lt;span class="s">Hampshire"&lt;/span>
    &lt;span class="s">citation&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s">ANB&lt;/span>
    &lt;span class="s">notes&lt;/span>                   &lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
      &lt;span class="no">Brownson spent some time at Brook Farm, which prepared him for &lt;/span>
      &lt;span class="no">Transcendentalism&lt;/span>

&lt;span class="pi">-&lt;/span>   &lt;span class="s">origin-religion&lt;/span>         &lt;span class="pi">:&lt;/span> &lt;span class="s">Unitarianism and Transcendentalism&lt;/span>
    &lt;span class="s">destination-religion&lt;/span>    &lt;span class="pi">:&lt;/span> &lt;span class="s">Catholicism&lt;/span>
    &lt;span class="s">date&lt;/span>                    &lt;span class="pi">:&lt;/span> &lt;span class="s">1844-10-19&lt;/span>
    &lt;span class="s">ritual&lt;/span>                  &lt;span class="pi">:&lt;/span> &lt;span class="s">baptism&lt;/span>
    &lt;span class="s">citation&lt;/span>                &lt;span class="pi">:&lt;/span> &lt;span class="s">ANB&lt;/span>
    &lt;span class="s">notes&lt;/span>                   &lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
      &lt;span class="no">Brownson studied after his conversion with a Sulpician priest.&lt;/span>

&lt;span class="s">source&lt;/span>          &lt;span class="pi">:&lt;/span>
&lt;span class="pi">-&lt;/span>   &lt;span class="err">@&lt;/span>&lt;span class="s">carey_orestes_2004&lt;/span>
&lt;span class="pi">-&lt;/span>   &lt;span class="s">American National Biography&lt;/span>

&lt;span class="s">comments&lt;/span>        &lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
  &lt;span class="no">This is a minimal example of what a model of a convert might look &lt;/span>
  &lt;span class="no">like. The historical data is hastily gathered, so only the model is &lt;/span>
  &lt;span class="no">of interest here.&lt;/span>

  &lt;span class="no">N.B. I would like to replace the citations with BibTeX keys.&lt;/span>
&lt;span class="nn">...&lt;/span></code></pre>
</figure>
I had to prove to myself that I could get at the data programmatically, so I wrote the Ruby script below. It's just a proof-of-concept, and it's the first Ruby script I've written, so there are ugly parts. The script creates a class `Converts`, which loads an array of YAML files into a hash. The class has a few methods to display the names of the converts and a list of all the conversions. Doubtless there are more interesting things that can be done.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="c1">#!/usr/bin/env ruby&lt;/span>
&lt;span class="c1"># A proof-of-concept script that outputs some simple data from YAML &lt;/span>
&lt;span class="c1"># files modeling conversions&lt;/span>
&lt;span class="c1">#&lt;/span>
&lt;span class="c1"># Author:: Lincoln Mullen (lincoln@lincolnmullen.com)&lt;/span>

&lt;span class="nb">require&lt;/span> &lt;span class="s1">&#39;yaml&#39;&lt;/span> 

&lt;span class="c1"># This class loads data from YAML files, and outputs some values&lt;/span>

&lt;span class="k">class&lt;/span> &lt;span class="nc">Converts&lt;/span>

  &lt;span class="kp">attr_accessor&lt;/span> &lt;span class="ss">:files&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="ss">:data&lt;/span>

  &lt;span class="k">def&lt;/span> &lt;span class="nf">initialize&lt;/span> &lt;span class="p">(&lt;/span>&lt;span class="n">files&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="kp">nil&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">data&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="kp">nil&lt;/span>&lt;span class="p">)&lt;/span>
    &lt;span class="vi">@files&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">files&lt;/span>
    &lt;span class="vi">@data&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Hash&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">new&lt;/span>

    &lt;span class="k">if&lt;/span> &lt;span class="vi">@files&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">nil?&lt;/span>
      &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"You didn&#39;t pass me any files."&lt;/span>
    &lt;span class="k">elsif&lt;/span> &lt;span class="vi">@files&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">respond_to?&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"each"&lt;/span>&lt;span class="p">)&lt;/span>
      &lt;span class="c1"># walk through the array of files, creating a hash with the &lt;/span>
      &lt;span class="c1"># file name as the key and the file data as the value&lt;/span>
      &lt;span class="vi">@files&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">each&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">file&lt;/span>&lt;span class="o">|&lt;/span>
        &lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">file&lt;/span>&lt;span class="p">]&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">YAML&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">load_file&lt;/span>&lt;span class="p">(&lt;/span> &lt;span class="n">file&lt;/span> &lt;span class="p">)&lt;/span>
      &lt;span class="k">end&lt;/span>
    &lt;span class="k">end&lt;/span>
  &lt;span class="k">end&lt;/span>

  &lt;span class="c1"># output the hash we can see what we&#39;re working with&lt;/span>
  &lt;span class="k">def&lt;/span> &lt;span class="nf">display_raw&lt;/span>
    &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="se">\n&lt;/span>&lt;span class="s2">This is the raw data we have loaded:"&lt;/span>
    &lt;span class="nb">p&lt;/span>&lt;span class="p">(&lt;/span> &lt;span class="vi">@data&lt;/span> &lt;span class="p">)&lt;/span>
  &lt;span class="k">end&lt;/span>

  &lt;span class="c1"># walk through the hash, outputting the names of each person&lt;/span>
  &lt;span class="k">def&lt;/span> &lt;span class="nf">display_names&lt;/span>
    &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="se">\n&lt;/span>&lt;span class="s2">These people converted:"&lt;/span>
    &lt;span class="vi">@data&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">each_key&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="o">|&lt;/span>
      &lt;span class="nb">puts&lt;/span> &lt;span class="s2">" - &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="p">][&lt;/span>&lt;span class="s2">"name-first"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="p">][&lt;/span>&lt;span class="s2">"name-last"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
    &lt;span class="k">end&lt;/span>
  &lt;span class="k">end&lt;/span>

  &lt;span class="c1"># walk through the hash, outputting the names and conversions of &lt;/span>
  &lt;span class="c1"># each person&lt;/span>
  &lt;span class="k">def&lt;/span> &lt;span class="nf">display_conversions&lt;/span>
    &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="se">\n&lt;/span>&lt;span class="s2">We know about these conversions:"&lt;/span>
    &lt;span class="vi">@data&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">each_key&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="o">|&lt;/span>
      &lt;span class="nb">puts&lt;/span> &lt;span class="s2">" - &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="p">][&lt;/span>&lt;span class="s2">"name-first"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="p">][&lt;/span>&lt;span class="s2">"name-last"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">:"&lt;/span>
      &lt;span class="c1"># each person has an array of conversions (even if there is &lt;/span>
      &lt;span class="c1"># only one conversion)&lt;/span>
      &lt;span class="vi">@data&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="n">key&lt;/span>&lt;span class="p">][&lt;/span>&lt;span class="s2">"conversions"&lt;/span>&lt;span class="p">].&lt;/span>&lt;span class="nf">each&lt;/span> &lt;span class="p">{&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">conversion&lt;/span>&lt;span class="o">|&lt;/span>
        &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"     + From &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">conversion&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"origin-religion"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> to &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">conversion&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"destination-religion"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> by &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">conversion&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"ritual"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> in &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">conversion&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"date"&lt;/span>&lt;span class="p">]&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">."&lt;/span>
      &lt;span class="p">}&lt;/span>
    &lt;span class="k">end&lt;/span>
  &lt;span class="k">end&lt;/span>

&lt;span class="k">end&lt;/span>

&lt;span class="c1"># get sample data by loading every YAML file in the directory&lt;/span>
&lt;span class="nb">puts&lt;/span> &lt;span class="s2">"Let&#39;s load all the YAML files in this directory:"&lt;/span>
&lt;span class="nb">puts&lt;/span> &lt;span class="no">Dir&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">glob&lt;/span>&lt;span class="p">(&lt;/span> &lt;span class="s1">&#39;*.yml&#39;&lt;/span>&lt;span class="p">).&lt;/span>&lt;span class="nf">join&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s1">&#39;, &#39;&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="n">c&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Converts&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">new&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="no">Dir&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">glob&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s1">&#39;*.yml&#39;&lt;/span>&lt;span class="p">))&lt;/span>

&lt;span class="c1"># call the methods to display the names and conversions&lt;/span>
&lt;span class="n">c&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">display_names&lt;/span>
&lt;span class="n">c&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">display_conversions&lt;/span></code></pre>
</figure>
If you run the script on the sample YAML files, you get the output below. (Yes---the script does output in [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295). I only know one trick.)
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">Let&lt;span class="s1">&#39;s load all the YAML files in this directory:
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
     + From Unitarianism and Transcendentalism to Catholicism by baptism in 1844-10-19.&lt;/span></code></pre>
</figure>
#### What's next?

If this model works for modeling conversions, it should also work for modeling other kinds of historical events. For example, suppose a labor historian is researching strikes and kept a YAML file for each strike ...
<figure class="highlight">
<pre><code class="language-yaml" data-lang="yaml">&lt;span class="s">id:&lt;/span>&lt;span class="err">    &lt;/span>&lt;span class="s">Pullman strike&lt;/span>
&lt;span class="s">location&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">Pullman, Illinois&lt;/span>
&lt;span class="s">date&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">1894-05-11&lt;/span>
&lt;span class="s">corporations&lt;/span>&lt;span class="pi">:&lt;/span>
&lt;span class="s">-&lt;/span>&lt;span class="err">  &lt;/span>&lt;span class="s">Pullman Palace Car Company&lt;/span>
&lt;span class="s">unions&lt;/span>&lt;span class="pi">:&lt;/span>
&lt;span class="s">-&lt;/span>&lt;span class="err">  &lt;/span>&lt;span class="s">American Railway Union&lt;/span>
&lt;span class="s">accounts&lt;/span>&lt;span class="pi">:&lt;/span>

&lt;span class="s">-&lt;/span>&lt;span class="err">  &lt;/span>&lt;span class="s">name&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">John A. Doe&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="s">source&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">Chicago Tribune&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="s">description&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="pi">&gt;&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">malesuada, purus vel posuere aliquam, enim orci tempor quam, ac&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">rutrum arcu arcu nec leo."&lt;/span>

&lt;span class="err">-&lt;/span>&lt;span class="no"> name: Jane B. Doe&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">source: New York Times&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">description: &gt;&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">"Maecenas in velit nulla, pretium vestibulum lacus. Morbi dui purus,&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">imperdiet ac aliquam sodales, gravida ut diam. Vestibulum nec erat a&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">ligula tincidunt dignissim in et diam. Quisque tincidunt&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="no">pellentesque lorem, a scelerisque quam lacinia vitae."&lt;/span></code></pre>
</figure>
and another for each union ...
<figure class="highlight">
<pre><code class="language-yaml" data-lang="yaml">&lt;span class="s">union&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">American Railway Union&lt;/span>
&lt;span class="s">leaders:&lt;/span>&lt;span class="err">   &lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="s">-&lt;/span>&lt;span class="err">   &lt;/span>&lt;span class="s">name&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">Eugene V. Debs&lt;/span>
&lt;span class="err">      &lt;/span>&lt;span class="s">start&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">1893-06-20&lt;/span>
&lt;span class="err">      &lt;/span>&lt;span class="s">end&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">~&lt;/span>
&lt;span class="s">founded&lt;/span>&lt;span class="pi">:&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="s">date&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">1893-06-20&lt;/span>
&lt;span class="err">  &lt;/span>&lt;span class="s">place&lt;/span>&lt;span class="pi">:&lt;/span> &lt;span class="s">Chicago, Illinois&lt;/span></code></pre>
</figure>
I asked about this idea at [Digital Humanities Questions & Answers](http://digitalhumanities.org/answers/topic/using-yaml-to-model-historical-lives-or-events) and on Twitter. [Chad Black](http://parezcoydigo.wordpress.com/), [Ben Brumfield](http://manuscripttranscription.blogspot.com/), Ethan Gruber, [Caleb McDaniel](http://www.owlnet.rice.edu/~wcm1/), and Conal Tuohy offered valuable advice about how to think about this problem and what tools might be helpful later in the project. The [TEI](http://www.tei-c.org/) markup for an [event](http://www.tei-c.org/release/doc/tei-p5-doc/en/html/ref-event.html) and [person](http://www.tei-c.org/release/doc/tei-p5-doc/en/html/ref-person.html) (recommended by Conal) seems promising because it can accommodate types of data that I know I'll need, such as uncertain dates and name changes.

For now, though, I'm going to work with YAML, since I can get started on it right away and since I'm completely sure it will work as research notes and reasonably sure it can be munged into another format later.

I'll be glad for any advice about how to improve the data model or script and about what considerations I should think about to make sure the data is useful. If you have any ideas about what to do with the data once I've gathered it, I'll be glad for those too.
