---
categories:
- rstats
date: '2014-09-20 14:03'
layout: post
title: A Vagrant Development Environment for R
...

Most of the time when I'm working in R I'm using Mac OS X and I have a
bunch of packages installed. But often I want to run my R code in a
clean environment, and when I'm developing a package I want to test it
on a Linux instance. The combination of [Vagrant][] and [VirtualBox][]
makes it easy to write a script which spools up a virtual machine for
this purpose. I've created a [set of files for Vagrant][] that created
an Ubuntu 14.04 virtual machine, provision it with all the development
tools I need, then install commonly used R packages. Creating a new
development environment is as simple as cloning the repository and
running `vagrant up`. You can get [these scripts at GitHub][set of files
for Vagrant].

Every so often, I also want to spool up a server at Amazon EC2 with far
more RAM and processing power than I have on my laptop. While the
`Vagrantfile` is unnecessary for this, the files `bootstrap.sh` and
`install-r-packages.R` can be adapted with just a little modification to
provision an EC2 instance.

  [Vagrant]: https://www.vagrantup.com/
  [VirtualBox]: https://www.virtualbox.org/
  [set of files for Vagrant]: https://github.com/lmullen/vagrant-r-dev
