package JX

import (
	"encoding/xml"
)

type Xake struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    Cake     `xml:"cake" json:"cake"`
}

type Cake []struct {
	XMLName     xml.Name    `xml:"cake" json:"-"`
	Name        string      `xml:"name" json:"name"`
	Stovetime   string      `xml:"stovetime" json:"stovetime"`
	Ingredients Ingredients `xml:"ingredients"json:"ingredients"`
}

type Ingredients struct {
	Text xml.Name `xml:"ingredients" json:"-"`
	Item []Item   `xml:"item" json:"item"`
}

type Item struct {
	XMLName   xml.Name `xml:"item" json:"-"`
	Itemname  string   `xml:"itemname"json:"itemname"`
	Itemcount string   `xml:"itemcount"json:"itemcount"`
	Itemunit  string   `xml:"itemunit"json:"itemunit"`
}
