package main

func destinationFromPatch(p planPatch) string {
	if p.Destination == nil {
		return ""
	}
	return *p.Destination
}
