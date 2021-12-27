/**
 * Draw maps when slides are activated
 */
function ready() {
  mystack = stack().on("activate", activate).on("deactivate", deactivate);
  congregationalistsIndex = findIndex("#congregationalists");
  diachronicIndex = findIndex("#diachronic");
  gaustadIndex = findIndex("#gaustad-slide");
  classMapIndex = findIndex("#class-map");
}

function findIndex(selector) {
  var section = d3.selectAll("section");
  var node = d3.select(selector);
  var index = section[0].indexOf(node.node());
  return index;
}

function activate(d, i) {
  switch(i) {
    case congregationalistsIndex:
      drawCongregationalists();
      break;
    case diachronicIndex:
      drawDiachronic();
      break;
    case gaustadIndex:
      drawGaustad();
      break;
    case classMapIndex:
      embedClassMap();
      break;
  }
}

function deactivate(d, i) {
  switch(i) {
    case congregationalistsIndex:
      deleteCongregationalists();
      break;
    case diachronicIndex:
      deleteDiachronic();
      break;
    case gaustadIndex:
      deleteGaustad();
      break;
    case classMapIndex:
      deleteClassMap();
      break;
  }
}

function hideCommentary() {
  d3.selectAll("section.text").remove();
  d3.selectAll("#title aside").remove()
  ready();
}

/**
 * Map of Congregationalists in 1854
 */

function drawCongregationalists() {
  congregationalistsMap = d3.carto.map();
  d3.select("#congregationalists-map").call(congregationalistsMap);
  congregationalistsMap
  .centerOn([-75.668006, 42.9], "latlong")
  .setScale(6);

  baseLayer = d3.carto.layer.tile();
  baseLayer
  .path("terrain-background")
  .tileType("stamen")
  .label("Base");

  states = d3.carto.layer.topojson();
  states
  .path("us.topojson")
  .renderMode("canvas")
  .cssClass("county")
  .label("States");

  churches = d3.carto.layer.csv();

  churchesModal = d3.carto.modal();
  churchesModal.formatter(function(d) {
    return "<p>" + d.church + " " + d.city + ", " + d.state + "</p>" +
      "<p>Members: " + d.members + "</p>" +
      "<p>Minister: " + d.minister + "</p>" +
      "<p>Organized: " + d.organized + "</p>" +
      "<p>Citation: <a target='_blank' href='" + d.url + "'>" + d.citation + "</a></p>"
  })

  churches
  .path("congregationalists.csv")
  .label("Churches")
  .cssClass("church")
  .renderMode("svg")
  .x("jlon")
  .y("jlat")
  .markerSize(0)
  .modal(churchesModal)
  .clickableFeatures(true)
  .on("load", scaleChurchCircles);

  congregationalistsMap
  .addCartoLayer(baseLayer)
  .addCartoLayer(states)
  .addCartoLayer(churches);

  d3.select("#united-states").on("click", function() {
    var us = [[-124.73292, 24.54512], [-66.94993, 49.38437]];
    congregationalistsMap.zoomTo(us, "latlong", 0.9, 5000)
  })


  d3.select("#new-england").on("click", function() {
    var ne = [[-67.701452, 47.133902], [-73.612097,41.067165]];
    congregationalistsMap.zoomTo(ne, "latlong", 0.9, 4000);
  })

  d3.select("#massachusetts").on("click", function() {
    var ma = [[-73.175799, 42.713157], [-69.855884, 41.555233]];
    congregationalistsMap.zoomTo(ma, "latlong", 0.9, 4000);
  })

  d3.select("#boston").on("click", function() {
    var gb = [[-71.202340, 42.409350], [-70.935921, 42.261137]];
    congregationalistsMap.zoomTo(gb, "latlong", 0.3, 4000);
  })

  d3.select("#connecticut").on("click", function() {
    var ct = [[-73.72778, 40.98886], [-71.78765, 42.04965]];
    congregationalistsMap.zoomTo(ct, "latlong", 0.9, 4000);
  })

  d3.select("#midwest").on("click", function() {
    var midwest = [[-97.056921, 46.955181], [-74.249304, 41.463548]];
    congregationalistsMap.zoomTo(midwest, "latlong", 0.9, 4000);
  })

  d3.select("#ohio").on("click", function() {
    var ohio = [[-84.82019, 38.41079], [-80.51880, 41.97753]];
    congregationalistsMap.zoomTo(ohio, "latlong", 0.7, 4000);
  })

}

function scaleChurchCircles() {
  var memberScale = d3.scale.sqrt().domain([0,800]).range([1,8]).clamp(true);
  churches.g().selectAll("circle")
  .transition()
  .attr("r", function(d) {
    if(d.members > 0) {
      return memberScale(d.members);
    } else {
      return 2;
    }
  })
  .duration(2500)
  .delay(function(d, i) { return i / 2500 * 2500; });

}

function deleteCongregationalists() {
  d3.select("#congregationalists-map").selectAll("*").remove();
}

/**
 * Map of Congregationalist churches over time
 */

function drawDiachronic() {
  diachronicMap = d3.carto.map();
  d3.select("#diachronic-map").call(diachronicMap);
  diachronicMap
  .centerOn([-71.668006, 42.239650], "latlong")
  .setScale(7);


  baseLayer = d3.carto.layer.tile();
  baseLayer
  .path("terrain-background")
  .tileType("stamen")
  .label("Base");

  churchesDiachronic = d3.carto.layer.csv();

  churchesDiachronic
  .path("congregationalists.csv")
  .label("Churches")
  .cssClass("church")
  .renderMode("svg")
  .x("jlon")
  .y("jlat")
  .markerSize(0)
  .clickableFeatures(false)
  .on("load", churchesOverTime);

  diachronicMap
  .addCartoLayer(baseLayer)
  .addCartoLayer(churchesDiachronic)
  .zoomable(false);

}

function churchesOverTime() {

  var date = 1630;

  timer = setInterval(function() {

    churchesDiachronic.g().selectAll("circle")
    .transition()
    .attr("r", function(d) {
      if(+d.organized <= date)
        return 4;
      else
        return 0;
    })
    .duration(250)
    .ease("linear");

    date++;
    d3.select("#diachronic-date").text(date);

    if(date >= 1853) clearInterval(timer);

  }, 250)
}


function deleteDiachronic() {
  clearInterval(timer);
  d3.select("#diachronic-map").selectAll("*").remove();
}


/**
 * Imitation of Gaustad map
 */

function drawGaustad() {
  gaustadMap = d3.carto.map();
  d3.select("#gaustad-map").call(gaustadMap);

  census = d3.carto.layer.topojson();
  census
  .path("census.topojson")
  .label("Census")
  .renderMode("svg")
  .cssClass("county")
  .clickableFeatures(true)
  .on("load", choroplethCensus);

  gaustadMap.addCartoLayer(census);

  gaustadMap.zoomTo([[-99.32, 25.68], [-69.75, 47.51]], "latlong", 0.55);

}

function choroplethCensus() {
  censusScale = d3.scale.threshold()
  .domain([1, 5, 15, 26])
  .range(["#ffffff", "#dfc27d", "#bf812d", "#8c510a", "#543005"])

  census.g().selectAll("path.county").style("fill", function(d) {
    return censusScale(d.properties.cong);
  }).style("fill-opacity", 1);

  // var albers = d3.geo.albersUsa()
  // .scale(1000)
  // .translate(gaustadMap.zoom().translate());

  // gaustadMap.mode("projection").projection(albers).refresh();


}

function deleteGaustad() {
  d3.select("#gaustad-map").selectAll("*").remove();
}


/**
 * Class map
 */

function embedClassMap() {
  var iframe = d3.select("#class-map")
  .append("iframe")
  .attr("src", "http://omeka.lts.brandeis.edu/neatline/fullscreen/mapping-bostons-religions")
}

function deleteClassMap() {
  d3.select("#class-map").select("iframe").remove();
}
