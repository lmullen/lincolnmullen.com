---
date: '2013-05-16T23:40:00+00:00'
slug: 'using-r-to-chart-the-historical-demography-of-american-judaism'
title: Using R to Chart the Historical Demography of American Judaism
---

The historical demography of American religion is in a sorry state. I've been intending to write about the reasons and possible fixes for this historiographical quandary for some time. My draft blog post has grown to epic length because there are really several topics to address: (1) what attempts have been made at understanding the historical demographics of American religion, (2) what sources exist for recovering these demographics (spoiler: it's pretty grim, but not as grim as you might think), (3) what historical and analytical techniques can we use to fill in the gaps, and (4) what digital humanities tools might help manage and analyze the data once recovered. Parts one through three must wait for more reading, but I want to offer some preliminary thoughts on digital humanities tools.

So what tools are available to historians who want to study historical demography?

As I see it, there are three main features that we need. One is a simple, future-proof way to store data. The second is a way to analyze the data mathematically. The third is a way to present the data visually.

Why not use Excel? For the love of your fellow scholars, *don't use Excel*. If you need a reason, see this example of [economists who made a mistake in Excel](http://www.nextnewdeal.net/rortybomb/researchers-finally-replicated-reinhart-rogoff-and-there-are-serious-problems) with the results that "that one of the core empirical points providing the intellectual foundation for the global move to austerity in the early 2010s was based on someone accidentally not updating a row formula in Excel" (via [Kieran Healy](http://kieranhealy.org/blog/archives/2013/04/17/new-tools-for-reproducible-research/)).

Thankfully scholars in the hard sciences and social sciences who deal with data all the time have developed very sophisticated tools, which historians can readily repurpose. The tool that I've been introduced to is the [statistical programming language R](http://www.r-project.org/). Unlike its competitors such as SPSS, R is free and open-source. It's a language that humanists are already using for interesting purposes: Matt Jockers [suggests](http://www.matthewjockers.net/2013/01/03/advice-for-dh-newbies/) that you learn R if you're interested in the digital humanities, and Ben Schmidt is doing great work with R at his blog [Sapping Attention](http://sappingattention.blogspot.com/). There are also many helpful resources for learning about R, such as Anthony Damico's [R Twotorials](http://www.twotorials.com/), which I [mentioned](http://chronicle.com/blogs/profhacker/learn-r-with-twotorials/45843) at ProfHacker, and [Try R](http://tryr.codeschool.com/), an interactive course.

R gives us the desired ability to analyze and plot data. For the data format, we can use the ubiquitous [comma-separated values](http://en.wikipedia.org/wiki/Comma-separated_values) (CSV) format, that is, a plain-text version of a spreadsheet which can be opened in Excel or Google Docs but also accessed via R or any other programming language.

As a first step towards using R as a historian, I want to present some basic data and plots. The goal is to demonstrate what R can do for the historian, rather than explain how to use R. (Perhaps I'll write up an introduction for how to use R as a historian later).

The data are American Jewish population estimates taken from the appendix to Jonathan D. Sarna, *[American Judaism: A History](http://books.google.com/books?id=Ujw_IawuIIgC&lpg=PA375&vq=appendix&dq=jonathan%20sarna%20american%20judaism&pg=PA375#v=onepage&q&f=false)* (New Haven: Yale University Press, 2004), 375--76. Because this data has some gaps, and because it includes upper and lower bound estimates, the data is more interesting than a trivial example.

I've entered this data into a [CSV file](https://github.com/lmullen/demographics-religion/blob/master/judaism/sarna.appendix.csv), part of which looks like this:
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">year,estimate-low,estimate-high,percentage-population-low,percentage-population-high
1660,50,50,,
1700,200,300,,
1776,1000,2500,0.04,0.1
1790,1300,3000,0.03,0.08
1800,2500,2500,0.04,0.04</code></pre>
</figure>
Next I've written [an R script](https://github.com/lmullen/demographics-religion/blob/master/judaism/judaism.overview.R) which will do some analysis and plotting.

Getting the data into R is easy:
<figure class="highlight">
<pre><code class="language-r" data-lang="r">&lt;span class="n">data&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="n">read.csv&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"sarna.appendix.csv"&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">comment.char&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"#"&lt;/span>&lt;span class="p">)&lt;/span></code></pre>
</figure>
Now we can use R to do some basic calculations. In this case, the data does not include a mid-point between the high and low estimates, so we need to calculate that.
<figure class="highlight">
<pre><code class="language-r" data-lang="r">&lt;span class="c1"># Calculate mean/midpoint and the height of the range bars
# -------------------------------------------------------------------
&lt;/span>&lt;span class="n">midpoint.population&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">estimate.high&lt;/span> &lt;span class="o">+&lt;/span> &lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">estimate.low&lt;/span>&lt;span class="p">)&lt;/span>&lt;span class="o">/&lt;/span>&lt;span class="m">2&lt;/span>
&lt;span class="n">error.population&lt;/span>    &lt;span class="o">&lt;-&lt;/span> &lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">estimate.high&lt;/span> &lt;span class="o">-&lt;/span> &lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">estimate.low&lt;/span>&lt;span class="p">)&lt;/span>&lt;span class="o">/&lt;/span>&lt;span class="m">2&lt;/span>
&lt;span class="n">midpoint.percentage&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">percentage.population.high&lt;/span> &lt;span class="o">+&lt;/span>
                        &lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">percentage.population.low&lt;/span>&lt;span class="p">)&lt;/span>&lt;span class="o">/&lt;/span>&lt;span class="m">2&lt;/span>
&lt;span class="n">error.percentage&lt;/span>    &lt;span class="o">&lt;-&lt;/span> &lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">percentage.population.high&lt;/span> &lt;span class="o">-&lt;/span>
                        &lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">percentage.population.low&lt;/span>&lt;span class="p">)&lt;/span>&lt;span class="o">/&lt;/span>&lt;span class="m">2&lt;/span></code></pre>
</figure>
Now we can plot the population estimates.
<figure class="highlight">
<pre><code class="language-r" data-lang="r">&lt;span class="n">png&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">filename&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"jewish-population.png"&lt;/span>&lt;span class="p">,&lt;/span>
&lt;span class="n">width&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="m">2000&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">height&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="m">1200&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">res&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="m">300&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="n">ggplot&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="o">+&lt;/span>
&lt;span class="n">geom_line&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">aes&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">year&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">midpoint.population&lt;/span>&lt;span class="p">))&lt;/span> &lt;span class="o">+&lt;/span>
&lt;span class="n">geom_pointrange&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">aes&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">x&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">data&lt;/span>&lt;span class="o">$&lt;/span>&lt;span class="n">year&lt;/span>&lt;span class="p">,&lt;/span>
                    &lt;span class="n">y&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">midpoint.population&lt;/span>&lt;span class="p">,&lt;/span>
                    &lt;span class="n">ymin&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">midpoint.population&lt;/span> &lt;span class="o">-&lt;/span> &lt;span class="n">error.population&lt;/span>&lt;span class="p">,&lt;/span>
                    &lt;span class="n">ymax&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">midpoint.population&lt;/span> &lt;span class="o">+&lt;/span> &lt;span class="n">error.population&lt;/span>&lt;span class="p">))&lt;/span> &lt;span class="o">+&lt;/span>
&lt;span class="n">labs&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">title&lt;/span>&lt;span class="o">=&lt;/span>&lt;span class="s2">"American Jewish Population Estimates, 1660-2000"&lt;/span>&lt;span class="p">,&lt;/span>
     &lt;span class="n">x&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"Year"&lt;/span>&lt;span class="p">,&lt;/span>
     &lt;span class="n">y&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"Population (millions)"&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="o">+&lt;/span>
&lt;span class="n">scale_y_continuous&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">breaks&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">seq&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="m">0&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="m">6e+06&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="m">1e+06&lt;/span>&lt;span class="p">),&lt;/span>
                   &lt;span class="n">labels&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">seq&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="m">0&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="m">6e+06&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="m">1e+06&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="o">/&lt;/span> &lt;span class="m">1e+06&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="n">attribution&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"Data: Jonathan Sarna, American Judaism"&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="n">dev.off&lt;/span>&lt;span class="p">()&lt;/span></code></pre>
</figure>
Here is the result:

<img class="center" src="//files.lincolnmullen.com/downloads/post/jewish-population.png" />

Notice that we've tried to make the plot indicate various features of the data. There are bars to indicate high and low estimates, which is important because the raw number of Jews is either growing or shrinking, depending on whether you use the high or low estimates. There are also dots to indicate which years we have data for, so we know where the line reflects an actual data point, and where it is an interpolation. And we've made the plot wider than it is tall, in order to avoid giving a false impression of the scale of the changes.

We can also plot the change in Jewish population as a percentage of the population as a whole. (If we wished, we could layer these plots to compare them.)

<img class="center" src="//files.lincolnmullen.com/downloads/post/jewish-population.percentage.png" />

The very small number of Jews before 1776 and their rapid growth after 1890 mean that the less drastic changes in the nineteenth century are obscured. Using the same data, we can take snapshots of the earlier and later nineteenth century.

<img class="center" src="//files.lincolnmullen.com/downloads/post/jewish-population.early-19c.png" />

<img class="center" src="//files.lincolnmullen.com/downloads/post/jewish-population.late-19c.png" />

This example is very basic---trivial, even. So why is this mode of working an improvement over, say, Excel?

First, this mode of working is transparent. My Brandeis colleague [Yoni Appelbaum](https://twitter.com/yappelbaum) makes a point of linking to primary and secondary sources in his [columns](http://www.theatlantic.com/yoni-appelbaum/) for *The Atlantic*, so that readers can see the evidence for themselves. In the same way, putting both the data and the R script [on GitHub](https://github.com/lmullen/demographics-religion) lets readers have access to the evidence. Much of the problem with existing demographic evidence is that the original assumptions and calculations are now irrecoverable and therefore unreliable (a point I'll demonstrate in a later post).

Second, this mode of working makes the work shareable. Much of the potential value from demographic work comes from the ability to combine and compare data from different sources. In this case I've taken the liberty of borrowing data from my dissertation adviser's book. It has always been possible to share data in print, of course. Digitizing data is laborious in its own way, but once data is digitized it can be combined and remixed easily.

Third, this mode of working is iterative. R makes it very easy to change the parameters of charts in order to investigate different aspects of whatever you are studying. Earlier, for example, we made charts that looked at the nineteenth century in more detail. R lets you be explicit about how you are iterating your investigation.

I'll be glad to hear more refined thoughts from anyone who knows more about R, the demography of American religion, or both.
