package beings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	_beings := loadFromXML()
	assert.Equal(t, 1, len(_beings))
	assert.Equal(
		t,
		"The cat is reasonably small, ginger, and chunky.",
		_beings["bubbles"].LookText,
	)
}
