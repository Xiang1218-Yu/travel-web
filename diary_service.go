package main

import "time"

func prepareDiaryForCreate(diary DiaryEntry, plans []TravelPlan, now time.Time) (DiaryEntry, error) {
	if err := validateDiaryForCreate(diary, plans); err != nil {
		return DiaryEntry{}, err
	}
	diary.ID = generateID()
	diary.CreatedAt = now.Format(time.RFC3339)
	diary.UpdatedAt = diary.CreatedAt
	if diary.Images == nil {
		diary.Images = []string{}
	}
	return diary, nil
}
