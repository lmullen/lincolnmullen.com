---
date: '2014-02-24T13:33:00+00:00'
slug: 'analyzing-historical-history-dissertations-gender'
title: 'Analyzing Historical History Dissertations: Gender'
---

*This series on "Analyzing Historical History Dissertations" is a work in progress and I've re-done some of these visualizations. If you would like to cite or link to this work in progress, please consider using the **[landing page](http://lincolnmullen.com/research/history-dissertations/)**, which will always have the most up-to-date information and a list of all the posts.*

**Update February 24, 2014 at 1:50 p.m.:** *I improved the regular expression that parses out first names, so the charts and table have been updated with fewer unknowns.*

One of the useful questions that can be asked of any data set is about the significance of gender. Unfortunately almost all of the data sets that I have gotten to analyze don't record gender. I'm willing to hazard a guess that that is true for many if not most data sets of interest to historians. However, many data sets, like the one I'm working on for history dissertations, do include a list of names, and first names are a pretty good indicator of gender.

But names are an imperfect indicator of gender. To begin with, you need a list of male and female names to assign a gender to a record. (There are a little under 12 thousand unique first names in the dissertations data set.) But the gender associated with names also varies over time. (Curiously, it seems that names tend to switch from male to female, but not the other way round.) [Cameron Blevins](http://www.cameronblevins.org/) had the really smart idea to look up the gender of names as they vary over time by using [Social Security Administration's names data set](http://catalog.data.gov/dataset/baby-names-from-social-security-card-applications-national-level-data). He and I are working together to compare the usefulness of his method in our respective data sets, so I'll write about the details later. (In the mean time, here is Cameron's [implementation in Python](https://github.com/cblevins/Gender-ID-By-Time) and an [R package](https://github.com/lmullen/gender) that I wrote. The READMEs explain more about the method.)

Using that method, I've guessed the gender of the first names in the dissertations data set, and produced this chart, which shows the number of dissertations produced in absolute terms.(The "unknowns" are names which did not appear in the Social Security Administration data for the given time period; as you would expect, in most cases they are non-European names.)

::: {.figure}
<p>
<img src="/downloads/historical-dissertations/plot-gender.svg" alt="plot of chunk plot-gender" /> <a href="/downloads/historical-dissertations/plot-gender.svg">SVG</a> \| <a href="/downloads/historical-dissertations/plot-gender.png">PNG</a> \| <a href="/downloads/historical-dissertations/plot-gender.pdf">PDF</a>
</p>
:::

Another way to examine this trend is to look at the proportion of male and female writers of history dissertations over time, which smooths out changes in the absolute number of dissertations produced. As you can see, the proportion of dissertations written by women has steadily increased since the 1950s. It appears to be plateauing with men producing a few percentage points more dissertations than women. (Keep in mind though that the number of names I was unable to identify is substantial, about 10% for each year through the 2000s.)

::: {.figure}
<p>
<img src="/downloads/historical-dissertations/dissertations-by-gender-proportion.svg" alt="plot of chunk dissertations-by-gender-proportion" /> <a href="/downloads/historical-dissertations/dissertations-by-gender-proportion.svg">SVG</a> \| <a href="/downloads/historical-dissertations/dissertations-by-gender-proportion.png">PNG</a> \| <a href="/downloads/historical-dissertations/dissertations-by-gender-proportion.pdf">PDF</a>
</p>
:::

If you want the percentages of dissertations written by men and women in tabular form, here it is.

**History Dissertations by Gender, 1950--2012**

<table>
<tr class="header">
<th style="text-align: right;">
year
</th>
<th style="text-align: center;">
proportion female
</th>
<th style="text-align: center;">
proportion male
</th>
<th style="text-align: center;">
proportion unknown
</th>
</tr>
<tr class="odd">
<td style="text-align: right;">
1950
</td>
<td style="text-align: center;">
0.14
</td>
<td style="text-align: center;">
0.76
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1951
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.85
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1952
</td>
<td style="text-align: center;">
0.12
</td>
<td style="text-align: center;">
0.84
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1953
</td>
<td style="text-align: center;">
0.05
</td>
<td style="text-align: center;">
0.90
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1954
</td>
<td style="text-align: center;">
0.06
</td>
<td style="text-align: center;">
0.91
</td>
<td style="text-align: center;">
0.03
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1955
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.88
</td>
<td style="text-align: center;">
0.03
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1956
</td>
<td style="text-align: center;">
0.09
</td>
<td style="text-align: center;">
0.86
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1957
</td>
<td style="text-align: center;">
0.07
</td>
<td style="text-align: center;">
0.90
</td>
<td style="text-align: center;">
0.03
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1958
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.86
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1959
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.82
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1960
</td>
<td style="text-align: center;">
0.08
</td>
<td style="text-align: center;">
0.90
</td>
<td style="text-align: center;">
0.02
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1961
</td>
<td style="text-align: center;">
0.08
</td>
<td style="text-align: center;">
0.88
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1962
</td>
<td style="text-align: center;">
0.13
</td>
<td style="text-align: center;">
0.82
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1963
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.84
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1964
</td>
<td style="text-align: center;">
0.08
</td>
<td style="text-align: center;">
0.86
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1965
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.85
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1966
</td>
<td style="text-align: center;">
0.10
</td>
<td style="text-align: center;">
0.84
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1967
</td>
<td style="text-align: center;">
0.11
</td>
<td style="text-align: center;">
0.85
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1968
</td>
<td style="text-align: center;">
0.11
</td>
<td style="text-align: center;">
0.84
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1969
</td>
<td style="text-align: center;">
0.11
</td>
<td style="text-align: center;">
0.83
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1970
</td>
<td style="text-align: center;">
0.13
</td>
<td style="text-align: center;">
0.82
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1971
</td>
<td style="text-align: center;">
0.14
</td>
<td style="text-align: center;">
0.82
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1972
</td>
<td style="text-align: center;">
0.16
</td>
<td style="text-align: center;">
0.80
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1973
</td>
<td style="text-align: center;">
0.15
</td>
<td style="text-align: center;">
0.80
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1974
</td>
<td style="text-align: center;">
0.20
</td>
<td style="text-align: center;">
0.74
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1975
</td>
<td style="text-align: center;">
0.20
</td>
<td style="text-align: center;">
0.76
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1976
</td>
<td style="text-align: center;">
0.20
</td>
<td style="text-align: center;">
0.76
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1977
</td>
<td style="text-align: center;">
0.23
</td>
<td style="text-align: center;">
0.73
</td>
<td style="text-align: center;">
0.04
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1978
</td>
<td style="text-align: center;">
0.24
</td>
<td style="text-align: center;">
0.71
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1979
</td>
<td style="text-align: center;">
0.26
</td>
<td style="text-align: center;">
0.69
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1980
</td>
<td style="text-align: center;">
0.28
</td>
<td style="text-align: center;">
0.67
</td>
<td style="text-align: center;">
0.05
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1981
</td>
<td style="text-align: center;">
0.27
</td>
<td style="text-align: center;">
0.67
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1982
</td>
<td style="text-align: center;">
0.28
</td>
<td style="text-align: center;">
0.66
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1983
</td>
<td style="text-align: center;">
0.30
</td>
<td style="text-align: center;">
0.63
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1984
</td>
<td style="text-align: center;">
0.31
</td>
<td style="text-align: center;">
0.62
</td>
<td style="text-align: center;">
0.07
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1985
</td>
<td style="text-align: center;">
0.30
</td>
<td style="text-align: center;">
0.62
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1986
</td>
<td style="text-align: center;">
0.30
</td>
<td style="text-align: center;">
0.61
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1987
</td>
<td style="text-align: center;">
0.32
</td>
<td style="text-align: center;">
0.60
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1988
</td>
<td style="text-align: center;">
0.32
</td>
<td style="text-align: center;">
0.57
</td>
<td style="text-align: center;">
0.11
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1989
</td>
<td style="text-align: center;">
0.32
</td>
<td style="text-align: center;">
0.59
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1990
</td>
<td style="text-align: center;">
0.34
</td>
<td style="text-align: center;">
0.54
</td>
<td style="text-align: center;">
0.12
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1991
</td>
<td style="text-align: center;">
0.35
</td>
<td style="text-align: center;">
0.53
</td>
<td style="text-align: center;">
0.12
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1992
</td>
<td style="text-align: center;">
0.38
</td>
<td style="text-align: center;">
0.50
</td>
<td style="text-align: center;">
0.12
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1993
</td>
<td style="text-align: center;">
0.38
</td>
<td style="text-align: center;">
0.51
</td>
<td style="text-align: center;">
0.11
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1994
</td>
<td style="text-align: center;">
0.38
</td>
<td style="text-align: center;">
0.52
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1995
</td>
<td style="text-align: center;">
0.39
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.12
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1996
</td>
<td style="text-align: center;">
0.40
</td>
<td style="text-align: center;">
0.50
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1997
</td>
<td style="text-align: center;">
0.40
</td>
<td style="text-align: center;">
0.50
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
1998
</td>
<td style="text-align: center;">
0.40
</td>
<td style="text-align: center;">
0.50
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
1999
</td>
<td style="text-align: center;">
0.41
</td>
<td style="text-align: center;">
0.51
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2000
</td>
<td style="text-align: center;">
0.42
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2001
</td>
<td style="text-align: center;">
0.42
</td>
<td style="text-align: center;">
0.50
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2002
</td>
<td style="text-align: center;">
0.43
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2003
</td>
<td style="text-align: center;">
0.42
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2004
</td>
<td style="text-align: center;">
0.44
</td>
<td style="text-align: center;">
0.46
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2005
</td>
<td style="text-align: center;">
0.45
</td>
<td style="text-align: center;">
0.46
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2006
</td>
<td style="text-align: center;">
0.42
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2007
</td>
<td style="text-align: center;">
0.45
</td>
<td style="text-align: center;">
0.46
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2008
</td>
<td style="text-align: center;">
0.41
</td>
<td style="text-align: center;">
0.49
</td>
<td style="text-align: center;">
0.10
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2009
</td>
<td style="text-align: center;">
0.45
</td>
<td style="text-align: center;">
0.47
</td>
<td style="text-align: center;">
0.08
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2010
</td>
<td style="text-align: center;">
0.44
</td>
<td style="text-align: center;">
0.47
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2011
</td>
<td style="text-align: center;">
0.43
</td>
<td style="text-align: center;">
0.48
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="odd">
<td style="text-align: right;">
2012
</td>
<td style="text-align: center;">
0.44
</td>
<td style="text-align: center;">
0.47
</td>
<td style="text-align: center;">
0.09
</td>
</tr>
<tr class="even">
<td style="text-align: right;">
2013
</td>
<td style="text-align: center;">
0.43
</td>
<td style="text-align: center;">
0.51
</td>
<td style="text-align: center;">
0.06
</td>
</tr>
</table>
