package main

import (
	"DBReader/DBReader"
	"flag"
)

func main() {
	//var a DBReader.Interface = &DBReader.Jake{}
	path := flag.String("f", "", "path to file with json/xml extension")
	flag.Parse() // after declaring flags we need to call it
	jdata := DBReader.OpenFile(path)
	DBReader.ReadFile(jdata)
	//a.ReadFile(jdata)
	//xdata := DBReader.OpenFile("../dock_for_parsing/original.xml")
	//var j, x Interface = &Jake{}, &Xake{}
	//DBReader.ReadFile(jdata)
}
