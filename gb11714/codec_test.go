package gb11714

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	n, err := dec("MA1FL0LY9")
	assert.NoError(t, err)
	assert.Equal(t, "10", enc(36))
	assert.Equal(t, "MA1FL0LY9", enc(n))
	assert.Equal(t, "MA1FL0LYA", enc(n+1))
	assert.Equal(t, "MA1FL0LZ9", enc(n+36))
}
