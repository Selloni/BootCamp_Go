package DBReader

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type Jake struct {
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

func (j *Jake) ReadFile(data []byte) {
	if err := json.Unmarshal(data, &j); err != nil {
		errors.New("not parsing")
	}
}

func (j *Jake) Write() []byte {
	data, err := xml.MarshalIndent(j, "", "\t")
	if err != nil {
		errors.New("no write xml")
	}
	return data
}

func (j *Jake) Create(data []byte) {
	file, err := os.Create("dock_for_parsing/stolen_database.xml")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}
