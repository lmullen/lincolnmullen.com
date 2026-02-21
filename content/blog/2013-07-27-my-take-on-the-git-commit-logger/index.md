---
title: 'My Take on the Git Commit Logger'
date: "2013-07-27"
slug: 'my-take-on-the-git-commit-logger'
description: "A post-commit Git hook that appends every commit message to a daily log—a lightweight alternative to Brett Terpstra's cron-based commit logger for tracking daily progress."
tags:
- technology
- git
draft: false
---

Almost all the work that I do on my computer that is worth doing goes through a Git repository. My Git commits are therefore a useful indication of what I've accomplished each day, except they are scattered across many directories.

When I saw that [Brett Terpstra](http://brettterpstra.com/) created [a script to make a log](http://brettterpstra.com/2012/05/21/git-logger-revisited/) of all his Git commits, I saw the solution to my problem. Terpstra has an implementation that scans a list of repositories once per day, filters certain kinds of commits, and exports them in different kinds of formats. That's more power than I needed, so I wrote my own script. The main difference is that where Terpstra's script is run by cron once per day and goes out to fetch commits, my script is run after every commit and pushes the information to the log.

The script is run using the magic of Git hooks, in this case, post-commit. Any script named `post-commit` in the `.git/hooks` directory will be called after each commit.

Here is the script (also [as a Gist](https://gist.github.com/lmullen/6095650)):

```ruby
#!/usr/bin/env ruby

# Write git commit messages to a log file
#
# Lincoln A. Mullen | lincoln@lincolnmullen.com | http://lincolnmullen.com
# MIT License <http://lmullen.mit-license.org/>
#
# You will have to install the git gem for this to work:
#   gem install git
#
# Name this file 'post-commit' and drop it in the directory '.git/hooks'  in
# any repository that you want to log. Make sure the file is executable. You
# can also add this to your git templates, which will put it in every new
# repository or to existing repositories by re-running git init.
#
# A commit message in the log should look this this:
#   2013-07-26-11-25-55 [repository:branch] First line of commit message

require "git"

log_file = "/home/lmullen/todo/time-use.txt"
repo = Git.open(Dir.pwd)
repo_name = Dir.pwd[%r{[\.\-\w]+$}]                       # top directory
date = repo.log.first.date.strftime("%Y-%m-%d-%H-%M-%S")  # format date
message = repo.log.first.message.lines.first.strip        # just first line
branch = repo.log.first.name
File.open(log_file, "a") do |log|
  log.puts "#{date} [#{repo_name}:#{branch}] #{message}"
end
```

You can drop that following script into the hooks directory of each repository that you want to watch. Or, learning from [Tim Pope](http://tbaggery.com/2011/08/08/effortless-ctags-with-git.html), you can create a Git templates directory which will be added to repositories when you run `git init`.

```bash
git config --global init.templatedir '~/.git_template'
mkdir -p ~/.git_template/hooks
```
