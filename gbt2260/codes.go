package gbt2260

import (
	_ "embed"
	"encoding/csv"
	"strconv"
	"strings"
)

//go:embed codes.csv
var codes string

type Record struct {
	Code string
	Name string
	Year int
}

func (r *Record) Parent() *Record {
	if r == nil {
		return nil
	}
	if strings.HasSuffix(r.Code, "0000") {
		return nil
	}
	if strings.HasSuffix(r.Code, "00") {
		return LookupCode(r.Code[0:2] + "0000")
	}
	return LookupCode(r.Code[0:4] + "00")
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

type store struct {
	records []*Record
	names   map[string]*Record
	codes   map[string]*Record
}

func (s *store) index() {
	s.names = make(map[string]*Record, len(s.records))
	s.codes = make(map[string]*Record, len(s.records))
	for _, r := range s.records {
		s.names[r.Name] = r
		s.codes[r.Code] = r
	}
}

var _store *store

func getStore() *store {
	if _store == nil {
		s := &store{}
		r := csv.NewReader(strings.NewReader(codes))
		r.ReuseRecord = true
		for {
			row, _ := r.Read()
			if row == nil {
				break
			}
			rec := &Record{
				Code: row[0],
				Year: mustAtoi(row[1]),
				Name: row[2],
			}
			s.records = append(s.records, rec)
		}
		s.index()
		_store = s
	}
	return _store
}

func mustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
