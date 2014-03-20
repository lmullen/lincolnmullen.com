---
layout: post
title: "A Figure Plugin for Jekyll"
date: 2014-03-19 21:27
categories: 
- jekyll
- Ruby
- blogging
---

This site runs on [Jekyll][], a static-site generator written in Ruby.
As I create [more][] and [more][1] [figures][] to go on this site, it
has become increasingly laborious to get the images, captions, and links
properly formatted. So I've written a [jekyll\_figure][] plugin to 
automate the figures.

A tag like this:

    {% raw %}
    {% figure filename svg,png,pdf 'Your caption here' %}
    {% endraw %}

Produces output like this:

{% figure demo/12names svg,pdf "Twelve names whose gender changed over time" %}

You can get the plugin at [GitHub][jekyll\_figure] or from [Rubygems][].

  [Jekyll]: http://jekyllrb.com
  [more]: /blog/charting-faculty-salaries/
  [1]: /blog/analyzing-historical-history-dissertations-page-counts/
  [figures]: /blog/historical-religion-data-in-the-nhgis/
  [jekyll\_figure]: https://github.com/lmullen/jekyll_figure
  [Rubygems]: https://rubygems.org/gems/jekyll_figure
