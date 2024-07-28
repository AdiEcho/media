package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

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

func (a AuthLogin) Entitlement(movie *FullMovie) (*Entitlement, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/entitlement/v2/asset/" + movie.DefaultPlayable.Id
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

type AuthLogin struct {
   Data []byte
   v *struct {
      Token string
   }
}
