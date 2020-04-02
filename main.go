package main

import (
	"html/template"
	"net/http"
)

type GData struct {
	Title string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("signup.html"))
}

func main() {

	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)

}

func signup(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "SIGNUP",
	}
	tpl.ExecuteTemplate(w, "signup.html", gd)
}