---
title: 'New release: tokenizers v0.2.0'
date: "2018-03-23"
slug: 'new-release-tokenizers-v0-2-0'
description: "tokenizers v0.2.0 adds tweet tokenization, document splitting, word counts, and Text Interchange Format support for better interoperability in R text analysis."
tags:
- R
- tokenizers
draft: false
---

A new v0.2.0 release of the [tokenizers package](http://lincolnmullen.com/software/tokenizers/) for R is now [available on CRAN](https://cran.r-project.org/package=tokenizers). The tokenizers package provides a fast, consistent set of functions to turn natural language text into tokens. This is a fairly substantial release which is a result of collaboration with Dmitry Selivanov, Os Keyes, Ken Benoit, and Jeffrey Arnold. This version adds these features (see the [changelog](http://lincolnmullen.com/software/tokenizers/news/index.html) for more details and attribution):

-   A new tokenizer for tweets that preserves usernames, hashtags, and URLS.
-   A new tokenizer for Penn Treebank style tokenization.
-   A new function to split long documents into pieces of equal length.
-   New functions to count words, characters, and sentences without tokenization.
-   The package now uses C++98 rather than C++11, so more users will be able to install it without upgrading their compiler. (No more e-mails from CentOS 6 users.)

Most important, the package implements the draft recommendations of the [Text Interchange Format](https://github.com/ropensci/tif). The TIF standards were drafted at the [2017 rOpenSci Text Workshop](http://textworkshop17.ropensci.org/). They define standards for tokens, corpora, and document-term matrices to allow R text analysis packages to interoperate with one another. I think these standards, once finalized and widely adopted, will be a very positive development for the coherence of the ecosystem of packages around text analysis in R. A [new vignette](http://lincolnmullen.com/software/tokenizers/articles/tif-and-tokenizers.html) explains how the tokenizers package fits into this ecosystem.

Finally, the package now has a [new website](http://lincolnmullen.com/software/tokenizers/index.html) for vignettes and documentation, thanks to [pkgdown](http://pkgdown.r-lib.org/). What the package does not have is a nice hex-sticker logo, but perhaps that can come in due time.
