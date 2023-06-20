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
		Cake := cake{j.Cake[i].Name,
			j.Cake[i].Time,
			nil}
		for k := 0; k < len(j.Cake[i].Ingredients); k++ {
			Ingredient := ingredients{j.Cake[i].Ingredients[k].IngredientName,
				j.Cake[i].Ingredients[k].IngredientCount,
				j.Cake[i].Ingredients[k].IngredientUnit}
			Cake.Ingredients = append(Cake.Ingredients, Ingredient)
		}
		r.Cake = append(r.Cake, Cake)
	}
}

func (r *recipes) TakeXml(x JX.Xake) {
	for i := 0; i < len(x.Cake); i++ {
		Cake := cake{x.Cake[i].Name,
			x.Cake[i].Stovetime,
			nil}
		for k := 0; k < len(x.Cake[i].Ingredients.Item); k++ {
			Ingredient := ingredients{
				x.Cake[i].Ingredients.Item[k].Itemname,
				x.Cake[i].Ingredients.Item[k].Itemcount,
				x.Cake[i].Ingredients.Item[k].Itemunit,
			}
			Cake.Ingredients = append(Cake.Ingredients, Ingredient)
		}
		r.Cake = append(r.Cake, Cake)
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

func CakeAdd(old recipes, new recipes) {
	for i := 0; i < len(new.Cake); i++ {
		if new.Cake[i].Name != old.Cake[i].Name {
			fmt.Printf("ADDED cake \\\"%s\\\"\n", new.Cake[i].Name)
		}
	}
}

func CakeRemove(old recipes, new recipes) {
	for i := 0; i < len(old.Cake); i++ {
		if old.Cake[i].Name != new.Cake[i].Name {
			fmt.Printf("REMOVED cake \\\"%s\\\"\n", old.Cake[i].Name)
		}
	}
}

func IngredientAndTime(old recipes, new recipes) {
	for i := 0; i < len(old.Cake); i++ {
		for j := 0; j < len(new.Cake); j++ {
			if old.Cake[i].Name == new.Cake[j].Name {
				if old.Cake[i].Time != new.Cake[j].Time {
					fmt.Printf("CHANGED cooking time for cake \\\"Red Velvet Strawberry Cake\\\""+
						" - \\\"%s\\\" instead of \\\"%s\\\"\n",
						old.Cake[i].Time, new.Cake[j].Time)
				}
				CheckIngredient(old.Cake[i].Ingredients, new.Cake[j].Ingredients, old.Cake[i].Name, new.Cake[j].Name)
			}
		}
	}
}

func CheckIngredient(old []ingredients, new []ingredients, oldName string, newName string) {
	oldV := make(map[string]map[string]string, len(old))
	newV := make(map[string]map[string]string, len(new))
	for i := 0; i < len(old); i++ {
		oldV[old[i].IngredientName] = map[string]string{old[i].IngredientCount: old[i].IngredientUnit}
	}
	for i := 0; i < len(new); i++ {
		newV[new[i].IngredientName] = map[string]string{new[i].IngredientCount: new[i].IngredientUnit}
		_, addCake := oldV[new[i].IngredientName]
		if !addCake {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", new[i].IngredientName, newName)
		}
	}
	for i := 0; i < len(old); i++ {
		mapCoUnt, haveCake := newV[old[i].IngredientName]
		if !haveCake {
			fmt.Printf("REMOVED ingredient \\\"%s\\\" for cake  \\\"%s\\\"\n", old[i].IngredientName, oldName)
		} else {
			//l, k := mapCoUnt

			fmt.Println(mapCoUnt)
			//checkUnit(oldV[old[i].IngredientName], newV[new[i].IngredientName])

		}
	}
}

func checkUnit(old map[string]string, new map[string]string) {

	//fmt.Printf()
}

//CHANGED unit for ingredient \"Flour\" for cake  \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"
//CHANGED unit count for ingredient \"Strawberries\" for cake  \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"
//REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake  \"Red Velvet Strawberry Cake\"

//./compareDB --old original_database.xml --new stolen_database.json
//ADDED cake \"Moonshine Muffin\"
//REMOVED cake \"Blueberry Muffin Cake\"
//CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"
//ADDED ingredient \"Coffee beans\" for cake  \"Red Velvet Strawberry Cake\"
//REMOVED ingredient \"Vanilla extract\" for cake  \"Red Velvet Strawberry Cake\"
