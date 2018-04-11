---
categories:
- digital history
date: '2014-05-22T00:13:00+00:00'
slug: 'map-projections-and-the-historical-imagination'
title: Map projections and the historical imagination
---

At the *New Republic*, [Susan Schulten](http://www.mappingthenation.com/) has a [fascinating article](http://www.newrepublic.com/article/117835/richard-edes-harrison-reinvented-mapmaking-world-war-2-americans) about maps made by Richard Edes Harrison during World War II. Schulten writes that Harrison, an artist and not a cartographer, changed the American public's perception of the war and world by refusing to use the Mercator projection. Instead, he drew maps from various perspectives above the earth, evoking "the perspective of a pilot, but one placed at an infinite distance."
<figure id="attachment_571" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-571" src="https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?resize=739%2C556" alt="" width="739" height="556" srcset="https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?resize=768%2C578 768w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?resize=300%2C226 300w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?resize=1024%2C771 1024w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?w=1536 1536w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/harrison-ussr.jpg?w=1478 1478w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Richard Edes Harrison, 'Europe as Viewed from the U.S.S.R.' (1944). Courtesy of [David Rumsey Map Collection](http://www.davidrumsey.com/luna/servlet/s/qpbh66).
</figcaption>
</figure>
Schulten [writes](http://www.newrepublic.com/article/117835/richard-edes-harrison-reinvented-mapmaking-world-war-2-americans):

> Most professional cartographers celebrated his provocative style for its ability to foster a more dynamic understanding of geographical relationships ... . \[H\]is goal was to wrench Americans out of a two-dimensional sense of geography, and embrace an understanding of perspective and direction.

Much of maps that historians make on the web is also limited by the Mercator projection. Google Maps, Stamen, and Map Box all use a "[Web Mercator](http://spatialreference.org/ref/sr-org/epsg3857-wgs84-web-mercator-auxiliary-sphere/)" projection for the tiles they provide. If you've used a web map that pans and zooms, it was almost certainly created in the Mercator projection. Even if there are [good technical reasons](http://www.mapthematics.com/forums/viewtopic.php?f=8&t=251) for Google to use that projection, historians have their own purposes that could be better served by alternative perspectives on the earth.
<figure id="attachment_570" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-570" src="https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/google.png?resize=739%2C358" alt="" width="739" height="358" srcset="https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/google.png?resize=768%2C372 768w, https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/google.png?resize=300%2C145 300w, https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/google.png?resize=1024%2C495 1024w, https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/google.png?w=1048 1048w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
The world seen from Google.
</figcaption>
</figure>
According to Schulten, Harrison never thought of himself as a cartographer, and "cartographers were quick to point out that no such perspective existed in nature." But thanks to the work of [Mike Bostock](http://bost.ocks.org/mike/) and [Jason Davies](https://www.jasondavies.com/), the kinds of perspectives that Harrison used are available to web mapmakers.[<sup>1</sup>](#fn1){\#fnref1.footnoteRef} As part of the [geographic projections](https://github.com/d3/d3-geo-projection) available in [D3](http://d3js.org/), Bostock has created a [satellite projection](http://bl.ocks.org/mbostock/3790444), with which the mapmaker can control the view of the earth as from a camera mounted on satellite. Such a map could be combined with any of the visualization techniques---choropleth shading, bubble maps, place names---for which [D3](http://d3js.org/) is known.
<figure id="attachment_567" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-567" src="https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/bostock-satellite.png?resize=739%2C539" alt="" width="739" height="539" srcset="https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/bostock-satellite.png?resize=768%2C560 768w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/bostock-satellite.png?resize=300%2C219 300w, https://i0.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/bostock-satellite.png?w=995 995w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
An example of Bostock's [satellite projection](http://bl.ocks.org/mbostock/3790444).
</figcaption>
</figure>
Jason Davies has gone a step further and figured out [how to re-project Mercator map tiles](https://www.jasondavies.com/maps/raster/) into any map projection.[<sup>2</sup>](#fn2){\#fnref2.footnoteRef}
<figure id="attachment_568" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-568" src="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/davies-satellite.png?resize=739%2C456" alt="" width="739" height="456" srcset="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/davies-satellite.png?resize=768%2C474 768w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/davies-satellite.png?resize=300%2C185 300w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/davies-satellite.png?resize=1024%2C631 1024w, https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/davies-satellite.png?w=1106 1106w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Davies's reprojected map tiles for a [satellite projection](https://www.jasondavies.com/maps/raster/satellite/).
</figcaption>
</figure>
This means that historians have the tools they need to re-imagine their historical maps outside of the constraints of any particular map projection.[<sup>3</sup>](#fn3){\#fnref3.footnoteRef} Just as Harrison changed the public view of the geopolitical situation with his perspective maps, so historians would have a powerful imaginative and rhetorical tool at their disposal if they could choose the perspective of their maps. One can imagine, for example, a map that redraws the Atlantic slave trade from the perspective of West Africa, or a map that literally [faces east from Indian country](http://www.amazon.com/gp/product/0674011171/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=0674011171&linkCode=as2&tag=thebacgla-20&linkId=PLQFRYLBZSK2JH5F). As is often the case when I read Schulten's work, the historical maps she studies provide a useful suggestion for how historians could make their own maps to imagine the past.
<figure id="attachment_569" style="max-width: 739px" class="wp-caption alignnone">
<img class="size-medium_large wp-image-569" src="https://i2.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/facing-east-768x359.png?resize=739%2C345" alt="" width="739" height="345" srcset="https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/facing-east.png?resize=768%2C359 768w, https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/facing-east.png?resize=300%2C140 300w, https://i1.wp.com/lincolnmullen.com/blog/wp-content/uploads/2017/01/facing-east.png?w=837 837w" sizes="(max-width: 739px) 100vw, 739px" data-recalc-dims="1" />
<figcaption class="wp-caption-text">
Looking east over the Chesapeake.
</figcaption>
</figure>
<section class="footnotes">

------------------------------------------------------------------------

<li id="fn1">
There are many other artistic elements to Harrison's work, such as the exaggerated topography, that it would be difficult to reproduce in a programmatic map.<a href="#fnref1">↩</a>
</li>
<li id="fn2">
An example showing <a href="http://www.jasondavies.com/maps/tile/">how Davies's method for re-projecting tiles works</a>.<a href="#fnref2">↩</a>
</li>
<li id="fn3">
Such perspectives have been available in Google Earth for many years, but asking viewers to drop into a proprietary application prevents the integration of web maps with prose narrative and argument.<a href="#fnref3">↩</a>
</li>
</section>
