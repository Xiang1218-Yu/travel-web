package main

func normalizePlanPatch(p planPatch) planPatch {
	return p
}

func patchHasChanges(p planPatch) bool {
	return p.Title != nil || p.Destination != nil || p.StartDate != nil ||
		p.EndDate != nil || p.Location != nil || p.Itinerary != nil ||
		changesVisibility(p)
}
