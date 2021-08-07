package gbt2659

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	r := LookupCode2("CN")
	assert.Equal(t, &Record{
		NameZh:     "中国",
		NameEn:     "China",
		FullNameEn: "the People's Republic of China",
		Code2:      "CN",
		Code3:      "CHN",
		Num:        156,
	}, r)
}
