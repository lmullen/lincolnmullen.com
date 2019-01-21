---
title: "Clio 2: Computational History"
semester: Spring
year: 2018
department: Department of History and Art History
university: George Mason University
coursenum: HIST 697-001
courseurl: https://historyarthistory.gmu.edu/courses/hist697/course_sections/34585
credits: 3
meets: Mondays, 7:20–10:00pm
room: RRCHNM conference room, Research Hall 402
instructor: Lincoln Mullen
instructorurl: https://lincolnmullen.com
email: lmullen@gmu.edu
office: Research Hall 457
officehours: "by appointment"
toc: true
draft: false
---

# Course description

In this course you will learn to apply computational methods to create historical arguments. You will learn to work with historical data, including finding, gathering, manipulating, analyzing, visualizing, and arguing from data, with special attention to geospatial, textual, and network data. These methods will be taught primarily through scripting in the [R programming language](https://www.r-project.org/). While historical methods can be applied to many topics and time periods, they cannot be understood separate from how the discipline forms meaningful questions and interpretations, nor divorced from the particularities of the sources and histories of some specific topic. You will therefore work through a series of example problems using datasets from the history of the nineteenth-century U.S. religion, and you will apply these methods to a dataset in your own field of research.

# Learning goals

After taking this course, you will be able to

-   perform exploratory data analysis; clean, tidy, and manipulate data; gather historical data from print and manuscript sources; use existing historical data sets; create common visualizations; work with geospatial, textual, and network data.
-   write scripts using the R programming language and its extensive set of packages.
-   understand the place of data analysis and visualization within humanities computing, digital history, and the discipline of history.
-   conceive of and execute a research project in computational history suitable for treatment in a dissertation chapter or journal article.
-   take the course "Programming in History/New Media," a.k.a. Clio 3, should you choose.

# Essential information

You are always welcome to [book an appointment](/page/office-hours/) during my office hours. If the times that are available do not work for you, feel free to contact me. All communication for this course will happen in [our Slack group](https://mason-dh-grad.slack.com/signup). Read this [getting started guide](https://get.slack.help/hc/en-us/articles/218080037-Getting-started-for-new-users) if you need help.

Bring a computer to each class meeting. We will use [R](https://cran.rstudio.com/) and [RStudio](https://www.rstudio.com/products/rstudio/#Desktop). [Install them](http://dh-r.lincolnmullen.com/installing-r-and-packages.html) on your own computer. You will also have access to an [RStudio Server instance](http://rstudio.chnm.org/) which will let you use R in your browser. Much of your work for the course will go on [GitHub](https://github.com/), so sign up for an account.

### Main texts

All required readings are available online for free or through the GMU libraries, though they can also be purchased (sometimes in more complete editions) in print or e-books. These are the books we will use most frequently.

- Shawn Graham, Ian Milligan, and Scott Weingart, [*Exploring Big Historical Data: The Historian's Macroscope*](http://www.themacroscope.org/2.0/) (Imperial College Press, 2015).
- Kieran Healy, [*Data Visualization: A Practical Introduction*](http://socviz.co/) (Princeton University Press, forthcoming 2018). 
- Roger Peng, [*Exploratory Data Analysis with R*](https://bookdown.org/rdpeng/exdata/) (Leanpub, 2016). 
- Julia Silge and David Robinson, [*Text Mining with R: A Tidy Approach*](https://www.tidytextmining.com/) (O'Reilly, 2017).
- Hadley Wickham and Garrett Grolemund, [*R for Data Science: Import, Tidy, Transform, Visualize, and Model Data*](http://r4ds.had.co.nz/) (O'Reilly, 2017).

# Assignments

**Be prepared**. Preparation and participation are expected as a matter of course in a graduate class. Complete all readings and assignments before class. If the readings include sample code or questions at the end, work through them as part of doing the readings.

**Worksheets and weekly assignments (20%).** Many classes will have an assignment due before class begins. Some will require you to do library research; others will be practice data analysis worksheets. Some of the questions on the worksheet will be easy; most will be difficult; some you may find nearly impossible. The aim is to *practice*. We will go over the worksheets in class each week. If you attempt a problem and can't solve it, you should still turn in whatever work you did on it. Students who complete all the easy and moderate difficulty questions, attempt the very difficult questions, and ask for help as needed will do just fine. These assignments will graded by completion, with three levels: "incomplete," "acceptable," "excellent." Unless otherwise specified, these assignments should be submitted as a PDF or a standalone HTML file, one file per assignment. Name them like this: `Mullen-worksheet-week02.pdf`. Submit them to [this Dropbox folder](https://www.dropbox.com/request/grN2YG0bS4ycamVCjHXX).

**Analysis assignments (3 @ 15% each).** You will do three analysis assignments, each demonstrating a specific skill in data analysis. For these assignments you will be given a historical dataset and asked some interpretative questions. You will prepare an [RMarkdown document](http://rmarkdown.rstudio.com/) containing prose, code, and tables or visualizations to answer the historical questions and, as necessary, explain your methods. You will be given a starter GitHub repository with the data and questions. Submit your final analysis as an HTML file along with your R Markdown file to [this Dropbox folder](https://www.dropbox.com/request/grN2YG0bS4ycamVCjHXX). <del>You will also be evaluated on the code in your GitHub repository, which I must be able to run on my computer.</del>

- [Mapping assignment](https://github.com/ClioGMU/clio2-mapping)
- [Text analysis assignment](https://github.com/ClioGMU/clio2-text)

**R package tutorial (10%).** At our second meeting, you will pick from a list of R packages not covered in this class. You will be assigned a week (beginning at week 7) during which you will teach the class for 15 minutes about the topic you selected. As part of that teaching, you will prepare a PDF handout. That handout should include these parts: (1) a one- to two-paragraph summary of what the package does and while it is useful; (2) a brief section of example code and results; (3) a bulleted list of examples (historical if possible) where the package was used. A draft of that handout is due to me one week before you are scheduled to teach. I will offer feedback, and you will give the class a revised version in Slack on the Friday before you teach.

**Research paper (25%).** You will write one research paper suitable for a presentation at a disciplinary or digital humanities conference (see for example the CFP for [Current Research in Digital History](http://crdh.rrchnm.org/), or the CFP for the major conference in your field). This paper must advance a historical argument using data analysis of a set of sources that you choose from your research interests. Submit this paper as a PDF or self-contained HTML file (if it includes interactive visualizations) to [this Dropbox folder](https://www.dropbox.com/request/grN2YG0bS4ycamVCjHXX). Further instructions will be given throughout the semester. **Due May 10.**

# Schedule

### Week 1 (Jan. 22): Introduction

Assignment: 

- Join the [Slack group](https://mason-dh-grad.slack.com/signup).
- [Install R](http://dh-r.lincolnmullen.com/installing-r-and-packages.html) and other necessary software.

Read:

- Arguing with Digital History working group, "[Digital History and Argument](https://rrchnm.org/argument-white-paper/)," white paper, Roy Rosenzweig Center for History and New Media (November 13, 2017).
- Read "[Introduction](http://dh-r.lincolnmullen.com/introduction.html)" and "[Getting Started](http://dh-r.lincolnmullen.com/getting-started.html)" from *Computational Historical Thinking*. 

### Week 2 (Jan. 29.): Data from history and historians

Assignment:

- [Getting familiar with R worksheet](http://dh-r.lincolnmullen.com/worksheets.html).
- Find primary source data tables, datasets, or corpora from your field of historical research. At least one of these must be a source which can be transcribed into a tabular dataset in a later week. Post full citations and URLs in the Slack group, along with a sentence or two explaining what you've found. Examine the links that other people post before class.

Read:

-   Wickham and Grolemund, *R for Data Science*, ch. 1, 4, 6, 8.
-   Shari Rabin, "'Let us Endeavor to Count Them Up': The Nineteenth-Century Origins of American Jewish Demography," *American Jewish History* 101, no 4 (2017): 419--440, <https://doi.org/10.1353/ajh.2017.0060>.
-   Roger Finke and Rodney Stark, [*The Churching of America, 1776-2005: Winners and Losers in Our Religious Economy*](http://www.jstor.org/stable/j.ctt5hhwnc) (Rutgers University Press, 2005), ch. 1, 5.
-   Laurie F. Maffly-Kipp, "[If It's South Dakota You Must Be Episcopalian: Lies, Truth-Telling, and the Mapping of U.S. Religion](http://www.jstor.org/stable/4146694)" *Church History* 71, no. 1 (2002): 132--42.

Browse:

-   Fletcher W. Hewes and Henry Garnett, [*Scribner's Statistical Atlas of the United States Showing by Graphic Methods Their Present Condition and Their Political, Social and Industrial Development*](http://www.davidrumsey.com/luna/servlet/s/rod36m) (New York: Charles Scribner's Sons, 1883), plates 58 to 61.
-   Herman Carl Weber, [*Presbyterian Statistics through One Hundred Years, 1826-1926*](http://catalog.hathitrust.org/Record/007109885) (Philadelphia: Presbyterian Church in the U.S.A., 1927), part II

### Week 3 (Feb. 5): Data manipulation

Assignment: 

-   [Data structures worksheet](http://dh-r.lincolnmullen.com/worksheets.html).
-   Transcribe at least 50 rows of the historical data you found last week. Be prepared to describe in class how you decided on the structure of your data.

Read:

-   Wickham and Grolemund, *R for Data Science*, ch. 5, 12.
-   Karl W. Broman and Kara H. Woo, "Data Organization in Spreadsheets" PeerJ Preprints 5:e3183v1 (2017): <https://doi.org/10.7287/peerj.preprints.3183v1>.
-   Frederick W. Gibbs, "[New Forms of History: Critiquing Data and Its Representations](http://tah.oah.org/february-2016/new-forms-of-history-critiquing-data-and-its-representations/)," *The American Historian* (February 2016).
-   Graham, Milligan, Weingart, *Macroscope*, ch. 1--2.
-   William G. Thomas III, "Computing and the Historical Imagination," in [*A Companion to Digital Humanities*](http://www.digitalhumanities.org/companion), ed. Susan Schreibman, Ray Siemens, John Unsworth (Blackwell, 2004).
- Susan Hockey, "The History of Humanities Computing," in [*A Companion to Digital Humanities*](http://www.digitalhumanities.org/companion), ed. Susan Schreibman, Ray Siemens, John Unsworth (Blackwell, 2004).

### Week 4 (Feb. 12): Data visualization

Assignment: 

- [First data manipulation worksheet](http://dh-r.lincolnmullen.com/worksheets.html).

Read:

-   Wickham and Grolemund, *R for Data Science*, ch. 3.
-   Healy, *Data Visualization*, ch. 1, 3.
-   Graham, Milligan, Weingart, *Macroscope*, ch 5.
-   Kieran Healy and James Moody, "[Data Visualization in Sociology](http://kieranhealy.org/files/papers/data-visualization.pdf)" *Annual Review of Sociology*, 40:105--128.

### Week 5 (Feb. 19): Data manipulation and visualizations

Assignment: 

- [First ggplot2 worksheet](http://dh-r.lincolnmullen.com/worksheets.html).

Read:

-   Wickham and Grolemund, *R for Data Science*, ch. 12--13.
-   Healy, *Data Visualization*, ch 4--5, 8.
-   John Theibault, "Visualizations and Historical Arguments," in *Writing History in the Digital Age*, ed. Kristen Nawrotzki and Jack Dougherty (University of Michigan Press, 2013), <https://doi.org/10.3998/dh.12230987.0001.001>.

### Week 6 (Feb. 26): Exploratory data analysis

Assignment: 

- [Second data manipulation worksheet](http://dh-r.lincolnmullen.com/worksheets.html). 

Read:

-   Wickham and Grolemund, *R for Data Science*, ch. 7, 19, 21.
-   Peng, *Exploratory Data Analysis*, ch. 1, 4--6.
-   Jenny Bryan et al, [*Happy Git and GitHub for the useR*](http://happygitwithr.com/).
-   Skim: Wickham and Grolemund, *R for Data Science*, ch. 27, 29.3-29.4, 30.

### Week 7 (Mar. 5): Mapping


Assignment: 

- [Functions worksheet](http://dh-r.lincolnmullen.com/worksheets.html).

Read:

-  Healy, *Data Visualization*, ch 7.
-  Documentation for the [leaflet package](https://rstudio.github.io/leaflet/).
-  Richard White "[What is Spatial History?](http://web.stanford.edu/group/spatialhistory/cgi-bin/site/pub.php?id=29)."
-  Cameron Blevins, "Space, Nation, and the Triumph of Region: A View of the World from Houston," *Journal of American History* 101, no. 1 (2014): 122--47, https://doi.org/10.1093/jahist/jau184.

### Week 8 (Mar. 19): Mapping

Read:

-   Todd Presner and David Shepard, "Mapping the Geospatial Turn," in _A New Companion to Digital Humanities_, ed. Susan Schreibman, Ray Siemens, and John Unsworth (Wiley, 2016), 201--212. [GMU library](http://mutex.gmu.edu/login?url=http://www.gmu.eblib.com/EBLWeb/patron/?target=patron&extendedid=P_4093339_0)
-   Bret E. Carroll, "Spatial Approaches to American Religious Studies, _Oxford Research Encyclopedia of Religion_ (Oxford University Press, 2015), <https://doi.org/10.1093/acrefore/9780199340378.013.13>.
-   Stephen Robertson, "[Putting Harlem on the Map](http://quod.lib.umich.edu/d/dh/12230987.0001.001/1:8/--writing-history-in-the-digital-age?g=dculture;rgn=div1;view=fulltext;xc=1#8.2)," in *Writing History in the Digital Age*, edited by Jack Dougherty and Kristen Nawrotzki (University of Michigan Press, 2013).

### Week 9 (Mar. 26): Text analysis


Assignment: 

- Mapping assignment due (see GitHub repository for data and instructions).

Read:

-   Silge and Robinson, *Tidy Text Mining with R*, ch. 1--2, 4--7.
-   Wickham and Grolemund, *R for Data Science*, 14.
-   Graham, Milligan, Weingart, *Macroscope*, chs. 3--4.
-   Tim Hitchcock and William J. Turkel, "The _Old Bailey Proceedings, 1674--1913_: Text Mining for Evidence of Court Behavior," _Law and History Review_ 34, no. 4 (2016): 929--955, <https://doi.org/10.1017/S0738248016000304>.

Tutorials:

- Caitlin: forcats
- Greta: lubridate

### Week 10 (Apr. 2): Text analysis

Read:

-  Matthew L. Jockers and Ted Underwood, "Text-Mining the Humanities" in _A New Companion to Digital Humanities_, ed. Susan Schreibman, Ray Siemens, and John Unsworth (Wiley, 2016), 291--306. [GMU library](http://mutex.gmu.edu/login?url=http://www.gmu.eblib.com/EBLWeb/patron/?target=patron&extendedid=P_4093339_0)
-   Matthew K. Gold et al., "[Forum: Text Analysis at Scale](http://dhdebates.gc.cuny.edu/debates/text/93)," in *Debates in the Digital Humanities 2016* (University of Minnesota Press, 2016), 525--568.
-   Ryan Cordell, "Reprinting, Circulation, and the Network Author in Antebellum Newspapers," *American Literary History* 27, no. 3 (2015): 417--445, <https://doi.org/10.1093/alh/ajv028>.
-   David A. Smith, Ryan Cordell, and Abby Mullen, "Computational Methods for Uncovering Reprinted Texts in Antebellum Newspapers," *American Literary History* 27, no. 3 (2015): E1--E15, <https://doi.org/10.1093/alh/ajv029>.

Tutorials:

- Andrew: DT
- Chris: stringr

### Week 11 (Apr. 9): Network analysis


Assignment: 

- Text analysis assignment due (see GitHub repository for data and instructions).

Read:

-   Graham, Milligan, Weingart, *Macroscope*, ch. 6--7.
- Rebecca Sutton Koeser, "[Trusting Others to 'Do the Math'](https://doi.org/10.1080/03080188.2016.1165454)" _Interdisciplinary Science Reviews_ 40, no. 4 (2015): 376--392, <https://doi.org/10.1080/03080188.2016.1165454>.
- Benjamin Schmidt, "[Do Digital Humanists Need to Understand Algorithms?](http://dhdebates.gc.cuny.edu/debates/text/99)" in _Debates in the Digital Humanities 2016_, ed. Matthew K. Gold and Lauren F. Klein (University of Minnesota Press, 2016).

Tutorials:

- Clarke: rvest
- Brian: ggrepel

### Week 12 (Apr. 16): Network analysis

Read:

- Matthew Lincoln, "Social Network Centralization Dynamics in Print Production in the Low Countries, 1550--1750," *International Journal for Digital Art History* 2 (2016): 134--157, <https://doi.org/10.11588/dah.2016.2.25337>.
-   "[AHR Forum: Mapping the Republic of Letters](https://academic-oup-com.mutex.gmu.edu/ahr/search-results?f_TocHeadingTitle=AHR%20Forum:%20Mapping%20the%20Republic%20of%20Letters)," _American Historical Review_ 122, no. 2 (2017): 399--463. 

Tutorials:

- Kenny: magick
- Jay: iheatmapr

### Week 13 (Apr. 23): TBD

Topic and readings to be determined by the needs of student research papers.

Tutorials:

- Alan: dygraphs
- John: Bookdown
- Spencer: Shiny

### Week 14 (Apr. 30): TBD

Topic and readings to be determined by the needs of student research papers.

# Fine print

This syllabus may be updated online as necessary. The online version of this syllabus is the only authoritative version.

Students must satisfactorily complete all assignments (including participation assignments) in order to pass this course. Your attendance is expected at every meeting. If you must be absent, I request that you notify me in advance of the class meeting. I am sometimes willing to grant extensions for cause, but you must request an extension before the assignment's due date. For every day or part of a day that an assignment is late without an extension, I may reduce your grade. No work (other than final exams and final projects) will be accepted later than the last day that the class meets. I will discuss grades only in person during office hours.

See the [George Mason University catalog](http://catalog.gmu.edu/) for general policies, as well as the university [statement on diversity](http://ctfe.gmu.edu/professional-development/mason-diversity-statement/). You are expected to know and follow George Mason's policies on [academic integrity](http://oai.gmu.edu/) and the [honor code](http://oai.gmu.edu/understanding-the-honor-code/). If you are a student with a disability and you need academic accommodations, please see me and contact the Office of Disability Services at 703-993-2474 or through [their website](http://ods.gmu.edu). You are responsible for verifying your enrollment status. All academic accommodations must be arranged through that office. Please note these dates from the [academic calendar](http://registrar.gmu.edu/calendars/spring-2018/).

- Last day to add a class or drop a class without penalty: January 29, 2018.
- Last day to drop a class without special permission: February 23, 2018.
