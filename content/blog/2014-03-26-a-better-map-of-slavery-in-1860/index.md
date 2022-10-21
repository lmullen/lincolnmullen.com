---
title: A better map of slavery in 1860
date: '2014-03-26T02:59:00+00:00'
slug: 'a-better-map-of-slavery-in-1860'
tags:
- mapping
- slavery
- digital history
---

When I finished working the other night I [tweeted](https://twitter.com/lincolnmullen/status/448282708095934464) the current state of the map of slavery that I had been making. Anthea Butler [retweeted](https://twitter.com/AntheaButler/status/448282851055009792) it, and then a lot of people saw it. (Not *that* many, but certainly more than will ever read the dissertation chapter the map is a part of.) I'm glad that people found the map interesting. But though there was nothing erroneous about the map, it certainly was not the best map of slavery possible. Here is the draft map.

{{< figure src="twitter-map.png" caption="Number of slaves by county in 1860 (quartile breaks)." >}}

It's easy to spot the biggest problem in that map: the values mapped to the colors are less than ideal. I suspect that most people who saw the map didn't pay any attention to the legend at the bottom. And why should they have? Until I changed the numbers to a humanist-readable format the legend was almost incomprehensible. What the legend means is that the lightest yellow represents counties where there were 450 or fewer slaves living; the dark red represents counties where there were more than 5,380 slaves and fewer than 37,300 slaves.

Those numbers should give a reader pause: why should a county with 5,380 slaves be classified the same as a county with almost seven times as many slaves? The breaks in the first map are not arbitrary, but divide the counties into [quartiles](http://en.wikipedia.org/wiki/Quartile). That is, I ran a function which divided the counties into four even groups. This was a rough and ready way to classify the counties.

The trouble is that quartiles are not a particularly meaningful way to classify the counties. You might even argue---and this certainly wasn't my intent---that it is a sensationalist way to classify the counties. By definition, using quartiles means that one-quarter of the counties on the map would be colored bright red. If this were a map of smokers, then one-fourth of the counties would be bright red; if it were a map of lung cancer, one-fourth of the counties would be bright red. That's because when using quartiles, the breaks are determined by the count of the observations (i.e., the number of counties) rather than the value of the observations (i.e., the number of slaves in each county). Below is a histogram of the distribution of counties: you can see that a few counties had very large numbers of slaves, while most counties had relatively smaller numbers.

{{< figure src="histogram.png" caption="Histogram of the number of slaves per county." >}}

But the question of how to categorize the counties is as much a historical question as it is a question for the techniques of data analysis. Though histories of slavery have often been written about large plantations where many slaves lived, historians have long known that many enslaved African Americans did not live on plantations, because most slaveholders owned only a few slaves. This is an important point, because the possibilities for slave culture and religion are very different on a farm with one or two enslaved African Americans than on a plantation with a hundred slaves. Below is a chart of the number of slaveholders by the number of slaves that they owned.[^1] (Notice also how widespread slave ownership was: 395,216 slaveholders according to the 1860 census.)

[^1]: In this case the categories come directly from the Census tables. As some people wrote to say, the proportion of African Americans in the total population is another way to measure this, but that's the subject of another map. 

{{< figure src="holders_1860.webp" caption="Slaveholders by number of slaves owned in 1860." >}}

What the charts of counties and slaveholders demonstrate is that dividing the counties into quartiles does not make for an accurate map. Fortunately there are better methods, in particular [George Jenks's algorithm](http://en.wikipedia.org/wiki/Jenks_natural_breaks_optimization) for finding breaks in the data set. The [Jenks method](http://support.esri.com/en/knowledgebase/techarticles/detail/26442) tries to make groups whose individual members are as close to each other as possible, but where each group in the aggregate is as much unlike the other groups as possible. Using that algorithm, we can divide the counties into more meaningful groups, as the chart below shows.

{{< figure src="histogram_breaks.png" caption="Histogram of slaves per county with breaks compared." >}}

Using the Jenks breaks, we get a much better map of where slaves lived in 1860. We can see all of the detail that was in the earlier map: the South's [fertile crescent](http://www.npr.org/blogs/krulwich/2012/10/02/162163801/obama-s-secret-weapon-in-the-south-small-dead-but-still-kickin) through Virginia, North Carolina, South Carolina, George, Alabama, and Mississippi; the Mississippi, Missouri, and Tennessee river valleys; South Carolina as the state with the highest concentration of slaves. But this revised map has a higher resolution (if you will). We can now see cities like Washington, Charleston, Nashville, Mobile, and New Orleans---important since slavery must be understood in terms of slave markets, commodities markets, and capitalism.[^2]  The hinterland of slavery is also more clearly defined---important since the expansion of slavery was the issue in the sectional crisis.[^3]

[^2]: See Walter Johnson, _Soul by Soul: Life Inside the Antebellum Slave Market_ (Harvard University Press, 1999); Walter Johnson, _River of Dark Dreams: Slavery and Empire in the Cotton Kingdom_ (Harvard University Press, 2013).

[^3]: The code and data for the map is [on GitHub](https://github.com/lmullen/slavery-nhgis).

{{< figure src="slaves_1860.webp" caption="Number of slaves by county in 1860 (Jenks breaks)." >}}

The lesson here is not that you should only make finished work public. But I hope that this look at the decisions that go into working with data demonstrates how a historian's knowledge is more important than technological skill in making a historical map.
