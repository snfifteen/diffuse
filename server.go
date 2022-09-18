package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	_ "modernc.org/sqlite"
	"database/sql"
	"fmt"
	"time"
	"github.com/google/uuid"
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
	
	tokenCookie, err := r.Cookie("sn-token")

	if err != nil && err.Error() == "http: named cookie not present" {
		expiration  := time.Now().Add(365 * 24 * time.Hour)
        tokenCookie := http.Cookie{Name: "sn-token",Value:uuid.New().String(),Expires:expiration}
        
		http.SetCookie(w, &tokenCookie)	
	}
	
	
	fmt.Println("\nPrinting cookie with name as token")
	fmt.Println(tokenCookie)

	/*
	fmt.Println("\nPrinting all cookies")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}
	fmt.Println()
	*/

	var templates = template.Must(template.ParseGlob("templates/*.html"))
	templates.ExecuteTemplate(w, "index.html", nil)
}

func runGenerate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	prompt := r.FormValue("prompt")
	w.Write([]byte(prompt))
}



func main() {
	
	db, _ := sql.Open("sqlite", "./db/processor.db")
	rows, _ := db.Query("select * from queue")
	
	fmt.Println("rows:",rows.Next())

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", runIndex)
	http.HandleFunc("/generate", runGenerate)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
