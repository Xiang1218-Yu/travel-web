package main

func changesVisibility(p planPatch) bool {
	return p.IsPublic != nil
}

func visibilityValue(p planPatch) (bool, bool) {
	if p.IsPublic == nil {
		return false, false
	}
	return *p.IsPublic, true
}
