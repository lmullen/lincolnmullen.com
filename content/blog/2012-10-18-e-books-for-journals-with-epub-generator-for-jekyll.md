---
date: '2012-10-18T15:24:00+00:00'
slug: 'e-books-for-journals-with-epub-generator-for-jekyll'
title: 'E-books for Journals with EPUB Generator for Jekyll'
---

The *JSR* [editors](http://jsr.fsu.edu/about/masthead.html) been making a lot of changes at the *[Journal of Southern Religion](http://jsr.fsu.edu)* in the past year. We've [rebuilt the site](http://jsr.fsu.edu/blog/announcement/site-update.html) from the ground up, we've [released our content under an open-access license](http://jsr.fsu.edu/blog/announcement/open-access-license.html) (written up [here](http://oaopenaccess.wordpress.com/2012/08/27/open-access-journal-of-southern-religion-adopts-creative-commons-attribution-license/)), we've added a [successful podcast](http://jsr.fsu.edu/new-media/) with seven episodes already, and we're getting close to adding DOIs and releasing a new issue. Sometime soon I want to write a full 'making of' post about the decisions we made and the tools we used. Today, though, I want to introduce a Ruby script I've written for the *JSR* to make e-books in [EPUB](http://idpf.org/epub) format of each issue.

The *JSR*'s website runs on [Jekyll](http://jekyllrb.com/), a 'simple, blog aware, static site generator' that converts Markdown files into a website. In other words, *JSR* articles are stored in plain text files like [these](https://github.com/lmullen/jsr/tree/master/_source/issues/vol13) from issue 13. Stephen Ramsay explains [why this is a good thing](http://stephenramsay.us/2011/06/13/back-to-the-roots-web-with-jekyll.html). Because the *JSR* content is portable and parsable, it's easy to use the content for purposes other than the website itself.

One of the new things I want to do with the *JSR* is make an e-book of each issue, so that scholars can read the journal on tablets and mobiles more immersively than is possible on a website. Whether there interest among *JSR* readers in this new format remains to be seen, but other people are [writing about EPUB for scholars](http://blogs.plos.org/mfenner/2011/01/23/beyond-the-pdf-%E2%80%A6-is-epub/) and other open-access journals, like the *[Journal of Digital Humanities](http://journalofdigitalhumanities.org/)* are releasing their content as EPUB.

I've written a Ruby script, `epub-jekyll`, which parses posts and pages in a Jekyll website and passes them through [Pandoc](johnmacfarlane.net/pandoc/) to create an EPUB. Because the script uses a YAML file as a table of contents, it should be useable on other [Jekyll](http://jekyllrb.com/) (or [Octopress](http://octopress.org/)) websites. You can **[get the code from GitHub](https://github.com/lmullen/jekyll-ebook)** and learn how to use the script in the README.

This my first project in Ruby, so I'll be glad for any improvements you'd like to suggest.

**Update 2012-11-01:** I learned that there is another project called [jekyll-epub](https://rubygems.org/gems/jekyll-epub) so I had to change the name of this project to `jekyll-ebook`. The links above have been corrected. And I've also published the project [as a gem](https://rubygems.org/gems/jekyll-ebook).
