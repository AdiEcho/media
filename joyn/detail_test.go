package joyn

import (
	"154.pages.dev/encoding"
	"fmt"
	"testing"
	"time"
)

func TestDetail(t *testing.T) {
	for _, test := range tests {
		detail, err := NewDetail(test.path)
		if err != nil {
			t.Fatal(err)
		}
		name, err := encoding.Name(Namer{detail})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%q\n", name)
		fmt.Printf("%+v\n", detail)
		time.Sleep(time.Second)
	}
}

var tests = []struct {
	key_id string
	path   string
}{
	{
		// joyn.de/filme/barry-seal-only-in-america
		key_id: "e+os9wvbQLpkvIFRuG3exA==",
		path:   "/filme/barry-seal-only-in-america",
	},
	{
		// joyn.de/serien/one-tree-hill/1-2-quaelende-angst
		path: "/serien/one-tree-hill/1-2-quaelende-angst",
	},
}
