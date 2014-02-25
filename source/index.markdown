---
layout: default
title: About
---

Greetings. I am a PhD candidate in the [history department][] at
[Brandeis University][] and a historian of religion in early America and
the nineteenth century.

My [dissertation][] is a history of conversion between religions in the
nineteenth-century United States. My [curriculum vitae][] has my
publications, talks, and workshops, usually in full text. My
[research][] is made public in databases and other digital projects.
Along with Erin Bartram, I've created the [American Converts
Database][], and I'm working on a project on the [historical
demographics of American religion][]. I [teach classes][] on American
religious history, the nineteenth century, and digital humanities. I'm a
regular contributor at at [Religion in American History][] and
[ProfHacker][].

On this site I write an academic blog about history, religion, and
digital humanities, which also aggregates my writings elsewhere. Here is
a list of [all of my posts][], and these are the most recent:

<ul id="posts">
{% for post in site.posts limit: 5 %}
<li>
<a href="{{ post.url }}">{{ post.title }}</a>
</li>
{% endfor %}
</ul>

Much of my research is supported or augmented by coding projects, which 
are listed on the [research page][research].

You can write to me at <lincoln@lincolnmullen.com> or find me on
[GitHub][] or [Twitter][].

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


<div class="posts">
  {% for post in paginator.posts %}
  <div class="post">
    <h1 class="post-title">
      <a href="{{ post.url }}">
        {{ post.title }}
      </a>
    </h1>

    <span class="post-date">{{ post.date | date_to_string }}</span>

    {{ post.content }}
  </div>
  {% endfor %}
</div>

<div class="pagination">
  {% if paginator.next_page %}
    <a class="pagination-item older" href="/page{{paginator.next_page}}">Older</a>
  {% else %}
    <span class="pagination-item older">Older</span>
  {% endif %}
  {% if paginator.previous_page %}
    {% if paginator.page == 2 %}
      <a class="pagination-item newer" href="/">Newer</a>
    {% else %}
      <a class="pagination-item newer" href="/page{{paginator.previous_page}}">Newer</a>
    {% endif %}
  {% else %}
    <span class="pagination-item newer">Newer</span>
  {% endif %}
</div>
