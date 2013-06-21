var _ts1 = _ts1 || [];
(function(){

  var accoutId,
      image = new Image(),
      w = window,
      n = navigator,
      s = screen

  function debug(message) {
    console.debug("[debug]: " + message)
  }

  function resolution() {
    return s.width + "x" + s.height
  }

  function dimension() {
    return s.availWidth + "x" + s.availHeight
  }

  function retina() {
    return (w.devicePixelRatio > 1)
  }

  function color() {
    return (s.pixelDepth || s.colorDepth)
  }

  function agent() {
    return n.userAgent
  }

  function os() {
    return n.platform
  }

  function cookie() {
    return n.cookieEnabled
  }

  function plugins() {
    var value = ""
    for(var i=0; i < n.plugins.length; i++) {
      debug("plugin[" + i + "]="+n.plugins[i].name)
      value += "&plugin[]=" + escape(n.plugins[i].name)
    }

    return value
  }

  function referrer() {
    return document.referrer
  }

  function title() {
	  return document.title
  }

  function url() {
  	return w.location.pathname
  }

  function lifespan() {
    return (new Date).getTime() - _ts1.start
  }

  function track(raw, escaped) {
    var params = "?"
    for(var key in raw) {
      debug(key + "=" + raw[key])
      params += "&" + key + "=" + escape(raw[key])
    }

    if(escaped) {
      params += escaped
    }

    debug(params)
    _ts1.start = (new Date).getTime()
    image.src = "/tracker.gif" + params
  }

  function push(value) {
    debug("push")
    switch(value[0]) {
    case "account": accoutId = args[1]; break
    case "track": track({ title: title(), url: url(), lifespan: lifespan() })
    }
  }

  track({
    resolution: resolution(),
    dimension: dimension(),
    retina: retina(),
    color: color(),
    agent: agent(),
    os: os(),
    cookie: cookie(),
    referrer: referrer(),
    title: title(),
    url: url()
  }, plugins())

  _ts1.push = push
})()
