---
date: '2015-03-11T08:00:00+00:00'
slug: 'detecting-text-reuse-in-legal-documents'
title: 'Detecting Text Reuse in Nineteenth-Century Legal Documents'
---

<figure id="attachment_552" style="max-width: 209px" class="wp-caption alignright">
<img class="size-full wp-image-552" src="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/ddfield.jpg?resize=209%2C300" alt="" width="209" height="300" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
David Dudley Field, leading member of the New York commission that drafted the Field Code of 1850. [Image from the New York Public Library.](http://digitalcollections.nypl.org/items/510d47df-a7c9-a3d9-e040-e00a18064a99)
</figcaption>
</figure>
How can you track changes in the law of nearly every state in the United States over the course of half a century? How can you figure out which states borrowed laws from one another, and how can you visualize the connections among the legal system as a whole?

[Kellen Funk](http://kellenfunk.org/), a historian of American law, is writing a dissertation on how codes of civil procedure spread across the United States in the second half of the nineteenth century. He and I have been collaborating on the digital part of this project, which involves identifying and visualizing the borrowings between these codes. The problem of text reuse is a common one in digital history/humanities projects. In this post I want to describe our methods and lay out some of our preliminary results. To get a fuller picture of this project, you should read the four posts that Kellen has written about his project:

<!--more-->
-   [The Influence of the Field Code: An Introduction to the Critical Issues](http://kellenfunk.org/field-code/the-influence-of-the-field-code-an-introduction/)
-   [Density Plots: Predicting the Influence of One Jurisdiction's Code on Another's](http://kellenfunk.org/field-code/density-plots/)
-   [Network Analysis: Finding the Regional Influences of Codification](http://kellenfunk.org/field-code/network-analysis/)

### The question stated

After the American Revolution, almost all states were common law jurisdictions. Beginning in the 1840s, legal reformers, led most prominently by the New York attorney David Dudley Field, attempted to revolutionize legal practice by preparing codes of civil procedure. For civil cases these codes did away with common law ways of bringing lawsuits and running trials and instead offered a simplified, rationalized, supposedly democratic way of conducting lawsuits. The first such code was submitted by a three-man commission to the New York state legislature in 1848. In 1850 the legislature enacted a revised version, commonly called the Field Code. Other states followed suit and codified their civil procedure. As Kellen [writes](http://kellenfunk.org/field-code/the-influence-of-the-field-code-an-introduction/):

> Over the course of the next several decades almost every state eventually adopted the Field reforms, eliminating their courts of chancery (equity) and requiring "fact pleading," party testimony, and unregulated lawyers' fees. Some states adopted their reforms piecemeal, but most enacted codes of their own, often following one or another draft (but usually the later, complete draft) of the Field Code.

From the first Field Code in New York in 1848 through the 1890s, thirty-two states enacted versions of the Field Code, and civil procedure in other states were influenced by it. States often revised their civil procedure, so there were multiple codes from many states. Sometimes the state commissions explicitly stated which codes they borrowed from. But in many cases, such borrowings have to be inferred from the text.

This has the makings of a classic digital history problem. To begin with, the codes are public documents from the nineteenth century so there are no copyright restrictions. The texts are readily available from Google Books or the Hathi Trust. Kellen is in the process of identifying and gathering all of the relevant texts and quite a few others besides into his [American Legislation Project](http://kellenfunk.org/legislation/). Second, we have a large volume of text: the 1850 New York code alone runs to 969 pages. These are texts which are amenable to distant reading. In digital literary studies one could at least make the argument that, say, novels are interesting when read closely.[<sup>1</sup>](#fn1){\#fnref1.footnoteRef} But (though I don't wish to disparage the sources that legal historians use) the codes themselves are mind-numbing. It is probably uncontroversial to say that we are all better off reading these texts at a distance instead of slogging through them subsection by subsection. And third, the relationships between these texts can easily be expressed in terms of network analysis---one of the common methods of digital history.

And most important, there is a real historical payoff to doing digital work on these codes. Legal historians have long known about the Field Code in New York; they have also long known that other states emulated that code. No one, however, has figured out exactly which parts of which codes were borrowed and whether states adapted their codes from states other than New York.[<sup>2</sup>](#fn2){\#fnref2.footnoteRef} In other words, this project is also a classic problem of historical revisionism, whereby an accepted historical narrative is made considerably more complex. And, as I'll report below, we have indeed found that the relationships between state legal codes is more complex than previously argued.

### An example borrowing

But first, let's look at one instance of the kinds of borrowings that we are trying to detect in order to see what the potential difficulties are. For this example, I'll show one place where California's 1851 code of civil procedure borrowed from New York's 1850 code. Keep in mind that California was admitted as a state in September 1850, so in 1851 it was revising its code of civil procedure for the first time as a state.

First, here is a [section on serving summonses from the 1850 New York code](http://books.google.com/books?id=9HEbAQAAIAAJ&dq=%22Documents%20of%20the%20Assembly%20of%20the%20State%20of%20New%20York%22%201850%20david%20graham&pg=PA257#v=onepage&q&f=false).[<sup>3</sup>](#fn3){\#fnref3.footnoteRef}
<figure id="attachment_557" style="max-width: 609px" class="wp-caption alignnone">
<img class="size-full wp-image-557" src="//files.lincolnmullen.com/uploads/2017/01/ny1850-summons.png?resize=609%2C266" alt="" width="609" height="266" srcset="//files.lincolnmullen.com/uploads/2017/01/ny1850-summons.png?w=609 609w, //files.lincolnmullen.com/uploads/2017/01/ny1850-summons.png?resize=300%2C131 300w" sizes="(max-width: 609px) 100vw, 609px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
An excerpt from the 1850 New York code of civil procedure.
</figcaption>
</figure>
Transcription:

> § 628. The summons must be served by delivering a copy thereof, as follows:
>
> 1.  If the suit be against a corporation, to the president or other head of the corporation, secretary, cashier, or managing agent thereof:
>
> 2.  If against a minor under the age of fourteen years, to such minor personally, and also to his father, mother, ...
>
And here is the [related section in the California code](http://books.google.com/books?id=4PHEAAAAIAAJ&ots=TDfCSa1x9C&dq=laws%20of%20california&pg=PA51#v=onepage&q&f=false).[<sup>4</sup>](#fn4){\#fnref4.footnoteRef} I have highlighted the changes in yellow. You can look at the surrounding text for both codes on Google Books and get a sense of the other borrowings.
<figure id="attachment_550" style="max-width: 647px" class="wp-caption alignnone">
<img class="size-full wp-image-550" src="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/ca1851-summons.png?resize=647%2C163" alt="" width="647" height="163" srcset="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/ca1851-summons.png?w=647 647w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/ca1851-summons.png?resize=300%2C76 300w" sizes="(max-width: 647px) 100vw, 647px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
An excerpt from the 1851 California code of civil procedure.
</figcaption>
</figure>
Transcription:

> § [29]{style="background-color: yellow;"}. The summons [shall]{style="background-color: yellow;"} be served by delivering a copy thereof [attached to the certified copy of the complaint]{style="background-color: yellow;"} as follows:
>
> [1st.]{style="background-color: yellow;"} If the suit be against a corporation, to the president or other head of the corporation, secretary, cashier, or managing agent thereof:
>
> [2d.]{style="background-color: yellow;"} If against a minor under the age of fourteen years, to such minor personally, and also to his father, mother, or guardian; or if there ...

The borrowing is obvious, but variations make it complicated to detect the similarity algorithmically. In the example above, California adds the words "attached to the certified copy of the complaint." Sometimes the variations are due to local circumstances, as in the example below, where the California code inserts the local names of its courts and sensibly requires that actions happen in San Francisco instead of New York. In addition to actual changes in the text, imperfect OCR adds noise. Though we have taken some pains to get good OCR and were mostly successful, I think, the noise still adds an additional challenge.
<figure id="attachment_549" style="max-width: 659px" class="wp-caption alignnone">
<img class="size-full wp-image-549" src="//files.lincolnmullen.com/uploads/2017/01/ca1851-sanfrancisco.png?resize=659%2C150" alt="" width="659" height="150" srcset="//files.lincolnmullen.com/uploads/2017/01/ca1851-sanfrancisco.png?w=659 659w, //files.lincolnmullen.com/uploads/2017/01/ca1851-sanfrancisco.png?resize=300%2C68 300w" sizes="(max-width: 659px) 100vw, 659px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Section 22 of the California code from 1851.
</figcaption>
</figure>
So the challenge is to figure out how to detect these borrowings at a large scale, both identifying which codes were most likely to be derived from others and also identifying which sections in particular were borrowed. This project is thus similar in purpose and method, though rather different in content, to the *[Viral Texts](http://viraltexts.org/)* project created by Ryan Cordell, Elizabeth Maddock Dillon, David Smith, Abby Mullen, Peter Roby, Kevin Smith, and Matthew Williamson to detect communities of reprinting in nineteenth-century newspapers.

### Using n-grams to detect borrowings

For our test run of the project we started with 61 codes, which comprise some ten thousand pages and four million words. This is only a middling-sized dataset, compared to some, but it is still nothing to sneeze at. Our first priority in investigating this corpus of civil procedures codes was to figure out whether we could detect (and visualize) where borrowings occurred between codes. Like the *Viral Texts* team (and many other people doing the same kind of work), we use n-grams to break the text into small chunks in order to detect reuse.[<sup>5</sup>](#fn5){\#fnref5.footnoteRef}

Our method worked like this. After stripping punctuation and tokenizing the text into words, we first split the texts into n-grams: we chose five-grams after some experimentation. We then filtered n-grams that seemed unreasonable: anything containing a non-alphabetic character is likely to be just OCR noise. We then found all of n-grams that were shared between codes (i.e., the intersection of the two sets).

Using n-grams in this way lets us detect shared passages while allowing for minor changes and errors in the OCR. I hope the following example does not belabor what is a fairly simple point. But let's take the example above where California 1851 borrowed from New York 1850, and split the two passages into five-grams. Splitting the texts into short sequences duplicates the words across the n-grams. This redundancy permits us to detect matches even when there are variations---meaningful or not---in the text. The table below shows the five-grams from the excerpts above and whether or not they match.

<table>
<tr class="header">
<th style="text-align: left;">
NY 1850
</th>
<th style="text-align: left;">
CA 1851
</th>
<th style="text-align: left;">
match?
</th>
</tr>
<tr class="odd">
<td style="text-align: left;">
the summons must be served
</td>
<td style="text-align: left;">
the summons shall be served
</td>
<td style="text-align: left;">
no
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
summons must be served by
</td>
<td style="text-align: left;">
summons shall be serve by
</td>
<td style="text-align: left;">
no
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
must be served by delivering
</td>
<td style="text-align: left;">
shall be served by delivering
</td>
<td style="text-align: left;">
no
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
be served by delivering a
</td>
<td style="text-align: left;">
be served by delivering a
</td>
<td style="text-align: left;">
yes
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
served by delivering a copy
</td>
<td style="text-align: left;">
served by delivering a copy
</td>
<td style="text-align: left;">
yes
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
by delivering a copy thereof
</td>
<td style="text-align: left;">
by delivering a copy thereof
</td>
<td style="text-align: left;">
yes
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
delivering a copy thereof as
</td>
<td style="text-align: left;">
delivering a copy thereof attached
</td>
<td style="text-align: left;">
no
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
a copy thereof as follows
</td>
<td style="text-align: left;">
a copy thereof attached to
</td>
<td style="text-align: left;">
no
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
...
</td>
<td style="text-align: left;">
...
</td>
<td style="text-align: left;">
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
as follows if the suit
</td>
<td style="text-align: left;">
as follows if the suit
</td>
<td style="text-align: left;">
yes
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
follows if the suit be
</td>
<td style="text-align: left;">
follows if the suit be
</td>
<td style="text-align: left;">
yes
</td>
</tr>
</table>
We then extended this technique to compare any two codes. That is, for each five-gram in the California 1851 code, we see whether it has a match in the New York 1850 code. The plot below shows a vertical line for each five-gram in the California code that has a match somewhere in the New York code.
<figure id="attachment_548" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-548" src="//files.lincolnmullen.com/uploads/2017/01/ca1851-density.jpg?resize=739%2C240" alt="" width="739" height="240" srcset="//files.lincolnmullen.com/uploads/2017/01/ca1851-density.jpg?resize=768%2C249 768w, //files.lincolnmullen.com/uploads/2017/01/ca1851-density.jpg?resize=300%2C97 300w, //files.lincolnmullen.com/uploads/2017/01/ca1851-density.jpg?w=861 861w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Places where the California 1851 code borrowed from the New York 1850 code.
</figcaption>
</figure>
One's reaction to a first look at that plot ought to be skepticism. It is interesting to note that there are definitely places in the code where the borrowing is heavier than in other places. I take that variability as an invitation <del>for me</del> for Kellen to go back and read those sections of the code. But one should be skeptical because the degree to which it looks like California borrowed from New York is almost entirely dependent on the level of transparency one assigns to the vertical lines indicating matches. I could make you think that everything or nothing was borrowed by adjusting that one setting.[<sup>6</sup>](#fn6){\#fnref6.footnoteRef}

However, I don't want you to look at that plot in isolation. Instead, compare it to the plot below which shows matches between Michigan's 1853 book of court rules, covering the same topics as a code of civil procedure, and New York's 1850 code.
<figure id="attachment_554" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-554" src="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?resize=739%2C277" alt="" width="739" height="277" srcset="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?resize=768%2C288 768w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?resize=300%2C113 300w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?resize=1024%2C384 1024w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?w=1536 1536w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/mi1853-density.png?w=1478 1478w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Places were the Michigan 1853 book of court rules borrowed from the New York 1850 code.
</figcaption>
</figure>
Just as we chose the California 1851 code because we knew it borrowed from New York, we also chose the 1853 Michigan code because knew it did *not* borrow from New York. And that is what the two plots show: California borrows heavily and Michigan borrows hardly at all. The noise in the Michigan plot is what we would expect. All of these codes are formulaic documents using a highly specialized (and verbose) vocabulary, so it is not surprising that there are phrases of even five words that occasionally show up in both codes. Exactly how reliable this method remains to be determined precisely, and of course these plots could be improved. But these plots do provide persuasive, if pragmatic, evidence that this method works: we can detect which codes borrow and which codes do not. Furthermore, by writing a kind of [keywords in context](http://en.wikipedia.org/wiki/Key_Word_in_Context) function, we were able to spot check the matches to see that they were in fact significant borrowings of meaningful length.[<sup>7</sup>](#fn7){\#fnref7.footnoteRef}

### Visualizing the network of borrowings

Once we were sure that we could usefully detect matches between codes, we wanted to get an overview of our entire corpus.[<sup>8</sup>](#fn8){\#fnref8.footnoteRef} To do this, we reduced the relationship between each code to a single number expressing the degree to which code B borrowed from code A. We tried this several ways. The simplest way is to figure out whether a given n-gram in code B also occurs in code A, then to calculate the proportion of matches to the total number of n-grams. We have to experiment further with this. Another way is was to calculate the ratio between the unique number of five-grams in both codes and the unique number of the five-grams in the destination code. This gives us a single number between 0 and 1 for each combination of codes.[<sup>9</sup>](#fn9){\#fnref9.footnoteRef} Conceptually, this is equivalent to the proportion of the destination code that is borrowed from the origin code. So for example the California 1851 code above has a score of .387 when compared to the New York 1850 code: a significant proportion of the California code was borrowed from the New York code. When a code is compared to itself, the pairing gets a score of 1, i.e., they are a perfect match. Two codes in which there were absolutely no borrowings would receive a score of 0. For comparison, when the Michigan 1853 is compared to New York 1850 (as plotted above) it receives a score of .032.

From there we computed a matrix of similarity scores between all the possible borrowings. The table excerpted below shows the degree of similarity between California's 1851 code and the other codes listed.

<table>
<caption>
CA 1851's similarity to other codes
</caption>
<tr class="header">
<th style="text-align: left;">
code
</th>
<th style="text-align: left;">
score
</th>
</tr>
<tr class="odd">
<td style="text-align: left;">
AZ 1865
</td>
<td style="text-align: left;">
.759
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
NV 1861
</td>
<td style="text-align: left;">
.695
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
NV 1869
</td>
<td style="text-align: left;">
.656
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
UT 1870
</td>
<td style="text-align: left;">
.476
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
NY 1850
</td>
<td style="text-align: left;">
.387
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
FL 1870
</td>
<td style="text-align: left;">
.188
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
NC 1868
</td>
<td style="text-align: left;">
.184
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
NY 1849
</td>
<td style="text-align: left;">
.174
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
CA 1850
</td>
<td style="text-align: left;">
.139
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
NY 1848
</td>
<td style="text-align: left;">
.115
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
OH 1853
</td>
<td style="text-align: left;">
.091
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
KY 1851
</td>
<td style="text-align: left;">
.065
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
MO 1849
</td>
<td style="text-align: left;">
.056
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
MO 1856
</td>
<td style="text-align: left;">
.045
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
GA 1860
</td>
<td style="text-align: left;">
.022
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
MI 1853
</td>
<td style="text-align: left;">
.013
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
CT 1879
</td>
<td style="text-align: left;">
.003
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
UT 1853
</td>
<td style="text-align: left;">
.002
</td>
</tr>
</table>
Since 1851 was very early in the codification movement, most of these comparisons are anachronistic: it is of course impossible for California in 1851 to have borrowed from Arizona in 1865. So we filtered our matrix so we did not look for earlier codes borrowing from later codes. Time's arrow alone is helpful for sorting out these relationships. The remaining matrix expresses a network graph, where the rows and columns are the nodes, and the values of the matrix are the weights of the edges. In such a matrix every code is connected to every other code, even if the connection (as betweeen CA 1851 and MO 1849 in the table above) so weak as to be just noise. So we again filtered the matrix to remove connections below an arbitrary weight.

Using the resulting matrix, we can created the following network visualization of how codes borrowed from one another.[<sup>10</sup>](#fn10){\#fnref10.footnoteRef} (Click the figure to get a bigger version and zoom in with your browser.)

 
<figure id="figure-7">
[!['Network](/figures/field-code/codes-s10.svg)]( "/figures/field-code/codes-s10.svg")
<figcaption>
Figure 7: Network graph of borrowings between codes of civil procedure in the nineteenth-century United States (weight ≥ .10). \[[SVG](/figures/field-code/codes-s10.svg)\]
</figcaption>
</figure>
The network graph below extracts the largest family of codes that were related to one another, and is much stricter about which connections are retained. Just as we expected, this graph centers on the 1850 New York code, which radiates connections out to the surrounding codes. (If the Field Code were not central, as we know it to be from other historical research, then we would have a had a strong presumption that our method had a problem.) But what is different about this visualization is the detailed connections between all the other codes.
<figure id="figure-8">
[!['The](/figures/field-code/code-family-s25.svg)]( "/figures/field-code/code-family-s25.svg")
<figcaption>
Figure 8: The largest family of codes which borrowed from one another (weight ≥ .25). \[[SVG](/figures/field-code/code-family-s25.svg)\]
</figcaption>
</figure>
### Preliminary findings and next steps

Our most significant findings so far both confirms and modifies the existing interpretations of the codification movement. The New York 1850 code, the original Field Code, was indeed the most important of the codes of civil procedure. But the NY 1850 code was only weakly to moderately related to its derivatives. Its significance was not in providing the model for each of the codes of civil procedure that followed. Rather, the Field Code provided the model for several regional families of codes. The Field Code, for instance, strongly influenced the California 1851 code. But the California 1851 code provided the archetype for a set of codes in the Western states. Utah, Montana, Arizona, Nevada, Colorado, and Idaho borrowed their codes almost entirely from California's. Other regional borrowings can be detected. North Carolina in 1868 borrowed from New York's 1849 code (an earlier version of the 1850 code), and Florida and South Carolina in 1870 borrowed from North Carolina. In 1862 Oregon borrowed from New York's 1850 code, and in 1900 Alaska borrowed from Oregon. The story is not that states adopted the Field Code, but that the first or more important state in each region used the Field Code as a model, and then the other states crafted codes from local models that worked for the particular needs and conditions of those places.

There is much more to be done with these codes. I'll mention the three most obvious next steps, then two possible further steps. First, the easiest thing to do is to extend this analysis to additional codes, which is simply a matter of identifying and OCRing them and then re-running the analysis. Second, it will be necessary to identify the content of the borrowings between codes. For example, which parts of the Field Code did California adopt, and how did California change those adoptions over time? Which parts of California's code did the western states adopt? And third, it will be necessary to improve our analysis and make it more rigorous.

On that last point. We are at the stage in the project where we have proven (to ourselves, at least) the usefulness of this method. I often find that I can divide data analysis projects like this one into two phases (usually iterated). The first is a period of quick and dirty building of code to discover what the main questions are and what possible approaches might work. The second is a period of working to answer specific questions using methods which can be justified step by step to other scholars. In between, it's necessary to burn down all the work that one has already done and start fresh. I think we have discovered the main lines of inquiry in the first phase, and now we are in the intermediate (read, burning down to rebuild better) phase.

Finally, two tentative further steps. First, it may prove worthwhile to create an interactive web version of some of these visualizations, and perhaps of a "code browser," in order to permit others to explore these borrowings for themselves. Whether there would be enough interest to justify the expense of time we have not yet determined. And second, we think this method is more broadly applicable to the field of legal history. We hope that readers of these posts might offer some critiques of our methods. Eventually we plan to write up our methods in a more formal way, as well as to extend our work to legislative borrowings more broadly.

::: {style="margin-top: 5rem;"}
:::

#### Appendix: check our work

All the code for this project is [available on GitHub](https://github.com/lmullen/civil-procedure-codes). In the `text` directory of that repository you can find the plain-text versions of all the codes that we have OCRed, if you wish to run your own analysis. See also Kellen's [American Legislation Project](http://kellenfunk.org/legislation/), which gathers the available scanned versions of the published codes. Several RPubs files ([1](http://rpubs.com/lmullen/24393), [2](http://rpubs.com/lmullen/24527), [3](http://rpubs.com/lmullen/24776), [4](http://rpubs.com/lmullen/24945), [5](http://rpubs.com/lmullen/52960)) show the earlier stages of our investigation. I'm not really proud of some of the code in those documents, but this is as close to [open notebook history](http://wcm1.web.rice.edu/open-notebook-history.html) as we can get. We will of course improve the code as we go.
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
Well, not all novels, and certainly not all nineteenth-century novels.<a href="#fnref1">↩</a>
</li>
<li id="fn2">
For that matter, no one has done the archival work on the legal reformers either. It is important to remember that Kellen's larger project is a mix of macroscopic digital historical work and old-fashioned archival <em>Sitzfleisch</em>.<a href="#fnref2">↩</a>
</li>
<li id="fn3">
<em>The Code of Civil Procedure of the State of New-York</em> (1850), title 4, section 628, p. <a href="http://books.google.com/books?id=9HEbAQAAIAAJ&dq=%22Documents%20of%20the%20Assembly%20of%20the%20State%20of%20New%20York%22%201850%20david%20graham&pg=PA257#v=onepage&q&f=false">257</a>.<a href="#fnref3">↩</a>
</li>
<li id="fn4">
"An Act to regulate proceedings in Civil Cases, in the Courts of Justice of this State," title 3, section 29, in <em>The Statutes of California, Passed at the Second Session of the Legislature</em> (1851), <a href="http://books.google.com/books?id=4PHEAAAAIAAJ&ots=TDfCSa1x9C&dq=laws%20of%20california&pg=PA51#v=onepage&q&f=false">55</a>.<a href="#fnref4">↩</a>
</li>
<li id="fn5">
Some of our code is modeled on the algorithms described in David A. Smith, Ryan Cordell, and Elizabeth Maddock Dillon, "<a href="http://www.ccs.neu.edu/home/dasmith/infect-bighum-2013.pdf">Infectious texts: Modeling Text Reuse in Nineteenth-century Newspapers</a>," <em>IEEE Workshop on Big Data and the Humanities</em> (2013).<a href="#fnref5">↩</a>
</li>
<li id="fn6">
For the record, the alpha setting for that plot is <code>0.1</code>.<a href="#fnref6">↩</a>
</li>
<li id="fn7">
A next step will be to extract not just phrases or sentences but the large sections of the codes which are heavily reused, just as <em>Viral Texts</em> does in extracting reprinted newspaper articles.<a href="#fnref7">↩</a>
</li>
<li id="fn8">
The charts below are from the first 61 codes that we found and OCRed. We will eventually extend this approach to every relevant code that can be found. At last count we are up to 76.<a href="#fnref8">↩</a>
</li>
<li id="fn9">
In our work so far, the difference between these two methods has not produced significant differences in the network graphs below. We used the second method most, and report on it in this post. But we think the simpler method is probably better and we are going to move to that. See the discussion of next steps below.<a href="#fnref9">↩</a>
</li>
<li id="fn10">
There are other ways of visualizing the results, including clustering the civil procedure codes, but that is a subject for further work.<a href="#fnref10">↩</a>
</li>
</section>
