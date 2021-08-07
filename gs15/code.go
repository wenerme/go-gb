package gs15

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func ValidDigitLen(n int) int {
	v := 1
	for i := 1; ; i++ {
		v *= 10
		if n < v {
			return i
		}
	}
}

func DigitStringToIntSlice(v string) ([]int, error) {
	s := make([]int, len(v))
	for i, r := range v {
		d := int(r - '0')
		if d < 0 || d > 9 {
			return nil, errors.Errorf("invalid digit %q in %q", string(r), v)
		}
		s[i] = d
	}
	return s, nil
}

func IntToIntSlice(v int) []int {
	l := ValidDigitLen(v)
	s := make([]int, l)
	n := v
	for i := range s {
		s[l-i-1] = n % 10
		n /= 10
	}
	return s
}

func Mod1110(n []int) int {
	m10 := 10
	for _, v := range n {
		m10 += v

		m10 %= 10
		if m10 == 0 {
			m10 = 10
		}
		m10 *= 2

		m10 %= 11
	}
	// 暴力 模反元素
	for i := 0; i < 10; i++ {
		if (m10+i)%10 == 1 {
			return i
		}
	}
	panic("unexpected")
}

type Code struct {
	FirstRegNo int // 首次登记机关码
	Sequence   int // 顺序码 8
	Sum        int // 校验码 1
}

func (c Code) String() string {
	return fmt.Sprintf("%06d%08d%01d", c.FirstRegNo, c.Sequence, c.Sum)
}

func (c Code) CalcSum() int {
	s := fmt.Sprintf("%06d%08d", c.FirstRegNo, c.Sequence)
	i, _ := DigitStringToIntSlice(s)
	return Mod1110(i)
}

func (c Code) IsValid() bool {
	return c.Sum == c.CalcSum()
}

func (c Code) Prev() Code {
	n := Code{
		FirstRegNo: c.FirstRegNo,
		Sequence:   c.Sequence - 1,
	}
	n.Sum = n.CalcSum()
	return n
}

func (c Code) Next() Code {
	n := Code{
		FirstRegNo: c.FirstRegNo,
		Sequence:   c.Sequence + 1,
	}
	n.Sum = n.CalcSum()
	return n
}

func ParseCode(s string) (c *Code, err error) {
	if len(s) != FullLen {
		return nil, fmt.Errorf("不是15位营业执照编号: %v", len(s))
	}
	c = &Code{}
	c.FirstRegNo, err = strconv.Atoi(s[0:6])
	if err == nil {
		c.Sequence, err = strconv.Atoi(s[6:14])
	}
	if err == nil {
		c.Sum, err = strconv.Atoi(s[14:15])
	}
	err = errors.Wrap(err, "解析营业执照编号失败")
	return
}
