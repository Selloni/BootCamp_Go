package main

import (
	"context"
	"log"
	"web/ex01/interal/author"
	"web/ex01/interal/author/DB"
	"web/ex01/pkg/db/psql"
)

func main() {
	//handleRequest()
	cnf := psql.Config{
		Name:     "grandpat",
		Pass:     "grandpat",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgre"}
	psqlClient, err := psql.NewClient(context.TODO(), cnf)
	if err != nil {
		log.Fatal("dont connect DB", err)
	}

	newText := author.Author{
		ID:   "1",
		Text: "text test coonect",
	}

	repository := DB.NewRepository(psqlClient)

	err = repository.Create(context.TODO(), &newText)
	if err != nil {
		log.Fatal(err)
	}
}
