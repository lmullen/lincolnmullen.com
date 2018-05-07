---
title: "Digital history tenure dossier"
toc: true
---

This page collects the materials for my tenure dossier which are available only or primarily in digital format. As appropriate, the project and peer review status of each element is described. 

# About this dossier

This digital history dossier is a supplement to my written tenure statement. That statement aims to explain how my digital history scholarship fits into my broader research agenda. This dossier is not primarily intended to include all of the published articles and essays derived from my digital scholarship: for those, please see the supplement of additional publications. Rather this dossier aims to surface the elements of those projects that are born-digital, and to gather them in one place.

The dossier is organized by project. Underneath each project I have listed the digital elements of that project. 

I have gone to some lengths to have my digital scholarship peer reviewed whenever possible, even though born digital content is seldom peer reviewed in the same way that traditional, print-based scholarship is. To that end, under each project I have listed the form of the peer review that the scholarship has taken.

Because digital scholarship is often non-linear, especially when it takes the form of a database, visualization, or software, I have highlighted various pieces of these projects to guide the reader in exploring how those projects work. For instance, under the content-based projects I have included links to the various software repositories that enable this scholarship, while under the software packages I have included links to specific elements under the software documentation. Following these links should give a better sense of those elements of these various projects.

# America's Public Bible

**Project website**: <http://americaspublicbible.org/>

**Project status**: The prototype version is available online. As described in the tenure statement, a digital monograph with expanded source base, more visualizations, and six prose chapters combined with visualizations are in progress.

**Peer review:** The project's prototype and the proposal for the digital monograph were peer reviewed by two external readers for Stanford University Press. The project won [first place](https://www.neh.gov/news/press-release/2016-07-25) in the NEH's *Chronicling America* Data Challenge. Additionally the project has been [reviewed](https://doi.org/10.1093/jahist/jax556) in the *Journal of American History*.

**Other parts of the project**:

- The various code repositories for this project are collected as a [GitHub organization](https://github.com/public-bible). These include repositories for [training the models](https://github.com/public-bible/public-bible-models) that find the biblical quotations, for the [interactive visualization](https://github.com/public-bible/public-bible-verse-explorer) in the prototype, and for the [website](https://github.com/public-bible/public-bible-website) (which includes visualizations). 

**Project description**: *America’s Public Bible* uncovers the presence of biblical quotations in the nearly 11 million newspaper pages in the Library of Congress’s Chronicling America collection. Using the techniques of machine learning I have identified over 866,000 quotations of the Bible or verbal allusions to specific biblical verses on those newspaper pages. For now, the project has looked only for quotations to the King James Version (or Authorized Version) of the English Bible, by far the most commonly used Bible among American Protestants during the nineteenth century.6 For over 1,700 of the most frequently quoted verses, this site offers a way to explore the trend in how frequently a biblical verse was used, with links to each quotation highlighted in the pages of Chronicling America. The site thus uncovers two contexts for each verse: the context of the newspaper article in which it was used, and the broader chronological context of quotations from that verse and the Bible as a whole.

America’s Public Bible brings together two strands of scholarship. On the one hand, its methodology is drawn from recent digital humanities projects which are concerned with tracking the reuse of texts. On the other hand, it draws on a deep scholarly literature on the Bible as a cultural text in American life. The Bible included tens of thousands of texts, each of which could be interpreted in many ways. The contribution this site makes is to show how thousands of biblical verses were used over nearly a century in some 56 billion words of text, revealing trends that are inaccessible to a single scholar’s reading of these documents, yet enabling a close reading of the ways in which verses were put to use.

For many Americans the Bible was a text whose meaning was self-evident, yet the Bible’s role in U.S. history cannot be understood apart from the ways that Americans actually put it to use. That is why I have called this site America’s Public Bible. By looking at uses of the Bible in newspapers, we can see which parts of the Bible were in common currency among Americans, as well as the range of interpretations that were given to those verses. Verses that could be cited without a reference (or used in jokes) indicated a kind of literacy or familiarity, and possibly a shared assumption about what those verses could be interpreted to mean. Verses that were used constantly were a shared cultural touchstone, while verses that were used only episodically reveal the tensions in a particular political or social situation. By looking at how the verses were actually used, we can see how the Bible was a contested yet common text.

# Mapping Early American Elections

# Spine of American Law

# Mapping American Religious Ecologies

# R packages

## tokenizers: Fast, consistent tokenization of natural language text

**Status:** stable, but actively maintained

**Peer review:** [peer reviewed by rOpenSci](https://github.com/ropensci/onboarding/issues/33)

**Published paper**: This package has been described in the [following paper](https://doi.org/10.21105/joss.00655):

Lincoln A. Mullen et al., "Fast, Consistent Tokenization of Natural Language Text," *Journal of Open Source Software* 3, no. 23 (2018): 655, <https://doi.org/10.21105/joss.00655>.

**Package documentation:**

- The "[Introduction to the tokenizers Package](https://lincolnmullen.com/software/tokenizers/articles/introduction-to-tokenizers.html)" gives an overview of how to use the package.
- The "[The Text Interchange Formats and the tokenizers Package](https://lincolnmullen.com/software/tokenizers/articles/tif-and-tokenizers.html)" vignette explains how this package fits into the R text analysis ecosystem.
- The [package website](https://lincolnmullen.com/software/tokenizers/) describes the package and documents of all its functionality.

**Package description:** This R package offers functions with a consistent interface to convert natural language text into tokens. It includes tokenizers for shingled n-grams, skip n-grams, words, word stems, sentences, paragraphs, characters, shingled characters, lines, tweets, Penn Treebank, and regular expressions, as well as functions for counting characters, words, and sentences, and a function for splitting longer texts into separate documents, each with the same number of words. The package is built on the stringi and Rcpp packages for fast yet correct tokenization in UTF-8.

This package complies with the standards for input and output recommended by the Text Interchange Formats. The TIF initiative was created at an rOpenSci meeting in 2017, and its recommendations are available as part of the tif package. See the “The Text Interchange Formats and the tokenizers Package” vignette for an explanation of how this package fits into that ecosystem.

## textreuse: Detect text reuse and document similarity

**Status:** in active development

**Peer review:** [peer reviewed by rOpenSci](https://github.com/ropensci/onboarding/issues/20)

**Package documentation:**

- The "[Introduction to the textreuse Package](https://lincolnmullen.com/software/tokenizers/articles/introduction-to-tokenizers.html)" gives an overview of how to use the package.
- The "[Minhash and locality-sensitive hashing](https://lincolnmullen.com/software/textreuse/articles/textreuse-minhash.html)" vignette explains how to measure the similarities between documents.
- The "[Pairwise comparisons for document similarity](https://lincolnmullen.com/software/textreuse/articles/textreuse-pairwise.html)" vignette explains how to apply similarities to pairs of documents.
- The "[Text Alignment](https://lincolnmullen.com/software/textreuse/articles/textreuse-alignment.html)" vignette explains how to fine the best subset of each document that aligns with one another.
- The [package website](https://lincolnmullen.com/software/textreuse/) describes the package and documents of all its functionality.

**Package description:** This R package provides a set of functions for measuring similarity among documents and detecting passages which have been reused. It implements shingled n-gram, skip n-gram, and other tokenizers; similarity/dissimilarity functions; pairwise comparisons; minhash and locality sensitive hashing algorithms; and a version of the Smith-Waterman local alignment algorithm suitable for natural language. It is broadly useful for, for example, detecting duplicate documents in a corpus prior to text analysis, or for identifying borrowed passages between texts. The classes provides by this package follow the model of other natural language processing packages for R, especially the NLP and tm packages. (However, this package has no dependency on Java, which should make it easier to install.)

## USAboundaries: Historical and contemporary boundaries of the United States of America

**Status:** stable, but actively maintained

**Peer review:** [peer reviewed by the *Journal of Open Source Software*](https://github.com/openjournals/joss-reviews/issues/314)

**Published paper**: This package has been described in the [following paper](https://doi.org/10.21105/joss.00314):

Lincoln A. Mullen and Jordan Bratt, "USAboundaries: Historical and Contemporary Boundaries of the United States of America," *Journal of Open Source Software* 3, no. 23 (2018): 314, <https://doi.org/10.21105/joss.00314>.

**Package documentation:**

- The “[A sample analysis using USAboundaries](https://lincolnmullen.com/software/usaboundaries/articles/usaboundaries-sample-analysis.html)” vignette gives example of how the package can be used for both historical and contemporary maps.
- The [package website](https://lincolnmullen.com/software/usaboundaries/) describes the package and documents of all its functionality.

**Package description:** This R package includes contemporary state, county, and Congressional district boundaries, as well as zip code tabulation area centroids. It also includes historical boundaries from 1629 to 2000 for states and counties from the Newberry Library’s Atlas of Historical County Boundaries, as well as historical city population data from Erik Steiner’s “United States Historical City Populations, 1790-2010.” The package has some helper data, including a table of state names, abbreviations, and FIPS codes, and functions and data to get State Plane Coordinate System projections as EPSG codes or PROJ.4 strings.

This package can serve a number of purposes. The spatial data can be joined to any other kind of data in order to make thematic maps. Unlike other R packages, this package also contains historical data for use in analyses of the recent or more distant past. 

## gender: Predict gender from names using historical data

**Status:** stable, but actively maintained

**Peer review:** An article describing this method has been peer reviewed by *Digital Humanities Quarterly*.

**Published paper**: The method behind this software has been described in the [following article](http://www.digitalhumanities.org/dhq/vol/9/3/000223/000223.html):

Cameron Blevins and Lincoln A. Mullen, “Jane, John … Leslie? A Historical Method for Algorithmic Gender Prediction,” *Digital Humanities Quarterly* 9, no. 3 (2015): <http://www.digitalhumanities.org/dhq/vol/9/3/000223/000223.html>.

**Package documentation:**

- The “[Predicting Gender Using Historical Data](https://lincolnmullen.com/software/gender/articles/predicting-gender.html)” vignette explains how to use the package.
- The [package website](https://lincolnmullen.com/software/gender/) describes the package and documents of all its functionality.

**Package description:** Data sets, historical or otherwise, often contain a list of first names but seldom identify those names by gender. Most techniques for finding gender programmatically rely on lists of male and female names. However, the gender associated with names can vary over time. Any data set that covers the normal span of a human life will require a historical method to find gender from names. This R package uses historical datasets from the U.S. Social Security Administration, the U.S. Census Bureau (via IPUMS USA), and the North Atlantic Population Project to provide predictions of gender for first names for particular countries and time periods.

## historydata: Datasets for historians

**Status:** in active development

**Package documentation:**

- The [package website](https://lincolnmullen.com/software/historydata/) describes the package and documents of all its included datasets.

**Package description:** These sample data sets are intended for historians learning R. They include population, institutional, religious, military, and prosopographical data suitable for mapping, quantitative analysis, and network analysis.

## internetarchive: An R client for the Internet Archive API

**Status:** stable, but actively maintained

**Package documentation:**

- The "[Internet Archive API Client](https://lincolnmullen.com/software/internetarchive/articles/internet-archive.html)" gives an overview of how to use the package to download metadata and files from the Internet Archive.
- The [package website](https://lincolnmullen.com/software/internetarchive/) describes the package and documents of all its included datasets.

**Package description:** This API client for the Internet Archive is intended primarily for searching for items, retrieving metadata for items, and downloading the files associated with items. The functions can be used with the pipe operator (%>%) from magrittr and the data manipulation verbs in dplyr to create pipelines from searching to downloading. For the full details of what is possible with the Internet Archive API, see their advanced search help.
