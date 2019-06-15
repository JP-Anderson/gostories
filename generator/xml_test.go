package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeechFromXml(t *testing.T) {
	tree := SpeechFromXml(simpleResponseXML)
	assert.Equal(t, "What is your favourite colour?", tree.Event.Speech)
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
