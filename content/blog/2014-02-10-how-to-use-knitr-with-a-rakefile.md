---
date: '2014-02-10T15:05:00+00:00'
slug: 'how-to-use-knitr-with-a-rakefile'
title: How to Use Knitr with a Rakefile
---

I wrote earlier about using [Rakefiles with Pandoc](http://lincolnmullen.com/blog/using-a-global-rakefile-with-pandoc/). Now that I'm engaged in writing a series of posts on historical history dissertations, I have another problem that Rake can solve. Each of the blog posts for the series is written in [R Markdown](http://www.rstudio.com/ide/docs/authoring/using_markdown?version=0.98.501&mode=desktop), which lets me mix R code and writing together in a kind of [literate programming](http://en.wikipedia.org/wiki/Literate_programming) or research notebook. Then the fabulous [knitr package](http://yihui.name/knitr/) by [Yihui Xie](http://yihui.name/) runs the code to create a combination of writing and output in a Markdown file.

Let's assume that you have a set of R Markdown files in your project directory, and that you want to knit them without loading any `.Rdata` or `.Rprofile` files so that your research is reproducible. You can do this by with a Rakefile like this:
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="c1"># Find the source files and generate list of corresponding output files&lt;/span>
&lt;span class="no">KNITR_FILES&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">FileList&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"*.rmd"&lt;/span>&lt;span class="p">]&lt;/span>
&lt;span class="no">OUTPUT_MDS&lt;/span>  &lt;span class="o">=&lt;/span> &lt;span class="no">KNITR_FILES&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">ext&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">".md"&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="n">rule&lt;/span> &lt;span class="s2">".md"&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="s2">".rmd"&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="sx">%[Rscript --vanilla -e "library(knitr); knit(&#39;#{t.source}&#39;);"]&lt;/span>
&lt;span class="k">end&lt;/span>

&lt;span class="n">desc&lt;/span> &lt;span class="s2">"Run knitr on all the analysis files"&lt;/span>
&lt;span class="n">task&lt;/span> &lt;span class="ss">:analysis&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="no">OUTPUT_MDS&lt;/span></code></pre>
</figure>
I have some additional options for this kind of project Rakefile, which you can [see here](https://github.com/lmullen/dissertations-data/blob/master/Rakefile). Mostly they run data munging scripts and define cleaning tasks.
