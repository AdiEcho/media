package roku

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func New_Content(id string) (*Content, error) {
   req, err := http.NewRequest(
      "GET", "https://therokuchannel.roku.com/api/v2/homescreen/content", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL = func() *url.URL {
      include := []string{
         "episodeNumber",
         "releaseDate",
         "seasonNumber",
         "series.title",
         "title",
         "viewOptions",
      }
      expand := url.URL{
         Scheme: "https",
         Host: "content.sr.roku.com",
         Path: "/content/v1/roku-trc/" + id,
         RawQuery: url.Values{
            "expand": {"series"},
            "include": {strings.Join(include, ",")},
         }.Encode(),
      }
      homescreen := url.PathEscape(expand.String())
      return req.URL.JoinPath(homescreen)
   }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var con Content
   if err := json.NewDecoder(res.Body).Decode(&con.s); err != nil {
      return nil, err
   }
   return &con, nil
}

func (c Content) DASH() *Video {
   for _, option := range c.s.View_Options {
      for _, vid := range option.Media.Videos {
         if vid.Video_Type == "DASH" {
            return &vid
         }
      }
   }
   return nil
}

func (Content) Owner() (string, bool) {
   return "", false
}

func (c Content) Season() (string, bool) {
   return c.s.Season_Number, c.s.Season_Number != ""
}

func (c Content) Episode() (string, bool) {
   return c.s.Episode_Number, c.s.Episode_Number != ""
}

func (c Content) Title() (string, bool) {
   return c.s.Title, true
}

// we have to embed to prevent clobbering Namer.Title
type Content struct {
   s struct {
      Series *struct {
         Title string
      }
      Title string
      Release_Date string `json:"releaseDate"` // 2007-01-01T000000Z
      View_Options []struct {
         Media struct {
            Videos []Video
         }
      } `json:"viewOptions"`
      Season_Number string `json:"seasonNumber"`
      Episode_Number string `json:"episodeNumber"`
   }
}

func (c Content) Show() (string, bool) {
   if c.s.Series != nil {
      return c.s.Series.Title, true
   }
   return "", false
}

func (c Content) Year() (string, bool) {
   if c.s.Series != nil {
      return "", false
   }
   year, _, _ := strings.Cut(c.s.Release_Date, "-")
   return year, true
}
