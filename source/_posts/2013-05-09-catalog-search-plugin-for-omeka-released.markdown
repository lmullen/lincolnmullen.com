---
layout: post
title: "Catalog Search Plugin for Omeka Released"
date: 2013-05-09 08:31
comments: true
categories: 
- American Converts Database
- digital humanities
tags:
- American Converts Database
- digital humanities
- code
- coding
- PHP
- Omeka
---

For my work on the [American Converts Database][] I've created an
[Omeka][] plugin that connects Omeka items to records in other catalogs.
My database, and probably other databases as well, have people or items
that are identified by Library of Congress subject headings. One could
manually search other catalogs by the subject heading, but its easy
enough to generate links that do the searching automatically. The plugin
adds links to Archive Grid, the DPLA, Google Books, Google Scholar, the
Hathi Trust, JSTOR, the Library of Congress, and WorldCat. The links
work best if you use Omeka's very good [Library of Congress Suggest
plugin][] to add your subject headings. So, for example, from this item
for [Warder Cresson][] in the Converts Database, you can click links to
the [Library of Congress][] to find books that Cresson wrote, to
[JSTOR][] to find articles about him, and to [Google Books][] to search
inside books for mentions of him.

I had written an earlier version of the plugin that basically hard wired
the links so that a user had to modify the plugin itself to change which
connections were made. This first public release of the plugin lets
users decide which catalog links to display, and also to add their own
custom links.

You can [download the plugin][] from the Omeka site, [see the code][] at
GitHub, or let me know [about any bugs][].

  [American Converts Database]: http://americanconverts.org
  [Omeka]: http://omeka.org
  [Library of Congress Suggest plugin]: http://omeka.org/add-ons/plugins/library-of-congress-suggest/
  [Warder Cresson]: http://americanconverts.org/items/show/7
  [Library of Congress]: http://catalog2.loc.gov/vwebv/search?searchArg=Cresson%2C+Warder%2C+1798-1860&searchCode=GKEY%5E*&searchType=0
  [JSTOR]: http://www.jstor.org/action/doBasicSearch?Query=Cresson+Warder
  [Google Books]: https://www.google.com/search?btnG=Search+Books&tbm=bks&tbo=1&q=Cresson+Warder
  [download the plugin]: http://omeka.org/add-ons/plugins/catalog-search/
  [see the code]: https://github.com/lmullen/plugin-CatalogSearch
  [about any bugs]: https://github.com/lmullen/plugin-CatalogSearch/issues
