Gasoline = {}
Gasoline.DB = {}
Gasoline.Graph = function(p, options) {
	var
	data   = options.data,
	width  = options.width,
	height = options.height,
	margin = options.margin || 0,
	parse = d3.time.format.iso.parse

	var
	x = d3.time.scale().range([0, width]),
	y = d3.scale.linear().range([height + margin, margin]),
	stack = d3.layout.stack().offset("wiggle")

	var
	line = d3.svg.line()
		.x(function(d) { return x(parse(d.Timestamp)) })
		.y(function(d) { return y(d.Value) })
		.interpolate("monotone")

	var
	svg = d3.select(p).append("svg")
		.attr("width", width)
		.attr("height", height + margin)

	var
	axis = d3.svg.axis()
		.scale(x)
		.orient("bottom")
		.ticks(10)
		.tickSize(-height, 0, 0)

	function domains() {
		var
		dxMin,
		dxMax,
		dyMin = 0,
		dyMax = 0

		for(var i=0; i<data.length; i++) {
			var
			n = data[i].Value.length - 1,
			xMin = parse(data[i].Value[n].Timestamp),
			xMax = parse(data[i].Value[0].Timestamp),
			yMax = data[i].Value[0].Value

			if(!dxMin || !dxMax) {
				dxMin = xMin
				dxMax = xMax
			}

			if(xMin < dxMin) {
				dxMin = xMin
			}

			if(xMax > dxMax) {
				dxMax = xMax
			}

			if(yMax > dyMax) {
				dyMax = yMax
			}
		}

		x.domain([dxMax, dxMin])
		y.domain([dyMin, dyMax])
	}

	// add lines
	domains()
	svg.selectAll(".line")
		.data(data)
		.enter().append("path")
		.attr("class", function(d) { return d.Name + " line" })
		.attr("d", function(d) { return line(d.Value) })

	svg.append("g")
		.attr("class", "x axis")
		.attr("transform", "translate(0," + height + ")")
		.call(axis)

	// animate
	function update() {
		/* if a chrome tab is not in focus the rendering stacks and
		 * the graph looks very wrong. remove any item that is older
		 * than 60 seconds
		 */

		for(var i=0; i<data.length; i++) {
			while(data[i].Value.length > 60) {
				data[i].Value.shift()
			}
		}

		domains()
		d3.selectAll(".line").transition()
			.duration(1000)
			.attr("d", function(d) { return line(d.Value) })

		d3.select("g").transition()
			.duration(1000)
			.call(axis)

		setTimeout(update, 1000)
	}

	if(options.update) {
		update()
	}
}

Gasoline.Value = function(p, options) {
	var el = document.querySelector(p)

	function update() {
		if(!el.innerText || options.value != el.innerText) {
			el.innerText = options.value
		}

		if(options.update) {
			setTimeout(update, 1000)
		}
	}

	update()
}

Gasoline.format = {
	number: function(n) {
		var
		s = n.toString(),
		v = "",
		i = s.length

		while(i > 3) {
			var x = i - 3
			v = s.substring(x, i) + v
			i = x

			if(i > 0) {
				v = "," + v
			}
		}

		if(i > 0) {
			v = s.substring(0, i) + v
		}

		return v
	}
}
