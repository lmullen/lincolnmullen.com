---
categories:
- digital history
date: '2016-03-04T05:00:00+00:00'
slug: 'a-course-in-computational-methods-and-nineteenth-century-religious-data'
title: 'A Course in Computational Methods and Nineteenth-Century Religious Data'
---

This semester I am teaching a graduate course on [Data and Visualization in Digital History](http://lincolnmullen.com/courses/data-dh.2016/). The aim of this course is to teach students how to do the kind of data analysis and visualization that they are likely to do for a dissertation chapter or a journal article. In my way of working, that means the first part of the semester is an introduction to scripting in R, focusing on the grammar of graphics with [ggplot2](http://ggplot2.org/) and the grammar of data manipulation with [dplyr](https://github.com/hadley/dplyr) and [tidyr](https://github.com/hadley/tidyr). Then the second part of the course is aimed at introducing specific kinds of analysis in the context of historical work. My aim is that this course will be the first in a two course sequence, where the second course (colloquially known as Clio 3) will have more programming in R (as opposed to scripting), will have more \*nix-craft, will tackle a more advanced historical problem, will possibly cover more machine learning, and will end up creating interactive analyses in Shiny.

There are a few things about the Data and Visualization course that I think are worth mentioning.

First, I've been creating [worksheets for historical data analysis](https://dh-r.lincolnmullen.com/worksheets.html) each week. These worksheets tend to demonstrate some technique, then ask students to build up an analysis step by step. The questions within each worksheet range in difficulty from the rote and mechanical to the very difficult. While for now these worksheets are aimed at this class in particular, I intend over time to write worksheets like these for any topic in R I end up teaching. I'm rather pleased with these worksheets as a method of teaching data analysis by example.[<sup>1</sup>](#fn1){\#fnref1.footnoteRef} If I'm judging my students' initial reactions correctly, they are also finding them helpful, if rather difficult at times.

<!--more-->
The second point is that I decided to make the theme of the course about data on nineteenth-century U.S. religion. This is an area in which I'm just beginning to do work (as described [here](http://usreligion.blogspot.com/2016/01/Where-the-Problem-with-Historical-Data-about-U-S-Religion-Really-Lies.html) and [here](http://usreligion.blogspot.com/2016/02/a-very-preliminary-taxonomy-of-sources.html) and [here](http://lincolnmullen.com/projects/asch-2015/)). I really wrestled with whether this was the best thing to do, or whether I should make students pick their own project in their areas of research, or deliberately spread my examples around U.S. history and beyond.

But here's why I chose to do it this way. It seems to me that when digital history methods are taught, we too often focus on the second to last part of the methods (i.e., visualization) and miss everything that has gone before. Making a visualization depends on wrangling the data into the correct form, and before that on actually getting the data from somewhere. Part of the reason we don't teach those other steps is that we don't have as good a theory of how those earlier steps should be done. It makes sense to talk about principles of data visualization and maybe of data manipulation. But principles of data cleaning and gathering? They are rules of thumb at best, especially once the data is historical and not contemporary. Sometimes most crucial method in my digital history toolbox is sitting on the couch watching Netflix and transcribing data from books in the Hathi Trust, but that would make for an awfully boring workshop. If scholars who work with data are [like carpenters](http://www.mimno.org/articles/carpentry/), then historians who work with data are also like lumberjacks: we have to go into the forest and fell our own trees if we want to have some data to work with. On the other side, how does one make an argument from historical data? It's not an easy question. So I concluded that whatever the methods course was about, it had to be about some discrete topic which we could get deep into as a group, learning computational methods alongside readings in the historiography of that field so we know what the interesting questions are. Among the few things I know, I'd like to think that I can guide students through the full complexities of nineteenth century U.S. religious data. If I do this course again I might try to change the topic: U.S. legal history is particularly appealing at the moment, and perhaps the history of slavery in the United States. Still, the course will have to be about something.[<sup>2</sup>](#fn2){\#fnref2.footnoteRef}

I'm really trying to make this a digital *history* class, as opposed to a digital humanities class. This is in the spirit, I think, of Matt Jocker's *Text Analysis in R for Students of Literature*, which is aimed more at disciplinary than interdisciplinary concerns. And so I have included a week on the history of nineteenth-century data and visualization, another on the historiography of the topic, as well as weeks on gathering data and interpreting it historically.

Third, for a rather technical course, I've gotten the technical requirements down to a bare minimum. I listed a few additional requirements in the software section, but students could really get by with a web browser. That's because we have a server (thanks to the generosity of RRCHNM) where the students are using R via [RStudio Server](https://www.rstudio.com/products/rstudio/#Server). Knowing how to install a development environment is an important skill, but it's not the historical work itself. Since I don't like [wasting my students' time](http://lincolnmullen.com/blog/dont-waste-students-time-installing-software/), I've decided to use RStudio Server and then help students install a development environment on their own computer if they want.
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
Of course, the idea is not original to me: I learned of it from Ben Schmidt and <a href="http://workbook.craftingdigitalhistory.ca/">Shawn Graham</a> first.<a href="#fnref1">↩</a>
</li>
<li id="fn2">
Students do have complete leeway, though, to do their final projects on some other topic which is more suitable for them.<a href="#fnref2">↩</a>
</li>
</section>
