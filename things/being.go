package things

type Being struct {
	Name    string
	Species string
	// Speech options represented as a list of strings for now. Will become a tree with options.
	Speech  []string
}