---
title: 'Convert Markdown files to PDF with Rake and Pandoc'
date: "2013-05-30"
slug: 'convert-markdown-files-to-pdf-with-rake-and-pandoc'
description: "A Rakefile that converts all Markdown files in a directory to PDFs and DOCXs using Pandoc—a Ruby-powered upgrade to the earlier Make-based workflow."
tags:
- technology
draft: false
---

A post I wrote last fall about [using Make to convert all the Markdown files in a directory to PDFs](/blog/make-and-pandoc/) continues to draw lots of traffic, so it must be a solution that people find useful. I've since started using Rake instead of Make. You can find any number of [comparisons](http://blog.jgc.org/2010/11/things-make-got-right-and-how-to-make.html) [between](http://www.trottercashion.com/2010/10/29/replacing-make-with-rake.html) Make and and Rake. I use Rake because I like having the power of Ruby, but Make handles dependencies much better than Rake.

Here's is a generic `Rakefile` to make PDFs and DOCXs of every Markdown file in a directory.

```ruby
require "rake/clean"

# Define inputs and outputs
MDFILES = FileList["*.md"]
PDFS = MDFILES.ext(".pdf")
DOCX = MDFILES.ext(".docx")

# Clobber only PDFs and DOCXs we've generated
CLOBBER.include(PDFS, DOCX)

# Define bibliography and CSL files
BIB = "/Users/lmullen/acad/research/bib/master.bib"
CSL = "chicago-mullen.csl"

desc "Build all documents in all formats."
task :default => [:pdfs, :docx]

desc "Build PDFs of all documents."
task :pdfs => PDFS

desc "Build DOCXs of all documents."
task :docx => DOCX

# Build PDFs from Markdown source
rule ".pdf" => ".md" do |t|
  sh "pandoc #{t.source} -o #{t.name} --csl=#{CSL} --bibliography=#{BIB}"
end

# Build DOCXs from Markdown source
rule ".docx" => ".md" do |t|
  sh "pandoc #{t.source} -o #{t.name} --csl=#{CSL} --bibliography=#{BIB}"
end
```

If you prefer that your filenames be `document.md.pdf` instead of `document.pdf`, then you have to do something [ugly](http://stackoverflow.com/questions/16122744/rake-does-not-recognize-rules-with-multiple-extensions) and replace the rules above with these rules:

```ruby
# Build PDFs from Markdown source
rule( /\.md\.pdf$/ => [
  proc {|task_name| task_name.sub(/\.md\.pdf$/, '.md') }
  ]) do |t|
  sh "pandoc #{t.source} -o #{t.name} --csl=#{CSL} --bibliography=#{BIB}"
end

# Build DOCXs from Markdown source
rule( /\.md\.docx$/ => [
  proc {|task_name| task_name.sub(/\.md\.docx$/, '.md') }
  ]) do |t|
  sh "pandoc #{t.source} -o #{t.name} --csl=#{CSL} --bibliography=#{BIB}"
end
```
