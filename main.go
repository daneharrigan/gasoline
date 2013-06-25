package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gasoline/db"
	"github.com/daneharrigan/perks/topk"
	"log"
	"net/http"
	"time"
)

var (
	port   = flag.String("p", "5000", "Web service port")
	params = []string{"i", "u", "p", "v", "r", "l"}
	event  = "event: %s\n\ndata: %s\n\nid: %s\n\n"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=gasoline ")
	flag.Parse()

	http.HandleFunc("/tracker", serveTracker)
	http.HandleFunc("/dashboard", serveDashboard)
	http.HandleFunc("/stream", serveStream)
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

		log.Printf("fn=serveTracker id=%s", id)

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
			case "l":
				rec.TopK.Insert(v)
			}
		}
	}
}

func serveStream(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "Method Not Allowed", 405)
	case "OPTIONS":
		w.Header().Set("Content-Length", "0")
		w.Header().Set("Allow", "OPTIONS, GET")
	case "GET":
		id := r.FormValue("i")
		if id == "" {
			http.Error(w, "Forbidden", 403)
			return
		}

		log.Printf("fn=serveStream id=%s", id)
		w.Header().Set("Content-Type", "text/event-stream")
		f, _ := w.(http.Flusher)

		for {
			var data struct {
				PageView      int64
				Visit         int64
				UniqueVisitor int64
				ReturnVisitor int64
				TopView       topk.Samples
			}

			rec := db.Get(id)
			if rec != nil {
				data.PageView = rec.PageView
				data.Visit = rec.Visit
				data.UniqueVisitor = rec.UniqueVisitor
				data.ReturnVisitor = rec.ReturnVisitor
				data.TopView = rec.TopK.Query()
			}

			payload, _ := json.Marshal(&data)
			stamp := time.Now().UTC().Format(time.RFC3339)

			fmt.Fprintf(w, event, "update", payload, stamp)
			f.Flush()

			time.Sleep(1 * time.Second)
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
		w.Header().Set("Content-Type", "text/html")
		log.Println("fn=serveDashboard")
		http.ServeFile(w, r, "assets/dashboard.html")
	}
}
