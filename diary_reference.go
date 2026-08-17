package main

func validateDiaryPlanReference(diary DiaryEntry, plans []TravelPlan) error {
	if !planExists(plans, diary.PlanID) {
		return errDiaryPlanMissing
	}
	return nil
}
