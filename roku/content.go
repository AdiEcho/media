package roku

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
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

// we have to embed to prevent clobbering the interface
type Content struct {
   s struct {
      Series *struct {
         Title string
      }
      Season_Number int64 `json:"seasonNumber,string"`
      Episode_Number int64 `json:"episodeNumber,string"`
      Title string
      Release_Date string `json:"releaseDate"` // 2007-01-01T000000Z
      View_Options []struct {
         Media struct {
            Videos []Video
         }
      } `json:"viewOptions"`
   }
}

func (c Content) Series() string {
   return c.s.Series.Title
}

func (c Content) Season() (int64, error) {
   return c.s.Season_Number, nil
}

func (c Content) Episode() (int64, error) {
   return c.s.Episode_Number, nil
}

func (c Content) Title() string {
   return c.s.Title
}

func (c Content) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, c.s.Release_Date)
}
