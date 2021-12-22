package inventoryManager

import (
	"strings"
	"testing"

	"textadventureengine/constants"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

// TODO: describe this test
func TestInitInventory(t *testing.T) {
	inventoryManagerInstance = nil
	code, err := InitInventoryManager(testObjects.TestInventory, -1)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	code, err = InitInventoryManager(testObjects.TestInventory, 0)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	InitInventoryManager(testObjects.TestInventory, 100)
	inv := inventoryManagerInstance.GetInventory()
	if len(inv) != 1 {
		t.Errorf("size of inventory %d differed from expected length of 1", len(inv))
	}
	if inv[0] != testObjects.Ax {
		t.Errorf("inventory %v differed from expected inventory of %v", inv, testObjects.TestInventory)
	}
}

func TestAddToInventory(t *testing.T) {
	inventoryManagerInstance = nil
	// setup with a test inventory of one item
	InitInventoryManager(testObjects.TestInventory, 2)
	// add an item to the inventory
	success := inventoryManagerInstance.AddToInventory(testObjects.Bow)
	// make sure the add was successful
	assert.Equal(t, 0, success)
	inv := inventoryManagerInstance.GetInventory()
	// the length of our inventory should now be 2
	assert.Len(t, inv, 2)
}

func TestSetInventoryLimit(t *testing.T) {
	inventoryManagerInstance = nil
	// setup the inventory with our test inventory
	InitInventoryManager(testObjects.TestInventory, 1)
	inventoryManagerInstance.SetInventory(testObjects.TestInventory)
	// set the inventory limit higher
	ejectedItems := inventoryManagerInstance.SetInventoryLimit(10)
	// nothing should be ejected
	assert.Equal(t, ejectedItems, []*constants.Entity(nil))
	// set the inventory limit to 0
	ejectedItems = inventoryManagerInstance.SetInventoryLimit(0)
	// the ejected items should be the whole inventory we put in
	assert.Len(t, ejectedItems, 1)
	assert.Equal(t, ejectedItems[0], testObjects.TestInventory[0])
	inv := inventoryManagerInstance.GetInventory()
	assert.Len(t, inv, 0)
	// now that we're at 0, we should also not be able to add any items
	success := inventoryManagerInstance.AddToInventory(testObjects.Bow)
	assert.Equal(t, -1, success)
}

func TestSetInventory(t *testing.T) {
	inventoryManagerInstance = nil
	InitInventoryManager(testObjects.TestInventory, 1)
	// test setting the inventory
	inventoryManagerInstance.SetInventory([]*constants.Entity{testObjects.Bow})
	inv := inventoryManagerInstance.GetInventory()
	assert.Len(t, inv, 1)
	assert.Equal(t, inv[0], testObjects.Bow)
}

func TestGetInventoryInstance(t *testing.T) {
	inventoryManagerInstance = nil
	// test getting the inventory instance
	// if we haven't set up the inventory, we should get nil
	i := GetInventoryManager()
	assert.Nil(t, i)

	// after we setup the inventory, we should be able to retrieve it with GetInventory
	// (i.e. it should not be nil anymore)
	InitInventoryManager(testObjects.TestInventory, 100)
	success := inventoryManagerInstance.AddToInventory(testObjects.Bow)
	if success >= 0 {
		i = GetInventoryManager()
		assert.NotNil(t, i)
	}
}

func TestInspectInventory(t *testing.T) {
	inventoryManagerInstance = nil
	InitInventoryManager(testObjects.TestInventory, 100)
	i := GetInventoryManager()
	// try to inspect a valid thing that's in the inventory
	err, output := i.InspectInventory(testObjects.TestInventory[0].Name)
	assert.Equal(t, 0, err)
	assert.Contains(t, output, strings.ToLower(testObjects.TestInventory[0].Name))
	// try to inspect something invalid
	err, output = i.InspectInventory("somethingthatdefinitelydoesn'texist")
	assert.Equal(t, -1, err)
	assert.Equal(t, len(output), 0)
}

func TestPrintInventory(t *testing.T) {
	inventoryManagerInstance = nil
	InitInventoryManager([]*constants.Entity{testObjects.Ax, testObjects.Bow}, 2)
	i := GetInventoryManager()
	output := i.PrintInventory()
	// make sure anything is output at all
	assert.Greater(t, len(output), 0)
}
