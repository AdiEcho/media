package peacock

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

type video_playouts struct {
   Asset struct {
      Endpoints []struct {
         CDN string
         URL string
      }
   }
   Protection struct {
      LicenceAcquisitionUrl string // wikipedia.org/wiki/License
   }
}

func (v video_playouts) RequestUrl() (string, bool) {
   return v.Protection.LicenceAcquisitionUrl, true
}

func (video_playouts) RequestHeader(b []byte) (http.Header, error) {
   h := make(http.Header)
   h.Set("x-sky-signature", sign("POST", "/drm/widevine/acquirelicense", nil, b))
   return h, nil
}

func (video_playouts) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (video_playouts) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

func (a auth_tokens) video(content_id string) (*video_playouts, error) {
   body, err := func() ([]byte, error) {
      type capability struct {
         Acodec string `json:"acodec"`
         Container string `json:"container"`
         Protection string `json:"protection"`
         Transport string `json:"transport"`
         Vcodec string `json:"vcodec"`
      }
      var s struct {
         ContentId string `json:"contentId"`
         Device struct {
            Capabilities []capability `json:"capabilities"`
         } `json:"device"`
      }
      s.ContentId = content_id
      s.Device.Capabilities = []capability{
         {
            Acodec: "AAC",
            Container: "ISOBMFF",
            Protection: "WIDEVINE",
            Transport: "DASH",
            Vcodec: "H264",
         },
      }
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://play.ovp.peacocktv.com/video/playouts/vod",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("x-skyott-usertoken", a.UserToken)
   // `application/json` fails
   req.Header.Set("content-type", "application/vnd.playvod.v1+json")
   req.Header.Set(
      "x-sky-signature", sign(req.Method, req.URL.Path, req.Header, body),
   )
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   video := new(video_playouts)
   if err := json.NewDecoder(res.Body).Decode(video); err != nil {
      return nil, err
   }
   return video, nil
}
