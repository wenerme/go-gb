package gb11643

func MustNext(s string) string {
	code, err := Next(s)
	if err != nil {
		panic(err)
	}
	return code
}

func MustPrev(s string) string {
	code, err := Prev(s)
	if err != nil {
		panic(err)
	}
	return code
}

func Next(s string) (string, error) {
	code, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	next, err := code.Next()
	if err != nil {
		return "", err
	}
	return next.String(), nil
}

func Prev(s string) (string, error) {
	code, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	next, err := code.Prev()
	if err != nil {
		return "", err
	}
	return next.String(), nil
}
