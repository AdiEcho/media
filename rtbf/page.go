package rtbf

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type AuvioPage struct {
	Content struct {
		AssetId  string
		Subtitle subtitle
		Title    title
	}
}

// its just not available from what I can tell
func (AuvioPage) Year() int {
	return 0
}

func new_page(path string) (*AuvioPage, error) {
	res, err := http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var s struct {
		Data AuvioPage
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data, nil
}

func (a AuvioPage) Season() int {
	return a.Content.Title.season
}

type title struct {
	season int
	title  string
}

type subtitle struct {
	episode  int
	subtitle string
}

// json.data.content.title = "Grantchester S01";
// json.data.content.title = "I care a lot";
func (t *title) UnmarshalText(text []byte) error {
	t.title = string(text)
	if before, after, ok := strings.Cut(t.title, " S"); ok {
		if season, err := strconv.Atoi(after); err == nil {
			t.title = before
			t.season = season
		}
	}
	return nil
}

// json.data.content.subtitle = "06 - Les ombres de la guerre";
// json.data.content.subtitle = "Avec Rosamund Pike";
func (s *subtitle) UnmarshalText(text []byte) error {
	s.subtitle = string(text)
	if before, after, ok := strings.Cut(s.subtitle, " - "); ok {
		if episode, err := strconv.Atoi(before); err == nil {
			s.episode = episode
			s.subtitle = after
		}
	}
	return nil
}

func (a AuvioPage) Episode() int {
	return a.Content.Subtitle.episode
}

func (a AuvioPage) Show() string {
	if v := a.Content.Title; v.season >= 1 {
		return v.title
	}
	return ""
}

func (a AuvioPage) Title() string {
	if v := a.Content.Subtitle; v.episode >= 1 {
		return v.subtitle
	}
	return a.Content.Title.title
}
