---
date: '2014-02-05T20:03:00+00:00'
slug: 'ghostery-on-your-machine-stop-sites-from-tracking-you'
title: 'Ghostery on Your Machine: Stop Sites from Tracking You'
---

*\[This is a preview of a post that will appear at [ProfHacker](http://chronicle.com/blogs/profhacker/) next week. I've been on a long ProfHacker hiatus to write other things, but Internet privacy is sufficiently important to me to make an exception.\]*

Suppose that you visited the website of the *[New York Times](http://nytimes.com)* this morning. If you were me, then you likely skimmed past ads for eyeglasses, an advertorial for Dell, a link to graduate programs at Columbia, and innumerable ads for home delivery of the *Times*. Those ads were targeted at me: I recently bought glasses, I recently looked for some computer hardware, and I'm a graduate student. Your ads were probably different, but they were almost certainly targeted at you in some way.

What's going on here? The *Times*, like every other website in the world, is tracking who you are. Websites definitely track what you do on their site: just think of the customized recommendations on Amazon's website. (Why yes, I do think I'll buy the next volume of the *Summa Theologica*.) But websites are also tracking where you go and what you do on other websites, *whether or not you give them permission to do so*. The *Times* probably knows my shopping patterns, if not as well as Amazon, then well enough to make reasonable guesses about what I might be interested in. The *Times* is an entirely reputable website; there plenty of disreputable websites that are doing the same thing.

How do websites keep track of who you are? One way is through cookies, small files of information that keep information about you that persist across visits. Another way is with the IP address that you are browsing from. It's even possible to figure out who people are by testing to see whether their browsers are uniquely configured. The Electronic Frontier Foundation has created [Panopticlick](https://panopticlick.eff.org) to test whether your browser uniquely identifies you. When I ran the test, my browser "appears to be unique among the 3,836,733." Many websites embed Javascript trackers for analytics. Other websites embed Flash videos (like ads) or use invisible images a pixel big to see whether you've visited the site. You can be certain that you are being tracked. Whether or not websites know that you are "Lincoln Mullen" or a long string of letters and numbers doesn't really matter much.

What can be done? In theory, modern browsers can send a piece of information (called an HTTP-Header) to request a website not to track the user. You can read about how to [set up your browser to use "do not track" here](http://donottrack.us/). But nothing obligates websites to honor the request, and most do not.

You need something more powerful: a tool where you make the decisions rather than the website. I suggest you try [Ghostery](http://www.ghostery.com/).

\[//www.youtube-nocookie.com/embed/EKzyifAvC\_U?rel=0\]

[Ghostery](http://www.ghostery.com/) is a plugin for Chrome, Firefox, Safari, Internet Explorer and Opera, which you can install from its website. (See the introductory video above.) Ghostery keeps a database of all the methods of tracking people on the Internet. The database is up to 1840 trackers so far. When you visit a website, the plugin pops up a small, unobtrusive window that tells you how which trackers are active on a site. For example, at the *New York Times*, nine trackers are active. Some of these are innocuous, and add useful features. For example, the website uses Typekit, a service that gives websites more interesting fonts. Others are what you would expect: the *Times* has its own tracker. Others are entirely undesirable: they are trackers from ad networks like DoubleClick and Facebook that spy on your consumer activities. (Facebook, by the way, is not a website to communicate with friends, but a website to record your consumer preferences.) Ghostery lets you toggle a switch to turn off each of the trackers.

If you choose to install Ghostery, I suggest that you configure it to block all trackers. Then you can enable the desirable trackers as you see fit. For example, the *Chronicle*'s website uses a third-party commenting service called Disqus; you would have to re-enable Disqus to leave a comment on this post. You can enable those settings by going to the Ghostery options page. I would check the box to "Enable tracker library auto-updating," so that the list of trackers is kept up-to-date. Then I would "select all" of the trackers, which will disable all of them. And on the options page I would check "Block new elements by default," which will make any new tracker disabled by default. If you notice that Ghostery breaks something on a web page that you find desirable (perhaps the comments on this page) then it's easy to re-enable them on a case by case basis. At least you'll be making an informed decision.

Ironically, Ghostery makes money by selling information about web trackers to businesses. You can opt-in to sharing the data that the plugin gathers about trackers on websites. Ghostery promises that this information is just about trackers, not about you. (You can decide whether you believe them.)

You need not be [RMS](http://stallman.org/) to want a more control over your privacy when browsing on the web. Ghostery is the best tool I've found to give you fine grained control over which websites track you and how.\*

-   *N.B.* Ghostery does not protect you from NSA tracking.
