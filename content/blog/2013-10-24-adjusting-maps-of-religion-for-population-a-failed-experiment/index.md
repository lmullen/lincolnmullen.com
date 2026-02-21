---
title: 'Adjusting Maps of Religion for Population (A Failed Experiment?)'
date: "2013-10-24"
slug: 'adjusting-maps-of-religion-for-population-a-failed-experiment'
description: "Dividing county-level religion data by population to control for density—but in 1860 the population was more evenly distributed, so the adjustment changes less than expected."
tags:
- digital history
- mapping
- R
draft: false
---

XKCD [points out a common problem](http://xkcd.com/1138/) in mapping data:

{{< figure src="heatmap.png" caption="XKCD: Heatmap" >}}

If it's true that maps of data tend to mirror maps of population, then it strikes me that this is particularly a problem for historical maps. Even if I have a rough mental map of what parts of the United States are populated today, I certainly don't have an accurate mental map of the U.S. population at any given point in history. Does a map of any given historical data follow the population? You couldn't know unless you had mapped the population.

To try to solve this problem I plotted some county-level [religion data from the NHGIS](http://usreligion.blogspot.com/2013/10/historical-religion-data-in-nhgis-and.html). Then I plotted the same data divided by the population. If we're concerned that the main, hidden variable that is actually driving our map is population, then dividing by population should control for that effect. I'll show you the maps that I made first (code [here](https://github.com/lmullen/nhgis-explore/tree/2013-10-24-blog-post)), then try to figure out what `map = x /  population` means as an analysis tool.

Here is a map of the population of the eastern United States in 1860. Notice that the scale used in all this maps is a square-root scale, meaning that the difference between a populous city and a less-populous farming county is compressed so that the cities don't wash out the map. I've mapped only the eastern United States not because the West is unimportant, but because the West is so much less populous than the East that adjusting for population skews the maps.

{{< figure src="2013-10-24-us-population.png" >}}

The population in 1860 was much more spread out than it is today, because much many more people were farmers. We can observe that the Northeast is more populous than most of the rest of the country, and that the rest of the North is more populous than the South. But already we can make an inference: if the population was more evenly distributed than it is today, then variations in our map are less likely to have been caused by population. (To put this in mathematical terms, in our pseudo-equation `map = x / population`, as the population approaches being the same everywhere, then our denominator has no influence on the map. It's like dividing by 1.)

Now let's map the number of churches for six denominations in 1860. These denominations were more or less selected at random, except that I excluded Baptists and Methodists because they have so many more churches that they would raise the upper bound of the scale and so hide some of the detail.

{{< figure src="2013-10-24-churches-absolute.png" >}}

There is a wide geographic variation between these denominations. Already this tells us that the variation cannot be determined primarily by population. Briefly stated, this is what we can observe about each denomination:

-   Christian (i.e., churches from the Stone-Campbell movement) are located away from the eastern seaboard, especially in Ohio, Indiana, and Kentucky, but also with a larger than usual number of churches in Rhode Island and Cape Cod.
-   Episcopalians have their centers in Eastern cities, but also in all of New York State.
-   Friends are mostly located in eastern Pennsylvania, but also the Hudson River valley and northwest New York, with some migrations to Ohio and Indiana.
-   Moravians are very small compared to these other denominations.
-   Presbyterians are the most ubiquitous of these denominations, and they show up almost everywhere except New England (where no doubt they are replaced by their cousins the Congregationalists) but in extra strength in Pennsylvania, New York, and Ohio.
-   Catholics mostly live in the North (and if we had mapped it, in Canada) though with a presence in Louisiana.

Next we can map the same denominations, but dividing the number of churches by the population of the county.

{{< figure src="2013-10-24-churches-relative.png" >}}

It is difficult to express in words what `map = x / population` really means. Notice that we're dividing by the population in the county, not the number of adherents in the county, so we're not mapping the sizes of congregations. What I think we're mapping is the places where a denomination is doing better than we would expect given the population. Another, more interpretative way to put this is that these are maps of places where denominations had an edge on their competitors.

If that's what these maps mean, then we can make three observations from them.

1.  No denomination was doing better than expected in the cities. This makes sense, because it was hard to keep up with the population in cities. These maps don't show it, but we would also expect that it would be harder to gain a plurality in cities.

2.  The corollary is that some denominations did better than expected in non-urban counties. Again these maps don't show this, but we could expect that a denomination could more easily gain dominance in a place where there were fewer people and thus fewer churches. The Christian churches were large relative to population in what is now the Midwest. Episcopalians did well outside of the cities. Friends did poorly outside of eastern Pennsylvania, and Moravians never kept up with population anywhere. Presbyterians did well everywhere outside New England, and Catholics (surprisingly, given their urban reputation) did well in the North and Louisiana outside of the cities. These locations of the churches of these last two denominations seems to have been less dependent on population than the other denominations.

3.  The maps adjusted for population have far fewer hot spots than the map of absolute number of churches. (The places that are hot spots are clearly counties where the population is so low that having any number of churches is remarkable.) The evening out of the maps means that population was certainly one of main explanations for where religions had churches, and adjusting for it has removed some of the geographic variation. If population explained everything about the location of churches, then by adjusting for it we would expect a perfectly even map. To the extent that these maps are not even, population is not an explanation for the geography of religion. We might infer from this that it was difficult, but not impossible, for any denomination to gain a foothold in a county of any signficant population.

There are many ways that these maps could be improved. We would want to have maps for as many decades as possible, so that we can see if the population affects the maps more than it did in 1860. We can guess that as the population becomes less rural and more urban, that the effect of population will matter more. And since historians are interested in change over time, we would want to know how the observations we've made for these maps change for other other times. But we would also want these maps to be interactive and animated. We would want them to be interactive so that we can switch back and forth between the absolute number of churches and the number of churches relative to population. And we would want those transitions to be animated, so that our eye can be drawn to the differences. I'm in the midst of learning [D3.js](http://d3js.org/) so that I'll be able to make these kinds of interactive visualizations.

We can also think of shortcomings in the data that we are using. Ideally, we would want to map a variable like number of adherents that is more closely correlated to population than the variable (number of church buildings) that we are using. Different denominations have different preferences for the size of congregations, so a denomination could have more but smaller congregations and be about the same size as a denomination with fewer but bigger congregations. Furthermore, the number of churches does not even correlate to the number of congregations, since Baptists, say, might have one small congregation per church, but a Catholics might have many large congregations attending Mass each Sunday. (Even in 1865, [Catholics were complaining](http://quod.lib.umich.edu/m/moajrnl/bac8387.0001.001/5:3?rgn=full+text;view=image) that the census takers counted church buildings instead of communicants.)

There is another way to try to figure out whether population affects the location of religion, and that is to plot the number of churches as a function of the size of the county. That is what I've done in this chart, excluding the twenty most populous cities.

{{< figure src="2013-10-24-churches-charts.png" >}}

I don't have much to say about these charts, except to make these few observations. Every denomination tends to have most of its churches in counties with small populations, which is not surprising since the median county population is less than 10,000, and third quartile county population is about 17,000. Presbyterians tend to have more churches for any given population size, as do Episcopalians and Catholics. Perhaps these charts indicate that Episcopalians, Presbyterians, and Catholics are more affected by population, but I'll have to think more about it.

This experiment in adjusting maps for population is inconclusive. Perhaps there is some benefit to controlling for population, but the benefit is not immediately obvious. I'll want to try this technique with more kinds of data at different moments in time. If this isn't exactly a failed experiment, still it mostly proves the null hypothesis: that when the population is more evenly distributed, it's less important to control for population.
