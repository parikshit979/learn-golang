package handlers

import (
	"fmt"
	"net/http"
)

func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	localFileURL := getLocalFileURL("/path/to/file")

	fmt.Fprintf(w, "Local File URL: %s", localFileURL)
}

func getLocalFileURL(filePath string) string {
	return fmt.Sprintf("http://localhost:3000/files/%s", filePath)
}
