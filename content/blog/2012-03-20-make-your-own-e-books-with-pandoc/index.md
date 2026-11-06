---
title: "Make your own e-books with Pandoc"
date: "2012-03-20"
slug: "make-your-own-e-books-with-pandoc"
description: "How to turn blog posts, open-access books, or journal articles into EPUB e-books using Pandoc."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/make-your-own-e-books-with-pandoc/39067"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="pages_to_cover-300x225.jpg" alt="Pages being bound into a book cover" >}}

As devices for reading e-books proliferate, it increasingly makes sense to make publications available in an e-book. There are a number of cases in which you might do this:

<!--more-->

- If you have a blog and want to make the best posts into an e-book. For example, sociologist [Kieran Healy](http://kieranhealy.org/blog/archives/2011/12/08/books-i-did-not-read-this-year-an-ebook/) created an [e-book of posts from his blog](http://kieranhealy.org/blog/archives/2011/12/08/books-i-did-not-read-this-year-an-ebook/).
- If you have content in one format that you want to read as an e-book instead. For example, our own [Mark Sample](http://www.samplereality.com/) took the open-access book [*Hacking the Academy*](http://hackingtheacademy.org/) and turned it into [several versions of an e-book](http://www.samplereality.com/2011/09/10/hacking-the-academy-the-ebook-volume/) for Nooks, iPads, and Kindles.
- If you want to give your readers the option to read your content as an e-book. For example, I’m the web editor for [*The Journal of Southern Religion*](http://jsr.fsu.edu/). As a supplement to the articles on the website, I intend to make an e-book of each issue with the articles “bound” together.

Making an e-book can be easy—almost trivially easy—using [Pandoc](http://johnmacfarlane.net/pandoc/), a tool [I’ve written about earlier](http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700) on ProfHacker. Of course, Pandoc isn’t the only way to do this. Mark has used [Sigil and written about it](http://chronicle.com/blogs/profhacker/create-your-own-e-book-with-open-source-sigil/36501), and another good option is [Anthologize](http://anthologize.org/), which [Julie wrote about](http://chronicle.com/blogs/profhacker/one-week-one-tool-anthologize/25972).

Since a good way to learn is by doing, let’s create an e-book of *My Favorite ProfHacker Posts*. You can follow along with the commands below and [these files hosted on GitHub](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo), and you can also look at [the Pandoc documentation](http://johnmacfarlane.net/pandoc/epub.html). We’ll use just three posts to keep it simple:

- Billie’s [“Writer’s Bootcamp: The DRAW Method”](http://chronicle.com/blogs/profhacker/writers-bootcamp-the-draw-method/23097)
- George’s [“The Simplicity of ‘Think-Pair-Share’”](http://chronicle.com/blogs/profhacker/the-simplicity-of-think-pair-share/36094)
- Mark’s [“Teaching for Uncoverage rather than Coverage”](http://chronicle.com/blogs/profhacker/teaching-for-uncoverage-rather-than-coverage/35459)

Here’s how to make an e-book with a few commands.

### 1. Get a clean copy of each of the “chapters.”

You’ll need a clean, plain-text copy of each of the documents that will go into your e-book. You might have them in your blog, in an HTML file, or perhaps in a [LaTeX](http://chronicle.com/blogs/profhacker/getting-started-with-latex/23092) file. Either Markdown or HTML is fine. In this case, since I don’t have the originals of these posts, I’m going to use Pandoc to grab them from the web and convert them to [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295). (I explained how to do this in [the previous post about Pandoc](http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700).) Let’s grab Billie’s post and turn it into Markdown; you’d do the same thing for the other two posts.

`pandoc -s -r html http://chronicle.com/blogs/profhacker/writers-bootcamp-the-draw-method/23097 -o ch01.hara.markdown`

You’ll have to delete the junk from the files, such as the *Chronicle’s* navigation bar, but this is easy. We’re left with three files, one each for the post by [Billie](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo/blob/master/ch01.hara.markdown), [George](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo/blob/master/ch02.williams.markdown), and [Mark](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo/blob/master/ch03.sample.markdown).

### 2. Create the front matter

Our e-book will need a little metadata. First we’ll make a simple text file which will be our title page. We’ll call it `title.txt`. It will have these two lines for the title and author. (See the [file on GitHub](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo/blob/master/title.markdown).)

```
% My Favorite ProfHacker Posts
% Team ProfHacker
```

We’ll also need to provide the equivalent of a copyright page, which gives some information about the book. This is also a simple two-line text file, which we’ll call `metadata.xml`. (See the [file on GitHub](https://github.com/lmullen/ProfHacker-Pandoc-E-book-Demo/blob/master/metadata.xml).)

```
<dc:rights>The copyright status of this e-book is ambiguous.</dc:rights>
<dc:language>en-US</dc:language>
```

### 3. Stitch the e-book together with Pandoc

Now that we have all the parts of the e-book, we can stitch it together with one Pandoc command. You would type this into the command line. (The `\`s tell the shell that you are breaking one long command into several lines.) The first line calls `pandoc`, tells it where to find `metadata.xml`. The second line tells `pandoc` to output (`-o`) an EPUB file named `my-favorite-profhacker-posts.epub`. The remaining lines list the parts of the book in order.

```
pandoc -S --epub-metadata=metadata.xml \
    -o my-favorite-profhacker-posts.epub \
    title.txt \
    ch01.hara.markdown \
    ch02.williams.markdown \
    ch03.sample.markdown
```

And that’s all there is to it: you’re a twenty-first-century Gutenberg. You can [*download our new e-book*](https://github.com/downloads/lmullen/ProfHacker-Pandoc-E-book-Demo/my-favorite-profhacker-posts.epub) from GitHub. Now you have an e-book in the [EPUB](http://idpf.org/epub) format, which works on Nooks, iPads, iPhones, and iPods, and which you can easily convert to the MOBI format for Kindle using [Calibre](http://chronicle.com/blogs/profhacker/converting-ebooks-with-calibre/29389).

**What uses do you have for e-books? Have you made your own e-book?** Let us know in the comments.

*Photo courtesy of Flickr user [Simon Zirkunow](http://www.fotopedia.com/items/flickr-506334817) / Creative Commons licensed*
