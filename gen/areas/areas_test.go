package areas

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	console "gostories/engine/io/console"
	"gostories/things/area"
)

func TestLoadFromXML(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	assert.Equal(t, 5, len(_areas))
	assert.Equal(
		t,
		"You are in some kind of stockroom. There is one shelf stacked high against one wall, across from the entrance. There's an exit to the North, and East.",
		_areas["store_room"].Look,
	)
}

func TestGet(t *testing.T) {
	
	t.Run("lower case name", func (t *testing.T) {
		room := Get("cat_room")
		assert.NotNil(t, room)
		assert.True(t, strings.Contains(room.Look, "cat"))
	})

	t.Run("upper case name", func (t *testing.T) {
		room := Get("Store_Room")
		assert.NotNil(t, room)
		assert.True(t, strings.Contains(room.Look, "stockroom"))
	})

}

func TestLoadLookTexts(t *testing.T) {
	room := Get("table_room")
	assert.Equal(
		t,
		"You are in a small room which is empty apart from a small wooden stand in the corner across from the only exit, back to the south.",
		room.LookText(),
	)
	assert.Len(t, room.LookTexts, 2)
}

func TestLoadLookText(t *testing.T) {
	room := Get("cat_room")
	assert.Equal(
		t,
		"You are in a small room, which is totally empty apart from a fat ginger cat, and a door to the west.",
		room.LookText(),
	)
	assert.Nil(t, room.LookTexts)
}

func TestLoadFromXMLLoadsBeings(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	catRoom := _areas["cat_room"]
	assert.Equal(t, 1, len(catRoom.Beings))

	bubbles := catRoom.Beings[0]
	assert.Equal(t, "Bubbles", bubbles.Name)
}

func TestLoadFromXMLLoadsFeatures(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_areas := loadFromXML()
	storeRoom := _areas["store_room"]
	assert.Equal(t, 1, len(storeRoom.Features))
	catRoom := _areas["cat_room"]
	assert.Equal(t, 0, len(catRoom.Features))
	kitchen := _areas["kitchen"]
	assert.Equal(t, 1, len(kitchen.Features))
	fridge := kitchen.Features[0]
	assert.Equal(t, "fridge", fridge.GetName())
	assert.Equal(t, "The fridge is empty apart from a tin of sardines.", fridge.GetLookText())
}

func TestLoadFromXMLLoadsExits(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
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
	io.Handler = console.NewConsoleInputOutputHandler()
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
