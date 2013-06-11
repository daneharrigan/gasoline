package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port   = flag.String("p", "5000", "Web server port")
	params = []string{
		"r",  // resolution 	`screen.width x screen.height`
		"w",  // browser width	`screen.availWidth`
		"h",  // browser height	`screen.availHeight`
		"d",  // display colors	`screen.pixelDepth || screen.colorDepth`
		"pr", // display retina	`window.devicePixelRatio > 1`
		"a",  // browser agent	`navigator.userAgent`
	}
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=ts1 ")
	flag.Parse()

	log.Println("fn=start")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", serveExample) // index
	http.HandleFunc("/tracker.gif", serveTracker)

	err := http.ListenAndServe(":"+*port, nil)
	log.Fatalf("fn=listen_and_serve error=%q", err)
}

func serveTracker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Content-Length", "0")

	for _, p := range params {
		println(p + ": " + r.FormValue(p))
	}
}

// example

func serveExample(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/example.html")
}
