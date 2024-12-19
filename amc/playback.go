package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

func (a *Authorization) Playback(nid string) (*Playback, error) {
   var value struct {
      AdTags struct {
         Lat int `json:"lat"`
         Mode string `json:"mode"`
         Ppid int `json:"ppid"`
         PlayerHeight int `json:"playerHeight"`
         PlayerWidth int `json:"playerWidth"`
         Url string `json:"url"`
      } `json:"adtags"`
   }
   value.AdTags.Mode = "on-demand"
   value.AdTags.Url = "-"
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/playback-id/api/v1/playback/" + nid
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Data.AccessToken},
      "content-type": {"application/json"},
      "x-amcn-device-ad-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-service-id": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-ccpa-do-not-sell": {"doNotPassData"},
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
   var play Playback
   err = json.NewDecoder(resp.Body).Decode(&play)
   if err != nil {
      return nil, err
   }
   play.AmcnBcJwt = resp.Header.Get("x-amcn-bc-jwt")
   return &play, nil
}

func (p *Playback) Dash() (*DataSource, bool) {
   for _, source := range p.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source, true
      }
   }
   return nil, false
}

func (p *Playback) RequestUrl() (string, bool) {
   if v, ok := p.Dash(); ok {
      return v.KeySystems.Widevine.LicenseUrl, true
   }
   return "", false
}

func (p *Playback) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("bcov-auth", p.AmcnBcJwt)
   return head, nil
}

type Playback struct {
   AmcnBcJwt string `json:"-"`
   Data struct {
      PlaybackJsonData struct {
         Sources []DataSource
      }
   }
}
