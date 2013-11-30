---
layout: post
title: "Aggregating Positions on Maps"
date: 2013-11-30 16:21
categories: 
---

How to aggregate positions on maps in R, or by extension, in any mapping 
place that uses data.

Leaflet.js does this, as the DPLA put to good effect. But this is for 
interactive, and 

A crude way to do this is to use alpha, but you're limited by whatever 
your denominator is, and the eye isn't good at picking these things up.

Sample data:


Use the summarize function ?? to count unique locations. This is a basic 
thing to do. Now we can scale the size of the points by locaiton.

But we might also want to do summarize by year and location. Now we can 
subset by year. How to do this for multiple years?

Now we want to aggregate by a variable, say communions or converts, so 
we can scale by that value instead. 
