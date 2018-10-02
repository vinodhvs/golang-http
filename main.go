package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("{yolo:%v}", 12)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
