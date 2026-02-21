---
title: TopoJSON in R with rgdal
date: "2014-09-26"
slug: 'topojson-in-r-with-rgdal'
description: "Recent versions of GDAL can read TopoJSON directly, meaning rgdal can too—bringing D3's compact, topologically clean map format into R's mapping workflow."
tags:
- R
- spatial humanities
draft: false
---

In making maps in R, the biggest point of pain for me has been reading in shapefiles. D3 has been far better for making maps because [TopoJSON's](https://github.com/mbostock/topojson) topological format for representing geometries is so much better than the format of shapefiles. Bob Rudis's post "[Overcoming D3 Cartographic Envy With R + ggplot](http://rud.is/b/2014/09/26/overcoming-d3-cartographic-envy-with-r-ggplot/)" taught me that the most recent version of [GDAL/OGR](http://www.gdal.org/) can read in TopoJSON, and thus that the R package [rgdal](http://cran.rstudio.org/web/packages/rgdal/) can as well. This will make mapping in R much easier.
