package stan

import (
   "net/http"
   "net/url"
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
   return json.NewDecoder(res.Body).Decode(p)
}

func (p legacy_program) Show() string {
   return p.SeriesTitle
}

func (p legacy_program) Season() int {
   return p.TvSeasonNumber
}

func (p legacy_program) Episode() int {
   return p.TvSeasonEpisodeNumber
}

type legacy_program struct {
   SeriesTitle string
   Title string
   TvSeasonEpisodeNumber int
   TvSeasonNumber int
}

func (p legacy_program) Title() string {
   return p.Title
}

func (legacy_program) Year() int {
   return 0
}
