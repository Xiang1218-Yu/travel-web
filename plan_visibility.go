package main

func changesVisibility(p planPatch) bool { return p.IsPublic != nil }
