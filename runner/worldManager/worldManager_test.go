package worldManager

import (
	"testing"
	"textadventureengine/runner/constants"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func initTestWorldInstance(noStartingRoom bool) {
	worldManagerInstance = nil
	startingRoom := testObjects.TestGame.StartingRoom
	if noStartingRoom {
		startingRoom = nil
	}
	InitWorldManager(testObjects.TestGame.WorldLayout, testObjects.TestGame.WorldWidth, startingRoom)
}

func TestInitMap(t *testing.T) {
	assert.Nil(t, worldManagerInstance)
	initTestWorldInstance(false)
	assert.NotNil(t, worldManagerInstance)
}

func TestGetWidth(t *testing.T) {
	initTestWorldInstance(false)
	w := worldManagerInstance.GetWidth()
	assert.Equal(t, w, testObjects.TestGame.WorldWidth)
}

func TestGetWorldLayout(t *testing.T) {
	initTestWorldInstance(false)
	l := worldManagerInstance.GetWorldLayout()
	assert.Equal(t, l, testObjects.TestGame.WorldLayout)
}

func TestGetRoomString(t *testing.T) {
	initTestWorldInstance(true)
	s := worldManagerInstance.PrintRoom(true)
	assert.Equal(t, constants.NO_ROOMS, s)
	initTestWorldInstance(false)
	s = worldManagerInstance.PrintRoom(true)
	assert.Contains(t, s, testObjects.TestGame.StartingRoom.Desc)
}
