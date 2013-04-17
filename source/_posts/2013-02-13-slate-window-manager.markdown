---
layout: post
title: "The Slate Window Manager for Mac Power Users"
date: 2013-02-13 12:00
comments: true
categories: 
- ProfHacker
tags: 
- ProfHacker
- Mac
- window manager
- hack
---

*[This post originally appeared at [ProfHacker][].]*

Your computer’s user interface is based on one metaphor or another. (At
least, it is if you’re not typing just 1s and 0s into your computer.)
More than likely your computer operates on the “windows” metaphor, even
if you use a Mac. Documents and applications float across the screen,
and you click and drag to arrange the windows on your screen. The
trouble with the window metaphor is that every second spent arranging
windows is a waste of time. I find this to be a frequent source of
frustration, and neither Windows nor Mac OS X handles it well. (The
exception is [xmonad][], a tiling windows manager that is a true thing
of beauty, but not one that you can use without Linux.) That’s why
George wrote about [Better Snap Tool][], which Brian [also likes][] in
addition to [Divvy][], both of which add more powerful features to
manage windows on a Mac. Either is a good, user-friendly option.

If you want a Mac windows manager that is more powerful and much more
customizable, you might want to check out [Slate][]. Slate is free and
open-source, and it has all of the features of Divvy or Better Snap
Tool, though implemented somewhat differently. What makes it stand out
is that you can customize it with a Unix-style configuration file. I’ve
used Slate to set up these features:

-   Keyboard shortcuts to resize windows and move them to various parts
    of the screen.
-   Keyboard shortcuts to move windows from one screen to another.
-   Keyboard shortcuts to open focus on particular apps (e.g., OPT +
    SHIFT + Z focuses on my text editor).
-   Layouts for one monitor (laptop screen) and two monitors (laptop
    screen + external monitor) that rearrange all the windows to
    predetermined positions. These are activated whenever I plug in or
    unplug a screen, or when I press a particular keystroke.

You’ll have to be comfortable editing the configuration file to use
Slate, though the [documentation][Slate] is thorough, and [this blog
post][] offers a helpful introduction. You can also take a look at my
[configuration file][].

**Have you used Slate, Better Snap Tool, Divvy, or another window
manager? What are your favorite features? What tips do you have to
share? Please let us know in the comments.**

  [ProfHacker]: http://chronicle.com/blogs/profhacker/the-slate-window-manager-for-mac-power-users/46203
  [xmonad]: http://xmonad.org/
  [Better Snap Tool]: http://chronicle.com/blogs/profhacker/a-tool-to-manage-your-on-screen-windows/33188
  [also likes]: http://chronicle.com/blogs/profhacker/resizing-and-rearranging-your-windows-with-two-keys/39500
  [Divvy]: http://chronicle.com/blogs/profhacker/divide-and-conquer-your-screen-with-divvy/38877
  [Slate]: https://github.com/jigish/slate
  [this blog post]: http://thume.ca/howto/2012/11/19/using-slate/
  [configuration file]: https://gist.github.com/lmullen/4736568
