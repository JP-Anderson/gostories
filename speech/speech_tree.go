package speech

type Tree struct {
	Event Event `xml:"Event"`
}

// An Event should have either a single next speech Event in `Next`. Or
// it should have 2 or more Response structs in `Responses` with user
// selectable options leading to speech Events.
type Event struct {
	Next      *Event `xml:"Next"`
	Responses []Response `xml:"Responses>Response"`
	Speech    string `xml:"Speech"`
}

type Response struct {
	ResponseStr string `xml:"ResponseStr"`
	Next        Event `xml:"Event"`
}
