---
date: '2014-10-11 13:31'
layout: post
title: Interview with the Setup
...

### 1) Who are you, and what do you do?

My name is [Lincoln Mullen][] and I'm a historian at [George Mason
University][]. I teach and write about American religious history; at
the moment I'm writing a history of people who converted between
religions in the nineteenth-century United States. I'm also a digital
historian at the [Roy Rosenzweig Center for History and New Media][].
Among the many projects created by my colleagues at RRCHNM do, I'm
responsible for teaching a [graduate course on computer programming for
historians][], for which I'm writing a book (still very much a rough
draft) on [using R for digital history][].

### 2) What hardware do you use?

I have a [15" MacBook Pro][] with an Intel i7 processor, an SSD, and 16
GB of RAM provided by my department. I also use a [ThinkPad T430][]
which was my main machine during graduate school. There is an
unremarkable Dell external monitor on my desk. I have an iPhone 5C but I
more and more dislike using a phone for anything.

I really like my [Timbuk2 Command Laptop messenger][]; I wouldn't change
a thing about it.

### 3) And what software?

I used to play around with software much too often. Now I've settled
down with a few organizing principles influenced by Mike Gancarz's
*[Linux and the Unix Philosophy][]*: I use Unix-style tools and store my
data in plain text or flat files. Everything that I write is formatted
in [Markdown][] using the [Pandoc][] extensions. [John MacFarlane's][]
[Pandoc][] is peerless for converting between markup formats. I use some
custom [LaTeX][] templates with Pandoc. I use [Vim][], often in the
terminal but usually in [MacVim][], for all text editing. Every
project---writing, coding, you name it---is kept under version control
with [Git][] and almost always made available on [GitHub][]. Almost
every project is built by [GNU Make][]. All of that is run through the
shell, usually [ZSH][]. My [dotfiles][] and [Vim files][] are available
on GitHub.

I used to write in several different programming languages, but now I
use [R][] if at all possible. It took me some getting used to R and
functional programming, but R really is a beautiful language for data
analysis. It is made even better by the "Hadleyverse" of packages
written by [Hadley Wickham][], such as [dplyr][], [tidyr][], and
[ggplot2][]. I use [Yihui Xie's][] [knitr][] package to write and code
together in [R Markdown][]. [RStudio][] (in Vim mode) is my only
exception to the rule about using Vim for everything; it is a very
nicely done IDE for R. Even better, I can use the server version of
RStudio on a much more powerful machine, or for teaching a workshop, and
get the same interface. If I can't write it in R, then I use [Ruby][],
which is also beautiful. JavaScript is not beautiful, but [Mike
Bostock's][] [D3.js][] is; I use it for interactive data visualizations
and [maps][].

For managing citations I use [Zotero][], an excellent open-source
application and web service by my colleagues at RRCHNM. My notes are
written in Markdown, stored in [Git][], and edited in Vim, but they are
turned into a wiki by [Gitit][].

For most websites, like my personal site or my course syllabi, I use
[Jekyll][]. Jekyll is a static site generator written in Ruby so I'm
comfortable writing plugins for it; it uses Markdown (with a Pandoc
plugin) so I can write in plain text and keep the site under version
control. For any kind of web exhibit or other scholarly website I use
[Omeka][] (think WordPress for galleries, libraries, archives, museums,
and scholars). Omeka is also created by my colleagues at RRCHNM. If I
assign students to write blog posts, then I use [WordPress][]. I've been
really pleased with [Reclaim Hosting][], which provides hosting to
educators and students.

The ThinkPad runs [Ubuntu 14.04 LTS][]. Whenever I spool up a cloud or
local virtual machine, it runs Ubuntu. I often use [Vagrant][] with
[VirtualBox][] to create development environments. [Homebrew][] is
obligatory for installing development dependencies on a Mac.

I use a few Mac or web apps: [iTerm 2][] as a terminal emulator;
[OmniFocus][] for task management, which I could probably use better;
[BackBlaze][] for online backups; [DropBox][] for file syncing;
[Caffeine][] to keep my monitor from going to sleep while I'm teaching;
[QGIS][] for GIS work if I really must; [Tranmit][] for FTP and Amazon
S3; [TextExpander][] for snippets; [Spectactle][] for window management;
[Feedly][] for reading feeds.

My open-source or open-access licenses of choice are [MIT][] for
software and [CC-BY][] for prose.

### 4) What would be your dream setup?

I have plenty of computing power and can rent more cheaply. I don't even
have much to complain about when it comes to battery power. I wish that
there was a way to teach students digital methods without going through
the grime of setting up a development environment.

I've never found an e-reader that I liked. I want a device with a
large-enough screen approaching an 8.5" by 11" piece of paper that could
read e-books and PDFs of journal articles, nineteenth-century books, and
students papers with equal ease, preferably with an e-ink screen. Is
that too much to ask?

And most of all, I'd like more and more academic work in history and the
humanities to be released online (in pre-prints like [arXiv][], or in
final versions) in open formats and hopefully under open-access
licenses.

  [Lincoln Mullen]: http://lincolnmullen.com
  [George Mason University]: http://historyarthistory.gmu.edu/
  [Roy Rosenzweig Center for History and New Media]: http://chnm.gmu.edu/
  [graduate course on computer programming for historians]: http://lincolnmullen.com/courses/clio3.2014/
  [using R for digital history]: http://dh-r.lincolnmullen.com/
  [15" MacBook Pro]: http://www.apple.com/macbook-pro/
  [ThinkPad T430]: http://shop.lenovo.com/us/en/laptops/thinkpad/t-series/t430/
  [Timbuk2 Command Laptop messenger]: http://www.amazon.com/gp/product/B005HEV52G/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=B005HEV52G&linkCode=as2&tag=thebacgla-20&linkId=6PG2AG73XNIA5ES4
  [Linux and the Unix Philosophy]: http://www.amazon.com/gp/product/1555582737/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=1555582737&linkCode=as2&tag=thebacgla-20&linkId=ANASKTZGVC7WT6NX
  [Markdown]: http://daringfireball.net/projects/markdown/syntax
  [Pandoc]: http://johnmacfarlane.net/pandoc/
  [John MacFarlane's]: http://johnmacfarlane.net/
  [LaTeX]: http://www.latex-project.org/
  [Vim]: http://www.vim.org/
  [MacVim]: https://code.google.com/p/macvim/
  [Git]: http://git-scm.com/
  [GitHub]: https://github.com/lmullen
  [GNU Make]: http://bost.ocks.org/mike/make/
  [ZSH]: http://www.zsh.org/
  [dotfiles]: https://github.com/lmullen/dotfiles
  [Vim files]: https://github.com/lmullen/vimrc
  [R]: http://www.r-project.org/
  [Hadley Wickham]: http://had.co.nz/
  [dplyr]: http://cran.r-project.org/web/packages/dplyr/index.html
  [tidyr]: http://cran.r-project.org/web/packages/tidyr/index.html
  [ggplot2]: http://ggplot2.org/
  [Yihui Xie's]: http://yihui.name/
  [knitr]: http://yihui.name/knitr/
  [R Markdown]: http://rmarkdown.rstudio.com/
  [RStudio]: http://www.rstudio.com/products/RStudio/
  [Ruby]: https://www.ruby-lang.org/en/
  [Mike Bostock's]: http://bost.ocks.org/mike/
  [D3.js]: http://d3js.org/
  [maps]: http://lincolnmullen.com/projects/slavery/
  [Zotero]: https://www.zotero.org/
  [Gitit]: http://gitit.net/
  [Jekyll]: http://jekyllrb.com/
  [Omeka]: http://omeka.org/
  [WordPress]: https://wordpress.org/
  [Reclaim Hosting]: http://reclaimhosting.com/
  [Ubuntu 14.04 LTS]: http://www.ubuntu.com/
  [Vagrant]: https://www.vagrantup.com/
  [VirtualBox]: https://www.virtualbox.org/
  [Homebrew]: http://brew.sh/
  [iTerm 2]: http://iterm2.com/
  [OmniFocus]: https://www.omnigroup.com/omnifocus
  [BackBlaze]: https://www.backblaze.com/
  [DropBox]: https://db.tt/f5srxqL
  [Caffeine]: http://lightheadsw.com/caffeine/
  [QGIS]: http://www.qgis.org/en/site/
  [Tranmit]: http://panic.com/transmit/
  [TextExpander]: http://smilesoftware.com/TextExpander/index.html
  [Spectactle]: http://spectacleapp.com/
  [Feedly]: http://feedly.com/
  [MIT]: http://opensource.org/licenses/MIT
  [CC-BY]: http://creativecommons.org/licenses/by/4.0/
  [arXiv]: http://arxiv.org/
