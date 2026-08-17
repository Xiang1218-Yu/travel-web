package main

import "time"

type planPatch struct {
	Title       string           `json:"title"`
	Destination string           `json:"destination"`
	StartDate   string           `json:"start_date"`
	EndDate     string           `json:"end_date"`
	Location    *Location        `json:"location"`
	Itinerary   *[]ItineraryItem `json:"itinerary"`
	IsPublic    bool             `json:"is_public"`
}

func applyPlanPatch(plan TravelPlan, patch planPatch, now time.Time) TravelPlan {
	if patch.Title != "" {
		plan.Title = patch.Title
	}
	if patch.Destination != "" {
		plan.Destination = patch.Destination
	}
	if patch.StartDate != "" {
		plan.StartDate = patch.StartDate
	}
	if patch.EndDate != "" {
		plan.EndDate = patch.EndDate
	}
	if patch.Location != nil {
		plan.Location = *patch.Location
	}
	if patch.Itinerary != nil {
		plan.Itinerary = *patch.Itinerary
	}
	plan.IsPublic = patch.IsPublic
	plan.UpdatedAt = now.Format(time.RFC3339)
	return plan
}
