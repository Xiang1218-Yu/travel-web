package main

import (
	"encoding/json"
	"net/http"
)

func decodePlanPatch(r *http.Request) (planPatch, error) {
	var patch planPatch
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patch); err != nil {
		return planPatch{}, err
	}
	return normalizePlanPatch(patch), nil
}
