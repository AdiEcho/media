package pluto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

func (w *WebAddress) Set(s string) error {
	for {
		var (
			key string
			ok  bool
		)
		key, s, ok = strings.Cut(s, "/")
		if !ok {
			return nil
		}
		switch key {
		case "episode":
			w.episode = s
		case "movies":
			w.series = s
		case "series":
			w.series, s, ok = strings.Cut(s, "/")
			if !ok {
				return fmt.Errorf("%q", w.series)
			}
		}
	}
}

type WebAddress struct {
	series  string
	episode string
}

func (w WebAddress) String() string {
	var b strings.Builder
	b.WriteString("https://pluto.tv/on-demand/")
	if w.episode != "" {
		b.WriteString("series")
	} else {
		b.WriteString("movies")
	}
	b.WriteByte('/')
	b.WriteString(w.series)
	if w.episode != "" {
		b.WriteString("/episode/")
		b.WriteString(w.episode)
	}
	return b.String()
}

func (w WebAddress) video() (*Video, error) {
	req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = url.Values{
		"appName":           {"web"},
		"appVersion":        {"9"},
		"clientID":          {"9"},
		"clientModelNumber": {"9"},
		"drmCapabilities":   {"widevine:L3"},
		"seriesIDs":         {w.series},
	}.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var start struct {
		VOD []Video
	}
	err = json.NewDecoder(res.Body).Decode(&start)
	if err != nil {
		return nil, err
	}
	demand := start.VOD[0]
	if demand.Slug.slug != w.series {
		if demand.ID != w.series {
			return nil, fmt.Errorf("%+v", demand)
		}
	}
	for _, s := range demand.Seasons {
		s.parent = &demand
		for _, e := range s.Episodes {
			err := e.Slug.atoi()
			if err != nil {
				return nil, err
			}
			e.parent = s
			if e.Episode == w.episode {
				return e, nil
			}
			if e.Slug.slug == w.episode {
				return e, nil
			}
		}
	}
	err = demand.Slug.atoi()
	if err != nil {
		return nil, err
	}
	return &demand, nil
}

// ex-machina-2015-1-1-ptv1
// head-first-1998-1-2
// king-of-queens
// pilot-1998-1-1-ptv8
func (s *slug) atoi() error {
	split := strings.Split(s.slug, "-")
	slices.Reverse(split)
	if strings.HasPrefix(split[0], "ptv") {
		split = split[1:]
	}
	var err error
	s.episode, err = strconv.Atoi(split[0])
	if err != nil {
		return err
	}
	s.season, err = strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	s.year, err = strconv.Atoi(split[2])
	if err != nil {
		return err
	}
	return nil
}

func (s *slug) UnmarshalText(text []byte) error {
	s.slug = string(text)
	return nil
}

type season struct {
	Episodes []*Video
	parent   *Video
}

func (n namer) Show() string {
	if v := n.v.parent; v != nil {
		return v.parent.Name
	}
	return ""
}

type slug struct {
	episode int
	season  int
	slug    string
	year    int
}

type namer struct {
	v *Video
}

func (n namer) Season() int {
	return n.v.Slug.season
}

func (n namer) Episode() int {
	return n.v.Slug.episode
}

type Video struct {
	Episode string `json:"_id"`
	ID      string
	Name    string
	Seasons []*season
	Slug    slug
	parent  *season
}

func (n namer) Title() string {
	return n.v.Name
}

func (n namer) Year() int {
	return n.v.Slug.year
}
