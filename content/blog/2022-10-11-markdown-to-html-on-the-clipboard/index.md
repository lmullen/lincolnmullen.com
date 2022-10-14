---
title: 'Markdown to HTML on the clipboard'
date: 2022-10-11T00:39:56-04:00
slug: 'markdown-to-html-on-the-clipboard'
draft: false
---

I often have small snippets of Markdown that I want to copy to the clipboard and then paste as HTML. I thought about writing an extension for Visual Studio Code, or a custom script for [Boop](https://github.com/IvanMathy/Boop). But that seemed like a lot of work for a simple task. And then I remembered: Unix.

```
pbpaste | pandoc | pbcopy
```

There is a one-liner which will work on a Mac to paste Markdown into [Pandoc](https://pandoc.org/) and then copy the resulting HTML back to the clipboard.

Can't get much simpler than that.

If you want do go from HTML to Markdown, the one-liner is a little longer:

```
pbpaste | pandoc -f html -t markdown | pbcopy
```
