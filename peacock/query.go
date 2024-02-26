package peacock

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (q *query_node) New(content_id string) error {
   req, err := http.NewRequest("GET", "https://atom.peacocktv.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/adapter-calypso/v3/query/node/content_id/" + content_id
   req.Header = http.Header{
      "X-Skyott-Proposition": {"NBCUOTT"},
      "X-Skyott-Territory": {"US"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(q)
}

func (query_node) Owner() (string, bool) {
   return "", false
}

func (query_node) Year() (string, bool) {
   return "", false
}

func (q query_node) Show() (string, bool) {
   return q.Attributes.SeriesName, true
}

func (q query_node) Title() (string, bool) {
   return q.Attributes.Title, true
}

func (q query_node) Season() (string, bool) {
   return strconv.Itoa(q.Attributes.SeasonNumber), true
}

func (q query_node) Episode() (string, bool) {
   return strconv.Itoa(q.Attributes.EpisodeNumber), true
}

type query_node struct {
   Attributes struct {
      EpisodeNumber int
      SeasonNumber int
      SeriesName string
      Title string
   }
}
