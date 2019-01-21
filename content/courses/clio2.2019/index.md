---
title: "Clio 2: Computational History"
semester: Spring
year: 2019
department: Department of History and Art History
university: George Mason University
coursenum: HIST 697-001
courseurl: https://historyarthistory.gmu.edu/courses/hist697/course_sections/38740
credits: 3
meets: Mondays, 7:20–10:00pm
room: RRCHNM conference room, Research Hall 402
instructor: Lincoln Mullen
instructorurl: https://lincolnmullen.com
email: lmullen@gmu.edu
office: Research Hall 457
officehours: "Mondays, 11:00–12:00pm; Wednesdays, 10:30–11:45am"
toc: true
draft: true
---

# Course description

In this course you will learn to apply computational methods to create historical arguments. You will learn to work with historical data, including finding, gathering, manipulating, analyzing, visualizing, and arguing from data, with special attention to geospatial, textual, and network data. These methods will be taught primarily through scripting in the [R programming language](https://www.r-project.org/). While historical methods can be applied to many topics and time periods, they cannot be understood separate from how the discipline forms meaningful questions and interpretations, nor divorced from the particularities of the sources and histories of some specific topic. You will therefore work through a series of example problems using datasets from the history of the nineteenth-century U.S. religion, and you will apply these methods to a dataset in your own field of research.


# Learning goals

After taking this course, you will be able to

- perform exploratory data analysis; clean, tidy, and manipulate data; gather historical data from print and manuscript sources; use existing historical data sets; create common visualizations; work with geospatial, textual, and network data.
- write scripts using the R programming language and its extensive set of packages.
- understand the place of data analysis and visualization within humanities computing, digital history, and the discipline of history.
- conceive of and execute a research project in computational history suitable for treatment in a dissertation chapter or journal article.
- take the course "Programming in History/New Media," a.k.a. Clio 3, should you choose.

# Essential information

You are always welcome to [book an appointment](https://calendar.google.com/calendar/selfsched?sstoken=UU94MzdPM1ZlU0ZqfGRlZmF1bHR8ZTg3NzRjZDZiYTMxNDU5ZjEzNjZkZTZkODAyODVmYWI) during my office hours. If the times that are available do not work for you, feel free to contact me. All communication for this course will happen in [our Slack group](https://mason-dh-grad.slack.com/signup). Read this [getting started guide](https://get.slack.help/hc/en-us/articles/218080037-Getting-started-for-new-users) if you need help.

Bring a computer to each class meeting. We will use [R](https://cran.rstudio.com/) and [RStudio](https://www.rstudio.com/products/rstudio/#Desktop). [Install them](http://dh-r.lincolnmullen.com/installing-r-and-packages.html) on your own computer. You will also have access to an [RStudio Server instance](http://rstudio.chnm.org/) which will let you use R in your browser. Much of your work for the course will go on [GitHub](https://github.com/), so sign up for an account.

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

Due Monday, May 13 at 5pm. 

Final grades will be calculated using the typical percentage-based grading scale (A = 93--100, A- = 90--92, B+ = 88--89, B = 83--87, B- = 80--82, ... F = 0--59).

# Schedule 

### Monday, January 28

### Monday, February 4

### Monday, February 11

### Monday, February 18

### Monday, February 25

### Monday, March 4

### Monday, March 11: Spring break

### Monday, March 18

### Monday, March 25

### Monday, April 1

### Monday, April 8

### Monday, April 15

### Monday, April 22

### Monday, April 29

### Monday, May 6

# Fine print

This syllabus may be updated online as necessary. The online version of this syllabus is the only authoritative version.

Students must satisfactorily complete all assignments (including participation assignments) in order to pass this course. Your attendance is expected at every meeting (exceptions made only for religious holidays). If you must be absent, I request that you notify me in advance of the class meeting. I am sometimes willing to grant extensions for cause, but you must request an extension before the assignment's due date. For every day or part of a day that an assignment is late without an extension, I may reduce your grade. No work (other than final projects) will be accepted after the last day that the class meets. I will discuss grades only in person during office hours.

See the [George Mason University catalog](http://catalog.gmu.edu/) for general policies, as well as the university [statement on diversity](http://ctfe.gmu.edu/professional-development/mason-diversity-statement/). You are expected to know and follow George Mason's policies on [academic integrity](http://oai.gmu.edu/) and the [honor code](http://oai.gmu.edu/understanding-the-honor-code/). If you are a student with a disability and you need academic accommodations, please see me and contact the Office of Disability Services at 703-993-2474 or through [their website](http://ods.gmu.edu). You are responsible for verifying your enrollment status. All academic accommodations must be arranged through that office. Please note the dates for dropping and adding courses from the [GMU academic calendar](http://registrar.gmu.edu/calendars/spring-2019/).
