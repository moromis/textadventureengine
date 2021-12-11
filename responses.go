package main

// RESPONSES
const (
	UNKNOWN          = "I'm not sure what you're saying there pal."
	GO               = "Go Where?"
	GOING            = "Going!" // @placeholder
	HOW_TO_GO        = "That move doesn't seem possible"
	LOOK_PLACEHOLDER = "You definitely look at that. For sure." // @placeholder
)

// FUNCTION RESPONSES
var WHERE_TO_GO = func(verb string) string { return "Where do you want to " + verb + "?" }
var DEFAULT_MOVE = func(verb string, direction string) string { return "You " + verb + " " + direction }
