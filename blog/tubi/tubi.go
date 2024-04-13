package tubi

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

func (c content) episode() bool {
   return c.Detailed_Type == "episode"
}

func (c *content) New(content_id int) error {
   req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(content_id)},
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
   c.set_parent()
   return nil
}

func (c *content) set_parent() {
   for _, child := range c.Children {
      child.parent = c
      child.set_parent()
   }
}

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
