---
date: '2013-05-09T12:31:00+00:00'
slug: 'catalog-search-plugin-for-omeka-released'
title: Catalog Search Plugin for Omeka Released
---

For my work on the [American Converts Database](http://americanconverts.org) I've created an [Omeka](http://omeka.org) plugin that connects Omeka items to records in other catalogs. My database, and probably other databases as well, have people or items that are identified by Library of Congress subject headings. One could manually search other catalogs by the subject heading, but its easy enough to generate links that do the searching automatically. The plugin adds links to Archive Grid, the DPLA, Google Books, Google Scholar, the Hathi Trust, JSTOR, the Library of Congress, and WorldCat. The links work best if you use Omeka's very good [Library of Congress Suggest plugin](http://omeka.org/add-ons/plugins/library-of-congress-suggest/) to add your subject headings. So, for example, from this item for [Warder Cresson](http://americanconverts.org/items/show/7) in the Converts Database, you can click links to the [Library of Congress](http://catalog2.loc.gov/vwebv/search?searchArg=Cresson%2C+Warder%2C+1798-1860&searchCode=GKEY%5E*&searchType=0) to find books that Cresson wrote, to [JSTOR](http://www.jstor.org/action/doBasicSearch?Query=Cresson+Warder) to find articles about him, and to [Google Books](https://www.google.com/search?btnG=Search+Books&tbm=bks&tbo=1&q=Cresson+Warder) to search inside books for mentions of him.

I had written an earlier version of the plugin that basically hard wired the links so that a user had to modify the plugin itself to change which connections were made. This first public release of the plugin lets users decide which catalog links to display, and also to add their own custom links.

You can [download the plugin](http://omeka.org/add-ons/plugins/catalog-search/) from the Omeka site, [see the code](https://github.com/lmullen/plugin-CatalogSearch) at GitHub, or let me know [about any bugs](https://github.com/lmullen/plugin-CatalogSearch/issues).
