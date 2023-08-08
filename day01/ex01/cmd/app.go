package main

import (
	"log"
	"root/ex01/interal"
)

func main() {
	err := interal.Start()
	if err != nil {
		log.Fatal(err)
	}

}
