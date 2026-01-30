---
title: 'Analyzing Historical History Dissertations: Gender'
date: "2014-02-24"
slug: 'analyzing-historical-history-dissertations-gender'
tags:
- digital history
draft: false
---

*This series on "Analyzing Historical History Dissertations" is a work in progress and I've re-done some of these visualizations. If you would like to cite or link to this work in progress, please consider using the **[landing page](http://lincolnmullen.com/research/history-dissertations/)**, which will always have the most up-to-date information and a list of all the posts.*

**Update February 24, 2014 at 1:50 p.m.:** *I improved the regular expression that parses out first names, so the charts and table have been updated with fewer unknowns.*

One of the useful questions that can be asked of any data set is about the significance of gender. Unfortunately almost all of the data sets that I have gotten to analyze don't record gender. I'm willing to hazard a guess that that is true for many if not most data sets of interest to historians. However, many data sets, like the one I'm working on for history dissertations, do include a list of names, and first names are a pretty good indicator of gender.

But names are an imperfect indicator of gender. To begin with, you need a list of male and female names to assign a gender to a record. (There are a little under 12 thousand unique first names in the dissertations data set.) But the gender associated with names also varies over time. (Curiously, it seems that names tend to switch from male to female, but not the other way round.) [Cameron Blevins](http://www.cameronblevins.org/) had the really smart idea to look up the gender of names as they vary over time by using [Social Security Administration's names data set](http://catalog.data.gov/dataset/baby-names-from-social-security-card-applications-national-level-data). He and I are working together to compare the usefulness of his method in our respective data sets, so I'll write about the details later. (In the mean time, here is Cameron's [implementation in Python](https://github.com/cblevins/Gender-ID-By-Time) and an [R package](https://github.com/lmullen/gender) that I wrote. The READMEs explain more about the method.)

Using that method, I've guessed the gender of the first names in the dissertations data set, and produced this chart, which shows the number of dissertations produced in absolute terms.(The "unknowns" are names which did not appear in the Social Security Administration data for the given time period; as you would expect, in most cases they are non-European names.)

{{< figure src="//files.lincolnmullen.com/downloads/historical-dissertations/plot-gender.svg" >}}

[SVG](//files.lincolnmullen.com/downloads/historical-dissertations/plot-gender.svg) | [PNG](//files.lincolnmullen.com/downloads/historical-dissertations/plot-gender.png) | [PDF](//files.lincolnmullen.com/downloads/historical-dissertations/plot-gender.pdf)

Another way to examine this trend is to look at the proportion of male and female writers of history dissertations over time, which smooths out changes in the absolute number of dissertations produced. As you can see, the proportion of dissertations written by women has steadily increased since the 1950s. It appears to be plateauing with men producing a few percentage points more dissertations than women. (Keep in mind though that the number of names I was unable to identify is substantial, about 10% for each year through the 2000s.)

{{< figure src="//files.lincolnmullen.com/downloads/historical-dissertations/dissertations-by-gender-proportion.svg" >}}

[SVG](//files.lincolnmullen.com/downloads/historical-dissertations/dissertations-by-gender-proportion.svg) | [PNG](//files.lincolnmullen.com/downloads/historical-dissertations/dissertations-by-gender-proportion.png) | [PDF](//files.lincolnmullen.com/downloads/historical-dissertations/dissertations-by-gender-proportion.pdf)

If you want the percentages of dissertations written by men and women in tabular form, here it is.

**History Dissertations by Gender, 1950--2012**

| year | proportion female | proportion male | proportion unknown |
|-----:|:-----------------:|:---------------:|:------------------:|
| 1950 | 0.14 | 0.76 | 0.10 |
| 1951 | 0.10 | 0.85 | 0.05 |
| 1952 | 0.12 | 0.84 | 0.04 |
| 1953 | 0.05 | 0.90 | 0.04 |
| 1954 | 0.06 | 0.91 | 0.03 |
| 1955 | 0.10 | 0.88 | 0.03 |
| 1956 | 0.09 | 0.86 | 0.05 |
| 1957 | 0.07 | 0.90 | 0.03 |
| 1958 | 0.10 | 0.86 | 0.04 |
| 1959 | 0.10 | 0.82 | 0.08 |
| 1960 | 0.08 | 0.90 | 0.02 |
| 1961 | 0.08 | 0.88 | 0.04 |
| 1962 | 0.13 | 0.82 | 0.06 |
| 1963 | 0.10 | 0.84 | 0.06 |
| 1964 | 0.08 | 0.86 | 0.06 |
| 1965 | 0.10 | 0.85 | 0.06 |
| 1966 | 0.10 | 0.84 | 0.05 |
| 1967 | 0.11 | 0.85 | 0.04 |
| 1968 | 0.11 | 0.84 | 0.04 |
| 1969 | 0.11 | 0.83 | 0.05 |
| 1970 | 0.13 | 0.82 | 0.06 |
| 1971 | 0.14 | 0.82 | 0.04 |
| 1972 | 0.16 | 0.80 | 0.04 |
| 1973 | 0.15 | 0.80 | 0.04 |
| 1974 | 0.20 | 0.74 | 0.05 |
| 1975 | 0.20 | 0.76 | 0.04 |
| 1976 | 0.20 | 0.76 | 0.05 |
| 1977 | 0.23 | 0.73 | 0.04 |
| 1978 | 0.24 | 0.71 | 0.05 |
| 1979 | 0.26 | 0.69 | 0.05 |
| 1980 | 0.28 | 0.67 | 0.05 |
| 1981 | 0.27 | 0.67 | 0.06 |
| 1982 | 0.28 | 0.66 | 0.06 |
| 1983 | 0.30 | 0.63 | 0.06 |
| 1984 | 0.31 | 0.62 | 0.07 |
| 1985 | 0.30 | 0.62 | 0.08 |
| 1986 | 0.30 | 0.61 | 0.09 |
| 1987 | 0.32 | 0.60 | 0.08 |
| 1988 | 0.32 | 0.57 | 0.11 |
| 1989 | 0.32 | 0.59 | 0.09 |
| 1990 | 0.34 | 0.54 | 0.12 |
| 1991 | 0.35 | 0.53 | 0.12 |
| 1992 | 0.38 | 0.50 | 0.12 |
| 1993 | 0.38 | 0.51 | 0.11 |
| 1994 | 0.38 | 0.52 | 0.10 |
| 1995 | 0.39 | 0.49 | 0.12 |
| 1996 | 0.40 | 0.50 | 0.10 |
| 1997 | 0.40 | 0.50 | 0.09 |
| 1998 | 0.40 | 0.50 | 0.10 |
| 1999 | 0.41 | 0.51 | 0.09 |
| 2000 | 0.42 | 0.49 | 0.09 |
| 2001 | 0.42 | 0.50 | 0.08 |
| 2002 | 0.43 | 0.49 | 0.08 |
| 2003 | 0.42 | 0.49 | 0.09 |
| 2004 | 0.44 | 0.46 | 0.09 |
| 2005 | 0.45 | 0.46 | 0.09 |
| 2006 | 0.42 | 0.49 | 0.09 |
| 2007 | 0.45 | 0.46 | 0.10 |
| 2008 | 0.41 | 0.49 | 0.10 |
| 2009 | 0.45 | 0.47 | 0.08 |
| 2010 | 0.44 | 0.47 | 0.09 |
| 2011 | 0.43 | 0.48 | 0.09 |
| 2012 | 0.44 | 0.47 | 0.09 |
| 2013 | 0.43 | 0.51 | 0.06 |
