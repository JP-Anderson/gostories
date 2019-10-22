package beings

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_beings := loadFromXML()
	assert.Equal(t, 1, len(_beings))
	//	println(fmt.Sprintf("%#v", _beings))
	assert.Equal(
		t,
		"The cat is reasonably small, ginger, and chunky.",
		_beings["bubbles"].LookText,
	)
}
