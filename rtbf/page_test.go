package rtbf

import (
	"154.pages.dev/text"
	"fmt"
	"testing"
	"time"
)

func TestPage(t *testing.T) {
	for _, medium := range media {
		page, err := NewPage(medium.path)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", page)
		name, err := text.Name(page)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%q\n", name)
		time.Sleep(time.Second)
	}
}

var media = []struct {
	id     int64
	key_id string
	path   string
	url    string
}{
	{
		id:     3201987,
		key_id: "o1C37Tt5SzmHMmEgQViUEA==",
		path:   "/media/i-care-a-lot-i-care-a-lot-3201987",
		url:    "auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987",
	},
	{
		path: "/media/grantchester-grantchester-s01-3194636",
		url:  "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
	},
	{
		path: "/emission/i-care-a-lot-27462",
		url:  "auvio.rtbf.be/emission/i-care-a-lot-27462",
	},
}
