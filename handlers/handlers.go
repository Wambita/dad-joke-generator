package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	logic "golang-api/api"
)

// struct to hold data passed to the template
type PageData struct {
	Joke string
}

// handlers
func Homehandler(w http.ResponseWriter, r *http.Request) {
	joke, err := logic.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Joke: joke,
	}

	tmplPath := filepath.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, data)
}

// // api handler
// func jokeHandler(w http.ResponseWriter, r *http.Request) {
// 	joke, err := logic.GetJoke()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Fprint(w, joke)
// }
