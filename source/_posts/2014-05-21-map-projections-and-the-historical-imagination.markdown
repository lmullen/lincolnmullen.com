---
categories:
- mapping
- d3
date: '2014-05-21 20:13'
layout: post
title: Map Projections and the Historical Imagination
...

At the *New Republic*, [Susan Schulten][] has a [fascinating article][]
about maps made by Richard Edes Harrison during World War II. Schulten
writes that Harrison, an artist and not a cartographer, changed the
American public's perception of the war and world by refusing to use the
Mercator projection. Instead, he drew maps from various perspectives
above the earth, evoking "the perspective of a pilot, but one placed at
an infinite distance."

{% figure harrison/harrison-ussr jpg "Richard Edes Harrison, 'Europe as Viewed from the U.S.S.R.' (1944). Courtesy of [David Rumsey][]." %}

Schulten [writes][fascinating article]:

> Most professional cartographers celebrated his provocative style for
> its ability to foster a more dynamic understanding of geographical
> relationships ... . [H]is goal was to wrench Americans out of a
> two-dimensional sense of geography, and embrace an understanding of
> perspective and direction.

Much of maps that historians make on the web is also limited by the
Mercator projection. Google Maps, Stamen, and Map Box all use a "[Web
Mercator][]" projection for the tiles they provide. If you've used a web 
map that pans and zooms, it was almost certainly created in the Mercator 
projection. Even if there are [good technical reasons][] for Google to 
use that projection, historians have their own purposes that could be 
better served by alternative perspectives on the earth.

{% figure harrison/google png "The world seen from Google." %} 

According to Schulten, Harrison never thought of himself as a
cartographer, and "cartographers were quick to point out that no such
perspective existed in nature." But thanks to the work of [Mike
Bostock][] and [Jason Davies][], the kinds of perspectives that Harrison
used are available to web mapmakers.[^1] As part of the [geographic
projections][] available in [D3][], Bostock has created a [satellite
projection][], with which the mapmaker can control the view of the earth
as from a camera mounted on satellite. Such a map could be combined with 
any of the visualization techniques---choropleth shading, bubble maps, 
place names---for which [D3][] is known.

{% figure harrison/bostock-satellite png "An example of Bostock's [satellite projection][]." %}

Jason Davies has gone a step further and figured out [how to reproject
Mercator map tiles][] into any map projection.[^2]

{% figure harrison/davies-satellite png "Davies's reprojected map tiles for a [satellite projection][1]." %}

This means that historians have the tools they need to re-imagine their
historical maps outside of the constraints of any particular map
projection.[^3] Just as Harrison changed the public view of the
geopolitical situation with his perspective maps, so historians would
have a powerful imaginative and rhetorical tool at their disposal if
they could choose the perspective of their maps. One can imagine, for
example, a map that redraws the Atlantic slave trade from the 
perspective of West Africa, or a map that literally [faces east from 
Indian country][]. As is often the case when I read Schulten's work, 
the historical maps she studies provide a useful suggestion for 
how historians could make their own maps to imagine the past.

{% figure harrison/facing-east png "Looking east over the Chesapeake." %}

[^1]: There are many other artistic elements to Harrison's work, such as
    the exaggerated topography, that it would be difficult to reproduce
    in a programmatic map.

[^2]: An example showing [how Davies's method for re-projecting tiles
    works][].

[^3]: Such perspectives have been available in Google Earth for many
    years, but asking viewers to drop into a proprietary application
    prevents the integration of web maps with prose narrative and
    argument.

  [Susan Schulten]: http://www.mappingthenation.com/
  [fascinating article]: http://www.newrepublic.com/article/117835/richard-edes-harrison-reinvented-mapmaking-world-war-2-americans
  [David Rumsey]: http://www.davidrumsey.com/luna/servlet/s/qpbh66
  [Web Mercator]: http://spatialreference.org/ref/sr-org/epsg3857-wgs84-web-mercator-auxiliary-sphere/
  [good technical reasons]: http://www.mapthematics.com/forums/viewtopic.php?f=8&t=251
  [Mike Bostock]: http://bost.ocks.org/mike/
  [Jason Davies]: https://www.jasondavies.com/
  [geographic projections]: https://github.com/d3/d3-geo-projection
  [D3]: http://d3js.org/
  [satellite projection]: http://bl.ocks.org/mbostock/3790444
  [how to reproject Mercator map tiles]: https://www.jasondavies.com/maps/raster/
  [1]: https://www.jasondavies.com/maps/raster/satellite/
  [faces east from Indian country]: http://www.amazon.com/gp/product/0674011171/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=0674011171&linkCode=as2&tag=thebacgla-20&linkId=PLQFRYLBZSK2JH5F
  [how Davies's method for re-projecting tiles works]: http://www.jasondavies.com/maps/tile/
