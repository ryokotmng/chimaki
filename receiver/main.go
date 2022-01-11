package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Printer)
	http.ListenAndServe(":8080", nil)
}

func Printer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received -----")
}
