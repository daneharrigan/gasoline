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
			"&rf=" + escape(referrer) +
			"&rv=" + Revisit()

		console.log("[ts1] debug: params " + params)
		return params
	}

	function Revisit() {
		var value = GetCookie("_ts1_revisit")
		SetCookie("_ts1_revisit", 1, 0) //, 60 * 60 * 24 * 356) // should i always set you
		return !!value
	}

	function SetCookie(name, value, lifeSpan) {
		var cookie = name + "=" + escape(value)

		if(lifeSpan > 0) {
			var timestamp = new Date()
			timestamp.setTime(timestamp.getTime() + (lifeSpan * 1000))
			cookie += "; expires=" + timestamp.toGMTString()
		}

		cookie += "; path=/"
		document.cookie = cookie
	}

	function GetCookie(name) {
		name += "="
		var cookies = document.cookie.split(";")
		for(var i=0; i < cookies.length; i++) {
			var cookie = cookies[i]
			if(cookie[0] == " ") {
				cookie = cookie.substring(1, cookie.length)
			}

			if(cookie.indexOf(name) == 0) {
				return unescape(cookie.substring(name.length, cookie.length))
			}
		}

		return
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
