---
title: 'Setting up a new Mac with dotfiles and Homebrew bundle'
date: 2021-12-27T13:32:26-05:00
slug: 'setting-up-a-new-mac-with-homebrew'
---

I recently had to set up a new Mac for work. Generally speaking, this happens so infrequently that it is worth setting up the new machine from scratch, rather than using Migration Assistant. I like to avoid carrying over the cruft that comes from several years of a constantly updated development environment, and all work files are in iCloud Drive or GitHub anyway. But still, that leaves a fair bit of set up to do to get things working correctly.

For a long time I've kept [my dotfiles in a GitHub repository](https://github.com/lmullen/dotfiles). This sets up configuration for ZSH, Neovim, Go, R, Homebrew, LaTeX, Git, and the like. While a lot of it is Mac specific, the shell and text editor configuration work fine on Linux machines, so I can easily bring settings to servers. This makes customizing my development environment fairly painless. (And Visual Studio Code now has great settings sync, so that takes care of itself.)

Of course not everything can go in a public GitHub repository. Recently, I've taken to having a single file (`~/.env.zsh`), which contains project and service credentials, as well as machine-specific settings, stored as environment variables. For example, all the projects that I create pull their database connection settings from environment variables. And setting the number of cores available on a particular makes scaling up parallel processing easier. This file, like SSH keys, is easy to move over to a new machine.

{{< figure src="env-screenshot.png" caption="Some machine-specific settings from my environment file." >}}

What was new to me this time was using Homebrew bundles for installing software and dependencies. While I've used Homebrew for a long time, I recently learned from a [blog post by Casey Liss](https://www.caseyliss.com/2019/10/8/brew-bundle) that Homebrew has for a while now supported creating a list of software to install. In addition to installing CLIs and other packages from Homebrew proper, and GUI applications as Homebrew Casks, it even supports (though not particularly reliably) installing applications from the Mac App Store.

So I set up a [Brewfile for my work machine](https://github.com/lmullen/dotfiles/blob/main/homebrew/Schaff.Brewfile). This worked great for setting up the new machine, and it is nice to have an explicit record of the software that I need to have installed.
