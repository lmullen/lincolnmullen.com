---
title: "This semester's syllabus for Data Analysis for History"
date: "2025-01-25T16:10:21-05:00"
slug: "this-semesters-syllabus-for-data-analysis-for-history"
draft: false
# newsletter: true
# tags:
# - 
# crosspost: 
#   url: ""
#   source: ""
# linkpost: ""
---

I've posted my [syllabus](/courses/data.2025/) for this semester's iteration of Data Analysis for History. I've been teaching this course for over a decade. (Previous iterations: [2014](https://files.lincolnmullen.com/syllabi/clio3.2014.pdf), [2016](https://files.lincolnmullen.com/syllabi/data-dh.2016.pdf), [2017](https://files.lincolnmullen.com/syllabi/clio2.2017.pdf), [2018](/courses/clio2.2018/), [2019](/courses/clio2.2019/), [2020](/courses/data.2020/), [2023](/courses/data.2023/).) This is the first time I have totally re-thought the syllabus, rather than iteratively improved it.

The course has always had two principles in its approach. The first is that the goal of teaching the course should be to enable students to make interpretative arguments about the past using data. That fundamental goal is unchanged. The second was that I would teach all of the methods (visualization, data manipulation, text analysis, etc.) using R. The reason was that students could learn a single language and set of libraries, as well as a core set of principles around the idea of "tidy data." They could then apply those principles in a language they knew to a wide variety of data analysis problems. I still think that is a justifiable approach, and I let it go only with some regret. There are excellent texts for teaching with that kind of approach, and my own working expertise is still mostly in that area.

But the world of data analysis is so very much different. In particular, it is so much easier to create high quality, interactive visualizations for the web using JavaScript libraries, especially [Observable Plot](https://observablehq.com/plot/). They are better than anything I know how to make in R, and much more accessible for being online. A major weakness of my previous approach to teaching was that I never really showed students how to use relational databases, which I intend to rectify with this course. My own approach to computational history has changed. Instead of trying to do everything in a single language, I tend to use [Go](/blog/a-programming-language-and-a-sense-of-self/) for high-performance computing tasks and JavaScript for visualizations, with a PostgreSQL database in the middle. I can't really see teaching students an approach that is different than the one I would use myself.  

There is one other change, of course. Generative AI is good---really good---at assisting with these kinds of data analysis. I use [Claude](https://claude.ai) to soften up all kinds of tasks in this area. For example, if you give Claude a database schema, you can ask your queries in natural language and get pretty good SQL back, along with explanations of how it works. So I think that generative AI will be a very useful teaching tool. Whereas before it was too much to ask students to tackle more than one kind of language, now I think it is possible for them to get targeted help via AI, allowing them to focus on the principles and the past, rather than on semicolons. (Or on significant whitespace: looking at you, Python.) And again, it's the same principle: I want to teach students the way I actually work, and the way I work has changed rather significantly.

We will see how the students respond to this new approach. I've had the good fortune to teach the course regularly, which allows an iterative approach from semester to semester.
