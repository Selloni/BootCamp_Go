package DBReader

import (
	"encoding/xml"
	"errors"
)

type Xake struct {
	Cake []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"time"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name" xml:"ingredientName"`
			IngredientCount string `json:"ingredient_count" xml:"ingredientCount"`
			IngredientUnit  string `json:"ingredient_unit,omitempty"xml:"ingredientUnit"`
		} `json:"ingredients" xml:"ingredients"`
	} `json:"cake"xml:"cake"`
}

func (x *Xake) ReadFile(data []byte) {
	if err := xml.Unmarshal(data, &x); err != nil {
		errors.New("not parsing")
	}
}
