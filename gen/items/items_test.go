package items

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFromList(t *testing.T) {
	items := buildItemsMap()
	assert.Equal(t, 3, len(items))
}

func TestBuildFromMap(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_items := loadFromXML()
	assert.Equal(t, 3, len(_items))
}
