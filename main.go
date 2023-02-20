package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Log(message string) {
	fmt.Println(message)
}
func main() {

	// read configfile
	Log("Reading configfile...")
	readConfigFile()

	// update zone cache
	Log("Updating zone cache...")
	UpdateZoneCache()

	// create router
	Log("Initializing API Server...")
	router := gin.Default()

	// add middleware
	router.Use(TokenAuthMapper())

	// handle fatal errors
	router.Use(gin.Recovery())

	// define routes
	//router.GET("/apiv3/status/:zonename", GetZoneStatus)

	// start server
	port := 8080
	Log("Starting API Server on port " + fmt.Sprint(port) + "...")
	router.Run("localhost:" + fmt.Sprint(port))
}
