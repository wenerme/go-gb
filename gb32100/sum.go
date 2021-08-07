package gb32100

import (
	"fmt"
)

// GB/T 17710

var _weight = []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}

func IsValid(s string) bool {
	if len(s) != FullLen {
		return false
	}
	sum, err := Sum(s[0:17])
	c := s[17:18]
	return err == nil && sum == c
}

func Sum(s string) (string, error) {
	if len(s) != ContentLen {
		return "", fmt.Errorf("need %v got: %v", ContentLen, len(s))
	}
	sum := 0
	for i, r := range s {
		d, err := rtoi(r)
		if err != nil {
			return "", err
		}
		sum += d * _weight[i]
	}
	sum %= 31
	sum = 31 - sum
	if sum == 30 {
		return "Y", nil
	}
	return string(itor(sum)), nil
}

func itor(i int) rune {
	switch {
	case i < 10:
		return rune('0' + i)
	default:
		return _intToRuneMap[i]
	}
}

func rtoi(r rune) (int, error) {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0'), nil
	case r >= 'A' && r <= 'Z':
		v, ok := _runeToIntMap[r]
		if ok {
			return v, nil
		}
		return 0, fmt.Errorf("invalid rune: %q", string(r))
	default:
		return 0, fmt.Errorf("invalid rune: %q", string(r))
	}
}

var _runeToIntMap = map[rune]int{
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'G': 16,
	'H': 17,
	'J': 18,
	'K': 19,
	'L': 20,
	'M': 21,
	'N': 22,
	'P': 23,
	'Q': 24,
	'R': 25,
	'T': 26,
	'U': 27,
	'W': 28,
	'X': 29,
	'Y': 30,
}

var _intToRuneMap = map[int]rune{
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
	16: 'G',
	17: 'H',
	18: 'J',
	19: 'K',
	20: 'L',
	21: 'M',
	22: 'N',
	23: 'P',
	24: 'Q',
	25: 'R',
	26: 'T',
	27: 'U',
	28: 'W',
	29: 'X',
	30: 'Y',
}
