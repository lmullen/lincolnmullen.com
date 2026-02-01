---
title: Resources for Mapping US Boundaries Over Time
date: "2014-04-23"
slug: 'resources-for-mapping-us-boundaries-over-time'
tags:
- digital history
- spatial humanities
draft: false
---

When making maps it's not hard to find contemporary boundaries. The [Census Bureau](http://www.census.gov/geo/maps-data/data/tiger-line.html) for the United States and [Natural Earth](http://www.naturalearthdata.com/) for the rest of the world offer reliable shapefiles. Languages like R include [packages](http://cran.r-project.org/web/packages/maps/index.html) which make accessing boundary data very easy. But when working on maps for history, using contemporary boundaries with historical data makes for glaring anachronisms. What I've wanted for research purposes are historical boundaries of the United States which include the changes for each year, and which I can trust because of the evidence of scholarship (e.g., citations) behind them. I also want these boundaries for teaching, so that geography can provide a backbone to the religious, cultural, social, and political changes I'm usually discussing. These are the sources that I've found with a few comments on their usefulness for research and teaching.

First, the single most useful resource that I've found is the *[Atlas of Historical County Boundaries](http://publications.newberry.org/ahcbp/index.html)* from the [Newberry Library](http://www.newberry.org/). This atlas includes boundary files for states and territories from 1783 to 2000, and also for counties from 1629 to 2000. The boundaries are available as both shapefiles and KMZ files suitable for Google Earth. (I've found that it's easier to get students up to speed using Google Earth than QGIS.) The boundary files are offered at various degrees of simplification, letting you pick between precision and file size without having to do the simplification yourself. For my purposes, the most important thing about the Newberry's boundary data is that it comes in a single file. This means that you can more easily filter the boundaries by date, for example, in QGIS or D3. Making a complex map showing change over time gets a lot simpler when you can filter one file instead of having to merge many files. The maps come with impressive annotations, including citations for the boundaries for each feature, and discussions of what changed about the boundaries. The data is licensed CC-NC-SA. The Newberry Library also offers an [online viewer](http://historical-county.newberry.org/website/USA/viewer.htm) for the files, which is clunky.[^1]

{{< figure src="us-1851.png" caption="Figure 1: Filtering the Newberry shapefiles to show the United States on January 1, 1851." >}}

Second, [Michael Porath](http://michaelporath.com/) has released a [series of state- and territory-level shapefiles and GeoJSON files](http://poezn.github.io/us-history-maps/) for 141 dates between March 4, 1789 and August 21, 1959. Porath [writes](http://michaelporath.com/projects/manifest-destiny/about.html), "All content for this visualization is taken verbatim from the original Wikipedia page [Territorial evolution of the United States](http://en.wikipedia.org/wiki/Territorial_evolution_of_the_United_States)." His shapefiles are especially useful if you want to make a map for a particular point in time, in which case you can just use the shapefile from the next earliest date. (Bonus: the shapefiles are available in a Git repository, and are licensed CC-SA.)

Porath has mapped the territorial changes himself in a project he calls "[Manifest Destiny](http://michaelporath.com/projects/manifest-destiny/)." The many maps let you page through the territorial changes in the United States. For example, try paging through the maps which show the order in which southern states seceded from the Union in 1860--61. I can see how these maps could be used to good effect in the U.S. survey.

{{< figure src="porath.png" caption="Figure 2: Michael Porath's maps displaying the United States on January 10, 1861." >}}

Third, the [National Historical Geographic Information System](https://www.nhgis.org/) offers [state- and county-level shapefiles](https://www.nhgis.org/documentation/gis-data) for each decade since 1790. The NHGIS shapefiles are meant for use with the historical census data offered by the NHGIS. Though I've used these shapefiles for general purpose mapping before,[^2] I can only recommend them for use when you are also using the NHGIS's census data.[^3] The maps are only offered every decade, so many of the more interesting boundary changes are grouped together. The shapefiles have a peculiar projection, so for many uses you'll have to re-project them.[^4] And the shapefiles are full resolution, so you'll have to do the simplification yourself, which is tedious and error-prone. But to take advantage of the wealth of historical census data in the [NHGIS](https://www.nhgis.org/), these shapefiles are quite useful.

Go forth and make maps without anachronism.

[^1]: I think the interface is provided by an ArcGIS web viewer.

[^2]: For [mapping Catholic dioceses](http://lincolnmullen.com/blog/maps-of-catholic-dioceses-in-the-us-canada-and-mexico-take-two/).

[^3]: As in this [map of slavery in 1860](http://lincolnmullen.com/blog/a-better-map-of-slavery-in-1860/).

[^4]: When I make maps in [ggplot2](http://ggplot2.org/), R sometimes chokes on fortified shapefiles that are not in EPSG 4326 or EPSG 3857, but otherwise does fine plotting and re-projecting the shapefiles itself.
