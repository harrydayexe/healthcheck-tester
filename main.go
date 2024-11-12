package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	readyz := flag.Bool("r", true, "Should the container return a 200 on /readyz")
	livez := flag.Bool("l", true, "Should the container return a 200 on /livez")
	healthz := flag.Bool("h", true, "Should the container return a 200 on /healthz")

	router := http.NewServeMux()

	router.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if *readyz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		if *livez {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if *healthz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/testcontainersz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
