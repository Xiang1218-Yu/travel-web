package main

func validateDiaryFields(diary DiaryEntry) error {
	if diary.PlanID == "" || diary.Date == "" || diary.Title == "" {
		return errDiaryRequired
	}
	return nil
}

func validateDiaryForCreate(diary DiaryEntry, plans []TravelPlan) error {
	if err := validateDiaryFields(diary); err != nil {
		return err
	}
	return validateDiaryPlanReference(diary, plans)
}
