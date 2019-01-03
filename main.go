package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}

		if req.URL.Path == "/__sleep__" {
			durationStr := req.URL.Query().Get("duration")
			duration, err := time.ParseDuration(durationStr)
			if err != nil {
				duration = time.Second
			}

			time.Sleep(duration)
			resp.WriteHeader(200)
			resp.Write([]byte("slept for " + duration.String()))
			return
		}

		http.FileServer(http.Dir(*dir)).ServeHTTP(resp, req)
	})))
}
