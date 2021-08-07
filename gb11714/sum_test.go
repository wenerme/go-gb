package gb11714_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wenerme/go-gb/gb11714"
)

func TestSum(t *testing.T) {
	{
		s, err := gb11714.Sum("D2143569")
		assert.NoError(t, err)
		assert.Equal(t, "X", s)
	}
	{
		s, err := gb11714.Sum("MA1FL0LY")
		assert.NoError(t, err)
		assert.Equal(t, "9", s)
	}
	assert.True(t, gb11714.IsValid("MA1FL0LY9"))
}
