package command

import (
	"fmt"

	"github.com/jaychoo/rdfgo/n3"
	"github.com/jaychoo/rdfgo/sql"
)

// ConvertTableToN3 reurn a list of N3 structs from sql.Table
func ConvertTableToN3(t *sql.Table) []n3.N3 {
	r := []n3.N3{}

	for _, v := range t.Rows {
		n := n3.N3{
			Prefix:   t.DataBaseURI,
			PrefixNS: t.Name,
			URI:      "Test",
		}

		for _, c := range v.Columns {
			n.Items = append(n.Items, n3.N3Element{
				Key:   c.Key,
				Value: c.Value,
			})
		}

		r = append(r, n)
	}

	return r
}

func main() {
	fmt.Println("Test")
}
