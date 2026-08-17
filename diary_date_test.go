package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateDiaryRejectsDateAfterPlanEnd(t *testing.T) {
	dbMutex.Lock()
	oldDB, oldPath := db, dataPath
	db = Database{Plans: []TravelPlan{{ID: "plan-1", Title: "北海道", Destination: "札幌", StartDate: "2026-09-10", EndDate: "2026-09-12"}}, Diaries: []DiaryEntry{}}
	dataPath = filepath.Join(t.TempDir(), "db.json")
	dbMutex.Unlock()
	t.Cleanup(func() {
		dbMutex.Lock()
		db, dataPath = oldDB, oldPath
		dbMutex.Unlock()
	})

	body, _ := json.Marshal(DiaryEntry{PlanID: "plan-1", Date: "2026-09-13", Title: "返程后补记"})
	req := httptest.NewRequest(http.MethodPost, "/api/diaries", bytes.NewReader(body))
	res := httptest.NewRecorder()
	createDiary(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body=%s", res.Code, http.StatusBadRequest, res.Body.String())
	}
	if !strings.Contains(res.Body.String(), "旅行计划范围") {
		t.Fatalf("unexpected response: %s", res.Body.String())
	}
}
