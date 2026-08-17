package main

func validateDiaryForCreate(diary DiaryEntry, plans []TravelPlan) error {
	if diary.PlanID == "" || diary.Date == "" || diary.Title == "" {
		return errDiaryRequired
	}
	if err := validateDiaryPlanReference(diary, plans); err != nil {
		return err
	}
	return nil
}

func diaryBelongsToKnownPlan(diary DiaryEntry, plans []TravelPlan) bool {
	return planExists(plans, diary.PlanID)
}
