---
title: "Fix PDFs quickly with pdftk"
date: "2011-11-17"
slug: "fix-pdfs-quickly-with-pdftk"
description: "The command-line tool pdftk makes it fast to join, split, and rotate the PDFs that scholars work with every day."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/fix-pdfs-quickly-with-pdftk/37332"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="pdftk_heading.gif" alt="pdftk logo" >}}

I work a lot with PDFs. Most of the journal articles I read are PDFs, when I’m lucky drafts of other people’s work come as PDFs rather than Word documents, and I scan much of the source material for my research into PDFs. At ProfHacker, we’ve given you lots of advice about how to hack PDFs: how to [OCR them so they’re searchable](http://chronicle.com/blogs/profhacker/ocr-those-pdfs/25653), how to [convert documents into PDFs](http://chronicle.com/blogs/profhacker/converting-files-with-primo-pdf/30911), how to [use Zotfile with PDFs in your Zotero library](http://chronicle.com/blogs/profhacker/zotfile-awesomifies-zotero-attachments/36650), and how to [organize and annotate your PDFs](http://chronicle.com/blogs/profhacker/how-do-you-organizeannotate-pdfs-reader-response-roundup/24117).

<!--more-->

So what do you do when you need to fix a problem with a PDF? There are a few problems I commonly have with PDFs. Sometimes I need to join several PDFs into one master file, perhaps because I have scanned a long text in smaller chunks, or perhaps because I need to add an appendix to a document. Other times I need to rotate a document because I’ve scanned it upside down. (Someday we will have a ProfHacker post on how to scan documents right side up.)

For any of these tasks I could use Acrobat Pro or Mac OS X’s built-in Preview.app, and doubtless there are dozens of other applications for every platform that would also work. These [GUI](http://en.wikipedia.org/wiki/Graphical_user_interface) applications have an undoubted advantage. I can poke around in the menus and dialog boxes, and two (or ten) minutes later I’ll have fixed the PDF. The trouble is that it doesn’t get any easier or quicker the next time I have to fix a PDF. It always takes [just as much time as the first time](http://lenz.unl.edu/2011/08/05/the-man-finger-aftermath.html).

That’s where the command line tool [`pdftk`](http://www.pdflabs.com/tools/pdftk-the-pdf-toolkit/) (PDF Tool Kit) comes in. As with [other command line tools](http://chronicle.com/blogs/profhacker/tag/the-profhacker-guide-to-the-command-line), it takes some effort to learn what to type at the blinking cursor. But once you’ve learned, you can accomplish complex tasks in the time it takes your computer to open Adobe Acrobat.

So, once you have [installed `pdftk`](http://www.pdflabs.com/docs/install-pdftk/), and fired up your terminal of choice, try performing these tasks, drawn from the `pdftk` [examples](http://www.pdflabs.com/docs/pdftk-cli-examples/) page.

First, let’s join two PDFs together into one file. The two files to be joined are called `1.pdf` and `2.pdf` and the combined file is called `combined.pdf`.

> `pdftk 1.pdf 2.pdf cat output combined.pdf`

This command is essentially a sentence. It tells `pdftk` to take the files `1.pdf` and `2.pdf` and concatenate (i.e., `cat`) them, then `output` the resulting file with the name `combined.pdf`. Easy enough, right?

Now let’s rotate a PDF:

> `pdftk in.pdf cat 1-endS output out.pdf`

This command is another sentence. It tells `pdftk` to take the file `in.pdf`. Next we’re going to do something (`cat`) with all the pages (`1-end`). As you’d expect you could also perform the actions on pages 2-28, and so on. Then we’re going to rotate the pages 180 degrees with the `S` tag. The `S` stands for “south,” and as you’d expect you could rotate the document 90 degrees with `E` and 270 degrees with `W`. Then we `output` the document to `out.pdf`.

Once you learn the very basic language to give `pdftk` commands, you can perform these PDF tasks and [others](http://www.pdflabs.com/docs/pdftk-cli-examples/) very quickly indeed.

**What common fixes do you need to make to PDFs? Have you tried `pdftk`? What other tools should we cover in our [“ProfHacker Guide to the Command Line”](http://chronicle.com/blogs/profhacker/tag/the-profhacker-guide-to-the-command-line)?**
