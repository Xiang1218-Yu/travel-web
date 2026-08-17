package main

func findPlanByID(plans []TravelPlan, id string) (TravelPlan, bool) {
	for _, p := range plans {
		if p.ID == id {
			return p, true
		}
	}
	return TravelPlan{}, false
}

func planExists(plans []TravelPlan, id string) bool {
	_, ok := findPlanByID(plans, id)
	return ok
}
