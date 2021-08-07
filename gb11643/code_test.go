package gb11643

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	{
		s, err := Sum("11010519491231002")
		assert.NoError(t, err)
		assert.Equal(t, "X", s)
	}
	{
		assert.True(t, IsValid("11010519491231002X"))
	}
	c, err := ParseCode("11010519491231002X")
	assert.NoError(t, err)
	assert.Equal(t, &Code{
		DivisionCode: "110105",
		Date:         MustParseDate("19491231"),
		Number:       2,
		Sum:          "X",
	}, c)
	assert.True(t, c.IsFemale())
	assert.True(t, c.IsValid())
	assert.Equal(t, "19491231", FormatDate(time.Date(1949, 12, 31, 0, 0, 0, 0, time.Local)))

	c, err = ParseCode("110105195001010039")
	assert.NoError(t, err)
	assert.True(t, c.IsMale())
	assert.True(t, c.IsValid())
}

func TestSeqEdge(t *testing.T) {
	for _, test := range []struct {
		a string
		b string
	}{
		{a: "11010519500101999X", b: "11010519500102000X"},
	} {
		assert.True(t, IsValid(test.a))
		assert.True(t, IsValid(test.b))
		assert.Equal(t, test.b, MustNext(test.a))
		assert.Equal(t, test.a, MustPrev(test.b))
	}
}
