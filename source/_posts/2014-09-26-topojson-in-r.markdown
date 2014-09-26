---
categories:
- 'r-stats'
date: '2014-09-26 11:58'
layout: post
link: 'http://rud.is/b/2014/09/26/overcoming-d3-cartographic-envy-with-r-ggplot/'
title: TopoJSON in R with rgdal
...

In making maps in R, the biggest point of pain for me has been reading
in shapefiles. D3 has been far better for making maps because
[TopoJSON's][] topological format for representing geometries is so much
better than the format of shapefiles. Bob Rudis's post "[Overcoming D3
Cartographic Envy With R + ggplot][]" taught me that the most recent
version of [GDAL/OGR][] can read in TopoJSON, and thus that the R
package [rgdal][] can as well. This will make mapping in R much easier.

  [TopoJSON's]: https://github.com/mbostock/topojson
  [Overcoming D3 Cartographic Envy With R + ggplot]: http://rud.is/b/2014/09/26/overcoming-d3-cartographic-envy-with-r-ggplot/
  [GDAL/OGR]: http://www.gdal.org/
  [rgdal]: http://cran.rstudio.org/web/packages/rgdal/
