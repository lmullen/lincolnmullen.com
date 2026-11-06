---
title: "Getting comfortable on the command line"
date: "2011-10-24"
slug: "getting-comfortable-on-the-command-line"
description: "A gentle introduction to the most basic commands for moving around the command line, for scholars facing a blinking cursor for the first time."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/getting-comfortable-on-the-command-line/36862"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="screen-shot-2011-10-24-at-12.59.41-am.png" alt="A terminal session running pwd, cd, ls, touch, open, and less" >}}

Last month we started a new series: The ProfHacker Guide to the Command Line. The first post tried to explain [why you might want to use the command line](http://chronicle.com/blogs/profhacker/the-profhacker-guide-to-the-command-line/36125) in the first place. Tomorrow Konrad will introduce you to `sort`, the first command line tool we will cover that can be of use to scholars. Today I want to introduce you to the most basic commands for getting around the command line.

<!--more-->

The idea of trying out these few simple commands is to give you a gentle introduction to the [command line interface](http://en.wikipedia.org/wiki/Command-line_interface) (CLI). One problem in starting out with the CLI is that the blinking cursor on the blank screen gives you absolutely no idea of what to do next. These suggestions should help. Another common worry I hear from people who don’t use the CLI is that they will destroy their computer. While it is probably true that if you set the proverbial monkey to typing at the command line, it would delete your file system before it wrote the works of Shakespeare, a human user will be perfectly safe using these commands. So take fifteen minutes and actually try these commands on your computer, and you’ll start to feel more comfortable on the CLI.

I’ve included [\*nix](http://en.wikipedia.org/wiki/Unix-like) commands for Unix, Linux, and Mac OS X users, and separate DOS commands for Windows users. You should type in the example commands below verbatim, then press enter on your keyboard. Clicking the image at the start of this post will show you a screen shot of all the commands in this post.

### Finding Your Command Line

The first step is to find your CLI. On Mac OS X, open the applications folder and launch Terminal.app. On a Linux distribution like Ubuntu, search your dash for “Terminal,” or look for “Terminal” in your applications menu under “Accessories” or “System.” On a Windows PC, click “Start,” then “All Programs,” then “Accessories,” then “Command Prompt,” or click “Start” and search for “Command Prompt.”

### Where Am I? Where Am I Going?

Once you’ve got your CLI up and running, we can begin with some simple commands. Your command line is working in a particular directory (or folder) just like a window on a Mac or a PC is open to a particular directory. Let’s find out where we are with this command:

For \*nix (command 1):

> `pwd`

For Windows (command 1):

> `cd`

You should get output that looks something like this:

> `/Users/lmullen`

This output tell us that we’re in my folder `lmullen` within the `Users` folder. Let’s change our working directory to be the Desktop folder (surprise! your desktop is just a directory like any other), so we that the things we do on the CLI will appear magically before our eyes. Type this command:

For \*nix (command 2):

> `cd ~/Desktop/`

For Windows (command 2):

> `cd Desktop`

Now you should be in the `Desktop` directory. You can test this by running command 1 again. Now that we’re here, let’s see what files are stored in this directory. Try this command:

For \*nix (command 3):

> `ls`

For Windows (command 3):

> `dir`

The output from that command should list all the files on your desktop, which you can verify by checking what is on your desktop.

### Let’s Make a Text File

Now that we know where we are and how to check which files are stored in a directory, let’s add a file. This is the command to create a blank text file:

For \*nix (command 4):

> `touch test.txt`

For Windows (command 4):

> `copy NUL test.txt`

Now if you run command 3 again, you should see the file `test.txt` where it wasn’t before. You should also be able to look at your desktop and see the new file.

A blank text file won’t do us much good, so let’s open the text file so we can edit it.

For \*nix (command 5):

> `open test.txt`

For Windows (command 5):

> `notepad test.txt`

On a Mac this command will open the text file in the default text editor (probably TextEdit); on a Windows PC, the command will open the file in Notepad. In other words, the file will be opened just as if you had double-clicked on it in the operating system. Feel free to add a few lines of text to the file, then save it.

Now that you’ve added some text to the file, here is a bonus command for \*nix users only. Let’s look at the text in the file directly from the command line. Type this command:

For \*nix (command 6):

> `less test.txt`

Now you should see the text that you saved in the file displayed in the same window as your command line. When you’re done, press `q` to quit the program `less`.

Congratulations! You’ve managed to navigate the command line and create a text file without breaking your computer. Now you’re ready to learn specific commands that will do useful things for you.

**Did you try your hand at the command line? Let us know how it went in the comments.**
