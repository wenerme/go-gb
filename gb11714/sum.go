package gb11714

import (
	"fmt"
	"strconv"
)

var _weight = []int{3, 7, 9, 10, 5, 8, 4, 2}

func Sum(s string) (string, error) {
	if len(s) != ContentLen {
		return "", fmt.Errorf("invalid 8 digit code: got %v", len(s))
	}
	sum := 0
	for i, r := range s {
		d, err := rtoi(r)
		if err != nil {
			return "", err
		}
		sum += d * _weight[i]
	}
	sum %= 11
	sum = 11 - sum
	if sum == 10 {
		return "X", nil
	}
	return strconv.Itoa(sum), nil
}
