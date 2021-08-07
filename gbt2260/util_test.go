package gbt2260

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithIn(t *testing.T) {
	assert.True(t, WithIn("101011", "101011"))
	assert.True(t, WithIn("101011", "101000"))
	assert.True(t, WithIn("101011", "100000"))
	assert.False(t, WithIn("101000", "101011"))
}

func TestLevel(t *testing.T) {
	assert.Equal(t, 3, Level("010203"))
	assert.Equal(t, 2, Level("010200"))
	assert.Equal(t, 1, Level("010000"))
}
