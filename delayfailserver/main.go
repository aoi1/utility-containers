package main

import (
	"net/http"
	"sync"
	"time"
  "fmt"
)

var (
	// Flag to determine whether to send an error or not
	sendError bool
	// Mutex to safely update and check the sendError flag
	mu sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex and defer its unlock until the function returns
	mu.Lock()
	defer mu.Unlock()

	// Check if the server should send an error
	if sendError {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Otherwise, respond normally
	w.Write([]byte("Hello, world!"))
}

func main() {
	// Set a timer to trigger 5 seconds after startup
	go func() {
		time.Sleep(10 * time.Second)
		mu.Lock()
		defer mu.Unlock()
		sendError = true
	}()

	// Set up the HTTP server
	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

