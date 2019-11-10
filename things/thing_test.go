package things

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesNameSingleName(t *testing.T) {
	thing := Thing{
		Names: []string{ "john" },
	}
	assert.True(t, thing.MatchesName("JOHN"))
}

func TestMatchesNameMultipleNames(t *testing.T) {
	thingWithTwoNames := Thing{
		Names: []string{ "Peter Parker","Spiderman" },
	}

	assert.True(t, thingWithTwoNames.MatchesName("peter parker"))
	assert.True(t, thingWithTwoNames.MatchesName("spiderman"))
}


