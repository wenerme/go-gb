package gb11643

import (
	"fmt"
	"strconv"
)

var (
	_weights    = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1}
	_mod11Table = []int{1, 0, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

func IsValid(s string) bool {
	if len(s) != FullLen {
		return false
	}
	c, err := Sum(s[0:MasterLen])
	return err == nil && c == s[MasterLen:]
}

func Sum(s string) (string, error) {
	if len(s) != MasterLen {
		return "", fmt.Errorf("need %v got: %v", MasterLen, len(s))
	}
	sum := 0
	for i, r := range s {
		d := int(r - '0')
		if d < 0 || d > 9 {
			return "", fmt.Errorf("invalid char %q", string(r))
		}
		sum += d * _weights[i]
	}
	sum %= 11
	sum = _mod11Table[sum]
	if sum == 10 {
		return "X", nil
	}
	return strconv.Itoa(sum), nil
}
