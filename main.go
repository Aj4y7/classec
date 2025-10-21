package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HTTP server for Render health checks
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "OK")
		})
		http.ListenAndServe(":8080", nil)
	}()

	StartScheduler()
}
