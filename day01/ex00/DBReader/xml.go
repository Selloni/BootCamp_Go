package DBReader

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type Xake struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    Cake     `xml:"cake"`
}

type Cake []struct {
	XMLName     xml.Name    `xml:"cake" json:"-"`
	Name        string      `xml:"name"`
	Stovetime   string      `xml:"stovetime"`
	Ingredients Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	Text xml.Name `xml:"ingredients" json:"-"`
	Item []Item   `xml:"item"`
}

type Item struct {
	XMLName   xml.Name `xml:"item" json:"-"`
	Itemname  string   `xml:"itemname"`
	Itemcount string   `xml:"itemcount"`
	Itemunit  string   `xml:"itemunit"`
}

func (x *Xake) ReadFile(data []byte) {
	fmt.Println(string(data))

	if err := xml.Unmarshal(data, &x); err != nil {
		errors.New("not parsing")
	}
	fmt.Println(x.Cake)
	fmt.Println(x)
}

func (x *Xake) Write() []byte {
	data, err := json.MarshalIndent(x, "", "\t")
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
