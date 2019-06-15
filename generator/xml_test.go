package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	
	"gostories/speech"
)

func TestSpeechFromXMLFile(t *testing.T) {
	path := "./speech_data/test.xml"
	loadedTree := SpeechFromXMLFile(path)
	assert.Equal(t, getExpectedSimpleResponseTree(), loadedTree)
}

func TestSpeechFromXml(t *testing.T) {
	tree := SpeechFromXml(simpleResponseXML)
	assert.Equal(t, getExpectedSimpleResponseTree(), tree)
}

func getExpectedSimpleResponseTree() speech.Tree {
	return speech.Tree{
		Event: speech.Event{
			Speech: "What is your favourite colour?",
			Responses: []speech.Response{
				{
					ResponseStr: "Red",
					Next: speech.Event{
						Speech: "Hmm okay...",
					},
				},
				{
					ResponseStr: "Blue",
					Next: speech.Event{
						Speech: "That makes sense...",
					},
				},
			},
		},
	}
}

var simpleResponseXML = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Tree>
<Event>
	<Speech>What is your favourite colour?</Speech>
	<Responses>
		<Response>
			<ResponseStr>Red</ResponseStr>
			<Event>
				<Speech>Hmm okay...</Speech>
			</Event>
		</Response>
		<Response>
			<ResponseStr>Blue</ResponseStr>
			<Event>
				<Speech>That makes sense...</Speech>
			</Event>
		</Response>
	</Responses>
</Event>
</Tree>`)
