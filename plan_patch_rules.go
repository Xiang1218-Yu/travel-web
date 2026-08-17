package main

func patchHasChanges(p planPatch) bool {
	return p.Title != "" || p.Destination != "" || p.StartDate != "" ||
		p.EndDate != "" || p.Location != nil || p.Itinerary != nil ||
		p.IsPublic != nil
}
