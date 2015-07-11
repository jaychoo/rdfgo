package n3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type N3 struct {
	Prefix   string
	PrefixNS string
	Uri      string
	Items    []N3Element
}

type N3Element struct {
	Key   string
	Value interface{}
}

func (n *N3) Print() {
	for _, v := range n.getStringList() {
		fmt.Println(v)
	}

}

func replaceSuffix(in string, r rune, i int) string {
	o := []rune(in)
	o[i] = r
	return string(o)
}

func (n *N3) getStringList() []string {
	r := []string{}
	r = append(
		r,
		fmt.Sprintf("@prefix %s: <%s>.", n.PrefixNS, n.Prefix),
		"",
		fmt.Sprintf("<%s>", n.Uri),
	)

	if len(n.Items) > 0 {
		items := []string{}
		for _, i := range n.Items {
			items = append(items, fmt.Sprintf("  %s:%s %s;", n.PrefixNS, i.Key, i.Value))
		}

		lastOne := items[len(items)-1]
		items[len(items)-1] = replaceSuffix(lastOne, '.', len(lastOne)-1)
		r = append(r, items...)
	}

	return r
}

func (n *N3) WriteN3() {
	u := strings.Split(n.Uri, "/")
	fn := fmt.Sprintf("%s.n3", u[len(u)-1])
	f, err := os.Create(fn)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range n.getStringList() {
		_, err := w.WriteString(v + "\n")

		if err != nil {
			fmt.Println(err)
		}
	}
	w.Flush()
}
