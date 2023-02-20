package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get Bearer Token from header
		authHeader := c.GetHeader("Authorization")
		token := strings.Split(authHeader, " ")[1]

		// check token against config
		// if token is not found, return 401
		// if token is found, set endcustomer, tsig, master, product

		customerConfig, ok := CONFIG.CustomerConfigs[token]

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// check ip against allowedipranges
		// if ip is not found, return 401
		ipallowed := false
		for _, ipRange := range customerConfig.AllowedIPRanges {
			clientip := c.ClientIP()
			parsedclientip := net.ParseIP(clientip)
			if ipRange.Contains(parsedclientip) {
				ipallowed = true
				break
			}
		}
		if !ipallowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// set endcustomer, tsig, master, product
		c.Set("endcustomer", customerConfig.Endcustomer)
		c.Set("forcedmasters", customerConfig.ForcedMasters)
		c.Set("forcedproduct", customerConfig.ForcedProduct)

	}
}
