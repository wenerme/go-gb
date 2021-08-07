package gbt2260

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	// 直辖市 Parent 级别不同
	// 310210 - 没有 310200 - 有 310000

	c := LookupCode("320107")
	p := LookupCode("320100")
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", c.Parent())
	fmt.Printf("%#v\n", c.Parent().Parent())
	cc := LookupName(c.Name)
	assert.Equal(t, c, cc)
	assert.Equal(t, p, c.Parent())
}
