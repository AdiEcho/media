package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (a auth_login) entitlement(f *full_movie) (*entitlement, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/entitlement/v2/asset/" + f.DefaultPlayable.ID
   req.Header.Set("authorization", "Bearer " + a.v.Token)
   magine_accesstoken.set(req.Header)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   title := new(entitlement)
   err = json.NewDecoder(res.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

func (a *auth_login) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}

type auth_login struct {
   data []byte
   v struct {
      Token string
   }
}

func (a *auth_login) New(identity, key string) error {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "identity": identity,
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a auth_login) playback(
   movie *full_movie, title *entitlement,
) (*playback, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.DefaultPlayable.ID
   magine_accesstoken.set(req.Header)
   magine_play_devicemodel.set(req.Header)
   magine_play_deviceplatform.set(req.Header)
   magine_play_devicetype.set(req.Header)
   magine_play_drm.set(req.Header)
   magine_play_protocol.set(req.Header)
   req.Header.Set("authorization", "Bearer " + a.v.Token)
   req.Header.Set("magine-play-deviceid", "!")
   req.Header.Set("magine-play-entitlementid", title.Token)
   x_forwarded_for.set(req.Header)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   play := new(playback)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type entitlement struct {
   Token string
}

type header struct {
   key string
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

type playback struct {
   Headers map[string]string
   Playlist string
}

type poster struct {
   auth auth_login
   play *playback
}

func (poster) RequestUrl() (string, bool) {
   return "https://client-api.magine.com/api/playback/v1/widevine/license", true
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (p poster) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   magine_accesstoken.set(head)
   head.Set("authorization", "Bearer " + p.auth.v.Token)
   for key, value := range p.play.Headers {
      head.Set(key, value)
   }
   return head, nil
}
