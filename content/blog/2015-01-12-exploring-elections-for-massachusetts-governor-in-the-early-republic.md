---
date: '2015-01-12T00:35:00+00:00'
slug: 'exploring-elections-for-massachusetts-governor-in-the-early-republic'
title: Exploring Elections for Massachusetts Governor in the Early Republic
---

The [New Nation Votes](http://elections.lib.tufts.edu/) database (NNV) offers election returns from the early American republic collected by Philip Lampi and digitized by Tufts University and the American Antiquarian Society. Several scholars writing in a 2013 issue of the *Journal of the Early Republic* have tackled questions such as voter turnout and measures of party competitiveness (Brooke), the resurgence of the Federalists after 1808 (Lampi), the expansion of the franchise (Ratcliffe), and families and the turnover of congressmen (Zagarri). My aim is much more preliminary: to see what kind of analysis, in particular mapping, might be done with the dataset.[<sup>1</sup>](#fn1){\#fnref1.footnoteRef} I have wanted to explore this dataset for some time, so here is a preliminary investigation into the Massachusetts gubernatorial elections up to 1824.

The first aim is to get an overview of party politics in the elections for governor. The chart below shows the percentage of votes won by the Federalist and Democratic-Republican parties from 1796 to 1824. The overall pattern in elections for governors is fairly plain. From 1797 until 1805, the Federalists had a strong hold on the office, putting Increase Sumner, Moses Gill, and Caleb Strong in the governor's chair.[<sup>2</sup>](#fn2){\#fnref2.footnoteRef} Caleb Strong came to office (after both Sumner and Gill had died in office) with some competition in 1800, but his hold was fairly secure until 1805. That year inaugurated stiff competition for the governorship, which switched hands repeatedly until the election of 1813. The War of 1812 and the rise of younger Federalists gave the Federalists the upper hand until they lost the 1823 election, never to win the Massachusetts governorship again. Note that there are a few oddities in this chart which I have not resolved. For example, John Brooks was listed as a Federalist every year from 1816 to 1821, except 1818 to 1819; I don't know whether that means Brooks really ran without an affiliation or whether it is an omission in the data. But this chart more or less confirms the argument of Philip Lampi (and earlier, of David Hackett Fischer).
<figure id="figure-1">
<a onclick="ga('send', 'event', { 'eventCategory': 'Figure', 'eventAction': 'View', 'eventLabel': 'ma-gov/ma-parties'});" href='/figures/ma-gov/ma-parties.svg'><img src='/figures/ma-gov/ma-parties.svg' alt='Percentage of the vote won by the Federalist and Democratic-Republican parties in elections for Massachusetts governor, 1796--1824. Data: \[A New Nation Votes\](http://elections.lib.tufts.edu/).'></a>
<figcaption>
Figure 1: Percentage of the vote won by the Federalist and Democratic-Republican parties in elections for Massachusetts governor, 1796--1824. Data:

[A New Nation Votes](http://elections.lib.tufts.edu/). \[<a onclick="ga(&#39;send&#39;, &#39;event&#39;, { &#39;eventCategory&#39;: &#39;Figure&#39;, &#39;eventAction&#39;: &#39;View&#39;, &#39;eventLabel&#39;: &#39;ma-gov/ma-parties&#39;});" href="/figures/ma-gov/ma-parties.svg">SVG</a>, <a onclick="ga(&#39;send&#39;, &#39;event&#39;, { &#39;eventCategory&#39;: &#39;Figure&#39;, &#39;eventAction&#39;: &#39;View&#39;, &#39;eventLabel&#39;: &#39;ma-gov/ma-parties&#39;});" href="/figures/ma-gov/ma-parties.png">PNG</a>\]

</figcaption>
</figure>
<!--more-->
The next task is to see which politicians were serious contenders for governor of Massachusetts. I'm arbitrarily defining a contender as someone who managed to win at least 10 percent of the vote in at least one election. More than one thousand people are listed in the gubernatorial election returns from 1787 to 1824, but only twenty got at least 10% of the vote, and only ten won office in their own right.[<sup>3</sup>](#fn3){\#fnref3.footnoteRef} The chart below shows the careers of those contenders. There are too many people on this chart for the colors to be much help, so I've labeled the lines for the more significant figures. John Hancock had a secure tenure, while Samuel Adams's was somewhat more rocky. But note that Increase Sumner, the first Federalist governor, also won high proportions of the popular vote. Starting in 1800 and certainly by 1805, elections were contested much more heavily; the nature of gubernatorial politics changed. We can see the arcs of people's careers. Federalist Caleb Strong won a close election in 1800 and gradually increased his margin of victory, but with the new regime of competition he was an on-again, off-again candidate until 1815. Federalist John Brooks enjoyed seven wins in a row, but his last three election were contested by William Eustis. Eustis never did defeat Brooks, but he did defeat Federalists Harrison Gray Otis in 1823 and Samuel Lathrop in 1824. Elbridge Gerry was the William Jennings Bryan of the early republic (except Gerry eventually won), running repeatedly for governor starting in 1788, but not even coming close until he won in 1810.
<figure id="figure-2">
<a onclick="ga('send', 'event', { 'eventCategory': 'Figure', 'eventAction': 'View', 'eventLabel': 'ma-gov/ma-gov'});" href='/figures/ma-gov/ma-gov.svg'><img src='/figures/ma-gov/ma-gov.svg' alt='Contenders in elections for Massachusetts governor who won more than 10% of the vote, 1787--1824. Data: \[A New Nation Votes\](http://elections.lib.tufts.edu/).'></a>
<figcaption>
Figure 2: Contenders in elections for Massachusetts governor who won more than 10% of the vote, 1787--1824. Data:

[A New Nation Votes](http://elections.lib.tufts.edu/). \[<a onclick="ga(&#39;send&#39;, &#39;event&#39;, { &#39;eventCategory&#39;: &#39;Figure&#39;, &#39;eventAction&#39;: &#39;View&#39;, &#39;eventLabel&#39;: &#39;ma-gov/ma-gov&#39;});" href="/figures/ma-gov/ma-gov.svg">SVG</a>\]

</figcaption>
</figure>
Next, I've created maps for three gubernatorial elections: 1800, 1807, and 1823. These maps are exceedingly rough-and-ready, intended for exploration rather than argument. I made them using my [cartographer package](https://github.com/lmullen/cartographer). The election returns are in the NNV dataset at the level of towns, so I geocoded the names of 869 towns in Massachusetts and Maine.[<sup>4</sup>](#fn4){\#fnref4.footnoteRef} This is not ideal, but since towns tend to split rather than move the locations should be more or less correct for these exploratory maps. The county boundaries are from the [Atlas of Historical County Boundaries](http://publications.newberry.org/ahcbp/) via my [USAboundaries package](https://github.com/ropensci/USAboundaries). The maps each include layers for the top two or three candidates. Red represents the Federalists; blue, the Democratic-Republicans. Each town is sized according to the number of votes for the candidate.[<sup>5</sup>](#fn5){\#fnref5.footnoteRef} Click on each town to get the number of votes. This way of layering the votes for each candidate is not ideal. Perhaps a better solution would be to show how many more votes each candidate won in a particular place; e.g., Strong won 142 more votes than Gerry in Brookfield.

Some general observations about the importance of space in these elections.

First, Boston was far and away the biggest city in Massachusetts, but it had little impact on the elections. In the 1800 election, Gerry got only 24 more votes than Strong in Boston, a difference of less than 1 percent of the turn out. In 1823, Otis got only 108 more votes than Eustis. Only in 1807 did Strong get significantly more votes than Sullivan (and Strong still lost the election). Even though Boston contributed more votes than any place, and though sometimes it went for Democratic-Republicans and sometimes it went for Federalists, it was not really a swing city because the two parties were usually closely tied in Boston.

Second, in the 1800 election Strong won because he won Berkshire and Hampshire Counties in Western Massachusetts. Gerry's support in those counties was virtually non-existent. [<sup>6</sup>](#fn6){\#fnref6.footnoteRef} Gerry, though, did much better in Maine, especially away from the coast. Strong also did well in Essex County, a Federalist center of strength.
<figure>
<iframe src="/figures/ma-gov/ma-gov-1800.html" style="height:500px;width:100%;">
<br />
</iframe>
<figcaption>
Figure 3: Election for Massachusetts governor, 1800. Data: [A New Nation Votes](http://elections.lib.tufts.edu/).

</figcaption>
</figure>
By 1807 the Democratic-Republican candidate, James Sullivan, did far better than Gerry had in the Western counties and in some Western towns he did better than Strong. Sullivan even made some inroads into Essex County and Cape Cod, though Strong made some inroads into Maine. This election was closely contested in nearly every town, and Sullivan narrowly defeated Strong by gaining votes in places that had gone heavily for Federalists in earlier elections. The change in politics from dominance by one party to heavily contested elections that we noted in the charts above also appears on this map.
<figure>
<iframe src="/figures/ma-gov/ma-gov-1807.html" style="height:500px;width:100%;">
<br />
</iframe>
<figcaption>
Figure 4: Election for Massachusetts governor, 1807. Data: [A New Nation Votes](http://elections.lib.tufts.edu/).

</figcaption>
</figure>
In 1823, Otis maintained some of the Federalist strength in western Massachusetts, though he also lost (I suspect that when Maine gained statehood in 1820, the Federalists benefited slightly from a decline of Republican votes). But Otis was defeated in most of the towns surrounding Boston.
<figure>
<iframe src="/figures/ma-gov/ma-gov-1823.html" style="height:500px;width:100%;">
<br />
</iframe>
<figcaption>
Figure 5: Election for Massachusetts governor, 1823. Data: [A New Nation Votes](http://elections.lib.tufts.edu/).

</figcaption>
</figure>
These maps show comparatively little of the split between "blue" cities and "red" country that we are accustomed to in modern electoral maps. This is hardly surprising, since mass urbanization happened much later. But what is surprising in these few maps is how close the vote was in many towns. The line between Federalists and Democratic-Republicans did not run between towns but through them. Elections were highly competitive at the state level, but that competition was also reflected in most towns.[<sup>7</sup>](#fn7){\#fnref7.footnoteRef} There is a lot more work to do, including figuring out a better way of representing votes by town, creating maps for all the Massachusetts gubernatorial elections, extending the analysis to other states and other types of elections, and taking on questions such as voter turnout and changing patterns of votes within particular towns.

#### Appendix

If you would like to look up a particular election or candidate, use the table below.
<figure>
<iframe src="/figures/ma-gov/gov-datatable.html" style="height:390px;width:100%;">
<br />
</iframe>
<figcaption>
Table 1: Top five contenders in each election for governor of Massachusetts, 1787--1824. Data: [A New Nation Votes](http://elections.lib.tufts.edu/). Search by year or candidate name.

</figcaption>
</figure>
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
<p>
See the summer 2013 issue of the <em>Journal of the Early Republic</em>, which includes the following articles: Caroline F. Sloat, "A New Nation Votes and the Study of American Politics, 1789-1824," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 183--86, doi:10.1353/jer.2013.0042; John L. Brooke, "'King George Has Issued Too Many Pattents for Us': Property and Democracy in Jeffersonian New York," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 187--217, doi:10.1353/jer.2013.0037; Donald Ratcliffe, "The Right to Vote and the Rise of Democracy, 1787-1828," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 219--54, doi:10.1353/jer.2013.0033; Philip J. Lampi, "The Federalist Party Resurgence, 1808-1816: Evidence from the New Nation Votes Database," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 255--81, doi:10.1353/jer.2013.0029; Rosemarie Zagarri, "The Family Factor: Congressmen, Turnover, and the Burden of Public Service in the Early American Republic," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 283--316, doi:10.1353/jer.2013.0026; Andrew W. Robertson, "Afterword: Reconceptualizing Jeffersonian Democracy," <em>Journal of the Early Republic</em> 33, no. 2 (2013): 317--34, doi:10.1353/jer.2013.0023.<a href="#fnref1">↩</a>
</p>
</li>
<li id="fn2">
<p>
Samuel Adams (won 1796) is listed as a Republican in NNV.<a href="#fnref2">↩</a>
</p>
</li>
<li id="fn3">
<p>
Governors Levi Lincoln Sr., Moses Gill, and Marcus Morton succeeded governors who died in office but did not win office in their own right. The turnover between parties must be attributed, at least in part, to weak successors running for governor.<a href="#fnref3">↩</a>
</p>
</li>
<li id="fn4">
<p>
This required an additional step to distinguish between Maine and Massachusetts, since until 1820, towns in what is today Maine were part of Massachusetts. A few populated places, such as "Number 8 and 9" in Maine could not be geocoded, but those places account for fewer than 100 votes per election.<a href="#fnref4">↩</a>
</p>
</li>
<li id="fn5">
<p>
Because many more people voted in later elections, the relationship between the size of the circles and the number of votes varies from map to map.<a href="#fnref5">↩</a>
</p>
</li>
<li id="fn6">
<p>
I am surprised that a Federalist did better in Western Massachusetts. Am I wrong to be surprised?<a href="#fnref6">↩</a>
</p>
</li>
<li id="fn7">
<p>
Of course there are some exceptions. Chesire and Adams stand out to me: both were home to mills, and Chesire had a glass factory. Did these mill towns have a different kind of politics?<a href="#fnref7">↩</a>
</p>
</li>
</section>
