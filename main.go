package main

import (
	"fmt"
	"net/http"

	handler "golang-api/handlers"
)

func main() {
	// http.HandleFunc("/", handler.MethodHandler(handler.Homehandler, "GET"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // http.HandleFunc("/joke", jokeHandler)

	// catch all route for unhandles paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/static/" {
			handler.ErrorHandler(w, r, http.StatusNotFound, "404 -  Not Found")
			return
		}
		handler.Homehandler(w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
