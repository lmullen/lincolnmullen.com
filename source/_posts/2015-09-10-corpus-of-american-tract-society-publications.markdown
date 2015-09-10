---
layout: post
title: Corpus of American Tract Society Publications
...

I've created a small to mid-sized corpus of publications by the American Tract Society up to the year 1900 in plain text. This corpus has been gathered from the [Internet Archive](https://archive.org/). It includes 641 documents containing just under sixty million words, along with a CSV file containing metadata for each of the files. I don't make any claims that this includes all of the ATS publications from that time period, and it is pretty obvious that the metadata from the Internet Archive is not much good. The titles are mostly correct; the dates are pretty far off in cases.

This corpus was created for the purpose of testing document similarity and text reuse algorithms. I need a corpus for testing the [textreuse](https://github.com/lmullen/textreuse), which is in very early stages of development. From reading many, many of these tracts, I already know the patterns of text reuse. (And of course, the documents are historically interesting in their own right, and might be a good candidate for text mining.) The ATS frequently republished tracts under the same title. Furthermore, they published volumes containing the entire series of tracts that they had published individually. So there are examples of entire documents which are reprinted, but also some documents which are reprinted inside others. Then as a extra wrinkle, the corpus contains the editions of the Bible published by the ATS, plus their edition of Cruden's concordance and a Bible dictionary. Likely all of the tracts quote the Bible, some at great length, so there are many examples of borrowing there.

Here is the corpus and its repository:

-   [ats\_corpus.zip](http://lincolnmullen.com/files/ats_corpus.zip) (122 Mb)
-   [GitHub repository to reproduce the corpus](https://github.com/lmullen/ats-corpus)
