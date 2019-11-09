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
}

func TestLoadTriggerStringsWithAndWithoutTarget(t *testing.T) {
	_features := loadFromXML()
	stand := _features["stand"].GetThing()
	triggers := stand.Triggers
	putTrigger := triggers["put"]
	assert.Equal(t, "shrubbery", putTrigger.Target)
	assert.NotEqual(t, "", putTrigger.Action)

	fridge := features["fridge"].GetThing()
	fTriggers := fridge.Triggers
	lookTrigger := fTriggers["look"]
	assert.Equal(t, "", lookTrigger.Target)
}

func TestLoadPutActionTrigger(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_features := loadFromXML()
	stand := _features["stand"].GetThing()
	assert.NotNil(t, stand)
	// TODO: actually test the put action with shrubbery on stand.
}
