package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateDiaryRejectsUnknownPlan(t *testing.T) {
	dbMutex.Lock()
	db = Database{Plans: []TravelPlan{{ID: "known-plan", Title: "旅行", Destination: "厦门", StartDate: "2026-09-01"}}}
	dbMutex.Unlock()
	t.Cleanup(func() { dbMutex.Lock(); db = Database{}; dbMutex.Unlock() })
	req := httptest.NewRequest(http.MethodPost, "/api/diaries", strings.NewReader(`{"plan_id":"missing-plan","date":"2026-09-01","title":"孤儿游记"}`))
	rr := httptest.NewRecorder()
	createDiary(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("未知计划仍被创建，status=%d body=%s", rr.Code, rr.Body.String())
	}
}
