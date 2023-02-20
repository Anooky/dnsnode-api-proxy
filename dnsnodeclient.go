package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Tsig struct {
	Name string `json:"name"`
	Alg  string `json:"alg"`
	Key  string `json:"key"`
}

type Master struct {
	IP   string `json:"ip"`
	Tsig string `json:"tsig"`
}

type Zone struct {
	Name        string   `json:"name"`
	Masters     []Master `json:"masters"`
	Product     string   `json:"product"`
	Endcustomer string   `json:"endcustomer"`
}

type Distmaster struct {
	IPv4Address string `json:"ipv4_address"`
	IPv6Address string `json:"ipv6_address"`
	Serial      int    `json:"serial"`
	Timestamp   int    `json:"timestamp"`
}

type Site struct {
	Name      string `json:"name"`
	Serial    int    `json:"serial"`
	Timestamp int    `json:"timestamp"`
}

type Zonestatus struct {
	ConfiguredSites       int          `json:"configured_sites"`
	ConfiguredDistmasters int          `json:"configured_distmasters"`
	CurrentSerial         int          `json:"current_serial"`
	CurrentTimestamp      int          `json:"current_timestamp"`
	Distmasters           []Distmaster `json:"distmasters"`
	Sites                 []Site       `json:"sites"`
}

const NETNOD_BASE_URL = "https://dnsnodeapi.netnod.se/apiv3/"

func DnsnodeMakeRequest(url string, method string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Token " + DNSNODE_TOKEN},
	}

	res, err := client.Do(req)
	// check http status
	if res.StatusCode != 200 {
		log.Fatal("HTTP status code: ", res.StatusCode, " for url: ", url, " method: ", method)
	}
	return res, err

}

func DnsnodeGetZone(zonename string) Zone {

	url := NETNOD_BASE_URL + "zone/" + zonename

	res, err := DnsnodeMakeRequest(url, "GET")
	// handle error
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var zone Zone
	jsonErr := json.Unmarshal(body, &zone)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return zone

}

func DnsnodeGetAllZones() []Zone {
	url := NETNOD_BASE_URL + "zone/"
	res, err := DnsnodeMakeRequest(url, "GET")
	// handle error
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var zones []Zone
	jsonErr := json.Unmarshal(body, &zones)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return zones
}

func DnsnodeZoneStatus(zonename string) Zonestatus {

	url := NETNOD_BASE_URL + "status/" + zonename

	res, err := DnsnodeMakeRequest(url, "GET")
	// handle error
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var zonestatus Zonestatus
	jsonErr := json.Unmarshal(body, &zonestatus)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return zonestatus

}
