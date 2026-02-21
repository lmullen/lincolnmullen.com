---
title: 'Writing in Vim with placeholders and quickfix'
date: "2012-12-08"
slug: 'writing-in-vim-with-placeholders-and-quickfix'
description: "A Vim workflow for keeping up writing momentum: insert standard placeholder markers like {TODO} and {PAGE}, highlight them in red, then collect them all in the quickfix list for cleanup."
tags:
- Vim
draft: false
---

In his famous [essay](http://www.wiu.edu/cas/history/pdf/HistoryasaLiteraryArt.pdf) on "History as a Literary Art," Samuel Eliot Morrison advises, "When you once are writing, go on writing as long as you can; there will be plenty of time later to shove in the footnotes or return to the library for extra information," and, "Don't stop to consult your notes for every clause or sentence: it is better to get what you have to say clearly in your mind and dash it off; then, after you have it down, return to your notes and compose your next few pages or paragraphs."

I try to follow this advice when I'm writing by leaving notes or markers in the text when I'm missing a fact or a reference. Following up on my [post](http://lincolnmullen.com/blog/cleanup-your-markdown-text-in-vim-with-equalprg/) about Vim yesterday, here is how I manage these markers in Vim.

First, I insert the markers in a standard form: `{TODO}` is my generic marker; `{PAGE}` marks a missing reference; `{FACT}` marks a missing date or the like. I insert these markers using [vim-snipmate](https://github.com/garbas/vim-snipmate) but any method will do.

Second, I use the line below to highlight these markers in the text. I have the line in `~/.vim/after/syntax/pandoc.vim`, but you could also put it in your `.vimrc`.

```vim
match Error '{\w\+}'
```

The result looks like this:

{{< figure src="2012-12-08-marker.png" >}}

These markers then come in handy when it's time to clean up the draft. I could search for the markers, but this keymapping makes things easier:

```vim
nnoremap <leader>{ :vimgrep /{\w\+}/ %<CR>:copen<CR>
```

That keymapping puts all of the markers into Vim's [quickfix](http://vimdoc.sourceforge.net/htmldoc/quickfix.html) buffer, which gives me a list of all the needed corrections and lets me jump from item to item. It looks like this:

{{< figure src="2012-12-08-quickfix.png" >}}
