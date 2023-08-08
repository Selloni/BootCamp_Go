package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func main() {
	p := UnknownPlant{"Alisa", "foot", 198}
	describePlant(p)
	s := AnotherUnknownPlant{222, "dog", 12}
	describePlant(s)
}

func describePlant(plan interface{}) {
	pType := reflect.TypeOf(plan)
	pValue := reflect.ValueOf(plan)
	for i := 0; i < pType.NumField(); i++ {
		field := pType.Field(i)
		value := pValue.Field(i)
		tag := field.Tag
		fmt.Print(field.Name)
		if tag != "" {
			fmt.Printf("(%s)", tag)
		}
		fmt.Printf(":%v\n", value)
	}
}
