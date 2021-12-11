package main

import (
	"fmt"
	"strings"
)

type MapInstance struct {
	getWidth       func() int
	getMapLayout   func() []*Entity
	getMapTable    func() [][]string
	getCurrentRoom func() *Entity
	printRoom      func(full bool) string
	canMove        func(move [2]int) bool
	move           func(move [2]int)
}

var mapInstance *MapInstance = nil

func getValidMovesString(room *Entity) string {
	ret := "\nFrom here, you can go:\n"
	i := 0
	for move := range room.validMoves {
		ret += "- " + move
		if i < len(room.validMoves)-1 {
			ret += "\n"
		}
		i++
	}
	return ret
}

// TODO: put all map management, map state, current room, etc. here
func initMapInstance(mapLayout []*Entity, mapWidth int, startingRoom *Entity) {
	var currentRoom = startingRoom
	getWidth := func() int {
		return mapWidth
	}
	getMapLayout := func() []*Entity {
		return mapLayout
	}
	getMapTable := func() [][]string {
		stringified := []string{}
		for _, room := range mapLayout {
			stringified = append(stringified, room.name)
		}
		ret := [][]string{}
		for i := 0; i < len(stringified); i += mapWidth {
			ret = append(ret, stringified[i:i+mapWidth])
		}
		return ret
	}
	getRoomString := func(full bool) string {
		details := ""
		if full {
			details = getValidMovesString(currentRoom)
		}
		return fmt.Sprintf("%s\n%s\n%s\n", strings.ToUpper(currentRoom.name), currentRoom.desc, details) // TODO: in the future we'll want to have varying levels of verbosity
	}
	getCurrentRoom := func() *Entity {
		return currentRoom
	}
	canMove := func(move [2]int) bool {
		potentialColChange := currentRoom.location[0] + move[0]
		potentialRowChange := currentRoom.location[1] + move[1]
		return potentialRowChange >= 0 && potentialColChange >= 0
	}
	move := func(move [2]int) {
		colMove := currentRoom.location[0] + move[0]
		rowMove := currentRoom.location[1] + move[1]
		currentRoom = mapLayout[colMove+rowMove*mapWidth]
	}

	// singleton pattern
	if mapInstance == nil {
		mapInstance = &MapInstance{
			getWidth,
			getMapLayout,
			getMapTable,
			getCurrentRoom,
			getRoomString,
			canMove,
			move,
		}
	}
}
