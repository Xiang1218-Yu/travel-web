package main

import (
	"fmt"
	"time"
)

func prepareDiaryForCreate(diary DiaryEntry, plans []TravelPlan, now time.Time) (DiaryEntry, error) {
	if diary.PlanID == "" || diary.Date == "" || diary.Title == "" {
		return DiaryEntry{}, fmt.Errorf("计划ID、日期和标题为必填项")
	}
	plan, ok := planForDiary(plans, diary.PlanID)
	if !ok {
		return DiaryEntry{}, fmt.Errorf("旅行计划不存在")
	}
	if err := validateDiaryDate(diary, plan); err != nil {
		return DiaryEntry{}, err
	}
	return finalizeDiaryForStorage(diary, now), nil
}

func finalizeDiaryForStorage(diary DiaryEntry, now time.Time) DiaryEntry {
	diary.ID = generateID()
	diary.CreatedAt = now.Format(time.RFC3339)
	diary.UpdatedAt = diary.CreatedAt
	if diary.Images == nil {
		diary.Images = []string{}
	}
	return diary
}
