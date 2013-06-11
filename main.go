package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type View struct {
	Title string
	URL   string
	Count int
}

type DB struct {
	Request    int
	Retina     int
	Resolution map[string]int
	PageViews  []*View
}

var (
	db     = &DB{Resolution: make(map[string]int), PageViews: []*View{}}
	port   = flag.String("p", "5000", "Web server port")
	params = []string{
		"r",  // resolution 	`screen.width x screen.height`
		"w",  // browser width	`screen.availWidth`
		"h",  // browser height	`screen.availHeight`
		"c",  // display colors	`screen.pixelDepth || screen.colorDepth`
		"p",  // display retina	`window.devicePixelRatio > 1`
		"a",  // browser agent	`navigator.userAgent`
		"t",  // page title     `document.title`
		"u",  // page uri	`window.location.pathname`
		"rf", // referrer	`document.referrer`
	}
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=ts1 ")
	flag.Parse()

	log.Println("fn=start")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/tracker.gif", serveTracker)
	http.HandleFunc("/", serveExample)
	http.HandleFunc("/dashboard", serveDashboard)
	http.HandleFunc("/events", serveEvents)

	err := http.ListenAndServe(":"+*port, nil)
	log.Fatalf("fn=listen_and_serve error=%q", err)
}

func serveTracker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Content-Length", "0")

	for _, p := range params {
		value := r.FormValue(p)
		switch p {
		case "r":
			db.Resolution[value]++
		case "p":
			db.Retina++
		}
	}

	db.Request++

	title := r.FormValue("t")
	uri := r.FormValue("u")

	for _, view := range db.PageViews {
		if view.URL == uri {
			view.Count++
			return
		}
	}

	db.PageViews = append(db.PageViews, &View{
		URL: uri, Title: title, Count: 1})
}

// example

func serveExample(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/example.html")
}

// dashboard

func serveDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/dashboard.html")
}

// events

func serveEvents(w http.ResponseWriter, r *http.Request) {
	f, _ := w.(http.Flusher)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var event string
	event += "event: %s\n\n"
	event += "data: %s\n\n"

	for _ = range time.Tick(500 * time.Millisecond) {
		payload, _ := json.Marshal(&db)
		fmt.Fprintf(w, event, "message", payload)
		f.Flush()
	}
}
