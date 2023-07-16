package main

import (
	"github.com/gin-gonic/gin"
)

//@tittle Candy Server
//@version 1.0
//@description summary of the candy order

// @host localhost:3333
// @BasePath  /buy_candy/
func main() {
	r := gin.Default()

	r.Run(":3333")

}
