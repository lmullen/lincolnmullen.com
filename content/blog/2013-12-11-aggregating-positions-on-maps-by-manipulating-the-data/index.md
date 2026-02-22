---
title: 'Aggregating positions on maps by manipulating the data'
date: "2013-12-11"
slug: 'aggregating-positions-on-maps-by-manipulating-the-data'
description: "Using R's plyr package to aggregate overlapping map points by city and year, turning Paulist mission records into clear, additive clusters instead of stacked dots."
tags:
- digital history
- R
- mapping
draft: false
---

I've recently been [mapping the missions of the Paulist Fathers](http://usreligion.blogspot.com/2013/10/historical-religion-data-in-nhgis-and.html) over the course of the nineteenth century. One problem with the data is that many of the points overlap with one another, since the Paulists were often in cities like New York, Philadelphia, Chicago, and Baltimore. When mapping these points, they overlap with one another.

This is a common problem in mapping, which [Leaflet.js](http://leafletjs.com/) solves admirably for web maps. See, for example, the [DPLA's map of items](http://dp.la/map). Another solution is to make the points transparent, so that overlapping points are darker than one another. While a judicious use of transparency can help in some places, it is generally a poor solution. It's hard to explain what the layered colors mean to users, and the eye is poor at detecting the difference anyway.

What I wanted to do is to sum together the overlapping points. So for example, instead of having a data file with a mission at St. Peter's Church in New York in September 1851 with 4,000 confessions, and another mission at St. Patrick's Cathedral in October 1851 with 7,000 confessions, I wanted a data file where those points are aggregated as New York missions in 1851 with 11,000 total confessions and 2 missions.

This task is a simple job for Hadley Wickham's [plyr](http://plyr.had.co.nz/) package in R.

```r
library(plyr)

raw <- read.csv("data/paulist-chronicles/paulist-missions.geocoded.csv")

aggregated <- ddply(raw, .(city, state, year), summarize,
                    long = max(geo.lon),
                    lat =  max(geo.lat),
                    confessions = sum(confessions_total, na.rm = TRUE),
                    converts = sum(converts_total, na.rm = TRUE),
                    number_missions = length(location))

write.csv(aggregated,
          "data/paulist-chronicles/paulist-missions.aggregated.csv",
          row.names = FALSE)
```

First we load the package and read in the raw data. One function, `ddply`, will accomplish our work for us. The first argument passes our data to the function. The second argument, `.(city, state, year)`, splits up our original data by finding all the unique combinations of the `city`, `state`, and `year` variables. In other words, `ddply` makes a new data frame for each combination, such as New York missions in 1851, New York missions in 1852, Chicago missions in 1851, Chicago missions in 1852, and so on. (We could leave out the `year` variable if we wanted to aggregate the missions just by place, not the combination of time and place.) Then `ddply` applies the `summarize` function to each of those split-up data frames. Essentially we're deciding on the columns for our new aggregated data frame. For `long` and `lat`, we're taking the maximum value for, say, each New York mission, but each New York mission should have the same latitude and longitude so it doesn't matter. For `confessions` and `converts` we are summing up the total. Then for the `number_missions` we count the `length` of the split-up data frame (in other words, the total number of observations that have the unique combination of `city`, `state`, and `year`).

You can see what this does by looking at the [original data](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/paulist-missions.geocoded.csv), which goes through [this script](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/aggregate-missions.r) to produce this [aggregated data](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/paulist-missions.aggregated.csv). I won't belabor the explanation, since Wickham explains it better than I can in his article "[The Split-Apply-Combine Strategy for Data Analysis](http://www.jstatsoft.org/v40/i01)."[^1] What I do want to point out is that this is a useful technique for making maps with overlapping points. And even if you are making the maps outside of R, perhaps in GIS software or in D3, R and plyr can still be powerful tools to get your data into the proper format.

[^1]: Hadley Wickham, "The Split-Apply-Combine Strategy for Data Analysis," *Journal of Statistical Software* 40, no. 1 (April 2011): 1--29, <http://www.jstatsoft.org/v40/i01>.
