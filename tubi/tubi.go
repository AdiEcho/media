package tubi

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (Content) Error() string {
   return "Content"
}

func (c Content) Get(id int) (*Content, bool) {
   if c.Id == id {
      return &c, true
   }
   for _, child := range c.Children {
      if v, ok := child.Get(id); ok {
         return v, true
      }
   }
   return nil, false
}

func (c Content) Marshal() ([]byte, error) {
   var buf bytes.Buffer
   enc := json.NewEncoder(&buf)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   err := enc.Encode(c)
   if err != nil {
      return nil, err
   }
   return buf.Bytes(), nil
}

func (c *Content) New(id int) error {
   req, err := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = c.Unmarshal(text)
   if err != nil {
      return err
   }
   if len(c.VideoResources) == 0 {
      return errors.New(string(text))
   }
   return nil
}

type Namer struct {
   Content *Content
}

func (c *Content) Unmarshal(text []byte) error {
   err := json.Unmarshal(text, c)
   if err != nil {
      return err
   }
   c.set(nil)
   return nil
}

type Content struct {
   Children       []*Content
   DetailedType   string `json:"detailed_type"`
   EpisodeNumber  int    `json:"episode_number,string"`
   Id             int    `json:",string"`
   SeriesId       int    `json:"series_id,string"`
   Title          string
   VideoResources []VideoResource `json:"video_resources"`
   Year           int
   parent         *Content
}

func (c Content) Episode() bool {
   return c.DetailedType == "episode"
}

func (c *Content) set(parent *Content) {
   c.parent = parent
   for _, child := range c.Children {
      child.set(c)
   }
}

func (n Namer) Episode() int {
   return n.Content.EpisodeNumber
}

func (n Namer) Season() int {
   if v := n.Content.parent; v != nil {
      return v.Id
   }
   return 0
}

func (n Namer) Show() string {
   if v := n.Content.parent; v != nil {
      return v.parent.Title
   }
   return ""
}

// S01:E03 - Hell Hath No Fury
func (n Namer) Title() string {
   if _, v, ok := strings.Cut(n.Content.Title, " - "); ok {
      return v
   }
   return n.Content.Title
}

func (n Namer) Year() int {
   return n.Content.Year
}
