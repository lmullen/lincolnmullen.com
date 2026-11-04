---
title: "Two kinds of computational historical thinking"
date: "2026-11-04T15:03:34-05:00"
slug: "two-kinds-of-computational-historical-thinking"
draft: false
# crosspost:
#   url: ""
#   source: ""
# linkpost: ""
# via: ""
tags:
- digital history
- computational history
- teaching
description: "Responding to Mandy Regan on teaching computational history in the age of AI, a distinction between two kinds of computational thinking: understanding the methodological choices built into code, and imagining what historical questions we can compute on our sources."
social:
  bluesky: "Mandy Regan's post on teaching computational history with AI got me thinking about two kinds of computational thinking: methodology and imagination. https://lincolnmullen.com/blog/two-kinds-of-computational-historical-thinking/"
  mastodon: "Mandy Regan's post on teaching computational history with AI got me thinking about two kinds of computational thinking: methodology and imagination. https://lincolnmullen.com/blog/two-kinds-of-computational-historical-thinking/"
  linkedin: "Mandy Regan's post on teaching computational history with AI got me thinking about two kinds of computational thinking: methodology and imagination. https://lincolnmullen.com/blog/two-kinds-of-computational-historical-thinking/"
---

The [first semester](https://files.lincolnmullen.com/syllabi/clio3.2014.pdf) that I taught a (rather incoherent) course on computational history, the course was titled "Programming in History/New Media." The then-current state of that wing of digital history was that graduate students should learn to code. (Witness [*The Programming Historian*](https://programminghistorian.org).) Even at the time, I think I was trying to articulate that programming was merely a means to an end and that what we were really after was a way to think about history through computation. A good thing too, because programming came hard to the graduate students, but now an LLM can easily program anything a first-year graduate student can think of. So we may as well teach them to think computationally. 

<!--more-->

Mandy Regan was a student in that first semester, now teaching computational history at Clemson. In a recent post she reflects on "[Teaching Computational History in the Age of AI](https://www.amanda-regan.com/posts/teaching-computational-history-age-of-ai/)":

> The kinds of digital work my students produce has grown dramatically over the last four years, and the growth is not incremental. It is a different category of work and they can produce substantial and sophisticated digital scholarship that make real historical interventions. It significantly lowers the barrier to doing sophisticated digital work. ...
> 
> But I want to be careful because I think we risk drawing the wrong conclusion from all of this. It is crucially important that we do not give up on teaching the computational thinking skills that all of this rests on. Our arguments rest on assumptions built into the code. The number of topics you choose for your topic model. The stop words you throw out. The blocking strategy you use before you compare records. The threshold at which you decide two things are the same thing. The projection you use to display points on a map. Every one of those decisions is an interpretive claim dressed up as a technical default, and every one of them shapes what you are able to see and ask. If you do not understand how and why the code works, the AI model makes those determinations on your behalf. At best that alters your results in ways you cannot account for. At worst it is flatly wrong and arrives looking authoritative.

One difficulty students have always had with the computational history course has been confusing methods and code. This showed most obviously when students said I was teaching a course "on R." If R is the only programming language you know, then filtering a data table seems like something you do in R. Until students see that you can do that in R or SQL or Python or whatever, it is hard for them to see that filtering is an action that happens to be implemented in many languages. Likewise it is hard for students to see that the argument `k = 30` to a function that runs a topic model is fundamentally a methodological decision, where another argument might simply be an implementation detail. A big challenge in the course has been teaching students to think at the right level of abstraction.

Mandy and I are using the term *computational thinking* in two different but compatible ways, and her post has sharpened my thinking about how to distinguish between them. In the paragraphs quoted above, her point is that working directly in code obligates one to make many technical decisions that affect the interpretation. That usage of the term points to methodology: how does the computation work (at some level of abstraction) and what choices do you need to make about it? Having to write computer programs that do some kind of computation is one way to come to understand methodology. This is akin to Ben Schmidt's idea of computation as transformation.[^1]

The way that I often use the term *computational thinking* is more akin to *imagination*. Given a set of sources, what questions can I ask of them? What can I see about the past in them? How do I compute on them to make a historical interpretation? This is the fundamental problem: how do I match a set of sources to a set of computational methods and to a form of presentation which can tell me something meaningful about the past?

I take Mandy's point that much is lost when students use AI to jump up several rungs on the ladder of abstraction, and never touch the lower levels of computational thinking. I worry about losing that when I teach the course next semester, [redesigned yet again](/blog/this-semesters-syllabus-for-data-analysis-for-history/) for AI. But my current thought is that programming was already too low a level of abstraction, and the sheer volume of choices to make when writing code---not to mention the difficulty of getting the code to run at all---obscured what kinds of thinking were genuine methodological decisions and which were implementation details. And I'm optimistic that the staggering speedup from idea to implementation enabled for students by AI will make it possible for them to think more imaginatively.

[^1]: Benjamin M. Schmidt, “[Do Digital Humanists Need to Understand Algorithms?](https://dhdebates.gc.cuny.edu/read/untitled/section/557c453b-4abb-48ce-8c38-a77e24d3f0bd),” in *Debates in the Digital Humanities 2016*, ed. Matthew K. Gold and Lauren F. Klein (University of Minnesota Press, 2016).
