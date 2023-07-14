package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/grailbio/base/tsv"
)

type Data struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	// Location elastic.GeoPoint `json:"location`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name     Name     `json:"name"`
	Address  Address  `json:"address"`
	Phone    Phone    `json:"phone"`
	Location Location `json:"location"`
}

type Name struct {
	D string `json:"type"`
}
type Address struct {
	D string `json:"type"`
}
type Phone struct {
	D string `json:"type"`
}
type Location struct {
	D string `json:"type"`
}

func main() {

	pathToDataSets := flag.String("dSet", "data.csv", "path to csv file")
	indexName := flag.String("iName", "places", "index name")
	countOfData := flag.Int("cData", 0, "count of data u want to upload")
	flag.Parse()

	dataSet := ReadCSV(*pathToDataSets)

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
	res, err = client.Indices.PutMapping(
		strings.NewReader(buf.String()),
		client.Indices.PutMapping.WithIndex("places"),
		client.Indices.PutMapping.WithDocumentType("place"),
		client.Indices.PutMapping.WithIncludeTypeName(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *countOfData; i++ {
		place := setData(dataSet[i])
		myJson, err := jsonPack(place)
		if err != nil {
			log.Fatal(err)
		}

		request := esapi.IndexRequest{
			Index:        *indexName,
			DocumentID:   strconv.Itoa(i + 1),
			DocumentType: "place",
			Body:         strings.NewReader(string(myJson)),
			Refresh:      "true",
		}
		response, err := request.Do(context.Background(), client)
		if err != nil {
			log.Fatal(err)
		}
		if response.IsError() {
			log.Fatalln("Error indexing document")
		}
		fmt.Println("status:", response.Status())
		response.Body.Close()
	}
}

func ReadCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("could not open the file", err)
	}
	defer file.Close()
	var d []byte
	reader := tsv.NewReader(file)
	reader.Read(d)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("could not read file ", err)
	}
	return data
}

func setData(data []string) Data {
	id, _ := strconv.Atoi(data[0])
	lon, _ := strconv.ParseFloat(data[4], 64)
	lat, _ := strconv.ParseFloat(data[5], 64)
	return Data{
		ID:        id,
		Name:      data[1],
		Address:   data[2],
		Phone:     data[3],
		Longitude: lon,
		Latitude:  lat,
	}
}

func jsonPack(place Data) ([]byte, error) {
	res, err := json.Marshal(place)
	if err != nil {
		return nil, err
	}
	return res, nil
}
