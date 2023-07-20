package main

import (
	"day04/ex00/swager/restapi"
	"day04/ex00/swager/restapi/operations"
	"errors"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
	"math"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned int i;
unsigned int argscharcount = 0;

char *ask_cow(char phrase[]) {
	int phrase_len = strlen(phrase);
	char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
	strcpy(buf, " ");

	for (i = 0; i < phrase_len + 2; ++i) {
		strcat(buf, "_");
	}

	strcat(buf, "\n< ");
	strcat(buf, phrase);
	strcat(buf, " ");
	strcat(buf, ">\n ");

	for (i = 0; i < phrase_len + 2; ++i) {
		strcat(buf, "-");
	}
	strcat(buf, "\n");
	strcat(buf, "        \\   ^__^\n");
	strcat(buf, "         \\  (oo)\\_______\n");
	strcat(buf, "            (__)\\       )\\/\\\n");
	strcat(buf, "                ||----w |\n");
	strcat(buf, "                ||     ||\n");
	return buf;
}
*/
import "C"

func CowSay() string {
	cs := C.CString("Thank you!")
	answer := C.GoString(C.ask_cow(cs))
	return answer
}

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

type candyInf struct {
	name  string
	price int
}

var (
	flagPort   = flag.Int("port", 3333, "Port to run this service on")
	candyPrice = make(map[string]int)
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandyServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *flagPort
	priceList := getPrise()
	fillMap(priceList)

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(
		func(params operations.BuyCandyParams) middleware.Responder {
			candy := struct {
				name  string
				count int64
				money int64
			}{
				swag.StringValue(params.Order.CandyType),
				swag.Int64Value(params.Order.CandyCount),
				swag.Int64Value(params.Order.Money),
			}
			_, ok := candyPrice[candy.name]
			if !ok {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: fmt.Sprintf("Incorrect input: we don't have %s", candy.name)})
			}

			if candy.count <= 0 || candy.money <= 0 {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: "Incorrect input: money or count can't be zero or negative"})
			}

			change, err := calcChange(candy.count, int64(candyPrice[candy.name]), candy.money)
			if err != nil {
				return operations.NewBuyCandyPaymentRequired().WithPayload(
					&operations.BuyCandyPaymentRequiredBody{Error: err.Error()})
			}

			return operations.NewBuyCandyCreated().WithPayload(
				&operations.BuyCandyCreatedBody{Change: change, Thanks: CowSay()})
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

func getPrise() []candyInf {
	return []candyInf{
		{name: "CE", price: 10},
		{name: "AA", price: 15},
		{name: "NT", price: 17},
		{name: "DE", price: 21},
		{name: "YR", price: 23},
	}
}

func fillMap(priceList []candyInf) {
	for _, i := range priceList {
		candyPrice[i.name] = i.price
	}
}

func calcChange(count, candy, money int64) (int64, error) {
	change := money - (candy * count)
	if change < 0 {
		return 0, errors.New(fmt.Sprintf("You need %d more money!", int(math.Abs(float64(change)))))
	}
	return change, nil
}
