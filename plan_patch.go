package main

import "time"

type planPatch struct {
	Title    string `json:"title"`
	IsPublic bool   `json:"is_public"`
}

func applyPlanPatch(plan TravelPlan, patch planPatch, now time.Time) TravelPlan {
	if patch.Title != "" {
		plan.Title = patch.Title
	}
	plan.IsPublic = patch.IsPublic
	plan.UpdatedAt = now.Format(time.RFC3339)
	return plan
}
