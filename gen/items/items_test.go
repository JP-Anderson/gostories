package items

import (
	"gostories/engine/io"
	console "gostories/engine/io/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	io.Handler = console.NewConsoleInputOutputHandler()
	_items := loadFromXML()
	assert.Equal(t, 3, len(_items))
}

func TestGet(t *testing.T) {
	
	t.Run("lowercase name", func (t *testing.T) {
		shrubbery := Get("shrubbery")
		assert.NotNil(t, shrubbery)
		assert.Equal(t, "shrubbery", shrubbery.GetThing().Name)
	})

	t.Run("uppercase name", func (t *testing.T) {
		shrubbery := Get("Shrubbery")
		assert.NotNil(t, shrubbery)
		assert.Equal(t, "shrubbery", shrubbery.GetThing().Name)
	})

}
