package tubi

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (c Content) Season() int {
   if v := c.parent; v != nil {
      return v.V.ID
   }
   return 0
}

// S01:E03 - Hell Hath No Fury
func (c Content) Title() string {
   if _, v, ok := strings.Cut(c.V.Title, " - "); ok {
      return v
   }
   return c.V.Title
}

type Content struct {
   Data []byte
   V struct {
      Children        []*Content
      DetailedType   string `json:"detailed_type"`
      EpisodeNumber  int `json:"episode_number,string"`
      ID              int `json:",string"`
      SeriesId       int `json:"series_id,string"`
      Title           string
      VideoResources []VideoResource `json:"video_resources"`
      Year            int
   }
   parent          *Content
}

func (c Content) Get(id int) (*Content, bool) {
   if c.V.ID == id {
      return &c, true
   }
   for _, child := range c.V.Children {
      if v, ok := child.Get(id); ok {
         return v, true
      }
   }
   return nil, false
}
func (c Content) EpisodeType() bool {
   return c.V.DetailedType == "episode"
}

func (c Content) Episode() int {
   return c.V.EpisodeNumber
}

// geo block VPN not x-forwarded-for
func (c *Content) New(id int) error {
   req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(id)},
      "deviceId":   {"!"},
      "platform":   {"android"},
      "video_resources[]": {
         "dash",
         "dash_widevine",
      },
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   c.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (c *Content) Unmarshal() error {
   err := json.Unmarshal(c.Data, &c.V)
   if err != nil {
      return err
   }
   c.set(nil)
   return nil
}

func (c *Content) set(parent *Content) {
   c.parent = parent
   for _, child := range c.V.Children {
      child.set(c)
   }
}

func (c Content) Year() int {
   return c.V.Year
}

func (c Content) Show() string {
   if v := c.parent; v != nil {
      return v.parent.V.Title
   }
   return ""
}

