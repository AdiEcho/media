package joyn

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func NewDetail(path string) (*DetailPage, error) {
	body, err := func() ([]byte, error) {
		var s struct {
			Query     string `json:"query"`
			Variables struct {
				Path string `json:"path"`
			} `json:"variables"`
		}
		s.Query = detail_page_static
		s.Variables.Path = path
		return json.Marshal(s)
	}()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", "https://api.joyn.de/graphql", bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"content-type":  {"application/json"},
		"joyn-platform": {"web"},
		"x-api-key":     {"4f0fd9f18abbe3cf0e87fdb556bc39c8"},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var s struct {
		Data struct {
			Page DetailPage
		}
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.Page, nil
}

func (n namer) Show() string {
	if v := n.d.Episode; v != nil {
		return v.Series.Title
	}
	return ""
}

func (n namer) Season() int {
	if v := n.d.Episode; v != nil {
		return v.Season.Number
	}
	return 0
}

func (n namer) Episode() int {
	if v := n.d.Episode; v != nil {
		return v.Number
	}
	return 0
}

func (n namer) Title() string {
	if v := n.d.Episode; v != nil {
		return v.Title
	}
	if v := n.d.Movie; v != nil {
		return v.Title
	}
	return ""
}

type DetailPage struct {
	Episode *struct {
		Video struct {
			ID string
		}
		Series struct {
			Title string
		}
		Season struct {
			Number int
		}
		Number int
		Title  string
	}
	Movie *struct {
		ProductionYear int `json:",string"`
		Title          string
		Video          struct {
			ID string
		}
	}
}

func (n namer) Year() int {
	if v := n.d.Movie; v != nil {
		return v.ProductionYear
	}
	return 0
}

type namer struct {
	d *DetailPage
}

const detail_page_static = `
query($path: String!) {
   page(path: $path) {
      ... on EpisodePage {
         episode {
            ... on Episode {
               video {
                  id
               }
               series {
                  title
               }
               season {
                  ... on Season {
                     number
                  }
               }
               number
               title
            }
         }
      }
      ... on MoviePage {
         movie {
            ... on Movie {
               productionYear
               title
               video {
                  id
               }
            }
         }
      }
   }
}
`

func (d DetailPage) content_id() (string, bool) {
	if v := d.Episode; v != nil {
		return v.Video.ID, true
	}
	if v := d.Movie; v != nil {
		return v.Video.ID, true
	}
	return "", false
}
