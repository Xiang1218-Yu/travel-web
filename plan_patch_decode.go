package main

import "net/http"

func decodePlanPatch(r *http.Request) (planPatch, error) {
	var p planPatch
	if err := parseJSON(r, &p); err != nil {
		return planPatch{}, err
	}
	return p, nil
}
