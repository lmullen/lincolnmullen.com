---
categories:
- digital history
date: '2016-03-05T05:00:00+00:00'
slug: 'makefiles-for-writing-data-analysis-ocr-and-converting-shapefiles'
title: 'Makefiles for Writing, Data Analysis, OCR, and Converting Shapefiles'
aliases:
- /blog/makefiles-for-ocr-and-converting-shapefiles/
---

I love Makefiles more than I ought to.[<sup>1</sup>](#fn1) If you haven't come across [GNU Make](https://www.gnu.org/software/make/) before, then Mike Bostock's "[Why Use Make](https://bost.ocks.org/mike/make/)" is a good introduction, as is the [section on Make](http://plain-text.co/make.html) at Kieran Healy's "[Plain Person's Guide to Plain Text Social Science](http://plain-text.co/)." I like Make for several reasons. It lets you specify how your final products (like a website or a PDF document) are related to inputs, and that discipline is invaluable for producing reproducible research and for structuring your project sensibly. For lots of tasks it provides free parallelization and rebuilds only what is absolutely necessary. Since my projects fit into several different genres, once I have created a Makefile for the genre, it is trivial to adapt it to different projects. Whether it is an article or a book manuscript, a data analysis project, a website, my CV, or some kind of file conversion process, all that I need to remember how to do is type `make` to build the project, `make deploy` to put it on the web, and `make clean` to start over.

<!--more-->
I often get asked how to do certain tasks related to digital humanities. Several of these queries came all at once recently, so it made sense to create some general purpose Makefiles that solve certain classes of problems.[<sup>2</sup>](#fn2){\#fnref2.footnoteRef} Below I point you to Makefiles for writing projects, for data analysis notebooks and other websites using R Markdown, for OCRing PDFs, and for converting shapefiles.

If you look over all these Makefiles you'll see that there are probably only five or six elements which are repeated over and over. It doesn't take many lines in a Makefile to get powerful results, yet I run the command `make` literally dozens of times per day in widely varying projects. GNU Make is a little peculiar, but picking it up has probably had the best return on my time for any technology I've learned.

### Writing projects

I wrote about using [Make with pandoc](http://lincolnmullen.com/blog/make-and-pandoc/) several years ago. That is still the basic template that I use for all my writing projects. Here is [an example](https://github.com/lmullen/dissertation/blob/master/Makefile) where the Makefile is basically the same even though the writing project is on a rather larger scale.

### Data analysis notebooks and websites with R Markdown

Almost all of my projects that involve some kind of data analysis end up having a research notebook (like a lab notebook), where I keep track of my work and present preliminary results along the way. These all use [R Markdown](http://rmarkdown.rstudio.com/). As an example, one ongoing project is tracking the migration of law in the nineteenth-century United States. You can see the project's [notebook here](http://lmullen.github.io/civil-procedure-codes/), along with the [Makefile for the project](https://github.com/lmullen/civil-procedure-codes/blob/master/Makefile). This is such a common pattern for my projects that I have abstracted it out into my [Simple R Markdown Notebook](https://lmullen.github.io/rmd-notebook/) ([repository here](https://github.com/lmullen/rmd-notebook)). This is intentionally as bare bones as possible, but the Makefile creates all the pages along with an home page.

I also have several standalone websites which are created with R Markdown. A good recent example is [the Makefile](https://github.com/lmullen/worksheets/blob/master/Makefile) for my worksheets for learning R, which creates [this site](https://dh-r.lincolnmullen.com/worksheets.html). Besides the tasks to actually create the website, check out the task to [deploy it using rsync](https://github.com/lmullen/worksheets/blob/master/Makefile#L14). (Here's an even simpler example of a [deploy task using scp](https://github.com/lmullen/CV/blob/master/Makefile#L14).)

### OCR

A common problem is to have a batch of PDFs for which one wants to generate OCR in plain text, either one file for each page or one file for each PDF. I've had this kind of task in several of my projects, but I recently abstracted it out into its own repository. The basic idea is that you drop a bunch of PDFs into the same directory as the Makefile, then generate the OCR text files in parallel. There is a fuller explanation in the [repository's README](https://github.com/lmullen/ocr-makefile). I even tell you how to install the dependencies. From time to time I've wanted to OCR so many files that I have thrown this Makefile up on a server, which works beautifully.

### Converting and simplifying shapefiles

[NHGIS](https://www.nhgis.org/) is awesome for providing high quality shapefiles keyed to each U.S. Census. But these are too high-resolution and in the wrong projection to upload to [CartoDB](https://cartodb.com/), which is often useful for teaching. [This is a Makefile](https://github.com/lmullen/nghis-simplifier) which will convert those NHGIS shapefiles into a format that you can use in CartoDB. It wouldn't take much tweaking to adapt this to other tasks involving the conversion of shapefiles.
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
One of my favorite parlor tricks is "deleting" my book manuscript in front of an audience, typing <code>make</code> and letting them watch it be recreated.<a href="#fnref1">↩</a>
</li>
<li id="fn2">
No, if you write to me I won't create a custom Makefile for your project. These were things I had already written for my own purposes and it was trivial to generalize them for people I have an obligation to.<a href="#fnref2">↩</a>
</li>
</section>
