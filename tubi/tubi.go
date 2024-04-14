package tubi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Content struct {
	Children        []*Content
	Detailed_Type   string
	Episode_Number  int `json:",string"`
	ID              int `json:",string"`
	Series_ID       int `json:",string"`
	Title           string
	Video_Resources []VideoResource
	Year            int
	parent          *Content
}

func (c *Content) New(id int) error {
	req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = url.Values{
		"content_id": {strconv.Itoa(id)},
		"deviceId":   {"!"},
		"platform":   {"android"},
		"video_resources[]": {
			"dash",
			"dash_widevine",
		},
	}.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(c); err != nil {
		return err
	}
	c.set(nil)
	return nil
}

////////

func (c Content) Get(id int) (*Content, bool) {
	if c.ID == id {
		return &c, true
	}
	for _, child := range c.Children {
		if v, ok := child.Get(id); ok {
			return v, true
		}
	}
	return nil, false
}

func (c *Content) set(parent *Content) {
	c.parent = parent
	for _, child := range c.Children {
		child.set(c)
	}
}

func (n namer) Episode() int {
	return n.c.Episode_Number
}

func (n namer) Season() int {
	if v := n.c.parent; v != nil {
		return v.ID
	}
	return 0
}

func (n namer) Show() string {
	if v := n.c.parent; v != nil {
		return v.parent.Title
	}
	return ""
}

func (n namer) Year() int {
	return n.c.Year
}

type namer struct {
	c *Content
}

func (c Content) episode() bool {
	return c.Detailed_Type == "episode"
}

// S01:E03 - Hell Hath No Fury
func (n namer) Title() string {
	if _, v, ok := strings.Cut(n.c.Title, " - "); ok {
		return v
	}
	return n.c.Title
}
