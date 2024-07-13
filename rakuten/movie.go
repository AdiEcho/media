package rakuten

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var classification_id = map[string]int{
   "dk": 283,
   "fi": 284,
   "fr": 23,
   "ie": 41,
   "no": 286,
   "se": 282,
   "ua": 276,
   "uk": 18,
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "rakuten.tv")
   s = strings.TrimPrefix(s, "/")
   var found bool
   a.market_code, a.content_id, found = strings.Cut(s, "/movies/")
   if !found {
      return errors.New("/movies/ not found")
   }
   a.classification_id, found = classification_id[a.market_code]
   if !found {
      return errors.New("market_code not found")
   }
   return nil
}

type Address struct {
   classification_id int
   content_id        string
   market_code       string
}

func (a Address) Movie() (*GizmoMovie, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.content_id
   req.URL.RawQuery = url.Values{
      "market_code":       {a.market_code},
      "classification_id": {strconv.Itoa(a.classification_id)},
      "device_identifier": {"atvui40"},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   movie := new(GizmoMovie)
   err = json.NewDecoder(resp.Body).Decode(movie)
   if err != nil {
      return nil, err
   }
   return movie, nil
}

func (a Address) String() string {
   var b strings.Builder
   if a.market_code != "" {
      b.WriteString("https://www.rakuten.tv/")
      b.WriteString(a.market_code)
   }
   if a.content_id != "" {
      b.WriteString("/movies/")
      b.WriteString(a.content_id)
   }
   return b.String()
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
