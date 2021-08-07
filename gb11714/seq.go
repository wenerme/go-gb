package gb11714

func Prev(s string) (string, error) {
	n, err := dec(s[0:ContentLen])
	return build(n, -1), err
}

func Next(s string) (string, error) {
	n, err := dec(s[0:ContentLen])
	return build(n, +1), err
}

func build(n int, d int) string {
	v := n
	for {
		v += d
		// GB 32100-2015
		switch (v % 36) - 10 + 'A' {
		case 'I':
		case 'O':
		case 'Z':
		case 'S':
		case 'V':
		default:
			goto done
		}
	}
done:
	o := enc(v)

	c, _ := Sum(o)
	o += c
	return o
}
