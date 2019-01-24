---
title: 'Chronicling America OCR debatcher'
date: 2019-01-24T18:01:12-05:00
slug: 'chronam-debatcher'
series:
subjects:
- go
- software
draft: false
---

This probably useful only for me, but I've made a small [utility](https://github.com/lmullen/chronam-ocr-debatcher) to help get the
Chronicling America OCR files. The batch files from the Chronicling America
[bulk data downloads](https://chroniclingamerica.loc.gov/about/api/#bulk-data)
are `.tar.bz2` files with both plain text and XML versions of the OCR text of the
newspaper pages. The files are slow to unzip and dump tens of thousands of
files, at least half of which you don't need, onto your disk. So the
utility process the batches without unzipping them and creates a CSV file with
the text and the IDs used elsewhere in Chronicling America. You can [get the
utility](https://github.com/lmullen/chronam-ocr-debatcher) at GitHub.
