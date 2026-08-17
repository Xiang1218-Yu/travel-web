package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestUpdatePlanChangesDestination(t *testing.T) {
	dbMutex.Lock()
	oldDB, oldPath := db, dataPath
	db = Database{Plans: []TravelPlan{{ID: "p1", Title: "周末", Destination: "厦门", StartDate: "2026-10-01"}}}
	dataPath = filepath.Join(t.TempDir(), "db.json")
	dbMutex.Unlock()
	t.Cleanup(func() { dbMutex.Lock(); db, dataPath = oldDB, oldPath; dbMutex.Unlock() })
	req := httptest.NewRequest(http.MethodPut, "/api/plans/p1", bytes.NewBufferString(`{"destination":"泉州"}`))
	res := httptest.NewRecorder()
	updatePlan(res, req, "p1")
	if res.Code != http.StatusOK {
		t.Fatal(res.Code)
	}
	dbMutex.RLock()
	got := db.Plans[0].Destination
	dbMutex.RUnlock()
	if got != "泉州" {
		t.Fatalf("destination=%q, want 泉州", got)
	}
}
