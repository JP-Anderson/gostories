package main

import (
	"gostories/engine"
	"gostories/gen/areas"
)

func main() {
	startRoom := areas.Area("cat_room")
	stage := engine.Stage{}
	stage.Start(startRoom)
}
