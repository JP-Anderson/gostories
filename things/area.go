package things

// An Area is a place in the world which contains characters, items, and routes to/from other areas.
type Area interface {

	// Activating Look on an area will usually give a brief overview of items of interest in the area
	Look()

}
