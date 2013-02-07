---
layout: post
title: "Parsing TEI Files in Ruby with Nokogiri"
date: 2013-02-06 16:52
comments: true
categories: 
- digital humanities
tags:
- Ruby
- TEI
- Nokogiri
- digital humanities
---

Yesterday when I [wrote][] about experimenting with TEI Boilerplate, I
mentioned that one of the impediments I'd found to using [TEI][] was
being able to do something with it immediately. [TEI Boilerplate][] lets
you see a TEI file in your browser immediately. But I also wanted to
experiment with analyzing a TEI file programmatically, so I found some
sample documents and wrote an easy script in [Ruby][] to serve as my own
proof of concept.

<!--more-->

For experimental purposes, I downloaded the Folger Shakespeare Library's
[Digital Texts][], a collection of Shakespeare's plays encoded in TEI. I
choose these texts because they had each speaker marked up, as in this
snippet from *[Macbeth][]*. For my purposes, a text that marked up
names, dates, or places would be more interesting, but the principles 
are identical.

``` xml
<speaker xml:id="spk-1490">
<w xml:id="w0221430">SECOND</w>
<c xml:id="c0221440"> </c>
<w xml:id="w0221450">WITCH</w>
</speaker>
```

I decided to write a Ruby script that identified all the speakers and
counted the number of times each spoke. The heavy lifting is done by the
Ruby library [Nokogiri][], "an HTML, XML, SAX, and Reader parser" able
"to search documents via XPath or CSS3 selectors." I learned about
Nokogiri from [this post][] by Jason Heppler, from whom I've learned
most of what I know about Ruby. (See his [Rubyist Historian][] for a
primer.)

[Nokogiri][] is very powerful---more powerful than I know what to do 
with. In this script, it does all the analytical work in two lines. 
One line opens the TEI file, and Nokogiri then parses the document.

``` ruby
doc = Nokogiri::XML(open(filename))
```

The other line finds each `<speaker>` element and cleans up the name 
of the speaker.

``` ruby
name = speaker.content.gsub(/\n/,"")
``` 

The rest of the script just sets up some scaffolding to keep track of 
the speakers and the number of their lines. Here is the whole thing.

{% include_code speakers.rb %}

The output for *Macbeth* looks like this. It's nothing too impressive, 
but it does show how Ruby and Nokogiri can be used to analyze TEI files.

``` text
[["MACBETH", 145],
 ["LADY MACBETH", 59],
 ["MACDUFF", 59],
 ["MALCOLM", 40],
 ["ROSS", 39],
 ["BANQUO", 33],
 ["FIRST WITCH", 23],
 ["LENNOX", 21],
 ["DOCTOR", 20],
 ["LADY MACDUFF", 19],
 ["DUNCAN", 18],
 ["SECOND WITCH", 15],
 ["SON", 14],
 ["ALL", 13],
 ["THIRD WITCH", 13],
 ["GENTLEWOMAN", 11],
 ["SIWARD", 11],
 ["FIRST MURDERER", 11],
 ["MURDERER", 7],
 ["SERVANT", 6],
 ["SECOND MURDERER", 6],
 ["THIRD MURDERER", 6],
 ["MESSENGER", 6],
 ["SEYTON", 5],
 ["MENTEITH", 5],
 ["YOUNG SIWARD", 4],
 ["PORTER", 4],
 ["ANGUS", 4],
 ["OLD MAN", 4],
 ["DONALBAIN", 3],
 ["CAITHNESS", 3],
 ["CAPTAIN", 3],
 ["LORDS", 3],
 ["LORD", 3],
 ["MURDERERS", 3],
 ["HECATE", 2],
 ["FLEANCE", 2],
 ["SECOND APPARITION", 2],
 ["THIRD APPARITION", 1],
 ["FIRST APPARITION", 1],
 ["SOLDIER", 1],
 ["MACBETH AND LENNOX", 1]]
```

  [Digital Texts]: http://www.folgerdigitaltexts.org/
  [Macbeth]: http://www.folgerdigitaltexts.org/?chapter=5&play=Mac&loc=p7
  [Nokogiri]: http://nokogiri.org/
  [Ruby]: http://www.ruby-lang.org/en/
  [Rubyist Historian]: http://hepplerj.github.com/rubyist-historian/
  [TEI Boilerplate]: http://dcl.slis.indiana.edu/teibp/
  [TEI]: http://www.tei-c.org/index.xml
  [this post]: http://jasonheppler.org/2012/10/12/better-web-scraping-with-nokogiri.html
  [wrote]: http://lincolnmullen.com/blog/tei-boilerplate-for-getting-started-with-tei/
