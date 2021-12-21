package gameFileIO

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"textadventureengine/structs"
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
func WriteGameFileToJson(game *structs.Game) {
	var filepath = game.FilePath
	if filepath == "" {
		filepath = game.Title
	}
	var f = openFileHandle(filepath, os.Create)
	defer f.Close()

	// point encoder at file
	var encoder = json.NewEncoder(f)

	// write game
	encoder.Encode(game)
}

func ReadGameFileFromJson(filename string) (game *structs.Game) {
	var f = openFileHandle(filename, os.Open)
	defer f.Close()

	// point encoder at file
	var decoder = json.NewDecoder(f)

	// write map
	decoder.Decode(&game)

	return game
}
