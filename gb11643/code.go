package gb11643

import (
	"fmt"
	"strconv"
	"time"
)

type Code struct {
	DivisionCode string // 6 登记管理机关行政区划码 GB/T 2260
	Date         time.Time
	Number       int
	Sum          string
}

func (c Code) IsMale() bool {
	return c.Number%2 == 1
}

func (c Code) IsFemale() bool {
	return c.Number%2 == 0
}

func (c Code) String() string {
	return c.MatserNumber() + c.Sum
}

func (c Code) MatserNumber() string {
	return fmt.Sprintf("%v%v%03d", c.DivisionCode, FormatDate(c.Date), c.Number)
}

func (c Code) IsValid() bool {
	sum, err := Sum(c.MatserNumber())
	return sum == c.Sum && err == nil
}

func (c Code) Next() (*Code, error) {
	o := &Code{
		DivisionCode: c.DivisionCode,
		Date:         c.Date,
		Number:       c.Number + 1,
	}
	if o.Number > 999 {
		o.Number = 0
		o.Date = o.Date.AddDate(0, 0, 1)
	}
	var err error
	o.Sum, err = Sum(o.MatserNumber())
	return o, err
}

func (c Code) Prev() (*Code, error) {
	o := &Code{
		DivisionCode: c.DivisionCode,
		Date:         c.Date,
		Number:       c.Number - 1,
	}
	if o.Number == -1 {
		o.Number = 999
		o.Date = o.Date.AddDate(0, 0, -1)
	}
	var err error
	o.Sum, err = Sum(o.MatserNumber())
	return o, err
}

func ParseCode(s string) (c *Code, err error) {
	c = &Code{
		DivisionCode: s[0:6],
		Sum:          s[17:18],
	}
	c.Date, err = ParseDate(s[6:14])
	if err == nil {
		c.Number, err = strconv.Atoi(s[14:17])
	}
	return
}

func MustParseDate(s string) time.Time {
	date, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return date
}

const DateLayout = "20060102"

func ParseDate(s string) (time.Time, error) {
	return time.Parse(DateLayout, s)
}

func FormatDate(d time.Time) string {
	return d.Format(DateLayout)
}
