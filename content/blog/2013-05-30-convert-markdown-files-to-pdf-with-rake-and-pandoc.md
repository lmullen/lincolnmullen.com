---
date: '2013-05-30T19:30:00+00:00'
slug: 'convert-markdown-files-to-pdf-with-rake-and-pandoc'
title: Convert Markdown Files to PDF with Rake and Pandoc
---

A post I wrote last fall about [using Make to convert all the Markdown files in a directory to PDFs](/blog/make-and-pandoc/) continues to draw lots of traffic, so it must be a solution that people find useful. I've since started using Rake instead of Make. You can find any number of [comparisons](http://blog.jgc.org/2010/11/things-make-got-right-and-how-to-make.html) [between](http://www.trottercashion.com/2010/10/29/replacing-make-with-rake.html) Make and and Rake. I use Rake because I like having the power of Ruby, but Make handles dependencies much better than Rake.

Here's is a generic `Rakefile` to make PDFs and DOCXs of every Markdown file in a directory.
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="nb">require&lt;/span> &lt;span class="s2">"rake/clean"&lt;/span>

&lt;span class="c1"># Define inputs and outputs&lt;/span>
&lt;span class="no">MDFILES&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">FileList&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="s2">"*.md"&lt;/span>&lt;span class="p">]&lt;/span>
&lt;span class="no">PDFS&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">MDFILES&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">ext&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">".pdf"&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="no">DOCX&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">MDFILES&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">ext&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">".docx"&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="c1"># Clobber only PDFs and DOCXs we&#39;ve generated&lt;/span>
&lt;span class="no">CLOBBER&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">include&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="no">PDFS&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="no">DOCX&lt;/span>&lt;span class="p">)&lt;/span>

&lt;span class="c1"># Define bibliography and CSL files&lt;/span>
&lt;span class="no">BIB&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"/Users/lmullen/acad/research/bib/master.bib"&lt;/span>
&lt;span class="no">CSL&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"chicago-mullen.csl"&lt;/span>

&lt;span class="n">desc&lt;/span> &lt;span class="s2">"Build all documents in all formats."&lt;/span>
&lt;span class="n">task&lt;/span> &lt;span class="ss">:default&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="p">[&lt;/span>&lt;span class="ss">:pdfs&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="ss">:docx&lt;/span>&lt;span class="p">]&lt;/span>
 
&lt;span class="n">desc&lt;/span> &lt;span class="s2">"Build PDFs of all documents."&lt;/span>
&lt;span class="n">task&lt;/span> &lt;span class="ss">:pdfs&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="no">PDFS&lt;/span>

&lt;span class="n">desc&lt;/span> &lt;span class="s2">"Build DOCXs of all documents."&lt;/span>
&lt;span class="n">task&lt;/span> &lt;span class="ss">:docx&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="no">DOCX&lt;/span>

&lt;span class="c1"># Build PDFs from Markdown source&lt;/span>
&lt;span class="n">rule&lt;/span> &lt;span class="s2">".pdf"&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="s2">".md"&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="s2">"pandoc &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">source&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> -o &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --csl=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">CSL&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --bibliography=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">BIB&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
&lt;span class="k">end&lt;/span>

&lt;span class="c1"># Build DOCXs from Markdown source&lt;/span>
&lt;span class="n">rule&lt;/span> &lt;span class="s2">".docx"&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="s2">".md"&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="s2">"pandoc &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">source&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> -o &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --csl=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">CSL&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --bibliography=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">BIB&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
&lt;span class="k">end&lt;/span></code></pre>
</figure>
If you prefer that your filenames be `document.md.pdf` instead of `document.pdf`, then you have to do something [ugly](http://stackoverflow.com/questions/16122744/rake-does-not-recognize-rules-with-multiple-extensions) and replace the rules above with these rules:
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="c1"># Build PDFs from Markdown source&lt;/span>
&lt;span class="n">rule&lt;/span>&lt;span class="p">(&lt;/span> &lt;span class="sr">/\.md\.pdf$/&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="p">[&lt;/span>
  &lt;span class="nb">proc&lt;/span> &lt;span class="p">{&lt;/span>&lt;span class="o">|&lt;/span>&lt;span class="n">task_name&lt;/span>&lt;span class="o">|&lt;/span> &lt;span class="n">task_name&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">sub&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="sr">/\.md\.pdf$/&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="s1">&#39;.md&#39;&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="p">}&lt;/span>
  &lt;span class="p">])&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="s2">"pandoc &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">source&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> -o &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --csl=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">CSL&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --bibliography=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">BIB&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
&lt;span class="k">end&lt;/span>

&lt;span class="c1"># Build DOCXs from Markdown source&lt;/span>
&lt;span class="n">rule&lt;/span>&lt;span class="p">(&lt;/span> &lt;span class="sr">/\.md\.docx$/&lt;/span> &lt;span class="o">=&gt;&lt;/span> &lt;span class="p">[&lt;/span>
  &lt;span class="nb">proc&lt;/span> &lt;span class="p">{&lt;/span>&lt;span class="o">|&lt;/span>&lt;span class="n">task_name&lt;/span>&lt;span class="o">|&lt;/span> &lt;span class="n">task_name&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">sub&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="sr">/\.md\.docx$/&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="s1">&#39;.md&#39;&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="p">}&lt;/span>
  &lt;span class="p">])&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="o">|&lt;/span>
  &lt;span class="n">sh&lt;/span> &lt;span class="s2">"pandoc &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">source&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> -o &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">t&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --csl=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">CSL&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> --bibliography=&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="no">BIB&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>
&lt;span class="k">end&lt;/span></code></pre>
</figure>
