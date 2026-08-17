package main

func collectMapMarkers(plans []TravelPlan, diaries []DiaryEntry, onlyPublic bool) []Marker {
	markers := make([]Marker, 0)
	visibility := visiblePlanIDs(plans)
	for _, plan := range plans {
		if !markerIsVisible(plan.ID, visibility, onlyPublic) || !hasMarkerLocation(plan.Location) {
			continue
		}
		lat, lng := markerCoordinates(plan.Location)
		markers = append(markers, Marker{ID: "plan-" + plan.ID, Type: "plan", PlanID: plan.ID, Title: plan.Title, Latitude: lat, Longitude: lng, Date: plan.StartDate})
	}
	for _, diary := range diaries {
		if !markerIsVisible(diary.PlanID, visibility, onlyPublic) || !hasMarkerLocation(diary.Location) {
			continue
		}
		lat, lng := markerCoordinates(diary.Location)
		markers = append(markers, Marker{ID: "diary-" + diary.ID, Type: "diary", PlanID: diary.PlanID, Title: diary.Title, Latitude: lat, Longitude: lng, Date: diary.Date})
	}
	return markers
}
