package speech

type Tree struct {
	Root Event
}

// An Event should have either a single next speech Event in `Next`. Or
// it should have 2 or more Response structs in `Responses` with user
// selectable options leading to speech Events.
type Event struct {
	Next *Event
	Responses []*Response
	Speech string
}

type Response struct {
	Response string
	Next *Event
}