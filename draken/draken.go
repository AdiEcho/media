package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (a AuthLogin) Entitlement(f *FullMovie) (*Entitlement, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/entitlement/v2/asset/" + f.DefaultPlayable.Id
   req.Header.Set("authorization", "Bearer "+a.v.Token)
   magine_accesstoken.set(req.Header)
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
   title := new(Entitlement)
   err = json.NewDecoder(resp.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

func (a AuthLogin) Playback(
   movie *FullMovie, title *Entitlement,
) (*Playback, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.DefaultPlayable.Id
   magine_accesstoken.set(req.Header)
   magine_play_devicemodel.set(req.Header)
   magine_play_deviceplatform.set(req.Header)
   magine_play_devicetype.set(req.Header)
   magine_play_drm.set(req.Header)
   magine_play_protocol.set(req.Header)
   req.Header.Set("authorization", "Bearer "+a.v.Token)
   req.Header.Set("magine-play-deviceid", "!")
   req.Header.Set("magine-play-entitlementid", title.Token)
   x_forwarded_for.set(req.Header)
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
   play := new(Playback)
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (a *AuthLogin) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.v)
}

type Entitlement struct {
   Token string
}

type header struct {
   key   string
   value string
}

var magine_accesstoken = header{
   "magine-accesstoken", "22cc71a2-8b77-4819-95b0-8c90f4cf5663",
}

var magine_play_devicemodel = header{
   "magine-play-devicemodel", "firefox 111.0 / windows 10",
}

var magine_play_deviceplatform = header{
   "magine-play-deviceplatform", "firefox",
}

var magine_play_devicetype = header{
   "magine-play-devicetype", "web",
}

var magine_play_drm = header{
   "magine-play-drm", "widevine",
}

var magine_play_protocol = header{
   "magine-play-protocol", "dashs",
}

// this value is important, with the wrong value you get random failures
var x_forwarded_for = header{
   "x-forwarded-for", "95.192.0.0",
}

func (h header) set(head http.Header) {
   head.Set(h.key, h.value)
}

type Playback struct {
   Headers  map[string]string
   Playlist string
}

type Poster struct {
   Auth AuthLogin
   Play *Playback
}

func (p Poster) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   magine_accesstoken.set(head)
   head.Set("authorization", "Bearer "+p.Auth.v.Token)
   for key, value := range p.Play.Headers {
      head.Set(key, value)
   }
   return head, nil
}

func (Poster) RequestUrl() (string, bool) {
   return "https://client-api.magine.com/api/playback/v1/widevine/license", true
}

func (Poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (Poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}
type AuthLogin struct {
   Data []byte
   v    struct {
      Token string
   }
}

func (a *AuthLogin) New(identity, key string) error {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "identity":  identity,
         "accessKey": key,
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://drakenfilm.se/api/auth/login", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   a.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}
