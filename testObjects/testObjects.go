package testObjects

import "tae.com/constants"

// ROOMS
var Here = &constants.Entity{
	Name:     "Here",
	Desc:     "A nice place",
	Location: [2]int{0, 0},
	ValidMoves: map[string]string{
		constants.SOUTH: "You amble from here to there",
	},
}
var There = &constants.Entity{
	Name:     "There",
	Desc:     "An okay place, I guess",
	Location: [2]int{1, 0},
	ValidMoves: map[string]string{
		constants.NORTH: "You mobilize from there to here",
	},
}

// ENTITIES
var RandomMan = &constants.Entity{
	Name:     "A Man",
	Desc:     "He looks a little sus, not gonna lie...",
	Location: There.Location,
}
var TestEntities = []*constants.Entity{
	RandomMan,
}

// MAP
var TestMapWidth = 1
var TestMap = []*constants.Entity{
	Here,
	There,
}

// ITEMS
var Ax = &constants.Entity{
	Name: "Ax",
	Desc: "An ax",
}
var Bow = &constants.Entity{
	Name: "Bow",
	Desc: "A bow",
}

// INVENTORY
var TestInventory = []*constants.Entity{
	Ax,
}
