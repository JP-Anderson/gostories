package areas

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	assert.Equal(t, 3, len(_areas))
	assert.Equal(
		t,
		"You are in some kind of stockroom. There is one shelf stacked high against one wall, across from the entrance.",
		_areas["store_room"].Look,
	)
}

func TestLoadFromXMLLoadsBeings(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	assert.Equal(t, 3, len(_areas))
	catRoom := _areas["cat_room"]
	assert.Equal(t, 1, len(catRoom.Beings))
}
