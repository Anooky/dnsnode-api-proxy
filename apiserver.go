package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureEndcustomer(context *gin.Context, zonename string) bool {

	// get zone from cache
	zone := GetZoneFromCache(zonename)

	// abort if zone is not found
	if zone.Name == "" {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Zone not found"})
		return false
	}

	// get endcustomer from context
	endcustomer := context.MustGet("endcustomer").(string)

	// abort if endcustomer is not allowed to access the zone
	if zone.Endcustomer != endcustomer {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized [Endcustomer not allowed to access zone]"})
		return false
	}
	return true

}

func GetZoneStatus(context *gin.Context) {
	zonename := context.Param("zonename")

	// ensure endcustomer is allowed to access the zone
	if !EnsureEndcustomer(context, zonename) {
		return
	}

	zonestatus := DnsnodeZoneStatus(zonename)
	context.IndentedJSON(http.StatusOK, zonestatus)

}

func CreateZone(context *gin.Context) {
	// parse zone from request body
	var zone Zone
	if err := context.ShouldBindJSON(&zone); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request [Invalid JSON]"})
		return
	}

	// force endcustomer to be the one from the context
	zone.Endcustomer = context.MustGet("endcustomer").(string)

	// force product to be the one from the context
	zone.Product = context.MustGet("forcedproduct").(string)

	// force masters to be the ones from the context
	zone.Masters = context.MustGet("forcedmasters").([]Master)

	// create zone
	ok, err := DnsnodeCreateZone(zone.Name, zone.Endcustomer, zone.Masters, zone.Product)

	// check if zone was created
	if !ok {
		// Log error
		Log("Error creating zone: " + zone.Name + " Error: " + err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}
	// return success message in body
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Zone created"})

	// add zone to cache
	RefreshZoneInCache(zone.Name)

}

func DeleteZone(context *gin.Context) {
	zonename := context.Param("zonename")

	// ensure endcustomer is allowed to access the zone
	if !EnsureEndcustomer(context, zonename) {
		return
	}

	// delete zone
	ok, err := DnsnodeDeleteZone(zonename)

	// check if zone was deleted
	if !ok {
		// Log error
		Log("Error deleting zone: " + zonename + " Error: " + err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	// return success message in body
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Zone deleted"})

	// remove zone from cache
	RemoveZoneFromCache(zonename)

}

func GetZones(context *gin.Context) {
	// get endcustomer from context
	endcustomer := context.MustGet("endcustomer").(string)

	// get zones from cache
	zones := GetZonesFromCache(endcustomer)

	// return zones
	context.IndentedJSON(http.StatusOK, zones)

}

// get single Zone
func GetZone(context *gin.Context) {
	zonename := context.Param("zonename")

	// ensure endcustomer is allowed to access the zone
	if !EnsureEndcustomer(context, zonename) {
		return
	}

	// get zone from cache
	zone := DnsnodeGetZone(zonename)

	// return zone
	context.IndentedJSON(http.StatusOK, zone)

}

// Zone statistics
func GetZoneStatistics(context *gin.Context) {
	zonename := context.Param("zonename")

	// ensure endcustomer is allowed to access the zone
	if !EnsureEndcustomer(context, zonename) {
		return
	}

	// get zone statistics
	statistics := DnsnodeZoneStatistics(zonename)

	// return statistics
	context.IndentedJSON(http.StatusOK, statistics)

}

// Zone anomalies
func GetZoneAnomaliesSerial(context *gin.Context) {
	zonename := context.Param("zonename")

	// ensure endcustomer is allowed to access the zone
	if !EnsureEndcustomer(context, zonename) {
		return
	}

	// get zone anomalies
	anomalies := DnsnodeZoneAnomaliesSerial(zonename)

	// return anomalies
	context.IndentedJSON(http.StatusOK, anomalies)

}
