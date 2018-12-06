package main

import (
	"fmt"
	"flag"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

const api = "https://api.iextrading.com/1.0"

func main() {

	port := flag.String("port", "8080", "server port")
	flag.Parse()
	
	router := gin.Default()
	router.Use(cors.Default())

	// Serve react build files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	
	api := router.Group("/api")
	api.GET("/stocks", StockHandler)
	
	// start and run the server
	router.Run(fmt.Sprintf(":%s", *port))

}
