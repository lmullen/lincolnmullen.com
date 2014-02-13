---
layout: post
title: "Analyzing Historical History Dissertations: Cleaning the Data Again"
date: 2014-02-13 10:55
categories: 
---

*This series on "Analyzing Historical History Dissertations" is a work
in progress. If you would like to cite or link to it, please consider
using the **[landing page][]**, which will always have the most
up-to-date information and a list of all the posts.*

  [landing page]: http://lincolnmullen.com/research/history-dissertations/

After posts about the length of dissertations about history university
by university and for the American academy as a whole, along with a map
of where history dissertations have been written, this post will be much
less interesting. In this post I want to write about some decisions that
I've made to further clean and tidy the data, and the consequent
revisions to earlier visualizations. I'm writing about this process
because I have conceived of this series of posts as being not just a
presentation of my findings about history dissertations, but a set of
reflections on the process of data analysis as performed by a historian.
I hope this series will be pedagogically useful. Certainly no one ever
taught me how to do this kind of data analysis in a history department,
and it's not that hard to find scientists complaining that no one ever
taught them how to do this kind of work. So I hope this series of posts
and the [accompanying code][] is one [resource][] [among][] [many][],
and in particular that it will be useful as a guide to the many steps in
digging through a mid- to large-sized data set.

  [accompanying code]: https://github.com/lmullen/dissertations-data
  [resource]: http://programminghistorian.org/
  [among]: http://geospatialhistorian.wordpress.com/
  [many]: http://software-carpentry.org/

First, what do I mean by cleaning and tidying data? By cleaning data I
mean standardizing the values of the data to make them more useful. For
example, I have filtered out at least twenty thousand irrelevant
dissertations from the data set. Cleaning might also mean performing an
operation like standardizing the name of a university, or joining all
the fields that hold paragraphs for an abstract into one field. By
tidying data I mean what [Hadley Wickham][] means in this fantastic
[journal article][] and [talk][], namely, changing the structure of data
so that "each variable is a column, each observation is a row, and each
type of observational unit is a table." To put it a different way,
tidying data is [normalizing the data][]; cleaning the data is akin
to proofreading the data.

  [Hadley Wickham]: http://had.co.nz/
  [journal article]: http://vita.had.co.nz/papers/tidy-data.pdf
  [talk]: http://vimeo.com/33727555
  [normalizing the data]: http://en.wikipedia.org/wiki/Database_normalization

So why do I need to clean and tidy the data after five posts? One
possible model of how data analysis should proceed is that you first
clean the data, then analyze. Another possible model is that you first
set up the basic structure of the analysis, then clean the data to get
the results that you want. In practice, learning from data is a process
that iterates between analysis and cleaning. The process is iterative
because you have to analyze the data before you know how the data must
be cleaned and tidied, and you have to clean and tidy the data to make
certain kinds of analysis available.

These are the kinds of tidying and cleaning that can happen at this
stage.

First, [mapping the dissertations][] made it obvious that the ProQuest
data set contains very incomplete information for dissertations written
outside North America. We may as well throw away those observations,
since they are too incomplete to tell us anything useful. The data set
also contains 4,608 dissertations from Canada. Looking at the quality of
the data recorded for those Canadian dissertations, it seems that they
are worth keeping. The deciding point for me is that the [AHA Directory
of History Dissertations][] also includes Canadian dissertations.

  [mapping the dissertations]: http://lincolnmullen.com/blog/analyzing-historical-history-dissertations-location/
  [AHA Directory of History Dissertations]: https://secure.historians.org/pubs/dissertations/index.cfm

Second, looking closely at the list of universities shows that there are
some, such as the "Yale University, School of Forestry and Environmental
Studies" which represent the occasional dissertation that might be about
history, but is probably so far afield from the norm of a dissertation
in history department that these may as well be excluded. Most of these
will be dropped anyway when we filter the dissertations by subject.

Now we have a more difficult decision to make, since we have to decide
what counts as a dissertation in history. There is a strong argument to
be made that the best definition is a dissertation written in a history
department. Unfortunately, since only 7.64 percent of the dissertations
in this data set identify their department, it is impossible to use that
definition. But there are intellectual justifications for a more
capacious definition. Certainly in my own field of American religious
history, many if not most of the dissertations are written by people in
religious studies departments, divinity schools, Judaic Studies
departments, and the like. Economic history, particularly of a
quantitative bent, has mostly moved to economics departments (though
history of capitalism is making a roaring comeback in history
departments). So I have no problem with using the ProQuest subject
headings to study a broader definition of what a history dissertation
is.

The sticking point for me is whether to include dissertations with the
subject "Education, History of" as dissertations in history. The
argument that such dissertations are usually written in schools of
education rather than colleges of arts and sciences, are therefore held
to different norms and so should be excluded from this study is
persuasive to me. But we can also test the hypothesis that they are held
to different norms.

Here is a chart of the page lengths of history of education
dissertations vs all other kinds of history dissertations.

{% img /downloads/historical-dissertations/history-of-education-vs-general.png %}

There is a very wide divergence between the history of education and the
other history dissertations, so I've decided to reverse my earlier
decision and exclude them from the data set.

There is no such divergence for history of religion dissertations, so
they can stay.

{% img /downloads/historical-dissertations/history-of-religion-vs-general.png %}

History of science has a bizarre trajectory compared to every other kind
of history. (Historians of science, any ideas what is going on here?)
But even though historians of science are often in different
departments, they can stay in the data set.

{% img /downloads/historical-dissertations/history-of-science-vs-general.png %}

History of economics diverges from general history (though not as
sharply as history of education). I'm betting that's because such
dissertations are more likely to use quantitative methods, so they can
stay.

{% img /downloads/historical-dissertations/history-of-economics-vs-general.png %}

Biography also seems close to the profession in general.

{% img /downloads/historical-dissertations/biography-vs-general.png %}

As a last step, I'm dropping all universities with fewer than five
dissertations that meet all these criteria. Looking through these
universities, many of their dissertations seem like false positives, so
this cut-off should clean up problems in the long tail.

Making these changes to the data set gives us fewer dissertations: 84428
dissertations before, 80575 after cleaning, which is 3853 fewer
dissertations. It's a messy process, but then the data is messy. Keep
that in mind the next time you're impressed by a slick visualization.

* * * * * * * * * * *

Now that the data is more thoroughly cleaned, I have redone some of the
more significant visualizations from the earlier posts.

First, the number of dissertations per year.

{% img /downloads/historical-dissertations/number-of-dissertations-per-year-2.png %}

Second, the mean and median page counts of history dissertations. These are slightly higher after cleaning the data more thoroughly.


{% img /downloads/historical-dissertations/pages-mean-median-2.png %}


Third, a map of the locations where history dissertations were written.

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

$.get('/downloads/historical-dissertations/location/universities-geocoded-2.csv', function(csvContents) {
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

Finally, a redo of all the page counts by university. (Because nothing
makes a post popular like giving every historian in North America a
little blue dot of their own to look for.)

* * * * *

Columbia University.
{% img /downloads/historical-dissertations/pages-2-columbia-university.png %}

The University of Wisconsin - Madison.
{% img /downloads/historical-dissertations/pages-2-the-university-of-wisconsin---madison.png %}

University of California, Berkeley.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-berkeley.png %}

Yale University.
{% img /downloads/historical-dissertations/pages-2-yale-university.png %}

University of California, Los Angeles.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-los-angeles.png %}

Harvard University.
{% img /downloads/historical-dissertations/pages-2-harvard-university.png %}

New York University.
{% img /downloads/historical-dissertations/pages-2-new-york-university.png %}

University of Michigan.
{% img /downloads/historical-dissertations/pages-2-university-of-michigan.png %}

University of Pennsylvania.
{% img /downloads/historical-dissertations/pages-2-university-of-pennsylvania.png %}

The University of Chicago.
{% img /downloads/historical-dissertations/pages-2-the-university-of-chicago.png %}

Princeton University.
{% img /downloads/historical-dissertations/pages-2-princeton-university.png %}

Indiana University.
{% img /downloads/historical-dissertations/pages-2-indiana-university.png %}

University of Minnesota.
{% img /downloads/historical-dissertations/pages-2-university-of-minnesota.png %}

Stanford University.
{% img /downloads/historical-dissertations/pages-2-stanford-university.png %}

The University of Texas at Austin.
{% img /downloads/historical-dissertations/pages-2-the-university-of-texas-at-austin.png %}

The University of North Carolina at Chapel Hill.
{% img /downloads/historical-dissertations/pages-2-the-university-of-north-carolina-at-chapel-hill.png %}

University of Virginia.
{% img /downloads/historical-dissertations/pages-2-university-of-virginia.png %}

The Ohio State University.
{% img /downloads/historical-dissertations/pages-2-the-ohio-state-university.png %}

Duke University.
{% img /downloads/historical-dissertations/pages-2-duke-university.png %}

The Johns Hopkins University.
{% img /downloads/historical-dissertations/pages-2-the-johns-hopkins-university.png %}

University of Illinois at Urbana-Champaign.
{% img /downloads/historical-dissertations/pages-2-university-of-illinois-at-urbana-champaign.png %}

Cornell University.
{% img /downloads/historical-dissertations/pages-2-cornell-university.png %}

City University of New York.
{% img /downloads/historical-dissertations/pages-2-city-university-of-new-york.png %}

Northwestern University.
{% img /downloads/historical-dissertations/pages-2-northwestern-university.png %}

University of Washington.
{% img /downloads/historical-dissertations/pages-2-university-of-washington.png %}

Rutgers The State University of New Jersey - New Brunswick.
{% img /downloads/historical-dissertations/pages-2-rutgers-the-state-university-of-new-jersey---new-brunswick.png %}

University of Maryland, College Park.
{% img /downloads/historical-dissertations/pages-2-university-of-maryland,-college-park.png %}

University of California, Santa Barbara.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-santa-barbara.png %}

The University of Iowa.
{% img /downloads/historical-dissertations/pages-2-the-university-of-iowa.png %}

Emory University.
{% img /downloads/historical-dissertations/pages-2-emory-university.png %}

The Florida State University.
{% img /downloads/historical-dissertations/pages-2-the-florida-state-university.png %}

Brown University.
{% img /downloads/historical-dissertations/pages-2-brown-university.png %}

University of Kansas.
{% img /downloads/historical-dissertations/pages-2-university-of-kansas.png %}

University of Pittsburgh.
{% img /downloads/historical-dissertations/pages-2-university-of-pittsburgh.png %}

Temple University.
{% img /downloads/historical-dissertations/pages-2-temple-university.png %}

Michigan State University.
{% img /downloads/historical-dissertations/pages-2-michigan-state-university.png %}

University of Notre Dame.
{% img /downloads/historical-dissertations/pages-2-university-of-notre-dame.png %}

Boston University.
{% img /downloads/historical-dissertations/pages-2-boston-university.png %}

Georgetown University.
{% img /downloads/historical-dissertations/pages-2-georgetown-university.png %}

University of Southern California.
{% img /downloads/historical-dissertations/pages-2-university-of-southern-california.png %}

Brandeis University.
{% img /downloads/historical-dissertations/pages-2-brandeis-university.png %}

University of Missouri - Columbia.
{% img /downloads/historical-dissertations/pages-2-university-of-missouri---columbia.png %}

State University of New York at Binghamton.
{% img /downloads/historical-dissertations/pages-2-state-university-of-new-york-at-binghamton.png %}

The Pennsylvania State University.
{% img /downloads/historical-dissertations/pages-2-the-pennsylvania-state-university.png %}

State University of New York at Buffalo.
{% img /downloads/historical-dissertations/pages-2-state-university-of-new-york-at-buffalo.png %}

Vanderbilt University.
{% img /downloads/historical-dissertations/pages-2-vanderbilt-university.png %}

University of California, San Diego.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-san-diego.png %}

The University of Arizona.
{% img /downloads/historical-dissertations/pages-2-the-university-of-arizona.png %}

The Catholic University of America.
{% img /downloads/historical-dissertations/pages-2-the-catholic-university-of-america.png %}

University of Colorado at Boulder.
{% img /downloads/historical-dissertations/pages-2-university-of-colorado-at-boulder.png %}

Syracuse University.
{% img /downloads/historical-dissertations/pages-2-syracuse-university.png %}

University of Florida.
{% img /downloads/historical-dissertations/pages-2-university-of-florida.png %}

Saint Louis University.
{% img /downloads/historical-dissertations/pages-2-saint-louis-university.png %}

The University of New Mexico.
{% img /downloads/historical-dissertations/pages-2-the-university-of-new-mexico.png %}

The Claremont Graduate University.
{% img /downloads/historical-dissertations/pages-2-the-claremont-graduate-university.png %}

The George Washington University.
{% img /downloads/historical-dissertations/pages-2-the-george-washington-university.png %}

Tulane University.
{% img /downloads/historical-dissertations/pages-2-tulane-university.png %}

University of Massachusetts Amherst.
{% img /downloads/historical-dissertations/pages-2-university-of-massachusetts-amherst.png %}

University of California, Irvine.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-irvine.png %}

State University of New York at Stony Brook.
{% img /downloads/historical-dissertations/pages-2-state-university-of-new-york-at-stony-brook.png %}

University of Rochester.
{% img /downloads/historical-dissertations/pages-2-university-of-rochester.png %}

University of Kentucky.
{% img /downloads/historical-dissertations/pages-2-university-of-kentucky.png %}

University of South Carolina.
{% img /downloads/historical-dissertations/pages-2-university-of-south-carolina.png %}

University of Delaware.
{% img /downloads/historical-dissertations/pages-2-university-of-delaware.png %}

University of Hawai'i at Manoa.
{% img /downloads/historical-dissertations/pages-2-university-of-hawai'i-at-manoa.png %}

The University of Oklahoma.
{% img /downloads/historical-dissertations/pages-2-the-university-of-oklahoma.png %}

University of California, Davis.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-davis.png %}

Arizona State University.
{% img /downloads/historical-dissertations/pages-2-arizona-state-university.png %}

University of Connecticut.
{% img /downloads/historical-dissertations/pages-2-university-of-connecticut.png %}

University of Oregon.
{% img /downloads/historical-dissertations/pages-2-university-of-oregon.png %}

Louisiana State University and Agricultural & Mechanical College.
{% img /downloads/historical-dissertations/pages-2-louisiana-state-university-and-agricultural-&-mechanical-college.png %}

University of Georgia.
{% img /downloads/historical-dissertations/pages-2-university-of-georgia.png %}

American University.
{% img /downloads/historical-dissertations/pages-2-american-university.png %}

Washington University in St. Louis.
{% img /downloads/historical-dissertations/pages-2-washington-university-in-st.-louis.png %}

Boston College.
{% img /downloads/historical-dissertations/pages-2-boston-college.png %}

The University of Nebraska - Lincoln.
{% img /downloads/historical-dissertations/pages-2-the-university-of-nebraska---lincoln.png %}

University of California, Riverside.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-riverside.png %}

Case Western Reserve University.
{% img /downloads/historical-dissertations/pages-2-case-western-reserve-university.png %}

Drew University.
{% img /downloads/historical-dissertations/pages-2-drew-university.png %}

Fordham University.
{% img /downloads/historical-dissertations/pages-2-fordham-university.png %}

Texas A&M University.
{% img /downloads/historical-dissertations/pages-2-texas-a&m-university.png %}

The University of Utah.
{% img /downloads/historical-dissertations/pages-2-the-university-of-utah.png %}

Rice University.
{% img /downloads/historical-dissertations/pages-2-rice-university.png %}

Washington State University.
{% img /downloads/historical-dissertations/pages-2-washington-state-university.png %}

Kent State University.
{% img /downloads/historical-dissertations/pages-2-kent-state-university.png %}

University of Illinois at Chicago.
{% img /downloads/historical-dissertations/pages-2-university-of-illinois-at-chicago.png %}

Purdue University.
{% img /downloads/historical-dissertations/pages-2-purdue-university.png %}

University of Cincinnati.
{% img /downloads/historical-dissertations/pages-2-university-of-cincinnati.png %}

Loyola University Chicago.
{% img /downloads/historical-dissertations/pages-2-loyola-university-chicago.png %}

Howard University.
{% img /downloads/historical-dissertations/pages-2-howard-university.png %}

Texas Christian University.
{% img /downloads/historical-dissertations/pages-2-texas-christian-university.png %}

Ohio University.
{% img /downloads/historical-dissertations/pages-2-ohio-university.png %}

Marquette University.
{% img /downloads/historical-dissertations/pages-2-marquette-university.png %}

The University of Tennessee.
{% img /downloads/historical-dissertations/pages-2-the-university-of-tennessee.png %}

Boston University Graduate School.
{% img /downloads/historical-dissertations/pages-2-boston-university-graduate-school.png %}

The University of Alabama.
{% img /downloads/historical-dissertations/pages-2-the-university-of-alabama.png %}

Graduate Theological Union.
{% img /downloads/historical-dissertations/pages-2-graduate-theological-union.png %}

The College of William and Mary.
{% img /downloads/historical-dissertations/pages-2-the-college-of-william-and-mary.png %}

Bryn Mawr College.
{% img /downloads/historical-dissertations/pages-2-bryn-mawr-college.png %}

University of California, Santa Cruz.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-santa-cruz.png %}

University of North Texas.
{% img /downloads/historical-dissertations/pages-2-university-of-north-texas.png %}

The Southern Baptist Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-the-southern-baptist-theological-seminary.png %}

West Virginia University.
{% img /downloads/historical-dissertations/pages-2-west-virginia-university.png %}

University of Houston.
{% img /downloads/historical-dissertations/pages-2-university-of-houston.png %}

Bowling Green State University.
{% img /downloads/historical-dissertations/pages-2-bowling-green-state-university.png %}

University of Denver.
{% img /downloads/historical-dissertations/pages-2-university-of-denver.png %}

Texas Tech University.
{% img /downloads/historical-dissertations/pages-2-texas-tech-university.png %}

Auburn University.
{% img /downloads/historical-dissertations/pages-2-auburn-university.png %}

State University of New York at Albany.
{% img /downloads/historical-dissertations/pages-2-state-university-of-new-york-at-albany.png %}

Northern Illinois University.
{% img /downloads/historical-dissertations/pages-2-northern-illinois-university.png %}

The Union Institute.
{% img /downloads/historical-dissertations/pages-2-the-union-institute.png %}

Southern Illinois University at Carbondale.
{% img /downloads/historical-dissertations/pages-2-southern-illinois-university-at-carbondale.png %}

Southwestern Baptist Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-southwestern-baptist-theological-seminary.png %}

Wayne State University.
{% img /downloads/historical-dissertations/pages-2-wayne-state-university.png %}

The University of Mississippi.
{% img /downloads/historical-dissertations/pages-2-the-university-of-mississippi.png %}

Miami University.
{% img /downloads/historical-dissertations/pages-2-miami-university.png %}

New School for Social Research.
{% img /downloads/historical-dissertations/pages-2-new-school-for-social-research.png %}

Princeton Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-princeton-theological-seminary.png %}

Oklahoma State University.
{% img /downloads/historical-dissertations/pages-2-oklahoma-state-university.png %}

The University of Maine.
{% img /downloads/historical-dissertations/pages-2-the-university-of-maine.png %}

University of Arkansas.
{% img /downloads/historical-dissertations/pages-2-university-of-arkansas.png %}

University of Miami.
{% img /downloads/historical-dissertations/pages-2-university-of-miami.png %}

George Mason University.
{% img /downloads/historical-dissertations/pages-2-george-mason-university.png %}

Georgia State University.
{% img /downloads/historical-dissertations/pages-2-georgia-state-university.png %}

Carnegie Mellon University.
{% img /downloads/historical-dissertations/pages-2-carnegie-mellon-university.png %}

Fuller Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-fuller-theological-seminary.png %}

Mississippi State University.
{% img /downloads/historical-dissertations/pages-2-mississippi-state-university.png %}

University of New Hampshire.
{% img /downloads/historical-dissertations/pages-2-university-of-new-hampshire.png %}

Baylor University.
{% img /downloads/historical-dissertations/pages-2-baylor-university.png %}

Kansas State University.
{% img /downloads/historical-dissertations/pages-2-kansas-state-university.png %}

Clark University.
{% img /downloads/historical-dissertations/pages-2-clark-university.png %}

Iowa State University.
{% img /downloads/historical-dissertations/pages-2-iowa-state-university.png %}

The University of Toledo.
{% img /downloads/historical-dissertations/pages-2-the-university-of-toledo.png %}

The University of Southern Mississippi.
{% img /downloads/historical-dissertations/pages-2-the-university-of-southern-mississippi.png %}

Ball State University.
{% img /downloads/historical-dissertations/pages-2-ball-state-university.png %}

Union Institute and University.
{% img /downloads/historical-dissertations/pages-2-union-institute-and-university.png %}

New School University.
{% img /downloads/historical-dissertations/pages-2-new-school-university.png %}

Lehigh University.
{% img /downloads/historical-dissertations/pages-2-lehigh-university.png %}

The University of Wisconsin - Milwaukee.
{% img /downloads/historical-dissertations/pages-2-the-university-of-wisconsin---milwaukee.png %}

The University of Texas at Arlington.
{% img /downloads/historical-dissertations/pages-2-the-university-of-texas-at-arlington.png %}

Tufts University.
{% img /downloads/historical-dissertations/pages-2-tufts-university.png %}

Brigham Young University.
{% img /downloads/historical-dissertations/pages-2-brigham-young-university.png %}

The University of Memphis.
{% img /downloads/historical-dissertations/pages-2-the-university-of-memphis.png %}

Northern Arizona University.
{% img /downloads/historical-dissertations/pages-2-northern-arizona-university.png %}

The University of Texas at Dallas.
{% img /downloads/historical-dissertations/pages-2-the-university-of-texas-at-dallas.png %}

Union Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-union-theological-seminary.png %}

Southern Methodist University.
{% img /downloads/historical-dissertations/pages-2-southern-methodist-university.png %}

Virginia Polytechnic Institute and State University.
{% img /downloads/historical-dissertations/pages-2-virginia-polytechnic-institute-and-state-university.png %}

Westminster Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-westminster-theological-seminary.png %}

Duquesne University.
{% img /downloads/historical-dissertations/pages-2-duquesne-university.png %}

Western Michigan University.
{% img /downloads/historical-dissertations/pages-2-western-michigan-university.png %}

University of Idaho.
{% img /downloads/historical-dissertations/pages-2-university-of-idaho.png %}

California Institute of Integral Studies.
{% img /downloads/historical-dissertations/pages-2-california-institute-of-integral-studies.png %}

Regent University.
{% img /downloads/historical-dissertations/pages-2-regent-university.png %}

The University of Akron.
{% img /downloads/historical-dissertations/pages-2-the-university-of-akron.png %}

The Jewish Theological Seminary of America.
{% img /downloads/historical-dissertations/pages-2-the-jewish-theological-seminary-of-america.png %}

University of Missouri - Kansas City.
{% img /downloads/historical-dissertations/pages-2-university-of-missouri---kansas-city.png %}

Florida International University.
{% img /downloads/historical-dissertations/pages-2-florida-international-university.png %}

Pacifica Graduate Institute.
{% img /downloads/historical-dissertations/pages-2-pacifica-graduate-institute.png %}

University of Nevada, Reno.
{% img /downloads/historical-dissertations/pages-2-university-of-nevada,-reno.png %}

Virginia Commonwealth University.
{% img /downloads/historical-dissertations/pages-2-virginia-commonwealth-university.png %}

The University of North Carolina at Greensboro.
{% img /downloads/historical-dissertations/pages-2-the-university-of-north-carolina-at-greensboro.png %}

Trinity Evangelical Divinity School.
{% img /downloads/historical-dissertations/pages-2-trinity-evangelical-divinity-school.png %}

New Orleans Baptist Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-new-orleans-baptist-theological-seminary.png %}

University of Louisville.
{% img /downloads/historical-dissertations/pages-2-university-of-louisville.png %}

Hebrew Union College - Jewish Institute of Religion (Ohio).
{% img /downloads/historical-dissertations/pages-2-hebrew-union-college---jewish-institute-of-religion-(ohio).png %}

Georgia Institute of Technology.
{% img /downloads/historical-dissertations/pages-2-georgia-institute-of-technology.png %}

Andrews University.
{% img /downloads/historical-dissertations/pages-2-andrews-university.png %}

Northeastern University.
{% img /downloads/historical-dissertations/pages-2-northeastern-university.png %}

Southeastern Baptist Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-southeastern-baptist-theological-seminary.png %}

Yeshiva University.
{% img /downloads/historical-dissertations/pages-2-yeshiva-university.png %}

The Iliff School of Theology and University of Denver.
{% img /downloads/historical-dissertations/pages-2-the-iliff-school-of-theology-and-university-of-denver.png %}

University of Nevada, Las Vegas.
{% img /downloads/historical-dissertations/pages-2-university-of-nevada,-las-vegas.png %}

Oregon State University.
{% img /downloads/historical-dissertations/pages-2-oregon-state-university.png %}

Rensselaer Polytechnic Institute.
{% img /downloads/historical-dissertations/pages-2-rensselaer-polytechnic-institute.png %}

Mid-America Baptist Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-mid-america-baptist-theological-seminary.png %}

Memphis State University.
{% img /downloads/historical-dissertations/pages-2-memphis-state-university.png %}

University of California, San Francisco.
{% img /downloads/historical-dissertations/pages-2-university-of-california,-san-francisco.png %}

Texas Woman's University.
{% img /downloads/historical-dissertations/pages-2-texas-woman's-university.png %}

Colorado State University.
{% img /downloads/historical-dissertations/pages-2-colorado-state-university.png %}

Salve Regina University.
{% img /downloads/historical-dissertations/pages-2-salve-regina-university.png %}

Indiana University of Pennsylvania.
{% img /downloads/historical-dissertations/pages-2-indiana-university-of-pennsylvania.png %}

University of South Florida.
{% img /downloads/historical-dissertations/pages-2-university-of-south-florida.png %}

Walden University.
{% img /downloads/historical-dissertations/pages-2-walden-university.png %}

Lutheran School of Theology at Chicago.
{% img /downloads/historical-dissertations/pages-2-lutheran-school-of-theology-at-chicago.png %}

Middle Tennessee State University.
{% img /downloads/historical-dissertations/pages-2-middle-tennessee-state-university.png %}

University of Wyoming.
{% img /downloads/historical-dissertations/pages-2-university-of-wyoming.png %}

Capella University.
{% img /downloads/historical-dissertations/pages-2-capella-university.png %}

Dallas Theological Seminary.
{% img /downloads/historical-dissertations/pages-2-dallas-theological-seminary.png %}

University of New Orleans.
{% img /downloads/historical-dissertations/pages-2-university-of-new-orleans.png %}

