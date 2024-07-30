package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

type Playback struct {
   Drm struct {
      Schemes struct {
         Widevine struct {
            LicenseUrl string
         }
      }
   }
   Manifest struct {
      Url Url
   }
}

func (u *Url) UnmarshalText(text []byte) error {
   u.Url = new(url.URL)
   err := u.Url.UnmarshalBinary(text)
   if err != nil {
      return err
   }
   query := u.Url.Query()
   manifest := query["r.manifest"]
   query["r.manifest"] = manifest[len(manifest)-1:]
   u.Url.RawQuery = query.Encode()
   return nil
}

type Url struct {
   Url *url.URL
}

func (d DefaultToken) Playback(web WebAddress) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var p playback_request
      p.ConsumptionType = "streaming"
      p.EditId = web.EditId
      return json.MarshalIndent(p, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.discomax.com",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b bytes.Buffer
      b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
      b.WriteString("/playbackInfo")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization": {"Bearer "+d.Body.Data.Attributes.Token},
      "content-type": {"application/json"},
      "x-wbd-session-state": {d.SessionState},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   play := new(Playback)
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type DefaultRoutes struct {
   Data struct {
      Attributes struct {
         Url WebAddress
      }
   }
   Included []route_include
}

func (w WebAddress) MarshalText() ([]byte, error) {
   var b bytes.Buffer
   if w.VideoId != "" {
      b.WriteString("/video/watch/")
      b.WriteString(w.VideoId)
   }
   if w.EditId != "" {
      b.WriteByte('/')
      b.WriteString(w.EditId)
   }
   return b.Bytes(), nil
}

type WebAddress struct {
   VideoId string
   EditId  string
}

func (w *WebAddress) UnmarshalText(text []byte) error {
   s := string(text)
   if !strings.Contains(s, "/video/watch/") {
      return errors.New("/video/watch/ not found")
   }
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "play.max.com")
   s = strings.TrimPrefix(s, "/video/watch/")
   var found bool
   w.VideoId, w.EditId, found = strings.Cut(s, "/")
   if !found {
      return errors.New("/ not found")
   }
   return nil
}

func (d DefaultToken) Routes(web WebAddress) (*DefaultRoutes, error) {
   address := func() string {
      path, _ := web.MarshalText()
      var b strings.Builder
      b.WriteString("https://default.any-")
      b.WriteString(home_market)
      b.WriteString(".prd.api.discomax.com/cms/routes")
      b.Write(path)
      return b.String()
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "include": {"default"},
      // this is not required, but results in a smaller response
      "page[items.size]": {"1"},
   }.Encode()
   req.Header = http.Header{
      "authorization": {"Bearer "+d.Body.Data.Attributes.Token},
      "x-wbd-session-state": {d.SessionState},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   route := new(DefaultRoutes)
   err = json.NewDecoder(resp.Body).Decode(route)
   if err != nil {
      return nil, err
   }
   return route, nil
}
