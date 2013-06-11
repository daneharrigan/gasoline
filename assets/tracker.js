_ts1 = _ts1 || [];
(function(){
	console.log("[ts1] debug: start");

	var acc = undefined

	function Track() {
		console.log("[ts1] debug: track")

		var img = new Image()
		img.src = "/tracker.gif?" + Params()
	}

	function Params() {
		var resolution   = screen.width + "x" + screen.height,
		    screenWidth  = screen.availWidth,
		    screenHeight = screen.availHeight,
		    colorRange   = screen.pixelDepth || screen.colorDepth,
		    pixelRatio   = window.devicePixelRatio,
		    userAgent    = navigator.userAgent,
		    pageTitle 	 = document.title,
		    pathName     = window.location.pathname,
		    referrer     = document.referrer

		var params = "" +
			"&ac=" + acc +
			"&r=" + resolution +
			"&w=" + screenWidth +
			"&h=" + screenHeight +
			"&c=" + colorRange +
			"&p=" + pixelRatio +
			"&a=" + escape(userAgent) +
			"&t=" + escape(pageTitle) +
			"&u=" + escape(pathName) +
			"&rf=" + escape(referrer)

		console.log("[ts1] debug: params " + params)
		return params
	}

	function Push(args) {
		console.log("[ts1] debug: push")
		switch(args[0]) {
			case "account": acc = args[1]; break
			case "track": Track(); break
		}
	}

	_ts1.push = Push
	while(args = _ts1.shift()) {
		_ts1.push(args)
	}

	console.log("[ts1] debug: finish");
})()
