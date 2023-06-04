package main

import (
	"flag"
	"log"
	"root/ex00/DBReader"
	"strings"
)

func main() {
	var j, x DBReader.Interface = &DBReader.Jake{}, &DBReader.Xake{}
	path := flag.String("f", "", "path to file with json/xml extension")
	flag.Parse() // after declaring flags we need to call it
	if strings.HasSuffix(*path, ".xml") {
		x.ReadFile(DBReader.OpenFile(path))
		x.Create(x.Write())
	} else if strings.HasSuffix(*path, ".json") {
		j.ReadFile(DBReader.OpenFile(path))
		j.Create(j.Write())
	} else {
		log.Fatal("no format file")
	}
}
