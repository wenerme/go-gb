package gb11714_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wenerme/go-gb/gb11714"
)

func TestSeq(t *testing.T) {
	for _, test := range []struct {
		a string
		b string
	}{
		{a: "MA1FL0JH7", b: "MA1FL0JJ3"}, // skip I
		{a: "MA1FL0L19", b: "MA1FL0L27"},
	} {
		vb, err := gb11714.Next(test.a)
		assert.NoError(t, err)
		assert.Equal(t, test.b, vb)

		va, err := gb11714.Prev(vb)
		assert.NoError(t, err)
		assert.Equal(t, test.a, va)
	}
}
