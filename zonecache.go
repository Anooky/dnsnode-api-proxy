package main

import "fmt"

// this file contains a cache for the zone data
// the key is the zone name, the value is the zone object
// the cache is updated every 5 minutes
// additionally, the cache is updated when a zone is added or deleted

var ZONECACHE map[string]Zone

func UpdateZoneCache() {
	// get all zones
	zones := DnsnodeGetAllZones()
	// create a new map
	zoneCache := make(map[string]Zone)
	// add all zones to the map
	for _, zone := range zones {
		zoneCache[zone.Name] = zone
	}
	ZONECACHE = zoneCache

	// log the number of Zones in the cache
	Log(fmt.Sprint(len(ZONECACHE)) + " zones in cache")
}

func GetZoneFromCache(zonename string) Zone {
	return ZONECACHE[zonename]
}

func RefreshZoneInCache(zonename string) {
	zone := DnsnodeGetZone(zonename)
	ZONECACHE[zonename] = zone
}

func DeleteZoneFromCache(zonename string) {
	delete(ZONECACHE, zonename)
}