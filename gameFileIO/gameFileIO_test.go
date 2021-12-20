package gameFileIO

import (
	"testing"

	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

func TestJsonIO(t *testing.T) {
	var testStartingRoom = testObjects.Here
	WriteGameFileToJson("test_data", testObjects.TestMap, testObjects.TestMapWidth, testStartingRoom, testObjects.TestEntities, testObjects.TestInventory)
	var mapLayout, mapWidth, startingRoom, entities, inventory = ReadGameFileFromJson("test_data")

	assert.Equal(t, testObjects.TestMapWidth, mapWidth)
	assert.Equal(t, testObjects.TestMap, mapLayout)
	assert.Equal(t, testStartingRoom, startingRoom)
	assert.Equal(t, testObjects.TestEntities, entities)
	assert.Equal(t, testObjects.TestInventory, inventory)
}
