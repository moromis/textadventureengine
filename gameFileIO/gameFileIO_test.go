package gameFileIO

import (
	"testing"

	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func TestJsonIO(t *testing.T) {
	WriteGameFileToJson(testObjects.TestGame)
	var gameFromDisk = ReadGameFileFromJson(testObjects.TestGame.Title)

	assert.Equal(t, testObjects.TestGame.MapWidth, gameFromDisk.MapWidth)
	assert.Equal(t, testObjects.TestGame.MapLayout, gameFromDisk.MapLayout)
	assert.Equal(t, testObjects.TestGame.StartingRoom, gameFromDisk.StartingRoom)
	assert.Equal(t, testObjects.TestGame.Inventory, gameFromDisk.Inventory)
}
