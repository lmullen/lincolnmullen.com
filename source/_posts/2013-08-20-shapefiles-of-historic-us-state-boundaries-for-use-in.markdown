---
layout: post
title: "Shapefiles of Historic U.S. State Boundaries for Use in R"
date: 2013-08-20 22:21
categories: 
- R-language
- GIS
- spatial-humanities
- digital-humanities
- history
---

In the past few months I've picked up [the R language][] for digital
analysis of historical sources. It's one of the most powerful digital
humanities tools that I've learned, because it's useful for so many
tasks, especially quantitative analysis and mapping.

While there are several mapping packages for R, including the excellent
[ggmap][], the boundaries that these packages come with are all
contemporary boundaries. While it's acceptable to plot data on physical
map of the United States, as I did for my [maps of Catholic dioceses][]
in North America, it's far more interesting to be able to include the
changing political boundaries of the United States.

The data for those historic state boundaries are available from the
[National Historical Geographic Information System][], which "provides,
free of charge, aggregate census data and GIS-compatible boundary files
for the United States between 1790 and 2011." The NHGIS has an
extraordinarily useful collection of historical data, including a fair
amount of data on religion, which I'll write about later. Among their
offerings are shapefiles for state boundaries by decade, beginning in
1790. These shapefiles are easy to use in a GIS program like QGIS.

The shapefiles don't, however, work out of the box for R. The shapefiles
are associated with a particular map projection, and to work with
[ggmap][] and [ggplot2][], they need to be converted to the more common
WGS84 projection. This conversion is much too time consuming to be
performed on-the-fly. The shapefiles also need to be converted to a data
frame to be plotted. For doing work in R, the most convenient form this
boundary data could take is as an R object for easy loading.

I've written an [R script to convert these shapefiles][] to R objects in
the correct projection as part of my project on the [historical
demography][] of American religions. You can download the shapefiles
from the NHGIS and use the script perform the conversion yourself. But
since this is non-trivial both in terms of the required dependencies
(see the script's documentation) and in the amount of time it takes to
process the files, you can download the already converted boundary data:

-   [historic-state-boundaries.zip][] (500 MB)

The NHGIS has kindly granted me permission to share these converted 
files, but if you use them, please be sure to [cite][] the NHGIS.  
(Please cite me too!) The NHGIS has also asked that you not redistribute 
these converted files yourself, but instead point people here or to 
their own site. This is the citation you can use for the NHGIS:

> Minnesota Population Center. National Historical Geographic
> Information System: Version 2.0. Minneapolis, MN: University of
> Minnesota 2011. <http://www.nhgis.org>

You can load this boundary data into R using the [R.utils][] package,
like so:

{% highlight r %}
library(R.utils)
states_1790 <- loadObject("path/to/us.state.1790.Rdata")
{% endhighlight %}

Then you can plot the boundaries using this minimal example:

{% highlight r %}
library(ggplot2)
ggplot() +
geom_polygon(data = states_1790, fill = 'gray', color = 'blue', alpha=0.3,
             aes(x = long, y = lat, group = group))
{% endhighlight %}

The result can be a map like this one of the United States in 1790,
1830, and 1860 ([code][]):

{% img center /downloads/demographics/demo.state-boundaries.png %}

If you're using R for historical mapping, please try out this data and 
let me know if you have any problems. Thanks to the NHGIS for providing 
the data in the first place.

  [code]: https://github.com/lmullen/demographics-religion/blob/master/shapefiles.demo.R
  [National Historical Geographic Information System]: https://www.nhgis.org/
  [R script to convert these shapefiles]: https://github.com/lmullen/demographics-religion/blob/master/historic.shapefiles.R
  [historical demography]: https://github.com/lmullen/demographics-religion/
  [cite]: https://www.nhgis.org/research/citation
  [R.utils]: http://cran.r-project.org/web/packages/R.utils/index.html
  [the R language]: http://www.r-project.org/
  [historic-state-boundaries.zip]: https://www.dropbox.com/s/xu0bug1x7hgm5h1/historic-state-boundaries.zip
  [ggmap]: https://sites.google.com/site/davidkahle/ggmap
  [ggplot2]: http://ggplot2.org/
  [maps of Catholic dioceses]: http://lincolnmullen.com/blog/mapping-catholic-dioceses-over-time/

