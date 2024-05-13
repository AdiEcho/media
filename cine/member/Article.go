package member

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// https://www.cinemember.nl/nl/films/american-hustle
func (a *ArticleSlug) Set(s string) error {
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "www.")
	s = strings.TrimPrefix(s, "cinemember.nl")
	s = strings.TrimPrefix(s, "/nl")
	s = strings.TrimPrefix(s, "/")
	*a = ArticleSlug(s)
	return nil
}

func (a ArticleSlug) String() string {
	return string(a)
}

type ArticleSlug string

func (a ArticleSlug) article() (*data_article, error) {
	body, err := func() ([]byte, error) {
		var s struct {
			Query     string `json:"query"`
			Variables struct {
				ArticleUrlSlug ArticleSlug `json:"articleUrlSlug"`
			} `json:"variables"`
		}
		s.Variables.ArticleUrlSlug = a
		s.Query = query_article
		return json.Marshal(s)
	}()
	if err != nil {
		return nil, err
	}
	res, err := http.Post(
		"https://api.audienceplayer.com/graphql/2/user",
		"application/json", bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var s struct {
		Data struct {
			Article data_article
		}
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	for _, asset := range s.Data.Article.Assets {
		asset.article = &s.Data.Article
	}
	return &s.Data.Article, nil
}

func (namer) Show() string {
	return ""
}

func (namer) Season() int {
	return 0
}

func (namer) Episode() int {
	return 0
}

func (n namer) Title() string {
	return n.d.CanonicalTitle
}

type data_article struct {
	Assets         []*article_asset
	CanonicalTitle string `json:"canonical_title"`
	ID             int
	Metas          []struct {
		Key   string
		Value string
	}
}

func (d data_article) year() (string, bool) {
	for _, meta := range d.Metas {
		if meta.Key == "year" {
			return meta.Value, true
		}
	}
	return "", false
}

type namer struct {
	d *data_article
}

func (n namer) Year() int {
	if v, ok := n.d.year(); ok {
		if v, err := strconv.Atoi(v); err == nil {
			return v
		}
	}
	return 0
}

const query_article = `
query($articleUrlSlug: String) {
   Article(full_url_slug: $articleUrlSlug) {
      ... on Article {
         assets {
            ... on Asset {
               id
               linked_type
            }
         }
         canonical_title
         id
         metas(output: html) {
            ... on ArticleMeta {
               key
               value
            }
         }
      }
   }
}
`

func (d data_article) film() (*article_asset, bool) {
	for _, asset := range d.Assets {
		if asset.LinkedType == "film" {
			return asset, true
		}
	}
	return nil, false
}

type article_asset struct {
	ID         int
	LinkedType string `json:"linked_type"`
	article    *data_article
}
