package gbt2659

import (
	_ "embed"
	"encoding/csv"
	"strconv"
	"strings"
)

type Record struct {
	NameZh     string
	NameEn     string
	FullNameEn string
	Code2      string // 两字母代码 - 例如用于 互联网的地理级别域名
	Code3      string // 三字母代码
	Num        int    // 数字代码 - 与联合国统计处定义的代码类似
}

//go:embed records.csv
var data string

type store struct {
	records []*Record
	code2   map[string]*Record
}

func (s *store) index() {
	s.code2 = make(map[string]*Record, len(s.records))
	for _, r := range s.records {
		s.code2[r.Code2] = r
	}
}

var _store *store

func getStore() *store {
	if _store == nil {
		s := &store{}
		r := csv.NewReader(strings.NewReader(data))
		r.ReuseRecord = true
		for {
			row, _ := r.Read()
			if row == nil {
				break
			}
			rec := &Record{
				NameZh:     row[0],
				NameEn:     row[1],
				FullNameEn: row[2],
				Code2:      row[3],
				Code3:      row[4],
				Num:        mustAtoi(row[5]),
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
