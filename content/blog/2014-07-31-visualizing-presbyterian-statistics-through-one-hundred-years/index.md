---
title: Visualizing Presbyterian Statistics Through One Hundred Years
date: "2014-07-31"
slug: 'visualizing-presbyterian-statistics-through-one-hundred-years'
description: "Herman Carl Weber's 1927 *Presbyterian Statistics through One Hundred Years* offers a surprising precedent for digital history: denominations have always needed their numbers visualized."
tags:
- data visualization
- American religious history
draft: false
---

The title *[Presbyterian Statistics through One Hundred Years, 1826--1926: Tabulated, Visualized, and Interpreted](http://catalog.hathitrust.org/Record/007109885)* sounds like contemporary digital history project. Click the link, and if you're an optimist about digital methods you might expect a revolutionary new methodology for doing history that overturns old models; if you are a pessimist you might expect to find some "Big Data" hubris. But what you'll find is a book from 1927, compiled the Presbyterian minister and employee of the General Council of the Presbyterian Church U.S.A., Herman Carl Weber.[^1] Weber compiled the data at the behest of the PCUSA, laboriously compiling the statistics because the General Council wanted to know what insights could be gathered from its record.

The first part of Weber's work included dozens of tables like the one below.[^2] Denominations often maintained records and published them in yearbooks and annual reports.[^3] What is especially useful with Weber's figures are that he compiled the numbers longitudinally---apparently over a period of three years of work---so that it is possible to see change over time.

{{< figure src="table.png" caption="Figure 1: A table of membership figures from Weber, p. 12." >}}

Having gathered the numbers, Weber turned them into a series of visualizations and accompanying interpretations aimed at improving both the national church and individual congregations. In turning records into tables, and tables into graphs, Weber had some confidence that "the circle of those who can understand visualizations readily is very large," though he also deemed it wise to use prose to "suggest what some of the visualizations mean" (3--4). In 1927 the idea of visualizations was not particularly innovative. Beginning in the 1780s the Scottish political economist [William Playfair](http://en.wikipedia.org/wiki/William_Playfair) invented some of the foundational visualizations, including line charts, bar charts, and [time series](http://en.wikipedia.org/wiki/Time_series). As [Susan Schulten has shown](http://www.mappingthenation.com/), maps and other kinds of data-driven visualizations were important techniques for nineteenth-century American science, engineering, and statecraft. Nor is it any coincidence that visualization was a technique for interpretation in the same period that the historians in America were turning themselves into professionals.

Weber's charts told a story. In the first visualization in the book, reproduced below, notice the rise in Presbyterian membership over a hundred years. (The vertical orientation of the chart exaggerates growth, but from what I can tell Weber chose that orientation because of the constraints of printing and not out of any intention to distort.) Weber called this a "bird's-eye view of the membership" (46).

{{< figure src="membership.png" caption="Figure 2: A chart of membership from Weber, p. 46." >}}

That first chart with its steep curve made the Presbyterians look good, but Weber was not so naive. He knew that in the nineteenth century United States everything was growing, and that some denominations were growing faster than the Presbyterians and had a larger share of the population. Weber included the "ratio chart" below, derived from the Yale economist and statistician [Irving Fisher](http://en.wikipedia.org/wiki/Irving_Fisher), that used a logarithmic scale to show "the same relative increases" with "the same slope" (203). The short explanation is that Weber found a way to show that the Presbyterians were only growing at the same rate as the population, and that the Baptists and Methodists were growing faster.

{{< figure src="ratio.png" caption="Figure 3: A chart from Weber comparing denominations, p. 49." >}}

Weber's most intriguing work was a pair of charts about people who joined the Presbyterian church. Weber was quite interested in whether or not the Presbyterians were fulfilling their mission to win converts. The chart below of "Accessions on Confession" was "the line of response from the young, adolescent life of the Church to the call of the Kingdom." Weber annotated the chart to show the underlying causes in the fluctuations: the peaks for the revivals of Finney, Moody, and Sunday; the falling off in conversions caused by denominational splits, heresy trials, and the controversy with science.

{{< figure src="accessions.png" caption="Figure 4: A chart of Accessions on Confession from Weber, p. 51." >}}

The chart of total new members showed a positive picture of Presbyterians' growth, but again, Weber knew enough to normalize the data to account for growth in population. In the chart below, he created an "Evangelistic Index," the "proportion of new members in the total membership." This was "the most significant of all the visualizations submitted in this volume" because it "portrays the actual record of the Church in the primary functional responsibility which has been committed to it" (57).

{{< figure src="evangelistic.png" caption="Figure 5: The Evangelistic Index, p. 56." >}}

My point in showing Weber's visualizations is two-fold. First, visualization as a method for history is much older than current debates would sometimes lead one to believe. Such methods are deeply rooted in the nineteenth and early twentieth century. The corollary is that the history of American religion, insofar as it makes use of quantitative methods, must depend on data collected, aggregated, visualized, and interpreted by the very people whom historians wish to study.

[^1]: The book is Herman C. Weber, *Presbyterian Statistics Through One Hundred Year, 1826--1926: Tabulated, Visualized, and Interpreted* (Philadelphia: The General Council, Presbyterian Church in the U.S.A., 1927), <http://catalog.hathitrust.org/Record/007109885>.

[^2]: For use in my dissertation, and for a future project on the demography of American religion, I've turned some of Weber's data tables into [a useable format](https://github.com/lmullen/demographics-religion/tree/master/data/weber-1927). If you want to see what I've started to do with the data---though this is very much a work in progress---you can see [this repository](https://github.com/lmullen/presbyterian-analysis).

[^3]: See for example [this run of Congregationalist yearbooks](http://www.congregationallibrary.org/periodicals/yearbooks) from 1854 to 1960 compiled by the fine folks at the [Congregational Library](http://www.congregationallibrary.org/).
