package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdatePlanPreservesVisibilityWhenOnlyTitleChanges(t *testing.T) {
	dbMutex.Lock()
	db = Database{Plans: []TravelPlan{{ID: "plan-public", Title: "旧标题", Destination: "厦门", StartDate: "2026-09-01", IsPublic: true}}}
	dbMutex.Unlock()
	t.Cleanup(func() { dbMutex.Lock(); db = Database{}; dbMutex.Unlock() })

	req := httptest.NewRequest(http.MethodPut, "/api/plans/plan-public", strings.NewReader(`{"title":"新标题"}`))
	rr := httptest.NewRecorder()
	updatePlan(rr, req, "plan-public")
	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	var got TravelPlan
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if !got.IsPublic {
		t.Fatalf("只更新标题后公开状态被改成 false: %#v", got)
	}
	if got.Title != "新标题" {
		t.Fatalf("title=%q", got.Title)
	}
}
