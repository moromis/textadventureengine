package constants

// RESPONSES
const (
	NO_ROOMS         = "You are floating in a void... There is nothing... (maybe create some rooms?)"
	GO               = "Go Where?"
	GOING            = "Going!"                                 // @placeholder
	LOOK_PLACEHOLDER = "You definitely look at that. For sure." // @placeholder
)

// ARRAY RESPONSES
var HOW_TO_GO = []string{"That move doesn't seem possible", "I don't think you can go that way!", "Not sure how you'd go there..."}
var UNKNOWN = []string{"I'm not sure what you're saying there, pal.", "Huh?", "Beep boop, cannot compute"}
var NO_INVENTORY = []string{"Your pockets are empty...", "You've got nothing!", "It looks like you don't own a thing..."}
var STUCK = []string{"There's no way out.", "You're stuck here... Forever?", "You can't see anywhere to go."}

// FUNCTION RESPONSES
var WHERE_TO_GO = func(verb string) string { return "Where do you want to " + verb + "?" }
var DEFAULT_MOVE = func(verb string, direction string) string { return "You " + verb + " " + direction }
