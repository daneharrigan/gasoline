function debug(message) {
	console.debug("app=gasoline " + message)
};

_gasoline = _gasoline || [];

(function(){
	debug("fn=start")
	var accountId,
	    image = new Image,
	    w = window,
	    d = document,
	    n = navigator

	function setCookie(name, value, lifespan) {
		var cookie = name + "=" + escape(value)

		if(lifespan) {
			cookie += "; expires=" + lifespan
		}

		cookie += "; path=/"
		document.cookie = cookie
	}

	function getCookie(name) {
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

	function existingCookie(name, days) {
		var v = getCookie(name),
		    d = new Date

		d.setDate(d.getDate() + days)
		setCookie(name, 1, d.toGMTString())

		return !!v
	}

	function pageView() {
		return true
	}

	function visit() {
		return d.referrer.indexOf(d.location.origin) != 0
	}

	function returnVisitor() {
		return !existingCookie("_gasoline_visitor", 1)
	}

	function uniqueVisitor() {
		return !existingCookie("_gasoline_unique", 365)
	}

	function features() {
		var value = []

		// Cookies
		if(n.cookieEnabled) {
			value.push("Cookies")
		}

		// Retina Display
		if(w.devicePixelRatio > 1) {
			value.push(escape("Retina Display"))
		}

		// Silverlight, Flash, QuickTime, Google Talk, Java Applet
		var p = n.plugins
		for(var i=0; i<p.length; i++) {
			var name = escape(p[i].name.split(/ plug-?in/i)[0])
			if(value.indexOf(name) == -1) {
				value.push(name)
			}
		}

		return value.join(",")
	}

	function track() {
		debug("fn=track")
		var payload = "?"
		var params = {
			i: accoutId,
			u: uniqueVisitor,
			p: pageView,
			v: visit,
			r: returnVisitor,
			l: escape(w.location.pathname),
			f: features()
		}

		for(var k in params) {
			var v = params[k]
			if(typeof v == "function") {
				if(v()) {
					payload += "&" + k + "=1"
				}
      } else {
				payload += "&" + k + "=" + v
			}
		}

		image.src = "/tracker" + payload
	}

	_gasoline.push = function(value) {
		switch(value[0]) {
		case "account": accoutId = value[1]; break
		case "track": track()
		}
	}

	debug("fn=shift")
	while(args = _gasoline.shift()) {
		_gasoline.push(args)
	}
})();
