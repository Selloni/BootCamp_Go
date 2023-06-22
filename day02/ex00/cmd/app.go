package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
)

type boolFlag struct {
	symlinks    *bool
	directories *bool
	file        *bool
	extension   *bool
}

func main() {
	var flags boolFlag
	if err := parsingFlag(&flags); err != nil {
		log.Fatal(err)
	}
	myFind(os.Args[len(os.Args)-1])
}

func parsingFlag(boolflg *boolFlag) error {
	boolflg.symlinks = flag.Bool("sl", false, "only symlinks")
	boolflg.directories = flag.Bool("d", false, "only directories")
	boolflg.file = flag.Bool("f", false, "only files")
	boolflg.extension = flag.Bool("ext", false, "extension for file")
	flag.Parse()
	if !*boolflg.file && *boolflg.extension {
		return errors.New("you need using -f and -ext together")
	}
	return nil
}

func myFind(pathDir string) error {
	cout := exec.Command("find", pathDir)
	cout.Stdout = os.Stdout
	cout.Run()
	return nil
}
