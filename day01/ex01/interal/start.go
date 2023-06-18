package interal

//
//func Start() error {
//	jo := DBReader.Jake{}
//	js := DBReader.Jake{}
//	//xo := DBReader.Xake{}
//	//xs := DBReader.Xake{}
//	pthOrigin := flag.String("old", "", "original recipe")
//	pthStolen := flag.String("new", "", "stolen recipe")
//	flag.Parse() // after declaring flags we need to call it
//	fmt.Println(*pthOrigin)
//	fmt.Println(*pthStolen)
//	if strings.HasSuffix(*pthOrigin, ".json") {
//		jo.ReadFile(OpenFile(pthOrigin))
//		//WriteJson(&jo)
//	} else if strings.HasSuffix(*pthOrigin, ".xml") {
//		jo.ReadFile(OpenFile(pthOrigin))
//		//WriteJson(&jo)
//	} else {
//		return errors.New("no format file")
//	}
//	if strings.HasSuffix(*pthStolen, ".xml") {
//		js.ReadFile(OpenFile(pthStolen))
//		//WriteJson(&js)
//	} else if strings.HasSuffix(*pthStolen, ".json") {
//		js.ReadFile(OpenFile(pthStolen))
//		//WriteJson(&js)
//	} else {
//		return errors.New("no format file")
//	}
//	fmt.Println(jo)
//	fmt.Println(js)
//	Comparison(&jo, &js)
//
//	return nil
//}
