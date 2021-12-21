package mapManager

import (
	"fmt"
	"strings"
	"textadventureengine/structs"
)

type MapInstance struct {
	GetWidth       func() int
	GetMapLayout   func() []*structs.Entity
	GetMapTable    func() [][]string
	GetCurrentRoom func() *structs.Entity
	PrintRoom      func(full bool) string
	CanMove        func(move [2]int) bool
	Move           func(move [2]int)
}

var mapInstance *MapInstance = nil

func getValidMovesString(room *structs.Entity) string {
	ret := "\nFrom here, you can go:\n"
	i := 0
	for move := range room.ValidMoves {
		ret += "- " + move
		if i < len(room.ValidMoves)-1 {
			ret += "\n"
		}
		i++
	}
	return ret
}

func GetMap() *MapInstance {
	return mapInstance
}

// TODO: put all map management, map state, current room, etc. here
func InitMapInstance(mapLayout []*structs.Entity, mapWidth int, startingRoom *structs.Entity) {
	var currentRoom = startingRoom
	getWidth := func() int {
		return mapWidth
	}
	getMapLayout := func() []*structs.Entity {
		return mapLayout
	}
	getMapTable := func() [][]string {
		stringified := []string{}
		for _, room := range mapLayout {
			stringified = append(stringified, room.Name)
		}
		ret := [][]string{}
		for i := 0; i < len(stringified); i += mapWidth {
			ret = append(ret, stringified[i:i+mapWidth])
		}
		return ret
	}
	getRoomString := func(full bool) string {
		if currentRoom == nil {
			return "You are floating in a void... There is nothing... (maybe create some rooms?)"
		}
		details := ""
		if full {
			details = getValidMovesString(currentRoom)
		}
		return fmt.Sprintf("%s\n%s\n%s\n", strings.ToUpper(currentRoom.Name), currentRoom.Desc, details) // TODO: in the future we'll want to have varying levels of verbosity
	}
	getCurrentRoom := func() *structs.Entity {
		return currentRoom
	}
	canMove := func(move [2]int) bool {
		potentialColChange := currentRoom.Location[0] + move[0]
		potentialRowChange := currentRoom.Location[1] + move[1]
		return potentialRowChange >= 0 && potentialColChange >= 0
	}
	move := func(move [2]int) {
		colMove := currentRoom.Location[0] + move[0]
		rowMove := currentRoom.Location[1] + move[1]
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
