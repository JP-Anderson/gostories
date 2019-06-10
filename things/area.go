package things

type Area struct {
	Look   string
	Exits  []Exit
	Beings []Being
}

type Exit struct {
	To   Area
	From Area
}


