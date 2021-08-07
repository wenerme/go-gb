package gbt2260

func LookupCode(code string) *Record {
	s := getStore()
	return dup(s.codes[code])
}

func LookupName(code string) *Record {
	s := getStore()
	return dup(s.names[code])
}

func dup(r *Record) *Record {
	if r != nil {
		rr := *r
		return &rr
	}
	return nil
}
