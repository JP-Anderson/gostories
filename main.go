package main

import (
	"gostories/engine"
	"gostories/gen/areas"
)

func main() {
	startRoom := areas.Get("stasis_pod")
	stage := engine.Stage{}
	stage.Start(*startRoom)
}
