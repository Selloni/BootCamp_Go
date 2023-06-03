package main

import (
	"DBReader/DBReader"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	var j, x DBReader.Interface = &DBReader.Jake{}, &DBReader.Xake{}
	path := flag.String("f", "", "path to file with json/xml extension")
	flag.Parse() // after declaring flags we need to call it
	if strings.HasSuffix(*path, ".xml") {
		data := DBReader.OpenFile(path)
		x.ReadFile(data)
		fmt.Println(x)
		x.Create(x.Write())
	} else if strings.HasSuffix(*path, ".json") {
		data := DBReader.OpenFile(path)
		j.ReadFile(data)
		fmt.Println(j)
		j.Create(j.Write())
	} else {
		log.Fatal("no format file")
	}

}
