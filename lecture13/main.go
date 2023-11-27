package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/static/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	filePath := "./lecture13/internal/files/" + strings.TrimPrefix(r.URL.Path, "/static/")
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("error openning file: %v", err), http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("getting file info error: %v", err), http.StatusNoContent)
		return
	}

	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
