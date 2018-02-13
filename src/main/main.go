package main

import (
	"html/template"
	"log"
	"net/http"
)

//C:\Users\rsvit\Go\TranspApp\bin\Templates
//C:\Users\rsvit\Go\TranspApp\src\main

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//funcoes que recebem e direcionam para pagina...

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Several error", http.StatusInternalServerError)
		return
	}

}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Several error", http.StatusInternalServerError)
		return
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Several error", http.StatusInternalServerError)
		return
	}

}
