package rakuten

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (w WebAddress) movie() (*GizmoMovie, error) {
	req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
	if err != nil {
		return nil, err
	}
	req.URL.Path = "/v3/movies/" + w.content_id
	req.URL.RawQuery = url.Values{
		"market_code":       {w.market_code},
		"classification_id": {strconv.Itoa(w.classification_id)},
		"device_identifier": {"atvui40"},
	}.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var b strings.Builder
		res.Write(&b)
		return nil, errors.New(b.String())
	}
	movie := new(GizmoMovie)
	err = json.NewDecoder(res.Body).Decode(movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (w WebAddress) String() string {
	var b strings.Builder
	if w.market_code != "" {
		b.WriteString("https://www.rakuten.tv/")
		b.WriteString(w.market_code)
	}
	if w.content_id != "" {
		b.WriteString("/movies/")
		b.WriteString(w.content_id)
	}
	return b.String()
}

func (w *WebAddress) Set(s string) error {
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "www.")
	s = strings.TrimPrefix(s, "rakuten.tv")
	s = strings.TrimPrefix(s, "/")
	var found bool
	w.market_code, w.content_id, found = strings.Cut(s, "/movies/")
	if !found {
		return errors.New("/movies/ not found")
	}
	w.classification_id, found = classification_id[w.market_code]
	if !found {
		return errors.New("market_code not found")
	}
	return nil
}

var classification_id = map[string]int{
	"dk": 283,
	"fi": 284,
	"fr": 23,
	"no": 286,
	"se": 282,
}

func (GizmoMovie) Show() string {
	return ""
}

func (GizmoMovie) Season() int {
	return 0
}

func (GizmoMovie) Episode() int {
	return 0
}

func (g GizmoMovie) Title() string {
	return g.Data.Title
}

type GizmoMovie struct {
	Data struct {
		Title string
		Year  int
	}
}

func (g GizmoMovie) Year() int {
	return g.Data.Year
}

type WebAddress struct {
	classification_id int
	content_id        string
	market_code       string
}
