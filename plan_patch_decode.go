package main

import "net/http"

func decodePlanPatch(r *http.Request) (planPatch, error) {
	var patch planPatch
	if err := parseJSON(r, &patch); err != nil {
		return planPatch{}, err
	}
	return patch, nil
}
