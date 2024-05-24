package rakuten

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

type web_address struct {
   content_id string
   market_code string
}

func (w web_address) String() string {
   var b strings.Builder
   b.WriteString("https://www.rakuten.tv/")
   b.WriteString(w.market_code)
   b.WriteString("/movies/")
   b.WriteString(w.content_id)
   return b.String()
}

func (g *gizmo_movie) New() error {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/v3/movies/jerry-maguire"
   req.URL.RawQuery = url.Values{
      "classification_id": {"23"},
      "market_code": {"fr"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(g)
}

func (gizmo_movie) Show() string {
   return ""
}

func (gizmo_movie) Season() int {
   return 0
}

func (gizmo_movie) Episode() int {
   return 0
}

func (g gizmo_movie) Title() string {
   return g.Data.Title
}

type gizmo_movie struct {
   Data struct {
      Title string
      Year int
   }
}

func (g gizmo_movie) Year() int {
   return g.Data.Year
}
