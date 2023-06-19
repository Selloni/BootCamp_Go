package interal

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"root/ex01/JX"
)

type recipes struct {
	Cake []cake
}

type cake struct {
	Name        string
	Time        string
	Ingredients []ingredients
}

type ingredients struct {
	IngredientName  string
	IngredientCount string
	IngredientUnit  string
}

func (r *recipes) TakeJson(j JX.Jake) {
	for i := 0; i < len(j.Cake); i++ {
		fmt.Println(i)
		Cake := cake{j.Cake[i].Name,
			j.Cake[i].Time,
			nil}
		for k := 0; k < len(j.Cake[i].Ingredients); k++ {
			Ingredient := ingredients{j.Cake[i].Ingredients[k].IngredientName,
				j.Cake[i].Ingredients[k].IngredientCount,
				j.Cake[i].Ingredients[k].IngredientCount}
		}

	}
}

func (r *recipes) TakeXml(x JX.Xake) {
	for i := 0; i < len(x.Cake); i++ {
		r.Cake[i].Name = x.Cake[i].Name
		r.Cake[i].Time = x.Cake[i].Stovetime
		r.Cake[i].Ingredients = nil
		for k := 0; k < len(x.Cake[i].Ingredients.Item); k++ {
			r.Cake[i].Ingredients[k].IngredientName = x.Cake[i].Ingredients.Item[k].Itemname
			r.Cake[i].Ingredients[k].IngredientCount = x.Cake[i].Ingredients.Item[k].Itemcount
			r.Cake[i].Ingredients[k].IngredientUnit = x.Cake[i].Ingredients.Item[k].Itemunit
		}
	}
}

var (
	newCakeNames []string
	oldCakeNames []string
)

func ReadFile(str *string) []byte {
	data, err := os.ReadFile(*str)
	if err != nil {
		errors.New("not open file")
	}
	return data
}

func ParsingJson(data []byte, j *JX.Jake) {
	if err := json.Unmarshal(data, &j); err != nil {
		errors.New("not parsing json")
	}
}

func ParsingXml(data []byte, x *JX.Xake) {
	if err := xml.Unmarshal(data, &x); err != nil {
		errors.New("not parsing xml")
	}
}

//func compIng(origin *DBReader.Jake, stolen *DBReader.Jake, count int) {
//	find := false
//	for i := range origin.Cake[count].Ingredients {
//		for j := range stolen.Cake[count].Ingredients {
//			org := origin.Cake[count].Ingredients[i]
//			stl := stolen.Cake[count].Ingredients[j]
//			if org.IngredientName == stl.IngredientName {
//				find = true
//			}
//			if find {
//				if org.IngredientCount != org.IngredientCount {
//					fmt.Printf("CHANGED unit count for ingredient %s for cake %s - %d instead of %d", org, origin.Cake[count], org.IngredientCount, stl.IngredientCount)
//				}
//			} else {
//				fmt.Printf("REMOVED ingredient \\%s\\ for cake \\%s\\", org, origin.Cake[count])
//			}
//		}
//	}
//}

//./compareDB --old original_database.xml --new stolen_database.json
//ADDED cake \"Moonshine Muffin\"
//REMOVED cake \"Blueberry Muffin Cake\"
//CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"
//ADDED ingredient \"Coffee beans\" for cake  \"Red Velvet Strawberry Cake\"
//REMOVED ingredient \"Vanilla extract\" for cake  \"Red Velvet Strawberry Cake\"
//CHANGED unit for ingredient \"Flour\" for cake  \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"
//CHANGED unit count for ingredient \"Strawberries\" for cake  \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"
//REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake  \"Red Velvet Strawberry Cake\"
