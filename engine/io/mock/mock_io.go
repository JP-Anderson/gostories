package mockio

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/parser"
)

// MockInputOutputHandler mocks the InputOutputHandler for the purposes of asserting data that
// has been Output from the engine.
type MockInputOutputHandler struct {
	output []string
}

// ExpectedStringEqualsNthOutputString returns true if the string output value created on the Nth call
// to NewLine or NewLinef equals the provided expected string.
func (c *MockInputOutputHandler) ExpectedStringEqualsNthOutputString(t *testing.T, expected string, n int) {
	assert.Equal(t, expected, c.output[n-1])
}

// NewMockInputOutputHandler creates a new MockInputOutputHandler for use in testing.
func NewMockInputOutputHandler() *MockInputOutputHandler {
	return &MockInputOutputHandler{}
}

// NewLine takes a string and adds it to the internal output tracker for assertions.
func (c *MockInputOutputHandler) NewLine(newOutput string) error {
	c.output = append(c.output, newOutput)
	return nil
}

// NewLinef takes a string and adds it to the internal output tracker for assertions.
func (c *MockInputOutputHandler) NewLinef(output string, args ...interface{}) error {
	return c.NewLine(fmt.Sprintf(output, args...))
}

// ReadInt isn't yet used but is needed for the interface
func (c *MockInputOutputHandler) ReadInt() (i int, e error) {
	// TODO: provide a method for pre-loading ReadInt with a series of int/error returns
	return -1, nil
}

// SimpleParse isn't yet used
func (c *MockInputOutputHandler) SimpleParse() (parser.Action, string) {
	// TODO: provide a method for pre-loading with a series of Action/string returns
	return parser.Unknown(), ""
}
