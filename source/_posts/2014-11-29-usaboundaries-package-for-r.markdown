---
categories:
- 'r-stats'
date: '2014-11-29 15:04'
layout: post
title: USAboundaries Package for R
...

The Dr. William M. Scholl Center for American History and Culture has
released a set of shapefiles for historical state and county boundaries
from 1629 to 2000 as part of their *[Atlas of Historical County
Boundaries][]*. These shapefiles are very useful for creating historical
maps, so I've bundled them with a few convenience functions as the
[USAboundaries package for R][]. This package makes it easy to make a
map for any arbitrary day in United States history.

    library(USAboundaries)
    library(sp)

    map_1844 <- us_boundaries(as.Date("1844-03-21"))
    plot(map_1844)

{% figure usaboundaries/map1844 png 'The United States on March 21, 1844' %}

For more, see the documentation. The package is available on [CRAN][].

  [Atlas of Historical County Boundaries]: http://publications.newberry.org/ahcbp/
  [USAboundaries package for R]: https://github.com/ropensci/usaboundaries
  [CRAN]: http://cran.rstudio.org/web/packages/USAboundaries/index.html
