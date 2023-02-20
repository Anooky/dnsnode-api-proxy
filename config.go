package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

var DNSNODE_TOKEN string
var CONFIG Configfile

type CustomerConfig struct {
	Token           string       `json:"token"`
	Endcustomer     string       `json:"endcustomer"`
	ForcedMasters   []Master     `json:"forcedmasters"`
	ForcedProduct   string       `json:"forcedproduct"`
	AllowedIPRanges []*net.IPNet `json:"allowedipranges"`
}

// use customer's token as key for customerconfig
type Configfile struct {
	DNSNodeToken    string                    `json:"dnsnodetoken"`
	CustomerConfigs map[string]CustomerConfig `json:"customerconfigs"`
}

// this function reads the config file and parses its json content into a configfile object
func readConfigFile() {
	var config Configfile
	file, _ := ioutil.ReadFile("config.json")
	json.Unmarshal([]byte(file), &config)
	CONFIG = config
	DNSNODE_TOKEN = config.DNSNodeToken
}
