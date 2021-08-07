package gb11714

import (
	"fmt"
)

// rtoi base36
func rtoi(v rune) (d int, err error) {
	switch {
	case v >= '0' && v <= '9':
		d = int(v - '0')
	case v >= 'A' && v <= 'Z':
		d = int(v-'A') + 10
	default:
		err = fmt.Errorf("invalid base36 char %q", string(v))
	}
	return
}

func enc(n int) string {
	s := ""
	v := n

	for v > 0 {
		d := v % 36
		v /= 36
		switch {
		case d < 10:
			s = string(rune('0'+d)) + s
		default:
			s = string(rune('A'+d-10)) + s
		}
	}
	return s
}

func dec(n string) (int, error) {
	// js Number.MAX_SAFE_INTEGER / Math.pow(36,10) ~= 2
	// int is enough for 8 base36
	s := 0
	b := 1
	l := len(n)
	for i := range n {
		r := n[l-i-1]
		v, err := rtoi(rune(r))
		if err != nil {
			return 0, err
		}
		s += v * b
		b *= 36
	}
	return s, nil
}
