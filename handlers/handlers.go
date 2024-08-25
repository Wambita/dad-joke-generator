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

// error data struct
type ErrorData struct {
	StatusCode int
	Message    string
}

// error handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.WriteHeader(statusCode)

	// prepare  the error data to pass to  pass the template
	data := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}
	// parse and execute the error template
	tmplPath := filepath.Join("templates", "error.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, data)
}

// method handler
func MethodHandler(h http.HandlerFunc, allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 - Method Not Allowed")
			return
		}
		h(w, r)
	}
}

//
