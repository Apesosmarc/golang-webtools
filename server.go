package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse map[string]int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		meth := r.Method

		fmt.Printf("Method: %v\n", meth)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse{"apple": 5, "lettuce": 19, "testData": 20009912})
	})
	port := ":3000"
	fmt.Printf("server starting on: %v\n", port)
	http.ListenAndServe(":3000", nil)
}
