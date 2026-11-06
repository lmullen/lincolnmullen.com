---
title: "Pandoc converts all your (text) documents"
date: "2012-02-23"
slug: "pandoc-converts-all-your-text-documents"
description: "Pandoc converts plain-text documents written in Markdown or LaTeX into word processor files, HTML, PDFs, and more, making plain-text writing practical."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="diagram.png" alt="Pandoc conversion network" >}}

For the past few months we ProfHackers have been running an occasional series about using the command line. I got us started with a couple posts [explaining why you might want to use the command line](http://chronicle.com/blogs/profhacker/the-profhacker-guide-to-the-command-line/36125) and [how to get started using it](http://chronicle.com/blogs/profhacker/getting-comfortable-on-the-command-line/36862). Konrad followed with a posts about [the `uniq` command](http://chronicle.com/blogs/profhacker/using-the-uniq-command/36962) and [the `sort` command](http://chronicle.com/blogs/profhacker/using-the-sort-command-2/36791) for working with text and data files. Amy added a post about [how the command line let her hack the NOOK Color](http://chronicle.com/blogs/profhacker/why-the-command-line-is-your-friend/38009), and I wrote about [using `pdftk` to manipulate PDFs](http://chronicle.com/blogs/profhacker/fix-pdfs-quickly-with-pdftk/37332).

<!--more-->

Taking up the command line is easier if you have a specific problem you’re trying to solve. For me, the problem was that I wanted to do all of my writing in a [plain text format, like Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295) or [LaTeX](http://chronicle.com/blogs/profhacker/getting-started-with-latex/23092). But I need to be able to share my writing in a variety of formats: HTML for the web, PDF for printed documents or academic writing, and occasionally RTF or Microsoft Word or OpenOffice.

The best way I’ve found to move between these formats is [Pandoc](http://johnmacfarlane.net/pandoc/). Pandoc is a command line tool written by a philosophy professor, [John MacFarlane](http://johnmacfarlane.net/). Its general use is to take a document in one format and convert it to another. You can get an idea of the wide variety of formats Pandoc can translate by looking at an [enlargement of the header diagram](http://johnmacfarlane.net/pandoc/diagram.png).

Here’s an example of how this works. Suppose that you have a Markdown document like the one we created for the post on Markdown. ([View `pandoc-example.markdown` on GitHub](https://gist.github.com/c26556c14a1834af5906#file_pandoc_example.markdown).) You can convert this to a number of text formats with a simple terminal command:

**Markdown to HTML ([HTML output on GitHub](https://gist.github.com/c26556c14a1834af5906#file_pandoc_example.html)):**

`pandoc pandoc-example.markdown -o pandoc-example.html`

**Markdown to LaTeX ([LaTeX output on GitHub](https://gist.github.com/c26556c14a1834af5906#file_pandoc_example.tex)):**

`pandoc pandoc-example.markdown -o pandoc-example.tex`

**Markdown to DOCX:**

`pandoc pandoc-example.markdown -o pandoc-example.docx`

**Markdown to PDF ([download PDF](http://chronicle.com/blogs/profhacker/files/2012/02/pandoc-example.pdf)):**

`pandoc pandoc-example.markdown -o pandoc-example.pdf`

That command calls `pandoc`, tells it which file to convert (`pandoc-example.markdown`) and tells it which file to export (e.g., `pandoc-example.html`). Pandoc figures out what types of files these are from the extension, or you can pass it additional arguments. For some of the formats, you can convert the other way. For example, you could convert LaTex to Markdown or to a Word DOCX, or HTML to Markdown or LaTeX.  To convert to PDF, though, you’ll need to have LaTeX installed on your system.

Another useful thing that Pandoc can do is take a URL and convert the webpage to another format. For example, this command turns a page on my website into Markdown.

`pandoc -s -r html http://lincolnmullen.com/writing.html -o test.markdown`

You can see many more uses for Pandoc on its [example page](http://johnmacfarlane.net/pandoc/demos.html), and you can try some conversions with its [online demo](http://johnmacfarlane.net/pandoc/try).

There are several pros to using Pandoc. It’s [easy to install](http://johnmacfarlane.net/pandoc/installing.html) if you use the binaries for Windows or Mac. (It was a bit of a pain for me to compile from source, but there’s no reason you’d need to do that.) The tool is under active development, so bugs are being fixed and occasionally new formats are added. And there are quite a few advanced things that you can do, like create EPUB e-books and automatically generate citations using citeproc-hs and bibliographies like BibTeX (which you can export from Zotero). There are some conversions that it would be nice if Pandoc could do, but it can’t. For example, Pandoc can turn Markdown into a Word DOCX, but it can’t turn a DOCX into Markdown, HTML, etc., because of the limitations of the DOCX format.

If you do your writing in plain text or a markup format like LaTeX, [Pandoc](http://johnmacfarlane.net/pandoc/) is an essential, everyday tool for moving between formats. And if you occasionally need to turn HTML into other formats, it’s handy to have [Pandoc](http://johnmacfarlane.net/pandoc/) in your toolkit.

**Have you tried Pandoc? What uses have you found for it?**
