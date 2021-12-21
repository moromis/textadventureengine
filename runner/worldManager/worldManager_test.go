package worldManager

import (
	"testing"
	"textadventureengine/runner/constants"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func initTestWorldInstance(noStartingRoom bool) {
	worldInstance = nil
	startingRoom := testObjects.TestGame.StartingRoom
	if noStartingRoom {
		startingRoom = nil
	}
	InitWorld(testObjects.TestGame.WorldLayout, testObjects.TestGame.WorldWidth, startingRoom)
}

func TestInitMap(t *testing.T) {
	assert.Nil(t, worldInstance)
	initTestWorldInstance(false)
	assert.NotNil(t, worldInstance)
}

func TestGetWidth(t *testing.T) {
	initTestWorldInstance(false)
	w := worldInstance.GetWidth()
	assert.Equal(t, w, testObjects.TestGame.WorldWidth)
}

func TestGetWorldLayout(t *testing.T) {
	initTestWorldInstance(false)
	l := worldInstance.GetWorldLayout()
	assert.Equal(t, l, testObjects.TestGame.WorldLayout)
}

func TestGetRoomString(t *testing.T) {
	initTestWorldInstance(true)
	s := worldInstance.PrintRoom(true)
	assert.Equal(t, constants.NO_ROOMS, s)
	initTestWorldInstance(false)
	s = worldInstance.PrintRoom(true)
	assert.Contains(t, s, testObjects.TestGame.StartingRoom.Desc)
}
