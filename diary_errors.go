package main

import "errors"

var errDiaryRequired = errors.New("计划ID、日期和标题为必填项")
var errDiaryPlanMissing = errors.New("旅行计划不存在")

func diaryCreateError(err error) string {
	if errors.Is(err, errDiaryPlanMissing) {
		return "旅行计划不存在"
	}
	return err.Error()
}
