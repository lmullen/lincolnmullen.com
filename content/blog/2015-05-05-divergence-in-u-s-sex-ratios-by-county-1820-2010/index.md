---
title: 'Divergence in U.S. Sex Ratios by County, 1820--2010'
date: "2015-05-05"
slug: 'divergence-in-u-s-sex-ratios-by-county-1820-2010'
description: "An interactive map of U.S. sex ratios by county from 1820 to 2010 reveals frontier imbalance, Civil War absences, and prison counties bulging with men since the 1990s."
tags:
- digital history
- mapping
- data visualization
draft: false
---

At a THATCamp New England last spring I re-purposed some earlier code to create a map of sex ratios by county from U.S. Census data. As far as I can remember, though, I never actually made the map public or wrote about it beyond showing it to a few people. I don't intend to do anything more with this data, so here is the interactive map on [Divergence in U.S. Sex Ratios by County, 1820--2010](https://projects.lincolnmullen.com/sex-ratios/).

{{< figure src="visualization-screenshot.png" caption="Screenshot of the interactive map of sex ratios in U.S. counties." >}}

Two general patterns are clear. First, heavily populated places tended to be slightly female, while lightly populated places[^1] tend to be heavily male. In any decade there were counties with more than 60% males; but there were only one or two counties with significant populations where the population was more than 60% female, and hardly any where the population was more than 55% female. It is easier to diverge from an equal number of males and females when the population is small.[^2] As an extreme example, Gaines County, Texas, in 1880 was 87.5% female, but only because its population as counted in the Census was seven females and one male. Thus, as the population increases, counties diverge less sharply from an even sex ratio though they do tend to have more females than males. You can see this same effect displayed a different way in this chart of sex ratios vs population.

{{< figure src="ratio-vs-population.png" caption="Sex ratios by population, U.S. counties 1820--2010. Data: NHGIS." >}}

The second pattern involves change over time and space. The spaces that tended to have more men than women were in the west. As late as 1950, most of the western counties were more than 51% male, with only a few counties in the east having populations that were more than 51% female. One exception was the 1870 census, where the absence of men in the states of the former confederacy can be seen on the map. From 1960 to the present most states with exceptions like Alaska, Wyoming, and North Dakota had majority female populations. But since the 1990s, the map shows isolated, lightly populated counties with heavily male populations: those are the counties with prisons.

[^1]: That is, places not populated by people counted in the Census. Most of the time the Census data does not take into account for Native American populations.

[^2]: Unlike these excellent [election maps](http://www.nytimes.com/interactive/2014/11/04/upshot/senate-maps.html) from the *New York Times*, this map makes no attempt to adjust the transparency of a county for its population.
