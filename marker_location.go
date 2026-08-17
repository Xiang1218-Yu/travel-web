package main

func hasMarkerLocation(location Location) bool {
	return location.Name != "" || location.Latitude != 0 || location.Longitude != 0
}

func markerCoordinates(location Location) (float64, float64) {
	return location.Latitude, location.Longitude
}
