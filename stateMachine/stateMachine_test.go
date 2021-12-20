package stateMachine

import (
	"testing"

	"tae.com/constants"
	"tae.com/testObjects"
)

func TestParseInput(t *testing.T) {
	SetupStateMachine(testObjects.TestMap, testObjects.TestMapWidth, testObjects.Here, testObjects.TestInventory)

	tests := []struct {
		testString string
		result     string
	}{
		{"", constants.UNKNOWN},
		{"asdf", constants.UNKNOWN},
	}

	for _, test := range tests {
		result := ParseInput(test.testString)
		if result != test.result {
			t.Errorf("Parsing %s failed, got: %s, want: %s.", test.testString, result, test.result)
		}
	}
}

func TestParseMovement(t *testing.T) {
	SetupStateMachine(testObjects.TestMap, testObjects.TestMapWidth, testObjects.Here, testObjects.TestInventory)

	emptyArray := []string{}
	southArray := []string{"go", "south"}
	northArray := []string{"go", "north"}
	println(testObjects.Here.ValidMoves[constants.SOUTH])
	tests := []struct {
		testString string
		splitInput []string
		result     string
	}{
		{"", emptyArray, constants.UNKNOWN},
		{"asdf", emptyArray, constants.UNKNOWN},
		{"go", emptyArray, constants.WHERE_TO_GO("go")},
		{"go", southArray, testObjects.Here.ValidMoves[constants.SOUTH]}, // TODO: in the future this will actually respond with the movement response from the room?
		{"walk", emptyArray, constants.WHERE_TO_GO("walk")},
		{"walk", northArray, testObjects.There.ValidMoves[constants.NORTH]}, // TODO: in the future this will actually respond with the movement response from the room?
	}

	for _, test := range tests {
		result := parseMovement(test.testString, test.splitInput, constants.UNKNOWN)
		if result != test.result {
			t.Errorf("Parsing %s, %v failed, got: %s, want: %s.", test.testString, test.splitInput, result, test.result)
		}
	}
}

// func TestMovePlayer(t *testing.T) {
// 	// just for testing TODO: remove
// 	var here = &constants.Entity{
// 		name:     "Here",
// 		desc:     "A nice place",
// 		location: [2]int{0, 0},
// 		validMoves: map[string]string{
// 			constants.SOUTH: "You amble from here to there",
// 		},
// 	}

// 	// just for testing TODO: remove
// 	var there = &constants.Entity{
// 		name:     "There",
// 		desc:     "An okay place, I guess",
// 		location: [2]int{0, 1},
// 		validMoves: map[string]string{
// 			constants.NORTH: "You mobilize from there to here",
// 		},
// 	}

// 	currentRoom = here
// 	emptyArray := []string{}
// 	northArray := []string{"go", "north"}
// 	tests := []struct {
// 		testString string
// 		splitInput []string
// 		result     string
// 	}{
// 		{"", emptyArray, constants.UNKNOWN},
// 		{"asdf", emptyArray, constants.UNKNOWN},
// 		{"go", emptyArray, WHERE_TO_GO("go")},
// 		{"go", northArray, GOING}, // TODO: in the future this will actually respond with the movement response from the room?
// 		{"walk", emptyArray, WHERE_TO_GO("walk")},
// 		{"walk", northArray, GOING}, // TODO: in the future this will actually respond with the movement response from the room?
// 	}

// 	for _, test := range tests {
// 		result := parseMovement(test.testString, test.splitInput, constants.UNKNOWN)
// 		if result != test.result {
// 			t.Errorf("Parsing %s, %v failed, got: %s, want: %s.", test.testString, test.splitInput, result, test.result)
// 		}
// 	}
// }
