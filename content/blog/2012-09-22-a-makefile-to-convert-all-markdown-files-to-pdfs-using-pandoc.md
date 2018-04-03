---
date: '2012-09-22T18:55:00+00:00'
slug: 'a-makefile-to-convert-all-markdown-files-to-pdfs-using-pandoc'
title: A Makefile to Convert All Markdown Files to PDFs Using Pandoc
---

I write all my documents in [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295) and pipe them through [Pandoc](http://johnmacfarlane.net/pandoc/) to make PDFs. Typing `pandoc myfile.md -o myfile.md.pdf` over and over at the command line gets tedious, especially if the [Pandoc](http://johnmacfarlane.net/pandoc/) command requires options for BibTeX and CSL files.

My solution is to use [GNU Make](http://www.gnu.org/software/make/) (helpful [introduction](http://www.cprogramming.com/tutorial/makefiles.html) here). Make lets you define a series of rules to compile files. You define 'targets' (the files to be made), 'dependencies' (the source files), and the commands that express the relation between the two. A typical rule using [Pandoc](http://johnmacfarlane.net/pandoc/) might look like this:
<figure class="highlight">
<pre><code class="language-make" data-lang="make">&lt;span class="nl">lesson.01.md.pdf &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">lesson.01.md&lt;/span>
    &lt;span class="err">pandoc&lt;/span> &lt;span class="err">lesson.01.md&lt;/span> &lt;span class="err">-o&lt;/span> &lt;span class="err">lesson&lt;/span> &lt;span class="err">01.md.pdf&lt;/span></code></pre>
</figure>
But in larger projects, such as the [class](http://lincolnmullen.com/courses/uws/) that I'm teaching, every time I added a Markdown file I also had to add a new rule to the `Makefile`. Besides the time this took, that process was error-prone, especially if I ever changed the name of the file. The `Makefile` for the class was over 60 lines long---and that's at the start of the semester.

So this afternoon I took some time to figure out how to write a Makefile which will run Pandoc on any Markdown file in a directory. I've put the file up as a [Gist](https://gist.github.com/3767386) on GitHub. Here is how it works.

First we define a variable with the list of all the files we want to make. We get that list by using the `wildcard` function to find all the `*.md` files, then by using the `patsubst` function to append `.pdf` to those file names.
<figure class="highlight">
<pre><code class="language-make" data-lang="make">&lt;span class="nv">PDFS&lt;/span> &lt;span class="o">:=&lt;/span> &lt;span class="err">$&lt;/span>&lt;span class="o">(&lt;/span>patsubst %.md,%.md.pdf,&lt;span class="err">$&lt;/span>&lt;span class="o">(&lt;/span>wildcard &lt;span class="k">*&lt;/span>.md&lt;span class="o">))&lt;/span></code></pre>
</figure>
Next we define a generic target, `all`, using that list.
<figure class="highlight">
<pre><code class="language-make" data-lang="make">&lt;span class="nl">all &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">$(PDFS)&lt;/span></code></pre>
</figure>
We need a rule for making PDFs. Using Make's string replacement macros, we define a generic rule:
<figure class="highlight">
<pre><code class="language-make" data-lang="make">&lt;span class="nl">%.md.pdf &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">%.md&lt;/span>
        &lt;span class="err">pandoc&lt;/span> &lt;span class="err">$&lt;&lt;/span> &lt;span class="err">-o&lt;/span> &lt;span class="err">$@&lt;/span></code></pre>
</figure>
Now running `make` creates or updates all the PDFs in the project. Four lines of code that are portable to any project beat 60+ lines that are specific to the project.

Here is the whole `Makefile` as a [Gist](https://gist.github.com/3767386) (with a few additional rules):
<figure class="highlight">
<pre><code class="language-make" data-lang="make">&lt;span class="c"># Produce PDFs from all Markdown files in a directory
# Lincoln Mullen | http://lincolnmullen.com | lincoln@lincolnmullen.com
&lt;/span>
&lt;span class="c"># List files to be made by finding all *.md files and appending .pdf
&lt;/span>&lt;span class="nv">PDFS&lt;/span> &lt;span class="o">:=&lt;/span> &lt;span class="err">$&lt;/span>&lt;span class="o">(&lt;/span>patsubst %.md,%.md.pdf,&lt;span class="err">$&lt;/span>&lt;span class="o">(&lt;/span>wildcard &lt;span class="k">*&lt;/span>.md&lt;span class="o">))&lt;/span>

&lt;span class="c"># The all rule makes all the PDF files listed
&lt;/span>&lt;span class="nl">all &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">$(PDFS)&lt;/span>
                        
&lt;span class="c"># This generic rule accepts PDF targets with corresponding Markdown 
# source, and makes them using pandoc
&lt;/span>&lt;span class="nl">%.md.pdf &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">%.md&lt;/span>
        &lt;span class="err">pandoc&lt;/span> &lt;span class="err">$&lt;&lt;/span> &lt;span class="err">-o&lt;/span> &lt;span class="err">$@&lt;/span>

&lt;span class="c"># Remove all PDF outputs
&lt;/span>&lt;span class="nl">clean &lt;/span>&lt;span class="o">:&lt;/span>
        &lt;span class="err">rm&lt;/span> &lt;span class="err">$(PDFS)&lt;/span>

&lt;span class="c"># Remove all PDF outputs then build them again
&lt;/span>&lt;span class="nl">rebuild &lt;/span>&lt;span class="o">:&lt;/span> &lt;span class="nf">clean all&lt;/span></code></pre>
</figure>
The filenames of my PDFs are a little peculiar: if the source file is `handout.md`, I name the PDF `handout.md.pdf` instead of `handout.pdf`. That's a deliberate choice, so that when my shell completes filenames I get the full name of the Markdown file instead of having to make a choice between `.md` and `.pdf`. Nine times out of ten I want to edit the Markdown file instead of doing something with the PDF.
