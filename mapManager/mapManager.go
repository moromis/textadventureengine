package mapManager

import (
	"fmt"
	"strings"

	"textadventureengine/constants"
)

type MapInstance struct {
	GetWidth       func() int
	GetMapLayout   func() []*constants.Entity
	GetMapTable    func() [][]string
	GetCurrentRoom func() *constants.Entity
	PrintRoom      func(full bool) string
	CanMove        func(move [2]int) bool
	Move           func(move [2]int)
}

var mapInstance *MapInstance = nil

func getValidMovesString(room *constants.Entity) string {
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
func InitMapInstance(mapLayout []*constants.Entity, mapWidth int, startingRoom *constants.Entity) {
	var currentRoom = startingRoom
	getWidth := func() int {
		return mapWidth
	}
	getMapLayout := func() []*constants.Entity {
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
		details := ""
		if full {
			details = getValidMovesString(currentRoom)
		}
		return fmt.Sprintf("%s\n%s\n%s\n", strings.ToUpper(currentRoom.Name), currentRoom.Desc, details) // TODO: in the future we'll want to have varying levels of verbosity
	}
	getCurrentRoom := func() *constants.Entity {
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
