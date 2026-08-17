package main

func patchChangesVisibility(p planPatch) bool { return p.IsPublic != nil }
