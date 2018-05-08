---
title: "Digital history dossier"
toc: true
---

# Introductory statement

This digital history dossier is a supplement to my tenure statement. That statement explains how my scholarship in traditional history and digital history fit into a broader research agenda. This online dossier aims to describe my digital history work which is born-digital. It features work which can only be presented fully on the web, or in the form of software or datasets.
This online dossier is not intended to include published articles and essays about digital history: for those, please see the appendix of additional publications.

The work featured in this dossier falls into two main categories: first, my work contributes to and enables historical interpretations using digital history methods; second, it implements those methods in the form of software and its documentation. My digital history work has focused on the area I call computational history, and specifically on digital mapping and computational text analysis.

My interpretative digital scholarship includes work in American religious history, early American political history, and American legal history. *America's Public Bible* is a monograph-length project that also includes interactive visualizations. The prize-winning prototype for this project is currently available as I develop the rest of the site for Stanford University Press. *Mapping Early American Elections* is in the process of providing the most detailed maps available for Congressional elections through the nineteenth Congress, as well as for select state legislative elections. "The Spine of American Law," is primarily a journal article based on text analysis, but our supplementary appendix provides materials for other legal historians to build on the base of our scholarship. "The Spread of American Slavery" is an interactive visualization of the populations of free and enslaved African Americans in the decennial censuses up to 1860. This map has had a very wide public impact. And *Mapping American Religious Ecologies* is my next major collaborative research project, currently in the planning and development stages.

These interpretative projects share several common features. First, they apply digital methods to make interpretative arguments. As the collaboratively written [white paper on digital history and argument](https://rrchnm.org/argument-white-paper/) that I co-organized points out, such practices are rare among digital historians. Second, my interpretative digital history scholarship centers around the nineteenth-century United States. This focus has allowed me to build up a body of scholarship within that chronological and geographic field of study. And third, whenever possible I have released data and code alongside my interpretative scholarship. My open-source ethos is motivated by a desire to allow other scholars to engage with and test my interpretations and the raw materials which make those interpretations possible, but also to enable other scholars to create their own interpretations from the datasets that I've created. For instance, *Mapping Early American Elections* provides datasets of electoral returns which are ready to be mapped, and the project will soon include tutorials on how to use those datasets. "The Spine of American Law" includes both our corpus of nineteenth-century codes of civil procedure, which allows digital historians to perform their own computations, and our intermediate clusters of sections of the law, which allows traditional historians to perform a close reading of changes in the law.

My methodological work which is born-digital has primarily taken the form of software packages for the R programming language. These software packages include extensive documentation to explain how they work and for what purposes they can be used. These packages are published on the [Comprehensive R Archive Network](https://cran.rstudio.com/) (CRAN), I have also published articles or papers about the software packages as appropriate. My R packages contribute to mapping and especially to text analysis, in addition to teaching R to historians. Current counts of the downloads of each package are available on the package websites (though these counts come only from the most popular of the CRAN mirrors, and leave out downloads from the dozens of other CRAN mirrors).

This dossier is organized by project. Underneath each project I have listed the project's current status and provided a project description explaining its contributions. More details about the projects can be found, of course, by visiting them at the URLs provided. In the case of software, I have provided a website with all the package documentations and examples.

I have gone to some lengths to have my digital scholarship peer reviewed whenever possible, even though peer review for born-digital content is seldom available in the same way that it is for traditional, print-based scholarship. To that end, under each project I have listed the form of the peer review that the scholarship has taken when applicable.

Because digital scholarship often has multiple elements, especially when it takes the form of a database, visualization, or software, I have highlighted various pieces of these projects to guide the reader in exploring how these projects work. For instance, under the content-based projects I have included links to the various software repositories that enable this scholarship, while under the software packages I have included links to specific elements under the software documentation. In addition to following the main link for each projects, these more detailed links should draw your attention to key elements of the project.

# America's Public Bible

**Project website**: <http://americaspublicbible.org/>

**Project status**: A prototype version of this site is available online. A digital monograph with an expanded source base, more visualizations, and six prose chapters combined with visualizations are in progress.

**Peer review:** The project's prototype and the proposal for the digital monograph were peer reviewed by two external readers for Stanford University Press. The project won [first place](https://www.neh.gov/news/press-release/2016-07-25) in the NEH's *Chronicling America* Data Challenge. Additionally the project has been [reviewed](https://doi.org/10.1093/jahist/jax556) in the *Journal of American History*.

**Other parts of the project**:

- The various code repositories for this project are collected as a [GitHub organization](https://github.com/public-bible). These include repositories for [training the models](https://github.com/public-bible/public-bible-models) that find the biblical quotations, for the [interactive visualization](https://github.com/public-bible/public-bible-verse-explorer) in the prototype, and for the [website](https://github.com/public-bible/public-bible-website) (which includes visualizations). 

**Project description**: *America’s Public Bible* uncovers the presence of biblical quotations in millions newspaper pages in the Library of Congress's *Chronicling America* collection and Gale's *Nineteenth-century Newspapers* dataset. Using the techniques of machine learning, I have identified quotations or verbal allusions to specific biblical verses on those newspaper pages. For now, the project has looked only for quotations to the King James Version (or Authorized Version) of the English Bible, by far the most commonly used Bible among American Protestants during the nineteenth century, but it will be expanded to other versions of the Bible for the digital monograph. For over 1,700 of the most frequently quoted verses, this site offers a way to explore the trend in how frequently a biblical verse was used, with links to each quotation highlighted in the pages of *Chronicling America*. The site thus uncovers two contexts for each verse: the context of the newspaper article in which it was used, and the broader chronological context of quotations from that verse and the Bible as a whole.

America's Public Bible brings together two strands of scholarship. On the one hand, its methodology is drawn from recent digital humanities projects which are concerned with tracking the reuse of texts. On the other hand, it draws on a deep scholarly literature on the Bible as a cultural text in American life. The Bible included tens of thousands of texts, each of which could be interpreted in many ways. The contribution this site makes is to show how thousands of biblical verses were used over nearly a century in well over 60 billion words of text, revealing trends that are inaccessible to a single scholar's reading of these documents, yet enabling a close reading of the ways in which verses were put to use.

For many Americans the Bible was a text whose meaning was self-evident, yet the Bible's role in U.S. history cannot be understood apart from the ways that Americans actually put it to use. That is why I have called this site *America's Public Bible*. By looking at uses of the Bible in newspapers, we can see which parts of the Bible were in common currency among Americans, as well as the range of interpretations that were given to those verses. Verses that could be cited without a reference (or used in jokes) indicated a kind of literacy or familiarity, and possibly a shared assumption about what those verses could be interpreted to mean. Verses that were used constantly were a shared cultural touchstone, while verses that were used only episodically reveal the tensions in a particular political or social situation. By looking at how the verses were actually used, we can see how the Bible was a contested yet common text.

# Mapping Early American Elections

**Project website**: <http://earlyamericanelections.org/>

**Project status**: *Mapping Early American Elections* began in May 2016, and will be funded through May 2019. The first two years have focused primarily on the data transformation, since it has required a great deal of effort to make it suitable for mapping. We have published 157 maps for the first ten Congresses, and will complete maps for the remaining nine Congresses in our period by the end of summer 2018. We are also in the process of editing six interpretative essays by leading scholars in early American political history, using the data and maps we have created. I will contribute an essay on party coherence over time and space, using the spatial data that we have created to show in which counties parties tended to remain dominant for long stretches of time, and where races tended to be close. This is a form of spatial analysis which, combined with the maps, enables a spatial interpretation of the main trends in early American politics.

**Peer review**: Like all NEH grants, the funding process for *Mapping Early American Elections* entailed a rigorous review by a panel of experts at the National Endowment for the Humanities. This project was funded in the Division of Preservation and Access's program for [Humanities Collections and Reference Resources](https://www.neh.gov/grants/preservation/humanities-collections-and-reference-resources).

**Project description**: *Mapping Early American Elections* is based on *The New Nation Votes* dataset, the only comprehensive record of elections in the early American republic, gathered over 45 years of research by Philip J. Lampi of the American Antiquarian Society. *Mapping Early American Elections* takes that dataset, adds a spatial element to it, fills in missing party affiliations and vote percentages, and maps the results. The *New Nation Votes* data include between ten and fifteen thousand place names, but those place names are not associated with actual spatial data that can be used in mapping software. To map the data, a historian must connect the place names in the dataset to the locations of cities and towns and factor in the constantly changing boundaries of counties and congressional districts. This process requires both technical knowledge of spatial and quantitative data and historical research into elections and boundaries. 

*Mapping Early American Elections* will offer expanded access to the early American election returns in the form of spatial data, maps, and tutorials that will allow researchers to generate their own maps. First, we are extending the *New Nation Votes* dataset by connecting it to existing spatial datasets of historical town, county, and district boundaries. The new dataset will be usable in many mapping programs. We will provide tutorials that teach users, including journalists, educators, and researchers, how to map the data for themselves. Second, we are creating a series of interactive maps of congressional and state legislative elections. These maps show changes over time and variation by place in party preference and other variables. These maps will be available at different geographic scales, from counties and towns in individual states to national maps. The project team and recognized experts in the field of early American history will write a series of essays that frame the historical context for understanding the maps and explain what can be learned from the various maps.

The data collected offer a window into the formative era of American politics. The era from the federal Constitutional Convention to the election of John Quincy Adams raised for the first time the perennial questions of American political life: Who could vote, and at what rates would they turn out at the polls? Which parties and candidates would they support? How would states and regions vote differently? Scholars have already used this dataset to transform our understanding of early American political history. They have not yet, however, used the dataset to create maps that visually represent the spatial dimensions of electoral process. *Mapping Early American Elections* will make historical election data accessible to broader audiences through a hands-on interrogation of sources. By providing maps and tutorials, researchers will be able to actively trace the evolution of American democracy. By encouraging users to ask their own questions about past elections, we hope to encourage citizens to become more engaged in current-day elections in their own towns, counties, and the nation. 

# Spine of American Law

**Digital components:** Our article "[The Spine of American Law: Digital Text Analysis and U.S. Legal Practice](https://academic.oup.com/ahr/article/123/1/132/4840258?guestAccessKey=3a271895-4af8-41e8-b3d9-15941dfe0bcf)" in the *American Historical Review* was primarily cast in the format of a journal article. (Please note, though, that the online PDF or HTML version, not the print version, is the version that best reflects our scholarship, since the visualizations that I created are presented online in full color.) 

But the digital elements that we created were essential to that article. They are available on the journal's website as a research compendium. This appendix includes the two R packages created for this article, “textreuse” and “tokenizers,” as well as the repository of code used to analyze the data and create the visualizations. Two versions of our dataset are available: one containing the OCR text of each code or statute in a separate text file, and the second containing those codes and statutes split into sections. A data table provides a list of all the matches between sections that we found, and another data table contains what we determined to be the best match for each section of the code. A set of text files provides the clusters of related sections. 

This research compendium serves two important purposes. First, it allows other scholars to validate our results. But second, it also allows other scholars to extend our results by using the corpus we gathered for themselves, or by examining the results in the clusters we provide to answer different questions than the ones that we foregrounded in our article.

# Mapping American Religious Ecologies

**Project status:** Mapping American Religious Ecologies is in the early stages of development, as John Turner and I use seed funding from the GMU College of Humanities and Social Sciences to inventory and digitize the federal 1926 Census of Religious Bodies. We are currently writing a grant to the NEH which, if funded, will begin in May 2019, and we have also submitted a letter of inquiry to a private funder. 

**Groundwork for this project:** I have been exploring what can be done with data about American religion before the 1940s for some time. My clearest programmatic statement, which includes sample visualizations of Congregationalist data that I gathered, is in a conference paper for the American Society of Church History titled "[Mapping Boston’s Religions: Next Steps in Mapping U.S. Religious History](https://lincolnmullen.com/projects/asch-2015/)." I have created maps of [Paulist missions](https://lincolnmullen.com/projects/paulists-map/), as well as a map of Methodist data for a presentation at the Society for Historians of the Early American Republic. Some of the data that I have gathered is publicly available in a [GitHub repository](https://github.com/lmullen/religious-statistics).

**Project description:** *Mapping American Religious Ecologies* is an interdisciplinary investigation into the ecology of American religious life that seeks to create, map, and interpret more detailed datasets of interactions between religious groups than have previously been available to scholars. The project will demonstrate---both visually and analytically---the ecology of American religion: how religious groups have lived alongside or avoided each, how they have been shaped by changing patterns of housing and transportation, and how Americans have created religious places out of urban space. 

This project aims to create three modules: (1) maps created with city- and congregational-level data (rather than county-level data) from the 1926 Census of Religious Bodies; (2) maps of the increasing diversity of American religion following immigration reform in 1965, with a particular focus on Muslim, Hindu, Buddhist, and Ethiopian and other Orthodox Christian groups; (3) maps utilizing previously unused denominational records from the nineteenth century. 

Scholars of American religion have often remarked on the significance of space and geography to their field. The distribution of different religious groups across the United States is highly uneven, due to both immigration and internal migration. Geographic considerations shaped where American religious groups settled, perhaps most famously in the migration of the Latter-day Saints (Mormons) to the Great Basin, but also in the concentration of both Catholics and Jews in urban areas, and in the movement of white Protestants to suburbs. Through their religious affiliations, moreover, people frequently turn geographic space into sacred and contested space. 

Yet for all of the obvious importance of space, scholars of American religion have made do with very limited quantitative data. Concerns over the implications for religious liberty meant that the federal government did not include any religious questions on the Census until 1850, and stopped asking religious questions after 1946. Even the questions that the Census did ask were fairly limited, and the published reports of the Census were always aggregated by state or county. The result is that most maps that have been made of American religion prior to the late twentieth century are dependent on county-level aggregates of the number of congregations. These maps present only a blurry image of American religion.

In a sharply worded review of several atlases of American religion titled “If It’s South Dakota You Must Be Episcopalian,” Laurie Maffly-Kipp critiques the representations of county-by-county counts of churches as undue simplifications of the complexity of religious experience. “It would be quite ironic,” she writes, “if our dazzling new technologies and spatial theories only brought us back to much more circumscribed definitions of religious experience.” 

Our project will build upon but surpass previous maps of the American religious landscape. We will begin doing so by using previously untapped sources that are more detailed, starting with the federal 1926 Census of Religious Bodies. The 1926 Census of Religious Bodies is a uniquely detailed source for American religion. The Census Bureau managed to contact virtually every American congregation in that year, and it is the only religious census for which the records survive. Such data will allow us to reconstruct the distribution and interactions between American religious groups, allow us to visually represent how the American religious experience has changed over time. 

# The Spread of American Slavery

**Project description**: "[The Spread of American Slavery](https://lincolnmullen.com/projects/slavery/)" is an interactive visualization that I created initially for my classes, but which has come to reach a far wider audience. The map allows users to examine various measures for the enslaved and free populations of African Americans in the decennial censuses.

The map has had a very wide public impact. I initially wrote about the map for [*Smithsonian* magazine](https://www.smithsonianmag.com/history/maps-reveal-slavery-expanded-across-united-states-180951452/?no-ist), and the map has also been written about in media outlets such as [Slate](http://www.slate.com/articles/life/the_history_of_american_slavery/2015/06/the_demographics_of_american_slavery_according_to_new_research_by_eric_foner.html), the [*Washington Post*](https://www.washingtonpost.com/opinions/how-to-learn-more-about-slavery-after-watching-the-new-roots/2016/06/03/c7b02b98-290d-11e6-a3c4-0724e8e24f3f_story.html?noredirect=on&utm_term=.23ba1ec16c49), and [*Vox*](https://www.vox.com/2016/3/26/11306912/slavery-map), as well as less probable venues such as the [*Daily Mail*](http://www.dailymail.co.uk/news/article-2633230/The-slaves-helped-build-nation-Maps-reveal-relentless-westward-spread-human-bondage-early-America.html) and [*PC Magazine*](https://www.pcmag.com/article2/0,2817,2491377,00.asp), and on the blogs of publicly-engaged scholars such as [Susan Schulten](http://www.mappingthenation.com/blog/an-interactive-map-of-slavery/) and [John Fea](https://thewayofimprovement.com/2014/05/16/the-spread-of-u-s-slavery-1790-1860/). The map has also been used in [elementary](https://www.tolerance.org/sites/default/files/2018-02/TT-Teaching-Hard-History-American-Slavery-IDM-How-did-Slavery-Shape-Feb2018.pdf) and [high school] (http://blogs.henrico.k12.va.us/hcps-ushst-resources/2014/05/16/the-spread-of-u-s-slavery-infographic/) curriculum and college syllabi. 

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

This package complies with the standards for input and output recommended by the [Text Interchange Formats](https://github.com/ropensci/tif). I helped create the TIF initiative at an rOpenSci meeting in 2017, and its recommendations are available as part of the tif package. See the “The Text Interchange Formats and the tokenizers Package” vignette for an explanation of how this package fits into that ecosystem.

## textreuse: Detect text reuse and document similarity

**Status:** in active development

**Peer review:** [peer reviewed by rOpenSci](https://github.com/ropensci/onboarding/issues/20)

**Package documentation:**

- The "[Introduction to the textreuse Package](https://lincolnmullen.com/software/tokenizers/articles/introduction-to-tokenizers.html)" gives an overview of how to use the package.
- The "[Minhash and locality-sensitive hashing](https://lincolnmullen.com/software/textreuse/articles/textreuse-minhash.html)" vignette explains how to measure the similarities between documents.
- The "[Pairwise comparisons for document similarity](https://lincolnmullen.com/software/textreuse/articles/textreuse-pairwise.html)" vignette explains how to apply similarities to pairs of documents.
- The "[Text Alignment](https://lincolnmullen.com/software/textreuse/articles/textreuse-alignment.html)" vignette explains how to fine the best subset of each document that aligns with one another.
- The [package website](https://lincolnmullen.com/software/textreuse/) describes the package and documents of all its functionality.

**Package description:** This R package provides a set of functions for measuring similarity among documents and detecting passages which have been reused. It implements shingled n-gram, skip n-gram, and other tokenizers; similarity/dissimilarity functions; pairwise comparisons; minhash and locality sensitive hashing algorithms; and a version of the Smith-Waterman local alignment algorithm suitable for natural language. It is broadly useful for, for example, detecting duplicate documents in a corpus prior to text analysis, or for identifying borrowed passages between texts.

## USAboundaries: Historical and contemporary boundaries of the United States of America

**Status:** stable, but actively maintained

**Peer review:** [peer reviewed by the *Journal of Open Source Software*](https://github.com/openjournals/joss-reviews/issues/314)

**Published paper**: This package has been described in the [following paper](https://doi.org/10.21105/joss.00314):

Lincoln A. Mullen and Jordan Bratt, "USAboundaries: Historical and Contemporary Boundaries of the United States of America," *Journal of Open Source Software* 3, no. 23 (2018): 314, <https://doi.org/10.21105/joss.00314>.

**Package documentation:**

- The “[A sample analysis using USAboundaries](https://lincolnmullen.com/software/usaboundaries/articles/usaboundaries-sample-analysis.html)” vignette gives example of how the package can be used for both historical and contemporary maps.
- The [package website](https://lincolnmullen.com/software/usaboundaries/) describes the package and documents of all its functionality.

**Package description:** This R package includes contemporary state, county, and Congressional district boundaries, as well as zip code tabulation area centroids. It also includes historical boundaries from 1629 to 2000 for states and counties from the Newberry Library’s *Atlas of Historical County Boundaries*, as well as historical city population data from Erik Steiner’s “United States Historical City Populations, 1790-2010.” The package has some helper data, including a table of state names, abbreviations, and FIPS codes, and functions and data to get State Plane Coordinate System projections as EPSG codes or PROJ.4 strings.

This package can serve a number of purposes. The spatial data can be joined to any other kind of data in order to make thematic maps. Unlike other R packages, this package also contains historical data for use in analyses of the recent or more distant past. 

## gender: Predict gender from names using historical data

**Status:** stable, but actively maintained

**Peer review:** An article describing this method has been peer reviewed by *Digital Humanities Quarterly*.

**Published paper**: The method behind this software has been described in the [following article](http://www.digitalhumanities.org/dhq/vol/9/3/000223/000223.html):

Cameron Blevins and Lincoln A. Mullen, “Jane, John … Leslie? A Historical Method for Algorithmic Gender Prediction,” *Digital Humanities Quarterly* 9, no. 3 (2015): <http://www.digitalhumanities.org/dhq/vol/9/3/000223/000223.html>.

**Package documentation:**

- The “[Predicting Gender Using Historical Data](https://lincolnmullen.com/software/gender/articles/predicting-gender.html)” vignette explains how to use the package.
- The [package website](https://lincolnmullen.com/software/gender/) describes the package and documents of all its functionality.

**Package description:** Datasets, historical or otherwise, often contain a list of first names but seldom identify those names by gender. Most techniques for finding gender programmatically rely on lists of male and female names. However, the gender associated with names can vary over time. Any data set that covers the normal span of a human life will require a historical method to find gender from names. This R package uses historical datasets from the U.S. Social Security Administration, the U.S. Census Bureau (via IPUMS USA), and the North Atlantic Population Project to provide predictions of gender for first names for particular countries and time periods.

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

**Package description:** This API client for the Internet Archive is intended primarily for searching for items, retrieving metadata for items, and downloading the files associated with items. The functions can be used with the pipe operator (`%>%`) from magrittr and the data manipulation verbs in dplyr to create pipelines from searching to downloading. For the full details of what is possible with the Internet Archive API, see their advanced search help.
