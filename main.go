package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gasoline/db"
	"gasoline/sum"
	"github.com/daneharrigan/perks/topk"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	port   = flag.String("p", "5000", "Web service port")
	params = []string{"i", "u", "p", "v", "r", "l", "f", "d", "o", "t"}
	event  = "event: %s\nid: %s\ndata: %s\n\n"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=gasoline ")
	flag.Parse()

	http.HandleFunc("/tracker", serveTracker)
	http.HandleFunc("/dashboard", serveDashboard)
	http.HandleFunc("/example", serveExample)
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
			case "l": // most popular url
				rec.TopK.Insert(v)
			case "f": // available features
				fs := strings.Split(v, ",")
				for _, f := range fs {
					rec.Features.Insert(f)
				}
			case "d": // resolution/screen dimension
				rec.Resolutions.Insert(v)
			case "o": // operating system
				rec.OS.Insert(v)
			case "t":
				t, err := strconv.ParseFloat(v, 64)
				if err != nil {
					log.Printf("fn=ParseFloat error=%q", err)
					continue
				}

				log.Printf("url=%s duration=%f", r.FormValue("url"), t)
				rec.ViewDuration(r.FormValue("url"), t)
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
				Features      sum.Results
				Resolutions   sum.Results
				OS            sum.Results
			}

			rec := db.Get(id)
			if rec != nil {
				data.PageView = rec.PageView
				data.Visit = rec.Visit
				data.UniqueVisitor = rec.UniqueVisitor
				data.ReturnVisitor = rec.ReturnVisitor
				data.TopView = rec.TopK.Query()
				data.Features = rec.Features.Query()
				data.Resolutions = rec.Resolutions.Query()
				data.OS = rec.OS.Query()
			}

			payload, _ := json.Marshal(&data)
			stamp := time.Now().UTC().Format(time.RFC3339)

			fmt.Fprintf(w, event, "update", stamp, payload)
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

func serveExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	log.Println("fn=serveExample")
	http.ServeFile(w, r, "assets/example.html")
}
