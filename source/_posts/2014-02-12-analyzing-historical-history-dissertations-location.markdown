---
layout: post
title: "Analyzing Historical History Dissertations: Location"
date: 2014-02-12 08:00
categories: 
- r-language
- leaflet-js
- digital-humanities
- historical-dissertations
---

*Earlier posts about "Analyzing Historical History Dissertations": 
[Beginnings][] | [Limitations of the Data][] | [Page Counts][] | [Page 
  Counts by University][].*

So far I have mostly looked at the number of dissertations, and the page
counts of dissertations. I'd like to refine the analysis of page counts
by seeing if the length of dissertations varies by [NRC rankings of
history departments][].  Today, though, I want to look at the locations
where history dissertations were produced.

Below is a map of the universities that have produced dissertations
about history. In case the map isn't self-explanatory, the numbers in
red boxes represent the total number of dissertations produced in a
geographic area. You can put your mouse over the box to see the
boundaries of the geographic area. Zooming in or clicking on the box
will cluster the points into smaller and smaller units. Eventually
you'll get individual markers for each university. Clicking on those
markers will tell you the university name and the number of history
dissertations is has produced.

<style>
#map {
  height: 500px;
  width: 800px;
  margin-top: 20px;
  margin-bottom: 20px;
  border: 1px solid gray;
}

#map img {
  border:0;
  -webkit-box-shadow: none;
  box-shadow: none;
}

.mycluster {
  width: 40px;
  background-color: lightcoral;
  text-align: center;
  font-size: 14px;
}
</style>

<script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
<script>
$("head").append('<link rel="stylesheet" href="http://cdn.leafletjs.com/leaflet-0.7.2/leaflet.css" />');

$("head").append('<link rel="stylesheet" href="/downloads/historical-dissertations/location/MarkerCluster.Default.css" />');
</script>

<script src="http://cdn.leafletjs.com/leaflet-0.7.2/leaflet.js"></script>
<link rel="stylesheet" href="/downloads/historical-dissertations/location/MarkerCluster.css" />
<script src="/downloads/historical-dissertations/location/leaflet.markercluster-src.js"></script>
<script src="/downloads/historical-dissertations/location/leaflet.geocsv.js"></script>

<div id="map"></div>

<script>
var map = L.map('map');
var osmUrl='http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
var osmAttrib='Map data (c) OpenStreetMap contributors';
var osm = new L.TileLayer(osmUrl, {attribution: osmAttrib});		

// start the map in the center of the United States
map.setView(new L.LatLng(39.8, -95), 4);
map.addLayer(osm);

var markers = new L.MarkerClusterGroup({
  iconCreateFunction: function (cluster) {
    var markers = cluster.getAllChildMarkers();
    var n = 0;
    for (var i = 0; i < markers.length; i++) {
      n += +markers[i].feature.properties["count"];
    }
    return L.divIcon({ html: n, className: 'mycluster', iconSize: L.point(50, 20) });
  },
});

$.get('/downloads/historical-dissertations/location/universities-geocoded.csv', function(csvContents) {
  var geoLayer = L.geoCsv(csvContents, {
    titles: ["","university","count","lon","lat"],
    firstLineTitles: true, 
    fieldSeparator: ',',
    latitudeTitle: 'lat',
    longitudeTitle: 'lon',
    onEachFeature: function (feature, layer) {
      var popup = '';
      popup += "<strong>" + feature.properties["university"] + "</strong>";
      popup += "<br/>";
      popup += "Dissertations about history: "
      popup += feature.properties["count"] + ".";
      layer.bindPopup(popup);
    }
  });
  markers.addLayer(geoLayer);
  map.addLayer(markers);
});
</script>

Unsurprisingly, the Northeast can be thought of as a mega-center, with 
Boston, New Haven, New York, Philadelphia, and the areas in between 
producing on the order of 28 thousand history dissertations. But there 
are many other smaller centers: Washington, Chicago, Michigan, the North 
Carolina research triangle, the University of California and Stanford. 
The University of Wisconsin is a behemoth unto itself, and other 
universities (e.g., the Indiana University at Bloomington) produce many 
dissertations. Columbia University has the most dissertations in the 
data set. 

The main observations I would make are these. First, the Northeast
produces a lot of dissertations. But second, the writing of history
dissertations is scattered across the country, and regions like the
Midwest, the upper Southeast, southern California, and Texas are sizable
in themselves. And finally, a few schools seem to produce most of the
output of history dissertations. Very roughly speaking, 25 schools have
produced about half the dissertations in history, while another 350 or
so have produced the other half. (See the chart at the end of [this
post][].) 

In future posts I want to look at not just the places where historians
write, but the places they write about. I'm curious to see if I can
detect some correlation between the universities that produce
dissertations and the places those dissertations are about. Surely there
is a reason for all these dissertations about Boston and Chicago.

* * * * *

Caveats: If you zoom the map out, you'll see that there is some cruft in
the data. There are 341 dissertations in Asia, 55 in Australia, 94 in
Europe, 6 in Africa, and 2 in South America. There are also 4,608
dissertations in Canada. Clearly this data set does not contain all the
dissertations from outside North America. It's not clear to me whether
the data set is comprehensive for Canadian dissertations or not. I knew
that there were some non-U.S. dissertations in the data set, but I
didn't know how many or from where exactly until I made this map. For
future analysis I'm going to trim all the dissertations from outside the
United States, and go back and redo the earlier analysis, though the
difference will likely be minimal. That's why I practice [reproducible
research][]. And as I observed in the post on [page counts by
university][], the data set only has information for some universities
after the 1980s or 1990s. That will skew the numbers on this map down.

  [reproducible research]: http://christophergandrud.github.io/RepResR-RStudio/
  [Beginnings]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-beginnings/
  [Limitations of the Data]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-limitations-of-the-data/
  [Page Counts]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts/
  [Page Counts by University]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts-by-university/
  [NRC rankings of history departments]: https://www.historians.org/publications-and-directories/perspectives-on-history/december-2010/nrc-report-provides-data-on-history-doctoral-programs
  [this post]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts-by-university/
  [page counts by university]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-page-counts-by-university/

