---
layout: post
title: "A Makefile to Convert All Markdown Files to PDFs Using Pandoc"
date: 2012-09-22 14:55
comments: true
categories: 
- hacks
tags:
- Makefile
- Pandoc
- Markdown
---

<!-- **Update May 30, 2013:** I now do the same thing in [Rake instead of
Make][]. -->

I write all my documents in [Markdown][] and pipe them through
[Pandoc][] to make PDFs. Typing `pandoc myfile.md -o myfile.md.pdf` over
and over at the command line gets tedious, especially if the [Pandoc][]
command requires options for BibTeX and CSL files.

My solution is to use [GNU Make][] (helpful [introduction][] here). Make
lets you define a series of rules to compile files. You define 'targets'
(the files to be made), 'dependencies' (the source files), and the
commands that express the relation between the two. A typical rule using
[Pandoc][] might look like this:

{% highlight make %}
lesson.01.md.pdf : lesson.01.md
    pandoc lesson.01.md -o lesson 01.md.pdf
{% endhighlight %}

But in larger projects, such as the [class][] that I'm teaching, every
time I added a Markdown file I also had to add a new rule to the
`Makefile`. Besides the time this took, that process was error-prone,
especially if I ever changed the name of the file. The `Makefile` for
the class was over 60 lines long---and that's at the start of the
semester. 

So this afternoon I took some time to figure out how to write a Makefile
which will run Pandoc on any Markdown file in a directory. I've put the
file up as a [Gist][] on GitHub. Here is how it works.

First we define a variable with the list of all the files we want to
make. We get that list by using the `wildcard` function to find all the
`*.md` files, then by using the `patsubst` function to append `.pdf` to
those file names. 

{% highlight make %}
PDFS := $(patsubst %.md,%.md.pdf,$(wildcard *.md))
{% endhighlight %}

Next we define a generic target, `all`, using that list.

{% highlight make %}
all : $(PDFS)
{% endhighlight %}

We need a rule for making PDFs. Using Make's string replacement
macros, we define a generic rule:

{% highlight make %}
%.md.pdf : %.md
        pandoc $< -o $@
{% endhighlight %}

Now running `make` creates or updates all the PDFs in the project. Four
lines of code that are portable to any project beat 60+ lines that are
specific to the project.

Here is the whole `Makefile` as a [Gist][] (with a few additional
rules):

{% highlight make %}
# Produce PDFs from all Markdown files in a directory
# Lincoln Mullen | http://lincolnmullen.com | lincoln@lincolnmullen.com

# List files to be made by finding all *.md files and appending .pdf
PDFS := $(patsubst %.md,%.md.pdf,$(wildcard *.md))

# The all rule makes all the PDF files listed
all : $(PDFS)
                        
# This generic rule accepts PDF targets with corresponding Markdown 
# source, and makes them using pandoc
%.md.pdf : %.md
        pandoc $< -o $@

# Remove all PDF outputs
clean :
        rm $(PDFS)

# Remove all PDF outputs then build them again
rebuild : clean all
{% endhighlight %}

The filenames of my PDFs are a little peculiar: if the source file is
`handout.md`, I name the PDF `handout.md.pdf` instead of `handout.pdf`.
That's a deliberate choice, so that when my shell completes filenames I
get the full name of the Markdown file instead of having to make a
choice between `.md` and `.pdf`. Nine times out of ten I want to edit
the Markdown file instead of doing something with the PDF.

[Markdown]: http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295
[Pandoc]: http://johnmacfarlane.net/pandoc/
[GNU Make]: http://www.gnu.org/software/make/
[introduction]: http://www.cprogramming.com/tutorial/makefiles.html
[class]: http://lincolnmullen.com/courses/uws/
[Gist]: https://gist.github.com/3767386
[Rake instead of Make]: /blog/rake-and-pandoc/

