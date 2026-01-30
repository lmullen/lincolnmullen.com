---
title: 'Analyzing Historical History Dissertations: Location'
date: "2014-02-12"
slug: 'analyzing-historical-history-dissertations-location'
tags:
- digital history
draft: false
---

*This series on "Analyzing Historical History Dissertations" is a work in progress and I've re-done some of these visualizations. If you would like to cite or link to this work in progress, please consider using the **[landing page](http://lincolnmullen.com/research/history-dissertations/)**, which will always have the most up-to-date information and a list of all the posts.*

So far I have mostly looked at the number of dissertations, and the page counts of dissertations. I'd like to refine the analysis of page counts by seeing if the length of dissertations varies by [NRC rankings of history departments](https://www.historians.org/publications-and-directories/perspectives-on-history/december-2010/nrc-report-provides-data-on-history-doctoral-programs). Today, though, I want to look at the locations where history dissertations were produced.

Below is a map of the universities that have produced dissertations about history. Once again, since I'm using the ProQuest data, these are dissertations about history, a bigger set than dissertations written in history departments. In case the map isn't self-explanatory, the numbers in red boxes represent the total number of dissertations produced in a geographic area. You can put your mouse over the box to see the boundaries of the geographic area. Zooming in or clicking on the box will cluster the points into smaller and smaller units. Eventually you'll get individual markers for each university. Clicking on a marker will tell you the university name and the number of history dissertations it is has produced.

<link rel="stylesheet" href="//files.lincolnmullen.com/downloads/historical-dissertations/location/MarkerCluster.css" />

::: {#map}
:::

Unsurprisingly, the Northeast can be thought of as a mega-center, with Boston, New Haven, New York, Philadelphia, and the areas in between producing on the order of 28 thousand history dissertations. But there are many other smaller centers: Washington, Chicago, Michigan, the North Carolina research triangle, the University of California and Stanford. The University of Wisconsin is a behemoth unto itself, and other universities (e.g., the Indiana University at Bloomington) produce many dissertations. Columbia University has the most dissertations in the data set.

The main observations I would make are these. First, the Northeast produces a lot of dissertations. But second, the writing of history dissertations is scattered across the country, and regions like the Midwest, the upper Southeast, southern California, and Texas are sizable in themselves. And finally, a few schools seem to produce most of the output of history dissertations. Very roughly speaking, 25 schools have produced about half the dissertations in history, while another 350 or so have produced the other half. (See the chart at the end of [this post](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts-by-university/).)

In future posts I want to look at not just the places where historians write, but the places they write about. I'm curious to see if I can detect some correlation between the universities that produce dissertations and the places those dissertations are about. Surely there is a reason for all these dissertations about Boston and Chicago.

---

Caveats: If you zoom the map out, you'll see that there is some cruft in the data. There are 341 dissertations in Asia, 55 in Australia, 94 in Europe, 6 in Africa, and 2 in South America. There are also 4,608 dissertations in Canada. Clearly this data set does not contain all the dissertations from outside North America. It's not clear to me whether the data set is comprehensive for Canadian dissertations or not. I knew that there were some non-U.S. dissertations in the data set, but I didn't know how many or from where exactly until I made this map. For future analysis I'm going to trim all the dissertations from outside the United States, and go back and redo the earlier analysis, though the difference will likely be minimal. That's why I practice [reproducible research](http://christophergandrud.github.io/RepResR-RStudio/). And as I observed in the post on [page counts by university](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts-by-university/), the data set only has information for some universities after the 1980s or 1990s. That will skew the numbers on this map down.
