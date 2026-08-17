package main

func planForDiary(plans []TravelPlan, id string) (TravelPlan, bool) {
	for _, plan := range plans {
		if plan.ID == id {
			return plan, true
		}
	}
	return TravelPlan{}, false
}

func planEndDate(plan TravelPlan) string {
	if plan.EndDate == "" {
		return plan.StartDate
	}
	return plan.EndDate
}
