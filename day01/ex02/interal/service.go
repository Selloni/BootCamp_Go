package interal

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func Parsing() {
	var old = flag.String("old", "*.txt", "snapshot old database")
	var new = flag.String("new", "*.txt", "snapshot new database")
	flag.Parse()
	CompareFile(openFile(*old), openFile(*new))
}

func openFile(pathFile string) []string {
	oldFile, err := os.Open(pathFile)
	var buffer []string
	if err != nil {
		log.Fatalf("not open %s file", pathFile)
	}
	defer oldFile.Close()
	line := bufio.NewScanner(oldFile)
	for line.Scan() {
		buffer = append(buffer, line.Text())
	}
	return buffer
}

func CompareFile(old []string, new []string) {
	sort.Strings(old)
	sort.Strings(new)
	for i := 0; i < len(old); i++ {
		find := findElem(new, old[i])
		if !find {
			fmt.Printf("REMOVED %s\n", old[i])
		}
	}
	for i := 0; i < len(new); i++ {
		find := findElem(old, new[i])
		if !find {
			fmt.Printf("ADDED %s\n", new[i])
		}
	}
}

// SearchStrings -- побаловаться, псомотреть как устроен в нутри
func findElem(buff []string, elem string) bool {
	i := sort.SearchStrings(buff, elem)
	if i < len(buff) && buff[i] == elem {
		return true
	}
	return false
}
