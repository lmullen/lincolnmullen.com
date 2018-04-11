---
date: '2013-12-11T21:21:00+00:00'
slug: 'aggregating-positions-on-maps-by-manipulating-the-data'
title: Aggregating Positions on Maps by Manipulating the Data
---

I've recently been [mapping the missions of the Paulist Fathers](http://usreligion.blogspot.com/2013/10/historical-religion-data-in-nhgis-and.html) over the course of the nineteenth century. One problem with the data is that many of the points overlap with one another, since the Paulists were often in cities like New York, Philadelphia, Chicago, and Baltimore. When mapping these points, they overlap with one another.

This is a common problem in mapping, which [Leaflet.js](http://leafletjs.com/) solves admirably for web maps. See, for example, the [DPLA's map of items](http://dp.la/map). Another solution is to make the points transparent, so that overlapping points are darker than one another. While a judicious use of transparency can help in some places, it is generally a poor solution. It's hard to explain what the layered colors mean to users, and the eye is poor at detecting the difference anyway.

What I wanted to do is to sum together the overlapping points. So for example, instead of having a data file with a mission at St. Peter's Church in New York in September 1851 with 4,000 confessions, and another mission at St. Patrick's Cathedral in October 1851 with 7,000 confessions, I wanted a data file where those points are aggregated as New York missions in 1851 with 11,000 total confessions and 2 missions.

This task is a simple job for Hadley Wickham's [plyr](http://plyr.had.co.nz/) package in R.
<figure class="highlight">
<pre><code class="language-r" data-lang="r">&lt;span class="n">library&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">plyr&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="n">raw&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="n">read.csv&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"data/paulist-chronicles/paulist-missions.geocoded.csv"&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="n">aggregated&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="n">ddply&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">raw&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="err">.&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">city&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">state&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">year&lt;/span>&lt;span class="p">),&lt;/span> &lt;span class="n">summarize&lt;/span>&lt;span class="p">,&lt;/span>
                    &lt;span class="n">long&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">max&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">geo.lon&lt;/span>&lt;span class="p">),&lt;/span>
                    &lt;span class="n">lat&lt;/span> &lt;span class="o">=&lt;/span>  &lt;span class="n">max&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">geo.lat&lt;/span>&lt;span class="p">),&lt;/span>
                    &lt;span class="n">confessions&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">sum&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">confessions_total&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">na.rm&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">TRUE&lt;/span>&lt;span class="p">),&lt;/span>
                    &lt;span class="n">converts&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">sum&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">converts_total&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="n">na.rm&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">TRUE&lt;/span>&lt;span class="p">),&lt;/span>
                    &lt;span class="n">number_missions&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">length&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">location&lt;/span>&lt;span class="p">))&lt;/span> 

&lt;span class="n">write.csv&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">aggregated&lt;/span>&lt;span class="p">,&lt;/span> 
          &lt;span class="s2">"data/paulist-chronicles/paulist-missions.aggregated.csv"&lt;/span>&lt;span class="p">,&lt;/span>
          &lt;span class="n">row.names&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">FALSE&lt;/span>&lt;span class="p">)&lt;/span></code></pre>
</figure>
First we load the package and read in the raw data. One function, `ddply`, will accomplish our work for us. The first argument passes our data to the function. The second argument, `.(city, state, year)`, splits up our original data by finding all the unique combinations of the `city`, `state`, and `year` variables. In other words, `ddply` makes a new data frame for each combination, such as New York missions in 1851, New York missions in 1852, Chicago missions in 1851, Chicago missions in 1852, and so on. (We could leave out the `year` variable if we wanted to aggregate the missions just by place, not the combination of time and place.) Then `ddply` applies the `summarize` function to each of those split-up data frames. Essentially we're deciding on the columns for our new aggregated data frame. For `long` and `lat`, we're taking the maximum value for, say, each New York mission, but each New York mission should have the same latitude and longitude so it doesn't matter. For `confessions` and `converts` we are summing up the total. Then for the `number_missions` we count the `length` of the split-up data frame (in other words, the total number of observations that have the unique combination of `city`, `state`, and `year`).

You can see what this does by looking at the [original data](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/paulist-missions.geocoded.csv), which goes through [this script](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/aggregate-missions.r) to produce this [aggregated data](https://github.com/lmullen/demographics-religion/blob/master/data/paulist-chronicles/paulist-missions.aggregated.csv). I won't belabor the explanation, since Wickham explains it better than I can in his article "[The Split-Apply-Combine Strategy for Data Analysis](http://www.jstatsoft.org/v40/i01)." [<sup>1</sup>](#fn1){\#fnref1.footnoteRef} What I do want to point out is that this is a useful technique for making maps with overlapping points. And even if you are making the maps outside of R, perhaps in GIS software or in D3, R and plyr can still be powerful tools to get your data into the proper format.
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
<p>
Hadley Wickham, "The Split-Apply-Combine Strategy for Data Analysis," <em>Journal of Statistical Software</em> 40, no. 1 (April 2011): 1--29, <a href="http://www.jstatsoft.org/v40/i01" class="uri">http://www.jstatsoft.org/v40/i01</a>.<a href="#fnref1">↩</a>
</p>
</li>
</section>
