package inventoryManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tae.com/constants"
	"tae.com/testObjects"
)

// TODO: describe this test
func TestInventory(t *testing.T) {
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
	success := inventoryInstance.AddToInventory(testObjects.Bow)
	assert.Equal(t, 0, success)
	inv = inventoryInstance.GetInventory()
	assert.Len(t, inv, 2)
	ejectedItems := inventoryInstance.SetInventoryLimit(10)
	assert.Equal(t, ejectedItems, []*constants.Entity(nil))
	ejectedItems = inventoryInstance.SetInventoryLimit(1)
	assert.Len(t, ejectedItems, 1)
	assert.Equal(t, ejectedItems[0], testObjects.Bow)
	inv = inventoryInstance.GetInventory()
	assert.Len(t, inv, 1)
	assert.Equal(t, inv[0], testObjects.Ax)
	success = inventoryInstance.AddToInventory(testObjects.Bow)
	assert.Equal(t, -1, success)
}
