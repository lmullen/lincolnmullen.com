---
title: 'Analyzing Historical History Dissertations: Cleaning the Data Again'
date: "2014-02-13"
slug: 'analyzing-historical-history-dissertations-cleaning-the-data-again'
tags:
- digital history
draft: false
---

*This series on "Analyzing Historical History Dissertations" is a work in progress. If you would like to cite or link to it, please consider using the **[landing page](http://lincolnmullen.com/research/history-dissertations/)**, which will always have the most up-to-date information and a list of all the posts.*

After posts about the length of dissertations about history university by university and for the American academy as a whole, along with a map of where history dissertations have been written, this post will be much less interesting. In this post I want to write about some decisions that I've made to further clean and tidy the data, and the consequent revisions to earlier visualizations. I'm writing about this process because I have conceived of this series of posts as being not just a presentation of my findings about history dissertations, but a set of reflections on the process of data analysis as performed by a historian. I hope this series will be pedagogically useful. Certainly no one ever taught me how to do this kind of data analysis in a history department, and it's not that hard to find scientists complaining that no one ever taught them how to do this kind of work. So I hope this series of posts and the [accompanying code](https://github.com/lmullen/dissertations-data) is one [resource](http://programminghistorian.org/) [among](http://geospatialhistorian.wordpress.com/) [many](http://software-carpentry.org/), and in particular that it will be useful as a guide to the many steps in digging through a mid- to large-sized data set.

First, what do I mean by cleaning and tidying data? By cleaning data I mean standardizing the values of the data to make them more useful. For example, I have filtered out at least twenty thousand irrelevant dissertations from the data set. Cleaning might also mean performing an operation like standardizing the name of a university, or joining all the fields that hold paragraphs for an abstract into one field. By tidying data I mean what [Hadley Wickham](http://had.co.nz/) means in this fantastic [journal article](http://vita.had.co.nz/papers/tidy-data.pdf) and [talk](http://vimeo.com/33727555), namely, changing the structure of data so that "each variable is a column, each observation is a row, and each type of observational unit is a table." To put it a different way, tidying data is [normalizing the data](http://en.wikipedia.org/wiki/Database_normalization); cleaning the data is akin to proofreading the data.

So why do I need to clean and tidy the data after five posts? One possible model of how data analysis should proceed is that you first clean the data, then analyze. Another possible model is that you first set up the basic structure of the analysis, then clean the data to get the results that you want. In practice, learning from data is a process that iterates between analysis and cleaning. The process is iterative because you have to analyze the data before you know how the data must be cleaned and tidied, and you have to clean and tidy the data to make certain kinds of analysis available.

These are the kinds of tidying and cleaning that can happen at this stage.

First, [mapping the dissertations](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-location/) made it obvious that the ProQuest data set contains very incomplete information for dissertations written outside North America. We may as well throw away those observations, since they are too incomplete to tell us anything useful. The data set also contains 4,608 dissertations from Canada. Looking at the quality of the data recorded for those Canadian dissertations, it seems that they are worth keeping. The deciding point for me is that the [AHA Directory of History Dissertations](https://secure.historians.org/pubs/dissertations/index.cfm) also includes Canadian dissertations.

Second, looking closely at the list of universities shows that there are some, such as the "Yale University, School of Forestry and Environmental Studies" which represent the occasional dissertation that might be about history, but is probably so far afield from the norm of a dissertation in history department that these may as well be excluded. Most of these will be dropped anyway when we filter the dissertations by subject.

Now we have a more difficult decision to make, since we have to decide what counts as a dissertation in history. There is a strong argument to be made that the best definition is a dissertation written in a history department. Unfortunately, since only 7.64 percent of the dissertations in this data set identify their department, it is impossible to use that definition. But there are intellectual justifications for a more capacious definition. Certainly in my own field of American religious history, many if not most of the dissertations are written by people in religious studies departments, divinity schools, Judaic Studies departments, and the like. Economic history, particularly of a quantitative bent, has mostly moved to economics departments (though history of capitalism is making a roaring comeback in history departments). So I have no problem with using the ProQuest subject headings to study a broader definition of what a history dissertation is.

The sticking point for me is whether to include dissertations with the subject "Education, History of" as dissertations in history. The argument that such dissertations are usually written in schools of education rather than colleges of arts and sciences, are therefore held to different norms and so should be excluded from this study is persuasive to me. But we can also test the hypothesis that they are held to different norms.

Here is a chart of the page lengths of history of education dissertations vs all other kinds of history dissertations.

![](//files.lincolnmullen.com/downloads/historical-dissertations/history-of-education-vs-general.png)

There is a very wide divergence between the history of education and the other history dissertations, so I've decided to reverse my earlier decision and exclude them from the data set.

There is no such divergence for history of religion dissertations, so they can stay.

![](//files.lincolnmullen.com/downloads/historical-dissertations/history-of-religion-vs-general.png)

History of science has a bizarre trajectory compared to every other kind of history. (Historians of science, any ideas what is going on here?) But even though historians of science are often in different departments, they can stay in the data set.

![](//files.lincolnmullen.com/downloads/historical-dissertations/history-of-science-vs-general.png)

History of economics diverges from general history (though not as sharply as history of education). I'm betting that's because such dissertations are more likely to use quantitative methods, so they can stay.

![](//files.lincolnmullen.com/downloads/historical-dissertations/history-of-economics-vs-general.png)

Biography also seems close to the profession in general.

![](//files.lincolnmullen.com/downloads/historical-dissertations/biography-vs-general.png)

As a last step, I'm dropping all universities with fewer than five dissertations that meet all these criteria. Looking through these universities, many of their dissertations seem like false positives, so this cut-off should clean up problems in the long tail.

Making these changes to the data set gives us fewer dissertations: 84428 dissertations before, 80575 after cleaning, which is 3853 fewer dissertations. It's a messy process, but then the data is messy. Keep that in mind the next time you're impressed by a slick visualization.

---

Now that the data is more thoroughly cleaned, I have redone some of the more significant visualizations from the earlier posts.

First, the number of dissertations per year.

![](//files.lincolnmullen.com/downloads/historical-dissertations/number-of-dissertations-per-year-2.png)

Second, the mean and median page counts of history dissertations. These are slightly higher after cleaning the data more thoroughly.

![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-mean-median-2.png)

Third, a map of the locations where history dissertations were written.

<link rel="stylesheet" href="//files.lincolnmullen.com/downloads/historical-dissertations/location/MarkerCluster.css" />

::: {#map}
:::

Finally, a redo of all the page counts by university. (Because nothing makes a post popular like giving every historian in North America a little blue dot of their own to look for.)

---

Columbia University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-columbia-university.png)

The University of Wisconsin -- Madison. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-wisconsin---madison.png)

University of California, Berkeley. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-berkeley.png)

Yale University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-yale-university.png)

University of California, Los Angeles. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-los-angeles.png)

Harvard University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-harvard-university.png)

New York University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-new-york-university.png)

University of Michigan. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-michigan.png)

University of Pennsylvania. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-pennsylvania.png)

The University of Chicago. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-chicago.png)

Princeton University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-princeton-university.png)

Indiana University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-indiana-university.png)

University of Minnesota. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-minnesota.png)

Stanford University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-stanford-university.png)

The University of Texas at Austin. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-texas-at-austin.png)

The University of North Carolina at Chapel Hill. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-north-carolina-at-chapel-hill.png)

University of Virginia. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-virginia.png)

The Ohio State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-ohio-state-university.png)

Duke University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-duke-university.png)

The Johns Hopkins University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-johns-hopkins-university.png)

University of Illinois at Urbana-Champaign. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-illinois-at-urbana-champaign.png)

Cornell University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-cornell-university.png)

City University of New York. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-city-university-of-new-york.png)

Northwestern University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-northwestern-university.png)

University of Washington. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-washington.png)

Rutgers The State University of New Jersey -- New Brunswick. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-rutgers-the-state-university-of-new-jersey---new-brunswick.png)

University of Maryland, College Park. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-maryland,-college-park.png)

University of California, Santa Barbara. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-santa-barbara.png)

The University of Iowa. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-iowa.png)

Emory University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-emory-university.png)

The Florida State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-florida-state-university.png)

Brown University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-brown-university.png)

University of Kansas. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-kansas.png)

University of Pittsburgh. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-pittsburgh.png)

Temple University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-temple-university.png)

Michigan State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-michigan-state-university.png)

University of Notre Dame. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-notre-dame.png)

Boston University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-boston-university.png)

Georgetown University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-georgetown-university.png)

University of Southern California. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-southern-california.png)

Brandeis University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-brandeis-university.png)

University of Missouri -- Columbia. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-missouri---columbia.png)

State University of New York at Binghamton. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-state-university-of-new-york-at-binghamton.png)

The Pennsylvania State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-pennsylvania-state-university.png)

State University of New York at Buffalo. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-state-university-of-new-york-at-buffalo.png)

Vanderbilt University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-vanderbilt-university.png)

University of California, San Diego. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-san-diego.png)

The University of Arizona. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-arizona.png)

The Catholic University of America. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-catholic-university-of-america.png)

University of Colorado at Boulder. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-colorado-at-boulder.png)

Syracuse University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-syracuse-university.png)

University of Florida. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-florida.png)

Saint Louis University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-saint-louis-university.png)

The University of New Mexico. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-new-mexico.png)

The Claremont Graduate University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-claremont-graduate-university.png)

The George Washington University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-george-washington-university.png)

Tulane University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-tulane-university.png)

University of Massachusetts Amherst. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-massachusetts-amherst.png)

University of California, Irvine. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-irvine.png)

State University of New York at Stony Brook. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-state-university-of-new-york-at-stony-brook.png)

University of Rochester. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-rochester.png)

University of Kentucky. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-kentucky.png)

University of South Carolina. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-south-carolina.png)

University of Delaware. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-delaware.png)

University of Hawai'i at Manoa. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-hawai'i-at-manoa.png)

The University of Oklahoma. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-oklahoma.png)

University of California, Davis. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-davis.png)

Arizona State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-arizona-state-university.png)

University of Connecticut. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-connecticut.png)

University of Oregon. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-oregon.png)

Louisiana State University and Agricultural & Mechanical College. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-louisiana-state-university-and-agricultural-&-mechanical-college.png)

University of Georgia. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-georgia.png)

American University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-american-university.png)

Washington University in St. Louis. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-washington-university-in-st.-louis.png)

Boston College. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-boston-college.png)

The University of Nebraska -- Lincoln. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-nebraska---lincoln.png)

University of California, Riverside. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-riverside.png)

Case Western Reserve University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-case-western-reserve-university.png)

Drew University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-drew-university.png)

Fordham University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-fordham-university.png)

Texas A&M University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-texas-a&m-university.png)

The University of Utah. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-utah.png)

Rice University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-rice-university.png)

Washington State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-washington-state-university.png)

Kent State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-kent-state-university.png)

University of Illinois at Chicago. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-illinois-at-chicago.png)

Purdue University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-purdue-university.png)

University of Cincinnati. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-cincinnati.png)

Loyola University Chicago. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-loyola-university-chicago.png)

Howard University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-howard-university.png)

Texas Christian University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-texas-christian-university.png)

Ohio University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-ohio-university.png)

Marquette University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-marquette-university.png)

The University of Tennessee. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-tennessee.png)

Boston University Graduate School. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-boston-university-graduate-school.png)

The University of Alabama. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-alabama.png)

Graduate Theological Union. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-graduate-theological-union.png)

The College of William and Mary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-college-of-william-and-mary.png)

Bryn Mawr College. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-bryn-mawr-college.png)

University of California, Santa Cruz. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-santa-cruz.png)

University of North Texas. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-north-texas.png)

The Southern Baptist Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-southern-baptist-theological-seminary.png)

West Virginia University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-west-virginia-university.png)

University of Houston. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-houston.png)

Bowling Green State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-bowling-green-state-university.png)

University of Denver. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-denver.png)

Texas Tech University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-texas-tech-university.png)

Auburn University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-auburn-university.png)

State University of New York at Albany. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-state-university-of-new-york-at-albany.png)

Northern Illinois University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-northern-illinois-university.png)

The Union Institute. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-union-institute.png)

Southern Illinois University at Carbondale. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-southern-illinois-university-at-carbondale.png)

Southwestern Baptist Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-southwestern-baptist-theological-seminary.png)

Wayne State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-wayne-state-university.png)

The University of Mississippi. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-mississippi.png)

Miami University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-miami-university.png)

New School for Social Research. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-new-school-for-social-research.png)

Princeton Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-princeton-theological-seminary.png)

Oklahoma State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-oklahoma-state-university.png)

The University of Maine. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-maine.png)

University of Arkansas. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-arkansas.png)

University of Miami. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-miami.png)

George Mason University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-george-mason-university.png)

Georgia State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-georgia-state-university.png)

Carnegie Mellon University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-carnegie-mellon-university.png)

Fuller Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-fuller-theological-seminary.png)

Mississippi State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-mississippi-state-university.png)

University of New Hampshire. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-new-hampshire.png)

Baylor University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-baylor-university.png)

Kansas State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-kansas-state-university.png)

Clark University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-clark-university.png)

Iowa State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-iowa-state-university.png)

The University of Toledo. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-toledo.png)

The University of Southern Mississippi. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-southern-mississippi.png)

Ball State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-ball-state-university.png)

Union Institute and University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-union-institute-and-university.png)

New School University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-new-school-university.png)

Lehigh University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-lehigh-university.png)

The University of Wisconsin -- Milwaukee. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-wisconsin---milwaukee.png)

The University of Texas at Arlington. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-texas-at-arlington.png)

Tufts University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-tufts-university.png)

Brigham Young University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-brigham-young-university.png)

The University of Memphis. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-memphis.png)

Northern Arizona University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-northern-arizona-university.png)

The University of Texas at Dallas. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-texas-at-dallas.png)

Union Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-union-theological-seminary.png)

Southern Methodist University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-southern-methodist-university.png)

Virginia Polytechnic Institute and State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-virginia-polytechnic-institute-and-state-university.png)

Westminster Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-westminster-theological-seminary.png)

Duquesne University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-duquesne-university.png)

Western Michigan University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-western-michigan-university.png)

University of Idaho. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-idaho.png)

California Institute of Integral Studies. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-california-institute-of-integral-studies.png)

Regent University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-regent-university.png)

The University of Akron. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-akron.png)

The Jewish Theological Seminary of America. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-jewish-theological-seminary-of-america.png)

University of Missouri -- Kansas City. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-missouri---kansas-city.png)

Florida International University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-florida-international-university.png)

Pacifica Graduate Institute. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-pacifica-graduate-institute.png)

University of Nevada, Reno. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-nevada,-reno.png)

Virginia Commonwealth University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-virginia-commonwealth-university.png)

The University of North Carolina at Greensboro. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-university-of-north-carolina-at-greensboro.png)

Trinity Evangelical Divinity School. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-trinity-evangelical-divinity-school.png)

New Orleans Baptist Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-new-orleans-baptist-theological-seminary.png)

University of Louisville. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-louisville.png)

Hebrew Union College -- Jewish Institute of Religion (Ohio). ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-hebrew-union-college---jewish-institute-of-religion-(ohio).png)

Georgia Institute of Technology. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-georgia-institute-of-technology.png)

Andrews University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-andrews-university.png)

Northeastern University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-northeastern-university.png)

Southeastern Baptist Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-southeastern-baptist-theological-seminary.png)

Yeshiva University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-yeshiva-university.png)

The Iliff School of Theology and University of Denver. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-the-iliff-school-of-theology-and-university-of-denver.png)

University of Nevada, Las Vegas. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-nevada,-las-vegas.png)

Oregon State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-oregon-state-university.png)

Rensselaer Polytechnic Institute. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-rensselaer-polytechnic-institute.png)

Mid-America Baptist Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-mid-america-baptist-theological-seminary.png)

Memphis State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-memphis-state-university.png)

University of California, San Francisco. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-california,-san-francisco.png)

Texas Woman's University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-texas-woman's-university.png)

Colorado State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-colorado-state-university.png)

Salve Regina University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-salve-regina-university.png)

Indiana University of Pennsylvania. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-indiana-university-of-pennsylvania.png)

University of South Florida. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-south-florida.png)

Walden University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-walden-university.png)

Lutheran School of Theology at Chicago. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-lutheran-school-of-theology-at-chicago.png)

Middle Tennessee State University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-middle-tennessee-state-university.png)

University of Wyoming. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-wyoming.png)

Capella University. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-capella-university.png)

Dallas Theological Seminary. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-dallas-theological-seminary.png)

University of New Orleans. ![](//files.lincolnmullen.com/downloads/historical-dissertations/pages-2-university-of-new-orleans.png)
