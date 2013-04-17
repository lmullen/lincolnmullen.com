---
layout: post
title: "Two Computers, One Keyboard and Mouse with Synergy"
date: 2013-03-17 12:00
comments: true
categories: 
- ProfHacker
tags: 
- Synergy
- hacks
- ProfHacker
---

*[This post originally appeared at [ProfHacker][].]*

The odds are good that there are multiple computers in your life. You
might have a laptop along with a desktop computer in your office or at
home. Or (like me) you might have an outmoded computer at home that
could still be useful but which isn’t your primary machine. Or maybe you
run a NASA launch center, or wish you did. If you have multiple
computers, chances are you’d like to use them at the same time. For me,
when I recently built a [standing][] [desk][] in hopes of avoiding undue
[health risks][], it made sense to make a space both for my laptop and
extra monitor *and* for an old Ubuntu desktop. (Yes, it’s pretty nice
screen-real-estate-wise, if not good-taste-in-furniture-wise.)

The big problem\* with using multiple computers  is switching between
keyboards and mice. This will drive you nuts about 10 seconds after you
try to do it. That’s where [Synergy][] comes in.

[Synergy][] is open-source software that lets you share a single
keyboard and mouse on multiple computers. You install the software on
both of your computers (actually, as many computers as you want). One of
the computers is set up to be a server, which shares its keyboard and
mouse with the other computers, which are set up as clients. Synergy
works with Windows, Macs and Linux, and it has a settings manager which
is fairly straightforward. You’ll have to figure out either the names of
your computers on the network (e.g., mullen-mac.local) or your local IP
address, but Google can help you with that.

![Synergy settings][]

Once you have [Synergy][] set up, you can move the mouse from one
computer to the next. The keyboard control follows the mouse, so that
you can type on either machine. It’s not flawless,\*\* but I’ve found
that it works well. You’ll quickly realize that your computers use
different shortcuts (e.g., copying is Ctrl+C on Windows and Linux but
Cmd+C on Macs) but you can change your settings if you really want to go
all in.

As a bonus, Synergy shares the system clipboard between the two
machines. Copy on one computer, paste on the other. You won’t even
realize you’re doing it.

**Do you use multiple computers at the same time? What software or
hardware do you use?**

\* The other big problem with using two computers at once is
synchronizing files, but ProfHacker favorites [Dropbox][] and [Spider
Oak][] have solved that.

\*\* If you have trouble getting the scroll wheel to work properly, try
[this fix][].

  [ProfHacker]: http://chronicle.com/blogs/profhacker/two-computers-one-keyboard-and-mouse-with-synergy/46943
  [standing]: http://chronicle.com/blogs/profhacker/the-portable-ninja-standing-desk/45197
  [desk]: http://chronicle.com/blogs/profhacker/stand-in-the-place-where-you-work-geekdesk-max-review/46257
  [health risks]: http://chronicle.com/blogs/profhacker/stand-up-in-your-office/36054
  [Synergy]: http://synergy-foss.org/
  [Synergy settings]: http://chronicle.com/blogs/profhacker/files/2013/02/Screen-Shot-2013-02-28-at-8.54.22-PM.png
  [Dropbox]: http://db.tt/f5srxqL
  [Spider Oak]: https://spideroak.com/download/referral/525c5656d1c56eccf12b5a466f7553ca
  [this fix]: https://code.google.com/p/synergy-plus/issues/detail?id=28#c17
