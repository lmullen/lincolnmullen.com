---
title: 'Analyzing Historical History Dissertations: Limitations of the Data'
date: "2014-02-08"
slug: 'analyzing-historical-history-dissertations-limitations-of-the-data'
tags:
- digital history
draft: false
---

After my post yesterday about [historical history dissertations](http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-beginnings/), [Yoni Appelbaum](https://twitter.com/YAppelbaum) sent me some useful questions about the completeness and accuracy of the data. I've had my own questions about exactly what data I am working with, and what are its limitations. The difficulty with a reasonably large data set like this one is that it is easier to know something about the data in aggregate than any of the particulars.

The main question boils down to this: why does the data from ProQuest have 84,428 dissertations about history (based on the filters I described earlier) when the AHA's *[Directory of History Dissertations](https://secure.historians.org/pubs/dissertations/index.cfm)* contains "29,421 dissertations that were completed or are currently in progress at 194 academic departments in Canada and the United States"?

The answer is that the AHA data is based on dissertations completed within history departments. Those dissertations also have to be reported to the AHA, perhaps by the authors themselves, but usually by member history departments. The ProQuest data, as far as I can tell, comes from reports by universities to ProQuest, as well as retrospective data entry from sources like Dissertations Abstracts International and UMI. It contains data from many departments, with subject fields that describe the content of the dissertation. I'm filtering the data set to get only disserations that match these subjects.

```r
historical_subjects <- c("Art History.",
                         "Biography.",
                         "Economics, History.",
                         "Education, History of.",
                         "History of Science.",
                         "History, African.",
                         "History, Ancient.",
                         "History, Asia, Australia and Oceania.",
                         "History, Black.",
                         "History, Canadian.",
                         "History, Church.",
                         "History, European.",
                         "History, General.",
                         "History, History of Oceania.",
                         "History, Latin American.",
                         "History, Medieval.",
                         "History, Middle Eastern.",
                         "History, Military.",
                         "History, Modern.",
                         "History, Russian and Soviet.",
                         "History, United States.",
                         "History, World History.",
                         "Religion, History of.")
```

So while I'm not selecting, for example, American Studies dissertations as a whole, I am getting American Studies dissertations that someone judged to be history. This certainly leads to some fuzziness, but I think it also gives a fuller picture about the writing of history. In other words, the AHA data is probably better for answering questions about professional academic historians in history departments; the ProQuest data is probably better for answering questions about writing about history (or the past) across the academy.

But there is another question: is the ProQuest data tolerably complete? I suspect, given the very low numbers of dissertations before the 1950s, that the data is patchy and unreliable before that time. (For example, I didn't find W. E. B. Du Bois's 1896 dissertation in either data set.) Other than that, I'm not sure how to answer the question definitively. But I was curious to see whether the two data sets contained the dissertations by faculty affiliated with [my own department](http://www.brandeis.edu/departments/history/). This kind of spot check is only suggestive, of course. Leaving out two faculty with dissertations after 2012, and four faculty with PhDs from British universities, I came up with the following table.

| Faculty | In ProQuest | In AHA | Year |
|:--------|:------------|:-------|:-----|
| Antler, Joyce | yes | yes | 1977 |
| Donahue, Brian | yes | no | 1995 |
| Engerman, David | yes | yes | 1998 |
| Fischer, David Hackett | no | no | |
| Freeze, ChaeRan | yes | no | 1997 |
| Freeze, Gregory | yes | no | 1972 |
| Hang, Xing | yes | yes | 2010 |
| Hansen, Karen | yes | no | 1989 |
| Hopper, Ian | yes | yes | 2011 |
| Hulliung, Mark L. | no | no | |
| James, Heyward Parker | yes | no | 2001 |
| Kamensky, Jane | yes | yes | 1993 |
| Kapelle, William E. | no | no | |
| Reinharz, Jehuda | yes | no | 1972 |
| Sarna, Jonathan D. | yes | yes | 1979 |
| Sheppard, Eugene | yes | yes | 2001 |
| Sohrabi, Nagmeh | yes | no | 2005 |
| Sreenivasan, Govind | yes | yes | 1995 |
| Williams, Chad | yes | yes | 2004 |
| Willrich, Michael | yes | yes | 1997 |

Dissertations by three faculty did not show up in either data set. I don't have exact dates that those faculty graduated, but all three are from the 1960s or 1970s. Otherwise, the ProQuest data set contained all of the Brandeis faculty. The AHA data set, on the other hand, missed 7 additional dissertations out of 20. Those 7 dissertations are all what I would call history dissertations, though a few of them were written in non-history departments.

In sum, I think the ProQuest data is sketchy before the 1950s, but at least as reliable as the AHA data after the 1950s, taking into account a much broader definition of what counts as a history dissertation.
