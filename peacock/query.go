package peacock

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (q *QueryNode) New(content_id string) error {
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

func (QueryNode) Owner() (string, bool) {
   return "", false
}

func (q QueryNode) Season() (string, bool) {
   if v := q.Attributes.SeasonNumber; v >= 1 {
      return strconv.Itoa(v), true
   }
   return "", false
}

func (q QueryNode) Episode() (string, bool) {
   if v := q.Attributes.EpisodeNumber; v >= 1 {
      return strconv.Itoa(v), true
   }
   return "", false
}

func (q QueryNode) Title() (string, bool) {
   return q.Attributes.Title, true
}

func (q QueryNode) Show() (string, bool) {
   if v := q.Attributes.SeriesName; v != "" {
      return v, true
   }
   return "", false
}

type QueryNode struct {
   Attributes struct {
      EpisodeNumber int
      SeasonNumber int
      SeriesName string
      Title string
      Year int
   }
}

func (q QueryNode) Year() (string, bool) {
   if q.Attributes.SeriesName != "" {
      return "", false
   }
   return strconv.Itoa(q.Attributes.Year), true
}
