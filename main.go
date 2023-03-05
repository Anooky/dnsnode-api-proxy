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
	Log("Loaded configuration for " + fmt.Sprint(len(CONFIG.CustomerConfigs)) + " customers.")

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
	router.GET("/apiv3/status/:zonename", GetZoneStatus)
	router.GET("/apiv3/zone/", GetZones)
	router.POST("/apiv3/zone", CreateZone)
	router.DELETE("/apiv3/zone/:zonename", DeleteZone)
	router.GET("/apiv3/statistics/graph/:zonename", GetZoneStatistics)

	// start server
	port := 8080
	Log("Starting API Server on port " + fmt.Sprint(port) + "...")
	router.Run("0.0.0.0:" + fmt.Sprint(port))
}
