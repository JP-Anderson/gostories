package beings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromXML(t *testing.T) {
	_beings := loadFromXML()
	assert.Equal(t, 2, len(_beings))
	bubbles := _beings["bubbles"]
	assert.Equal(
		t,
		"The cat is reasonably small, ginger, and chunky.",
		bubbles.LookText,
	)
	assert.Equal(
		t,
		"Meow. Meeeeeeeeew! Mew.",
		bubbles.AltSpeech.Event.Speech,
	)
	assert.Equal(
		t,
		"Good day! I don't suppose you have any food I could eat do you? I'm famished!",
		bubbles.Speech.Event.Speech,
	)
}

func TestGet(t *testing.T) {
	
	t.Run("lower case name", func (t *testing.T) {
		being := Get("bubbles")
		assert.NotNil(t, being)
		assert.Equal(t, "Bubbles", being.GetThing().Name)
	})

	t.Run("upper case name", func (t *testing.T) {
		being := Get("Bubbles")
		assert.NotNil(t, being)
		assert.Equal(t, "Bubbles", being.GetThing().Name)
	})

}
