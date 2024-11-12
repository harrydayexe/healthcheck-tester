package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get Environment Variables
	readyzString := os.Getenv("READYZ")
	livezString := os.Getenv("LIVEZ")
	healthzString := os.Getenv("HEALTHZ")

	var readyz, livez, healthz bool = true, true, true
	if readyzString == "FALSE" {
		readyz = false
	}
	if livezString == "FALSE" {
		livez = false
	}
	if healthzString == "FALSE" {
		healthz = false
	}

	router := http.NewServeMux()

	router.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if readyz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		if livez {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if healthz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/testcontainersz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
