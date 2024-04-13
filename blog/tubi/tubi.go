package tubi

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

type content struct {
   Children []*content
   Detailed_Type string
   Episode_Number int `json:",string"`
   ID int `json:",string"`
   Series_ID int `json:",string"`
   Title string
   Video_Resources []video_resource
   Year int
   parent *content
}

func (c *content) New(id int) error {
   req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(id)},
      "deviceId": {"!"},
      "platform": {"android"},
      "video_resources[]": {"dash_widevine"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := json.NewDecoder(res.Body).Decode(c); err != nil {
      return err
   }
   c.set(nil)
   return nil
}

func (c content) episode() bool {
   return c.Detailed_Type == "episode"
}

func (c content) get(id int) (*content, bool) {
   if c.ID == id {
      return &c, true
   }
   for _, child := range c.Children {
      if v, ok := child.get(id); ok {
         return v, true
      }
   }
   return nil, false
}

func (c *content) set(parent *content) {
   c.parent = parent
   for _, child := range c.Children {
      child.set(c)
   }
}

type namer struct {
   c *content
}

func (n namer) Episode() int {
   return n.c.Episode_Number
}

func (n namer) Season() int {
   if v := n.c.parent; v != nil {
      return v.ID
   }
   return 0
}

func (n namer) Show() string {
   if v := n.c.parent; v != nil {
      return v.parent.Title
   }
   return ""
}

func (n namer) Title() string {
   return n.c.Title
}

func (n namer) Year() int {
   return n.c.Year
}
