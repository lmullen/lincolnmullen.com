---
categories:
- digital history
- rstats
date: '2017-07-10T16:07:29+00:00'
slug: 'usaboundaries-v0-3-0-released'
title: 'New release: USAboundaries v0.3.0'
---

I've recently published [version 0.3.0 of my USAboundaries R package](http://ropensci.github.io/USAboundaries/) to CRAN. USAboundaries provides access to spatial data for U.S. counties, states, cities, congressional districts, and zip codes. Of course you can easily get contemporary boundaries from lots of places, but this package lets you specify dates and get the locations for historical county and state boundaries as well as city locations.

{{< figure src="usaboundaries-sample.png" caption="Sample boundaries from the USAboundaries package." >}}

This version of the package has a number of new features. It has jumped on the [Simple Features](https://github.com/edzer/sfr) bandwagon, so now all boundary data is returned as an sf object. This version also includes updated shapefiles from the U.S. Census for contemporary data, as well as new centroids for Zipcode Tabulation Areas and historical city populations courtesy of Erik Steiner's [project from CESTA](https://github.com/cestastanford/historical-us-city-populations).

I'm especially glad that the package has added a new author: [Jordan Bratt](http://jfbratt.org/), a PhD student at George Mason and a collaborator on [Mapping Early American Elections](http://earlyamericanelections.org/). Jordan added functionality to the package that lets users [get projections](http://ropensci.github.io/USAboundaries/reference/state_plane.html) from the State Plane Coordinate System, so that they can make locally accurate maps at the level of the state or below.

The package has a [new website](/software/usaboundaries/) thanks to [pkgdown](http://hadley.github.io/pkgdown/). You can read the full release notes at the [package website](http://ropensci.github.io/USAboundaries/news/index.html).
