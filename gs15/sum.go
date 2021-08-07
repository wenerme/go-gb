package gs15

func Sum(s string) (int, error) {
	i, err := DigitStringToIntSlice(s)
	return Mod1110(i), err
}
