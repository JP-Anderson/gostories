package features

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_features := loadFromXML()
	assert.Equal(t, 3, len(_features))
}

func TestGet(t *testing.T) {
	lowerCaseName := "stand"
	upperCaseName := "Fridge"

	stand := Get(lowerCaseName)
	assert.NotNil(t, stand)
	assert.Equal(t, "stand", stand.GetThing().Name)

	fridge := Get(upperCaseName)

	assert.NotNil(t, fridge)
	assert.Equal(t, "fridge", fridge.GetThing().Name)
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

func TestLoadNames(t *testing.T) {
	_features := loadFromXML()
	stand := _features["stand"].GetThing()
	assert.True(t, stand.MatchesName("stand"))
	assert.True(t, stand.MatchesName("table"))
}

func TestLoadPutActionTrigger(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_features := loadFromXML()
	stand := _features["stand"].GetThing()
	assert.NotNil(t, stand)
}
