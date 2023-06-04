package interal

import (
	"errors"
	"flag"
	"root/ex00/DBReader"
	"strings"
)

func app() error {
	pthOrigin := flag.String("old", "", "original recipe")
	pthStolen := flag.String("new", "", "stolen recipe")
	jo := DBReader.Jake{}
	js := DBReader.Jake{}
	xo := DBReader.Jake{}
	xs := DBReader.Jake{}
	if strings.HasSuffix(*pthOrigin, ".json") {
		jo.ReadFile(DBReader.OpenFile(pthOrigin))
		println(jo)
	} else if strings.HasSuffix(*pthOrigin, ".xml") {
		xo.ReadFile(DBReader.OpenFile(pthOrigin))
	} else {
		return errors.New("no format file")
	}
	if strings.HasSuffix(*pthStolen, ".xml") {
		js.ReadFile(DBReader.OpenFile(pthOrigin))
		println(jo)
	} else if strings.HasSuffix(*pthStolen, ".json") {
		xs.ReadFile(DBReader.OpenFile(pthOrigin))
	} else {
		return errors.New("no format file")
	}

	return nil
}
