---
layout: post
title: "Cleanup your Markdown text in Vim with equalprg"
date: 2012-12-07 15:05
comments: true
categories: 
- hacks
tags:
- Vim
- Pandoc
- hacks
- writing
---

[Jason Heppler][], [Caleb McDaniel][], and [Chad Black][] often post
Vim hacks, so here is one of mine.

One of the downsides to writing prose in Markdown is that the formatting
of the text can become quite messy, with lines of varying lengths,
footnotes scattered throughout the document with incorrect numbering,
links in the text, and so on. For files like HTML or Javascript, Vim can
automatically indent text using the `=` command over a range or motion,
or the `==` command on a single line. By setting the `equalprg`
variable, you can clean up your Markdown text using an external program.

<!--more-->

First, add this line to a `~/.vim/ftplugin/pandoc.vim` file. This code 
tells Vim to use `=` to call Pandoc as a filter to format the text.

{% codeblock %}
setlocal equalprg=pandoc\ -t\ markdown\ --reference-links\ --atx-headers
{% endcodeblock %}

Now you can select a range of text and press `=`, or go to the start of 
the document and press `=G` to filter the whole document.

I've also added a keymapping to filter the whole document automatically. 
You can put this line in the same file as above, or in your `.vimrc`.

{% codeblock %}
nnoremap <silent> <leader>= mpgg/\v^(\w\|\#)<CR>=G`p :let @/ = ""<CR>
{% endcodeblock %}

That line drops a mark at your current cursor location, goes to the
beginning of the document, finds the first line that begins with a word
or a header (to avoid the Pandoc title block), filters the text through
Pandoc to the end of the document, return the cursor to the mark, and
clears the search highlighting.

I tend to clean up the text after every paragraph I write. One caution:
if Pandoc thinks your markup is invalid, you can lose text, so be
careful. Also, keep your documents under version control.

  [Jason Heppler]: http://jasonheppler.org
  [Caleb McDaniel]: http://wcm1.web.rice.edu/hacks.html
  [Chad Black]: http://parezcoydigo.wordpress.com/
