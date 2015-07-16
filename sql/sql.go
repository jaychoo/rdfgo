package sql

import (
	"fmt"
	"github.com/jaychoo/rdfgo/n3"
)

// Table - Simple representation of SQL table
type Table struct {
	DataBaseName string
	DataBaseURI  string
	Name         string
	Rows         map[int]Row
}

// Row - Table Row
type Row struct {
	Columns []Entry
}

// Entry - Rwo Entry
type Entry struct {
	Key   string
	Value interface{}
}

// CreateN3 - Create list of N3 struct instances from SQL Table
func (t *Table) CreateN3() []n3.N3 {
	r := []n3.N3{}
	for k, v := range t.Rows {
		fmt.Println(k, v)
		n := n3.N3{}
		r = append(r, n)
	}

	return r
}
