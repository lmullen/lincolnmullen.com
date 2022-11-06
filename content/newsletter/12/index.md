---
title: "Books for Labor Day"
date: "2021-09-07"
slug: 12
draft: false
tags:
- newsletter
- Roy Rosenzweig
- tools of the trade
---

As promised, here is the first installment of an occasional series on my tech stack. If you want to jump straight to the history, see below for two book recommendations on labor history and religious history.

-------------------

[Last week](/newsletter/11/) I mentioned that I would write about the technology stack that I use to do digital history work. For this week, let me briefly introduce the concept.

Talking about "a stack" of technologies is not the same thing as talking about digital tools. The whole discussion of "tools" in digital humanities, or the broader cultural contexts of life hacking, productivity pr0n, and their ilk, is not something I want to get into right now. 

Here is how a technology stack and "tools" are different. A tool is a specific means to a specific ends. You want to make a map, therefore you use X. You need to clean data, therefore you use Y. You need to make a network, therefore you use Z. There's nothing wrong with that, of course. But too often that approach leads to a shallow understanding of the method being used. Networks and maps are different, of course, but they can also be worked on using the same technologies. The short-term convenience of using a highly specialized tool buys a lot of long-term pain because then data and outputs are rarely reusable or interoperable.

So what I want to talk about for a few newsletter issues are the components of an overall approach to the kind of digital history I do (computational, spatial, data visualization and so forth), and how they work together. These are parts of a system for doing research that work together across projects. The goal is to build fluency with using these parts of the system, so that future projects can build on the data and systems I have developed along the way.

For example, almost all of my projects store their data in PostgreSQL, whether or not the data is spatial or textual, whether it comes from an API or bulk download. If I need to analyze that data, I do it in R. If it needs to go on the web, I use some combination of Go and JavaScript. Those different parts of the stack are all pulling from the same database.

Here are a few of the pieces I'm sure I want to talk about. Perhaps there will be others too.

- Data store: PostgreSQL, PostGIS
- Web and general-purpose programming: Go
- Data analysis and visualization: R and JavaScript
- Websites: Hugo
- Web servers and hosting

More in the coming weeks.

-------------------

It was just Labor Day in the U.S., so here are two book recommendations.

The first is Roy Rosenzweig's _[Eight Hours for What We Will: Workers and Leisure in an Industrial City, 1870–1920](http://services.cambridge.org/us/academic/subjects/history/american-history-general-interest/eight-hours-what-we-will-workers-and-leisure-industrial-city-18701920?format=PB&isbn=9780521313971)_. I now work at the Roy Rosenzweig Center for History and New Media, though I never met the man Roy Rosenzweig. (I was once assigned to hand out educational materials at the Organization of American Historians on behalf of RRCHNM. I was mostly unsuccessful at that task, but at least four people came up to me to tell me, at great length, how much Rosenzweig meant to them. It made an impression.) But I knew _Eight Hours for What We Will_ a long time before I came to RRCHNM, and not just because it still belongs on the exam list for every PhD in American history. One of my textbooks as an undergrad had the forbidding title, _Historiography: Ancient, Medieval, and Modern_. In reviewing labor history, it mentioned Rosenzweig's history of Worcester, Massachusetts. Since Worcester is the closest city to the town where I grew up, I looked it up in the university library and loved it. The book still sparkles with historical imagination and appreciation for its subjects, and it is very much worth reading.

The second is Heath Carter's _[Union Made: Working People and the Rise of Social Christianity in Chicago](https://global.oup.com/academic/product/union-made-9780190847371?lang=en&cc=us)_. You may detect that I love history books which are carefully anchored in a particular place. Carter's book is very much a history of Chicago, but it also unpacks the class divisions in specific congregations. The book begins with labor union representatives preaching Labor Day sermons in church pulpits in 1910. I think you will find that practice a contrast with practically any American church a century later.

-------------------

###  Updates

Working: Containerizing and adapting my prediction model for _America's Public Bible_ so that I can use it for [Computing Cultural Heritage in the Cloud](https://blogs.loc.gov/thesignal/2019/11/introducing-the-computing-cultural-heritage-in-the-cloud-project/).

Listening: Tennessee Ernie Ford, _Nearer the Cross_ (1958): found an LP for 99¢ at the used bookstore. 

Reading: Just finished Margaret O'Mara's _[The Code: Silicon Valley and the Remaking of America](https://www.penguinrandomhouse.com/books/534709/the-code-by-margaret-omara/)_, which I thought was both really good and a really good read.

Hiking: [Sky Meadows](https://www.dcr.virginia.gov/state-parks/sky-meadows) state park with my family.
