package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type response struct {
	Headers   http.Header `json:"headers"`
	Path      string      `json:"path"`
	Queries   url.Values  `json:"queries"`
	UserAgent string      `json:"userAgent"`
}

func main() {
	host := "127.0.0.1:8000"
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Headers:   r.Header,
			Path:      r.URL.Path,
			Queries:   r.URL.Query(),
			UserAgent: r.UserAgent(),
		}

		respBytes, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(respBytes)
	})

	log.Println("running at " + host)
	log.Fatal(http.ListenAndServe(host, h))
}
