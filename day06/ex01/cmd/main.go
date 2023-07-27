package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Login    string
	Password string
}

// w: обращнеие к страничке r: параметр для передачи
func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hi world")
	tmpl, err := template.ParseFiles("../ui/html/home.html")
	if err != nil {
		log.Fatal("dont read .html")
	}
	tmpl.Execute(w, nil)
}

func adminPanel(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hi world")
	tmpl, err := template.ParseFiles("../ui/html/admin.html")
	if err != nil {
		log.Fatal("dont read .html")
	}
	tmpl.Execute(w, nil)
}

func savePost(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("pass")
	text := r.FormValue("text")
	if login == "admin" && pass == "admin" {
		fmt.Fprintf(w, text)
	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/admin/", adminPanel)
	http.HandleFunc("/post/", savePost)
	http.ListenAndServe(":8888", nil)
}

func main() {
	handleRequest()
}
