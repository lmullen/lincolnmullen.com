queue()
  .defer(d3.json, "us.json")
  .defer(d3.csv,  "gender-census.csv")
  .await(ready);

var margin = {top: 10, right: 10, bottom: 10, left: 10},
    width = 960 - margin.left - margin.right,
    height = 500 - margin.top - margin.bottom,
    sliderHeight = 20,
    sliderWidth = 600,
    brewer = ["", "q0-9", "q1-9", "q2-9", "q3-9", "q4-9",
              "q5-9", "q6-9", "q7-9", "q8-9"],
    data = {},
    percentageFormat = d3.format(".2%");

// Map of maps
var maps = {
  "sexRatio": {
    "field": "diff",
    "label": "Divergence in Sex Ratio",
    "color": "BrBG",
    "scale": d3.scale.threshold()
            .domain([-0.5,-0.25,-0.1,-0.05,-0.01,0.01,0.05,0.1,0.25, 0.5])
            .range(brewer),
    "scaleMale": d3.scale.linear()
            .domain([-0.5, 0])
            .range([d3.rgb("#00441b"), d3.rgb("#ffffff")]),
    "scaleFemale": d3.scale.linear()
            .domain([0.5, 0])
            .range([d3.rgb("#40004b"), d3.rgb("#ffffff")]),
  }
};

var current = { "year": 1820, "map" : maps.sexRatio };

var svg = d3.select("#viz").append("svg")
      .attr("width", width + margin.left + margin.right)
      .attr("height", height + margin.top + margin.bottom)
    .append("g")
      .attr("transform", "translate(" + margin.left + "," + margin.top + ")")
      .classed("YlOrRd", true);

var loading = svg.append("text")
    .attr("x", width / 2)
    .attr("y", height / 2)
    .attr("text-anchor", "middle")
    .text("Loading the map ...");

// Year slider
var sliderContainer = d3.select("#slider").append("svg")
      .attr("width", sliderWidth + margin.left + margin.right)
      .attr("height", sliderHeight + margin.top + margin.bottom)
    .append("g")
      .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

var x = d3.scale.linear()
      .domain([1815, 2014])
      .range([0, sliderWidth])
      .clamp(true);

var brushToYear = d3.scale.threshold()
      .domain(d3.range(1820,2020,10))
      .range(d3.range(1820,2020,10));

var brush = d3.svg.brush()
  .x(x)
  .extent([current.year, current.year])
  .on("brush", brushed);

sliderContainer.append("g")
  .attr("class", "x axis")
  .attr("transform", "translate(0," + sliderHeight / 2 + ")")
  .call(d3.svg.axis()
          .scale(x)
          .orient("bottom")
          .ticks(20)
          .tickFormat(function(d) { return d; })
          .tickSize(0)
          .tickPadding(12))
          .select(".domain")
          .select(function() { 
            return this.parentNode.appendChild(this.cloneNode(true)); 
          })
          .attr("class", "halo");

var slider = sliderContainer.append("g")
  .attr("class", "slider")
  .call(brush);

slider.selectAll(".extent, .resize").remove();

slider.select(".background").attr("height", sliderHeight);

var handle = slider.append("circle")
  .attr("class", "handle")
  .attr("transform", "translate(0," + sliderHeight / 2 + ")")
  .attr("r", 9);

// Field selector
var fieldSelector = d3.select("#field-selector")
      .on("change", fieldSelected);

for (var key in maps) {
  fieldSelector.append("option")
  .attr("value", key)
  .text(maps[key].label);
}

var labels = ["75-100% female", "60-75% female", "55-60% female", "51-55% female","even","51-55% male", "55-60% male", "60-75% male", "75-100% male"];
//
// Legend
var legend = svg.append("g")
  .attr("transform", "translate(" + (width - 190) + "," + 200 + ")")
  .classed("legend", true);

var legendDate = legend.append("g")
  .attr("transform", "translate(0,0)")
  .append("text")
  .classed("date-label", "true");

var legendField = legend.append("g")
  .attr("transform", "translate(0,20)")
  .append("text");

var legendColors = legend.append("g")
  .attr("class","legend-colors")
  .attr("transform", "translate(10,40)");

// Map functions
var path = d3.geo.path().projection(null);

// Tooltip
var tooltip = d3.select("body").append("div")
  .classed("tooltip", true)
  .classed("hidden", true);

function ready(error, us, census) { 

  if (error) {
    loading.text("Sorry, there has been an error. " +
                 "Please refresh and try again.");
    console.log(error);
  }

  data.us     = us;
  data.census = d3.nest()
    .key( function(d) {return d.gisjoin;})
    .key( function(d) {return d.year;})
    .map(census, d3.map);

  // Draw the map for the first time
  slider.call(brush.event).call(brush.extent([1820, 1820])).call(brush.event);
  drawMap(current.year, current.map);
  loading.remove();

}  

function brushed() {
  var value = brush.extent()[0];
  if (d3.event.sourceEvent) { // not a programmatic event
    value = x.invert(d3.mouse(this)[0]);
    brush.extent([value, value]);
  }
  handle.attr("cx", x(value));
  var brushDate = brushToYear(value-5);

  // Only redraw the map when we cross the threshold to a new year
  if (brushDate !== current.year) {
    legendDate.text(brushDate + " Census");
    current.year = brushDate;
    drawMap(brushDate, current.map);
  }
}

function drawMap(date, map) {

  var counties = topojson.feature(data.us,
                                  data.us.objects["county_" + current.year]);

  svg.attr("class", map.color);
  svg.selectAll(".counties, .states, .country").remove();

  svg.selectAll(".counties") 
    .data(counties.features)
    .enter()
    .append("path")
    .attr("class", function(d) {
      if(data.census.has(d.id) && data.census.get(d.id).has(current.year.toString())) { 
        return map.scale(data.census.get(d.id).get(current.year.toString())[0][map.field]);
      }
    })
    .classed("na", function(d) {
      return !data.census.has(d.id) || !data.census.get(d.id).has(current.year.toString());
    })
    .classed("counties", true)
    .attr("id", function(d) { return d.id; })
    .attr("d", path)
    .on("mousemove", function(d, i) {

      var mouse = d3.mouse(d3.select("body").node());

      tooltip
      .classed("hidden", false)
      .attr("style", "left:" + (mouse[0] + 20) + "px; top:" + (mouse[1] - 40) + "px")
      .html(tooltipText(d)); 
    })
    .on("mouseout", function(d, i) {
      tooltip.classed("hidden", true);
    });

  svg.append("path")
    .datum(topojson.mesh(data.us, data.us.objects["county_" + current.year],
                          function(a, b) { 
                            return a.properties.s !== b.properties.s; 
                          }))
    .attr("class", "decade-" + current.year)
    .classed("states", true) .attr("d", path);

  svg.append("path")
    .datum(topojson.mesh(data.us, data.us.objects["county_" + current.year],
                        function(a, b) { return a === b; }))
    .attr("class", "decade-" + current.year)
    .classed("country", true) 
    .attr("d", path);

  current.map = map;
  updateLegend(map);
}

function fieldSelected() {
  field = fieldSelector.node().value;
  drawMap(current.year, maps[field]);
}

function updateLegend(map) {

  legendDate.text(current.year + " Census");
  legendField.text(map.label);

  legendColors.selectAll("circle, text").remove();

  for(i = 0; i <= 8; i++) {
    legendColors
      .append("circle")
      .attr("cx", 0)
      .attr("cy", i * 20)
      .attr("r", 7)
      .attr("class", map.color + " q" + i + "-9");

    var bounds = map.scale.invertExtent("q" + i + "-9");

    legendColors
      .append("text")
      .attr("x", 12)
      .attr("y", i * 20 + 5)
      .text(labels[i]);
  }

  legendColors.append("circle")
    .attr("cx", 0).attr("cy", 180).attr("r", 7)
    .classed("na", true);

  legendColors.append("text")
    .attr("x", 12).attr("y", 185)
    .text("Not available");

}

function tooltipText(d) {

  var male       = "n/a",
      female     = "n/a",
      malePerc   = "n/a",
      femalePerc = "n/a",
      val;

  if (data.census.has(d.id) && data.census.get(d.id).has(current.year.toString())) {

    val = data.census.get(d.id).get(current.year.toString())[0];

    male = +val.male;
    female = +val.female;
    malePerc = male / (male + female);
    femalePerc = female / (male + female);
  }

 return "<h5>" + d.properties.c + ", " + d.properties.s + "</h5>" +
   "<table>" +
   "<tr>" +
   "<td class='field'>Female: </td>" +
   "<td>" + female.toLocaleString() + "</td>" +
   "<td style='width:65px;'>" + percentageFormat(femalePerc) + "</td>" +
   "</tr><tr>"+
   "<td class='field'>Male: </td>" +
   "<td>" + male.toLocaleString() + "</td>" +
   "<td style='width:65px;'>" + percentageFormat(malePerc) + "</td>" +
   "</tr></table>";
}

