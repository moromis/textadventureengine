package inventoryManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tae.com/constants"
)

// TODO: describe this test
func TestInventory(t *testing.T) {
	var ax = &constants.Entity{
		Name: "Ax",
		Desc: "An ax",
	}
	bow := &constants.Entity{
		Name: "Bow",
		Desc: "A bow",
	}
	stuff := []*constants.Entity{
		ax,
	}
	code, err := InitInventory(stuff, -1)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	code, err = InitInventory(stuff, 0)
	if code != -1 && err != nil {
		t.Errorf("bad code (failure) should have been received (-1), but got %d", code)
	}
	InitInventory(stuff, 100)
	inv := inventoryInstance.getInventory()
	if len(inv) != 1 {
		t.Errorf("size of inventory %d differed from expected length of 1", len(inv))
	}
	if inv[0] != ax {
		t.Errorf("inventory %v differed from expected inventory of %v", inv, stuff)
	}
	success := inventoryInstance.addToInventory(bow)
	assert.Equal(t, 0, success)
	inv = inventoryInstance.getInventory()
	assert.Len(t, inv, 2)
	ejectedItems := inventoryInstance.setInventoryLimit(10)
	assert.Equal(t, ejectedItems, []*constants.Entity(nil))
	ejectedItems = inventoryInstance.setInventoryLimit(1)
	assert.Len(t, ejectedItems, 1)
	assert.Equal(t, ejectedItems[0], bow)
	inv = inventoryInstance.getInventory()
	assert.Len(t, inv, 1)
	assert.Equal(t, inv[0], ax)
	success = inventoryInstance.addToInventory(bow)
	assert.Equal(t, -1, success)
}
