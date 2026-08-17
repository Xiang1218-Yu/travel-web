package main

func planForDiary(plans []TravelPlan, id string) (TravelPlan, bool) {
	for _, plan := range plans {
		if plan.ID == id {
			return plan, true
		}
	}
	return TravelPlan{}, false
}
