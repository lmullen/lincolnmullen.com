---
title: 'The Quantified Dissertation: Using Bash and cron to Track Word Counts'
date: "2012-08-17"
slug: 'the-quantified-dissertation-using-bash-and-cron-to-track-word-counts'
description: "A Bash script and cron job that logs hourly word counts for every dissertation chapter to a CSV file, so you can chart your writing progress over time."
tags:
- technology
draft: false
---

Here is a hack for those of you writing a dissertation (or book, or other long form project) in plain text. Since I am writing my dissertation in [Markdown](http://chronicle.com/blogs/profhacker/markdown-the-syntax-you-probably-already-know/35295) and [Pandoc](http://chronicle.com/blogs/profhacker/pandoc-converts-all-your-text-documents/38700), it's easy to find out the word count for a chapter from the command line:

```bash
$ wc -w ch4.dissertation.Mullen.md
```

I wanted to track my progress by regularly getting a word count for every chapter. To do that, I wrote a Bash script which runs `wc` and stores the data to a CSV file. Here is the script:

```bash
#!/bin/bash
source /Users/lmullen/.bash_profile

EPOCH=$(date +"%s")
DATE=$(date)
WCCH1=$(wc -w /Users/lmullen/acad/dissertation/ch1.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch1.dissertation.Mullen.md//' | sed 's/ //g')
WCCH2=$(wc -w /Users/lmullen/acad/dissertation/ch2.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch2.dissertation.Mullen.md//' | sed 's/ //g')
WCCH3=$(wc -w /Users/lmullen/acad/dissertation/ch3.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch3.dissertation.Mullen.md//' | sed 's/ //g')
WCCH4=$(wc -w /Users/lmullen/acad/dissertation/ch4.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch4.dissertation.Mullen.md//' | sed 's/ //g')
WCCH5=$(wc -w /Users/lmullen/acad/dissertation/ch5.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch5.dissertation.Mullen.md//' | sed 's/ //g')
WCCH6=$(wc -w /Users/lmullen/acad/dissertation/ch6.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch6.dissertation.Mullen.md//' | sed 's/ //g')
WCCH7=$(wc -w /Users/lmullen/acad/dissertation/ch7.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch7.dissertation.Mullen.md//' | sed 's/ //g')
WCCH8=$(wc -w /Users/lmullen/acad/dissertation/ch8.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/ch8.dissertation.Mullen.md//' | sed 's/ //g')
WCPRE=$(wc -w /Users/lmullen/acad/dissertation/preface.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/preface.dissertation.Mullen.md//' | sed 's/ //g')
WCINTRO=$(wc -w /Users/lmullen/acad/dissertation/introduction.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/introduction.dissertation.Mullen.md//' | sed 's/ //g')
WCCONCL=$(wc -w /Users/lmullen/acad/dissertation/conclusion.dissertation.Mullen.md | sed 's/\/Users\/lmullen\/acad\/dissertation\/conclusion.dissertation.Mullen.md//' | sed 's/ //g')
echo $EPOCH,$DATE,$WCCH1,$WCCH2,$WCCH3,$WCCH4,$WCCH5,$WCCH6,$WCCH7,$WCCH8,$WCPRE,$WCINTRO,$WCCONCL >> /Users/lmullen/acad/dissertation/wordcount-data.csv
```

First, the script finds out the date in two formats (lines 4--5), the Unix "epoch" time (or number of seconds since midnight on January 1, 1970) and a human readable time. I thought I should store both since I didn't know which would be most useful later on. Second, the script counts the number of words in each chapter and in the front matter and back matter. It also strips away all of the output from that command except the word count itself (lines 6--16). Someone with better scripting skills than me could probably compress those eleven lines into one or two. Finally, the script appends a line with the data to a CSV file (line 17).

When the script is executed, it returns data like this (from the very first line in my data file):

```
EPOCH,DATETIME,CH1WC,CH2WC,CH3WC,CH4WC,CH5WC,CH6WC,CH7WC,CH8WC,PREWC,INTROWC,CONCLWC
1332736387,Mon Mar 26 00:33:07 EDT 2012,0,0,0,12412,0,0,0,0,0,0,0
```

As you can see, by March 26 I had written 12,412 words for chapter 4, and not much else. Apparently I also wrote this script after midnight.

I run the script automatically each hour by adding this line to my crontab:

```bash
0   *   *   *   *   /Users/lmullen/acad/dissertation/wordcount.sh
```

I haven't done much with the data yet, but the point is to collect it while I am writing. It's easy enough to open the CSV file in Excel and generate graphs of my progress. Eventually I may also analyze the data in [R](http://www.r-project.org/).

Gathering the data is one problem. Getting the word counts to increase at an acceptable pace is another problem altogether.
