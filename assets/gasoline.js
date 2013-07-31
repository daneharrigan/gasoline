Gasoline = {}
Gasoline.DB = {}
Gasoline.Graph = function(q, options) {
	var parse = d3.time.format.iso.parse

	function domains() {
		var
		dxMin,
		dxMax,
		dyMin = 0,
		dyMax = 0

		for(var i=0; i<data.length; i++) {
			var
			n = data[i].Values.length - 1,
			xMin = parse(data[i].Values[n].Timestamp),
			xMax = parse(data[i].Values[0].Timestamp),
			yMax = data[i].Values[0].Count

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

	function update() {
		console.log("at=update fn=Graph")
		for(var i=0; i<data.length; i++) {
			while(data[i].Values.length > 60) {
				data[i].Values.shift()
			}
		}

		domains()
		d3.selectAll("path.line")
			.data(data)
			.transition()
			.duration(1000)
			.attr("d", function(d) { return line(d.Values) })

		d3.select("g.axis").transition()
			.duration(1000)
			.call(axis)

		setTimeout(update, 1000)
	}

	var
	width = options.width,
	height = options.height,
	data = options.data,
	xMax = width,
	yMax = height - 45,
	lIndent = 4,
	lHeight = height - 3,
	tHeight = -(height - 40),
	aHeight = height - 40

	var
	svg = d3.select(q).append("svg")
		.attr("width", width)
		.attr("height", height)
		.attr("class", "lines")

	var
	x = d3.time.scale().range([0, xMax]),
	y = d3.scale.linear().range([yMax, 0])

	var
	line = d3.svg.line()
		.x(function(d) { return x(parse(d.Timestamp)) })
		.y(function(d) { return y(d.Count) })
		.interpolate("monotone")

	var
	lines = svg.append("g")
		.attr("class", "lines")

	var
	legend = svg.selectAll(".legend")
		.data(data)
		.enter().append("g")
			.attr("class", "legend")
			.attr("transform", "translate(" + lIndent + "," + lHeight + ")")

	var
	axis = d3.svg.axis()
		.scale(x)
		.orient("bottom")
		.ticks(10)
		.tickSize(tHeight, 0, 0)

	domains()
	lines.selectAll("path.line")
		.data(data)
	  .enter().append("path")
		.attr("class", function(d) { return "line " + d.Class })
		.attr("d", function(d) { return line(d.Values) })

	legend.append("circle")
		.attr("class", function(d) { return d.Class })
		.attr("r", 4)
		.attr("cy", -4)
		.attr("cx", function(_, i) { return i * 100 })

	legend.append("text")
		.text(function(d) { return d.Name })
		.attr("x", function(_, i) { return 10 + (i * 100) })

	svg.append("g")
		.attr("class", "axis")
		.attr("transform", "translate(0," + aHeight + ")")
		.call(axis)

	if(options.update) {
		update()
	}
}

Gasoline.Value = function(p, options) {
	var el = document.querySelector(p)

	function update() {
		var v = options.value
		if(typeof options.value == "function") {
			v = v()
		}
		console.log("at=update fn=Value")
		if(!el.innerText || v != el.innerText) {
			el.innerText = v
		}

		if(options.update) {
			setTimeout(update, 1000)
		}
	}

	update()
}

Gasoline.fmt = {
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
