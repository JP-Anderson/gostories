package things

// A Being is a character of some species which typically be interacted with through Speech
type Being interface {

	// Activating Look on a being will give a detailed description of the being. Some beings don't like being looked at.
	Look()

	// Activating Speak on a being will make the player (attempt to) initiate a conversation between the being.
	Speak()
}
