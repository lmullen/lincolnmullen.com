---
date: '2012-08-17T18:14:00+00:00'
slug: 'the-quantified-dissertation-using-bash-and-cron-to-track-word-counts'
title: 'The Quantified Dissertation: Using Bash and cron to Track Word Counts'
---

Here is a hack for those of you writing a dissertation (or book, or other long form project) in plain text. Since I am writing my dissertation in [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295) and [Pandoc](http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700), it's easy to find out the word count for a chapter from the command line:
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">&lt;span class="gp">$ &lt;/span>wc -w ch4.dissertation.Mullen.md</code></pre>
</figure>
I wanted to track my progress by regularly getting a word count for every chapter. To do that, I wrote a Bash script which runs `wc` and stores the data to a CSV file. Here is the script:
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">&lt;span class="c">#!/bin/bash          &lt;/span>
&lt;span class="nb">source&lt;/span> /Users/lmullen/.bash_profile

&lt;span class="nv">EPOCH&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>date +&lt;span class="s2">"%s"&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">DATE&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>date&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH1&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch1.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch1.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH2&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch2.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch2.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH3&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch3.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch3.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH4&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch4.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch4.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH5&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch5.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch5.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH6&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch6.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch6.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH7&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch7.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch7.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCH8&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/ch8.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/ch8.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCPRE&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/preface.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/preface.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCINTRO&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/introduction.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/introduction.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nv">WCCONCL&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="k">$(&lt;/span>wc -w /Users/lmullen/acad/dissertation/conclusion.dissertation.Mullen.md | sed &lt;span class="s1">&#39;s/\/Users\/lmullen\/acad\/dissertation\/conclusion.dissertation.Mullen.md//&#39;&lt;/span> | sed &lt;span class="s1">&#39;s/ //g&#39;&lt;/span>&lt;span class="k">)&lt;/span>
&lt;span class="nb">echo&lt;/span> &lt;span class="nv">$EPOCH&lt;/span>,&lt;span class="nv">$DATE&lt;/span>,&lt;span class="nv">$WCCH1&lt;/span>,&lt;span class="nv">$WCCH2&lt;/span>,&lt;span class="nv">$WCCH3&lt;/span>,&lt;span class="nv">$WCCH4&lt;/span>,&lt;span class="nv">$WCCH5&lt;/span>,&lt;span class="nv">$WCCH6&lt;/span>,&lt;span class="nv">$WCCH7&lt;/span>,&lt;span class="nv">$WCCH8&lt;/span>,&lt;span class="nv">$WCPRE&lt;/span>,&lt;span class="nv">$WCINTRO&lt;/span>,&lt;span class="nv">$WCCONCL&lt;/span> &gt;&gt; /Users/lmullen/acad/dissertation/wordcount-data.csv</code></pre>
</figure>
First, the script finds out the date in two formats (lines 4--5), the Unix "epoch" time (or number of seconds since midnight on January 1, 1970) and a human readable time. I thought I should store both since I didn't know which would be most useful later on. Second, the script counts the number of words in each chapter and in the front matter and back matter. It also strips away all of the output from that command except the word count itself (lines 6--16). Someone with better scripting skills than me could probably compress those eleven lines into one or two. Finally, the script appends a line with the data to a CSV file (line 17).

When the script is executed, it returns data like this (from the very first line in my data file):
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">EPOCH,DATETIME,CH1WC,CH2WC,CH3WC,CH4WC,CH5WC,CH6WC,CH7WC,CH8WC,PREWC,INTROWC,CONCLWC
1332736387,Mon Mar 26 00:33:07 EDT 2012,0,0,0,12412,0,0,0,0,0,0,0</code></pre>
</figure>
As you can see, by March 26 I had written 12,412 words for chapter 4, and not much else. Apparently I also wrote this script after midnight.

I run the script automatically each hour by adding this line to my crontab:
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">0   &lt;span class="k">*&lt;/span>   &lt;span class="k">*&lt;/span>   &lt;span class="k">*&lt;/span>   &lt;span class="k">*&lt;/span>   /Users/lmullen/acad/dissertation/wordcount.sh</code></pre>
</figure>
I haven't done much with the data yet, but the point is to collect it while I am writing. It's easy enough to open the CSV file in Excel and generate graphs of my progress. Eventually I may also analyze the data in [R](http://www.r-project.org/).

Gathering the data is one problem. Getting the word counts to increase at an acceptable pace is another problem altogether.
