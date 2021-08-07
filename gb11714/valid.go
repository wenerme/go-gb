package gb11714

// IsValid check a full org code is valid
func IsValid(s string) bool {
	if s == "" || len(s) != FullLen {
		return false
	}
	l := len(s)
	c := s[l-1:]
	sum, err := Sum(s[0 : l-1])
	return c == sum && err == nil
}
