package stateMachine

import (
	"strings"

	"textadventureengine/helpers"
	"textadventureengine/runner/constants"
	"textadventureengine/runner/inventoryManager"
	"textadventureengine/runner/worldManager"
	"textadventureengine/structs"
)

// TODO: do these belong here, and should they be global?
const VERBOSE = false // @global -- user defined, settings

func SetupStateMachine(game *structs.Game) {
	inventoryManager.InitInventoryManager(game.Inventory, 100) // TODO: read limit from preferences/file
	worldManager.InitWorldManager(game.WorldLayout, game.WorldWidth, game.StartingRoom)
}

// TODO: allow for arrays of responses, and randomly select if the type is an array
// TODO: research static (post-compile) type-checking in code
func ParseInput(input string) string {
	lowerInput := strings.ToLower(input)
	splitInput := strings.Fields(lowerInput)
	var output = helpers.PickStringRandomly(constants.UNKNOWN)
	if len(splitInput) == 0 {
		return output // TODO: should just assign and then return at the end of this whole thing
	}
	verb := splitInput[0]

	output = parseInventory(verb, output)
	output = parseMovement(verb, splitInput, output)
	output = parseLook(verb, splitInput, output)
	return output
}

func parseMovement(verb string, splitInput []string, output string) string {
	// check if the command is a movement command
	validCommand := constants.MOVEMENT_COMMANDS[verb]
	if validCommand > 0 {
		// check if there's anything following the movement command
		if len(splitInput) > 1 {
			// get the move direction
			direction := splitInput[1]
			// move the player
			return movePlayer(direction)
		}
		// if there's nothing to do, query the user -- incomplete command
		return constants.WHERE_TO_GO(verb)
	}
	return output
}

func movePlayer(direction string) string {
	worldInstance := worldManager.GetWorldManager()
	validMoves := worldInstance.GetCurrentRoom().ValidMoves
	if validMoves[direction] != "" {
		// get the cardinal direction move array [colMove, rowMove]
		movement := structs.CARDINAL_DIRECTIONS[direction]
		// if the movement is not malformed/exists
		if len(movement) == 2 {
			// ensure the move is possible
			if worldInstance.CanMove(movement) {
				// store move
				moveDesc := validMoves[direction]
				if VERBOSE {
					moveDesc += "\n\n" + worldInstance.PrintRoom(false)
				}
				// move
				worldInstance.Move(movement)
				// return the description of the move
				return moveDesc
			}
		} else {
			return constants.HOW_TO_GO
		}
	}
	return constants.HOW_TO_GO
}

func parseLook(verb string, splitInput []string, output string) string {
	worldInstance := worldManager.GetWorldManager()
	inventory := inventoryManager.GetInventoryManager()
	// check if the command is a movement command
	validCommand := constants.INSPECT_COMMANDS[verb]
	if validCommand > 0 {
		if len(splitInput) > 1 {
			maybeThing := splitInput[1]
			// look through our inventory first
			success, message := inventory.InspectInventory(maybeThing)
			if success == 0 {
				return message
			}
			// TODO: then look through the room, only looking at visible things (direct descendants [stuff] of the room?)

			// just look around -- unhandled in other words
			return constants.LOOK_PLACEHOLDER
		}
		return worldInstance.PrintRoom(true)
	}
	return output
}

func parseInventory(verb string, output string) string {
	// check if the command is a movement command
	validCommand := constants.INVENTORY_COMMANDS[verb]
	if validCommand > 0 {
		inventory := inventoryManager.GetInventoryManager()
		return inventory.PrintInventory()
	}
	return output
}
