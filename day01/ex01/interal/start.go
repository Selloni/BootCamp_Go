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
		old = recipes{}
		new = recipes{}
	)

	ParsingJson(ReadFile(pthOrigin), &jake)
	ParsingXml(ReadFile(pthStolen), &xake)
	old.TakeJson(jake)
	new.TakeXml(xake)
	CakeAdd(old, new)
	CakeRemove(old, new)
	ChangeTime(old, new)
	return nil
}
