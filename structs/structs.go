package structs

type Game struct {
	FilePath     string
	Title        string
	MapLayout    []*Entity
	MapWidth     int
	StartingRoom *Entity
	Inventory    []*Entity // TODO: should/can inventory just be `stuff` in a player Entity? (maybe with EntityType == PLAYER_TYPE)
}

const (
	ROOM_TYPE int = iota
	ANIMATE_TYPE
	INANIMATE_TYPE
)

type Entity struct {
	Name       string            // just an identifier
	Desc       string            // the description of the entity -- what if this changes based on triggers or items -- maybe have a list of strings? Or an object/enum?
	Stuff      []Entity          // stuff that the entity has -- an inventory of sorts
	Location   [2]int            // where is this entity currently located
	ValidMoves map[string]string // valid movement commands (TODO: should this only accept cardinal directions), mapped to movement responses
	EntityType int               // what type of entity this is, ROOM_TYPE, ANIMATE_TYPE, or INANIMATE_TYPE
}

// TODO: what about up, down, around, into, etc?
// Kevin Dec. 10 2021: up/down could go to a new 2d map,
// other movement ideas could be dealt with in a one-off way or
// just generally not treated like actual map movement -- or you could code it?
// i.e. "go around dragon" is actually just transformed into "go north"
var CARDINAL_DIRECTIONS = map[string][2]int{
	"north":     {-1, 0},
	"northeast": {-1, 1},
	"east":      {0, 1},
	"southeast": {1, 1},
	"south":     {1, 0},
	"southwest": {1, -1},
	"west":      {0, -1},
	"northwest": {-1, -1},
}

const (
	NORTH     string = "north"
	SOUTH     string = "south"
	EAST      string = "east"
	WEST      string = "west"
	NORTHEAST string = "northeast"
	NORTHWEST string = "northwest"
	SOUTHEAST string = "southeast"
	SOUTHWEST string = "southwest"
)
