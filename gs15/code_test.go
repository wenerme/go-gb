package gs15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert.Equal(t, Code{}.String(), "000000000000000")
	assert.Equal(t, Code{FirstRegNo: 1, Sequence: 2, Sum: 3}.String(), "000001000000023")

	c, err := ParseCode("123456123456781")
	assert.NoError(t, err)
	assert.Equal(t, 123456, c.FirstRegNo)
	assert.Equal(t, 12345678, c.Sequence)
	assert.Equal(t, 1, c.Sum)
}

func TestCodeSeq(t *testing.T) {
	c, err := ParseCode("310230001716146")
	assert.NoError(t, err)
	assert.Equal(t, "310230001716154", c.Next().String())
}

func TestValidDigitLen(t *testing.T) {
	assert.Equal(t, 3, ValidDigitLen(123))
	assert.Equal(t, 5, ValidDigitLen(12345))

	s, err := DigitStringToIntSlice("12345")
	assert.NoError(t, err)
	assert.Equal(t, IntToIntSlice(12345), s)

	assert.Equal(t, 6, Mod1110(IntToIntSlice(11010800000001)))
	assert.Equal(t, 6, Mod1110(IntToIntSlice(31023000171614)))
	assert.Equal(t, 4, Mod1110(IntToIntSlice(31023000171615)))
}
