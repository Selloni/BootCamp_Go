package interal

import (
	"errors"
	"flag"
	"os"
)

type buff struct {
}

func Parsing() {
	var old = flag.String("--old", "*.txt", "snapshot old database")
	var new = flag.String("--new", "*.txt", "snapshot new database")
	flag.Parse()
	openFile(*old, *new)
	//chekFiles(old, new)
}

func openFile(old string, new string) {
	oldFile, err := os.Open(old)
	if err != nil {
		errors.New("not old open file")
	}
}

//func chekFiles(old string, new string) {
//
//}
