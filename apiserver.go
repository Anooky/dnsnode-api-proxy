package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureEndcustomer(context *gin.Context, zonename string) {

	// get zone from cache
	zone := GetZoneFromCache(zonename)

	// abort if zone is not found
	if zone.Name == "" {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Zone not found"})
		return
	}

	// get endcustomer from context
	endcustomer := context.MustGet("endcustomer").(string)

	// abort if endcustomer is not allowed to access the zone
	if zone.Endcustomer != endcustomer {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

}

func GetZoneStatus(context *gin.Context, zonename string) {

	// ensure endcustomer is allowed to access the zone
	EnsureEndcustomer(context, zonename)

	zonestatus := DnsnodeZoneStatus(zonename)
	context.IndentedJSON(http.StatusOK, zonestatus)

}
