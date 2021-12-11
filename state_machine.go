package main

import (
	"strings"
)

// just for testing TODO: remove
var here = &Entity{
	name:     "Here",
	desc:     "A nice place",
	location: [2]int{0, 0},
	validMoves: map[string]string{
		SOUTH: "You amble from here to there",
	},
}

// just for testing TODO: remove
var there = &Entity{
	name:     "There",
	desc:     "An okay place, I guess",
	location: [2]int{1, 0},
	validMoves: map[string]string{
		NORTH: "You mobilize from there to here",
	},
}

// just for testing TODO: remove
var ax = &Entity{
	name: "Ax",
	desc: "An ax",
}

// just for testing TODO: remove
var bow = &Entity{
	name: "Bow",
	desc: "A bow",
}

// TODO: do these belong here, and should they be global?
const VERBOSE = false                  // @global -- user defined, settings
var mapLayout = []*Entity{here, there} // @global -- user defined
var mapWidth = 1                       // @global -- user defined
var startingRoom = here                // @global -- user defined
// var mapInstance = getMapInstance(mapLayout, mapWidth, startingRoom)

func setupStateMachine() {
	// TODO: read from file
	initInventory([]*Entity{ax, bow}, 100)
	initMapInstance(mapLayout, mapWidth, startingRoom)
}

// TODO: allow for arrays of responses, and randomly select if the type is an array
// TODO: research static (post-compile) type-checking in code
func parseInput(input string) string {
	lowerInput := strings.ToLower(input)
	splitInput := strings.Fields(lowerInput)
	var output string = UNKNOWN
	if len(splitInput) == 0 {
		return output // TODO: should just assign and then return at the end of this whole thing
	}
	verb := splitInput[0]

	println("verb: ", verb)

	output = parseInventory(verb, output)
	output = parseMovement(verb, splitInput, output)
	output = parseLook(verb, splitInput, output)
	return output
}

func parseMovement(verb string, splitInput []string, output string) string {
	// check if the command is a movement command
	validCommand := MOVEMENT_COMMANDS[verb]
	println(validCommand)
	if validCommand > 0 {
		// check if there's anything following the movement command
		if len(splitInput) > 1 {
			// get the move direction
			direction := splitInput[1]
			// move the player
			return movePlayer(direction)
		}
		// if there's nothing to do, query the user -- incomplete command
		return WHERE_TO_GO(verb)
	}
	return output
}

func movePlayer(direction string) string {
	validMoves := mapInstance.getCurrentRoom().validMoves
	if validMoves[direction] != "" {
		// get the cardinal direction move array [colMove, rowMove]
		movement := CARDINAL_DIRECTIONS[direction]
		// if the movement is not malformed/exists
		if len(movement) == 2 {
			// ensure the move is possible
			if mapInstance.canMove(movement) {
				// store move
				moveDesc := validMoves[direction]
				if VERBOSE {
					moveDesc += "\n\n" + mapInstance.printRoom(false)
				}
				// move
				mapInstance.move(movement)
				// return the description of the move
				return moveDesc
			}
		} else {
			return HOW_TO_GO
		}
	}
	return HOW_TO_GO
}

func parseLook(verb string, splitInput []string, output string) string {
	// check if the command is a movement command
	validCommand := INSPECT_COMMANDS[verb]
	if validCommand > 0 {
		if len(splitInput) > 1 {
			maybeThing := splitInput[1]
			// look through our inventory first
			success, message := inventoryInstance.inspectInventory(maybeThing)
			if success == 0 {
				return message
			}
			// TODO: then look through the room, only looking at visible things (direct descendants [stuff] of the room?)

			// just look around -- unhandled in other words
			return LOOK_PLACEHOLDER
		}
		return mapInstance.printRoom(true)
	}
	return output
}

func parseInventory(verb string, output string) string {
	// check if the command is a movement command
	validCommand := INVENTORY_COMMANDS[verb]
	if validCommand > 0 {
		return inventoryInstance.printInventory()
	}
	return output
}
