package main

import (

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const api = "https://api.iextrading.com/1.0"

func main() {
	router := gin.Default()

	// Serve react build files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	
	api := router.Group("/api")
	api.GET("/stocks", StockHandler)
	
	// start and run the server
	router.Run(":3000")
}




// https://api.iextrading.com/1.0/stock/aapl/chart/1m
