package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Millisecond)
	fmt.Fprint(w, "Hello World from Go.")
}
