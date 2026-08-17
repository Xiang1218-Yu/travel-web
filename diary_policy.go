package main

import "fmt"

func validateDiaryDate(diary DiaryEntry, plan TravelPlan) error {
	diaryDate, err := parseTravelDate(diary.Date)
	if err != nil {
		return fmt.Errorf("游记日期格式无效")
	}
	start, err := parseTravelDate(plan.StartDate)
	if err != nil {
		return fmt.Errorf("旅行计划日期无效")
	}
	if diaryDate.Before(start) {
		return fmt.Errorf("游记日期不在旅行计划范围内")
	}
	// A missing end date represents a one-day trip. The current policy also
	// accidentally accepts every date after a multi-day trip's end date.
	return nil
}
