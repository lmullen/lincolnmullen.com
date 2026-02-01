---
title: 'New Release of the gender Package (v0.5.0)'
date: "2015-08-24"
slug: 'new-release-of-the-gender-package-v0-5-0'
tags:
- R
- digital history
draft: false
---

I have released a new version (v0.5.0) of the [gender package](http://www.r-pkg.org/pkg/gender) for the [R programming language](https://www.r-project.org/). The gender package predicts gender from first names and dates using historical datasets. You can install the package from [CRAN](https://cran.rstudio.com/web/packages/gender/) or see the [code at GitHub](https://github.com/ropensci/gender).

This release brings a number of improvements. First, there are significant performance improvements, which make the package more useful for anyone using it for large datasets. I have also simplified the package so it always returns data frames. This was my first R package that I published on CRAN, and let's just say that I have since found a lot of low-hanging fruit when it came to performance and usability.

<!--more-->

Second, I have added a dataset from the [North Atlantic Population Project](https://www.nappdata.org/napp/) which provides a dataset of names for Canada, Great Britain, Germany, Iceland, Norway, Sweden for the nineteenth century. This extends the package's usefulness beyond its original focus on American history. (If you have suitable datasets for other times and places, I'd welcome contributions since the gender package can be easily extended.)

Third, I have added a new function `gender_df()` which makes it easier to use gender with a common research problem. The `gender()` function is vectorized on names but not on dates. In other words, it is easy to pass `gender()` many names, but not many dates. Suppose, for example, that we have a list of names and wish to guess their genders for birth years in the 1930s. We can do that like this:

    gender(c("Susan", "Madison", "Hillary"), years = c(1930, 1939), method = "ssa")

    ## Source: local data frame [3 x 6]
    ##
    ##      name proportion_male proportion_female gender year_min year_max
    ## 1 Hillary          1.0000            0.0000   male     1930     1939
    ## 2 Madison          1.0000            0.0000   male     1930     1939
    ## 3   Susan          0.0041            0.9959 female     1930     1939

But suppose, as is commonly the case, that you have a dataset where there are birth years (or a guess at a range of birth years) associated with each name, as in this example.

    demo_df

    ## Source: local data frame [7 x 3]
    ##
    ##   first_names last_names years
    ## 1       Susan          A  1930
    ## 2       Susan          B  2000
    ## 3     Madison          C  1930
    ## 4     Madison          D  2000
    ## 5     Hillary          E  1930
    ## 6     Hillary          F  2000
    ## 7     Hillary          G  1930

While it has always been possible to use the `gender()` function with `Map()` or `dplyr::do()`, those are not easy to use choices, and most code that I have seen (especially my own at first) has had a naive approach that calls the `gender()` function many more times than necessary. So the `gender_df()` function lets you pass it a data frame with a column of first names and a column of years (or two columns specifying the minimum and maximum of a range of years).

    gender_df(demo_df, name_col = "first_names", year_col = "years", method = "ssa")

    ##      name proportion_male proportion_female gender year_min year_max
    ## 1 Hillary          1.0000            0.0000   male     1930     1930
    ## 2 Madison          1.0000            0.0000   male     1930     1930
    ## 3   Susan          0.0000            1.0000 female     1930     1930
    ## 4 Hillary          0.0000            1.0000 female     2000     2000
    ## 5 Madison          0.0069            0.9931 female     2000     2000
    ## 6   Susan          0.0000            1.0000 female     2000     2000

Besides simplifying the code, this function does the least amount of work possible, and should speed up gender prediction for large datasets considerably.

There are more details on how to use the package in [the package vignette](https://cran.rstudio.com/web/packages/gender/vignettes/predicting-gender.html). And [Cameron Blevins](http://www.cameronblevins.org/) and I will soon publish an article justifying the historical method and applying it to a case study.
