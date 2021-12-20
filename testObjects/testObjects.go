package testObjects

import "textadventureengine/structs"

// ROOMS
var Here = &structs.Entity{
	Name:     "Here",
	Desc:     "A nice place",
	Location: [2]int{0, 0},
	ValidMoves: map[string]string{
		structs.SOUTH: "You amble from here to there",
	},
}
var There = &structs.Entity{
	Name:     "There",
	Desc:     "An okay place, I guess",
	Location: [2]int{1, 0},
	ValidMoves: map[string]string{
		structs.NORTH: "You mobilize from there to here",
	},
}

// ENTITIES
var RandomMan = &structs.Entity{
	Name:     "A Man",
	Desc:     "He looks a little sus, not gonna lie...",
	Location: There.Location,
}
var TestEntities = []*structs.Entity{
	RandomMan,
}

// MAP
var TestMapWidth = 1
var TestMap = []*structs.Entity{
	Here,
	There,
}

// ITEMS
var Ax = &structs.Entity{
	Name: "Ax",
	Desc: "An ax",
}
var Bow = &structs.Entity{
	Name: "Bow",
	Desc: "A bow",
}

// INVENTORY
var TestInventory = []*structs.Entity{
	Ax,
}
