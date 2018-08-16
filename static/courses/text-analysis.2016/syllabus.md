---
bibliography: 'text-analysis.bib'
csl: 'chicago-syllabus.csl'
css:
- 'tufte.css'
- 'custom.css'
link-citations: False
nocite: '@ramsay_reading_2011, @james_introduction_2013, @dalgaard_introductory_2008, @shotts_linux_2016, @kuhn_applied_2013, @xu_detecting_2014, @moretti_graphs_2005, @moretti_distant_2013, @archer_whats_2009, @silge_tidy_2016, @michel_quantitative_2011, @wickham_r_2016, @wickham_advanced_2014, @goldstone_quiet_2014, @underwood_theorizing_2014, @underwood_literary_2015, @blei_probablistic_2012, @milligan_automated_2012, @knox_understanding_2013, @robertson_signs_1998, @jockers_macroanalysis:_2013], @binder_alien_2016, @bryan_happy_2016'
smart: True
title: Text Analysis for Historians
type: index
---


[Schedule](#schedule) | [Assignments](#assignments-and-expectations) | [Fine Print](#acknowledgments-and-fine-print) | [References](#references)


<p class="sans">
Independent study, fall 2016. [Department of History and Art History](http://historyarthistory.gmu.edu/), George Mason University. Meets every other Wednesday at 1 p.m. for discussion; work sessions on alternate weeks as necessary. Instructor: [Lincoln Mullen](http://lincolnmullen.com) &lt;<lmullen@gmu.edu>&gt;. Office: Research Hall 457.
</p>

This independent study is an advanced course in the theory and practice of text analysis for historians. You will read current research in digital history and cognate fields. The aim is to learn the methods of text analysis which are most likely to produce insights useful for historical interpretation. You will work primarily in the [R programming language](https://www.r-project.org/) and with common Unix-style utilities, performing an analysis using each method we discuss week by week. By the end of the semester, you will write a paper which uses text analysis of a corpus to make a historical argument or interpretation.

## Schedule

### Week 1 (August 31): Introduction

-   [@gold_forum:_2016]
-   [@robertson_differences_2016]
-   [@cohen_searching_2010]
-   [@underwood_seven_2015]
-   [@jockers_text_2014, chs. 1--2]
-   [@arnold_humanities_2015, chs. 1--2]

Practicum: Before our first meeting, install [R](https://www.r-project.org/) and the [RStudio Desktop IDE](https://www.rstudio.com/products/rstudio/#Desktop) (you may wish to install the [preview version](https://www.rstudio.com/products/rstudio/download/preview/) to get notebook support). Start to become familiar with the basics of R as described in the introductory chapters of either Jockers or Arnold and Tilton. You should also become start to become familiar with the basics of the Unix-style command line (see Shotts, *Linux Command Line*, as a reference).

### Week 2 (September 7): Working with corpus metadata

-   [@robertson_searching_2016]
-   [@hitchcock_old_2016]
-   [@schmidt_women_2012]
-   [@flanders_data_2016]
-   [@sinclair_text_2016]

Practicum: Download at least one of the provided corpora. Extract the metadata into tables. Do an exploratory data analysis of the corpus metadata, paying special attention to the question of which texts actually belong in the corpus. See the [R Markdown documentation](http://rmarkdown.rstudio.com/) for help getting started with your first notebook. See Jenny Bryan, [*Happy Git and GitHub for the useR*](http://happygitwithr.com/), for guidance using Git/GitHub.

### Week 3 (September 21): Vector space models 

-   [@simpson_rise_1981]
-   [@graham_exploring_2015, ch. 3]
-   [@arnold_humanities_2015, chs. 9-10]
-   [@witmore_text:_2012]
-   [@schmidt_age_2011]
-   [@leskovec_mining_2014, ch. 1]

Practicum: Using [text2vec](https://cran.rstudio.com/web/packages/text2vec/), create a document-term model of a corpus. Experiment with using words and n-grams, filtering terms and stemming, applying transformations such as TF-IDF, and applying distance measures such as cosine distance. Produce plots of how terms are used. (You may want to work ahead on unsupervised clustering or principal component analysis.)

### Week 4 (October 5): Clustering and classification

-   [@friedman_history_1985, part 2]
-   [@jockers_text_2014, chs. 11--12]
-   [@jockers_text-mining_2016]
-   [@sculley_meaning_2008]
-   [@leskovec_mining_2014, ch. 7]
-   [@james_introduction_2013, chs. 2, 3, 10]

Practicum: Try a variety of unsupervised classification methods (e.g., K-means) and dimensionality reduction methods (e.g., PCA) if you did not do so last week. Using the [caret package](http://topepo.github.io/caret/index.html) make a supervised classifier (using various machine learning methods) to predict some aspect of your text. You should use the corpus metadata for your classification labels.

### Week 5 (October 19): Text reuse

-   [@friedman_history_1985, part 3]
-   [@cordell_reprinting_2015]
-   [@smith_computational_2015]
-   [@smith_detecting_2014]
-   [@leskovec_mining_2014, ch. 3]
-   Kellen Funk and Lincoln Mullen, "The Spine of American Law: Digital Text Analysis and U.S. Legal Practice." See also our [working paper](http://lmullen.github.io/civil-procedure-codes/talks/dh-working-group/Funk-Mullen.Migration-Field-Code.working-paper.pdf) and the [notebooks](http://lmullen.github.io/civil-procedure-codes/) for the project.

Practicum: Using either the [textreuse](https://github.com/ropensci/textreuse) or [LSHR](https://github.com/dselivanov/lshr) packages, look for reused passages in the corpus you are working with. Create clusters or networks of those reuses.

### Week 6 (November 2): Topic modeling

-   [@hoeflich_legal_2010, chs. 2, 3, 7]
-   *Journal of Digital Humanities* 2, no. 1 (winter 2012): <http://journalofdigitalhumanities.org/2-1/>.
-   [@nelson_mining_2011]
-   [@newman_probabilistic_2006]
-   [@graham_exploring_2015, ch. 4]
-   [@jockers_text_2014, ch. 13]
 
Practicum: Using the [textmineR](https://github.com/TommyJones/textmineR), create topic models of the corpus that you are working with. Make plots of changes in topics over time.

### Week 7 (November 16): Word-embedded models

-   [@welke_law_2010, 1-94]
-   [@schmidt_vector_2015]
-   [@schmidt_rejecting_2015]
-   [@gavin_arithmetic_2015]
-   [@mikolov_distributed_2013]
-   [@pennington_glove:_2014]
-   [@kusner_word_2015]
-   GloVe vignette from [text2vec](https://cran.rstudio.com/web/packages/text2vec/) and [wordVectors](https://github.com/bmschmidt/wordVectors) documentation.

Practicum: Train a word-embedded model on your corpus, or train several models on chronological or thematic subsets of your corpus. Use the model(s) to find distinctive use of language. More ambitious: use some of the more cutting-edge algorithms that use word vectors as inputs, for example, for finding document distances.

### Week 8 (November 30): Named-entity recognition

-   [@welke_law_2010, 95-158]
-   [@blevins_space_2014]
-   [@schmidt_language_2015]
-   [@fraas_mapping_2015]

Practicum: Using the [openNLP](https://cran.rstudio.com/web/packages/openNLP/) or [coreNLP](https://cran.rstudio.com/web/packages/coreNLP/) packages, run named-entity recognition on your corpus. Plot the entities over time or space.

## Assignments and expectations

You are encouraged to work with the [*Making of Modern Law: Legal Treatises*](http://gdc.gale.com/products/the-making-of-modern-law-legal-treatises-1800-1926/acquire/faqs/) corpus, which the library will provide. Alternatives include [*Chronicling America*](http://chroniclingamerica.loc.gov/), [*19th Century U.S. Newspapers*](http://www.gale.com/19th-century-us-newspapers/), and the [WPA former slave narratives](https://github.com/lmullen/WPAnarratives). You may substitute another corpus after consulting with me.

Come prepared to discuss the readings assigned for each meeting. There are several kinds of reading. Works on legal history set the stage for the corpus we are investigating, so that you know what questions are worth asking. If you decide to work on a corpus other than MOML, plan on substituting your own readings for works in that category. Works on other historical subjects demonstrate how the methods we are studying can be applied to historical or literary historical questions. The methodological pieces, well, explain how the methods work. Don't get hung up on the details of implementation: read to understand what the transformation is that the method accomplishes. Some of the methodological pieces are more practical, while others provide the theoretical basis. There is much more to read than I can assign, but the references below (also in a [Zotero group library](https://www.zotero.org/groups/text_analysis_for_historians)) contain many additional works. 

Before each meeting, you will create an [R Markdown notebook](http://rmarkdown.rstudio.com/r_notebooks.html) which---in prose, code, and figures---explores your chosen corpus of texts using the specified methodology. Get as far as you can with each method. You are always welcome to share code and ideas with other people in the class, though each person must turn in his or her own notebook. Submit all of the notebooks for the semester in a single GitHub repository. Name each notebook something sensible, like `08-named-entity-recognition.nb.html`. 

At the end of the semester, you will write a conference paper which makes a historical argument on the basis of the corpus you have chosen. Submit these papers as a separate GitHub repository. Prepare a one-page proposal by week 4, share a one-page statement of progress and problems by week 7, and submit the draft by 5 p.m. on December 16.

Grades will be assigned with 50% of the weight given to completing the readings and notebooks, and 50% to the final paper.

# Acknowledgments and fine print

I have benefited from syllabi by [Ted Underwood](https://tedunderwood.com/2015/01/16/syllabus-for-a-graduate-seminar/), [Andrew Goldstone](http://www.rci.rutgers.edu/~ag978/litdata/), and [Ben Schmidt](http://benschmidt.org/HDA15/). [Kellen Funk](http://kellenfunk.org/) provided advice about legal history readings.

See the [George Mason University catalog](http://catalog.gmu.edu/) for general policies, as well as the university [statement on diversity](http://ctfe.gmu.edu/professional-development/mason-diversity-statement/). You are expected to know and follow George Mason’s policies on [academic integrity](http://oai.gmu.edu/) and the [honor code](http://oai.gmu.edu/understanding-the-honor-code/). If you are a student with a disability and you need academic accommodations, please see me and contact the Office of Disability Services at 703-993-2474 or through [their website](http://ods.gmu.edu). All academic accommodations must be arranged through that office.

# References

