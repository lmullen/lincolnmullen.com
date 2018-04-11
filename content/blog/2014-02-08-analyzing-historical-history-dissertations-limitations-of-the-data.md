---
date: '2014-02-08T15:01:00+00:00'
slug: 'analyzing-historical-history-dissertations-limitations-of-the-data'
title: 'Analyzing Historical History Dissertations: Limitations of the Data'
---

*This series on "Analyzing Historical History Dissertations" is a work in progress and I've re-done some of these visualizations. If you would like to cite or link to this work in progress, please consider using the **[landing page](http://lincolnmullen.com/research/history-dissertations/)**, which will always have the most up-to-date information and a list of all the posts.*

After my post yesterday about [historical history dissertations](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-beginnings/), [Yoni Appelbaum](https://twitter.com/YAppelbaum) sent me some useful questions about the completeness and accuracy of the data. I've had my own questions about exactly what data I am working with, and what are its limitations. The difficulty with a reasonably large data set like this one is that it is easier to know something about the data in aggregate than any of the particulars.

The main question boils down to this: why does the data from ProQuest have 84,428 dissertations about history (based on the filters I described earlier) when the AHA's *[Directory of History Dissertations](https://secure.historians.org/pubs/dissertations/index.cfm)* contains "29,421 dissertations that were completed or are currently in progress at 194 academic departments in Canada and the United States"?

The answer is that the AHA data is based on dissertations completed within history departments. Those dissertations also have to be reported to the AHA, perhaps by the authors themselves, but usually by member history departments. The ProQuest data, as far as I can tell, comes from reports by universities to ProQuest, as well as retrospective data entry from sources like Dissertations Abstracts International and UMI. It contains data from many departments, with subject fields that describe the content of the dissertation. I'm filtering the data set to get only disserations that match these subjects.
<figure class="highlight">
<pre><code class="language-r" data-lang="r">&lt;span class="n">historical_subjects&lt;/span> &lt;span class="o">&lt;-&lt;/span> &lt;span class="n">c&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"Art History."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"Biography."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"Economics, History."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"Education, History of."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History of Science."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, African."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Ancient."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Asia, Australia and Oceania."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Black."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Canadian."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Church."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, European."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, General."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, History of Oceania."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Latin American."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Medieval."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Middle Eastern."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Military."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Modern."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, Russian and Soviet."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, United States."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"History, World History."&lt;/span>&lt;span class="p">,&lt;/span>
                         &lt;span class="s2">"Religion, History of."&lt;/span>&lt;span class="p">)&lt;/span></code></pre>
</figure>
So while I'm not selecting, for example, American Studies dissertations as a whole, I am getting American Studies dissertations that someone judged to be history. This certainly leads to some fuzziness, but I think it also gives a fuller picture about the writing of history. In other words, the AHA data is probably better for answering questions about professional academic historians in history departments; the ProQuest data is probably better for answering questions about writing about history (or the past) across the academy.

But there is another question: is the ProQuest data tolerably complete? I suspect, given the very low numbers of dissertations before the 1950s, that the data is patchy and unreliable before that time. (For example, I didn't find W. E. B. Du Bois's 1896 dissertation in either data set.) Other than that, I'm not sure how to answer the question definitively. But I was curious to see whether the two data sets contained the dissertations by faculty affiliated with [my own department](http://www.brandeis.edu/departments/history/). This kind of spot check is only suggestive, of course. Leaving out two faculty with dissertations after 2012, and four faculty with PhDs from British universities, I came up with the following table.

<table>
<tr class="header">
<th style="text-align: left;">
Faculty
</th>
<th style="text-align: left;">
In ProQuest
</th>
<th style="text-align: left;">
In AHA
</th>
<th style="text-align: left;">
Year
</th>
</tr>
<tr class="odd">
<td style="text-align: left;">
Antler, Joyce
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1977
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Donahue, Brian
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
1995
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Engerman, David
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1998
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Fischer, David Hackett
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Freeze, ChaeRan
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
1997
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Freeze, Gregory
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
1972
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Hang, Xing
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
2010
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Hansen, Karen
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
1989
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Hopper, Ian
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
2011
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Hulliung, Mark L.
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
James, Heyward Parker
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
2001
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Kamensky, Jane
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1993
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Kapelle, William E.
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Reinharz, Jehuda
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
1972
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Sarna, Jonathan D.
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1979
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Sheppard, Eugene
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
2001
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Sohrabi, Nagmeh
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
no
</td>
<td style="text-align: left;">
2005
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Sreenivasan, Govind
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1995
</td>
</tr>
<tr class="odd">
<td style="text-align: left;">
Williams, Chad
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
2004
</td>
</tr>
<tr class="even">
<td style="text-align: left;">
Willrich, Michael
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
yes
</td>
<td style="text-align: left;">
1997
</td>
</tr>
</table>
Dissertations by three faculty did not show up in either data set. I don't have exact dates that those faculty graduated, but all three are from the 1960s or 1970s. Otherwise, the ProQuest data set contained all of the Brandeis faculty. The AHA data set, on the other hand, missed 7 additional dissertations out of 20. Those 7 dissertations are all what I would call history dissertations, though a few of them were written in non-history departments.

In sum, I think the ProQuest data is sketchy before the 1950s, but at least as reliable as the AHA data after the 1950s, taking into account a much broader definition of what counts as a history dissertation.
