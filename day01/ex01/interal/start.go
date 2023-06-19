package interal

import (
	"errors"
	"flag"
	"fmt"
	"root/ex01/JX"
	"strings"
)

func Start() error {

	pthOrigin := flag.String("old", "", "original recipe")
	pthStolen := flag.String("new", "", "stolen recipe")
	flag.Parse() // after declaring flags we need to call it
	fmt.Println(*pthOrigin)
	fmt.Println(*pthStolen)
	if !strings.HasSuffix(*pthOrigin, ".json") && !strings.HasSuffix(*pthStolen, ".xml") {
		return errors.New("no format file")
	}
	jake := JX.Jake{}
	xake := JX.Xake{}
	var (
		cakeJ = recipes{}
		cakeX = recipes{}
	)

	ParsingJson(ReadFile(pthOrigin), &jake)
	ParsingXml(ReadFile(pthStolen), &xake)
	cakeJ.TakeJson(jake)
	cakeX.TakeXml(xake)
	fmt.Println("test")
	fmt.Println(cakeJ)
	fmt.Println(cakeX)
	return nil
}
