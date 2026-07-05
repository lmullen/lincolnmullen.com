---
title: "Interview at *Uses This*"
description: "An interview for *Uses This* on the tools a historian-programmer relies on: MacBook Pro, R, Vim, and a general preference for things that last over things that are new."
date: '2014-12-11'
slug: 'uses-this'
draft: false
crosspost:
  url: "https://usesthis.com/interviews/lincoln.mullen/"
  source: "*Uses This*"
tags:
- tools
---

## Who are you, and what do you do?

My name is [Lincoln Mullen](http://lincolnmullen.com/) and I'm a historian at [George Mason University](http://historyarthistory.gmu.edu/). I teach and write about American religious history; at the moment I'm writing a history of people who converted between religions in the nineteenth-century United States. I'm also a digital historian at the [Roy Rosenzweig Center for History and New Media](http://chnm.gmu.edu/). I'm responsible for teaching a [graduate course on computer programming for historians](https://files.lincolnmullen.com/syllabi/clio3.2014.pdf), for which I'm writing a book (still very much a rough draft) on [using R for digital history](http://dh-r.lincolnmullen.com/).

## What hardware do you use?

I have a [15" MacBook Pro](https://www.apple.com/macbook-pro/) with an Intel i7 processor, an SSD, and 16 GB of RAM provided by my department. I also use a [ThinkPad T430](http://shop.lenovo.com/us/en/laptops/thinkpad/t-series/t430/) which was my main machine during graduate school. There is an unremarkable Dell external monitor on my desk. I have an [iPhone 5C](https://en.wikipedia.org/wiki/IPhone_5C) but I more and more dislike using a phone for anything.

I really like my [Timbuk2 Command Laptop messenger](https://www.amazon.com/Timbuk2-Command-Laptop-TSA-Friendly-Messenger/dp/B0068DTASQ); I wouldn't change a thing about it.

## And what software?

I used to play around with software much too often. Now I've settled down with a few organizing principles influenced by Mike Gancarz's [*Linux and the Unix Philosophy*](http://www.amazon.com/gp/product/1555582737/): I use Unix-style tools and store my data in plain text or flat files. Everything that I write is formatted in [Markdown](https://daringfireball.net/projects/markdown/) using the [Pandoc](https://pandoc.org/) extensions. [John MacFarlane's](http://john.macfarlane.usesthis.com/) Pandoc is peerless for converting between markup formats. I use some custom [LaTeX](https://www.latex-project.org/) templates with Pandoc. I use [Vim](https://www.vim.org/), often in the terminal but usually in [MacVim](https://github.com/macvim-dev/macvim), for all text editing. Every project---writing, coding, you name it---is kept under version control with [Git](https://git-scm.com/) and almost always made available on [GitHub](https://github.com/). Almost every project is built by [GNU Make](http://www.gnu.org/software/make/manual/make.html). All of that is run through the shell, usually [ZSH](http://www.zsh.org/). My [dotfiles](https://github.com/lmullen/dotfiles/) and [Vim files](https://github.com/lmullen/vimrc/) are available on GitHub.

I used to write in several different programming languages, but now I use [R](http://www.r-project.org/) if at all possible. It took me some getting used to R and functional programming, but R really is a beautiful language for data analysis. It is made even better by the "Hadleyverse" of packages written by [Hadley Wickham](http://hadley.wickham.usesthis.com), such as [dplyr](https://github.com/hadley/dplyr/), [tidyr](https://github.com/hadley/tidyr/), and [ggplot2](http://ggplot2.org/). I use [Yihui Xie's](http://yihui.xie.usesthis.com/) [knitr](https://github.com/yihui/knitr) package to write and code together in [R Markdown](https://support.rstudio.com/hc/en-us/articles/205368677-R-Markdown-Dynamic-Documents-for-R). [RStudio](https://www.rstudio.com/) (in Vim mode) is my only exception to the rule about using Vim for everything; it is a very nicely done IDE for R. Even better, I can use the server version of RStudio on a much more powerful machine, or for teaching a workshop, and get the same interface. If I can't write it in R, then I use [Ruby](https://www.ruby-lang.org/en/), which is also beautiful. [JavaScript](https://en.wikipedia.org/wiki/JavaScript) is not beautiful, but Mike Bostock's [D3.js](https://d3js.org/) is; I use it for interactive data visualizations and [maps](https://projects.lincolnmullen.com/slavery/).

For managing citations I use [Zotero](https://www.zotero.org/), an excellent open-source application and web service by my colleagues at RRCHNM. My notes are written in Markdown, stored in Git, and edited in Vim, but they are turned into a wiki by [Gitit](http://gitit.net/).

For most websites, like my personal site or my course syllabi, I use [Jekyll](https://jekyllrb.com/). Jekyll is a static site generator written in Ruby so I'm comfortable writing plugins for it; it uses Markdown (with a Pandoc plugin) so I can write in plain text and keep the site under version control. For any kind of web exhibit or other scholarly website I use Omeka (think WordPress for galleries, libraries, archives, museums, and scholars). [Omeka](https://omeka.org/) is also created by my colleagues at RRCHNM. If I assign students to write blog posts, then I use [WordPress](https://wordpress.com/). I've been really pleased with [Reclaim Hosting](https://reclaimhosting.com/), which provides hosting to educators and students.

The ThinkPad runs [Ubuntu 14.04 LTS](https://www.ubuntu.com/). Whenever I spool up a cloud or local virtual machine, it runs Ubuntu. I often use [Vagrant](https://www.vagrantup.com/) with [VirtualBox](https://www.virtualbox.org/) to create development environments. [Homebrew](http://brew.sh) is obligatory for installing development dependencies on a Mac.

I use a few Mac or web apps: [iTerm2](https://iterm2.com/) as a terminal emulator; [OmniFocus](https://www.omnigroup.com/omnifocus/) for task management, which I could probably use better; [BackBlaze](https://www.backblaze.com/cloud-backup.html) for online backups; [DropBox](https://www.dropbox.com/) for file syncing; [Caffeine](http://lightheadsw.com/caffeine/) to keep my monitor from going to sleep while I'm teaching; [QGIS](https://qgis.org/en/site/) for GIS work if I really must; [Transmit](https://panic.com/transmit/) for FTP and Amazon S3; [TextExpander](https://smilesoftware.com/textexpander) for snippets; [Spectacle](https://www.spectacleapp.com/) for window management; [Feedly](https://feedly.com/) for reading feeds.

My open-source or open-access licenses of choice are MIT for software and CC-BY for prose.

## What would be your dream setup?

I have plenty of computing power and can rent more cheaply. I don't even have much to complain about when it comes to battery power. I wish that there was a way to teach students digital methods without going through the grime of setting up a development environment.

I've never found an e-reader that I liked. I want a device with a large-enough screen approaching an 8.5" by 11" piece of paper that could read e-books and PDFs of journal articles, nineteenth-century books, and students papers with equal ease, preferably with an e-ink screen. Is that too much to ask?

And most of all, I'd like more and more academic work in history and the humanities to be released online (in pre-prints like [arXiv](http://arxiv.org/), or in final versions) in open formats and hopefully under open-access licenses.
