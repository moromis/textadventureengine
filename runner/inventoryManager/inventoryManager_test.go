package inventoryManager

import (
	"strings"
	"testing"

	"textadventureengine/structs"
	"textadventureengine/testObjects"

	"github.com/stretchr/testify/assert"
)

// TODO: describe this test
func TestInitInventory(t *testing.T) {
	inventoryInstance = nil
	code, err := InitInventory(testObjects.TestInventory, -1)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	code, err = InitInventory(testObjects.TestInventory, 0)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	InitInventory(testObjects.TestInventory, 100)
	inv := inventoryInstance.GetInventory()
	if len(inv) != 1 {
		t.Errorf("size of inventory %d differed from expected length of 1", len(inv))
	}
	if inv[0] != testObjects.Ax {
		t.Errorf("inventory %v differed from expected inventory of %v", inv, testObjects.TestInventory)
	}
}

func TestAddToInventory(t *testing.T) {
	inventoryInstance = nil
	// setup with a test inventory of one item
	InitInventory(testObjects.TestInventory, 2)
	// add an item to the inventory
	success := inventoryInstance.AddToInventory(testObjects.Bow)
	// make sure the add was successful
	assert.Equal(t, 0, success)
	inv := inventoryInstance.GetInventory()
	// the length of our inventory should now be 2
	assert.Len(t, inv, 2)
}

func TestSetInventoryLimit(t *testing.T) {
	inventoryInstance = nil
	// setup the inventory with our test inventory
	InitInventory(testObjects.TestInventory, 1)
	inventoryInstance.SetInventory(testObjects.TestInventory)
	// set the inventory limit higher
	ejectedItems := inventoryInstance.SetInventoryLimit(10)
	// nothing should be ejected
	assert.Equal(t, ejectedItems, []*structs.Entity(nil))
	// set the inventory limit to 0
	ejectedItems = inventoryInstance.SetInventoryLimit(0)
	// the ejected items should be the whole inventory we put in
	assert.Len(t, ejectedItems, 1)
	assert.Equal(t, ejectedItems[0], testObjects.TestInventory[0])
	inv := inventoryInstance.GetInventory()
	assert.Len(t, inv, 0)
	// now that we're at 0, we should also not be able to add any items
	success := inventoryInstance.AddToInventory(testObjects.Bow)
	assert.Equal(t, -1, success)
}

func TestSetInventory(t *testing.T) {
	inventoryInstance = nil
	InitInventory(testObjects.TestInventory, 1)
	// test setting the inventory
	inventoryInstance.SetInventory([]*structs.Entity{testObjects.Bow})
	inv := inventoryInstance.GetInventory()
	assert.Len(t, inv, 1)
	assert.Equal(t, inv[0], testObjects.Bow)
}

func TestGetInventoryInstance(t *testing.T) {
	inventoryInstance = nil
	// test getting the inventory instance
	// if we haven't set up the inventory, we should get nil
	i := GetInventory()
	assert.Nil(t, i)

	// after we setup the inventory, we should be able to retrieve it with GetInventory
	// (i.e. it should not be nil anymore)
	InitInventory(testObjects.TestInventory, 100)
	success := inventoryInstance.AddToInventory(testObjects.Bow)
	if success >= 0 {
		i = GetInventory()
		assert.NotNil(t, i)
	}
}

func TestInspectInventory(t *testing.T) {
	inventoryInstance = nil
	InitInventory(testObjects.TestInventory, 100)
	i := GetInventory()
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
	inventoryInstance = nil
	InitInventory([]*structs.Entity{testObjects.Ax, testObjects.Bow}, 2)
	i := GetInventory()
	output := i.PrintInventory()
	// make sure anything is output at all
	assert.Greater(t, len(output), 0)
}
