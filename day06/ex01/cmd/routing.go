package main

import (
	"context"
	"fmt"
	"web/ex01/pkg/db/psql"

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
	///////////
	cnf := psql.Config{
		Name:     "grandpat",
		Pass:     "grandpat",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
	}
	psqlClient, err := psql.NewClient(context.Background(), cnf)
	if err != nil {
		log.Fatal("failed to connect to the database", err)
	}
	defer psqlClient.Close()
	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	login := r.Form.Get("login")
	pass := r.Form.Get("pass")
	text := r.Form.Get("text")
	if login == "admin" && pass == "admin" {
		fmt.Printf("ll %s, %s, %s", login, pass, text)
		err = psql.InsertText(psqlClient, text)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/admin/", adminPanel)
	http.HandleFunc("/post/", savePost)

	http.ListenAndServe(":8888", nil)
}
