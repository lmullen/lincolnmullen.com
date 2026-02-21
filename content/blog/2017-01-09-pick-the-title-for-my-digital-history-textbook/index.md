---
title: Pick the title for my digital history textbook
date: "2017-01-09"
slug: 'pick-the-title-for-my-digital-history-textbook'
description: "A computational history textbook restarted with a clearer vision: historical thinking first, R techniques second, tidy data throughout—plus a request for help choosing the title."
tags:
- digital history
- R
- teaching
draft: false
---

In my [first semester](https://files.lincolnmullen.com/syllabi/clio3.2014.pdf) teaching one of my department's graduate methods courses in digital history, I realized that there was not a lot good material for teaching computer programming and data analysis in R for historians. So I started writing up a series of tutorials for my students, which they said were helpful. It seemed like those materials could be the nucleus of a textbook, so I started [writing one](https://dh-r.lincolnmullen.com/) with the title *Digital History Methods in R*.

It was too soon to start writing, though. Besides needing to spend my time on more pressing projects, I didn't really have a clear conception of how to teach the material. And in the past few years, the landscape for teaching computational history has been transformed. There are many more books available, some specifically aimed at humanists, such as Graham, Milligan, and Weingart's [*Exploring Big Historical Data*](http://www.themacroscope.org/2.0/) and Arnold and Tilton's [*Humanities Data in R*](http://humanitiesdata.org/), and others aimed at teaching a modern version of R, such as Hadley Wickham's [*Advanced R*](http://adv-r.had.co.nz/) and [*R for Data Science*](http://r4ds.had.co.nz/). The "[tidyverse](https://blog.rstudio.org/2016/09/15/tidyverse-1-0-0/)" of R packages has made a consistent approach to data analysis possible, and the set of packages for text analysis in R is now much better. [R markdown](http://rmarkdown.rstudio.com/) and [bookdown](https://bookdown.org/) have made writing a technical book about R much easier, and [Shiny](http://shiny.rstudio.com/) has made it much easier to demonstrate concepts interactively.

After teaching these courses [a](http://www.dhsi.org/archive.php?year=2016) [few](https://files.lincolnmullen.com/syllabi/data-dh.2016.pdf) [times](https://files.lincolnmullen.com/syllabi/text-analysis.2016.pdf), I have a clearer conception of what the textbook needs to accomplish and how I want it to look.

<!--more-->
-   It will first and foremost be a book about historical thinking. The most difficult thing to teach is not how to left join two tables, but how to ask historical questions and answer them computationally. So the first part of the book will take readers through a set of example projects, beginning with historical sources and questions, applying data analysis techniques, and ending up with historical insight.
-   The second part of the book will contain a series of chapters explaining various techniques. Some will be very practical, such as dealing with dates or strings, and others will be more conceptual, such as explaining networks.
-   The book will take a "tidy data" approach whenever possible, meaning everywhere except text analysis. And when necessary, it will explain concepts with embedded Shiny applications. Except maybe a bit of shell scripting, the only language or tool explained will be R.
-   The scope of the book will be computational history. So, it will be unapologetically a book for historians rather than humanists, though I hope there will be material of value to, say, literary scholars. And it will only take up that one narrow slice of digital history/humanities.

I'm not entirely sure about the title, though. What do you think of these possible titles?

1.  *Computational Historical Thinking: With Applications in R*
2.  *The Compleat Computational Historian*
3.  *Computational History Methods in R*
