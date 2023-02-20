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
