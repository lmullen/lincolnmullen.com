---
layout: post
title: "Using R to Chart the Historical Demography of American Judaism"
date: 2013-05-16 19:40
comments: true
sidebar: collapse
categories: 
- research
- digital humanities
- demographics of religion
- history
tags:
- research
- digital humanities
- demographics of religion
- history
- R language
- statistical programming
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

{% highlight bash %}
year,estimate-low,estimate-high,percentage-population-low,percentage-population-high
1660,50,50,,
1700,200,300,,
1776,1000,2500,0.04,0.1
1790,1300,3000,0.03,0.08
1800,2500,2500,0.04,0.04
{% endhighlight %}


Next I've written [an R script](https://github.com/lmullen/demographics-religion/blob/master/judaism/judaism.overview.R) which will do some analysis and plotting.

Getting the data into R is easy:

{% highlight r %}
data <- read.csv("sarna.appendix.csv", comment.char = "#")
{% endhighlight %}

Now we can use R to do some basic calculations. In this case, the data does not include a mid-point between the high and low estimates, so we need to calculate that. 

{% highlight r %}
# Calculate mean/midpoint and the height of the range bars
# -------------------------------------------------------------------
midpoint.population <- (data$estimate.high + data$estimate.low)/2
error.population    <- (data$estimate.high - data$estimate.low)/2
midpoint.percentage <- (data$percentage.population.high +
                        data$percentage.population.low)/2
error.percentage    <- (data$percentage.population.high -
                        data$percentage.population.low)/2
{% endhighlight %}

Now we can plot the population estimates. 

{% highlight r %}
png(filename = "jewish-population.png",
width=2000, height=1200, res=300)
ggplot(data) +
geom_line(aes(data$year, midpoint.population)) +
geom_pointrange(aes(x = data$year,
                    y = midpoint.population,
                    ymin = midpoint.population - error.population,
                    ymax = midpoint.population + error.population)) +
labs(title="American Jewish Population Estimates, 1660-2000",
     x = "Year",
     y = "Population (millions)") +
scale_y_continuous(breaks = seq(0, 6e+06, 1e+06),
                   labels = seq(0, 6e+06, 1e+06) / 1e+06)
attribution("Data: Jonathan Sarna, American Judaism")
dev.off()
{% endhighlight %}

Here is the result:

{% img center /downloads/img/jewish-population.png %}

Notice that we've tried to make the plot indicate various features of the data. There are bars to indicate high and low estimates, which is important because the raw number of Jews is either growing or shrinking, depending on whether you use the high or low estimates. There are also dots to indicate which years we have data for, so we know where the line reflects an actual data point, and where it is an interpolation. And we've made the plot wider than it is tall, in order to avoid giving a false impression of the scale of the changes.

We can also plot the change in Jewish population as a percentage of the population as a whole. (If we wished, we could layer these plots to compare them.)

{% img center /downloads/img/jewish-population.percentage.png %}

The very small number of Jews before 1776 and their rapid growth after 1890 mean that the less drastic changes in the nineteenth century are obscured. Using the same data, we can take snapshots of the earlier and later nineteenth century.

{% img center /downloads/img/jewish-population.early-19c.png %}

{% img center /downloads/img/jewish-population.late-19c.png %}

This example is very basic---trivial, even. So why is this mode of working an improvement over, say, Excel?

First, this mode of working is transparent. My Brandeis colleague [Yoni Appelbaum](https://twitter.com/yappelbaum) makes a point of linking to primary and secondary sources in his [columns](http://www.theatlantic.com/yoni-appelbaum/) for *The Atlantic*, so that readers can see the evidence for themselves. In the same way, putting both the data and the R script [on GitHub](https://github.com/lmullen/demographics-religion) lets readers have access to the evidence. Much of the problem with existing demographic evidence is that the original assumptions and calculations are now irrecoverable and therefore unreliable (a point I'll demonstrate in a later post).

Second, this mode of working makes the work shareable. Much of the potential value from demographic work comes from the ability to combine and compare data from different sources. In this case I've taken the liberty of borrowing data from my dissertation adviser's book. It has always been possible to share data in print, of course. Digitizing data is laborious in its own way, but once data is digitized it can be combined and remixed easily. 

Third, this mode of working is iterative. R makes it very easy to change the parameters of charts in order to investigate different aspects of whatever you are studying. Earlier, for example, we made charts that looked at the nineteenth century in more detail. R lets you be explicit about how you are iterating your investigation. 

I'll be glad to hear more refined thoughts from anyone who knows more about R, the demography of American religion, or both.
