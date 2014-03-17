---
layout: post
title: "Charting Faculty Salaries by Discipline, Rank, and Institution"
date: 2014-03-17 13:07
categories: 
- r-stats
- visualization
---

Today the *Chronicle of Higher Education* published a table of data on
the "[Average Salaries of Tenured and Tenure-Track Faculty at 4-Year
Colleges, 2013--14][]," with an [accompanying article][]. Except for one
chart in the article about pay increases, the *Chronicle* did not
present the data in visual form. So I offer you two charts you can send
to your loved ones to justify your life choices. You'll probably want to
click through to the SVG or PDF versions of these charts so that you can
scroll and zoom.

The first chart ([SVG][], [PDF][]) shows salaries grouped by institution
type:

<img src="/downloads/post/faculty-salaries-by-institution.svg">

The second chart ([SVG][1], [PDF][2]) shows salaries grouped by
discipline:

<img src="/downloads/post/faculty-salaries-by-discipline.svg">

*Note:* The labels on the bars are average salaries for assistant,
associate, and full professors; I've left out new assistant professors.
The vertical lines in the first chart are the average salaries by rank
for all disciplines. You'll notice that in a number of disciplines, new
assistant professors make more on average than assistant professors. I
assume this is because there are fewer new assistant professors and so
the average is more likely to be skewed, because wealthier institutions
are the ones that are hiring new assistant professors, because raises
have stagnated during the economic downturn, and because of thing they
call "[salary compression][]".

You can [check my work here][].

  [Average Salaries of Tenured and Tenure-Track Faculty at 4-Year
  Colleges, 2013--14]: http://chronicle.com/article/Average-Salaries-of-Tenured/145283/
  [accompanying article]: http://chronicle.com/article/Remedial-Educators-Contest/145349
  [SVG]: /downloads/post/faculty-salaries-by-institution.svg
  [PDF]: /downloads/post/faculty-salaries-by-institution.pdf
  [1]: /downloads/post/faculty-salaries-by-discipline.svg
  [2]: /downloads/post/faculty-salaries-by-discipline.pdf
  [salary compression]: http://www.insidehighered.com/news/2013/02/11/university-tries-deal-salary-compression-among-faculty-members
  [check my work here]: https://gist.github.com/lmullen/9606187
