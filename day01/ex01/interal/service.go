package interal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"root/ex00/DBReader"
)

func OpenFile(str *string) []byte {
	data, err := os.ReadFile(*str)
	if err != nil {
		errors.New("not open")
	}
	return data
}

func WriteJson(bd interface{}) error {
	_, err := json.Marshal(bd)
	if err != nil {
		errors.New("the file is not written")
	}
	return nil
}

func Comparison(jo *DBReader.Jake, js *DBReader.Jake) {
	println("sdfs")
	for i := range jo.Cake {
		for j := range js.Cake {
			if jo.Cake[i].Name == js.Cake[j].Name {
				println("eqal")
				compIng(jo, js, i)
			}
		}
	}
}

func compIng(origin *DBReader.Jake, stolen *DBReader.Jake, count int) {
	find := false
	for i := range origin.Cake[count].Ingredients {
		for j := range stolen.Cake[count].Ingredients {
			org := origin.Cake[count].Ingredients[i]
			stl := stolen.Cake[count].Ingredients[j]
			if org.IngredientName == stl.IngredientName {
				find = true
			}
			if find {
				if org.IngredientCount != org.IngredientCount {
					fmt.Printf("CHANGED unit count for ingredient %s for cake %s - %d instead of %d", org, origin.Cake[count], org.IngredientCount, stl.IngredientCount)
				}
			} else {
				fmt.Printf("REMOVED ingredient \\%s\\ for cake \\%s\\", org, origin.Cake[count])
			}
		}
	}
}

//./compareDB --old original_database.xml --new stolen_database.json
//ADDED cake \"Moonshine Muffin\"
//REMOVED cake \"Blueberry Muffin Cake\"
//CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"
//ADDED ingredient \"Coffee beans\" for cake  \"Red Velvet Strawberry Cake\"
//REMOVED ingredient \"Vanilla extract\" for cake  \"Red Velvet Strawberry Cake\"
//CHANGED unit for ingredient \"Flour\" for cake  \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"
//CHANGED unit count for ingredient \"Strawberries\" for cake  \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"
//REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake  \"Red Velvet Strawberry Cake\"
