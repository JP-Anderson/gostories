package mockio

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gostories/engine/io"
	"gostories/engine/parser"
)

// MockInputOutputHandler mocks the InputOutputHandler for the purposes of asserting data that
// has been Output from the engine.
type MockInputOutputHandler struct {
	output []string
	suppressOutput bool
}

// ExpectedStringEqualsNthOutputString returns true if the string output value created on the Nth call
// to NewLine or NewLinef equals the provided expected string.
func (c *MockInputOutputHandler) ExpectedStringEqualsNthOutputString(t *testing.T, expected string, n int) {
	assert.Equal(t, expected, c.output[n-1])
}

// NewMockHandler creates a MockHandler which also outputs to the console in tests.
func NewMockHandler() *MockInputOutputHandler {
	mock := &MockInputOutputHandler {
		suppressOutput: false,
	}
	io.Handler = mock
	return mock
}

// NewLine takes a string and adds it to the internal output tracker for assertions.
func (c *MockInputOutputHandler) NewLine(newOutput string) error {
	c.output = append(c.output, newOutput)
	if c.suppressOutput {
		return nil
	}
	println(newOutput)
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

// ReadIntInRange is implemented for the interface
func (c *MockInputOutputHandler) ReadIntInRange(x, y int) int { return -1 }

// SimpleParse isn't yet used
func (c *MockInputOutputHandler) SimpleParse() (parser.Action, []string) {
	// TODO: provide a method for pre-loading with a series of Action/string returns
	return parser.Unknown(), []string{""}
}
