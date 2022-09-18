package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func runIndex(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseGlob("templates/*.html"))
	templates.ExecuteTemplate(w, "index.html", nil)
}

func runGenerate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	prompt := r.FormValue("prompt")
	w.Write([]byte(prompt))
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", runIndex)
	http.HandleFunc("/generate", runGenerate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
