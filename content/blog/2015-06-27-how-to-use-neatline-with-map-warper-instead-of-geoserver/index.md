---
title: How to use Neatline with Map Warper instead of GeoServer
date: "2015-06-27"
slug: 'how-to-use-neatline-with-map-warper-instead-of-geoserver'
description: "Map Warper's WMS service can substitute for GeoServer in Neatline, making georectified historical maps available without running your own tile server."
aliases:
- /blog/how-to-use-neatline-without-geoserver/
tags:
- digital history
- mapping
draft: false
---

> **TL;DR.** You don't need GeoServer to use historic map layers in Neatline: you just need any [WMS server](https://en.wikipedia.org/wiki/Web_Map_Service), such as [Map Warper](http://mapwarper.net/). Map Warper's base WMS URL is `http://mapwarper.net/maps/wms/{MAPID}`; the layer ID is whatever title you gave the map in Map Warper.

For any map that needs to connect its elements to full metadata and/or narratives, [Omeka](http://omeka.org/) and the [Neatine](http://neatline.org/) plugin are a useful combination. I used Omeka and Neatline for a class on [mapping Boston's religions](http://omeka.lts.brandeis.edu/neatline/fullscreen/mapping-bostons-religions)---the kind of class I hope to teach again---and most recently in a [workshop session on deep mapping](https://projects.lincolnmullen.com/spatial-workshop/deep-maps.html).

The trouble with Neatline is its recommended dependency, [GeoServer](http://geoserver.org/). Omeka and Neatline require only a [LAMP stack](https://en.wikipedia.org/wiki/LAMP_(software_bundle)), which you can to find on any shared web host. (Shout out to [Reclaim Hosting](https://portal.reclaimhosting.com/aff.php?aff=004), which is far and away the best host for academics and students.) Installing Omeka and the [main Neatline plugin](http://omeka.org/add-ons/plugins/neatline/) is very easy. GeoServer, however, is a Java application and is a pain to install. At the very least you'll need a virtual private server and some webmaster chops to get it up and running.

I've figured out how to use all of Neatline's features without installing GeoServer by using a different WMS server instead. The people at [THATCamp Prime](http://prime2015.thatcamp.org) this weekend indicated that an explanation would be helpful, so I'm writing up how to do that.

<!--more-->

Neatline has many parts, and you have to understand what functionality each of them provides. The [Neatline documentation](http://docs.neatline.org/installing-neatline.html) is clear that with just the main Neatline plugin you can add points or lines or polygons associated with Omeka items on top of common map layers such as Google Maps or Open Street Maps.[^1] To take a trivial example, here is a single point in Neatline with an Open Street Maps base layer. If this is all you want to do, then you don't need GeoServer or another WMS server at all.

{{< figure src="point.png" caption="A single point in Neatline." >}}

But one of the compelling features of Neatline is the ability to add georectified maps. These maps are not just additional map layers, but are also items which can be associated with metadata. As an example, in the screenshot below I have included a map titled "[Historic Canterbury](http://www.davidrumsey.com/luna/servlet/s/sgm48r)." The [Neatline documentation](http://docs.neatline.org/installing-neatline.html) says that to include such maps you need GeoServer: "For projects that need to incorporate custom imagery, though---for example, historical projects that want to overlay georeferenced maps on top of the default layers, or art history projects that need to use scans of paintings as base layers in exhibits---a third piece of software called GeoServer is necessary."

{{< figure src="overlay.png" caption="A historical map layer in Neatline." >}}

What the Neatline documentation does not clearly explain is why GeoServer is necessary. GeoServer provides you with the ability to serve georectified maps via the [Web Map Service](https://en.wikipedia.org/wiki/Web_Map_Service) (WMS) protocol. In the [most common use case](http://docs.neatline.org/style-tab-imagery.html) for Neatline, GeoServer takes a georectified TIFF of a map located on your server and serves up the map for your users' browsers.

But [many services](http://trac.osgeo.org/openlayers/wiki/AvailableWMSServices) on the web provide geospatial data via WMS. We could, if we took time to figure it out, overlay NASA temperature data in Neatline, or topographical maps. All we need is the base URL of a WMS service and the layer IDs of the maps that are being served.

[Map Warper](http://mapwarper.net/) is a service that provides a web interface for georectifying maps.[^2] For instance, [here is the map of Canterbury](http://mapwarper.net/maps/9592#Preview_Map_tab) that I displayed above. Map Warper provides several export options, including WMS. You can dig through the XML returned by [its WMS capabilities query](http://mapwarper.net/maps/wms/9592?request=GetCapabilities&service=WMS&version=1.1.1) in order to figure out the correct settings for Neatline, or you can follow these instructions. The base URL will take this form: `http://mapwarper.net/maps/wms/{MAPID}`. The map ID is the numeric value assigned to your map at Map Warper, which you can figure out from the Map Warper URL or metadata. For the Canterbury map, the ID is `9592`, so the full WMS base URL is `http://mapwarper.net/maps/wms/9592`. Then you need the layer ID. The layer ID is whatever title you assigned to your map at Map Warper, in this case, `Historic Canterbury`. Once you've gathered that information, you can create a new item in Neatline, and enter that into the "Style" tab on the item interface.

{{< figure src="form.png" caption="Adding WMS settings to a Neatline item." >}}

Omeka and Neatline together are very useful for certain kinds of maps, so I hope this brief explanation helps people who might otherwise be deterred by the GeoServer dependency.

#### Bonus: Installing GeoServer via Docker

Once you grok [Docker](https://www.docker.com/), GeoServer is not that hard to install, though it is still work to maintain. This [Docker image for GeoServer](https://github.com/kartoza/docker-geoserver) is well documented. I was able to get it working on a low-powered [Digital Ocean instance](https://www.digitalocean.com/?refcode=ae1e7cf83916) without any fuss.

[^1]: I can't find this in the Neatline documentation itself, but [according to the Neatline forums](http://omeka.org/forums/topic/adding-osm-base-layer-to-neatline) actually it is possible to use just the Neatline plugin with any tile in the common `{Z}/{X}/{Y}` format.

[^2]: You may also have seen it at the New York Public Library or at Harvard's WorldMap.
