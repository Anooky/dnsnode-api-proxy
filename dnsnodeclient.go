package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func DnsnodeMakeRequest(url string, method string, body string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Token " + DNSNODE_TOKEN},
	}

	// add body if method is POST
	if method == "POST" {
		req.Body = ioutil.NopCloser(strings.NewReader(body))
	}

	res, err := client.Do(req)
	// check http status
	if res.StatusCode < 200 || res.StatusCode > 299 {
		// get the response body
		body, _ := ioutil.ReadAll(res.Body)
		stringbody := string(body)
		log.Println("HTTP status code: ", res.StatusCode, " for url: ", url, " method: ", method, " body: ", stringbody)

		// return stringbody as errormessage
		return nil, errors.New(stringbody)

	}
	return res, err

}

func DnsnodeGetZone(zonename string) Zone {

	url := NETNOD_BASE_URL + "zone/" + zonename

	res, err := DnsnodeMakeRequest(url, "GET", "")
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
	res, err := DnsnodeMakeRequest(url, "GET", "")
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

	res, err := DnsnodeMakeRequest(url, "GET", "")
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

func DnsnodeCreateZone(zonename string, endcustomer string, masters []Master, product string) (bool, error) {
	url := NETNOD_BASE_URL + "zone/"

	zone := Zone{
		Name:        zonename,
		Masters:     masters,
		Product:     product,
		Endcustomer: endcustomer,
	}

	zoneJson, err := json.Marshal(zone)
	if err != nil {
		return false, err
	}

	// log tze zoneJson
	log.Println(string(zoneJson))

	_, err = DnsnodeMakeRequest(url, "POST", string(zoneJson))

	// handle error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DnsnodeDeleteZone(zonename string) (bool, error) {
	url := NETNOD_BASE_URL + "zone/" + zonename

	_, err := DnsnodeMakeRequest(url, "DELETE", "")

	// handle error
	if err != nil {
		return false, err
	}
	return true, nil
}
