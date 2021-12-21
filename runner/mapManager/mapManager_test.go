package mapManager

import (
	"testing"
	"textadventureengine/runner/constants"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func initTestMapInstance(noStartingRoom bool) {
	mapInstance = nil
	startingRoom := testObjects.TestGame.StartingRoom
	if noStartingRoom {
		startingRoom = nil
	}
	InitMapInstance(testObjects.TestGame.MapLayout, testObjects.TestGame.MapWidth, startingRoom)
}

func TestInitMap(t *testing.T) {
	assert.Nil(t, mapInstance)
	initTestMapInstance(false)
	assert.NotNil(t, mapInstance)
}

func TestGetWidth(t *testing.T) {
	initTestMapInstance(false)
	w := mapInstance.GetWidth()
	assert.Equal(t, w, testObjects.TestGame.MapWidth)
}

func TestGetMapLayout(t *testing.T) {
	initTestMapInstance(false)
	l := mapInstance.GetMapLayout()
	assert.Equal(t, l, testObjects.TestGame.MapLayout)
}

func TestGetRoomString(t *testing.T) {
	initTestMapInstance(true)
	s := mapInstance.PrintRoom(true)
	assert.Equal(t, constants.NO_ROOMS, s)
	initTestMapInstance(false)
	s = mapInstance.PrintRoom(true)
	assert.Contains(t, s, testObjects.TestGame.StartingRoom.Desc)
}
