---
layout: post
title: "Paper Machines Visualizes Your Zotero Library"
date: 2012-11-08 12:00
comments: true
categories: 
- ProfHacker
tags: 
- ProfHacker
- Zotero
- Paper Machines
- visualization
- data mining
---

*[This post originally appeared at [ProfHacker][].]*

In the past few months there has been a lot of attention paid to a
[Zotero][] plugin called [Paper Machines][]. Created by Brown grad
student [Chris Johnson-Roberson][] as a Google Summer of Code
project, Paper machines was coded by Brown grad student [Chris
Johnson-Roberson][], and [Jo Guldi][] and [Matthew Battles][] directed
the project. Paper Machines uses the data in a Zotero collection to
generate analyses and visualizations. If you have a sizeable collection
of documents with good metadata and full text (for example, PDFs with
text layers) then Paper Machines can run its analysis. The most basic
output is a word cloud. More advanced analyses provide phrases matching
a pattern, maps of place names, annotations of people, places, and
organizations. Paper Machines can also perform topic modeling using
MALLET.

I ran Paper Machines on the digitized primary sources for a chapter of
my dissertation. I didn’t do any work to clean up the data, and this is
obviously just a first pass without any analysis. But I was impressed by
this map of place names, which generated the kind of
American-British-German connection that I thought I saw in the sources.
But the plot also contains places that I didn’t expect, which I hope
will be a fruitful area to research further.

{% img center http://chronicle.com/blogs/profhacker/files/2012/11/Screen-Shot-2012-11-07-at-8.13.41-PM.png %}

For a much fuller description of Paper Machines, you can see the
[documentation][Paper Machines], Johnson-Roberson’s [write up][Paper
Machines Map] at Harvard’s metaLAB, or a two [blog][] [posts][] by
Sarita Alami from the Emory libraries.

**Have you tried Paper Machines with your Zotero library?**

**Update (11/8):** In the original post I neglected to give credit to Jo
Guldi and Matthew Battles who started and directed the project. You can
read a post about Paper Machines by Guldi [here][].

**Update (11/9):** Chris Johnson-Roberson sends this update about the
contributions of the various collaborators: ”Paper Machines was coded by
Brown grad student Chris Johnson-Roberson in collaboration with
historian Jo Guldi, under the guidance of Matthew Battles at
Harvard’s [metaLAB][].”

  [ProfHacker]: http://chronicle.com/blogs/profhacker/paper-machines-visualizes-your-zotero-library/44056
  [Zotero]: http://chronicle.com/blogs/profhacker/tag/zotero
  [Paper Machines]: https://github.com/chrisjr/papermachines
  [Chris Johnson-Roberson]: http://www.chrisjr.org/
  [Jo Guldi]: http://www.joguldi.com/
  [Matthew Battles]: http://metalab.harvard.edu/people/
  [blog]: http://web.library.emory.edu/blog/supercharge-your-zotero-library-using-paper-machines-part-i
  [posts]: http://web.library.emory.edu/blog/supercharge-your-zotero-library-using-paper-machines-part-ii-6
  [here]: http://landscape.blogspot.com/2012/03/understanding-paper-machines.html
  [metaLAB]: http://metalab.harvard.edu/
