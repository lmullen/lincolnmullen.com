---
date: '2014-02-05T22:34:00+00:00'
slug: 'using-a-global-rakefile-with-pandoc'
title: Using a Global Rakefile with Pandoc
---

I wrote a post a while ago about [how to use Rake to automate the process of building PDFs using Pandoc](http://lincolnmullen.com/blog/rake-and-pandoc/). The idea is that you drop the Rakefile into each project that has Markdown files, and by running `rake pdfs` you can generate the PDFs without typing long and awkward Pandoc commands. It works great.

But since I'm lazy (in the [sense used by Larry Wall](http://threevirtues.com/)), it's too much work to copy the Rakefile into each project. I just [learned](http://www.stuartellis.eu/articles/rake/) that you can use global Rakefiles, which will then be available in any directory. Here is [how to do it](http://blog.envylabs.com/post/62063421092/rake-global-tasks).

First, create the directory `~/.rake`. Then in that directory create a file `pandoc.rake`. The file should contain something like this:
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="no">MDFILES&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">FileList&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"*.md"&lt;/span>&lt;span class="p">]&lt;/span>
&lt;span class="no">PDFS&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">MDFILES&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">ext&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">".pdf"&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="n">desc&lt;/span> &lt;span class="s2">"Use Pandoc to build PDFs of all Markdown documents"&lt;/span>
&lt;span class="n">task&lt;/span> &lt;span class="ss">:pandoc&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="no">PDFS&lt;/span>

&lt;span class="n">rule&lt;/span> &lt;span class="s2">".pdf"&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="s2">".md"&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="s2">"pandoc &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">source&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> -o &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
&lt;span class="k">end&lt;/span>

&lt;span class="c1"># Only clobber generated PDFs&lt;/span>
&lt;span class="nb">require&lt;/span> &lt;span class="s2">"rake/clean"&lt;/span>
&lt;span class="no">CLOBBER&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">include&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="no">PDFS&lt;/span>&lt;span class="p">)&lt;/span></code></pre>
</figure>
Now in each directory you should be able to run `rake -g pandoc` to generate a PDF from each Markdown file. You can customize the shell command if you have more complicated needs, or you can create a project-specific Rakefile if the project is more complicated. For important projects you should definitely include a Rakefile in the repository, since the Rakefile records how to reproduce your outputs.

But for simple projects this is sufficient.
