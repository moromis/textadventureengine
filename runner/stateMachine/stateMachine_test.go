package stateMachine

import (
	"testing"

	"textadventureengine/helpers"
	"textadventureengine/runner/constants"
	"textadventureengine/structs"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	SetupStateMachine(testObjects.TestMap, testObjects.TestMapWidth, testObjects.Here, testObjects.TestInventory)

	tests := []struct {
		testString string
		result     interface{}
	}{
		{"", constants.UNKNOWN},
		{"asdf", constants.UNKNOWN},
	}

	for _, test := range tests {
		result := ParseInput(test.testString)
		assert.Contains(t, test.result, result)
	}
}

func TestParseMovement(t *testing.T) {
	SetupStateMachine(testObjects.TestMap, testObjects.TestMapWidth, testObjects.Here, testObjects.TestInventory)

	emptyArray := []string{}
	southArray := []string{"go", "south"}
	northArray := []string{"go", "north"}
	tests := []struct {
		testString string
		splitInput []string
		result     interface{}
	}{
		{"", emptyArray, constants.UNKNOWN},
		{"asdf", emptyArray, constants.UNKNOWN},
		{"go", emptyArray, constants.WHERE_TO_GO("go")},
		{"go", southArray, testObjects.Here.ValidMoves[structs.SOUTH]}, // TODO: in the future this will actually respond with the movement response from the room?
		{"walk", emptyArray, constants.WHERE_TO_GO("walk")},
		{"walk", northArray, testObjects.There.ValidMoves[structs.NORTH]}, // TODO: in the future this will actually respond with the movement response from the room?
	}

	for _, test := range tests {
		result := parseMovement(test.testString, test.splitInput, helpers.PickStringRandomly(constants.UNKNOWN))
		assert.Contains(t, test.result, result)
	}
}
