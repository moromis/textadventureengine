package worldManager

import (
	"fmt"
	"strings"
	"textadventureengine/runner/constants"
	"textadventureengine/structs"
)

type WorldManager struct {
	GetWidth       func() int
	GetWorldLayout func() []*structs.Entity
	GetWorldTable  func() [][]string
	GetCurrentRoom func() *structs.Entity
	PrintRoom      func(full bool) string
	CanMove        func(move [2]int) bool
	Move           func(move [2]int)
}

var worldManagerInstance *WorldManager = nil

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

func GetWorldManager() *WorldManager {
	return worldManagerInstance
}

// TODO: put all map management, map state, current room, etc. here
func InitWorldManager(worldLayout []*structs.Entity, worldWidth int, startingRoom *structs.Entity) {
	var currentRoom = startingRoom
	getWidth := func() int {
		return worldWidth
	}
	getWorldLayout := func() []*structs.Entity {
		return worldLayout
	}
	getWorldTable := func() [][]string {
		stringified := []string{}
		for _, room := range worldLayout {
			stringified = append(stringified, room.Name)
		}
		ret := [][]string{}
		for i := 0; i < len(stringified); i += worldWidth {
			ret = append(ret, stringified[i:i+worldWidth])
		}
		return ret
	}
	printRoom := func(full bool) string {
		if currentRoom == nil {
			return constants.NO_ROOMS
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
		currentRoom = worldLayout[colMove+rowMove*worldWidth]
	}

	// singleton pattern
	if worldManagerInstance == nil {
		worldManagerInstance = &WorldManager{
			getWidth,
			getWorldLayout,
			getWorldTable,
			getCurrentRoom,
			printRoom,
			canMove,
			move,
		}
	}
}
