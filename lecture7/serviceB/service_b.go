package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/to-a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "From B to A")
	})
	http.HandleFunc("/from-a", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://localhost:5000/to-b")
		if err != nil {
			http.Error(w, "error contacting with service a", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "error reading response body from a", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	})

	http.ListenAndServe(":5001", nil)
}
