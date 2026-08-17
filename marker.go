package main

type Marker struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	PlanID    string  `json:"plan_id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Date      string  `json:"date"`
}
