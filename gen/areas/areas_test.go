package areas

import (
	"fmt"
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"gostories/things/area"
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

func TestLoadFromXMLLoadsFeatures(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	storeRoom := _areas["store_room"]
	assert.Equal(t, 1, len(storeRoom.Features))
	catRoom := _areas["cat_room"]
	assert.Equal(t, 0, len(catRoom.Features))
	kitchen := _areas["kitchen"]
	assert.Equal(t, 1, len(kitchen.Features))
	println(fmt.Sprintf("%#v", kitchen))
	fridge := kitchen.Features[0]
	assert.Equal(t, "fridge", fridge.GetName())
	assert.Equal(t, "The fridge is empty apart from a tin of sardines.", fridge.GetLookText())
}

func TestLoadFromXMLLoadsExits(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	storeRoom := _areas["store_room"]
	catRoom := _areas["cat_room"]
	kitchen := _areas["kitchen"]

	// Assert cat room exit
	assert.Equal(t, catRoom.Exits[area.West].To, storeRoom)
	assert.Equal(t, catRoom.Exits[area.West].From, catRoom)

	// Assert kitchen exit
	assert.Equal(t, kitchen.Exits[area.South].To, storeRoom)
	assert.Equal(t, kitchen.Exits[area.South].From, kitchen)

	// Assert store room exits
	assert.Equal(t, storeRoom.Exits[area.North].To, kitchen)
	assert.Equal(t, storeRoom.Exits[area.North].From, storeRoom)

	assert.Equal(t, storeRoom.Exits[area.East].To, catRoom)
	assert.Equal(t, storeRoom.Exits[area.East].From, storeRoom)
}

func TestLoadFromXMLLoadsItems(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	storeRoom := _areas["store_room"]

	collar, err := storeRoom.Items.GetItemWithName("collar")
	c := *collar
	// Assert store room has collar
	assert.NoError(t, err)
	assert.NotNil(t, collar)
	assert.False(t, c.GetThing().Visible)
	assert.Equal(t, "collar", c.GetName())
}
