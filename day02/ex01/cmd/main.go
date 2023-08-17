package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type boolFlag struct {
	line  *bool
	char  *bool
	words *bool
}

func main() {

	var buff []string
	var flags boolFlag
	err := parsingFlag(&flags)
	if err != nil {
		log.Fatal(err)
	}
	i := 2
	if !*flags.char && !*flags.words && !*flags.line {
		i = 1
	}
	for _, file := range os.Args[i:] {
		err = countElem(file, flags, &buff)
		if err != nil {
			log.Fatal(err)
		}
	}
	//time.Sleep(3 * time.Second)
	//fmt.Println(buff)
	for _, i := range buff {
		fmt.Println(i)
	}
	//fmt.Println(buff)

}

func parsingFlag(fl *boolFlag) error {
	fl.words = flag.Bool("w", false, "counting words")
	fl.char = flag.Bool("m", false, "counting characters")
	fl.line = flag.Bool("l", false, "counting lines")
	flag.Parse()
	count := 0
	if *fl.line {
		count++
	}
	if *fl.char {
		count++
	}
	if *fl.words {
		count++
	}
	if count > 1 {
		return errors.New("you should use one flag")
	}
	return nil
}

func countElem(pathFile string, fl boolFlag, buffer *[]string) error {
	//var wg sync.WaitGroup
	//wg.Add(1)
	var mu sync.Mutex
	file, err := os.Open(pathFile)
	if err != nil {
		return errors.New("not open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if *fl.line {
		countLine(pathFile, scanner, buffer, mu)
	} else if *fl.char {
		countChar(pathFile, scanner, buffer, mu)
	} else {
		countWord(pathFile, scanner, buffer, mu)
	}
	return nil
}

func countLine(fileName string, scanner *bufio.Scanner, buffer *[]string, mu sync.Mutex) {
	var count int
	for scanner.Scan() {
		count++

	}
	mu.Lock()
	*buffer = append(*buffer, fmt.Sprintf("%d\t%s", count, fileName))
	//fmt.Println(*buffer)
	mu.Unlock()
}

func countWord(fileName string, scanner *bufio.Scanner, buffer *[]string, mu sync.Mutex) {
	var count int
	for scanner.Scan() {
		world := strings.Split(scanner.Text(), " ")
		count += len(world)

	}
	mu.Lock()
	*buffer = append(*buffer, fmt.Sprintf("%d\t%s", count, fileName))
	//fmt.Println(*buffer)
	mu.Unlock()
}

func countChar(fileName string, scanner *bufio.Scanner, buffer *[]string, mu sync.Mutex) {
	var count int
	for scanner.Scan() {
		count += len([]rune(scanner.Text()))

	}
	mu.Lock()
	*buffer = append(*buffer, fmt.Sprintf("%d\t%s", count, fileName))
	//fmt.Println(*buffer)
	mu.Unlock()
}
