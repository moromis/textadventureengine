package inventoryManager

import (
	"errors"
	"fmt"
	"strings"

	"textadventureengine/constants"
)

type Inventory struct {
	GetInventory      func() []*constants.Entity
	AddToInventory    func(thing *constants.Entity) int
	SetInventory      func(stuff []*constants.Entity)
	SetInventoryLimit func(newSize int) []*constants.Entity
	InspectInventory  func(itemName string) (int, string)
	PrintInventory    func() string
}

var inventoryInstance *Inventory = nil

func GetInventory() *Inventory {
	return inventoryInstance
}

func InitInventory(initialStuff []*constants.Entity, maxSize int) (int, error) {
	if maxSize < 0 {
		return -1, errors.New("max inventory size of less than 0 passed")
	}
	if len(initialStuff) > maxSize {
		return -1, errors.New("amount of initial inventory is larger than max inventory size passed")
	}

	var inventoryLimit = maxSize
	var inventory = initialStuff

	getInventory := func() []*constants.Entity {
		return inventory
	}
	addToInventory := func(thing *constants.Entity) int {
		if len(inventory) < inventoryLimit {
			inventory = append(inventory, thing)
			return 0
		} else {
			return -1
		}
	}
	setInventory := func(stuff []*constants.Entity) {
		inventory = stuff
	}
	setInventoryLimit := func(newSize int) []*constants.Entity {
		inventoryLimit = newSize
		if newSize < len(inventory) {
			ejectedItems := inventory[newSize:]
			inventory = inventory[:newSize]
			return ejectedItems
		}
		return nil
	}
	inspectInventory := func(itemName string) (int, string) {
		for _, thing := range inventory {
			thingName := strings.ToLower(thing.Name)
			if thingName == itemName {
				return 0, fmt.Sprintf("You examine the %s:\n%s", thingName, thing.Desc)
			}
		}
		return -1, ""
	}
	printInventory := func() string {
		ret := "Your inventory:\n"
		for index, item := range inventory {
			ret += "- " + item.Name
			if index < len(inventory)-1 {
				ret += "\n"
			}
		}
		return ret
	}

	// singleton pattern
	if inventoryInstance == nil {
		inventoryInstance = &Inventory{
			getInventory,
			addToInventory,
			setInventory,
			setInventoryLimit,
			inspectInventory,
			printInventory,
		}
	}

	return 0, nil
}
