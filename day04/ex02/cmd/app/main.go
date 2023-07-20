package main

import (
	"day04/ex00/swager/restapi"
	"day04/ex00/swager/restapi/operations"
	"day04/ex02/cmd/caw-say"
	"errors"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
	"math"
)

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
				&operations.BuyCandyCreatedBody{Change: change, Thanks: caw_say.CowSay()})
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
