package main

import (
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(100) % 5 {
	case 1:
		w.WriteHeader(http.StatusOK)
	case 2:
		w.WriteHeader(http.StatusFound)
	case 3:
		w.WriteHeader(http.StatusBadRequest)
	case 4:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
