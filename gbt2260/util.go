package gbt2260

import "strings"

func WithIn(child string, parent string) bool {
	a := Split(child)
	b := Split(parent)
	if a == nil || b == nil {
		return false
	}
	return a[0] == b[0] && (a[1] == b[1] || b[1] == "00") && (a[2] == b[2] || b[2] == "00")
}

func Split(s string) []string {
	if len(s) != 6 {
		return nil
	}
	return []string{s[0:2], s[2:4], s[4:6]}
}

func Level(s string) int {
	if strings.HasSuffix(s, "0000") {
		return 1
	}
	if strings.HasSuffix(s, "00") {
		return 2
	}
	return 3
}
