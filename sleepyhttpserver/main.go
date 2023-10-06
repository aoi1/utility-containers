package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 30秒間スリープしてサーバーの起動を遅延させる
	fmt.Println("Sleeping for 30 seconds before starting the server...")
	time.Sleep(30 * time.Second)

	http.HandleFunc("/", handler)
	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
