---
author: Lincoln Mullen
categories:
- digital history
- teaching
date: '2017-02-08T18:40:15+00:00'
slug: 'a-confirmation-of-andrew-goldstone-on-teaching-quantitative-methods'
title: 'A confirmation of Andrew Goldstone on "Teaching Quantitative Methods"'
---

At his blog, Andrew Goldstone has [posted a pre-print](https://andrewgoldstone.com/blog/ddh2018preprint/) of his essay on "Teaching Quantitative Methods: What Makes It Hard (in Literary Studies)" for the forthcoming *Debates in DH* *2018*. It's a "lessons learned" essay from one of his courses that is well worth reading if you're teaching or taking that kind of a course in a humanities discipline. This semester I'm teaching my fourth course that fits into that category (fifth, if you count DHSI), and I can co-sign nearly everything that Goldstone writes, having committed many of the same mistakes and learned some of the same lessons. (Except over time I've relaxed my \*nix-based fundamentalism and repealed my ban on Windows.) Here is a response to Goldstone's main points.

> Cultivating technical facility with computer tools---including programming languages---should receive less attention than methodologies for analyzing quantitative or aggregative evidence. Despite the widespread DH interest in the former, it has little scholarly use without the latter.

The [first time](http://lincolnmullen.com/courses/clio3.2014/) I taught a course like this, emphasizing programming over methods was my biggest mistake. Teaching programming concepts in JavaScript and then asking students to switch to R was the other big mistake, though some students redeemed my error by doing their final projects in D3.js. If you make a distinction between "programming" as the act of solving generalized, abstract problems and "scripting" as the act of solving a concrete, specific research problem in a computer language, students in these courses aren't learning programming anyway, though of course they are learning programming concepts. In later versions of the course ([2016](http://lincolnmullen.com/courses/data-dh.2016/), [2017](http://lincolnmullen.com/courses/clio2.2017/)), I've boiled the programming prolegomenon down to a week and a half.

In its place I've put one week on how the historical sources of data that we are using were created and another week on how historians have used data analysis. That's in addition to methodological readings which also deal with questions of historical thinking with data. The course as taught is increasingly organized around how to frame and answer historical questions. Ideally, the course would let a student who wants to use one of these methods for a dissertation get a head start. But the course is still inadequate to those purposes. One problem is that the field of digital history generally doesn't have a good approach to producing disciplinary argumentation. Another problem is that students in history might want to use a broad range of computational techniques, so it seems to me that the course should survey a range of approaches rather than go deep into one of them.

> Studying method requires pedagogically suitable material for study, but good teaching datasets do not exist. It will require communal effort to create them on the basis of existing research.

My approach to this problem has been three-pronged. First, my [historydata](https://github.com/ropensci/historydata) R package has a number of cleaned up, more-or-less contrived datasets which can be used to demonstrate specific techniques. Second, when those are inadequate, I have a few datasets or corpora available to me from my own research, such as election returns from *Mapping Early American Elections* */ A New Nation Votes*, a corpus of codes of civil procedure and another of legal treatises, the WPA former slave narratives corpus, and so on. And third, I always let students substitute their own data from their own historical period and topic whenever feasible, but in practice it rarely is feasible given how difficult it is to create a useable dataset. So the danger here is that I'm just imposing my eclectic mix of nineteenth-century religious, political, and legal history on the students and rehashing my own approach. (And, I suppose, there is the corollary danger of being one of those people who can only talk about their own work.) So this is pretty far from the "communal effort" that Goldstone wants, though I could probably make better use of the datasets that people like [Miriam Posner](http://miriamposner.com/dh101f15/index.php/assignments/final-project/datasets/) or the *[Historian's Macroscope](http://www.themacroscope.org/2.0/data-downloads/)* authors have put together.

> Following the "theory" model, DH has typically been inserted into curricula as a single-semester course. Yet as a training in method, the analysis of aggregate data will undoubtedly require more time, and a different rationale, than that offered by what Gerald Graff calls "the field-coverage principle" in the curriculum.

I'm fortunate that I'm not obligated to stick to a single course thanks to the way that the Mason PhD in history is arranged. The field is plagued by an expectation that digital history is something which can be learned while standing on one foot, but in the graduate program we have the opportunity to do more than that for students. All the PhD students and many of the master's students have to take a course vernacularly known as "Clio 1," which amounts to an introduction to the theory and practice of digital history. I haven't taught this course, but I benefit from the fact that most people in my classes have already learned from my colleagues. All the PhD students and some master's students also have to take "Clio 2," which could be about several things. But since in the short term I'm likely to be the only person teaching it every year, it amounts to "Data and Visualization in Digital History." Then students who elect to do a minor field in digital history (a minority of grad students, but a sizable one) can choose from a number of courses taught by different faculty. One of those is my "Clio 3," which I originally taught as a programming course, but now construing the catalog's description rather broadly I'm going to treat it as a course that goes deep into one method like [text analysis](http://lincolnmullen.com/courses/text-analysis.2016/) or that covers some other advanced topic useful to a specific group of students.

To sum up, I have some of the same misgivings about my course that Goldstone does. What's changed most about my approach, besides the re-framing it around methods rather than programming, is that I'm thinking about how they fit into a broader curriculum, including an ongoing relationship with students. Of course not every student who is required to take my course in data and visualization will want to take additional steps, and that's fine. But for those who do, my intention is to support them with a second course that is a deeper dive, to keep up on-going support for their specific research topics, and to help people who elect to do one of the [digital or hybrid dissertations](http://historyarthistory.gmu.edu/graduate/phd-history/digital-dissertation-guidelines) permitted by my department.

Goldstone, Andrew. "Teaching Quantitative Methods: What Makes It Hard (in Literary Studies)." In Matthew K. Gold, Lauren F. Klein, eds., <em>Debates in the Digital Humanities 2018.</em> Minneapolis: University of Minnesota Press, forthcoming 2018. Pre-print DOI: <a href="http://dx.doi.org/doi:10.7282/T3G44SKG">10.7282/T3G44SKG</a>.