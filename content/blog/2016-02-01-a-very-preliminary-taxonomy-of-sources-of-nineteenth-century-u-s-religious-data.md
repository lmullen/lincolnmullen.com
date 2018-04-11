---
categories:
- digital history
- religious history
date: '2016-02-01T05:00:00+00:00'
slug: 'a-very-preliminary-taxonomy-of-sources-of-nineteenth-century-u-s-religious-data'
tags:
- religion data
title: 'A Very Preliminary Taxonomy of Sources of Nineteenth-Century U.S. Religious Data'
---

In my [last post](http://usreligion.blogspot.com/2016/01/Where-the-Problem-with-Historical-Data-about-U-S-Religion-Really-Lies.html) I explained that historians of U.S. religion have barely begun to scratch the surface of the data (meaning, sources that are amenable to computation) that are available to them. To demonstrate this I gave the example of a single source, the *Minutes of the Annual Conferences of the Methodist Episcopal Church*.

In this post I want to attempt a very preliminary taxonomy of the kinds of sources that are available to religious historians who wish to use mapping or quantitative analysis of some kind or another. Let's call this a taxonomy instead of a catalog, because I'm going to list the kinds of sources that I've come about rather than try to give a bibliography of all of the sources themselves. I'd love to be able to list all the sources, but I haven't done all that work yet. And let's say this is very preliminary, because I hope this post is an example of the so-called [Cunningham's Law](https://meta.wikimedia.org/wiki/Cunningham%27s_Law): "the best way to get the right answer on the Internet is not to ask a question; it's to post the wrong answer." That is to say, if you know of a source or category of source that I don't know about, I hope you'll correct me in the comments. Finally, I should mention that I'm teaching [a course](http://lincolnmullen.com/courses/data-dh.2016/) this semester on "Data and Visualization in Digital History" where we are working on nineteenth-century U.S. religious statistics. I'm indebted to the excellent students in that course, who have already turned up many sources that I didn't know about.

Enough throat clearing.

All U.S. religious statistics are divided into two parts, those from the Census, and those not from the Census.

<!--more-->
The Census Bureau began asking [questions about religion](https://ask.census.gov/faq.php?id=5000&faqId=29) in 1850, and stopped in 1936. From 1850 to 1900, the decennial censuses asked a few questions about religion as part of their "Social Statistics" section. Some years are more interesting than others because of the questions asked. Some years asked only questions about congregations, such as the number of churches and their seating capacity, while other years asked questions about membership.

More valuable are the Censuses of Religious Bodies from [1906](http://catalog.hathitrust.org/Record/001408064), [1916](http://catalog.hathitrust.org/Record/001408065), [1926](http://catalog.hathitrust.org/Record/001408066) , and [1936](http://catalog.hathitrust.org/Record/001408067). Apparently the Census Bureau also performed a census in 1946 but didn't have the funds for tabulating the results. These censuses were gathered by sending questionnaires to congregations and following up with visits from census takers.[<sup>1</sup>](#fn1){\#fnref1.footnoteRef}

The key to the Census data, though, is the level of aggregation for which the data has been reported. This may not be immediately intuitive, so an elaboration. Let's say the Census asked questions of every congregation it could find in 1926 (which it did). The form that the Census [used in 1926](http://www2.census.gov/prod2/decennial/documents/13949806v1ch1.pdf) looked like this.
<figure id="figure-1">
[!['Form](/figures/taxonomy-religion/1926-census-form.png)](/figures/taxonomy-religion/1926-census-form.png)
<figcaption>
Figure 1: Form from the 1926 Census of Religious Bodies. \[[PNG](/figures/taxonomy-religion/1926-census-form.png)\]
</figcaption>
</figure>
The Census did not simply transcribe those forms and report them, of course. Instead it aggregated them and reported the results by, for example, denomination and city, or denomination and county, or denomination and state. The level of aggregation matters because it affects the level of detail with which we can see religious groups. If this isn't obvious now, take it on faith until a later post where I can demonstrate how an exploratory data analysis works.

Most of the religion data that the Census gathered can be accessed through [NHGIS](https://nhgis.org/) (which I wrote [somewhat naively about here](http://usreligion.blogspot.com/2013/10/historical-religion-data-in-nhgis-and.html)). Those data are aggregated to the level of the county. However, the volumes that the Census published for the Censuses of Religious Bodies report statistics for cities. This seems to me to be a far more useful level of aggregation. And I hope that it may be possible to access the raw returns to find congregation level data, but I have not yet done the archival work to verify that. If there is one thing that can be said about the Census, it is that it really went into the highways and byways to find even the smallest denominations, and is very valuable for that reason.

That leaves all the kinds of data which did not come from the Census. One broad category is data which denominations gathered for themselves.

Occasionally denominations compiled volumes which covered long spans of time (the century mark was particularly popular). I [wrote earlier](http://usreligion.blogspot.com/2014/08/visualizing-presbyterian-statistics.html) about H. C. Weber's *[Presbyterian Statistics through One Hundred Years, 1826--1926: Tabulated, Visualized, and Interpreted](http://catalog.hathitrust.org/Record/007109885)* (1927); another example is C. C. Goss's *[Statistical History of the First Century of American Methodism](http://catalog.hathitrust.org/Record/008626507)* (1866). (Alyssa Toby Fahringer found the latter source.) These volumes are useful for providing a diachronic perspective on a denomination, since they did the laborious work of gathering the year-by-year statistics into a single volume. Again, though it is a question of aggregation. In this case the retrospective volumes gather, for example, the membership totals for individual congregations in a given year and sum it up into a single number for the year.

Much more useful, if more voluminous, are the year-by-year records that denominations kept. These volumes are a level down the stack from the retrospective volumes; indeed, they were the sources that the retrospective volumes used. Many denominations kept yearly statistics. Sometimes these were reported in yearbooks, other times in the journals of general conventions, or in denominational newspapers. In general they tend to be organized by congregation, circuit, diocese, or some other administrative unit within the denomination. Sometimes exasperated denominational secretaries tried to elicit reports from congregations. Other denominations regularized the procedure: for example, Episcopalians had set Sundays throughout the liturgical year for taking attendance. Sometimes these volumes could be quite detailed in terms of the kinds of records they kept, including membership, sacramental, and financial records. They often have very detailed information on ministers. Of course, it requires a lot of work to figure out what is available for individual denominations. But here are a few examples of this kind of record:

-   The Congregational Library has found [links to Congregationalist yearbooks](http://www.congregationallibrary.org/periodicals/yearbooks)
-   I've gathered [links](http://notebook.lincolnmullen.com/Methodist%20Minutes%20of%20the%20Annual%20Conferences) to the Methodists' Minutes of the Annual Conferences.
-   And I've gathered [links](http://notebook.lincolnmullen.com/Protestant%20Episcopal%20Church%20yearbooks) to the yearbooks of the Protestant Episcopal Church.
-   The Seventh Day Adventists have gathered [their meeting minutes and other records](http://documents.adventistarchives.org/default.aspx). (Jeri Wieringa, who is doing some sophisticated work with these records for her dissertation on the Adventists, found these records.)
-   Jordan Bratt has found the General Conference Reports from the Church of Jesus Christ of Latter-day Saints, for example, [1911's report](Conference%20reports%20of%20The%20Church%20of%20Jesus%20Christ%20of%20Latter-day%20Saints).

Sometimes there is an even more detailed level of records in the form of regional reports. For example, the Protestant Episcopal Church published yearbooks and general convention reports which had data at the diocesan level. But individual dioceses also published [journals of their conventions](http://catalog.hathitrust.org/Search/Home?lookfor=annual%20convention%20diocese&searchtype=all&ft=ft&setft=true) which have data at the level of the congregation (an [example](http://babel.hathitrust.org/cgi/pt?id=nyp.33433070790765;view=1up;seq=106)).

Of course not every group fit well into the denominational label, nor did every denomination have much authority over its member congregations. For example the Union of American Hebrew Congregations had to solicit individual congregations to create this volume on *[The Statistics of the Jews of the United States](http://babel.hathitrust.org/cgi/pt?id=hvd.hnrqec;view=1up;seq=5)* (1880). (That source was found by Candace Gray.)

Besides denominational statistics, benevolent organizations kept track of the numbers of tracts printed, Bibles distributed, dollars raised, Sunday school scholars taught, dram shops shut down, and souls won. I'm inclined to treat these statistics with even more skepticism than the denominational statistics. Nevertheless, there is a lot of work to be done in understanding the reach of these benevolent organizations. For instance, I'd like to do some work following the money trail of some of these organizations. (George Oberle is doing some similar, valuable work on learned societies in the early republic.)

Then besides the denominational statistics, there are sources which were not created by religious actors themselves but which contain useful information. There are innumerable sources which do not appear on their face to be amenable to computation, but which can be used for mapping or other kinds of analysis with a little ingenuity. For instance, scholars like Jeanne Halgren Kilde and Chris Cantwell have found city directories and insurance maps useful for locating congregations. Digitized newspapers like the *[Chronicling America](http://chroniclingamerica.loc.gov/)* collection, or Michigan State's *[Sunday School Books in Nineteenth Century America Dataset](https://www.lib.msu.edu/ssbdata/)* are accessible for text mining.

This concludes my rambling and decidedly idiosyncratic tour of sources which I've encountered which might be amenable to mapping or other computational kinds of analysis. So I put the question to you, gentle reader: what sources in this vein have you encountered? Where are the tables of data yearning for someone to come and interpret them?

There is much more work to be done before we can go from sources to historical interpretations. In future posts I'll write about how this data can be transcribed and located on a map, how an exploratory data analysis works, and make at least one map for you.

In the meantime, as a bonus for making it to the end of the post, here is a [table](http://babel.hathitrust.org/cgi/pt?id=wu.89072975139;view=1up;seq=110) of people who pledged money to the Protestant Episcopal Church's Diocese of Central Pennsylvania, failed to pay, and were mentioned by name and outstanding amount in the 1880 journal of the diocese.
<figure id="figure-2">
[!['Pay](/figures/taxonomy-religion/debtors.png)](/figures/taxonomy-religion/debtors.png)
<figcaption>
Figure 2: Pay me that thou owest. \[[PNG](/figures/taxonomy-religion/debtors.png)\]
</figcaption>
</figure>
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
See Rodney Stark, "The Reliability of Historical United States Census Data on Religion," <em>Sociology of Religion</em> 53, no. 1 (1992): 91--95, doi:10.2307/3711631.<a href="#fnref1">↩</a>
</li>
</section>
