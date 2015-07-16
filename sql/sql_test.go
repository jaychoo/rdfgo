package sql

import (
	"fmt"
	"testing"
)

func getTestTable() {
}

func TestEntry(t *testing.T) {
	e := Entry{
		Key:   "id",
		Value: 1,
	}
	fmt.Println(e)
}

func TestRow(t *testing.T) {
	e := Entry{
		Key:   "id",
		Value: 1,
	}

	r := Row{
		Columns: []Entry{
			e,
		},
	}

	fmt.Println(r)
}

func TestTable(t *testing.T) {
	e1 := Entry{
		Key:   "id",
		Value: 1,
	}

	e2 := Entry{
		Key:   "name",
		Value: "Test",
	}

	r := map[int]Row{
		1: Row{
			Columns: []Entry{
				e1,
				e2,
			},
		},
	}

	tb := Table{
		DataBaseName: "testDB",
		DataBaseURI:  "local:local@localhost/local",
		Name:         "test",
		Rows:         r,
	}

	fmt.Println(tb)
}
