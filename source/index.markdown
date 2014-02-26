---
layout: default
title: About
---

Greetings. I am a PhD candidate in the [history department][] at
[Brandeis University][] and a historian of religion in early America and
the nineteenth century.

#### Research

My [dissertation][] is a history of conversion between religions in the
nineteenth-century United States. My [curriculum vitae][] has my
publications, talks, and workshops, usually in full text. 

My [research][] is made public in databases and other digital projects.  
Along with Erin Bartram, I've created the [American Converts 
Database][], and I'm working on a project on the [historical 
demographics of American religion][]. I [teach classes][] on American 
religious history, the nineteenth century, and digital humanities. 

#### Blog

On this site I write an academic blog about history, religion, and
digital humanities, and I'm also a regular contributor at at [Religion
in American History][] and [ProfHacker][]. Here is a list of [all of my
posts][], and these are the most recent:

<ul id="posts">
{% for post in site.posts limit: 5 %}
<li>
<time>{{ post.date | date: "%b %-d:" }}</time> <a href="{{ post.url }}">{{ post.title }}</a>
</p>
</li>
{% endfor %}
</ul>

#### Code

Much of my research is augmented by coding projects or expressed in
code. These projects are listed on the [research page][research].

  [history department]: http://www.brandeis.edu/departments/history/
  [Brandeis University]: http://www.brandeis.edu
  [dissertation]: /research/#dissertation
  [curriculum vitae]: /cv/
  [research]: /research/
  [American Converts Database]: http://americanconverts.org
  [historical demographics of American religion]: https://github.com/lmullen/demographics-religion
  [teach classes]: /teaching/
  [Religion in American History]: http://usreligion.blogspot.com/
  [ProfHacker]: http://chronicle.com/blogs/profhacker/
  [all of my posts]: /blog/
  [GitHub]: http://github.com/lmullen
  [Twitter]: http://twitter.com/lincolnmullen
