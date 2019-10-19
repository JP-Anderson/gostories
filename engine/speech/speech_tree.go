package speech

// Tree represents a tree of conversation events and responses. This struct
// contains the root Event node for a single in-game conversation. The Tree
// should be loaded via XML and is currently passed into a Being on creation.
type Tree struct {
	Event      Event `xml:"Event"`
	checkpoint *Event
}

// Start fetches the conversational event the conversation will start from.
// If a past checkpoint has been set, it will start there, otherwise, it will
// start from the beginning.
func (t *Tree) Start() Event {
	if t.checkpoint == nil {
		return t.Event
	}
	return *t.checkpoint
}

// SetStart sets a new checkpoint event in the speech tree. Future conversations
// using this tree will start from this point.
func (t *Tree) SetStart(e *Event) {
	t.checkpoint = e
}

// Event represents a conversational utterance an NPC/Being in-game has made
// to the player. It contains the text of the utterance, as well as either a
// single next speech Event in `Next`, or two or more Response structs in the
// `Responses` attribute with user selectable options leading to further speech
// Events.
type Event struct {
	// optional
	Next *Event `xml:"Event"`
	// optional
	Responses  []Response `xml:"Responses>Response"`
	Speech     string     `xml:"Speech"`
	Condition  string     `xml:"Condition"`
	Trigger    string     `xml:"Trigger"`
	Checkpoint *Event     `xml:"Checkpoint"`
}

// Response represents a conversational response the player can choose in
// response to a speech Event. A Response contains the text of the response
// utterance, can lead to further SpeechEvents, and can also contain Triggers
// to lead to further actions. A Response can be hidden/displayed from the
// user depending on its Condition/func attribute.
type Response struct {
	ResponseStr string `xml:"ResponseStr"`
	Next        Event  `xml:"Event"`
	Condition   string `xml:"Condition"`
	Trigger     string `xml:"Trigger"`
}
