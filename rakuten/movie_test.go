package rakuten

import (
	"154.pages.dev/encoding"
	"fmt"
	"testing"
	"time"
)

type movie_test struct {
	content_id string
	key_id     string
	url        string
}

var tests = map[string]movie_test{
	"fr": {
		content_id: "Y2YzNGEwM2JiYjRhYTg5OWRmNDJjM2NmN2E2Y2I5MjUtbWMtMC0xMzctMC0w",
		key_id:     "zzSgO7tKqJnfQsPPemy5JQ==",
		url:        "rakuten.tv/fr/movies/jerry-maguire",
	},
	"se": {
		content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
		key_id:     "mlNKHxLWjhojWfOHEP3bZQ==",
		url:        "rakuten.tv/se/movies/i-heart-huckabees",
	},
}

func TestMovie(t *testing.T) {
	for _, test := range tests {
		var web WebAddress
		err := web.Set(test.url)
		if err != nil {
			t.Fatal(err)
		}
		movie, err := web.movie()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", movie)
		name, err := encoding.Name(movie)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%q\n", name)
		time.Sleep(time.Second)
	}
}
