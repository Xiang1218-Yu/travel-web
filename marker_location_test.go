package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapMarkersKeepsNamedEquatorPrimeMeridianLocation(t *testing.T) {
	dbMutex.Lock()
	db = Database{Plans: []TravelPlan{{ID: "equator", Title: "赤道旅行", StartDate: "2026-11-03", Location: Location{Name: "加纳阿克拉", Latitude: 0, Longitude: 0}}}}
	dbMutex.Unlock()
	t.Cleanup(func() { dbMutex.Lock(); db = Database{}; dbMutex.Unlock() })
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/markers", nil)
	getMapMarkers(rr, req)
	var got []Marker
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].ID != "plan-equator" {
		t.Fatalf("具名的 0,0 坐标旅行没有出现在地图标记中: %#v", got)
	}
}
