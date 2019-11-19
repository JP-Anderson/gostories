package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIDString(t *testing.T) {
	
	t.Run("no spaces", func (t *testing.T) {
		assert.Equal(t, "idnospaces",ToIDString( "idnospaces"))
	})

	t.Run("capital letters", func (t *testing.T) {
		assert.Equal(t, "idwithcapitals", ToIDString("IDWithCapitals"))
	})

	t.Run("capitals and spaces", func (t *testing.T) {
		assert.Equal(t, "idwithcapsandspaces", ToIDString("ID with caps AND spaces"))
	})

	t.Run("leading and trailing spaces", func (t *testing.T) {
		assert.Equal(t, "id", ToIDString("              id  "))
	})
}
