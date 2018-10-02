package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

func main() {
	http.HandleFunc("/", TodoIndex)
	http.ListenAndServe(":8080", nil)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}
