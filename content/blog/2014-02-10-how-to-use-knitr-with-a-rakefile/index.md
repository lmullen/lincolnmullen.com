---
title: 'How to Use Knitr with a Rakefile'
date: "2014-02-10"
slug: 'how-to-use-knitr-with-a-rakefile'
tags:
- technology
- R
draft: false
---

I wrote earlier about using [Rakefiles with Pandoc](http://lincolnmullen.com/blog/using-a-global-rakefile-with-pandoc/). Now that I'm engaged in writing a series of posts on historical history dissertations, I have another problem that Rake can solve. Each of the blog posts for the series is written in [R Markdown](http://www.rstudio.com/ide/docs/authoring/using_markdown?version=0.98.501&mode=desktop), which lets me mix R code and writing together in a kind of [literate programming](http://en.wikipedia.org/wiki/Literate_programming) or research notebook. Then the fabulous [knitr package](http://yihui.name/knitr/) by [Yihui Xie](http://yihui.name/) runs the code to create a combination of writing and output in a Markdown file.

Let's assume that you have a set of R Markdown files in your project directory, and that you want to knit them without loading any `.Rdata` or `.Rprofile` files so that your research is reproducible. You can do this by with a Rakefile like this:

```ruby
# Find the source files and generate list of corresponding output files
KNITR_FILES = FileList["*.rmd"]
OUTPUT_MDS  = KNITR_FILES.ext(".md")

rule ".md" => ".rmd" do |t|
  sh %[Rscript --vanilla -e "library(knitr); knit('#{t.source}');"]
end

desc "Run knitr on all the analysis files"
task :analysis => OUTPUT_MDS
```

I have some additional options for this kind of project Rakefile, which you can [see here](https://github.com/lmullen/dissertations-data/blob/master/Rakefile). Mostly they run data munging scripts and define cleaning tasks.
