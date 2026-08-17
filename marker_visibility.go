package main

func visiblePlanIDs(plans []TravelPlan) map[string]bool {
	result := make(map[string]bool, len(plans))
	for _, plan := range plans {
		result[plan.ID] = plan.IsPublic
	}
	return result
}

func markerIsVisible(planID string, visibility map[string]bool, onlyPublic bool) bool {
	return !onlyPublic || visibility[planID]
}
