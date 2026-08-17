package main

func validateDiaryPlanReference(diary DiaryEntry, plans []TravelPlan) error {
	if !planExists(plans, diary.PlanID) {
		return errDiaryPlanMissing
	}
	return nil
}

func diaryBelongsToKnownPlan(diary DiaryEntry, plans []TravelPlan) bool {
	return planExists(plans, diary.PlanID)
}
