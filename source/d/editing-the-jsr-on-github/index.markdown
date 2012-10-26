---
layout: blank
title: "Editing the JSR on GitHub"
date: 2012-10-26 14:32
---

These are brief instructions on how to edit the content of the *Journal
of Southern Religion*. The *JSR*'s code is hosted on [GitHub][]. These
instructions will tell you how to get set up at GitHub, how to access
the content you want to edit, and how to make changes. Don't worry: *you
can't break anything*. Git keeps track of every change that is ever
made, so it's easy to go back and fix mistakes. And the files that you
will be editing have to be uploaded to the site, so no one will see any
mistakes.

## Get set up at GitHub

1.  [Sign up for a free account][] on GitHub.

2.  [E-mail me][] your GitHub username and password so that I can give
    you permission to edit the *JSR*'s content.

3.  Sign in to GitHub.

## Find the content you want to edit

1.  The *JSR*'s GitHub repository is here:
    <https://github.com/lmullen/jsr>

2.  One you've gone there, you have to get on the right Git branch. A
    branch is just a different version of the site, probably an
    unpublished version. If the issue you want to edit has already be
    made public at <http://jsr.fsu.edu> then you don't have to switch
    branches at all because you're already on the `master` branch. If
    the issue you want has not been published, you need to switch to the
    branch named after that issue. Switching branches will look like
    this:

    {% img center /d/editing-the-jsr-on-github/jsr-2.png %}

3.  All of the issues are stored in the directory `_source/issues`,
    which you can navigate to in GitHub. The page will look like this:

    {% img center /d/editing-the-jsr-on-github/jsr-1.png %}

    From there you can click on the directory for the issue that you
    want, e.g., `vol14`. Click on that directory to get to the issue
    contents.

4.  You'll now see a list of all the articles in the issue. Click on the
    article you want to edit.

## Edit an article

1.  Near the top at the right of the page you want to edit is a button
    that says `Edit`. Click that.

    {% img center /d/editing-the-jsr-on-github/jsr-3.png %}

2.  Now you'll be in an editor that you use to change the content. It
    will look like this.

    {% img center /d/editing-the-jsr-on-github/jsr-4.png %}

3.  If you want to change information such as the author or title, then
    you'll have to edit the metadata. The metadata is the information at
    the top of the file in between lines that look like this: `---`. For
    example, the line for the author might look like this:
    `author: Art Remillard`. Change whatever you need to.

4.  If you need to change the body of the article, then you can change
    the rest of the file. The *JSR* articles are written in a format
    called Markdown, which is pretty straightforward. Here is an
    [explanation of Markdown][].

    Italics look like this: `The book *Your Title Here* is good.`

    Footnote markers look like this: `end of sentence.[^14]`

    Footnotes themselves look like this (they're at the end of the
    file): `[^14]: Author, *Book* (2012), 23.`

    Em dashes: `word---another word`. En dashes: `from 1776--1780`.

    Headers: `## My Header`

5.  Once you've made your corrections, you need to save the changes. Git
    calls this a "commit." At the bottom of the file, you'll see a box
    labeled `Commit summary`. In that box type a brief description of
    what you did, e.g., `Fixed typos`. It will look like this.

    {% img center /d/editing-the-jsr-on-github/jsr-5.png %}

6.  Click `Commit Changes` and you're done! You can go on to edit other
    files if you wish.

7.  I'll probably notice that you made changes. But to be safe, please
    send me an e-mail saying that there are changes to be added to the
    site.

  [GitHub]: http://https://github.com/
  [Sign up for a free account]: https://github.com/signup/free
  [E-mail me]: mailto:lincoln@lincolnmullen.com
  [explanation of Markdown]: http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295
