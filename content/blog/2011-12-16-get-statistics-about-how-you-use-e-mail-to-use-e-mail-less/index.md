---
title: "Get statistics about how you use e-mail to use e-mail less"
date: "2011-12-16"
slug: "get-statistics-about-how-you-use-e-mail-to-use-e-mail-less"
description: "Profile, then optimize: two ways to gather data about how you use Gmail so you can find where to spend less time on e-mail."
draft: false
crosspost:
  url: "http://chronicle.com/blogs/profhacker/get-statistics-about-how-you-use-e-mail-to-use-e-mail-less/37753"
  source: "ProfHacker"
tags:
- tools
---

{{< image src="2-300x177.png" alt="Chart of e-mail statistics" >}}

There is a principle in writing software, which also applies to hacking your life: “profile, then optimize.” The idea is that you should gather data first, which will then let you know where the biggest gains are to be found by making optimizations. That’s why the data junkies at ProfHacker are big on gathering information about yourself with tools like [Ask My Every](http://chronicle.com/blogs/profhacker/learn-about-yourself-with-askmeevery-com/36382) and [750 Words](http://chronicle.com/blogs/profhacker/writers-boot-camp-using-750wordscom/23383).

<!--more-->

One area where it’s fairly easy to gather data is your e-mail inbox. Here are two ways to gather statistics about how you use e-mail. (Sorry, but these two only work for users of Gmail or Google Apps e-mail. With a little looking, you can probably find other ways of doing the same thing with different e-mail programs.)

{{< image src="email-report.png" alt="Sample e-mail statistics" >}}

One option is to use a [script for Gmail and Google Docs](http://code.google.com/googleapps/appsscript/articles/gmail-stats.html) that gives you a monthly report on what happened in your inbox. The Gmail Meter script is very easy to set up by [following the instructions](http://code.google.com/googleapps/appsscript/articles/gmail-stats.html) that come with it. You simply create a new Google spreadsheet (with the same account that is linked to your Gmail), add the script through the Google Docs toolbar, and grant it permission to trawl through your e-mail archive. A couple hours later you’ll receive an e-mail with some key statistics, attractively charted. The e-mail will also contain a link to a table of raw data letting you know how many e-mails you have sent and received from particular people. You can see a sample report in the image to the right ([click to enlarge](http://chronicle.com/blogs/profhacker/files/2011/12/email-report.png)). At the start of every month, you should get an updated report for the past month.

If you want a report for your entire e-mail archive, not just for the past month, you can use [mail-trends](http://code.google.com/p/mail-trends/), a command line Python script. If you have a basic familiarity with the [command line](http://chronicle.com/blogs/profhacker/tag/the-profhacker-guide-to-the-command-line), you can probably get along by following the instructions. (I had to patch a file, but the wiki explained what the problem was and how to fix it.) The program scans the headers of your Gmail archive, which took about 10 minutes for me, then outputs a helpful set of statistics. You can [see a sample report here](http://persistent.info/mail-trends/enron/). Particularly interesting are the stacked line charts that show how your volume of traffic to particular people ebbs and flows over time.

What can you do with this data? I was surprised to find out who sent me the most e-mail. Armed with that information, I can unsubscribe from some lists that aren’t worth the effort (probably true of most e-mail lists) and filter some other senders. I also learned that I send a lot of e-mail on Saturdays and Sundays, which I’m probably better off leaving till Monday. Some of the threads that I participate in are very long: usually about finding a time for a meeting. This is more incentive to use a service like [Google appointment slots](http://chronicle.com/blogs/profhacker/scheduling-101-appointment-slots-in-google-calendar/34575), [Tungle.me](http://chronicle.com/blogs/profhacker/scheduling-101-using-tungleme-for-committee-meetingsstudent-appointments/22944), or [Doodle](http://chronicle.com/blogs/profhacker/scheduling-101-using-doodle-for-student-appointments/22783). And I learned that while I’m better than the people who send mail to me at keeping messages short, 25 percent of the messages I send are still more than 200 words. Maybe next month I can cut that percentage back.

The point is that you don’t know what you need to optimize until you gather the data about yourself. The end of the semester and of the year is a particularly good time for data gathering to support the New Year’s efforts at optimization.

**Have you gathered data about your e-mail habits? What did you learn?**
