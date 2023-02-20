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

		// abort if no token is found
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized [No Token provided]"})
			return
		}

		// abort if token is not valid
		if !strings.HasPrefix(authHeader, "Token ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized	[Token not valid]"})
			return
		}
		token := strings.Split(authHeader, " ")[1]

		// check token against config
		// if token is not found, return 401
		// if token is found, set endcustomer, tsig, master, product

		customerConfig, ok := CONFIG.CustomerConfigs[token]

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized [Token not found]"})
			return
		}

		// check ip against allowedipranges
		// if ip is not found, return 401
		clientip := c.ClientIP()
		parsedclientip := net.ParseIP(clientip)
		ipallowed := false
		for _, ipRange := range customerConfig.AllowedIPRanges {
			// parse iprange into ipnet using parseCIDR
			_, ipnet, err := net.ParseCIDR(ipRange)
			if err != nil {
				Log("Error parsing IP range in config: " + ipRange)
				continue
			}

			if ipnet.Contains(parsedclientip) {
				ipallowed = true
				break
			}
		}
		if !ipallowed {
			Log("IP not allowed: " + c.ClientIP() + " for endcustomer: " + customerConfig.Endcustomer)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized [IP not allowed]"})
			return
		}

		// set endcustomer, tsig, master, product
		c.Set("endcustomer", customerConfig.Endcustomer)
		c.Set("forcedmasters", customerConfig.ForcedMasters)
		c.Set("forcedproduct", customerConfig.ForcedProduct)

	}
}
