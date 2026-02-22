---
title: 'Analyzing historical history dissertations: cleaning the data again'
date: "2014-02-13"
slug: 'analyzing-historical-history-dissertations-cleaning-the-data-again'
description: "Data analysis is iterative: cleaning reveals what must be re-analyzed, and analysis reveals what must be cleaned—lessons from filtering and tidying a ProQuest dissertation dataset."
tags:
- digital history
draft: false
---

After posts about the length of dissertations about history university by university and for the American academy as a whole, along with a map of where history dissertations have been written, this post will be much less interesting. In this post I want to write about some decisions that I've made to further clean and tidy the data, and the consequent revisions to earlier visualizations. I'm writing about this process because I have conceived of this series of posts as being not just a presentation of my findings about history dissertations, but a set of reflections on the process of data analysis as performed by a historian. I hope this series will be pedagogically useful. Certainly no one ever taught me how to do this kind of data analysis in a history department, and it's not that hard to find scientists complaining that no one ever taught them how to do this kind of work. So I hope this series of posts and the [accompanying code](https://github.com/lmullen/dissertations-data) is one [resource](http://programminghistorian.org/) [among](http://geospatialhistorian.wordpress.com/) [many](http://software-carpentry.org/), and in particular that it will be useful as a guide to the many steps in digging through a mid- to large-sized data set.

First, what do I mean by cleaning and tidying data? By cleaning data I mean standardizing the values of the data to make them more useful. For example, I have filtered out at least twenty thousand irrelevant dissertations from the data set. Cleaning might also mean performing an operation like standardizing the name of a university, or joining all the fields that hold paragraphs for an abstract into one field. By tidying data I mean what [Hadley Wickham](http://had.co.nz/) means in this fantastic [journal article](http://vita.had.co.nz/papers/tidy-data.pdf) and [talk](http://vimeo.com/33727555), namely, changing the structure of data so that "each variable is a column, each observation is a row, and each type of observational unit is a table." To put it a different way, tidying data is [normalizing the data](http://en.wikipedia.org/wiki/Database_normalization); cleaning the data is akin to proofreading the data.

So why do I need to clean and tidy the data after five posts? One possible model of how data analysis should proceed is that you first clean the data, then analyze. Another possible model is that you first set up the basic structure of the analysis, then clean the data to get the results that you want. In practice, learning from data is a process that iterates between analysis and cleaning. The process is iterative because you have to analyze the data before you know how the data must be cleaned and tidied, and you have to clean and tidy the data to make certain kinds of analysis available.

These are the kinds of tidying and cleaning that can happen at this stage.

First, [mapping the dissertations](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-location/) made it obvious that the ProQuest data set contains very incomplete information for dissertations written outside North America. We may as well throw away those observations, since they are too incomplete to tell us anything useful. The data set also contains 4,608 dissertations from Canada. Looking at the quality of the data recorded for those Canadian dissertations, it seems that they are worth keeping. The deciding point for me is that the [AHA Directory of History Dissertations](https://secure.historians.org/pubs/dissertations/index.cfm) also includes Canadian dissertations.

Second, looking closely at the list of universities shows that there are some, such as the "Yale University, School of Forestry and Environmental Studies" which represent the occasional dissertation that might be about history, but is probably so far afield from the norm of a dissertation in history department that these may as well be excluded. Most of these will be dropped anyway when we filter the dissertations by subject.

Now we have a more difficult decision to make, since we have to decide what counts as a dissertation in history. There is a strong argument to be made that the best definition is a dissertation written in a history department. Unfortunately, since only 7.64 percent of the dissertations in this data set identify their department, it is impossible to use that definition. But there are intellectual justifications for a more capacious definition. Certainly in my own field of American religious history, many if not most of the dissertations are written by people in religious studies departments, divinity schools, Judaic Studies departments, and the like. Economic history, particularly of a quantitative bent, has mostly moved to economics departments (though history of capitalism is making a roaring comeback in history departments). So I have no problem with using the ProQuest subject headings to study a broader definition of what a history dissertation is.

The sticking point for me is whether to include dissertations with the subject "Education, History of" as dissertations in history.

<!--more-->

The argument that such dissertations are usually written in schools of education rather than colleges of arts and sciences, are therefore held to different norms and so should be excluded from this study is persuasive to me. But we can also test the hypothesis that they are held to different norms.

Here is a chart of the page lengths of history of education dissertations vs all other kinds of history dissertations.

![](history-of-education-vs-general.png)

There is a very wide divergence between the history of education and the other history dissertations, so I've decided to reverse my earlier decision and exclude them from the data set.

There is no such divergence for history of religion dissertations, so they can stay.

![](history-of-religion-vs-general.png)

History of science has a bizarre trajectory compared to every other kind of history. (Historians of science, any ideas what is going on here?) But even though historians of science are often in different departments, they can stay in the data set.

![](history-of-science-vs-general.png)

History of economics diverges from general history (though not as sharply as history of education). I'm betting that's because such dissertations are more likely to use quantitative methods, so they can stay.

![](history-of-economics-vs-general.png)

Biography also seems close to the profession in general.

![](biography-vs-general.png)

As a last step, I'm dropping all universities with fewer than five dissertations that meet all these criteria. Looking through these universities, many of their dissertations seem like false positives, so this cut-off should clean up problems in the long tail.

Making these changes to the data set gives us fewer dissertations: 84428 dissertations before, 80575 after cleaning, which is 3853 fewer dissertations. It's a messy process, but then the data is messy. Keep that in mind the next time you're impressed by a slick visualization.

---

Now that the data is more thoroughly cleaned, I have redone some of the more significant visualizations from the earlier posts.

First, the number of dissertations per year.

![](number-of-dissertations-per-year-2.png)

Second, the mean and median page counts of history dissertations. These are slightly higher after cleaning the data more thoroughly.

![](pages-mean-median-2.png)

Finally, a redo of all the page counts by university. (Because nothing makes a post popular like giving every historian in North America a little blue dot of their own to look for.)

---

{{< figure src="pages-2-columbia-university.png" caption="Columbia University" >}}

{{< figure src="pages-2-the-university-of-wisconsin---madison.png" caption="The University of Wisconsin -- Madison" >}}

{{< figure src="pages-2-university-of-california,-berkeley.png" caption="University of California, Berkeley" >}}

{{< figure src="pages-2-yale-university.png" caption="Yale University" >}}

{{< figure src="pages-2-university-of-california,-los-angeles.png" caption="University of California, Los Angeles" >}}

{{< figure src="pages-2-harvard-university.png" caption="Harvard University" >}}

{{< figure src="pages-2-new-york-university.png" caption="New York University" >}}

{{< figure src="pages-2-university-of-michigan.png" caption="University of Michigan" >}}

{{< figure src="pages-2-university-of-pennsylvania.png" caption="University of Pennsylvania" >}}

{{< figure src="pages-2-the-university-of-chicago.png" caption="The University of Chicago" >}}

{{< figure src="pages-2-princeton-university.png" caption="Princeton University" >}}

{{< figure src="pages-2-indiana-university.png" caption="Indiana University" >}}

{{< figure src="pages-2-university-of-minnesota.png" caption="University of Minnesota" >}}

{{< figure src="pages-2-stanford-university.png" caption="Stanford University" >}}

{{< figure src="pages-2-the-university-of-texas-at-austin.png" caption="The University of Texas at Austin" >}}

{{< figure src="pages-2-the-university-of-north-carolina-at-chapel-hill.png" caption="The University of North Carolina at Chapel Hill" >}}

{{< figure src="pages-2-university-of-virginia.png" caption="University of Virginia" >}}

{{< figure src="pages-2-the-ohio-state-university.png" caption="The Ohio State University" >}}

{{< figure src="pages-2-duke-university.png" caption="Duke University" >}}

{{< figure src="pages-2-the-johns-hopkins-university.png" caption="The Johns Hopkins University" >}}

{{< figure src="pages-2-university-of-illinois-at-urbana-champaign.png" caption="University of Illinois at Urbana-Champaign" >}}

{{< figure src="pages-2-cornell-university.png" caption="Cornell University" >}}

{{< figure src="pages-2-city-university-of-new-york.png" caption="City University of New York" >}}

{{< figure src="pages-2-northwestern-university.png" caption="Northwestern University" >}}

{{< figure src="pages-2-university-of-washington.png" caption="University of Washington" >}}

{{< figure src="pages-2-rutgers-the-state-university-of-new-jersey---new-brunswick.png" caption="Rutgers The State University of New Jersey -- New Brunswick" >}}

{{< figure src="pages-2-university-of-maryland,-college-park.png" caption="University of Maryland, College Park" >}}

{{< figure src="pages-2-university-of-california,-santa-barbara.png" caption="University of California, Santa Barbara" >}}

{{< figure src="pages-2-the-university-of-iowa.png" caption="The University of Iowa" >}}

{{< figure src="pages-2-emory-university.png" caption="Emory University" >}}

{{< figure src="pages-2-the-florida-state-university.png" caption="The Florida State University" >}}

{{< figure src="pages-2-brown-university.png" caption="Brown University" >}}

{{< figure src="pages-2-university-of-kansas.png" caption="University of Kansas" >}}

{{< figure src="pages-2-university-of-pittsburgh.png" caption="University of Pittsburgh" >}}

{{< figure src="pages-2-temple-university.png" caption="Temple University" >}}

{{< figure src="pages-2-michigan-state-university.png" caption="Michigan State University" >}}

{{< figure src="pages-2-university-of-notre-dame.png" caption="University of Notre Dame" >}}

{{< figure src="pages-2-boston-university.png" caption="Boston University" >}}

{{< figure src="pages-2-georgetown-university.png" caption="Georgetown University" >}}

{{< figure src="pages-2-university-of-southern-california.png" caption="University of Southern California" >}}

{{< figure src="pages-2-brandeis-university.png" caption="Brandeis University" >}}

{{< figure src="pages-2-university-of-missouri---columbia.png" caption="University of Missouri -- Columbia" >}}

{{< figure src="pages-2-state-university-of-new-york-at-binghamton.png" caption="State University of New York at Binghamton" >}}

{{< figure src="pages-2-the-pennsylvania-state-university.png" caption="The Pennsylvania State University" >}}

{{< figure src="pages-2-state-university-of-new-york-at-buffalo.png" caption="State University of New York at Buffalo" >}}

{{< figure src="pages-2-vanderbilt-university.png" caption="Vanderbilt University" >}}

{{< figure src="pages-2-university-of-california,-san-diego.png" caption="University of California, San Diego" >}}

{{< figure src="pages-2-the-university-of-arizona.png" caption="The University of Arizona" >}}

{{< figure src="pages-2-the-catholic-university-of-america.png" caption="The Catholic University of America" >}}

{{< figure src="pages-2-university-of-colorado-at-boulder.png" caption="University of Colorado at Boulder" >}}

{{< figure src="pages-2-syracuse-university.png" caption="Syracuse University" >}}

{{< figure src="pages-2-university-of-florida.png" caption="University of Florida" >}}

{{< figure src="pages-2-saint-louis-university.png" caption="Saint Louis University" >}}

{{< figure src="pages-2-the-university-of-new-mexico.png" caption="The University of New Mexico" >}}

{{< figure src="pages-2-the-claremont-graduate-university.png" caption="The Claremont Graduate University" >}}

{{< figure src="pages-2-the-george-washington-university.png" caption="The George Washington University" >}}

{{< figure src="pages-2-tulane-university.png" caption="Tulane University" >}}

{{< figure src="pages-2-university-of-massachusetts-amherst.png" caption="University of Massachusetts Amherst" >}}

{{< figure src="pages-2-university-of-california,-irvine.png" caption="University of California, Irvine" >}}

{{< figure src="pages-2-state-university-of-new-york-at-stony-brook.png" caption="State University of New York at Stony Brook" >}}

{{< figure src="pages-2-university-of-rochester.png" caption="University of Rochester" >}}

{{< figure src="pages-2-university-of-kentucky.png" caption="University of Kentucky" >}}

{{< figure src="pages-2-university-of-south-carolina.png" caption="University of South Carolina" >}}

{{< figure src="pages-2-university-of-delaware.png" caption="University of Delaware" >}}

{{< figure src="pages-2-university-of-hawai'i-at-manoa.png" caption="University of Hawai'i at Manoa" >}}

{{< figure src="pages-2-the-university-of-oklahoma.png" caption="The University of Oklahoma" >}}

{{< figure src="pages-2-university-of-california,-davis.png" caption="University of California, Davis" >}}

{{< figure src="pages-2-arizona-state-university.png" caption="Arizona State University" >}}

{{< figure src="pages-2-university-of-connecticut.png" caption="University of Connecticut" >}}

{{< figure src="pages-2-university-of-oregon.png" caption="University of Oregon" >}}

{{< figure src="pages-2-louisiana-state-university-and-agricultural-&-mechanical-college.png" caption="Louisiana State University and Agricultural & Mechanical College" >}}

{{< figure src="pages-2-university-of-georgia.png" caption="University of Georgia" >}}

{{< figure src="pages-2-american-university.png" caption="American University" >}}

{{< figure src="pages-2-washington-university-in-st.-louis.png" caption="Washington University in St. Louis" >}}

{{< figure src="pages-2-boston-college.png" caption="Boston College" >}}

{{< figure src="pages-2-the-university-of-nebraska---lincoln.png" caption="The University of Nebraska -- Lincoln" >}}

{{< figure src="pages-2-university-of-california,-riverside.png" caption="University of California, Riverside" >}}

{{< figure src="pages-2-case-western-reserve-university.png" caption="Case Western Reserve University" >}}

{{< figure src="pages-2-drew-university.png" caption="Drew University" >}}

{{< figure src="pages-2-fordham-university.png" caption="Fordham University" >}}

{{< figure src="pages-2-texas-a&m-university.png" caption="Texas A&M University" >}}

{{< figure src="pages-2-the-university-of-utah.png" caption="The University of Utah" >}}

{{< figure src="pages-2-rice-university.png" caption="Rice University" >}}

{{< figure src="pages-2-washington-state-university.png" caption="Washington State University" >}}

{{< figure src="pages-2-kent-state-university.png" caption="Kent State University" >}}

{{< figure src="pages-2-university-of-illinois-at-chicago.png" caption="University of Illinois at Chicago" >}}

{{< figure src="pages-2-purdue-university.png" caption="Purdue University" >}}

{{< figure src="pages-2-university-of-cincinnati.png" caption="University of Cincinnati" >}}

{{< figure src="pages-2-loyola-university-chicago.png" caption="Loyola University Chicago" >}}

{{< figure src="pages-2-howard-university.png" caption="Howard University" >}}

{{< figure src="pages-2-texas-christian-university.png" caption="Texas Christian University" >}}

{{< figure src="pages-2-ohio-university.png" caption="Ohio University" >}}

{{< figure src="pages-2-marquette-university.png" caption="Marquette University" >}}

{{< figure src="pages-2-the-university-of-tennessee.png" caption="The University of Tennessee" >}}

{{< figure src="pages-2-boston-university-graduate-school.png" caption="Boston University Graduate School" >}}

{{< figure src="pages-2-the-university-of-alabama.png" caption="The University of Alabama" >}}

{{< figure src="pages-2-graduate-theological-union.png" caption="Graduate Theological Union" >}}

{{< figure src="pages-2-the-college-of-william-and-mary.png" caption="The College of William and Mary" >}}

{{< figure src="pages-2-bryn-mawr-college.png" caption="Bryn Mawr College" >}}

{{< figure src="pages-2-university-of-california,-santa-cruz.png" caption="University of California, Santa Cruz" >}}

{{< figure src="pages-2-university-of-north-texas.png" caption="University of North Texas" >}}

{{< figure src="pages-2-the-southern-baptist-theological-seminary.png" caption="The Southern Baptist Theological Seminary" >}}

{{< figure src="pages-2-west-virginia-university.png" caption="West Virginia University" >}}

{{< figure src="pages-2-university-of-houston.png" caption="University of Houston" >}}

{{< figure src="pages-2-bowling-green-state-university.png" caption="Bowling Green State University" >}}

{{< figure src="pages-2-university-of-denver.png" caption="University of Denver" >}}

{{< figure src="pages-2-texas-tech-university.png" caption="Texas Tech University" >}}

{{< figure src="pages-2-auburn-university.png" caption="Auburn University" >}}

{{< figure src="pages-2-state-university-of-new-york-at-albany.png" caption="State University of New York at Albany" >}}

{{< figure src="pages-2-northern-illinois-university.png" caption="Northern Illinois University" >}}

{{< figure src="pages-2-the-union-institute.png" caption="The Union Institute" >}}

{{< figure src="pages-2-southern-illinois-university-at-carbondale.png" caption="Southern Illinois University at Carbondale" >}}

{{< figure src="pages-2-southwestern-baptist-theological-seminary.png" caption="Southwestern Baptist Theological Seminary" >}}

{{< figure src="pages-2-wayne-state-university.png" caption="Wayne State University" >}}

{{< figure src="pages-2-the-university-of-mississippi.png" caption="The University of Mississippi" >}}

{{< figure src="pages-2-miami-university.png" caption="Miami University" >}}

{{< figure src="pages-2-new-school-for-social-research.png" caption="New School for Social Research" >}}

{{< figure src="pages-2-princeton-theological-seminary.png" caption="Princeton Theological Seminary" >}}

{{< figure src="pages-2-oklahoma-state-university.png" caption="Oklahoma State University" >}}

{{< figure src="pages-2-the-university-of-maine.png" caption="The University of Maine" >}}

{{< figure src="pages-2-university-of-arkansas.png" caption="University of Arkansas" >}}

{{< figure src="pages-2-university-of-miami.png" caption="University of Miami" >}}

{{< figure src="pages-2-george-mason-university.png" caption="George Mason University" >}}

{{< figure src="pages-2-georgia-state-university.png" caption="Georgia State University" >}}

{{< figure src="pages-2-carnegie-mellon-university.png" caption="Carnegie Mellon University" >}}

{{< figure src="pages-2-fuller-theological-seminary.png" caption="Fuller Theological Seminary" >}}

{{< figure src="pages-2-mississippi-state-university.png" caption="Mississippi State University" >}}

{{< figure src="pages-2-university-of-new-hampshire.png" caption="University of New Hampshire" >}}

{{< figure src="pages-2-baylor-university.png" caption="Baylor University" >}}

{{< figure src="pages-2-kansas-state-university.png" caption="Kansas State University" >}}

{{< figure src="pages-2-clark-university.png" caption="Clark University" >}}

{{< figure src="pages-2-iowa-state-university.png" caption="Iowa State University" >}}

{{< figure src="pages-2-the-university-of-toledo.png" caption="The University of Toledo" >}}

{{< figure src="pages-2-the-university-of-southern-mississippi.png" caption="The University of Southern Mississippi" >}}

{{< figure src="pages-2-ball-state-university.png" caption="Ball State University" >}}

{{< figure src="pages-2-union-institute-and-university.png" caption="Union Institute and University" >}}

{{< figure src="pages-2-new-school-university.png" caption="New School University" >}}

{{< figure src="pages-2-lehigh-university.png" caption="Lehigh University" >}}

{{< figure src="pages-2-the-university-of-wisconsin---milwaukee.png" caption="The University of Wisconsin -- Milwaukee" >}}

{{< figure src="pages-2-the-university-of-texas-at-arlington.png" caption="The University of Texas at Arlington" >}}

{{< figure src="pages-2-tufts-university.png" caption="Tufts University" >}}

{{< figure src="pages-2-brigham-young-university.png" caption="Brigham Young University" >}}

{{< figure src="pages-2-the-university-of-memphis.png" caption="The University of Memphis" >}}

{{< figure src="pages-2-northern-arizona-university.png" caption="Northern Arizona University" >}}

{{< figure src="pages-2-the-university-of-texas-at-dallas.png" caption="The University of Texas at Dallas" >}}

{{< figure src="pages-2-union-theological-seminary.png" caption="Union Theological Seminary" >}}

{{< figure src="pages-2-southern-methodist-university.png" caption="Southern Methodist University" >}}

{{< figure src="pages-2-virginia-polytechnic-institute-and-state-university.png" caption="Virginia Polytechnic Institute and State University" >}}

{{< figure src="pages-2-westminster-theological-seminary.png" caption="Westminster Theological Seminary" >}}

{{< figure src="pages-2-duquesne-university.png" caption="Duquesne University" >}}

{{< figure src="pages-2-western-michigan-university.png" caption="Western Michigan University" >}}

{{< figure src="pages-2-university-of-idaho.png" caption="University of Idaho" >}}

{{< figure src="pages-2-california-institute-of-integral-studies.png" caption="California Institute of Integral Studies" >}}

{{< figure src="pages-2-regent-university.png" caption="Regent University" >}}

{{< figure src="pages-2-the-university-of-akron.png" caption="The University of Akron" >}}

{{< figure src="pages-2-the-jewish-theological-seminary-of-america.png" caption="The Jewish Theological Seminary of America" >}}

{{< figure src="pages-2-university-of-missouri---kansas-city.png" caption="University of Missouri -- Kansas City" >}}

{{< figure src="pages-2-florida-international-university.png" caption="Florida International University" >}}

{{< figure src="pages-2-pacifica-graduate-institute.png" caption="Pacifica Graduate Institute" >}}

{{< figure src="pages-2-university-of-nevada,-reno.png" caption="University of Nevada, Reno" >}}

{{< figure src="pages-2-virginia-commonwealth-university.png" caption="Virginia Commonwealth University" >}}

{{< figure src="pages-2-the-university-of-north-carolina-at-greensboro.png" caption="The University of North Carolina at Greensboro" >}}

{{< figure src="pages-2-trinity-evangelical-divinity-school.png" caption="Trinity Evangelical Divinity School" >}}

{{< figure src="pages-2-new-orleans-baptist-theological-seminary.png" caption="New Orleans Baptist Theological Seminary" >}}

{{< figure src="pages-2-university-of-louisville.png" caption="University of Louisville" >}}

{{< figure src="pages-2-georgia-institute-of-technology.png" caption="Georgia Institute of Technology" >}}

{{< figure src="pages-2-andrews-university.png" caption="Andrews University" >}}

{{< figure src="pages-2-northeastern-university.png" caption="Northeastern University" >}}

{{< figure src="pages-2-southeastern-baptist-theological-seminary.png" caption="Southeastern Baptist Theological Seminary" >}}

{{< figure src="pages-2-yeshiva-university.png" caption="Yeshiva University" >}}

{{< figure src="pages-2-the-iliff-school-of-theology-and-university-of-denver.png" caption="The Iliff School of Theology and University of Denver" >}}

{{< figure src="pages-2-university-of-nevada,-las-vegas.png" caption="University of Nevada, Las Vegas" >}}

{{< figure src="pages-2-oregon-state-university.png" caption="Oregon State University" >}}

{{< figure src="pages-2-rensselaer-polytechnic-institute.png" caption="Rensselaer Polytechnic Institute" >}}

{{< figure src="pages-2-mid-america-baptist-theological-seminary.png" caption="Mid-America Baptist Theological Seminary" >}}

{{< figure src="pages-2-memphis-state-university.png" caption="Memphis State University" >}}

{{< figure src="pages-2-university-of-california,-san-francisco.png" caption="University of California, San Francisco" >}}

{{< figure src="pages-2-texas-woman's-university.png" caption="Texas Woman's University" >}}

{{< figure src="pages-2-colorado-state-university.png" caption="Colorado State University" >}}

{{< figure src="pages-2-salve-regina-university.png" caption="Salve Regina University" >}}

{{< figure src="pages-2-indiana-university-of-pennsylvania.png" caption="Indiana University of Pennsylvania" >}}

{{< figure src="pages-2-university-of-south-florida.png" caption="University of South Florida" >}}

{{< figure src="pages-2-walden-university.png" caption="Walden University" >}}

{{< figure src="pages-2-lutheran-school-of-theology-at-chicago.png" caption="Lutheran School of Theology at Chicago" >}}

{{< figure src="pages-2-middle-tennessee-state-university.png" caption="Middle Tennessee State University" >}}

{{< figure src="pages-2-university-of-wyoming.png" caption="University of Wyoming" >}}

{{< figure src="pages-2-capella-university.png" caption="Capella University" >}}

{{< figure src="pages-2-dallas-theological-seminary.png" caption="Dallas Theological Seminary" >}}

{{< figure src="pages-2-university-of-new-orleans.png" caption="University of New Orleans" >}}
