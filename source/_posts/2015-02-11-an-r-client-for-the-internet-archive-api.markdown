---
categories:
- rstats
date: '2015-02-11 09:15'
layout: post
title: An R Client for the Internet Archive API
...

{% img right /figures/internet-archive/internet-archive-logo.png "Internet Archive Logo" %}

In support of some of my research projects, I created a simple R package to access the [Internet Archive's](https://archive.org/) API. The package is intended to search for items, to retrieve their metadata in a usable form, and to download the files associated with the items. The package, called `internetarchive`, is [available on GitHub](https://github.com/lmullen/internetarchive). The README and the [vignette](http://lincolnmullen.com/research/internetarchive/) have a full explanation, but here is a brief overview.

First, you can do keyword searches:

    library("internetarchive")
    ia_keyword_search("isaac hecker")

Next, you can do advanced searches, specifying which fields you want to search:

    ia_search(c("publisher" = "american tract society",
                date = "1840 TO 1850"))

Having retrieve a list of items using either of those search functions, you can get get the items and their associated metadata in a data frame. Here we use [magrittr's](https://github.com/smbache/magrittr) pipe operator (`%>%`) to create a pipeline:

    ia_keyword_search("isaac hecker") %>%
      ia_get_items() %>%
      ia_metadata()

My intended use is for downloading the files associated with items to create a corpus of texts. In this example, we search for items and download only the plain text files. The filtering is provided by [dplyr](https://github.com/hadley/dplyr).

    library("dplyr")
    ia_keyword_search("isaac hecker") %>%
      ia_get_items() %>%
      ia_files() %>%
      filter(type == "txt") %>%
      ia_download(dir = "texts")

The functions `ia_metadata()`, `ia_files()`, and `ia_download()` all return data frames, which should be easily filtered, reshaped, and joined as necessary. I hope the package is useful for creating corpora for text mining as well as for downloading sources to read in batches.

Doubtless there are parts of the Internet Archive API, especially in the advanced search and file types, that I haven't adequately explored. I'll be glad for [bug reports](https://github.com/lmullen/internetarchive/issues).
