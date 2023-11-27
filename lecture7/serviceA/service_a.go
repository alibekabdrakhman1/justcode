package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/to-b", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "From A to B")
	})
	http.HandleFunc("/from-b", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://localhost:5001/to-a")
		if err != nil {
			http.Error(w, "error contacting with service b", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "error reading response body from b", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	})

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
}
