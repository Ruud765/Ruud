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
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/signup", signup)
}

func signup(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "SIGNUP",
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", gd)
}