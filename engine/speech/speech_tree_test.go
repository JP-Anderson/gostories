package speech

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeechTreeNext(t *testing.T) {
	tree := buildTestTree()
	one := tree.Event
	two := one.Next
	assert.Equal(t, "Hello", one.Speech)
	assert.Equal(t, "How are you", two.Speech)
}

func TestSpeechTreeResponses(t *testing.T) {
	tree := buildTestTree()
	responseEvent := tree.Event.Next

	resp1 := responseEvent.Responses[0]
	assert.Equal(t, "Yes!", resp1.ResponseStr)
	assert.Equal(t, "Okay", resp1.Next.Speech)

	resp2 := responseEvent.Responses[1]
	assert.Equal(t, "That would be an ecumenical matter.", resp2.ResponseStr)
	assert.Equal(t, "Yes, I suppose it would...", resp2.Next.Speech)
}

func buildTestTree() Tree {
	root := Event{
		Speech: "Hello",
		Next: &Event{
			Speech: "How are you",
			Responses: []Response{
				{
					ResponseStr: "Yes!",
					Next:        Event{Speech: "Okay"},
				},
				{
					ResponseStr: "That would be an ecumenical matter.",
					Next:        Event{Speech: "Yes, I suppose it would..."},
				},
			},
		},
	}
	return Tree{Event: root}
}
