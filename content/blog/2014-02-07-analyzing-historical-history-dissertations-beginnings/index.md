---
title: 'Analyzing Historical History Dissertations: Beginnings'
date: "2014-02-07"
slug: 'analyzing-historical-history-dissertations-beginnings'
description: "Opening a series using a ProQuest dump of 84,000 history dissertations since 1878 to ask how historians have studied the past—in quantity, length, subject, and generation."
tags:
- digital history
draft: false
---

Some scholars whose work I admire have done some useful data analysis about dissertations. Michael Beck [wrote a post](http://beckmw.wordpress.com/2013/04/15/how-long-is-the-average-dissertation/) where he scraped dissertation data from the University of Minnesota since 2007, and found that history dissertations are the longest. Ben Schmidt [asked how the length of history dissertations](http://rpubs.com/benmschmidt/princetonhistorydiss) at Princeton has varied since 1988, and some other interesting questions about the age of the author. Ben went on to ask "[What years do historians write about?](http://sappingattention.blogspot.com/2013/05/what-years-do-historians-write-about.html)" using the "all dissertation titles in history the American Historical Association knows about from the last 120 years." It's a great post: go [read it](http://sappingattention.blogspot.com/2013/05/what-years-do-historians-write-about.html).

I want to ask some related questions, but I hoped to gather as comprehensive a data set as possible. My first inclination was to scrape the data from ProQuest's [Dissertations and Theses Database](http://www.proquest.com/en-US/catalogs/databases/detail/pqdt.shtml). But in response to one of my tweets (followed up by several months of e-mails), ProQuest was kind enough to send me two dumps from that database. One contains all dissertations since 2000. The other contains all U.S. and Canadian dissertations in history and cognate fields since about 1878. While the terms of my agreement with ProQuest don't allow me to post the raw data (or, I suppose, the cleaned up data) I have put my code for exploring these datasets in a [GitHub repostory](https://github.com/lmullen/dissertations-data). This post is the first in a series of explorations of the data, starting with the historical rather than the contemporary dissertations.

From the MARC records that ProQuest gave me, I've been able to extract a number of fields, only some of which are available for each record. There is a ProQuest assigned ID number and sometimes an ISBN. Items always have an author and a title, and almost always a page count and a year of graduation. Many of the dissertations have an abstract, which I think will be useful for mining the topics that historians have studied. There are also Library of Congress subject fields, but these are usually very generic. Almost always a university is listed, along with a "school code" which I think standardizes university names; sometimes a department is listed as well. Some items have the lead adviser and other committee members listed separately; some mash them together; still others don't have the information at all. I'm hoping the adviser data will let me trace scholarly generations. The degree conferred is always listed. And finally there is a URL to a [ProQuest record](http://gateway.proquest.com/openurl?url_ver=Z39.88-2004&rft_val_fmt=info:ofi/fmt:kev:mtx:dissertation&res_dat=xri:pqm&rft_dat=xri:pqdiss:7928093). For reasons I don't understand, the ProQuest online record has different fields than the MARC dump.

The historical dissertations data contains 132,360 records. But not all of those are dissertations. There are 94,530 PhD dissertations and 23,362 MA theses; the rest are scattered among many different degrees. Furthermore, not all of the dissertations and theses are historical, whether one limits that category narrowly to dissertations produced in history departments or more broadly to dissertations with a historical methodology. Limiting the data to items whose primary subject heading is somehow related to history, we are left with 86,237 dissertations and theses. In summary, the data contain 84,428 PhD dissertations and 20,528 MA theses about history.

For a first plot of the data, we can look at the number of dissertations and theses over time.

<!--more-->

Because there were so few dissertations in the late nineteenth and early twentieth century (I suspect some are missing), let's limit the plot to the years after 1950.

![](2014-02-07-number-dissertations.png)

I'll do some more analysis soon. If you have any questions that you think the data can answer, let me know.
