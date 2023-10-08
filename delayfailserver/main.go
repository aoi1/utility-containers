package main

import (
	"log"
	"net/http"
	"time"
)

var startTime time.Time

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if time.Since(startTime) < 10*time.Second {
		log.Println("Health Check: OK")
		w.Write([]byte("OK"))
	} else {
		log.Println("Error: Service Unhealthy")
		// Respond with an error status after 10 seconds have passed
		http.Error(w, "Service Unhealthy", http.StatusServiceUnavailable)
	}
}

func main() {
	// Record the start time of the server
	startTime = time.Now()

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthCheckHandler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
