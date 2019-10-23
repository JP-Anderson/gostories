package features

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_features := loadFromXML()
	assert.Equal(t, 3, len(_features))
	shelf := _features["shelf"].GetThing()
	assert.Equal(t, "reveal-item(collar)", shelf.Triggers["look"])
}
