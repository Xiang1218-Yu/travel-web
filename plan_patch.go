package main

import "time"

type planPatch struct {
	Title       *string `json:"title"`
	Destination *string `json:"destination"`
	IsPublic    *bool   `json:"is_public"`
}

func applyPlanPatch(plan TravelPlan, patch planPatch, now time.Time) TravelPlan {
	if patch.Title != nil {
		plan.Title = *patch.Title
	}
	if patchHasDestination(patch) {
		plan.Destination = destinationFromPatch(patch)
	}
	if patchChangesVisibility(patch) {
		plan.IsPublic = *patch.IsPublic
	}
	plan.UpdatedAt = now.Format(time.RFC3339)
	return plan
}
