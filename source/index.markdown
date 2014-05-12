---
layout: default
title: About
---

I am a PhD candidate in the [history department][] at [Brandeis
University][] and a historian of religion in early America and the
nineteenth century. In the fall I will be an assistant professor in the
[Department of History and Art History][] at [George Mason
University][]. You can write to me at `lincoln@lincolnmullen.com`.

#### Blog

On this site I write an academic blog about history, religion, and
digital humanities, and I'm also a regular contributor at at [Religion
in American History][] and [ProfHacker][]. Here is a list of [all of my
posts][], and these are the most recent:

<ul id="posts">
{% for post in site.posts limit: 5 %}
<li>
<time>{{ post.date | date: "%b %-d:" }}</time>
<a href="{{ post.url }}">{{ post.title }}</a>
</p>
</li>
{% endfor %}
</ul>

#### Research

My [dissertation][], *The Varieties of Religious Conversion: The Origins
of Religious Choice in the United States*, is a history of conversion
between religions in the nineteenth-century United States. The
[publications page][] lists my publications and talks and provides the
full text if possible.

Below are some of the historical maps, visualizations, database, and
publication projects that I am working on, which are also listed on the
[research page][].

<div style="display: inline-block;">

<p>
<a
href="http://lincolnmullen.com/projects/us-boundaries/"><img 
style="display: inline-block;
margin-right: 6%; border: 1px solid #aaa"
src="/downloads/post/project-us-boundaries.png" width="45%" /></a>
<a href="/projects/paulists-map/"><img
style="display: inline-block; border: 1px solid #aaa"
src="/downloads/post/project-paulists-map.png" width="45%" /></a>
</p>

<p>
<a
href="http://americanconverts.org"><img style="display: inline-block;
margin-right: 6%; border: 1px solid #aaa"
src="/downloads/post/project-converts.png" width="45%" /></a>
<a href="http://lincolnmullen.com/projects/slavery/"><img
style="display: inline-block; border: 1px solid #aaa"
src="/downloads/post/project-slavery.png" width="45%" /></a>
</a>
</p>


<p>
<a href="/research/history-dissertations/"><img style="display: 
inline-block;
margin-right: 6%; border: 1px solid #aaa" 
src="/downloads/post/project-dissertations.png" width="45%" /></a>
<a href="http://jsr.fsu.edu/"><img style="display: 
inline-block;
border: 1px solid #aaa" 
src="/downloads/post/project-jsr.png" width="45%" /></a>
</p>

</div>

Much of my research is augmented by coding projects or expressed in
code. These projects are listed on the [research page][]. Most recently
I've been working with on the [gender package][] for R, which infers
gender based on first names and a range of years, along with a few Omeka
plugins and an [Omeka API client][] written in Ruby.

I [teach classes][] on American religious history, the nineteenth
century, and digital humanities.

  [history department]: http://www.brandeis.edu/departments/history/
  [Brandeis University]: http://www.brandeis.edu
  [Department of History and Art History]: http://historyarthistory.gmu.edu/
  [George Mason University]: http://www.gmu.edu/
  [Religion in American History]: http://usreligion.blogspot.com/
  [ProfHacker]: http://chronicle.com/blogs/profhacker/
  [all of my posts]: /blog/
  [dissertation]: /research/#dissertation
  [publications page]: /publications/
  [research page]: /research/
  [gender package]: https://github.com/lmullen/gender
  [Omeka API client]: https://github.com/lmullen/omeka_client
  [teach classes]: /teaching/
