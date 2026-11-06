---
title: "How to hack URLs for faster searches in your browser"
date: "2012-09-06"
slug: "how-to-hack-urls-for-faster-searches-in-your-browser"
description: "Custom search engines in your browser save a step every time you search a library catalog, dictionary, or database like JSTOR."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/how-to-hack-urls-for-faster-searches-in-your-browser/42304"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="4350621750_97552795ee_z-300x225.jpeg" alt="A library card catalog" >}}

A very large part of my work involves searching. I enter queries into library catalogs, into dictionaries, into search engines, and into academic databases. Anything I can do to reduce the number of steps it takes to perform a search will save me lots of time in the long run. So here is a way to set up custom search engines in your browser. For example, suppose I want to search for a word in JSTOR. In my browser I could

<!--more-->

1. click my JSTOR bookmark,
2. click the search box,
3. type my query and hit enter.

But using this hack, I can reduce that process by one step, so that I

1. go to search box (in Google Chrome, CMD + L gets you there fast) and
2. type ‘JSTOR’ plus tab plus my query and hit enter.

To make this work, you have to know a little bit about how URLs work. URLs can hold all sorts of useful (and hackable) information. For example, suppose I go to JSTOR and search for the word “colporteur” (since it comes up a lot in my sources). The resulting URL is complicated: `http://www.jstor.org.resources.library.brandeis.edu/action/doAdvancedSearch?q0=colporteur&f0=all&c1=AND&q1=&f1=all&acc=on&wc=on&Search=Search&sd=&ed=&la=&jo=`

That’s hard to read. But I do know that I searched for the word “colporteur,” and sure enough, that word shows up in the URL. The relevant part is `?q0=colporteur`. I can deduce that if I replaced that word with another search term (say, “American Tract Society”) and kept everything else the same, I would get a JSTOR query on that term. That is indeed what happens.\*

Armed with that information, I can now create my custom search engine. I’ll explain how to do this in Google Chrome ([manufacturer’s instructions here](http://support.google.com/chrome/bin/answer.py?hl=en&answer=95653)), though you can do something [similar in Mozilla Firefox](http://support.mozilla.org/en-US/kb/search-bar-easily-choose-your-search-engine#w_adding-search-engines), which Brian [wrote about two years ago](http://chronicle.com/blogs/profhacker/speed-up-your-searching-in-firefox/23105), in [Alfred](http://chronicle.com/blogs/profhacker/alfred-obeys-your-commands-to-launch-apps-and-more/41840), and doubtless in other applications.

In Chrome, you can go to the settings and click “Manage Search Engines.” If you scroll to the bottom of the list, you’ll see a box like this.

{{< image src="screen-shot-2012-09-02-at-11.18.02-pm.png" alt="Chrome's form for adding a new search engine, with fields for name, keyword, and URL" >}}

In the first two boxes, we can add the basic information. Let’s call this search engine “JSTOR” (first box) and use the keyword `jstor` (second box) to signal Chrome when we want to use this search engine. Then in the third box, we can tell Chrome what a typical search URL looks like for the JSTOR website. Into that box we can paste the long JSTOR URL from above. But where the word “colporteur” appeared, we’re going to substitute the characters `%s`. Whenever you use this search engine, Chrome will put the query that you type there.

{{< image src="screen-shot-2012-09-02-at-11.31.32-pm.png" alt="A custom JSTOR search engine with the keyword jstor and a search URL" >}}

After adding the search engine, I can go to the search bar in Chrome, type `jstor` and hit tab, and I’m ready to use my custom search engine.

{{< image src="screen-shot-2012-09-02-at-11.32.15-pm.png" alt="Chrome's address bar searching JSTOR for colporteur" >}}

I have custom search engines set up for LouFind (my library catalog), the Library of Congress, WorldCat, Google Books, BookFinder, Google Scholar, ProfHacker, the Oxford English Dictionary, and a number of subject-specific databases. Setting up a custom search engine like this is worth the trouble for any database or site you’ll search more than twice. You’ll only save 2 to 5 seconds each query, but multiplying that times thousands of searches (not to mention avoiding unnecessary drag on your work) gives you a big savings.

**What custom search engines do you use? Do you have any favorite URL hacks?**

\* If you want a fuller explanation, read about [query strings](http://en.wikipedia.org/wiki/Query_string). Bonus tip: the `.resources.library.brandeis.edu` part of the URL sends me to JSTOR through my university’s proxy, so I actually get access to JSTOR. Your institution probably has something similar set up.

*Image courtesy of Flickr user [dfulmer](http://www.flickr.com/photos/annarbor/4350621750/) / Creative Commons licensed*[Return to Top](#)
