---
title: 'USAboundaries v0.2.0 available on CRAN'
date: "2016-01-04"
slug: 'usaboundaries-v0-2-0-available-on-cran'
description: "USAboundaries v0.2.0 adds current Census boundaries to the package's historical county and state data, with higher resolution options and a cleaner interface."
tags:
- R
- spatial humanities
draft: false
---

A new release of the [USAboundaries package](https://github.com/ropensci/USAboundaries) (v0.2.0) for R is [available on CRAN](https://cran.r-project.org/package=USAboundaries). This package continues to provide historical boundaries of U.S. counties and states from 1629 to 2000, thanks to the Newberry Library's *[Atlas of Historical County Boundaries](http://publications.newberry.org/ahcbp/)*. In this release I have added current county, state, and congressional district boundaries from the [U.S. Census Bureau](https://www.census.gov/geo/maps-data/data/tiger-cart-boundary.html). Both the historical and contemporary boundaries data gain higher resolution versions suitable for mapping at the level of the state rather than the nation. This higher resolution data is optional, and will be installed the first time that a user requests it. Finally, the entire package interface has been improved, adding geography-specific functions (e.g., `us_states()`, `us_counties()`) instead of forcing everything through a single function, and removing a bunch of needless package dependencies.
