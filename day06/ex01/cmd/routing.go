package main

import (
	"fmt"
	//_ "github.com/lib/pq"
	"log"
	"net/http"
	"text/template"
)

// w: обращнеие к страничке r: параметр для передачи
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../ui/html/home.html")
	if err != nil {
		log.Fatal("dont read home.html")
	}
	tmpl.Execute(w, nil)
}

func adminPanel(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../ui/html/admin.html")
	if err != nil {
		log.Fatal("dont read admin.html")
	}
	tmpl.Execute(w, nil)
}

func savePost(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("pass")
	text := r.FormValue("text")
	if login == "admin" && pass == "admin" {
		//psql.OpenSql(text)
		fmt.Println(text)
	}
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/admin/", adminPanel)
	http.HandleFunc("/post/", savePost)
	http.ListenAndServe(":8888", nil)
}
