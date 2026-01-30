---
title: 'Maps of Catholic Dioceses in the US, Canada, and Mexico, Take Two'
date: "2013-09-08"
slug: 'maps-of-catholic-dioceses-in-the-us-canada-and-mexico-take-two'
tags:
- digital history
- mapping
- American religious history
draft: false
---

In an [earlier post](/blog/mapping-catholic-dioceses-over-time/), I shared some maps of Catholic dioceses in the United States, Canada, and Mexico over time. My R-fu has since gotten stronger. In particular, I've learned [how to use shapefiles](/blog/shapefiles-of-historic-us-state-boundaries-for-use-in/). While the R packages `ggmaps` is good when one wants a physical or street map, it's not capable of displaying changes in political boundaries, and it forces you to use a square aspect ratio. In these updated maps of Catholic dioceses, I've used shapefiles of historic state boundaries from the [NHGIS](https://www.nhgis.org/) to give a better view of where new dioceses fit into the United States. I've also animated the maps rather than display them separately. The [data](https://github.com/lmullen/demographics-religion/blob/master/data/clean/catholic.dioceses.geocoded.csv) along with the [R code](https://github.com/lmullen/demographics-religion/blob/master/catholic.dioceses.r) for making the maps is available on GitHub.

Red dot are the seats of dioceses; purple dots are the seats of archdioceses. The data includes both Roman rite and various eastern rite dioceses. I've only displayed dioceses in Mexico and Canada that are close to the border within the boundaries of the map, though the data has all the Canadian and Mexican dioceses.

{{< figure src="dioceses.animation.gif" >}}
