package main

import (
	"fmt"
	"net/http"

	handler "golang-api/handlers"
)

func main() {
	http.HandleFunc("/", handler.Homehandler)
	// http.HandleFunc("/joke", jokeHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
