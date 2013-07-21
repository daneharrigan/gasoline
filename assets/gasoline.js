// temporary: start
function range(min, max) {
  return Math.random() * (max - min) + min;
}

// temporary: end

Gasoline = {}
Gasoline.Graph = function(p, options) {
  var data   = options.data,
      width  = options.width,
      height = options.height,
      margin = options.margin || 0

  var x = d3.time.scale().range([0, width]),
      y = d3.scale.linear().range([height + margin, margin]),
      stack = d3.layout.stack().offset("wiggle")

  var line = d3.svg.line()
    .x(function(d) { return x(d.x) })
    .y(function(d) { return y(d.y) })
    .interpolate("monotone")

  var svg = d3.select(p).append("svg")
    .attr("width", width)
    .attr("height", height + margin)

  var axis = d3.svg.axis()
    .scale(x)
    .orient("bottom")
    .ticks(10)
    .tickSize(-height, 0, 0)

  x.domain(d3.extent(data, function(d) { return d.x }))
  y.domain([0, 20])

  svg.append("path")
    .datum(data)
    .attr("class", "line total")
    .attr("d", line)

  svg.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + height + ")")
    .call(axis)

  if(options.live) {
    function animate() {
      /* if a chrome tab is not in focus the rendering stacks and
       * the graph looks very wrong. remove any item that is older
       * than 60 seconds
       */
      while(data.length > 60) {
        data.shift()
      }

      x.domain(d3.extent(data, function(d) { return d.x }))

      d3.select("path").transition()
        .duration(1000)
        .attr("d", line)

      d3.select("g").transition()
        .duration(1000)
	.call(axis)

      data.push({
        "x": (new Date).setMilliseconds(0),
        "y": range(10, 12)
      })

      setTimeout(animate, 1000)
    }

    animate()
  }
}

