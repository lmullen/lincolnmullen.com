---
title: "Writing power tools: text editors"
date: "2012-03-12"
slug: "writing-power-tools-text-editors"
description: "Writing in plain text means you can use a programmer's text editor like Vim, with powerful features for moving through and transforming prose."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/writing-power-tools-text-editors/38940"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="screen-shot-2012-03-10-at-3.22.59-pm-300x195.png" alt="Vim screen shot" >}}

One of the recent [themes](http://chronicle.com/blogs/profhacker/profhacking-the-ultimate-word-processor/33369) on ProfHacker has been the virtues of using [plain](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295) [text](http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700) to do your work, especially your writing.  For me there is one big advantage: I write in plain text so that I can use a text editor to do my writing.

<!--more-->

By [text editor](http://en.wikipedia.org/wiki/Text_editor), I mean a program that is meant to edit code—the type of tool that programmers use. My text editor of choice is [Vim](http://www.vim.org/), but there are many similar programs: [BBEdit](http://www.barebones.com/products/bbedit/index.html), [Emacs](http://www.gnu.org/software/emacs/), [Notepad++](http://notepad-plus-plus.org/), [TextMate](http://macromates.com/), [TextWrangler](http://www.barebones.com/products/textwrangler/), and [SciTE](http://www.scintilla.org/SciTE.html) are just a few of the options. These all share features that are intended to help you write code. They are not primarily software for word processing, like Microsoft Word or [Libre Office](http://www.libreoffice.org/).

So why would someone want to write using software that is meant for writing code rather than writing words? I’ve found that thinking about my words as if they were code has improved my process of writing. There are the specific reasons:

1. *Text editors let me leave comments in my writing.* If you look at well written source code for a computer program, you’ll find many comments that do not affect how the program is run, but which do explain how it works. In the same way, I leave comments to myself about my writing. If I leave a gap between sections, a comment reminds me what remains to be written. I leave comments evaluating paragraphs or sections: “rework this paragraph to take into account argument X” and the like. The most important use I have for the comment function is cutting text. Most writers have a hard time bringing themselves to cut the words over which they’ve labored. Instead, I comment out text which I need to cut. For example, when writing this post the first reason I listed was my irrational hatred of MS Word ever since it killed my puppy, which I got the satisfaction of writing and you got the satisfaction of not reading, because it’s now a comment. This technique has the twin virtues of sparing me the psychic pain of deleting text, and of leaving the text in its context in case I later want to revise it.
2. *Text editors use color to highlight syntax.* Text editors color code the syntax of programing languages: one color for a function, another for a variable, and so on. This feature can be very powerful for writing. For example, I wrote this post in [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295). The syntax of various parts of the document were highlighted for me as I wrote, such as the links, the ordered list, and the headings. My text editor also highlights the words I search for. Which brings us to the next point …
3. *Text editors have amazing find and replace*. Any useful text editor has some version of regular expressions built in. You’ll have to read Konrad’s excellent posts on [regular](http://chronicle.com/blogs/profhacker/reformatting-confucius-with-regular-expressions/38596) [expressions](http://chronicle.com/blogs/profhacker/finding-the-women-of-heimskringla-with-regular-expressions/38631) to learn why these are useful, but text editors have an ability to search that far exceeds that of any word processor.
4. *Text editors keep your fingers on the keyboard, because they have good keyboard shortcuts.* In a word processor, if you want to move to the start of the next sentence, you’ll have to take your hand off the keyboard, use the mouse to click on the start of the next sentence, then put your hand back on the keyboard. In Vim, I type one character ‘)’, which takes me to the next sentence without ever taking my hands off the keyboard. This doesn’t seem like it would matter until you get used to it. Now when I do have to use Word, I feel like I’m running through water.
5. *Text editors are more stable.* Text editors are the tool programmers use most, and unsurprisingly they have taken care of themselves by making text editors highly reliable. For example: I currently have 52 documents open in my text editor (some are writing, a lot are code). But if I were to open even 10 documents in Word, my computer would grind to a halt. I haven’t restarted my text editor in several weeks; it’s more likely that my Mac will crash than that my text editor will.

Using a text editor is not the easiest hack to get started with, since there is a steep learning curve. Text editors are power tools, and they’ll save you lots of work only if you learn to use them properly. If you’re on a Mac, you might try [TextWrangler](http://www.barebones.com/products/textwrangler/), and if you’re on Windows you might try [Notepad++](http://notepad-plus-plus.org/) and using the tutorials they provide. For me, at least, using a text editor has been the most useful hack I’ve ever done.

**Have you used a text editor for writing? Which do you prefer, and why?**
