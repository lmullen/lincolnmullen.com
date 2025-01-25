---
coursenum: 'HIST 697-001'
courseurl: 'https://historyarthistory.gmu.edu/courses/hist697/course_sections/110323'
credits: 3
department: Department of History and Art History
instructor: Lincoln Mullen
instructorurl: 'https://lincolnmullen.com'
meets: 'Mondays, 7:20--10:00pm'
office: Research Hall 483
officehours: By appointment
room: Krug Hall 209
semester: Spring
title: Data Analysis for History
toc: true
university: George Mason University
year: 2025
---

## Course description

In this course you will learn to use computational methods to create historical interpretations. You will work with historical data, which includes finding, gathering, manipulating, analyzing, visualizing, and arguing from datasets, with special attention to geospatial and network data. These methods will be taught primarily using programming languages for data analysis. While data analysis methods can be applied to many topics and time periods, they cannot be understood separately from how the discipline forms meaningful questions and interpretations, nor divorced from the particularities of the sources and histories of some specific topic. You will therefore work through a series of example problems using provided datasets and then apply the methods to write a research paper using a dataset from your own historical field.

## Learning goals

After taking this course, you will be able to

- gather historical data from print and manuscript sources; use existing historical data sets; clean, tidy, and manipulate data; perform exploratory data analysis; create common visualizations; work with geospatial and network data.
- create visualizations for the web and for print in open-source JavaScript libraries.
- understand the place of data analysis and visualization within the field of digital history and the discipline of history.
- conceive of and execute a short research project in computational history.

## Essential information

All of the required readings for this course are available online or through the GMU libraries. Readings are listed on the syllabus below. This is a graduate methods course in a field that moves quickly. The syllabus is likely to change over the course of the semester, and I may change or add to the readings at any time.

Communication for this course will happen in a channel in the RRCHNM Slack group. You will receive an invitation to that Slack channel. Read this [getting started guide](https://get.slack.help/hc/en-us/articles/218080037-Getting-started-for-new-users) if you need help. The Slack group is your primary place to ask for help. Please ask for help in the public channel rather than private messages. You are almost certainly not the only person to have your question, and asking and answering questions publicly benefits everyone. When you ask a question, help me help you by including the code that you are asking about and any error messages that are relevant. If possible, share a link to the work directly, such as to an Observable notebook.

You are always welcome to talk with me during office hours. My [office hours page](/page/office-hours/) has instructions on how to book an appointment. Make an effort to meet with me in person before you request a meeting via Zoom. If the scheduled times don't work for you, please contact me and suggest a few other times that would work for you.

Bring a computer to each class meeting. Some, but not all, of the software we are using can be accessed through a web browser, but you will still need a computer. I also recommend that you install some key software. See the list under the heading for the first week. (I only know how to help  you with a computer with some kind of Unix-like operating system available: macOS or a Linux distribution. But all of the software we will use has Windows support.)

## Assignments

Submit each assignment to Canvas unless otherwise instructed. If the assignment lives online, send me a URL through Canvas but make sure that I have permissions to see it. Submit the assignments before the start of class on the day on which they are due.

Preparation and participation are expected as a matter of course in a graduate class. Complete all readings and submit all assignments before class. If the readings include sample code or questions at the end, work through them as part of doing the readings, though you do not need to submit them and I will not check them.

Final grades will be calculated using the typical percentage-based grading scale (A = 93--100, A- = 90--92, B+ = 88--89, B = 83--87, B- = 80--82, ... F = 0--59).

**Weekly assignments (30%).** Many classes will have an assignment due before class begins. Most of those will require you to work through the material for the previous week. The aim is to _practice_. We will go over these assignments in class each week as a part of the learning process. If you attempt a problem and can't solve it, you should still turn in whatever work you did on it. These assignments will graded by completion.

**Analysis assignments (3 × 10% = 30%).** You will do three analysis assignments, each demonstrating a specific skill in data analysis. For these assignments you will use a historical dataset and be asked some interpretative questions. You will prepare an Observable notebook containing prose, code, and tables or visualizations to answer the historical questions and, as necessary, explain your methods. For these assignment I will provide a dataset that you can work with, but you can substitute your own dataset if you prefer.

**Final project (draft = 10 %; final = 30%).** For your final project, you will write a paper, based on data analysis and visualization, that could be presented at a historical conference. Like the analysis assignments, this paper will be written in an Observable notebook, which you can use to create the visualizations and tables. This paper should include more prose and citations, not to exceed 2,000 words. The prose, visualizations, and tables should be refined to the level of quality that would be expected in a published article. Each table and figure must have a caption written in complete sentences. The primary point of your paper should be advancing some well-argued interpretation of the past. Explain your methods as needed, but write in a way which would be understandable and compelling to any historian working in your field. Rough draft to be circulated before class on **Monday, May 5.** Final draft due **Monday, May 12.**

## Schedule

### Week 1 (January 27): History of data visualization

Readings:

- Lauren Klein et al., [_Data by Design: An Interactive History of Data Visualization, 1789--1900_](https://dataxdesign.io/) (MIT Press, 2024).
- Frederick W. Gibbs, "[New Forms of History: Critiquing Data and Its Representations](http://tah.oah.org/february-2016/new-forms-of-history-critiquing-data-and-its-representations/)," _The American Historian_ (February 2016).
- Taylor Arnold and Lauren Tilton, "[New Data: The Role of Statistics in DH](https://dhdebates.gc.cuny.edu/read/untitled-f2acf72c-a469-49d8-be35-67f9ac1e3a60/section/a2a6a192-f04a-4082-afaa-97c76a75b21c#ch24)," in _Debates in DH 2019_, ed. Matthew K. Gold and Lauren F. Klein (University of Minnesota Press, 2019).

We will set these up the first day of class:

- Join the class Slack channel.
- Get a [GitHub](https://github.com/) account and post it to the Slack group (e.g., I am `lmullen` and this is [my GitHub profile](https://github.com/lmullen)).
- Get an account at [Observable](https://observablehq.com).
- Install [DB Browser for SQLite](https://sqlitebrowser.org).
- Install [Visual Studio Code](https://code.visualstudio.com).
- Install a terminal. For Mac I recommend either [iTerm2](https://iterm2.com) or [Warp](https://www.warp.dev).
- (Optional) Install [Homebrew](https://brew.sh), only if you use macOS.
- (Optional) Consider [Claude](https://claude.ai/) or its ilk for support writing code and queries.

### Week 2 (February 3): From sources to data to history

Assignment:

- Find at least three primary source data tables, datasets, or corpora from your field of historical research. These could include sources that are in print or manuscript, as well as datasets that have already been created. Post citations and URLs in the Slack group, along with a sentence or two explaining what you've found. Examine the links that other people post before class.
- Find one example of a digital history project that uses visualization or data analysis. Post the URL to the Slack group and be ready in class with brief answer to these questions: What is interesting or insightful about this project? What did this project do that you would like to learn how to do for your own research?

Readings:

- Hadley Wickham, "Tidy Data," _Journal of Statistical Software_ 59 (2014): 1--23: <https://doi.org/10.18637/jss.v059.i10>.
- Karl W. Broman and Kara H. Woo, "Data Organization in Spreadsheets," _American Statistician_ 72, no. 1 (2018): 2--10, <https://doi.org/10.1080/00031305.2017.1375989>.
- Chad Gaffield, "Words, Words, Words: How the Digital Humanities Are Integrating Diverse Research Fields to Study People," _Annual Review of Statistics and Its Application_ 5, no. 1 (2018): 119--39, <https://doi.org/10.1146/annurev-statistics-031017-100547>.
- Abraham Gibson and Cindy Ermus, "The History of Science and the Science of History: Computational Methods, Algorithms, and the Future of the Field," _Isis_ 110, no. 3 (2019): 555--66, <https://doi.org/10.1086/705543>.
- Jessica Marie Johnson, "Markup Bodies: Black \[Life\] Studies and Slavery \[Death\] Studies at the Digital Crossroads," _Social Text_ 36, no. 4 (2018): 57--79, <https://doi.org/10.1215/01642472-7145658>.
- Roger Finke and Rodney Stark, [_The Churching of America, 1776-2005: Winners and Losers in Our Religious Economy_](http://www.jstor.org/stable/j.ctt5hhwnc) (Rutgers University Press, 2005), ch. 1.

Browse:

- Herman Carl Weber, [_Presbyterian Statistics through One Hundred Years, 1826-1926_](http://catalog.hathitrust.org/Record/007109885) (Philadelphia: Presbyterian Church in the U.S.A., 1927), part II.

### Week 3 (February 10): Data visualization basics

Assignment:

- Either use a primary source dataset that you found last week or, as a backup, the [_Minutes_](https://catalog.hathitrust.org/Record/006771726) of the Methodist Episcopal Church from after 1851. Create a well-structured spreadsheet and transcribe some substantial portion of the data. Upload a CSV file to Slack before class. Be prepared to describe in class how you decided on the structure of your data and how you identified what the observations and variables were. Use the Broman and Woo article as a guide.

Readings:

- [Observable Notebooks documentation](https://observablehq.com/documentation/notebooks/). Focus on "Cells > Data table," "Cells > Chart," and the "Data" sections in particular.
- Claus O. Wilke, _Fundamentals of Data Visualization_ (O'Reilly, 2019). At a minimum, read ch. 2, 3, 4, 6, 7, 12, 13.
- Kieran Healy and James Moody, "[Data Visualization in Sociology](http://kieranhealy.org/files/papers/data-visualization.pdf)" _Annual Review of Sociology_, 40:105--128.
- John Theibault, "Visualizations and Historical Arguments," in _Writing History in the Digital Age_, ed. Kristen Nawrotzki and Jack Dougherty (University of Michigan Press, 2013), <https://doi.org/10.3998/dh.12230987.0001.001>.

### Week 4 (February 17): Databases

Assignment:

- Create an Observable notebook. Make at least three visualizations using the Chart feature. You can use either the data you have transcribed, if it is meaningfully complete, or a sample dataset you can find here or elsewhere.

Readings:

- Jay A. Kreibich, _Using SQLite_ (O'Reilly, 2010), ch. 1, 2, 4, 5, 6.
- Stephen Ramsay, "[Databases](https://companions.digitalhumanities.org/DH/?chapter=content/9781405103213_chapter_15.html)" in _A Companion to Digital Humanities_, ed. Susan Schriebman, Ray Siemens, and John Unsworth (Blackwell, 2008).
- "[Spreasheet Thinking vs. Database Thinking](https://eagereyes.org/blog/2016/spreadsheet-thinking-vs-database-thinking)"

For future reference, if you are looking to use PostgreSQL instead of SQLite: Anthony DeBarros, _Practical SQL_, 2nd ed. (No Starch Press, 2022).

### Week 5 (February 24): Data manipulation

Assignment:

- Create your own SQLite database using historical data, and come prepared to class to explain your design decisions. Use the following features: (1) multiple tables, (2) appropriately typed data, (3) primary keys (and possibly natural primary keys), (4) foreign keys, (5) indices.

Readings:

- Arquero [documentation](https://idl.uw.edu/arquero/), but especially the "[Introducing Arquero](https://observablehq.com/@uwdata/introducing-arquero)" notebook. See the [other notebooks](https://observablehq.com/collection/@uwdata/arquero) too.
- Note that while Arquero is written in JavaScript and suitable for Observable Notebooks, you may wish to read the chapter on "[Data Tidying](https://r4ds.hadley.nz/data-tidy.html)" in Hadley Wickham, Mine Çetinkaya-Rundel, and Garrett Grolemund, _R for Data Science_, 2nd ed. (O'Reilly, 2023), even though we are not using R.
- [Observable JavaScript documentation](https://observablehq.com/documentation/cells/observable-javascript). You may wish to consult Nick Morgan, _JavaScript Crash Course_ (No Starch Press, 2024) or Marijn Haverbeke, _Eloquent JavaScript_, 4th ed. (No Starch Press, 2024).

### Week 6 (March 3): Observable Plot

Assignment:

- I will provide you with a notebook with a sample dataset and a set of problems that can be solved by manipulating the data. Fork this worksheet and complete it.

Readings:

- Nathan Yau, _Visualize This_, 2nd ed. (Wiley, 2024), ch. 1, 3, 4, 5.
- [Observable Plot documentation](https://observablehq.com/plot/).
- [Observable Plot gallery](https://observablehq.com/@observablehq/plot-gallery).

### Spring break (March 10)

### Week 7 (March 17): Exploratory data analysis

Assignment:

- Create an Observable notebook. Make at least three visualizations using Observable plot feature. You can use either the data you have transcribed, if it is meaningfully complete, or a sample dataset you can find here or elsewhere.

Readings:

- "[Exploratory Data Analysis](https://r4ds.hadley.nz/eda)" in Hadley Wickham, Mine Çetinkaya-Rundel, and Garrett Grolemund, _R for Data Science_, 2nd ed. (O'Reilly, 2023). Focus on the questions, rather than the code, since we are using other tools besides R.
- Jordan F. Bratt, "[Congressional Incumbency in the Early Republic](https://earlyamericanelections.org/essays/05-bratt-incumbency.html)," _Mapping Early American Elections_ (RRCHNM, 2019).

### Week 8 (March 24): Maps

Assignment:

- Complete the exploratory data analysis assignment.
- Sign up for an account at [NHGIS](https://www.nhgis.org) and download some data of interest to you.

Readings:

- "[Mapping in Plot](https://observablehq.com/@observablehq/plot-mapping)"
- "[Build Your First Map with Observable Plot](https://observablehq.com/@observablehq/build-your-first-map-with-observable-plot)"
- "[Build Your First Choropleth Map with Observable Plot](https://observablehq.com/@observablehq/build-your-first-choropleth-map-with-observable-plot)"
- Robert K. Nelson and Edward L. Ayers, eds., [_American Panorama: An Atlas of United States History_](https://dsl.richmond.edu/panorama/) (Digital Scholarship Lab, University of Richmond).
- Richard White, "[What is Spatial History?](https://web.stanford.edu/group/spatialhistory/static/publications/pub29.html)," _Spatial History Project_ (Stanford University, 2010).
- Cameron Blevins, "Space, Nation, and the Triumph of Region: A View of the World from Houston," _Journal of American History_ 101, no. 1 (2014): 122--47, <https://doi.org/10.1093/jahist/jau184>.

### Week 9 (March 31): Maps

Readings:

- Nathan Yau, _Visualize This_, 2nd ed. (Wiley, 2024), ch. 7.
- Claus O. Wilke, _Fundamentals of Data Visualization_ (O'Reilly, 2019), ch. 15, 22.
- Greta Swain, "[Maryland's Political Geography in the Early Republic](https://earlyamericanelections.org/essays/06-swain-geography-of-politics.html)," _Mapping Early American Elections_ (RRCHNM, 2019).

### Week 10 (April 7): Networks

Assignment:

- Complete the mapping analysis assignment.

Readings:

- "[Graph Visualization Introduction](https://observablehq.com/@bstaats/graph-visualization-introduction)" and "[Networks](https://observablehq.com/@nyuvis/networks)."
- Mark E. J. Newman, [_Networks: An Introduction_](https://math.bme.hu/~gabor/oktatas/SztoM/Newman_Networks.pdf) (Oxford University Press, 2010), ch. 1, 3, 4. Skim ch. 6, 7.
- Joshua Catalano, "Digitally Analyzing the Uneven Ground: Language Borrowing Among Indian Treaties," _Current Research in Digital History_ 1 (2018): <https://doi.org/10.31835/crdh.2018.02>.
- Matthew Lincoln, "Social Network Centralization Dynamics in Print Production in the Low Countries, 1550--1750," _International Journal for Digital Art History_ 2 (2016): 134--157, <https://doi.org/10.11588/dah.2016.2.25337>.

### Week 11 (April 14): LLMs for historical research

Assignment:

- Complete the networks analysis assignment.

Readings:

- [Forum](https://academic.oup.com/ahr/issue/128/3) on "Artificial Intelligence and the Practice of History," _American Historical Review_ 128, no. 3 (2023): 1345--1389.
- [Documentation for Llamafile](https://github.com/Mozilla-Ocho/llamafile) and [Llama](https://www.llama.com).

### Week 12 (April 21): Image analysis

Readings:

- Taylor Arnold and Lauren Tilton, _Distant Viewing: Computational Exploration of Digital Images_ (MIT Press, 2023): <https://doi.org/10.7551/mitpress/14046.001.0001>.

### Week 13 (April 28): Project workshop

Topics for this week will be determined by students' needs for the final research project. Come to class prepared to discuss your work in progress.

Readings:

- Jeri Wieringa, "[A Gospel of Health and Salvation](http://dissertation.jeriwieringa.com)" (PhD dissertation, George Mason University, 2019).
- Rebecca Sutton Koeser, "[Trusting Others to 'Do the Math'](https://doi.org/10.1080/03080188.2016.1165454)" _Interdisciplinary Science Reviews_ 40, no. 4 (2015): 376--392, <https://doi.org/10.1080/03080188.2016.1165454>.
- Benjamin Schmidt, "[Do Digital Humanists Need to Understand Algorithms?](http://dhdebates.gc.cuny.edu/debates/text/99)" in _Debates in the Digital Humanities 2016_, ed. Matthew K. Gold and Lauren F. Klein (University of Minnesota Press, 2016).

### Week 14 (May 5): Project workshop

Assignment:

- Circulate a draft of your final project in Slack by Friday, April 28. Be prepared to present your work in class for approximately ten minutes. Read each person's draft and come prepared to offer helpful comments on their work.

## Fine print

This syllabus may be updated online as necessary. The online version of this syllabus is the only authoritative version.

Students must satisfactorily complete all assignments in order to pass this course. I am sometimes willing to grant extensions on assignments for cause, but you must request an extension before the assignment's due date. For graduate students, I never penalize late work, but falling behind is not a good idea. No work (other than final projects) will be accepted after the last day that the class meets. I will discuss grades only in person during office hours.

You are welcome to use generative AI for any part of this class; in fact, I hope to model for you how to do so. There are two requirements, however. The first is that you must be willing to put your name on anything you submit for the class. The second is that you must include a brief note explaining how you used generative AI alongside (but not as a part of) your assignment. Submit that brief note in Canvas.

Canvas will calculate all grades and I will not round them or perform separate calculations.

You are expected to attend each class and to participate actively (exceptions made only for health reasons, religious holidays, and other university-approved excuses). Whether or not students attend class consistently is the best indicator of how well they will do in the class. If you wish to be excused for an absence, please email me before the absence if possible, or as soon as possible after the absence. I understand that life happens, and I will do my best to work with you.

See the [George Mason University catalog](http://catalog.gmu.edu/) for general policies, as well as the university [statement on diversity](http://ctfe.gmu.edu/professional-development/mason-diversity-statement/). You are expected to know and follow George Mason's policies on [academic integrity](http://oai.gmu.edu/) and the [honor code](http://oai.gmu.edu/understanding-the-honor-code/). If you are a student with a disability and you need academic accommodations, please see me and contact the Office of Disability Services through [their website](http://ods.gmu.edu). You are responsible for verifying your enrollment status. All academic accommodations must be arranged through that office. Please note the dates for dropping and adding courses from the [GMU academic calendar](https://registrar.gmu.edu/calendars/spring_2025/).

This syllabus draws ideas and assignments from many people and syllabi, including Taylor Arnold, Andrew Goldstone, Kieran Healy, Jason Heppler, Ben Schmidt, and Lauren Tilton.
