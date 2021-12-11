package main

// commands are maps of the command verb to ints 1 or greater
// kevin Dec. 10 2021: could potentially do something more useful than a nondescript int

var MOVEMENT_COMMANDS = map[string]int{"go": 1, "walk": 2}

var INSPECT_COMMANDS = map[string]int{"l": 1, "look": 2, "inspect": 3}

var INVENTORY_COMMANDS = map[string]int{"i": 1, "inventory": 2}

var SAVE_COMMANDS = []string{"s", "save"}
