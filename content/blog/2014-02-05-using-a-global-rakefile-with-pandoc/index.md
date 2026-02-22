---
title: 'Using a global Rakefile with Pandoc'
date: "2014-02-05"
slug: 'using-a-global-rakefile-with-pandoc'
description: "A global Rakefile in ~/.rake makes Pandoc PDF generation available in any directory without copying boilerplate into every project folder."
tags:
- Rake
- Pandoc
draft: false
---

I wrote a post a while ago about [how to use Rake to automate the process of building PDFs using Pandoc](http://lincolnmullen.com/blog/rake-and-pandoc/). The idea is that you drop the Rakefile into each project that has Markdown files, and by running `rake pdfs` you can generate the PDFs without typing long and awkward Pandoc commands. It works great.

But since I'm lazy (in the [sense used by Larry Wall](http://threevirtues.com/)), it's too much work to copy the Rakefile into each project. I just [learned](http://www.stuartellis.eu/articles/rake/) that you can use global Rakefiles, which will then be available in any directory. Here is [how to do it](http://blog.envylabs.com/post/62063421092/rake-global-tasks).

First, create the directory `~/.rake`. Then in that directory create a file `pandoc.rake`. The file should contain something like this:

```ruby
MDFILES = FileList["*.md"]
PDFS = MDFILES.ext(".pdf")

desc "Use Pandoc to build PDFs of all Markdown documents"
task :pandoc => PDFS

rule ".pdf" => ".md" do |t|
  sh "pandoc #{t.source} -o #{t.name}"
end

# Only clobber generated PDFs
require "rake/clean"
CLOBBER.include(PDFS)
```

Now in each directory you should be able to run `rake -g pandoc` to generate a PDF from each Markdown file. You can customize the shell command if you have more complicated needs, or you can create a project-specific Rakefile if the project is more complicated. For important projects you should definitely include a Rakefile in the repository, since the Rakefile records how to reproduce your outputs.

But for simple projects this is sufficient.
