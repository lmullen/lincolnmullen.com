---
title: "How to fork a syllabus on GitHub"
date: "2012-04-12"
slug: "how-to-fork-a-syllabus-on-github"
description: "A step-by-step guide to putting the idea of forking a syllabus into practice using GitHub's version control and sharing."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/how-to-fork-a-syllabus-on-github/39447"
  source: "ProfHacker"
tags:
- teaching
- tools
---

{{< image src="2330044065_03682e251a_n.jpg" alt="Fork in the road" >}}

A few weeks ago [Brian](http://www.briancroxall.net/) wrote a great post on [“Forking Your Syllabus.”](http://chronicle.com/blogs/profhacker/forking-your-syllabus/39137) Borrowing from discussions with [Kathy Harris](http://triproftri.wordpress.com/2012/03/08/acknowledgments-on-syllabi/) and [Trevor Owen](https://twitter.com/#!/tjowens/status/177893949992738816), Brian [advanced the idea](http://chronicle.com/blogs/profhacker/forking-your-syllabus/39137) that

<!--more-->

> syllabi could learn a trick or two from GitHub. GitHub is a repository for open source code that supports version control . . . . What this means in plain terms is that developers can share code using GitHub and then other developers can add on to that code, with the repository tracking all the changes. If a developer wants to take a piece of code down a different line of development, he or she “forks” the code. The fork shows the provenance of the code while still allowing you to adapt it to your own needs. Finding a platform to “fork your syllabus” would not only allow you to give acknowledgments to those whose work you drew on, but it would invite others to make use of your syllabus for their own development.

As Brian noted in his post, the idea of forking a syllabus is related to many different ideas we’ve supported at ProfHacker: [collaborating](http://chronicle.com/search/?contextId=5&searchQueryString=collaboration&facetName%5B0%5D=content&facetName%5B1%5D=blog&facetValue%5B0%5D=blogPost&facetValue%5B1%5D=27&facetCaption%5B0%5D=BlogPost&facetCaption%5B1%5D=ProfHacker&omni_mfs=true), [teaching by borrowing (or theft)](http://chronicle.com/blogs/profhacker/so-now-youre-a-teacher/22687), [using open-access licenses](http://chronicle.com/blogs/profhacker/using-creative-commons-licensed-material-in-your-classroom/23871), and even [using version control](http://chronicle.com/blogs/profhacker/a-gentle-introduction-to-version-control/23064).

In that spirit, in this post I want to explain how to go about forking a document, such as a syllabus, on [GitHub](http://github.com). I also want to suggest some best practices that will make it easier to fork shared documents.

## How to Fork a GitHub Repository

Using [GitHub](http://github.com) is pretty straightforward. The idea of version control itself is a little complicated, however, but Julie’s [gentle introduction](http://chronicle.com/blogs/profhacker/a-gentle-introduction-to-version-control/23064) can ease you into it. While ProfHacker is no stranger to the[command line](http://chronicle.com/blogs/profhacker/tag/the-profhacker-guide-to-the-command-line), in this case you can make your life easier by using the editing tools on GitHub’s website, or by using a graphical Git client like [GitHub for Mac](http://chronicle.com/blogs/profhacker/github-for-mac/37502). Of course, you’ll have to create a user name at GitHub, which [their documentation](http://help.github.com/) can help with.

Once you have a GitHub account, you can find a repository (a collection of documents or code) that you want to fork. For example, Karl Stolley has [shared his syllabus](https://github.com/karlstolley/ht) for a course title “Humanizing Technology.”

You can fork that repository with the click of a button. Near the top of the repository’s page you’ll find a button that says “Fork.” Click that, and GitHub will give you a copy of the repository that you can edit on your own. Now you’ve “forked a class” as [Brian suggests](http://chronicle.com/blogs/profhacker/forking-your-syllabus/39137).

{{< image src="screen-shot-2012-04-11-at-1.58.06-pm.png" alt="GitHub fork" >}}

If you want to practice, you can [fork the ProfHacker e-book repository](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo) that I set up for an earlier ProfHacker post.

## Best Practices for Sharing Documents on GitHub

If you’re going to share your documents on GitHub, there are a few things you can do to make them easier for others to use.

1. **Share your documents in plain text.** GitHub works by tracking how files change line by line, and it can only do this for files it can see inside. MS Word documents, PDFs, and other files like that are opaque to GitHub. Good formats to write in are HTML and [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295).
2. **Break your documents into smaller files.** For a syllabus, you might put your course description, policies, schedule, and each assignment all into their own files. Making each part of the syllabus modular make it easier for someone to borrow an assignment or a policy that she likes.
3. **Give it a license.** I like the [Creative Commons BY](http://creativecommons.org/licenses/by/3.0/) license. (Here is [a post by Bethany Nowviskie](http://nowviskie.org/2011/why-oh-why-cc-by/) explaining why scholars should use that license.) Whatever license you pick, though, it is important that you be explicit about what people are allowed to do with your work and what conditions they have to meet. Otherwise, people have to assume that they can’t do anything at all with your work—and that’s probably not what you want.

**Do you have any plans to share a syllabus or other documents on GitHub? Have you forked other peoples’ documents?**

*[Image](http://www.flickr.com/photos/y_i/2330044065/) courtesy of Flickr user [i\_yudai](http://www.flickr.com/photos/y_i/) // Creative Commons BY licensed*
