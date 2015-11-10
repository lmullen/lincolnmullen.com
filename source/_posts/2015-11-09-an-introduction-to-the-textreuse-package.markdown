---
layout: post
title: 'An Introduction to the textreuse Package, with Suggested Applications'
...

A number of problems in digital history/humanities require one to calculate the similarity of documents or to identify how one text borrows from another. To give one example, the [Viral Texts](http://viraltexts.org/) project, by Ryan Cordell, David Smith, et al., has been very successful at identifying reprinted articles in American newspapers.[^1] [Kellen Funk](http://kellenfunk.org/) and I have been working on a text reuse problem in nineteenth-century legal history, [where we seek to track](http://lincolnmullen.com/blog/detecting-text-reuse-in-legal-documents/) how codes of civil procedure were [borrowed and modified](http://kellenfunk.org/field-code/) in jurisdictions across the United States.

As part of that project, I have recently released the [textreuse package](https://github.com/ropensci/textreuse) for [R](https://www.r-project.org/) to [CRAN](https://cran.rstudio.com/web/packages/textreuse/). (Thanks to [Noam Ross](http://www.noamross.net/) for giving this package a [very thorough open peer review](https://github.com/ropensci/onboarding/issues/20) for [rOpenSci](https://ropensci.org/), to whom I've contributed the package.) This package is a general purpose implementation of several algorithms for detecting text reuse, as well as classes and functions for investigating a corpus of texts. Put most simply, full text goes in and measures of similarity come out.[^2] Put more formally, here is the package description:

> Tools for measuring similarity among documents and detecting passages which have been reused. Implements shingled n-gram, skip n-gram, and other tokenizers; similarity/dissimilarity functions; pairwise comparisons; minhash and locality- sensitive hashing algorithms; and a version of the Smith-Waterman local alignment algorithm suitable for natural language.

This package is intended for use with what one might label "medium data." (Thankfully, I think we're past the stage where we were obsessing about "[big data](http://tedunderwood.com/2013/05/10/why-it-matters-that-we-dont-know-what-we-mean-by-big-data/).") By that I mean the package works on data which is small enough to fit into memory on a single machine. But because I've written parts of the package in [C++ via Rcpp](http://www.rcpp.org/) and implemented the algorithms listed above, it also works quickly on corpora which are big enough to make naive approaches choke. While scholars who work with the largest corpora are doing fascinating work, I find that the type of scholarly questions I'm most interested in can be asked of "medium data" corpora which have been assembled for that purpose.

In this post I give an overview of how you might apply the textreuse package to a couple of kinds of historical problems. I'll also explain how some of the algorithms work, but you can skip that if you want. I won't explain how to use the package: for that you can start with the [README](https://github.com/lmullen/textreuse/blob/master/README.md) and the [introductory vignette](), or any of the other topical vignettes. It's a long post, so here are the topics at hand:

-   [Basic tools and concepts for document similarity](#basic-tools-and-concepts-for-document-similarity)
-   [How locality-sensitive hashing works](#how-locality-sensitive-hashing-works)
-   [De-duplicating a corpus with LSH](#de-duplicating-a-corpus-with-lsh)
-   [Local alignment for natural language](#local-alignment-for-natual-language)
-   [Tracing borrowings within a corpus](#tracing-borrowings-within-a-corpus)
    -   [Networks of borrowings](#networks-of-borrowings)
    -   [Heat maps of borrowed sections within a document](#heat-maps-of-borrowed-sections-within-a-document)
    -   [Clustering borrowings with affinity propagation](#clustering-borrowings-with-affinity-propagation)
-   [Conclusion](#conclusion)

### Basic tools and concepts for document similarity

Let's begin with a tiny corpus that ships with the package. It contains eight documents drawn from American Tract Society publications. We create a corpus object, which contains eight different document objects ([see them here](https://github.com/lmullen/textreuse/tree/master/inst/extdata/ats)). Each of these documents has an ID (based on the filename), the text of the document, and tokens generated from that text---in this case, 5-grams.[^3]

``` {.r}
library("textreuse")
library("dplyr")
dir <- system.file("extdata/ats/", package = "textreuse")
corpus <- TextReuseCorpus(dir = dir, tokenizer = tokenize_ngrams, n = 5,
                          keep_tokens = TRUE)
corpus
```

    ## TextReuseCorpus
    ## Number of documents: 8 
    ## hash_func : hash_string 
    ## tokenizer : tokenize_ngrams

We can inspect the corpus by getting the IDs of the documents, and looking at the first few tokens in *The Life of the Rev. Richard Baxter*.

``` {.r}
names(corpus)
```

    ## [1] "calltounconv00baxt"        "gospeltruth00whit"        
    ## [3] "lifeofrevrichard00baxt"    "memoirjamesbrai00ricegoog"
    ## [5] "practicalthought00nev"     "remember00palm"           
    ## [7] "remembermeorholy00palm"    "thoughtsonpopery00nevi"

``` {.r}
doc <- corpus[["lifeofrevrichard00baxt"]]
tokens(doc)[200:210]
```

    ##  [1] "t be king attempts to"                    
    ##  [2] "be king attempts to reconcile"            
    ##  [3] "king attempts to reconcile 4"             
    ##  [4] "attempts to reconcile 4 contents"         
    ##  [5] "to reconcile 4 contents the"              
    ##  [6] "reconcile 4 contents the conflicting"     
    ##  [7] "4 contents the conflicting parties"       
    ##  [8] "contents the conflicting parties declines"
    ##  [9] "the conflicting parties declines a"       
    ## [10] "conflicting parties declines a bishopric" 
    ## [11] "parties declines a bishopric for"

We can then use comparison functions to determine how similar the documents are. Here we use the [Jaccard similarity](https://en.wikipedia.org/wiki/Jaccard_index) function (which returns a ratio between `0` and `1`), though the package implements several other similarity measures. (The Jaccard similarity is defined as the intersection of the sets divided by the union of the sets.) For instance, two of the documents are almost the same. But neither of them is much like the *Life of Richard Baxter*.

``` {.r}
jaccard_similarity(corpus[["remember00palm"]],
                   corpus[["remembermeorholy00palm"]])
```

    ## [1] 0.7006189

``` {.r}
jaccard_similarity(corpus[["lifeofrevrichard00baxt"]],
                   corpus[["remembermeorholy00palm"]]) 
```

    ## [1] 0.0000359887

Most of the package works with `apply`-like functions that take a corpus of documents and apply various functions to it. Here we use the pairwise comparison function, which finds all the combination of documents in the corpus and compares them using a function that you specify. The result is a matrix of scores (which we will convert to a data frame of pairwise comparisons for legibility).

``` {.r}
pairwise_compare(corpus, jaccard_similarity) %>% 
  round(3) %>% 
  pairwise_candidates() %>% 
  arrange(desc(score))
```

    ## Source: local data frame [28 x 3]
    ## 
    ##                            a                         b score
    ##                        (chr)                     (chr) (dbl)
    ## 1             remember00palm    remembermeorholy00palm 0.701
    ## 2      practicalthought00nev    thoughtsonpopery00nevi 0.463
    ## 3         calltounconv00baxt    lifeofrevrichard00baxt 0.281
    ## 4         calltounconv00baxt memoirjamesbrai00ricegoog 0.002
    ## 5  memoirjamesbrai00ricegoog     practicalthought00nev 0.002
    ## 6         calltounconv00baxt         gospeltruth00whit 0.001
    ## 7         calltounconv00baxt     practicalthought00nev 0.001
    ## 8         calltounconv00baxt    thoughtsonpopery00nevi 0.001
    ## 9          gospeltruth00whit memoirjamesbrai00ricegoog 0.001
    ## 10         gospeltruth00whit     practicalthought00nev 0.001
    ## ..                       ...                       ...   ...

So far, so good. We identified three matches between pairs (all correct, by the way) and have no false positives in the other 25 pairs of potential matches.

Unfortunately, this naive approach won't work for a corpus much bigger than the one we are using here. The problem is that the number of comparisons to be made grows geometrically with the size of the corpus. If *n* is the number of documents in the corpus, then we need to make (*n*<sup>2</sup> – *n*) / 2 comparisons. Our 8 document corpus need 28 comparisons. A corpus with 100 documents would require 4,950 comparisons. The full ATS corpus has 641 documents requiring 205,120 comparisons; later we will work with a corpus of about 90,000 documents, requiring a little over 4 billion pairwise comparisons. Even if we could compute one thousand comparisons per second,[^4] we would still need 50 or more days worth of computation to get results on that last corpus.[^5]

## How locality-sensitive hashing works

We can get around the problem of geometric growth in the number of pairwise comparisons with one of the primary methods offered by the package, the use of minhashes and locality-sensitive hashing (LSH). This technique was first invented at AltaVista to detect duplicated documents on the web, but as far as I know there is no other implementation of LSH on CRAN. Here is an explanation of how LSH works (skip ahead to the applications, if you wish.) For a full explanation, see chapter 3 of *[Mining of Massive Datasets](http://www.mmds.org/)*, or for suggestions for implementing it, see "[MinHash for Dummies](http://matthewcasperson.blogspot.com/2013/11/minhash-for-dummies.html)."[^6]

Hashing is a technique for mapping arbitrary values to compact, uniform representations.[^7] You may have heard of password hashes or cryptographic hashes for verifying downloads, for instance. As an example, let's take a line from a song, turn it into word tokens, and then represent each word as an integer.

``` {.r}
song <- tokenize_words("the answer is blowin' in the wind")
hashed_tokens <- hash_string(song)
hashed_tokens
```

    ## [1]   -78149361 -1880366041  -845898986 -1592008830  -845898995   -78149361
    ## [7]  -904825919

The key idea is that the same tokens will always be matched to the same hashes. For instance, if we hash just the word "the," we will get the same integer that we did in the first word above.

``` {.r}
hash_string("the")
```

    ## [1] -78149361

``` {.r}
hash_string("the") == hashed_tokens[1]
```

    ## [1] TRUE

But even the slightest variation will result in a different hash. So the word "The" (with an initial capital) has a different integer representation altogether.[^8]

``` {.r}
hash_string("The") == hashed_tokens[1]
```

    ## [1] FALSE

Now suppose that we select the lowest integer in our vector of hashes. Which word does it represent?

``` {.r}
min(hashed_tokens)
```

    ## [1] -1880366041

``` {.r}
song[which.min(hashed_tokens)]
```

    ## [1] "answer"

The resulting integer has the value -1880366041, which corresponds to the word "answer" in our set of tokens. Assuming that our hash function has certain properties, we have in effect randomly selected a token from our document.

If we had a bunch of hash functions, we could select the minimum hash (= minhash) generated by each, and so we could randomly select many of the tokens in the document. We can do this by generating as many random integers as we want there to be minhashes and shifting the bits in the integers generated by our original hash function. Here we will create a function which generates 20 minhashes (we set a seed so that we can get the same results later). Then we apply that function to our document.

``` {.r}
minhash <- minhash_generator(n = 20, seed = 12231)
minhash(song) 
```

    ##  [1]    35226184 -2015303753 -1815664729 -1962848079 -1690615111
    ##  [6] -2015078269 -1832536439 -1975996581   180289931 -1925792566
    ## [11]   161460535 -1926044820   137848423 -2086997468 -1853820671
    ## [16]   328168917    63234299 -1899537726   447399568 -1827767748

We have thus selected 20 random tokens from our document. Notice that we are representing the document by 20 minhashes regardless of the length of our original document. So, we can represent a line from a Bob Dylan song in 20 minhashes, or all of *Moby Dick* in 20 minhashes. All of the documents can be represented uniformly and compactly with a random selection of tokens.

Minhashes have a curious property. Because they select random tokens, if we compare the minhashes, we can start to approximate the Jaccard similarity of the two documents. Suppose we modify the lyrics slightly, substituting "breeze" for "wind." The Jaccard similarity of these documents is high.

``` {.r}
song_mod <- tokenize_words("the answer is blowin' in the breeze")
jaccard_similarity(song, song_mod)
```

    ## [1] 0.7142857

If we generate 20 minhashes for each of these lyrics, we can compare their minhashes in a table that looks like this:

      song        song\_mod    equal   band  band\_match
  ------------- ------------- ------- ------ -------------
    35226184      94136964     FALSE    1    FALSE
   -2015303753   -2015303753   TRUE     1    
   -1815664729   -1815664729   TRUE     2    FALSE
   -1962848079   -1937459098   FALSE    2    
   -1690615111   -1690615111   TRUE     3    TRUE
   -2015078269   -2015078269   TRUE     3    
   -1832536439   -2138832230   FALSE    4    FALSE
   -1975996581   -1916975220   FALSE    4    
    180289931     180289931    TRUE     5    TRUE
   -1925792566   -1925792566   TRUE     5    
    161460535     236607968    FALSE    6    FALSE
   -1926044820   -1926044820   TRUE     6    
    137848423     137848423    TRUE     7    TRUE
   -2086997468   -2086997468   TRUE     7    
   -1853820671   -1853820671   TRUE     8    FALSE
    328168917     336239874    FALSE    8    
    63234299      63234299     TRUE     9    TRUE
   -1899537726   -1899537726   TRUE     9    
    447399568     489631303    FALSE    10   FALSE
   -1827767748   -1802792213   FALSE    10   

Some of the minhashes for the two documents are equal. We can assign those minhashes to "bands": that is, we can think of the first two rows as a unit, the second two rows as a unit, and so on. If all of the minhashes in any band match, then we can regard the documents a potential match. In this case, we have used 10 bands of 2 rows each, and we see that several of the bands have all of their rows match. So we are justified in regarding these two documents as a probable match.

What we have done here is called locality-sensitive hashing (LSH). We have turned our minhashes into a different representation of bands and rows, which can be thought of as hashes. But unlike other hashing techniques, where we try to put different inputs into different buckets, here we are trying to put different inputs into the same buckets if the inputs are reasonably close to one another.

This technique has several advantages. First, we can represent the documents compactly as minhashes. Second, we can run through documents and create minhash/LSH signatures once for each document. In other words, the amount of computation grows linearly with the size of our corpus, rather than geometrically. And we can add new documents to the corpus later using the same minhash function without having to redo all of our computation.

The question arises, then: how likely are we to detect a match using LSH? The answer depends on the actual Jaccard similarity of the documents (the thing we want to find out), as well as the number of minhashes and the number of bands, which controls the sensitivity of our search. (The number of minhashes must equal the number of bands multiplied by the number of rows.) We used 20 minhashes in 10 bands, and we were measuring documents with a Jaccard similarity of 0.71. So the probability that we would have detected a match is given here:

``` {.r}
lsh_probability(h = 20, b = 10, s = 0.71)
```

    ## [1] 0.9991006

So we are virtually guaranteed to get a match for these two documents which are very similar. If the documents were much less similar, say with a Jaccard similarity of 0.2, the probability of marking the documents as a candidate pair would be lower:

``` {.r}
lsh_probability(h = 20, b = 10, s = 0.2)
```

    ## [1] 0.3351674

In other words, we can decide on the number of minhashes and bands that we want to use based on the threshold Jaccard similarity that we want to detect. The relationship between the number of minhashes, the actual Jaccard similarity of the documents we are measuring, and the probability that we will detect a match, looks something like a butte, with a steep threshold. In other words, we can determine the number of bands, rows, and minhashes in such a way that we can be fairly confident that we will identify candidate matches above a certain threshold similarity, and not identify candidate matches for documents below that threshold.

{% figure field-code/threshold png "The threshold for detecting matches with LSH." %}

The package gives a function to help you decide on how many minhashes and bands to use. The resulting number is the threshold Jaccard similarity, above which matches are likely to be detected.

``` {.r}
lsh_threshold(100, 20)
```

    ## [1] 0.5492803

``` {.r}
lsh_threshold(1000, 200)
```

    ## [1] 0.3465724

``` {.r}
lsh_threshold(2000, 500)
```

    ## [1] 0.2114743

## De-duplicating a corpus with LSH

Now that we grok the basic idea behind locality-sensitive hashing, we can use LSH to solve our text reuse problems. The first example that I will give shows how the LSH functions in the textreuse package can be used to de-duplicate a messy corpus prior to analysis. This is, after all, what LSH was originally invented to do, and I find that I often have to clean up my corpora before I can do other work. We will use a [corpus of American Tract Society publications](http://lincolnmullen.com/blog/corpus-of-american-tract-society-publications/) which I downloaded from the Internet Archive using [another of my R packages](http://lincolnmullen.com/blog/an-r-client-for-the-internet-archive-api/). It contains 641 documents and about 60 million words: not big by any means, but large enough that a naive approach won't work.

Let's load and tokenize the corpus using 5-grams. We also create a minhash function and generate minhash signatures of the documents.

``` {.r}
options("mc.cores" = 6L) # Set this only if you want to use more CPU cores
ats_minhash <- minhash_generator(n = 200, seed = 923)
ats <- TextReuseCorpus(dir = "~/dev/ats-corpus/corpus/",
                       tokenizer = tokenize_ngrams, n = 5,
                       minhash_func = ats_minhash)
```

We will use 200 hashes and 50 bands with 4 rows in each band. That should detect any pairs of matches above a Jaccard similarity of about 0.1. That is a very low Jaccard similarity score: in other words, we have set the sensitivity for potential matches very high.

``` {.r}
lsh_threshold(h = 200, b = 100) %>% round(2)
```

    ## [1] 0.1

We can find all of the candidate matches with the LSH function. This takes less than 10 seconds on my computer.

``` {.r}
buckets <- lsh(ats, bands = 100)
```

The object that I'm calling `buckets` is a collection of documents and a representation of their banded minhashes.[^9] For now this is just a data frame and will only work for in-memory data (which can still be quite a lot), but I've implemented the package with an eye towards making it possible to use a database back end which will permit even bigger datasets.

We can extract potential matches by querying for a specific document. Let's look for matches to our old friend Richard Baxter.

``` {.r}
lsh_query(buckets, "lifeofrevrichard00baxt")
```

    ## Source: local data frame [2 x 3]
    ## 
    ##                        a                     b score
    ##                    (chr)                 (chr) (dbl)
    ## 1 lifeofrevrichard00baxt    calltounconv00baxt    NA
    ## 2 lifeofrevrichard00baxt calltounconverte00bax    NA

Or we can get all of the potential matches. Once we have potential matches, we must apply a comparison function to them. (This takes about 12 seconds on my computer.) Here are the results, looking at the highest pairs of matches.[^10]

``` {.r}
ats_matches <- buckets %>% 
  lsh_candidates() %>% 
  lsh_compare(ats, jaccard_similarity) 

ats_matches %>% arrange(desc(score))
```

    ## Source: local data frame [384 x 3]
    ## 
    ##                            a                          b     score
    ##                        (chr)                      (chr)     (dbl)
    ## 1  ecumenicalmissio003136mbp  ecumenicalmissio014136mbp 1.0000000
    ## 2  powerofchristian005274mbp  powerofchristian017371mbp 1.0000000
    ## 3           cu31924029353509      unclejohnvassar00vass 0.9607856
    ## 4     themorningcometh00burr     themorningcometh01burr 0.9582093
    ## 5     daughtersofarmen00whee     daughtersofarmen01whee 0.9574196
    ## 6           cu31924023494127 twentysixyearsof00stotrich 0.9555602
    ## 7         ministrywene00swee       ministryweneed00swee 0.9550249
    ## 8      twostjohnsofnew00stal     twostjohnsofnewt01stal 0.9506407
    ## 9      revivalsketches00hump     revivalsketchesm00hump 0.9430790
    ## 10        ministrywene00swee        ministrywenee00swee 0.9302349
    ## ..                       ...                        ...       ...

Out of the 205,120 pairs of documents in the corpus, we only had to compute similarity scores for 384 pairs, or 0.19 percent of all pairs.

Now that we have all the pairs of potential documents in the corpus with scores for their similarity, we can cluster the documents and decide which are in fact duplicates and which we might wish to remove. The textreuse package has no opinion about which clustering method you should use. For the sake of completeness, here is a plain-vanilla example using hierarchical clustering from base R.

We can construct a similarity matrix from our pairs of candidate matches. For hierarchical clustering we need a *dissimilarity* matrix, so we subtract our Jaccard scores from 1 (the maximum Jaccard score), turn the dissimilarity matrix into a distance object, and use hierarchical clustering.

``` {.r}
sim <- as.matrix(ats_matches)
dissim <- 1 - sim

clusters <- dissim %>% 
  as.dist() %>% 
  hclust(method = "average")

plot(clusters, labels = FALSE, main  = "Clusters of similar ATS documents")
```

{% figure field-code/unnamed-chunk-23-1 png "A dendrogram of similar ATS documents." %}

I've left the labels off so that the dendrogram is more legible. We can pick a height then cut the dendrogram into clusters.

``` {.r}
cuts <- cutree(clusters, h = 0.6)
data_frame(doc = names(cuts), cluster = cuts) %>% 
  arrange(cluster) %>% 
  head(16)
```

    ## Source: local data frame [16 x 2]
    ## 
    ##                          doc cluster
    ##                        (chr)   (int)
    ## 1   afflictedmanscom00will_0       1
    ## 2   afflictedmanscom00will_1       1
    ## 3  africapastandpr00moisgoog       2
    ## 4   africapastpresen00mois_0       2
    ## 5  ahistoryenglish00martgoog       3
    ## 6     historyofenglish00mart       3
    ## 7      alarmtounconver00alle       4
    ## 8   alarmtounconvert00alle_0       4
    ## 9   alarmtounconvert00alle_1       4
    ## 10     alarmtounconvert00all       5
    ## 11    religioneternall00pike       5
    ## 12 anewmemoirhanna02kniggoog       6
    ## 13          cu31924013194380       6
    ## 14          hannahmore00knig       6
    ## 15          hannahmore01knig       6
    ## 16    newmemoirofhanna01knig       6

Here we see the first six clusters. Judging by the titles they are accurately clustered. We might wish for cluster 5 to be a part of cluster 4, but we could refine our technique or iterate on the clustering depending on what we wanted. We would then select one document from each cluster to be used.

This method could be broadly useful for anyone who wants to do text analysis on a clean corpus of documents.

## Local alignment for natual language

Another major feature in this initial release of the package is a version of the Smith-Waterman local alignment algorithm suitable for natural language processing.[^11] Local alignment is the process of taking two substrings and finding the best possible matching substring, where "best" is determined by a scoring algorithm that assigns values for matches, mismatches, and gaps. The Smith-Waterman dynamic programming algorithm is used for aligning DNA or protein sequences. There are dozens and dozens of explanations of how the algorithm works online: it seems to be an obligatory lecture in any bioinformatics course; the [Wikipedia explanation](https://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm) is unusually lucid for a description of an algorithm. The only implementation of this algorithm in R is in Bioconductor's [Biostrings package](https://bioconductor.org/packages/release/bioc/html/Biostrings.html), so I have implemented a version for natural language.

We can see how this works by aligning two documents. We might notice, for instance, that the documents `childshymnbook00caul` and `historyoflucyjac00amer` have a Jaccard score of `0.024`, low enough that we cannot conclude they are simple reprints, but high enough that the match isn't just noise or boilerplate (like the Google Books notice at the beginning of many documents). If we align the two documents we will get the longest borrowing in the document. The changes between the documents are also marked with `###`.

``` {.r}
align_local(ats[["childshymnbook00caul"]], ats[["historyoflucyjac00amer"]])
```

    ## TextReuse alignment
    ## Alignment score: 274 
    ## Document A:
    ## Early will I seek Thee # # # Now that my journey's just begun My road
    ## so little trod I'll come before I further run And give myself to God
    ## ######## ### ##### # What sorrows may my steps attend I never can
    ## foretell But if the Lord will be my friend I know that all is well
    ## 62 # THE CHILD'S 56 If all my earthly friends should die And leave me
    ## mourning here Since God can hear the orphan's cry O # what have I to
    ## fear # If I am poor he can supply Who has my table spread Who feeds the
    ## ravens when they cry And fills his poor with bread # If I am rich he'll
    ## guard my heart Temptation to withstand And make me willing to impart The
    ## bounties of his hand # But Lord whatever grief or ill For me may be in
    ## store Make me submissive to thy will #### And I would ask no more
    ## 
    ## Document B:
    ## EARLY WILL I SEEK THEE C M 1 NOW that my journey's just begun My road
    ## so little trod I'll come before I further run And give myself to God
    ## CHILDREN AND YOUTH 2 What sorrows may my steps attend I never can
    ## foretell But if the Lord will be my friend I know that all is well
    ## ## 3 ### ####### ## If all my earthly friends should die And leave me
    ## mourning here Since God can hear the orphan's cry # 0 what have I to
    ## fear 4 If I am poor he can supply Who has my table spread Who feeds the
    ## ravens when they cry And fills his poor with bread 5 If I am rich he'll
    ## guard my heart Temptation to withstand And make me willing to impart The
    ## bounties of his hand 6 But Lord whatever grief or ill For me may be in
    ## store Make me submissive to thy #### wilJ And I would ask no more

What we find is that *The History of Lucy Jackson* quotes a hymn from a children's hymnbook.[^12]

It is possible to use the Smith-Waterman algorithm to determine matches, as described by the *Viral Texts* project. I have found the minhash algorithm suitable for my purposes since I can more easily divide the documents in my corpus into sections. So I tend to use the local alignment to evaluate rather than detect the borrowings.

## Tracing borrowings within a corpus

A more interesting problem than de-duplicating a corpus is tracing the borrowings within it. Here the problem is not simply to cluster the documents in a corpus, but to figure out what those clusters mean and how documents have been modified, to see the borrowings changing over time, and to visualize the structure of the borrowings.

Let me give you a concrete example, using a project on which I am currently collaborating with Kellen Funk. Beginning in earnest in the 1850s, there was a movement in American law to codify civil procedure. Originating in two New York civil procedure codes of 1848 and 1850, named Field codes after their principal drafter [David Dudley Field](https://en.wikipedia.org/wiki/David_Dudley_Field_II), these codes were adopted and adapted by many U.S. states by the end of the nineteenth century.

To illustrate the process of borrowing, here are page images from New York's 1850 procedure code and from California's 1851 code. California has clearly lifted the section from New York, but it has also amended it. In many cases the amendments can be quite substantial.

New York's 1850 code:

{% figure field-code/ny1850-summons png "An excerpt from the 1850 New York code of civil procedure." %}

California's 1851 code:

{% figure field-code/ca1851-summons png "An excerpt from the 1851 California code of civil procedure." %}

The corpus we are working with currently contains 111 codes of civil procedure. While we could simply run the comparison functions on those documents, as we did for the ATS corpus above, that would reveal only a little about the structure of the borrowings. A far more useful approach is to split the civil procedure codes into their different sections and treat each of those sections as its own document. That gives us slightly more than 93,000 documents to deal with. Using the LSH functions in the textreuse package, it takes about an hour and half to detect a couple hundred thousand pairs of matches on my laptop. The end result is the same as with the ATS corpus: we get a data frame of matching pairs with their scores.

Full text goes in, similarity scores come out. The question, then, is what kinds of things can you learn from those similarity measures?

### Networks of borrowings

There are many things that you can do with that data. One obvious approach is to treat the borrowings as a network. The pair to pair measures of similarity provided by this package are the edge list for a network, so you can do anything with the textreuse package results that you can do with a network, whether that is visualization or more sophisticated ways of analyzing a network.

In an [earlier post](http://lincolnmullen.com/blog/detecting-text-reuse-in-legal-documents/) Kellen and I created a code-to-code network of borrowings within these codes of civil procedure. That network was based on comparing entire codes to entire codes. We have found that we get a much more accurate understanding of the borrowings by distilling the section-to-section borrowings down into code-to-code measures. This method matches how we know the code commissions actually worked when borrowing from other codes. I'm not going to post our current visualization since we are working on an article and don't want to give away everything prematurely, but you can poke around the [repository for our project](https://github.com/lmullen/civil-procedure-codes) if you're interested.

### Heat maps of borrowed sections within a document

We can also create heat maps of which sections within a code are borrowed from other codes. In the examples below, I've traced the development of some of the civil procedure codes in the West. The sections within a code are ordered from left to right, top to bottom. These visualizations are a work in progress: I need to refine the decision-making process for picking the best match for each section, and the colors will eventually have to be standardized across the visuals. But we can observe some interesting patterns.

California (1851) borrows most heavily from New York 1850, and occasionally from other New York codes. But we can also identify the sections which were lifted from the California code a year earlier.

{% figure field-code/CA1851 png "Borrowed sections in California (1851)." %}

Washington (1855) borrows large contiguous sections from Indiana and Oregon.

{% figure field-code/WA1855 png "Borrowed sections in Washington (1855)." %}

And Utah (1859) in turn borrows large, somewhat contiguous sections from the Oregon and Washington codes above.

{% figure field-code/UT1859 png "Borrowed sections in Utah (1859)." %}

These visualizations (once they are finished, anyway) are a tool for investigating which parts of each code were borrowed, as well as which parts of each code were an innovation. We can also ask what the contiguous borrowed sections were, and why contiguous sections were borrowed from different codes. Even in the absence of other historical evidence (which we can often use to corroborate our findings) we can deduce which codes were used by the various state code commissions.

### Clustering borrowings with affinity propagation

Another useful thing that we can do with the similarity measures is to cluster the borrowings. Above we converted that the set of pairs for the ATS corpus into a similarity matrix which we fed to a clustering algorithm. But with a corpus of this size there are simply too many matches to create a regular matrix: you'll run out of memory very quickly. Instead we can create a sparse matrix to represent the similarities compactly.[^13]

This means that we need to pick a clustering algorithm which can work with sparse similarity matrices. Fortunately, not only does [affinity propagation clustering](http://www.psi.toronto.edu/index.php?q=affinity%20propagation) work out of the box with a sparse matrix, it is also a clustering algorithm well suited to our particular problem. (See the [APCluster package](http://www.bioinf.jku.at/software/apcluster/) for an R implementation.) When solving a problem, you have to find an algorithm whose assumptions match the characteristics of your data. David Robinson explains this well in [his post on the "no free lunch theorem."](http://varianceexplained.org/r/kmeans-free-lunch/) Affinity propagation works by finding "exemplars among data points and forms clusters of data points around these exemplars." This is precisely the problem that we are trying to solve: there are exemplar documents which have been borrowed by other codes, around which clusters can be formed.[^14] I have found the affinity propagation to quickly produce clusters with virtually no false positives: I am very impressed by how well it did with the task.

In fact, the affinity propagation clustering was, if anything, too timid about putting together documents that should be together. But again, because the clustering algorithm is suitable to the problem that we are solving, there is a simple solution. Each cluster has an exemplar document, and we already have computed the similarities between those exemplar documents. I found it helpful to join clusters together that have exemplar documents above a threshold Jaccard similarity score. This produces the occasional false positive, but it also unites many clusters that ought to be together.

You can see what the clustering looks like in these text files. **[This cluster](/figures/field-code/NY1848-001560.txt)** has the cluster which contains the two example documents above serving a summons. You can see that it traces the spread of the section through 37 different codes, including some codes (e.g., NV1869) that more than double the length of the section by adding other material. The **[second cluster](/figures/field-code/NY1850-018680.txt)** has a variation of even greater interest. This section provides a list of people who are disqualified from being witnesses. The original section from New York (1850) disqualifies people of an "unsound mind" and children. California (1851) further disqualifies "Indians, or persons having one fourth or more of Indian blood, in an action or proceeding to which a white person is a party" and "Negroes, or persons having one half or more Negro blood, in an action or proceeding to which a white person is a party." While historians have long known that California introduced this disqualification on the basis of race, it is possible to trace here borrowings and variations and holdouts in other jurisdictions, and to set those borrowings in the context of the structural borrowings of codes as a whole.

The textreuse package thus allows you to deform the text, as Stephen Ramsay calls it in *[Reading Machines](http://www.press.uillinois.edu/books/catalog/75tms2pw9780252036415.html)*, by taking individual sections out of the context the particular codes from which they come.[^16] Stringing the related sections in chronological order makes the borrowings and changes far more legible. Those section-by-section borrowings can then be re-contextualized by setting them within the broader networks of borrowings within the codes.

## Conclusion

To sum up this over-long post: Detecting text reuse is a common problem in digital history and digital humanities. The basic concepts are quite simple at a small scale, but at more than a few hundred documents the geometric growth of the number of comparisons to be made requires more sophisticated methods. The textreuse package provides those methods, along with helpful classes for working easily with a corpus. Researchers working with "medium data" will find this package helpful. Finally, the package's similarity measures can be used to answer different kinds of questions which I have illustrated in this post. There are many different kinds of historical corpora which one could use with this package. Kellen and I may extend our analysis to other legal documents, such as treatises. And I am hoping to use this method on nineteenth-century hymnbooks if I can amass a sufficient corpus. If you have a use for the [textreuse package](https://github.com/ropensci/textreuse), or an idea for how it can be improved, I would be glad to hear about it.

[^1]: I have learned a great deal from the [methodological papers](http://viraltexts.org/publications-and-press/) published by the [Viral Texts](http://viraltexts.org/publications-and-press/) team: David A. Smith, Ryan Cordell, and Elizabeth Maddock Dillon, “Infectious Texts: Modeling Text Reuse in Nineteenth-Century Newspapers,” in *2013 IEEE International Conference on Big Data*, 2013, 86–94, doi:10.1109/BigData.2013.6691675; David A. Smith et al., “Detecting and Modeling Local Text Reuse,” in *Proceedings of the 14th ACM/IEEE-CS Joint Conference on Digital Libraries* (IEEE Press, 2014), 183–92, http://dl.acm.org/citation.cfm?id=2740800; Shaobin Xu et al., “Detecting and Evaluating Local Text Reuse in Social Networks,” *ACL 2014*, 2014, 50.

[^2]: A clarification: these are measures of similarity based on borrowed words rather than measures of similarity in content.

[^3]: Under the hood the `TextReuseTextDocument` class stores the tokens more efficiently as hashes---that is, integer representations of the tokens---rather than the tokens themselves.

[^4]: For documents about 10,000 words long on my computer, I get 735 comparisons per second.

[^5]: You could cut that time down by parallelizing the computation, but the problem isn't computing power but the naive method.

[^6]: Jure Leskovec, Anand Rajaraman, and Jeff Ullman, *[Mining of Massive Datasets](http://www.mmds.org/)* (Cambridge University Press, 2011), ch. 3.

[^7]: The savings can be considerable. For instance, if we take Project Gutenberg's version of *Moby Dick* (roughly 215,000 words), 7-grams of the text takes 20.1 MB of memory, but the integer hashes take only 859 kB.

[^8]: Note that any hash function has the possibility of a collision, that is, two different inputs may be mapped to the same output. In the 7-gram representations of *Moby Dick* using the `hash_string()` function, there are 3 collisions, or 0.0014 percent of the tokens. This has a minuscule effect compared to bad OCR or other problems with the text.

[^9]: I have again hashed the banded minhashes with a cryptographic hash: there are many different kinds of hashing in this package!

[^10]: Note that we did detect some pairs of documents with very low similarity scores: remember that the threshold is only probable, and we might want to filter these results before doing further work with them.

[^11]: I first learned of this algorithm from the *Viral Texts* papers cited above. Their implementation is more complex in that it uses an affine gap, while mine uses the simpler linear gap, and in that they prune the search tree. They are also working with differently structured data, so they use Smith-Waterman to identify matches, while I have found LSH to be sufficient for my purposes. I intend that future releases will make my implementation more full featured.

[^12]: Now there is a project idea: an analysis of borrowings and variations within hymnbooks, like Mark A. Noll, "'All Hail the Power of Jesus' Name': Significant Variations on a Significant Theme," in *Sing Them Over Again to Me: Hymns and Hymnbooks in America*, edited by Mark A. Noll and Edith L. Blumhofer (University of Alabama Press, 2006), 43--76, but bigger.

[^13]: The next release of textreuse will have a function to generate a sparse matrix automatically, but it is trivial to create one yourself.

[^14]: It is possible to provide a set of preferences to affinity propagation, identifying certain documents a priori to be more likely to be exemplars. In this particular problem, I could mark the New York codes as being more likely to be the exemplars. In practice, I have not found that to be necessary, though I have set the preferences globally to produce as few clusters as possible.

[^16]: See also Ryan Cordell's post on "[Scale as Deformance](http://ryancordell.org/research/scale-as-deformance/)."
