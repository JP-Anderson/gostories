package speech

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeechTreeNext(t *testing.T) {
	tree := buildTestTree()
	one := tree.Root
	two := one.Next
	assert.Equal(t, "Hello", one.Speech)
	assert.Equal(t, "How are you", two.Speech)
}

func TestSpeechTreeResponses(t *testing.T) {
	tree := buildTestTree()
	responseEvent := tree.Root.Next

	resp1 := responseEvent.Responses[0]
	assert.Equal(t, "Yes!", resp1.Response)
	assert.Equal(t, "Okay", resp1.Next.Speech)

	resp2 := responseEvent.Responses[1]
	assert.Equal(t, "That would be an ecumenical matter.", resp2.Response)
	assert.Equal(t, "Yes, I suppose it would...", resp2.Next.Speech)
}

func buildTestTree() Tree {
	root := Event{
		Speech: "Hello",
		Next: &Event{
			Speech: "How are you",
			Responses: []*Response{
				{
					Response: "Yes!",
					Next:     &Event{Speech: "Okay"},
				},
				{
					Response: "That would be an ecumenical matter.",
					Next:     &Event{Speech: "Yes, I suppose it would..."},
				},
			},
		},
	}
	return Tree{Root: root}
}
