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
	// An empty end date represents a one-day trip, so the trip ends on
	// the start date. Reject any diary dated after the plan's end date
	// so post-trip entries cannot be mixed into a plan.
	end := start
	if plan.EndDate != "" {
		end, err = parseTravelDate(plan.EndDate)
		if err != nil {
			return fmt.Errorf("旅行计划日期无效")
		}
	}
	if diaryDate.After(end) {
		return fmt.Errorf("游记日期不在旅行计划范围内")
	}
	return nil
}
