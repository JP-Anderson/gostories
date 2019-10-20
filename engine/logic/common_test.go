package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFuncName(t *testing.T) {
	input := "talk(your_mum)"
	assert.Equal(t, "your_mum", parseFuncParam(input))
}

func TestParseFuncParam(t *testing.T) {
	input := "talk(your_mum)"
	assert.Equal(t, "talk", parseFuncName(input))
}
