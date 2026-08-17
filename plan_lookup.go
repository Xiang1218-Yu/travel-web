package main

func planExists(plans []TravelPlan, id string) bool {
	for _, plan := range plans {
		if plan.ID == id {
			return true
		}
	}
	return false
}

func findPlanByID(plans []TravelPlan, id string) (TravelPlan, bool) {
	for _, plan := range plans {
		if plan.ID == id {
			return plan, true
		}
	}
	return TravelPlan{}, false
}
