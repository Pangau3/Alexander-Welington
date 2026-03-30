package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", UploadHandler)

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
