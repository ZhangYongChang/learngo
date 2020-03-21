package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	apiHandler := func(w http.ResponseWriter, req *http.Request) {
		type ColorGroup struct {
			ID     int
			Name   string
			Colors []string
		}
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		b, err := json.Marshal(group)
		if err != nil {
			log.Fatal("error:", err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}

	http.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
