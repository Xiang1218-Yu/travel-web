package main

import "time"

type planPatch struct {
	Title       string           `json:"title"`
	Destination string           `json:"destination"`
	StartDate   string           `json:"start_date"`
	EndDate     string           `json:"end_date"`
	Location    *Location        `json:"location"`
	Itinerary   *[]ItineraryItem `json:"itinerary"`
	// IsPublic 使用指针类型，以便区分“字段未提交”（保留原值）与“显式设置为 false”（设为私密）。
	IsPublic    *bool            `json:"is_public"`
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
	// 仅当客户端显式提交了 is_public 字段时才更新公开状态，未提交则保留原值。
	if patch.IsPublic != nil {
		plan.IsPublic = *patch.IsPublic
	}
	plan.UpdatedAt = now.Format(time.RFC3339)
	return plan
}
