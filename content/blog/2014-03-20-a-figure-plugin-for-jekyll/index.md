---
title: A Figure Plugin for Jekyll
date: "2014-03-20"
slug: 'a-figure-plugin-for-jekyll'
tags:
- Jekyll
draft: false
---

This site runs on [Jekyll](http://jekyllrb.com), a static-site generator written in Ruby. As I create more and more figures to go on this site, it has become increasingly laborious to get the images, captions, and links properly formatted. So I've written a [jekyll_figure](https://github.com/lmullen/jekyll_figure) plugin to automate the figures.

A tag like this:

    {% figure filename svg,png,pdf 'Your caption here' %}

Produces output like this:

{{< figure src="12names.svg" caption="Figure 1: Twelve names whose gender changed over time" >}}

You can get the plugin at [GitHub](https://github.com/lmullen/jekyll_figure) or from [Rubygems](https://rubygems.org/gems/jekyll_figure).
