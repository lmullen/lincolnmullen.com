---
title: "Teaching computational history"
date: "2020-01-22"
slug: 2
draft: false
tags:
- newsletter
- computational history
---

Greetings. And welcome back to "Working on It."

---------------------

This semester I will be teaching my class on "[Computational History](/courses/data.2020/)" for the sixth time. The official title in the catalog is something else, but I've gone rogue since my second year at GMU, and even the unofficial title has changed over time.

This is the course I was hired by GMU to teach. Or, more precisely, my department vaguely had the idea that they wanted a course on programming for historians, and I vaguely had an idea of how it could be taught. I shudder a bit looking back at the [first iteration](/courses/clio3.2014/) of the course. The bones of the course are similar, but my approach to the subject, including the wildly unrealistic idea that I should teach JavaScript as an introduction to R, now seems fundamentally mistaken. I was fortunate to have a great group of students who made the most of that first course as I learned the material a week ahead of them.

Now that I understand---or hope I understand---how to teach the course, perhaps I can explain some of the reasons the syllabus has changed over time.

The most important consideration is that the course has been required for Ph.D. students in history at George Mason. The chance to offer the course every year has let me improve it regularly. But it has also meant that the course had to be designed to meet the department's goals, which I have done by trying to give students a clear sense of how they can progress from this course both to work on their dissertations and to work at RRCHNM. Former students have taken both of those routes. The fact that the course is required has also meant that at least some students feel like it is not within their interests, and so I've worked hard with individual students to figure out connections to what they care about.

Besides the course's local context, it has also been influenced by developments in the historical profession. The place of computational history is, I think, needlessly polarized. After all, even one of the leading practitioners of computational history has argued that "[computational history is dead for good](https://benschmidt.org/post/2019-12-05-totc/2019-aha/)." At issue is whether computational history has promised a revolution in historical understanding, possibly a revolution in which all historians must participate willingly or unwillingly, and whether its interpretative results have been underwhelming.

The way that I present the course to students is hardly radical. My approach to historical methods is to "multiply your ways of knowing"---a [principle](https://dh-r.lincolnmullen.com/introduction.html) I learned from an advisor in grad school. While historians revise and discard interpretations, as a profession they tend to accumulate methods, adding the new but retaining the old. The usefulness of computational history is not in provoking a revolution of all historical understanding, but in creating a new angle of vision on sources that we can also read with existing historical methods. This also happens to be the way, I have found, that students approach computational history: a tiny handful want to upend the discipline, but a great many find it a useful set of techniques to weave into their other forms of historical research.

Another concern is the distinction between programming and methodology. In an article about a similar course in literary studies, Andrew Goldstone [argues](https://dhdebates.gc.cuny.edu/read/untitled-f2acf72c-a469-49d8-be35-67f9ac1e3a60/section/620caf9f-08a8-485e-a496-51400296ebcd#ch19) that "cultivating technical facility with computer tools---including programming languages---should receive less attention than methodologies for analyzing quantitative or aggregative evidence. Despite the widespread DH interest in the former, it has little scholarly use without the latter." Certainly, my course (originally titled "Programming in History/New Media") came into the curriculum at the same moment some digital humanists were claiming that everyone needed to learn how to code and that resources like [*The Programming Historian*](https://its.gmu.edu/service/wordpress-course-blogs/) were being created. That's a moment that the field has moved past for some good reasons: the idea that all humanists would learn to program was not reasonable, feasible, or desirable. It is also a moment we have moved past for dubious reasons, including the proliferation of "tools" for humanities data analysis, most of which are of poor quality and have little methodological sophistication, and which enable analysis without understanding.

Over time my syllabus has shifted from emphasizing programming towards emphasizing methodology. The recasting of the course as computational history rather than programming for historians reflects that change. (I have also taught a course on [web programming](https://lincolnmullen.com/courses/programming.2019/), which is a topic that is highly specific to the culture of the Mason/RRCHNM Ph.D. program.) In the data analysis course I tend to teach programming concepts---say, iteration---in only a cursory way because they are not critical to understanding and using the methods being studied. And methods, rather than programming, is what produces an evidentiary base for historical interpretations.

That said, it is still important that this course is the first place where most of the students encounter some kind of programming. I emphasize to them that, although they are learning R, the same concepts and methods can be expressed in other languages such as Python and SQL. Students with whom I continue working after the class is over not infrequently have a moment when they realize that some difficult problem in a different domain is conceptually the same as what they have already learned. But that kind of abstraction takes time. In a sense, if a programming language is already an abstraction, understanding the shared concepts is an abstraction of an abstraction, and beginners find it extremely difficult to grasp both. In the meantime, learning a specific programming language gives them, well, a language to conceptualize and reason about historical data and the inferences that can be made from it.

Many of the changes to the course have improved the sequencing of what I teach. Like nineteenth-century theories of evolution in which an embryo recapitulates the entire process of evolution, one bad pattern is teaching students is the same way that you happened to learn. That was a trap I sometimes fell into in my early semesters of teaching. Like almost everyone teaching R these days, I focus on the [tidyverse](https://www.tidyverse.org) as a set of interoperable packages and principles. The course is structured around understanding quantitative data analysis and visualization in generic terms, then applying it to the peculiarities of spatial, network, and textual data. Most of the changes I've made are not obvious on the syllabus itself, but by revising the course regularly I am much more confident that the sequence of materials comes to the students in a useful way, even if they still feel it comes at them much too quickly. 

Finally, there is the consideration of what kind of work the students should produce. In the earliest iterations, I was content to let students created exploratory graphics. I thought of these methods primarily as tools of exploration and, frankly, hadn't made the leap from exploration to argumentation myself. But now I try hard to help students move from a dataset through methods to a historical argument. Tutorials or textbooks are generally of little help in this regard, as they use contrived examples where, as Goldstone puts it, "the effort is out of all proportion to the result." His [suggestion](https://dhdebates.gc.cuny.edu/read/untitled-f2acf72c-a469-49d8-be35-67f9ac1e3a60/section/620caf9f-08a8-485e-a496-51400296ebcd#ch19), which I am increasingly adopting, is to have them follow and even to reproduce existing scholarship so that students can see how an argument is put together.

In its current version, the course is not so dissimilar from other courses in data analysis for history or literary studies. That convergence reflects what I take to be a growing consensus on how best to teach these methods. But there is one thing that distinguishes this iteration of the course from other offerings: this syllabus assigns work by several former students in the class.

-----------

### Random screenshot

{{< image src="random-screenshot.png" alt="A random screenshot" >}}

-----------

### Around the web

- What's a *pettifogger*? Kellen Funk [explains](https://twitter.com/kellenfunk/status/1219975657083625475).
- Matthew Lincoln (no relation, obviously) has been blogging like it is the early 2000s about his [first year at Carnegie Mellon](https://matthewlincoln.net/tags/cmu_dh.html).

-----------

### Updates

Reading: N.T. Wright's biography of Paul; Neal Stephenson's *Seveneves*.

Watching: *The Expanse*.

Working: Thinking about the infrastructure for presenting computational history on the web.
