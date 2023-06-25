package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type boolFlag struct {
	symlinks    *bool
	directories *bool
	file        *bool
	extension   *bool
	extStr      *string
}

func main() {
	var flags boolFlag
	if err := parsingFlag(&flags); err != nil {
		log.Fatal(err)
	}
	myFind(os.Args[len(os.Args)-1], flags)
}

func parsingFlag(boolflg *boolFlag) error {
	boolflg.symlinks = flag.Bool("sl", false, "only symlinks")
	boolflg.directories = flag.Bool("d", false, "only directories")
	boolflg.file = flag.Bool("f", false, "only files")
	boolflg.extension = flag.Bool("ext", false, "special file extension")
	//boolflg.extStr = flag.String("ext", "", "special file extension")
	flag.Parse()
	//puck := false
	//boolflg.extension = &puck
	//println(*boolflg.extStr)
	//if *boolflg.extStr == "" {
	//	*boolflg.extension = true
	//}
	if !*boolflg.file && *boolflg.extension {
		return errors.New("you need using -f and -ext together")
	}
	return nil
}

func myFind(pathDir string, fl boolFlag) error {
	err := filepath.Walk(pathDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().Type() == os.ModeDir && *fl.directories && os.ModePerm != 0 {
			fmt.Println(path)
		} else if info.Mode().Type() == os.ModeSymlink && *fl.symlinks && os.ModePerm != 0 {
			originFile, err := os.Readlink(info.Name())
			if err != nil {
				originFile = "[broken]"
			}
			fmt.Println(path, "->", originFile)
		} else if os.ModePerm != 0 && *fl.file && *fl.extension && info.Mode().Type() != os.ModeSymlink && !info.IsDir() {
			exp := flag.String("ext", "", "special file extension")
			if strings.HasSuffix(path, "."+*exp) {
				fmt.Println(path)
			}
		} else if os.ModePerm != 0 && !info.IsDir() && info.Mode().Type() != os.ModeSymlink && *fl.file {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}
