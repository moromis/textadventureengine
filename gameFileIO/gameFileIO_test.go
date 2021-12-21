package gameFileIO

import (
	"testing"

	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func TestJsonIO(t *testing.T) {
	WriteGameFileToJson(testObjects.TestGame)
	var gameFromDisk = ReadGameFileFromJson(testObjects.TestGame.Title)

	assert.Equal(t, testObjects.TestGame.WorldWidth, gameFromDisk.WorldWidth)
	assert.Equal(t, testObjects.TestGame.WorldLayout, gameFromDisk.WorldLayout)
	assert.Equal(t, testObjects.TestGame.StartingRoom, gameFromDisk.StartingRoom)
	assert.Equal(t, testObjects.TestGame.Inventory, gameFromDisk.Inventory)
}
