package gbt2659

func LookupCode2(code string) *Record {
	s := getStore()
	return dup(s.code2[code])
}

func dup(r *Record) *Record {
	if r != nil {
		rr := *r
		return &rr
	}
	return nil
}
