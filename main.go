package main

import (
	"flag"
	"fmt"
	"gasoline/db"
	"log"
	"net/http"
)

var (
	port   = flag.String("p", "5000", "Web service port")
	params = []string{"i", "u", "p", "v", "r"}
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=gasoline ")
	flag.Parse()

	http.HandleFunc("/tracker", serveTracker)
	http.HandleFunc("/dashboard", serveDashboard)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("fn=ListenAndServe")
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("fn=ListenAndServe error=%q", err)
	}
}

func serveTracker(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "Method Not Allowed", 405)
	case "OPTIONS":
		w.Header().Set("Content-Length", "0")
		w.Header().Set("Allow", "OPTIONS, GET")
	case "GET":
		w.Header().Set("Content-Length", "0")
		w.Header().Set("Content-Type", "image/gif")

		id := r.FormValue("i")
		if id == "" {
			http.Error(w, "Forbidden", 403)
			return
		}

		rec := db.Get(id)
		if rec == nil {
			rec = db.New(id)
		}

		log.Printf("page=tracker id=%s", id)

		for _, k := range params {
			v := r.FormValue(k)
			if v == "" {
				continue
			}

			switch k {
			case "p": // page view
				rec.PageView++
			case "u": // unique visitor
				rec.UniqueVisitor++
			case "v": // visits
				rec.Visit++
			case "r": // return visitor
				rec.ReturnVisitor++
			}
		}
	}
}

func serveDashboard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "Method Not Allowed", 405)
	case "OPTIONS":
		w.Header().Set("Content-Length", "0")
		w.Header().Set("Allow", "OPTIONS, GET")
	case "GET":
		w.Header().Set("Content-Type", "text/plain")

		id := r.FormValue("i")
		if id == "" {
			http.Error(w, "Forbidden", 403)
			return
		}

		log.Printf("page=dashboard id=%s", id)

		rec := db.Get(id)
		if rec == nil {
			http.Error(w, "Unprocessable Entity", 422)
			return
		}

		fmt.Fprintf(w, "%s: %d\n", "PageView", rec.PageView)
		fmt.Fprintf(w, "%s: %d\n", "Visit", rec.Visit)
		fmt.Fprintf(w, "%s: %d\n", "UniqueVisitor", rec.UniqueVisitor)
		fmt.Fprintf(w, "%s: %d\n", "ReturnVisitor", rec.ReturnVisitor)
	}
}
