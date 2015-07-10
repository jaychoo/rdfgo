package n3

import "testing"

func getTestObject() N3 {
	n := N3{
		Prefix:   "http://purl.org/dc/elements/1.1/",
		PrefixNS: "dc",
		Uri:      "http://en.wikipedia.org/wiki/Tony_Benn",
		Items: []N3Element{
			N3Element{
				Key:   "title",
				Value: "Tony Benn",
			},
			N3Element{
				Key:   "publisher",
				Value: "Wikipedia",
			},
		},
	}
	return n
}

func TestN3(t *testing.T) {
	n := getTestObject()
	n.Print()
	return
}

func TestWriteN3(t *testing.T) {
	n := getTestObject()
	n.WriteN3()
	return
}
