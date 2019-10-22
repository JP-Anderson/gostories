package beings

import (
<<<<<<< HEAD
=======
	"gostories/engine/io"
	console "gostories/engine/io/console"
>>>>>>> ebac53d7a94430573efd18872cdebc28e467cb71
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
<<<<<<< HEAD
	_beings := loadFromXML()
	assert.Equal(t, 1, len(_beings))
=======
	io.ActiveInputOutputHandler = console.NewConsoleInputOutputHandler()
	_beings := loadFromXML()
	assert.Equal(t, 1, len(_beings))
	//	println(fmt.Sprintf("%#v", _beings))
>>>>>>> ebac53d7a94430573efd18872cdebc28e467cb71
	assert.Equal(
		t,
		"The cat is reasonably small, ginger, and chunky.",
		_beings["bubbles"].LookText,
	)
}
