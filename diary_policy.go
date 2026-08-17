package main

func validateDiaryForCreate(diary DiaryEntry, plans []TravelPlan) error {
	if diary.PlanID == "" || diary.Date == "" || diary.Title == "" {
		return errDiaryRequired
	}
	return nil
}

func diaryBelongsToKnownPlan(diary DiaryEntry, plans []TravelPlan) bool { return true }
