---
title: 'Lower resolution shapefiles of historic state boundaries'
date: "2013-09-11"
slug: 'lower-resolution-shapefiles-of-historic-state-boundaries'
description: "Simplified versions of NHGIS historic state boundary shapefiles for faster iteration in R—reducing the 1860 file from 27.5 MB to 1.7 MB without visible quality loss."
tags:
- digital history
- mapping
- R
draft: false
---

A few weeks ago I [posted some shapefiles](/blog/shapefiles-of-historic-u-s-state-boundaries-for-use-in-r/) of historic state boundaries, converted to be used in R. These shapefiles are very high resolution, and so can be slow to work with when iteratively making maps. I've made some lower resolution shapefiles. (The [Ramer-Douglas-Peucker Algorithm](http://karthaus.nl/rdp/) that simplifies the shapefiles is quite interesting.) The new files are much quicker to work with. The R object for the 1860 boundaries, for example, has 1.8 million points and weighs in at 27.5 MB; the simplified object has 122,000 points and is only 1.7 MB. I think I've made the reduction so that there is no noticeable loss of quality when making maps at the national scale at full page print resolution. But you could also substitute the high-resolution version when you're ready to make a final version.

-   [Low-resolution shapefiles](https://dl.dropboxusercontent.com/u/1125028/historic.states.low-res.zip)
-   [High resolution shapefiles](https://www.dropbox.com/s/xu0bug1x7hgm5h1/historic-state-boundaries.zip)
-   [Code to generate the shapefiles](https://github.com/lmullen/demographics-religion/blob/master/historic.shapefiles.r)

The [NHGIS](https://nhgis.org) has kindly granted me permission to share these converted files, but if you use them, please be sure to cite the NHGIS. (Please cite me too!) The NHGIS has also asked that you not redistribute these converted files yourself, but instead point people here or to their own site. This is the citation you can use for the NHGIS:

> Minnesota Population Center. National Historical Geographic Information System: Version 2.0. Minneapolis, MN: University of Minnesota 2011. <http://www.nhgis.org>

You can load this boundary data into R using the [R.utils](http://cran.r-project.org/web/packages/R.utils/index.html) package, like so:

```r
library(R.utils)
states_1790 <- loadObject("path/to/us.state.1790.Rdata")
```
