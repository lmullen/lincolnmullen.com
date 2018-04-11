---
date: '2013-07-27T18:06:00+00:00'
slug: 'my-take-on-the-git-commit-logger'
title: My Take on the Git Commit Logger
---

Almost all the work that I do on my computer that is worth doing goes through a Git repository. My Git commits are therefore a useful indication of what I've accomplished each day, except they are scattered across many directories.

When I saw that [Brett Terpstra](http://brettterpstra.com/) created [a script to make a log](http://brettterpstra.com/2012/05/21/git-logger-revisited/) of all his Git commits, I saw the solution to my problem. Terpstra has an implementation that scans a list of repositories once per day, filters certain kinds of commits, and exports them in different kinds of formats. That's more power than I needed, so I wrote my own script. The main difference is that where Terpstra's script is run by cron once per day and goes out to fetch commits, my script is run after every commit and pushes the information to the log.

The script is run using the magic of Git hooks, in this case, post-commit. Any script named `post-commit` in the `.git/hooks` directory will be called after each commit.

Here is the script (also [as a Gist](https://gist.github.com/lmullen/6095650)):
<figure class="highlight">
<pre><code class="language-ruby" data-lang="ruby">&lt;span class="c1">#!/usr/bin/env ruby&lt;/span>

&lt;span class="c1"># Write git commit messages to a log file&lt;/span>
&lt;span class="c1"># &lt;/span>
&lt;span class="c1"># Lincoln A. Mullen | lincoln@lincolnmullen.com | http://lincolnmullen.com&lt;/span>
&lt;span class="c1"># MIT License &lt;http://lmullen.mit-license.org/&gt;&lt;/span>
&lt;span class="c1">#&lt;/span>
&lt;span class="c1"># You will have to install the git gem for this to work:&lt;/span>
&lt;span class="c1">#   gem install git&lt;/span>
&lt;span class="c1">#&lt;/span>
&lt;span class="c1"># Name this file &#39;post-commit&#39; and drop it in the directory &#39;.git/hooks&#39;  in&lt;/span>
&lt;span class="c1"># any repository that you want to log. Make sure the file is executable. You&lt;/span>
&lt;span class="c1"># can also add this to your git templates, which will put it in every new&lt;/span>
&lt;span class="c1"># repository or to existing repositories by re-running git init.  &lt;/span>
&lt;span class="c1"># &lt;/span>
&lt;span class="c1"># A commit message in the log should look this this:&lt;/span>
&lt;span class="c1">#   2013-07-26-11-25-55 [repository:branch] First line of commit message&lt;/span>

&lt;span class="nb">require&lt;/span> &lt;span class="s2">"git"&lt;/span>

&lt;span class="n">log_file&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="s2">"/home/lmullen/todo/time-use.txt"&lt;/span>
&lt;span class="n">repo&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Git&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">open&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="no">Dir&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">pwd&lt;/span>&lt;span class="p">)&lt;/span>
&lt;span class="n">repo_name&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="no">Dir&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">pwd&lt;/span>&lt;span class="p">[&lt;/span>&lt;span class="sr">%r{[&lt;/span>&lt;span class="se">\.\-\w&lt;/span>&lt;span class="sr">]+$}&lt;/span>&lt;span class="p">]&lt;/span>                       &lt;span class="c1"># top directory&lt;/span>
&lt;span class="n">date&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">repo&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">log&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">first&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">date&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">strftime&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="s2">"%Y-%m-%d-%H-%M-%S"&lt;/span>&lt;span class="p">)&lt;/span>  &lt;span class="c1"># format date&lt;/span>
&lt;span class="n">message&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">repo&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">log&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">first&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">message&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">lines&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">first&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">strip&lt;/span>        &lt;span class="c1"># just first line&lt;/span>
&lt;span class="n">branch&lt;/span> &lt;span class="o">=&lt;/span> &lt;span class="n">repo&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">log&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">first&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">name&lt;/span>
&lt;span class="no">File&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">open&lt;/span>&lt;span class="p">(&lt;/span>&lt;span class="n">log_file&lt;/span>&lt;span class="p">,&lt;/span> &lt;span class="s2">"a"&lt;/span>&lt;span class="p">)&lt;/span> &lt;span class="k">do&lt;/span> &lt;span class="o">|&lt;/span>&lt;span class="n">log&lt;/span>&lt;span class="o">|&lt;/span>                        
  &lt;span class="n">log&lt;/span>&lt;span class="p">.&lt;/span>&lt;span class="nf">puts&lt;/span> &lt;span class="s2">"&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">date&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2"> [&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">repo_name&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">:&lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">branch&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">] &lt;/span>&lt;span class="si">#{&lt;/span>&lt;span class="n">message&lt;/span>&lt;span class="si">}&lt;/span>&lt;span class="s2">"&lt;/span>  
&lt;span class="k">end&lt;/span></code></pre>
</figure>
You can drop that following script into the hooks directory of each repository that you want to watch. Or, learning from [Tim Pope](http://tbaggery.com/2011/08/08/effortless-ctags-with-git.html), you can create a Git templates directory which will be added to repositories when you run `git init`.
<figure class="highlight">
<pre><code class="language-bash" data-lang="bash">git config --global init.templatedir &lt;span class="s1">&#39;~/.git_template&#39;&lt;/span>
mkdir -p ~/.git_template/hooks</code></pre>
</figure>
