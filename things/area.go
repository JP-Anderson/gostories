package things

type Area struct {
	Look   string
	Exits  map[Direction]Exit
	Items []Item
	Beings []Being
}

type Exit struct {
	To   Area
	From Area
}

type Direction string
const (
	North = "north"
	East = "east"
	South = "south"
	West = "west"
)