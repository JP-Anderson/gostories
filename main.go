package main

import (
	"gostories/engine"
	"gostories/gen/areas"
)

func main() {
	startRoom := areas.Get("cat_room")
	stage := engine.Stage{}
	stage.Start(*startRoom)
}
