---
date: '2013-02-06T21:52:00+00:00'
slug: 'parsing-tei-files-in-ruby-with-nokogiri'
title: Parsing TEI Files in Ruby with Nokogiri
---

Yesterday when I [wrote](http://lincolnmullen.com/blog/tei-boilerplate-for-getting-started-with-tei/) about experimenting with TEI Boilerplate, I mentioned that one of the impediments I'd found to using [TEI](http://www.tei-c.org/index.xml) was being able to do something with it immediately. [TEI Boilerplate](http://dcl.slis.indiana.edu/teibp/) lets you see a TEI file in your browser immediately. But I also wanted to experiment with analyzing a TEI file programmatically, so I found some sample documents and wrote an easy script in [Ruby](http://www.ruby-lang.org/en/) to serve as my own proof of concept.

For experimental purposes, I downloaded the Folger Shakespeare Library's [Digital Texts](http://www.folgerdigitaltexts.org/), a collection of Shakespeare's plays encoded in TEI. I choose these texts because they had each speaker marked up, as in this snippet from *[Macbeth](http://www.folgerdigitaltexts.org/?chapter=5&play=Mac&loc=p7)*. For my purposes, a text that marked up names, dates, or places would be more interesting, but the principles are identical.

<figure class="highlight">
<pre><code class="language-xml" data-lang="xml">&lt;span class="nt">&lt;speaker&lt;/span> &lt;span class="na">xml:id=&lt;/span>&lt;span class="s">"spk-1490"&lt;/span>&lt;span class="nt">&gt;&lt;/span>
&lt;span class="nt">&lt;w&lt;/span> &lt;span class="na">xml:id=&lt;/span>&lt;span class="s">"w0221430"&lt;/span>&lt;span class="nt">&gt;&lt;/span>SECOND&lt;span class="nt">&lt;/w&gt;&lt;/span>
&lt;span class="nt">&lt;c&lt;/span> &lt;span class="na">xml:id=&lt;/span>&lt;span class="s">"c0221440"&lt;/span>&lt;span class="nt">&gt;&lt;/span> &lt;span class="nt">&lt;/c&gt;&lt;/span>
&lt;span class="nt">&lt;w&lt;/span> &lt;span class="na">xml:id=&lt;/span>&lt;span class="s">"w0221450"&lt;/span>&lt;span class="nt">&gt;&lt;/span>WITCH&lt;span class="nt">&lt;/w&gt;&lt;/span>
&lt;span class="nt">&lt;/speaker&gt;&lt;/span></code></pre>
</figure>

I decided to write a Ruby script that identified all the speakers and counted the number of times each spoke. The heavy lifting is done by the Ruby library [Nokogiri](http://nokogiri.org/), "an HTML, XML, SAX, and Reader parser" able "to search documents via XPath or CSS3 selectors." I learned about Nokogiri from [this post](http://jasonheppler.org/2012/10/12/better-web-scraping-with-nokogiri.html) by Jason Heppler, from whom I've learned most of what I know about Ruby. (See his [Rubyist Historian](http://hepplerj.github.com/rubyist-historian/) for a primer.)

[Nokogiri](http://nokogiri.org/) is very powerful---more powerful than I know what to do with. In this script, it does all the analytical work in two lines. One line opens the TEI file, and Nokogiri then parses the document.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="n">doc&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Nokogiri&lt;/span>&lt;span class="o">::&lt;/span>&lt;span class="no">XML&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="nb">open&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">filename&lt;/span>&lt;span class="p">))&lt;/span></code></pre>
</figure>
The other line finds each `<speaker>` element and cleans up the name of the speaker.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="nb">name&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">speaker&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">content&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">gsub&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="sr">/\n/&lt;/span>&lt;span class="p">,&lt;/span>&lt;span class="s2">""&lt;/span>&lt;span class="p">)&lt;/span></code></pre>
</figure>
The rest of the script just sets up some scaffolding to keep track of the speakers and the number of their lines. Here is the whole thing.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="c1">#!/usr/bin/env ruby&lt;/span>
&lt;span class="c1"># encoding: utf-8&lt;/span>

&lt;span class="c1"># Name::          speakers.rb&lt;/span>
&lt;span class="c1"># Author::        Lincoln Mullen (mailto:lincoln@lincolnmullen.com)&lt;/span>
&lt;span class="c1"># Copyright::     Copyright (c) 2013 Lincoln Mullen &lt;/span>
&lt;span class="c1"># License::       MIT License | http://lmullen.mit-license.org/&lt;/span>

&lt;span class="c1"># This program finds all of the speakers in a TEI file and lists them &lt;/span>
&lt;span class="c1"># by the number of times that they speak.&lt;/span>
&lt;span class="c1"># Usage: ./speakers.rb my-tei-file.xml&lt;/span>

&lt;span class="nb">require&lt;/span> &lt;span class="s1">&#39;nokogiri&#39;&lt;/span>          &lt;span class="c1"># for xml parsing&lt;/span>
&lt;span class="nb">require&lt;/span> &lt;span class="s1">&#39;pp&#39;&lt;/span>                &lt;span class="c1"># for a nicer output&lt;/span>

&lt;span class="c1"># Get the file name to open&lt;/span>
&lt;span class="n">filename&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">ARGV&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="mi">0&lt;/span>&lt;span class="p">]&lt;/span>

&lt;span class="c1"># Open a hash to store our data&lt;/span>
&lt;span class="n">speakers&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Hash&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">new&lt;/span>

&lt;span class="k">begin&lt;/span> 
  &lt;span class="c1"># Open the file and parse it with Nokogiri&lt;/span>
  &lt;span class="n">doc&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Nokogiri&lt;/span>&lt;span class="o">::&lt;/span>&lt;span class="no">XML&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="nb">open&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">filename&lt;/span>&lt;span class="p">))&lt;/span>
  &lt;span class="c1"># Find each instance of a &lt;speaker&gt; tag&lt;/span>
  &lt;span class="n">doc&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">search&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s1">&#39;speaker&#39;&lt;/span>&lt;span class="p">).&lt;/span>&lt;span class="nf">each&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">speaker&lt;/span>&lt;span class="o">|&lt;/span>
    &lt;span class="c1"># Clean up the line breaks in the speaker&#39;s name&lt;/span>
    &lt;span class="nb">name&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">speaker&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">content&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">gsub&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="sr">/\n/&lt;/span>&lt;span class="p">,&lt;/span>&lt;span class="s2">""&lt;/span>&lt;span class="p">)&lt;/span>
    &lt;span class="k">if&lt;/span> &lt;span class="n">speakers&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">has_key?&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="nb">name&lt;/span>&lt;span class="p">)&lt;/span>
      &lt;span class="c1"># If the speaker is already in our hash then add 1 to the count &lt;/span>
      &lt;span class="c1"># of utterances&lt;/span>
      &lt;span class="n">speakers&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="nb">name&lt;/span>&lt;span class="p">]&lt;/span> &lt;span class="o">+=&lt;/span> &lt;span class="mi">1&lt;/span>
    &lt;span class="k">else&lt;/span>
      &lt;span class="c1"># If the speaker is not already in the hash then add the speaker &lt;/span>
      &lt;span class="c1"># with a count of 1&lt;/span>
      &lt;span class="n">speakers&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="nb">name&lt;/span>&lt;span class="p">]&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="mi">1&lt;/span>
    &lt;span class="k">end&lt;/span>
  &lt;span class="k">end&lt;/span>
&lt;span class="k">rescue&lt;/span> &lt;span class="no">Errno&lt;/span>&lt;span class="o">::&lt;/span>&lt;span class="no">ENOENT&lt;/span>
  &lt;span class="c1"># If the file we&#39;ve been passed doesn&#39;t exist, catch the error&lt;/span>
  &lt;span class="nb">puts&lt;/span> &lt;span class="s2">"That file does not exist."&lt;/span>
&lt;span class="k">end&lt;/span>

&lt;span class="c1"># Sort the hash of speakers by the number of times they speak, in &lt;/span>
&lt;span class="c1"># descending order, then print the output&lt;/span>
&lt;span class="n">pp&lt;/span> &lt;span class="n">speakers&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">sort_by&lt;/span> &lt;span class="p">{&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="nb">name&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">lines&lt;/span>&lt;span class="o">|&lt;/span> &lt;span class="n">lines&lt;/span> &lt;span class="p">}.&lt;/span>&lt;span class="nf">reverse&lt;/span></code></pre>
</figure>
The output for *Macbeth* looks like this. It's nothing too impressive, but it does show how Ruby and Nokogiri can be used to analyze TEI files.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="p">[[&lt;/span>&lt;span class="s2">"MACBETH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">145&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"LADY MACBETH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">59&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MACDUFF"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">59&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MALCOLM"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">40&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"ROSS"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">39&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"BANQUO"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">33&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"FIRST WITCH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">23&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"LENNOX"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">21&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"DOCTOR"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">20&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"LADY MACDUFF"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">19&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"DUNCAN"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">18&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SECOND WITCH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">15&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SON"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">14&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"ALL"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">13&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"THIRD WITCH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">13&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"GENTLEWOMAN"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">11&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SIWARD"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">11&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"FIRST MURDERER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">11&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MURDERER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">7&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SERVANT"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">6&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SECOND MURDERER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">6&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"THIRD MURDERER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">6&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MESSENGER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">6&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SEYTON"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">5&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MENTEITH"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">5&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"YOUNG SIWARD"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">4&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"PORTER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">4&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"ANGUS"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">4&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"OLD MAN"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">4&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"DONALBAIN"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"CAITHNESS"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"CAPTAIN"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"LORDS"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"LORD"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MURDERERS"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">3&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"HECATE"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">2&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"FLEANCE"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">2&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SECOND APPARITION"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">2&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"THIRD APPARITION"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">1&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"FIRST APPARITION"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">1&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"SOLDIER"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">1&lt;/span>&lt;span class="p">],&lt;/span>
 &lt;span class="p">[&lt;/span>&lt;span class="s2">"MACBETH AND LENNOX"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="mi">1&lt;/span>&lt;span class="p">]]&lt;/span></code></pre>
</figure>
