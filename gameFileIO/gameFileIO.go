package gameFileIO

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"textadventureengine/constants"
)

func openFileHandle(filename string, fileMethod func(string) (*os.File, error)) *os.File {
	// open file handle
	var _filename = filename
	if !strings.Contains(_filename, ".tae") {
		_filename = _filename + ".tae"
	}
	var f *os.File = nil
	f, err := fileMethod(_filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// JSON
func WriteGameFileToJson(filename string, mapLayout []*constants.Entity, mapWidth int, startingRoom *constants.Entity, entities []*constants.Entity, inventory []*constants.Entity) {
	var f = openFileHandle(filename, os.Create)
	defer f.Close()

	// point encoder at file
	var encoder = json.NewEncoder(f)

	// write map
	encoder.Encode(mapLayout)
	encoder.Encode(mapWidth)
	encoder.Encode(startingRoom)

	// write entities
	encoder.Encode(entities)

	// write inventory
	encoder.Encode(inventory)
}

func ReadGameFileFromJson(filename string) (mapLayout []*constants.Entity, mapWidth int, startingRoom *constants.Entity, entities []*constants.Entity, inventory []*constants.Entity) {
	var f = openFileHandle(filename, os.Open)
	defer f.Close()

	// point encoder at file
	var decoder = json.NewDecoder(f)

	// set up what we intend to read out of the file
	// TODO: can we spec this so if we change it it changes for writing as well?
	mapWidth = 1

	// write map
	decoder.Decode(&mapLayout)
	decoder.Decode(&mapWidth)
	decoder.Decode(&startingRoom)
	decoder.Decode(&entities)
	decoder.Decode(&inventory)

	return mapLayout, mapWidth, startingRoom, entities, inventory
}
