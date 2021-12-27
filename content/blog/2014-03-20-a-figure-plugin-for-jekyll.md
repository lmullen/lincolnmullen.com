---
date: '2014-03-20T01:27:00+00:00'
slug: 'a-figure-plugin-for-jekyll'
title: A Figure Plugin for Jekyll
---

This site runs on [Jekyll](http://jekyllrb.com), a static-site generator written in Ruby. As I create [more](/blog/charting-faculty-salaries/) and [more](/blog/analyzing-historical-history-dissertations-page-counts/) [figures](/blog/historical-religion-data-in-the-nhgis/) to go on this site, it has become increasingly laborious to get the images, captions, and links properly formatted. So I've written a [jekyll\_figure](https://github.com/lmullen/jekyll_figure) plugin to automate the figures.

A tag like this:

    {% figure filename svg,png,pdf &#39;Your caption here&#39; %}

Produces output like this:
<figure id="figure-1">
<a onclick="ga('send', 'event', { 'eventCategory': 'Figure', 'eventAction': 'View', 'eventLabel': 'demo/12names'});" href='//files.lincolnmullen.com/figures//demo/12names.svg'><img src='//files.lincolnmullen.com/figures//demo/12names.svg' alt='Twelve names whose gender changed over time'></a>
<figcaption>
Figure 1: Twelve names whose gender changed over time \[

<a onclick="ga(&#39;send&#39;, &#39;event&#39;, { &#39;eventCategory&#39;: &#39;Figure&#39;, &#39;eventAction&#39;: &#39;View&#39;, &#39;eventLabel&#39;: &#39;demo/12names&#39;});" href="//files.lincolnmullen.com/figures//demo/12names.svg">SVG</a>, <a onclick="ga(&#39;send&#39;, &#39;event&#39;, { &#39;eventCategory&#39;: &#39;Figure&#39;, &#39;eventAction&#39;: &#39;View&#39;, &#39;eventLabel&#39;: &#39;demo/12names&#39;});" href="//files.lincolnmullen.com/figures//demo/12names.pdf">PDF</a>\]

</figcaption>
</figure>
You can get the plugin at [GitHub](https://github.com/lmullen/jekyll_figure) or from [Rubygems](https://rubygems.org/gems/jekyll_figure).
