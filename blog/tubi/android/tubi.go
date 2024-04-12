package tubi

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

type content_management struct {
   Video_Resources []struct {
      Manifest struct {
         URL string
      }
   }
}

func (c *content_management) New(content_id int) error {
   req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(content_id)},
      "deviceId": {"ab55452c-66e0-4021-9619-5bdc25f26ae8"},
      "platform": {"android"},
      "video_resources[]": {"dash_widevine"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(c)
}
