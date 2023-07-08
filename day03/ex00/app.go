package main

import (
	// "fmt"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"

	// "github.com/olivere/elastic/v7"
	"github.com/elastic/go-elasticsearch/esapi"
)

type Data struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	// Location elastic.GeoPoint `json:"location`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"katitude"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

type Properties []struct {
	Name     Name     `json:"name"`
	Address  Address  `json:"address"`
	Phone    Phone    `json:"phone"`
	Location Location `json:"location"`
}

type Name struct {
	S string
}

type Address struct {
	S string
}
type Phone struct {
	S string
}
type Location struct {
	S string
}

func main() {

	pathToDataSets := flag.String("dSet", "data.csv", "path to csv file")
	indexName := flag.String("iName", "places", "index name")
	countOfData := flag.Int("cData", 0, "count of data u want to upload")
	flag.Parse()

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal("failed to create a client")
	}
	res, err := client.Indices.Create("places")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var buf bytes.Buffer
	b := Schema{Properties{
		Name:     Name{"text"},
		Address:  Address{"text"},
		Phone:    Phone{"text"},
		Location: Location{"geo_point"},
	}}
	json.NewEncoder(&buf).Encode(b)
}

func ReadCSV(path string) [][]string {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal("didnt read fileCSV")
	}
	defer data.Close()
	reader := csv.NewReader(data)
	dataCSV, err := reader.ReadAll()
	if err != nil {
		log.Fatal("failed read all")
	}
	return dataCSV
}
