---
title: 'Jane, John ... Leslie? A Historical Method for Algorithmic Gender Prediction'
date: "2016-01-06"
slug: 'jane-john-leslie-a-historical-method-for-algorithmic-gender-prediction'
description: "The gender package for R infers names' genders from historical data, accounting for change over time—and reveals a persistent gender gap in *AHR* book reviews."
tags:
- digital history
- R
draft: false
---

[Cameron Blevins](http://www.cameronblevins.org/) and I recently published an article in *Digital Humanities Quarterly* titled "[Jane, John ... Leslie? A Historical Method for Algorithmic Gender Prediction](http://www.digitalhumanities.org/dhq/vol/9/3/000223/000223.html)." The article has two related goals. First we explain the historical method behind the [gender package](https://github.com/ropensci/gender) for R, showing how it takes into account changes in the associations between names and genders. This method can be used by historians and other scholars to guess genders from first names as reliably as possible. Then, to show how the method can actually be used to make an argument, we apply the method to show that, while the number of history dissertations written by men and women is nearly equal, there continues to be a gap between the number of books on history reviewed in the *American Historical Review* written by men and women.

Here is the abstract:

> This article describes a new method for inferring the gender of personal names using large historical datasets. In contrast to existing methods of gender prediction that treat names as if they are timelessly associated with one gender, this method uses a historical approach that takes into account how naming practices change over time. It uses historical data to measure the likelihood that a name was associated with a particular gender based on the time or place under study. This approach generates more accurate results for sources that encompass changing periods of time, providing digital humanities scholars with a tool to estimate the gender of names across large textual collections. The article first describes the methodology as implemented in the gender package for the R programming language. It goes on to apply the method to a case study in which we examine gender and gatekeeping in the American historical profession over the past half-century. The [gender package](https://github.com/ropensci/gender) illustrates the importance of incorporating historical approaches into computer science and related fields.
