package main

import (
	"flag"
	"log"
	"net/http"
	"gasoline/db"
)

var (
	port   = flag.String("p", "5000", "Web service port")
	params = []string{"i", "u", "pv", "v", "vr", "rv"}
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=gasoline ")
	flag.Parse()

	http.HandleFunc("/tracker", serveTracker)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("fn=ListenAndServe")
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("fn=ListenAndServe error=%q", err)
	}
}

func serveTracker(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("i")
	if id == "" {
		http.Error(w, "Forbidden", 403)
		return
	}

	rec := db.Get(id)
	if rec == nil {
		rec = db.New(id)
	}

	for _, k := range params {
		v := r.FormValue(k)
		if v == "" {
			continue
		}

		switch k {
		case "pv": // page view
			rec.Pageview++
		case "vr": // visitor
			rec.Visitor++
		case "v": // visit
			rec.Visit++
		case "r": // return visitor
			rec.ReturnVisitor++
		}
	}
}
