---
title: "Don't Waste Students' Time Installing Software in DH Courses"
date: "2016-01-08"
slug: 'dont-waste-students-time-installing-software-in-dh-courses'
tags:
- digital history
- AHA
- conference
- RStudio
- teaching
draft: false
---

Here is a summary of what I said in the [DH pedagogy lightening sessions](http://www.historians.org/annual-meeting/resources-and-guides/digital-history-at-aha16/digital-history-lightning-rounds/aha16-digital-pedagogy-lightning-round-lineup) at the AHA.

Simple idea 1: Installing software takes [a lot of time](http://www.pgbovine.net/command-line-bullshittery.htm), and installing software can often be harder and require more technological skill than actually using the software.

Simple idea 2: You must *scaffold* your digital history courses, so that one assignment leads into the next, and so that students build the methodological and technical skills that they need as they go through the course.

The problem is that students need to install the software before they can use it. The most technologically difficult, and the least pedagogically or historically interesting task, happens at the beginning of the course. This presents a tremendous barrier to student involvement. It wastes course time early in the semester, when building momentum is crucial.

Not so simple solution: My solution to this problem is to try to take the burden of installing software on myself so as to not waste students' time. For my "Data and DH" course next semester, as well as for previous courses, I have relied on an RRCHNM installation of [RStudio Server](https://www.rstudio.com/products/rstudio/). This lets students access a full development environment through their browser: no installing or configuring software.

You might object: there is no way my institution will give me a server of my own, and installing RStudio Server might be too difficult for me. For RStudio, at least, the [analogsea package](https://github.com/sckott/analogsea) can help you get a server up and running at Digital Ocean. Assuming you already have a [Digital Ocean](https://www.digitalocean.com/) account, it can be as simple as these few lines of code (though you will also have to add users and configure the memory).

```r
library(analogsea)

docklet_create() %>%
  docklet_rstudio()
```

The point is not that you should use RStudio Server (though it's great), and there are other options like Anaconda for Python. The point is to find a way to reduce or eliminate the waste of student time and attention that comes from installing software. Find a way to scaffold your courses so that you can get straight into the digital history.
