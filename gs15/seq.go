package gs15

func Prev(s string) (string, error) {
	c, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	return c.Prev().String(), nil
}

func Next(s string) (string, error) {
	c, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	return c.Next().String(), nil
}
