package gb32100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCode(t *testing.T) {
	u, err := ParseCode("91350100M000100Y43")
	assert.NoError(t, err)
	assert.Equal(t, &Code{
		RegDeptCode:  "9",
		OrgTypeCode:  "1",
		DivisionCode: "350100",
		OrgCode:      "M000100Y4",
		Sum:          "3",
	}, u)
	assert.True(t, u.IsValid())
	assert.True(t, IsValid(u.String()))
	assert.True(t, IsValid("91350100M000100Y43"))
	{
		next, err := Next("91310230MA1K314J7C")
		assert.NoError(t, err)
		assert.Equal(t, "91310230MA1K314K57", next)
	}
	{
		next, err := Prev("91310230MA1K314K57")
		assert.NoError(t, err)
		assert.Equal(t, "91310230MA1K314J7C", next)
	}
}
