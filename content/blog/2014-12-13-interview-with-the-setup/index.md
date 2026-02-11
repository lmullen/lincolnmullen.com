---
title: Interview with the Setup
date: "2014-12-13"
slug: 'interview-with-the-setup'
tags:
- personal
draft: false
---

*\[I was interviewed by [The Setup](http://lincoln.mullen.usesthis.com/), "a collection of nerdy interviews asking people from all walks of life what they use to get the job done." The interview is [there](http://lincoln.mullen.usesthis.com/) and below.\]*

#### Who are you, and what do you do?

My name is [Lincoln Mullen](http://lincolnmullen.com) and I'm a historian at [George Mason University](http://historyarthistory.gmu.edu/). I teach and write about American religious history; at the moment I'm writing a history of people who converted between religions in the nineteenth-century United States. I'm also a digital historian at the [Roy Rosenzweig Center for History and New Media](http://chnm.gmu.edu/). I'm responsible for teaching a [graduate course on computer programming for historians](https://files.lincolnmullen.com/syllabi/clio3.2014.pdf), for which I'm writing a book (still very much a rough draft) on [using R for digital history](http://dh-r.lincolnmullen.com/).

#### What hardware do you use?

I have a [15" MacBook Pro](http://www.apple.com/macbook-pro/) with an Intel i7 processor, an SSD, and 16 GB of RAM provided by my department. I also use a [ThinkPad T430](http://shop.lenovo.com/us/en/laptops/thinkpad/t-series/t430/) which was my main machine during graduate school. There is an unremarkable Dell external monitor on my desk. I have an iPhone 5C but I more and more dislike using a phone for anything.

I really like my [Timbuk2 Command Laptop messenger](http://www.amazon.com/gp/product/B005HEV52G/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=B005HEV52G&linkCode=as2&tag=thebacgla-20&linkId=6PG2AG73XNIA5ES4); I wouldn't change a thing about it.

#### And what software?

I used to play around with software much too often. Now I've settled down with a few organizing principles influenced by Mike Gancarz's *[Linux and the Unix Philosophy](http://www.amazon.com/gp/product/1555582737/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=1555582737&linkCode=as2&tag=thebacgla-20&linkId=ANASKTZGVC7WT6NX)*: I use Unix-style tools and store my data in plain text or flat files. Everything that I write is formatted in [Markdown](http://daringfireball.net/projects/markdown/syntax) using the [Pandoc](http://johnmacfarlane.net/pandoc/) extensions. [John MacFarlane's](http://johnmacfarlane.net/) [Pandoc](http://johnmacfarlane.net/pandoc/) is peerless for converting between markup formats. I use some custom [LaTeX](http://www.latex-project.org/) templates with Pandoc. I use [Vim](http://www.vim.org/), often in the terminal but usually in [MacVim](https://code.google.com/p/macvim/), for all text editing. Every project---writing, coding, you name it---is kept under version control with [Git](http://git-scm.com/) and almost always made available on [GitHub](https://github.com/lmullen). Almost every project is built by [GNU Make](http://bost.ocks.org/mike/make/). All of that is run through the shell, usually [ZSH](http://www.zsh.org/). My [dotfiles](https://github.com/lmullen/dotfiles) and [Vim files](https://github.com/lmullen/vimrc) are available on GitHub.

I used to write in several different programming languages, but now I use [R](http://www.r-project.org/) if at all possible. It took me some getting used to R and functional programming, but R really is a beautiful language for data analysis. It is made even better by the "Hadleyverse" of packages written by [Hadley Wickham](http://had.co.nz/), such as [dplyr](http://cran.r-project.org/web/packages/dplyr/index.html), [tidyr](http://cran.r-project.org/web/packages/tidyr/index.html), and [ggplot2](http://ggplot2.org/). I use [Yihui Xie's](http://yihui.name/) [knitr](http://yihui.name/knitr/) package to write and code together in [R Markdown](http://rmarkdown.rstudio.com/). [RStudio](http://www.rstudio.com/products/RStudio/) (in Vim mode) is my only exception to the rule about using Vim for everything; it is a very nicely done IDE for R. Even better, I can use the server version of RStudio on a much more powerful machine, or for teaching a workshop, and get the same interface. If I can't write it in R, then I use [Ruby](https://www.ruby-lang.org/en/), which is also beautiful. JavaScript is not beautiful, but [Mike Bostock's](http://bost.ocks.org/mike/) [D3.js](http://d3js.org/) is; I use it for interactive data visualizations and [maps](https://projects.lincolnmullen.com/slavery/).

For managing citations I use [Zotero](https://www.zotero.org/), an excellent open-source application and web service by my colleagues at RRCHNM. My notes are written in Markdown, stored in [Git](http://git-scm.com/), and edited in Vim, but they are turned into a wiki by [Gitit](http://gitit.net/).

For most websites, like my personal site or my course syllabi, I use [Jekyll](http://jekyllrb.com/). Jekyll is a static site generator written in Ruby so I'm comfortable writing plugins for it; it uses Markdown (with a Pandoc plugin) so I can write in plain text and keep the site under version control. For any kind of web exhibit or other scholarly website I use [Omeka](http://omeka.org/) (think WordPress for galleries, libraries, archives, museums, and scholars). Omeka is also created by my colleagues at RRCHNM. If I assign students to write blog posts, then I use [WordPress](https://wordpress.org/). I've been really pleased with [Reclaim Hosting](http://reclaimhosting.com/), which provides hosting to educators and students.

The ThinkPad runs [Ubuntu 14.04 LTS](http://www.ubuntu.com/). Whenever I spool up a cloud or local virtual machine, it runs Ubuntu. I often use [Vagrant](https://www.vagrantup.com/) with [VirtualBox](https://www.virtualbox.org/) to create development environments. [Homebrew](http://brew.sh/) is obligatory for installing development dependencies on a Mac.

I use a few Mac or web apps: [iTerm 2](http://iterm2.com/) as a terminal emulator; [OmniFocus](https://www.omnigroup.com/omnifocus) for task management, which I could probably use better; [BackBlaze](https://www.backblaze.com/) for online backups; [DropBox](https://db.tt/f5srxqL) for file syncing; [Caffeine](http://lightheadsw.com/caffeine/) to keep my monitor from going to sleep while I'm teaching; [QGIS](http://www.qgis.org/en/site/) for GIS work if I really must; [Transmit](http://panic.com/transmit/) for FTP and Amazon S3; [TextExpander](http://smilesoftware.com/TextExpander/index.html) for snippets; [Spectacle](http://spectacleapp.com/) for window management; [Feedly](http://feedly.com/) for reading feeds.

My open-source or open-access licenses of choice are [MIT](http://opensource.org/licenses/MIT) for software and [CC-BY](http://creativecommons.org/licenses/by/4.0/) for prose.

#### What would be your dream setup?

I have plenty of computing power and can rent more cheaply. I don't even have much to complain about when it comes to battery power. I wish that there was a way to teach students digital methods without going through the grime of setting up a development environment.

I've never found an e-reader that I liked. I want a device with a large-enough screen approaching an 8.5" by 11" piece of paper that could read e-books and PDFs of journal articles, nineteenth-century books, and students papers with equal ease, preferably with an e-ink screen. Is that too much to ask?

And most of all, I'd like more and more academic work in history and the humanities to be released online (in pre-prints like [arXiv](http://arxiv.org/), or in final versions) in open formats and hopefully under open-access licenses.
