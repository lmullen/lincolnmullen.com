---
title: 'Historical Religion Data in the NHGIS and What You Can Do with It'
date: "2013-10-01"
slug: 'historical-religion-data-in-the-nhgis-and-what-you-can-do-with-it'
tags:
- digital history
- mapping
- American religious history
draft: false
crosspost:
  url: "http://usreligion.blogspot.com/2013/10/historical-religion-data-in-nhgis-and.html"
  source: "[Religion in American History](http://usreligion.blogspot.com/)"
---

In 1865, the French observer of American Catholicism E. Rameur opened the inaugural issue of the *Catholic World* with an article on "[The Progress of the Church in the United States](http://quod.lib.umich.edu/m/moajrnl/bac8387.0001.001/5:3?rgn=full+text;view=image)." In assessing the current population of the church, Rameur's task was very difficult, as he pointed out: "The number of the faithful it is not easy to determine accurately; for a false delicacy prevents the Americans from including the statistics of religious belief in their census-tables. Estimates are very variable." Rameur was right: in the United States, unlike in many other countries, the census does not track religious data. The omission is justified by the separation of church and state, but every religious historian I know would gladly suffer a breach in the wall of separation for the sake of reliable data for the history of religion.

{{< figure src="2013-10-01-rameur.png" >}}

But Rameur understated the amount of data about religion in the U.S. census. By the time he wrote, the 1850 and 1860 census had gathered some data about churches, though not about their membership. (Rameur dismissed this data, but historians can't afford to be so discriminating.) Gaustad and Barlow summarize what is available in the census for their note on the sources for their standard work, *[The New Historical Atlas of American Religion](https://www.amazon.com/dp/B0044KMUNC/ref=as_li_ss_til?tag=thebacgla-20&camp=0&creative=0&linkCode=as4&creativeASIN=B0044KMUNC&adid=1VWHJBJNDXA101GHTNYP&)*:

> Although the U.S. Census began in 1790, compilation of religious data did not begin until 1850. From 1850 to 1880, the government gathered facts concerning the number of churches, their locations, and their seating accommodations. It included data on membership for the first time in 1890, and in the years 1906, 1916, 1926, and 1936 devised elaborate tables and published separate volumes dealing with all phases of church life that gave themselves to statistical representation. Since 1936---and unlike, for instance, Canada---the U.S. government has undertaken no federal census of religious bodies, largely on the (debatable) grounds that this would compromise civil liberty and church-state separation (395).

That's slim pickings compared to what might have been. But if you are a religious historian working from the second half of the nineteenth century to the first half of the twentieth century, there is at least some data available to you.

Fortunately this census data is easily accessible through the [National Historical Geographic Information System](https://www.nhgis.org/), a scholarly project of the Minnesota Population Center at the University of Minnesota. (The same center also puts out the North Atlantic Population Project and [other useful data](http://www.ipums.org/).) The NHGIS has digitized the census data into a format that is usable in a statistical program [like R](http://www.r-project.org/) or SAS, a GIS program like [QGIS](http://www.qgis.org/en/site/) or ArcGIS, or just plain Excel. You can browse through their site and select the data by topic and year. All told, there are 63 tables of information about religion in the NHGIS. You can download this data at three different levels of detail: national, state, and county.

{{< figure src="2013-10-01-nhgis-downloads.png" >}}

The NHGIS offers two kinds of data. The first is the boundary lines of states, territories, and counties for each decade since 1790. (These are called shapefiles.) Plotting nineteenth-century data onto a twenty-first century map is an anachronism; these shapefiles fix that. In the animation below of Catholic dioceses since 1790, you can watch the boundaries of the states change over time.

{{< figure src="dioceses.animation.gif" >}}

The second kind of data is the tables of quantitative information themselves. You can open these in a spreadsheet program to look up particular data: for example, the number of Catholic churches in Albany County in 1850 (answer: six). Unfortunately the NHGIS does not offer time series data for religion; that is to say, they have not linked together data on religious topics across census years. But if you want to make comparisons across time, you can do that on your own if somewhat laboriously.

{{< figure src="2013-10-01-spreadsheet.png" >}}

The real power of the NHGIS data comes when you start combining quantitative data with spatial data. (This isn't the place to explain how to do that, but you can look at Fred Gibbs's helpful tutorial on "[Historical Mapping with QGIS](http://fredgibbs.net/historical-mapping-with-qgis/)" if you want to try it for yourself.) Gaustad and Barlow offer beautiful maps of county level data for each denomination, usually for the years 1850, 1890, 1950, and 1990. They were limited to two years in the nineteenth-century because of the economics of print publication.

{{< figure src="2013-10-01-gaustad.JPG" >}}

We can reproduce these maps for ourselves for any year for which we have data. Here are maps of the number of Catholic churches in 1850, 1860, and 1870.

{{< figure src="2013-10-01-catholics-1850.png" >}}

{{< figure src="2013-10-01-catholics-1860.png" >}}

{{< figure src="2013-10-01-catholics-1870.png" >}}

There is a lot to criticize about these maps, which I made in a couple hours. To begin with, there is the problem of attaching the quantitative data to the shape of the county, which distorts the data we are actually interested in (the number of churches) according to the accidents of political boundaries. In these maps, for example, New York City is barely visible, though it was a primary center of American Catholicism, but some of the huge boundaries in the west are over represented.

But we've also corrected one the major problems with the Gaustad and Barlow maps. These kinds of maps (called [choropleths](http://en.wikipedia.org/wiki/Choropleth_map)) rely on changing the color within a geographic boundary according to changes in the data. The maps above use the same scale for each year, so the darkest red represents more than 40 churches per county, and the medium red represents about 20 churches per county. Gaustad and Barlow, however, almost always change the scale of their maps from year to year. For example, their map of Catholic churches in 1850 (based on the same data as above) uses the darkest color for counties that have over 10 churches. But the map for the 1890 uses that same color for counties with more than 20 churches, the map for the 1950s uses that color for counties with more than 30 churches, and the 1950s (bizarrely) uses that color for counties with more than 10 churches. Gaustad and Barlow were simply picking the best scale for the data in each year. But this deceives your eye when you make comparisons across years. You can get the impression that the number of Catholic churches shrinks in some places from 1850 to 1890, which was not the case; but you can also gain the impression that Catholic churches are everywhere in large numbers by the 1950s, but this is mostly an effect of the change in scale. This problem is more obvious but still difficult to overcome when comparing the changing scales of various denominations.

I don't mean to be overly critical of Gaustad and Barlow's atlas, which [is useful](http://usreligion.blogspot.com/2013/07/atlases-of-american-religion-print-and.html) as a final *presentation* of this historical data. I keep it always within arm's reach. But what is more useful for research is the *exploration* of data. When you dig into the data for yourself, you are able to make as many maps as you like. If you're interested in the interactions between Catholics and Episcopalians, as I am, then you can generate a map of Episcopalian churches in 1850 using the same scale as the Catholic map.

{{< figure src="2013-10-01-episcopalians-1850.png" >}}

You can also layer the two for a direct comparison, as in the map below (more red means more Catholic dominance, more blue means more Episcopal dominance, and more purple means they both have a strong presence). The prevalence of purple tells me that Catholics and Episcopalians were in direct competition in many places, but that Episcopalians were more widespread than Catholics, and that Catholics had the upper hand around the Canadian border.

{{< figure src="2013-10-01-catholics-episcopalians.png" >}}

This last map is particularly ugly and unsuitable as a presentation version. But it is illuminating as part of an iterative, adductive process of asking questions and trying out answers. That process, after all, is the fundamental way that historians work. You'll also notice that in the last map we've zoomed to the places that interest us most. If we wished we could zoom to finest level of detail that our data supports, instead of being stuck at the national level.

Finally, the real power of getting into data for yourself comes out when we start mixing data sets together. These maps would all be better if I got the equivalent data from Canada, since many people crossed the borders. Most recently I was able to combine data in a useful way for a chapter that I just finished. I'm in the process of cataloging all of the missions given by the Paulist Fathers. When I plotted their missions for a decade and a half after the Civil War, I got this map.

{{< figure src="2013-10-01-paulist-missions.png" >}}

On its own that map tells me something about the spread of Paulist missions and conversions. But I wondered why the Paulists went to the places that they did. Thanks to William Thomas and his team for *[Railroads and the Making of Modern America](http://railroads.unl.edu/)*, I was able to plot these missions on top of a map of railroads in 1870.

{{< figure src="2013-10-01-paulist-missions-rr.png" >}}

Using the railroad data explains a great deal about where the Paulists went. Except for staying out of most of the South, they went anywhere the railroads took them.

Religious historians, I think, are naturally suspicious of "data" because we are mostly humanists rather than social scientists. Even if we are interested, the field has mostly suffered from a lack of good data before, say, the middle of the twentieth century. But my experience is that there are many more sources amenable to mapping and quantification than we have assumed, even if getting the data out of the sources requires many hours of transcribing manuscripts into spreadsheets, and if using it requires picking up some tools that have not typically been the province of religious historians. If that kind of approach appeals to you, then the NHGIS census data is a good place to start.
