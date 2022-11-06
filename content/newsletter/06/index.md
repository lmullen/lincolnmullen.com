---
title: "Honesty in visualizations"
date: "2020-08-16"
slug: 6
draft: false
tags:
- newsletter
- visualization
---

Greetings. And welcome back to "Working on It."

-------------------

Everyone who writes is familiar with the process of revision. Less familiar is the process of revising a visualization. The goal is to make the visualization display as much meaning as the data will support, and no more. One also has a responsibility to do the best one can to avoid having the visualization misinterpreted. There is a way of talking about visualizations---a mistake, in my view---that focuses primarily on their distortions and deceptions, as if people never lie with words. Nevertheless, most people are less sophisticated at reading visualizations than prose, and so authors bear more of a burden to do right by the reader. 

As an example of trying to make an honest visualization, here is a series of screenshots showing revisions I made to *America's Public Bible*. These visualizations are part of the expanded version of the site and aren't available yet; only the [prototype](https://americaspublicbible.org) is up. The goal of this particular visualization is to let the reader pick a Bible verse, and then show the trend in the rate of quotations over time in American newspaper.

Making a visualization like this entails innumerable small choices along the way: here are a few of the salient ones. Let's start with what the dataset looks like. 

{{< image src="predictions.png" alt="Table of predictions" >}}

The prediction model I trained starts with newspaper pages. After identifying potential quotations, it makes a prediction: is this a quotation from the Bible or not? The result is table like the one above. Each row indicates that a particular verse appears on a particular newspaper page, with the probability that it is actually a quotation and not just a false positive. That kind of machine learning, called supervised classification, has some well-established procedures to ensure that the predictions are honest and not the result of self-deception. Important as they are, those rules need not detain us now.

This dataset has a few specific considerations for analyzing it honestly. One consideration is what the correct threshold is for determining whether something is a genuine quotation or not. Obviously it has to be more than 50%, but how much more? In the prototype I set that threshold at 90% because, frankly, I didn't want to be embarrassed by people noticing entries that weren't quotations. But that's not quite honest either: it leaves out many, many quotations which are genuine but just have a lower probability. (The tradeoff here is similar to that of a medical diagnostic test, albeit with rather different stakes.) The revised version of the site will include quotations above 58%. Another consideration is the fact that the OCR for *Chronicling America* can be quite dreadful in places; obviously it is not possible to find a quotation in a bunch of gibberish. In calculating the rate of quotations, the number of quotations is the numerator but I have also had to figure out how to exclude junk OCR in order to get the denominator right.

Once we've got a dataset of quotations, it is straightforward to aggregate them into the rate of quotations per word per year. Visualizing the rate, instead of the number, of quotations is a key decision. In this case it is an obvious choice. The number of quotations almost always goes up over time, simply because there were a lot more newspaper from the end of the nineteenth century that were digitized than from the beginning. The trend is what we really care about.

{{< image src="trend-lines.png" alt="Comparison of trend lines" >}}

Drawing a trend line is not as simple as it might seem. The contrived screenshot above shows two different options for how to visualize the trend. The gray line is the simplest option, but not very good; the red line is a better option. The gray line is very spiky because it shows each individual year, and because each pair of points is joined by a straight line segment. Year to year the data is noisy and we can only expect to find meaning in the long-term trends. I have seen people try to divine meaning from every little spike, when they are just noise. So the red line is a better, if not entirely satisfactory approach. For that line, each data point is shown as a five-year rolling average (e.g., the data point for 1860 is the average of the years 1858–1862). And the data points are joined with a smooth curve. The red trend line is the best I can come up with to represent the nature of the actual trend: in most instances quite gradual, with occasional genuine spikes.

The interface around the visualization is unfinished, but there are a few other decisions that had to be made. One is which verses to include in the list the user can select. This is somewhat subjective, but in many instances there are verses which have a very high rate of false positives. For example, the phrase "went into the city" frequently appears in both the Bible and (who would have guessed?) in newspapers. Then too, the Bible frequently repeats itself: the synoptic Gospels and quotations of the Old Testament in the New Testament are just the most obvious examples. In this case, the verse "Suffer little children to come unto me" was one of the most popular verses in the nineteenth century, but there is no reliable way to computationally distinguish between a quotation to Matthew 19:14, Mark 10:14, and Luke 18:16. In such instances, I collapse all the quotations into a single reference which stands for all three.

{{< image src="cutoff-date.png" alt="Cutoff date" >}}

The next consideration is which cutoff dates to pick for the visualization. The visualization above shows the trend to the maximum chronological extent of the *Chronicling America* corpus in the 1960s, whereas the visualizations above extend only to 1922. There appears to be a huge increase in the rate of quotations in the mid twentieth century, but displaying that trend would be misleading. There are next to no newspapers in *Chronicling America* after 1922. (Unreasonably long copyright terms rear their ugly heads again.) Thus we face what we might call the batting average problem: if you only have a few plate appearances and get lucky, you could end up with an abnormally good batting average. So, it's more honest to represent the trend line only when there is a substantial underlying corpus. For *Chronicling America*, the correct minimum and maximum years are easily determined.

{{< image src="complete.png" alt="Completed visualization" >}}

This last screenshot is what the mostly complete visualization looks like. This version adds another feature to help readers accurately interpret the trends. In addition to showing the trend line for *Chronicling America*, it also shows the trend line for a separate corpus, Gale's *Nineteenth-century Newspapers*. Since the corpora don't completely overlap, one would not expect the trend lines to be identical. But one would also expect them to be close, as they are here.

Many difficult decisions go into making a visualization as honest as possible, and then there is more work to be done interpreting it honestly. More on the problem of interpretation some other time, but I will leave you with my [parlor trick](https://lincolnmullen.com/blog/isnt-it-obvious/) for giving talks about this project. 

----------------------------------

### Brief book note

Two well-written, recently published works take up the themes of masculinity and femininity within American evangelicalism or megachurch Christianity.

Kristen Kobes Du Mez's [*Jesus and John Wayne: How White Evangelicals Corrupted a Faith and Fractured a Nation*](https://kristindumez.com/books/jesus-and-john-wayne/) chronicles the stream of masculinity that has pervaded American evangelicalism. No mere monograph, this history is quite sweeping in its interpretation  of the movement in terms of its advocacy for (or perhaps, obsession with?) a particular configuration of family and gender roles. The next time that I teach twentieth-century evangelicalism, I will likely start with Du Mez's book to provide the framework. (Don't miss this tongue-in-cheek [review](https://www.patheos.com/blogs/anxiousbench/2020/08/jesus-and-john-wayne-a-guide-for-biblical-manhood/).)

Kate Bowler's [*The Preacher's Wife: The Precarious Power of Evangelical Women Celebrities*](https://press.princeton.edu/books/hardcover/9780691179612/the-preachers-wife) is a fascinating study of celebrity Christian women, often the wives of megachurch pastors. The main dynamic she explores is that evangelical women operate under much stricter theological constraints governing gender roles than their liberal or mainline counterparts, but evangelical women have much great access to the marketplace by which celebrity power is created. It's a fitting sequel to Bowler's *Blessed*, on the history of the prosperity gospel.

------------------------

### Random screen shot

Too many screenshots above, so here is a photo of my tomato plants. No fig tree, but I'm thankful that I can sit in peace and safety under my own vines.

{{< image src="tomatoes.png" alt="Tomato plants" >}}

-------------------------

### Updates

Reading: Already finished it, but next time I'll write briefly about Tara Isabella Burton's *Strange Rites*. Currently reading Diana Butler Bass's *Standing Against the Whirlwind: Evangelical Episcopalians in Nineteenth-Century America*.

Working: Getting reading to teach a minor field readings for PhD students in American religion, the DH Practicum for incoming students at RRCHNM, as well as a course on "Capitalism and American Religion." Here's the [syllabus](https://lincolnmullen.com/courses/capitalism.2020/) for that last one.

Watching: [*The Expanse*](https://www.imdb.com/title/tt3230854/). I promised my colleagues at RRCHNM that I would finish it before we reopen, but I'm not going to meet my goal.

Playing: Some PGA Tour game on the Xbox, which is infinitely more stressful than any first-person shooter.
