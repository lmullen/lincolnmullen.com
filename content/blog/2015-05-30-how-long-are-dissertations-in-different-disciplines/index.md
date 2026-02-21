---
title: 'How Long Are Dissertations in Different Disciplines?'
date: "2015-05-30"
slug: 'how-long-are-dissertations-in-different-disciplines'
description: "ProQuest data on dissertation lengths across disciplines shows history with the longest median at 320 pages, while economics hovers near the bottom."
tags:
- digital history
- data analysis
draft: false
---

About a year ago when I was finishing my dissertation I got intensely interested in seeing what I could figure out about history dissertations from data. I'll let you draw your own conclusions as to why that seemed more urgent than writing my own dissertation. I wrote a series of posts on the topic using a dataset of historical history dissertations from ProQuest. But ProQuest also gave me a dataset of contemporary dissertations from all fields that I have not yet done anything with.

Recently I rediscovered Michael Beck's [post on the length of dissertations by department](https://beckmw.wordpress.com/2014/07/15/average-dissertation-and-thesis-length-take-two/) at the University of Minnesota (perhaps via Hacker News?). I thought I would redo Beck's visualization of dissertation lengths using the ProQuest dataset to see how the dissertations compare at most universities in the United States.

Here is a chart of the distribution of dissertation lengths by subject for universities in the United States and Canada from 2001 to 2013.

{{< figure src="diss-length.png" caption="Median length of PhD dissertations by subject in U.S. and Canadian universities since 2001." >}}

Note that the key difference between Beck's chart and mine that he is grouping is dissertations by "major," which I assume is the nearly the same thing as department. The ProQuest data does not include the department for about three-quarters of the dissertations, so I have grouped the dissertations by their primary subject field. Whether that is a good proxy for discipline or department, you can decide for yourself, but I think it's a reasonable approximation for this purpose. I've also filtered out diploma mills (why does ProQuest let them into their database?), non-PhD dissertations (no disrespect intended though), universities from outside the US and Canada (because they are a tiny part of the data) and any subject that had fewer than twenty-five dissertations. The ProQuest data seems not to include many dissertations from some disciplines, such as chemistry.

The chart is broadly similar to Beck's. Both show history (and in my case, art history) as the discipline with the longest dissertations, and economics with the shortest or second shortest. In the ProQuest data the median history dissertation is 320 pages, while in the Minnesota data it is about 290. That's not surprising, given [how much the median page count can vary](/blog/analyzing-historical-history-dissertations-page-counts-by-university/) for history dissertations from department to department. Both charts show that some disciplines can have quite a wide variability in the length of their dissertations. Take music, for instance, which can have very short dissertations (possibly compositions) and very long dissertations (perhaps histories of music). I wouldn't put too much stock in either chart. But they do show how different academic disciplines approach the dissertation with, broadly speaking, humanistic disciplines on the long end of the spectrum and STEM disciplines on the short end. The larger, broader data set here mostly confirms what Beck found.
