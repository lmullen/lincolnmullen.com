---
title: 'TEI Boilerplate for getting started with TEI'
date: "2013-02-05"
slug: 'tei-boilerplate-for-getting-started-with-tei'
description: "TEI Boilerplate uses XSL and jQuery to render TEI files directly in the browser—a much lower barrier to entry than learning XSLT just to see your encoded text."
tags:
- digital humanities
- TEI
draft: false
---

I sat in on [John Unsworth's](https://twitter.com/unsworth) digital humanities class on [TEI](http://www.tei-c.org/index.xml) this morning. I'd investigated TEI before, but TEI was frustrating because of the difficulty of transforming the TEI into something immediately visible and useful to check my work. Last night I discovered [TEI Boilerplate](http://dcl.slis.indiana.edu/teibp/), a basic XSL stylesheet to transform a TEI file so it is readable in a browser, with accompanying CSS and Javascript. The idea is to let you use your jQuery skills instead of learning XSLT:

> But the power of XSLT comes at a cost. XSLT is a far more complex language than CSS and has a far steeper learning curve. There are relatively few expert XSLT programmers compared to the numbers of expert Web developers with extensive knowledge of HTML/CSS/JavaScript.

> On the other hand, JavaScript is an extremely capable language for doing fancy things with XML content. JavaScript frameworks like JQuery and EXT JS, and the major browsers provide excellent support for manipulating the DOM (Document Object Model) of the web page. With TEI Boilerplate, the TEI document becomes part of the DOM, and JavaScript may then be used to perform more complex transformations, movements, and other manipulations of the TEI content than can be achieved with CSS alone.

By the way, you might spend an hour in frustration, including some ill-advised messing around in the guts of your system, because your browser simply will not display the TEI file when it's hosted locally, though it will display it when it's hosted on a server. It turns out that Chrome [disables](http://stackoverflow.com/a/6251757) local XSL transformations for security reasons, so just use another browser. I suppose you could also put the files on a locally hosted [XAMPP](http://www.apachefriends.org/en/xampp.html) server.

Before I downloaded it several times, TEI Boilerplate had only [8 downloads](http://sourceforge.net/projects/teiboilerplate/files/) for the current version. Can this really be?
