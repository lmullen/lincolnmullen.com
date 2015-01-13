---
categories:
- rstats
date: '2014-11-15 09:26'
layout: post
title: An API Client for Omeka in R
link: https://github.com/lmullen/omekaR
...

This weekend I wrote [omekaR](https://github.com/lmullen/omekaR), an API client for [Omeka](http://omeka.org/) in [R](http://www.r-project.org/). (Yes, I had earlier written an [Omeka API client in Ruby](https://github.com/lmullen/omeka_client). But since the main reason that I would want to access an Omeka site is to do some kind of data analysis, R is a much more suitable language.) The client is very bare bones at the moment, but it can get any kind of resource, fetch all items, and extract metadata from items in a way that is idiomatic to R. I will likely finish the client when I write the [section on APIs](http://dh-r.lincolnmullen.com/apis.html) for *[Digital History Methods in R](http://dh-r.lincolnmullen.com/)*.
