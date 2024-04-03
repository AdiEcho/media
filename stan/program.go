package stan

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (p *legacy_program) New(id int64) error {
   address := func() string {
      b := []byte("https://api.stan.com.au/programs/v1/legacy/programs/")
      b = strconv.AppendInt(b, id, 10)
      b = append(b, ".json"...)
      return string(b)
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(&p.v)
}

type legacy_program struct {
   v struct {
      ReleaseYear int
      SeriesTitle string
      Title string
      TvSeasonEpisodeNumber int
      TvSeasonNumber int
   }
}

func (p legacy_program) Episode() int {
   return p.v.TvSeasonEpisodeNumber
}

func (p legacy_program) Show() string {
   return p.v.SeriesTitle
}

func (p legacy_program) Season() int {
   return p.v.TvSeasonNumber
}

func (p legacy_program) Title() string {
   return p.v.Title
}

func (p legacy_program) Year() int {
   return p.v.ReleaseYear
}
