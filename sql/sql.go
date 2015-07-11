package sql

import (
	"fmt"
	"github.com/jaychoo/rdfgo"
)

type SqlTable struct {
	DataBaseName string
	DataBaseUri  string
	Name         string
	Rows         map[int]SqlTableRow
}

type SqlTableRow struct {
	Entry SqlEntry
}

type SqlEntry struct {
	Key   string
	Value interface{}
}

func (s *SqlTable) CreateN3() []N3 {
	r = []N3{}
	for k, v := range s.Rows {
		n = N3{}
		r = append(r, N3)
	}
}
