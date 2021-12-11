package main

import "testing"

func TestParseInput(t *testing.T) {
	setupStateMachine()

	tests := []struct {
		testString string
		result     string
	}{
		{"", UNKNOWN},
		{"asdf", UNKNOWN},
	}

	for _, test := range tests {
		result := parseInput(test.testString)
		if result != test.result {
			t.Errorf("Parsing %s failed, got: %s, want: %s.", test.testString, result, test.result)
		}
	}
}

func TestParseMovement(t *testing.T) {
	var here = &Entity{
		name:     "Here",
		desc:     "A nice place",
		location: [2]int{0, 0},
		validMoves: map[string]string{
			SOUTH: "You amble from here to there",
		},
	}
	var there = &Entity{
		name:     "There",
		desc:     "An okay place, I guess",
		location: [2]int{1, 0},
		validMoves: map[string]string{
			NORTH: "You mobilize from there to here",
		},
	}

	setupStateMachine()

	emptyArray := []string{}
	southArray := []string{"go", "south"}
	northArray := []string{"go", "north"}
	println(here.validMoves[SOUTH])
	tests := []struct {
		testString string
		splitInput []string
		result     string
	}{
		{"", emptyArray, UNKNOWN},
		{"asdf", emptyArray, UNKNOWN},
		{"go", emptyArray, WHERE_TO_GO("go")},
		{"go", southArray, here.validMoves[SOUTH]}, // TODO: in the future this will actually respond with the movement response from the room?
		{"walk", emptyArray, WHERE_TO_GO("walk")},
		{"walk", northArray, there.validMoves[NORTH]}, // TODO: in the future this will actually respond with the movement response from the room?
	}

	for _, test := range tests {
		result := parseMovement(test.testString, test.splitInput, UNKNOWN)
		if result != test.result {
			t.Errorf("Parsing %s, %v failed, got: %s, want: %s.", test.testString, test.splitInput, result, test.result)
		}
	}
}

// func TestMovePlayer(t *testing.T) {
// 	// just for testing TODO: remove
// 	var here = &Entity{
// 		name:     "Here",
// 		desc:     "A nice place",
// 		location: [2]int{0, 0},
// 		validMoves: map[string]string{
// 			SOUTH: "You amble from here to there",
// 		},
// 	}

// 	// just for testing TODO: remove
// 	var there = &Entity{
// 		name:     "There",
// 		desc:     "An okay place, I guess",
// 		location: [2]int{0, 1},
// 		validMoves: map[string]string{
// 			NORTH: "You mobilize from there to here",
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
// 		{"", emptyArray, UNKNOWN},
// 		{"asdf", emptyArray, UNKNOWN},
// 		{"go", emptyArray, WHERE_TO_GO("go")},
// 		{"go", northArray, GOING}, // TODO: in the future this will actually respond with the movement response from the room?
// 		{"walk", emptyArray, WHERE_TO_GO("walk")},
// 		{"walk", northArray, GOING}, // TODO: in the future this will actually respond with the movement response from the room?
// 	}

// 	for _, test := range tests {
// 		result := parseMovement(test.testString, test.splitInput, UNKNOWN)
// 		if result != test.result {
// 			t.Errorf("Parsing %s, %v failed, got: %s, want: %s.", test.testString, test.splitInput, result, test.result)
// 		}
// 	}
// }
