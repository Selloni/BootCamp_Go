package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
	"web/ex01/pkg/db/psql"
)

var flagPage uint = 1

// w: обращнеие к страничке r: параметр для передачи
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../ui/html/home.html")
	if err != nil {
		log.Fatal("dont read home.html")
	}
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
	tmpl.Execute(w, psql.PostText(psqlClient, flagPage))
}

func adminPanel(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../ui/html/admin.html")
	if err != nil {
		log.Fatal("dont read admin.html")
	}
	tmpl.Execute(w, nil)
}

func savePost(w http.ResponseWriter, r *http.Request) {
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
		err = psql.InsertText(psqlClient, text)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func showPost(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../ui/html/articl.html")
	vars := mux.Vars(r)
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
	tmpl.Execute(w, psql.OpenPagePost(psqlClient, vars))

}

func backPage(w http.ResponseWriter, r *http.Request) {
	//flagPage = flagPage - 2
	flagPage--
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func forwardPage(w http.ResponseWriter, r *http.Request) {
	//flagPage = flagPage + 2
	flagPage++
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRequest() {
	rout := mux.NewRouter()
	rout.HandleFunc("/", homePage).Methods("GET")
	rout.HandleFunc("/admin/", adminPanel).Methods("GET")
	rout.HandleFunc("/post/", savePost).Methods("POST")
	rout.HandleFunc("/post/{id:[0-9]+}", showPost).Methods("GET")
	rout.HandleFunc("/back/", backPage).Methods("POST", "GET")
	rout.HandleFunc("/forward/", forwardPage).Methods("POST", "GET")

	http.Handle("/", rout)
	http.ListenAndServe(":8888", nil)
}
