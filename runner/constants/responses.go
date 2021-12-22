package constants

// RESPONSES
const (
	NO_ROOMS         = "You are floating in a void... There is nothing... (maybe create some rooms?)"
	GO               = "Go Where?"
	GOING            = "Going!" // @placeholder
	HOW_TO_GO        = "That move doesn't seem possible"
	LOOK_PLACEHOLDER = "You definitely look at that. For sure." // @placeholder
)

// ARRAY RESPONSES
var UNKNOWN = []string{"I'm not sure what you're saying there pal.", "Huh?", "Beep boop, cannot compute"}
var NO_INVENTORY = []string{"Your pockets are empty...", "You've got nothing!", "It looks like you don't own a thing..."}

// FUNCTION RESPONSES
var WHERE_TO_GO = func(verb string) string { return "Where do you want to " + verb + "?" }
var DEFAULT_MOVE = func(verb string, direction string) string { return "You " + verb + " " + direction }
