package main

import "time"

type planPatch struct {
	Title       *string          `json:"title"`
	Destination *string          `json:"destination"`
	StartDate   *string          `json:"start_date"`
	EndDate     *string          `json:"end_date"`
	Location    *Location        `json:"location"`
	Itinerary   *[]ItineraryItem `json:"itinerary"`
	IsPublic    *bool            `json:"is_public"`
}

func applyPlanPatch(plan TravelPlan, patch planPatch, now time.Time) TravelPlan {
	if patch.Title != nil {
		plan.Title = *patch.Title
	}
	if patch.Destination != nil {
		plan.Destination = *patch.Destination
	}
	if patch.StartDate != nil {
		plan.StartDate = *patch.StartDate
	}
	if patch.EndDate != nil {
		plan.EndDate = *patch.EndDate
	}
	if patch.Location != nil {
		plan.Location = *patch.Location
	}
	if patch.Itinerary != nil {
		plan.Itinerary = *patch.Itinerary
	}
	if changesVisibility(patch) {
		plan.IsPublic = *patch.IsPublic
	}
	plan.UpdatedAt = now.Format(time.RFC3339)
	return plan
}
