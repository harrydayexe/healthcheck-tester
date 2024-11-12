package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	readyz := flag.Bool("readyz", false, "If provided, the application will return a 200 on /readyz, otherwise it will return a 503")
	livez := flag.Bool("livez", false, "If provided, the application will return a 200 on /livez, otherwise it will return a 503")
	healthz := flag.Bool("healthz", false, "If provided, the application will return a 200 on /healthz, otherwise it will return a 503")
	flag.Parse()

	// If the help flag is used, it will print usage and exit
	if len(os.Args) == 1 {
		// Check if flag.Parse() results in an error, which usually happens for invalid flags
		if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
			log.Fatal(err)
		}
	}

	router := http.NewServeMux()

	router.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !*readyz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		if !*livez {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if !*healthz {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	router.HandleFunc("/testcontainersz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3333",
		Handler: router,
	}

	fmt.Fprintf(os.Stdout, "Starting server at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
