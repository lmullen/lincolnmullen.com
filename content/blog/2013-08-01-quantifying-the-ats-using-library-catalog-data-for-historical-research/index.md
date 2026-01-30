---
title: 'Quantifying the ATS: Using Library Catalog Data for Historical Research'
date: "2013-08-01"
slug: 'quantifying-the-ats-using-library-catalog-data-for-historical-research'
tags:
- digital history
- American religious history
draft: false
---

*[This post originally appeared at [Religion in American History](http://usreligion.blogspot.com/2013/08/quantifying-american-tract-society.html).]*

The [American Antiquarian Society](http://www.americanantiquarian.org/) was generous enough to offer me a fellowship this summer, so I took a month to research in the [AAS's wonderful collections](http://usreligion.blogspot.com/2013/06/know-your-archives-american-antiquarian.html). A fair bit of my time was spent reading through the nearly complete print run of American Tract Society pamphlets from the early to mid-nineteenth century. I wanted to answer a question: did the ATS pamphlets include "the sinner's prayer," a common feature of twentieth century evangelical tracts? If it did, when did the genre develop? I think I have an answer to that question: yes, the ATS did contribute to the development of the genre of the sinner's prayer. More on what I found another time, both here on the blog and in chapter two of the dissertation.

Answering that question required a sense of how the ATS tracts were distributed over time, but reading each pamphlets was a poor way to get that sense. Down in the weeds of the pamphlets---most of which sounded alike---it was hard to see how the pamphlets were distributed over time. And whatever patterns I could discern were likely to be unreliable: I was, after all, looking for a particular pattern and was thus at risk of committing a classic historians' fallacy. What I needed was an external way of looking at the texts to supplement my close reading.

There was a second problem: it's really hard to date the ATS tracts. While most of them print a place of publication and a series number, none of them include a date. The date was left off because the tracts were frequently reprinted, sometimes in millions of impressions. For example a popular tract, Legh Richmond's *The African Servant*, was printed in multiple editions ranging at least from 1816 to 1846, though all bear the series number 60. But S. J. Wolfe of the Antiquarian Society figured out [how to date](http://americanantiquarian.org/ATS.htm) the tracts based on the changing addresses of the ATS. Where I brought a knowledge of theology and religious history to my reading of the ATS tract, and even a general knowledge of the history, economics, and methods of the American Tract Society drawn from other historians' work, I had no knowledge of the printing history of the ATS pamphlets. The catalogers at the Antiquarian Society, on the other hand, knew a lot of hard, cold facts about when the pamphlets were published, and who the anonymous authors were. Just as important, they knew when the facts were actually soft and fuzzy, and they had taken pains to represent the ambiguities of the data in the catalog.

The solution to the problem of dating the tracts was also the solution to the problem of getting an overview of the tracts. Because the Antiquarian Society's catalog lets users export bibliographic information in a variety of formats, I set out to visualize the ATS printing history using the catalog data. Other scholars have are also using library data to visualize printing: Jim Egan and Jean Bauer have a much more [ambitious project](http://www.stg.brown.edu/projects/mapping-genres/) using collections from the American Antiquarian Society, the John Carter Brown Library, and the John Hay Library.

### Method

I made a series of visualizations about various questions using Ruby and R. I'll gloss over the code here, but the code and the cleaned-up data is [available on GitHub](https://github.com/lmullen/ATS-bibliography). Here's how I did it:

1.  I exported the records for all the relevant items from the AAS catalog. In this case I searched for the records by call number, but for other questions one could search by subject heading, by publisher, and so on. I exported to MARC, since it offered the fullest data, and since the Ruby library for MARC was better than for other formats.

2.  I used a [Ruby script](https://github.com/lmullen/ATS-bibliography/blob/master/munge.rb) to munge the MARC records into usable [CSV files](https://github.com/lmullen/ATS-bibliography/tree/master/data-clean).

3.  I used an [R script](https://github.com/lmullen/ATS-bibliography/blob/master/analysis.R) to analyze the data.

The kind of analysis that I'm writing about here is not an argument itself, but a step along the way to an argument. Nor do I mean to imply that there is some kind of "scientific" value in this analysis. Take this for what it is: a behind-the-scenes musing about methodology.

### Dates

One of my questions was when the ATS published its tracts. This is a histogram of the tracts by date.

{{< figure src="ats-pub-date.png" >}}

The American Tract Society was founded in 1825 from the merger of the New-York Religious Tract Society and the Boston-based American Tract Society (formerly, the New England Religious Tract Society). The chart shows that the ATS quickly ramped up the publication of tracts in its first five years. Many of these early pamphlets were reprints from British tract societies, such as the Religious Tract Society of London founded in 1799. As was the case with most evangelical institutions from this period, the British had a head start, and evangelicalism was an Anglophone movement with debts to German, French, and Swiss Protestants. The ATS republished works from proven British authors, such as Legh Richmond and Hannah More. After the early series of tracts, beginning in the 1830s, the ATS published more of its own original tracts.

The above chart shows the earliest possible date that a given tract was published. But the ATS republished these tracts in many editions throughout the nineteenth century. The point is that the ATS publications tended to be first published in the decade between 1825 and 1835---exactly the period when Charles Finney was most active in New York. But unlike Finney, these tracts kept getting the message out for the entirety of the nineteenth century. (The ATS is [still publishing](http://www.crossway.org/group/ats) today.) And though much of Finney's powerful presentation came from being in his presence, and sitting on his anxious bench, the ATS tracts were carried across the country by colporteurs to places like the pine barrens of New Jersey which the government could barely reach.

### Subjects

A second, obvious, question is, what were these tracts about? Though I had a pretty good sense of the proportions of the various topics, counting the subject headings is a useful correction to possibly haphazard impressions. The ATS catalog entries have 186 distinct subject headings. The distribution of all the subject headings is the familiar "long tail." Keep in mind, though, that each tract can have more than one subject heading, so [a tract](http://catalog.mwa.org/vwebv/holdingsInfo?bibId=229772) might be about both one of the most obscure topics ("Wine---poetry") and one of the most common topics ("Temperance").

{{< figure src="ats-all-subjects.png" >}}

The chart is more useful if we limit it to only the twenty most frequently used subject headings.

{{< figure src="ats-subjects.png" >}}

These headings fall into several groups:

-   Tracts which promote conversion: headings like salvation, conversion, repentance, future punishment, and converts certainly fit into this group.

-   Tracts about family life and the education of children. These tracts are also about how to prepare children for conversion.

-   The rest of subject headings, which considered together makes the biggest grouping, are instructions in doctrine and the means of grace. The ATS seems to have thought of its mission as being to edify Christians as much as to evangelize unbelievers.

There is one topic which I found to be ubiquitous in reading the tracts, but which is barely noticeable in the subject headings, namely, warnings against infidelity, atheism, and liberal religion. Tracts with the subject heading "future punishment" were mostly written against Universalists. Other tracts were written against Unitarians. But even though only a few tracts are explicitly about Deists or atheists, they are mentioned or appear as characters in, I would guess, two out of three tracts. Curiously, Catholics are barely mentioned except in a very few "patriotic" tracts. Historians have long recognized the nineteenth century as the century of Protestant anti-Catholicism, and of course it was. But for the ATS the real bogeyman was not Catholicism but infidelity. On the importance of atheism, skepticism, and free thought for shaping American religion, see my recent [historiographic essay](http://usreligion.blogspot.com/2013/06/where-are-histories-of-american.html) and the soon-to-be last chapter of my dissertation.

### Length

Anyone who is familiar with the contemporary evangelical tracts would be surprised at long the ATS tracts are. Where a contemporary tract might be six or eight very small pages with several illustrations, the ATS tracts average sixteen pages of closely printed type and were commonly thirty-two pages long. Besides the main tract that comprised the body, the tracts usually had a paper wrapper which had supplementary materials or another tract in its own right. Bill Bright's well-known tract from the 1960s, *Do You Know the Four Spiritual Laws*, could fit on the wrapper of one of these hefty ATS tracts.

Since I am making an argument about how the normative experience of conversion became simplified, I wondered whether the tracts got shorter over time. They did not.

{{< figure src="ats-length.png" >}}

### Authors

The most useful analysis was figuring out who wrote these tracts. Not only were the authors frequently British or occasionally Swiss or German, the authors were frequently from the eighteenth or even the seventeenth centuries. In this histogram I've charted the number of tracts by the death date of their authors.[^1]

{{< figure src="ats-authors.png" >}}

We can get a sense of who wrote the ATS tracts by looking at most frequently published authors.[^2]

| rank | name                 | died | tracts |
|------|----------------------|------|--------|
| 1    | Doddridge, Philip    | 1751 | 11     |
| 2    | Fuller, Andrew       | 1815 | 11     |
| 3    | Richmond, Legh       | 1827 | 8      |
| 4    | McIlvaine, Charles   | 1873 | 7      |
| 5    | Watts, Isaac         | 1748 | 7      |
| 6    | Edwards, Justin      | 1853 | 6      |
| 7    | More, Hannah         | 1833 | 6      |
| 8    | Newman, William      | 1835 | 6      |
| 9    | Alexander, Archibald | 1851 | 5      |
| 10   | Baxter, Richard      | 1691 | 5      |
| 11   | Bennett, James       | 1862 | 5      |
| 12   | Burder, George       | 1832 | 5      |
| 13   | Clark, Daniel A.     | 1840 | 5      |
| 14   | Malan, Cesar         | 1864 | 5      |

The chart and the table show that most of the authors lived in the nineteenth century contemporaneous with the ATS. This chart does not include anonymous tracts, most of which were from the nineteenth century. Yet the ATS also reprinted works from Puritan divines or English Nonconformists such as John Flavel, John Owen, and Richard Baxter, and also revivalists and evangelicals like theologian Jonathan Edwards, hymn writer Isaac Watts, and most frequently the Nonconformist pastor Phillipp Doddridge.

In excerpting and re-framing these works of earlier generations, the ATS was remaking an old tradition into a new tradition. The history of the ATS publications is thus a story of long continuity in an evangelical tradition of heart religion and vital piety, centered on conversion, being gradually remade---sometimes intentionally but often through subtle shifts in the meanings of words and drifting in the substance of doctrine or practice---into a new forms of conversion.

Here's to my unsung collaborators: the diligent catalogers at the American Antiquarian Society, some of whom I had the privilege of meeting.

[^1]: I used the death date since there is necessarily quite a few years between one's birth and the time one is publishing.

[^2]: The prominence of Isaac Watts is a bit misleading, because he is listed as an author on tracts that contain his hymns, though he might not always be the primary author. But on the other hand, the frequency with which his hymns appear, probably understated in the catalog, is itself a sign of influence.
