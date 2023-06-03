package DBReader

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
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

func (x *Xake) Write() []byte {
	data, err := xml.Marshal(x)
	if err != nil {
		errors.New("no write xml")
	}
	return data
}

func (x *Xake) Create(data []byte) {
	file, err := os.Create("dock_for_parsing/stolen_database.json")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}

func (x *Xake) ReadFile(data []byte) {
	if err := xml.Unmarshal(data, &x); err != nil {
		errors.New("not parsing")
	}
}
