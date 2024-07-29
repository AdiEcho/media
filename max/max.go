package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
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

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.Drm.Schemes.Widevine.LicenseUrl, true
}

type playback_request struct {
   AppBundle            string `json:"appBundle"`            // required
   ApplicationSessionId string `json:"applicationSessionId"` // required
   Capabilities         struct {
      Manifests struct {
         Formats struct {
            Dash struct{} `json:"dash"` // required
         } `json:"formats"` // required
      } `json:"manifests"` // required
   } `json:"capabilities"` // required
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo      struct {
      Player struct {
         MediaEngine struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"mediaEngine"` // required
         PlayerView struct {
            Height int `json:"height"` // required
            Width  int `json:"width"`  // required
         } `json:"playerView"` // required
         Sdk struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"sdk"` // required
      } `json:"player"` // required
   } `json:"deviceInfo"` // required
   EditId            string   `json:"editId"`
   FirstPlay         bool     `json:"firstPlay"`         // required
   Gdpr              bool     `json:"gdpr"`              // required
   PlaybackSessionId string   `json:"playbackSessionId"` // required
   UserPreferences   struct{} `json:"userPreferences"`   // required
}
// note you can use other keys, but you need to change home_market to match
var default_key = hmac_key{
   Id:  "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
}

type session_state map[string]string

func (s session_state) Set(text string) error {
   for text != "" {
      var key string
      key, text, _ = strings.Cut(text, ";")
      key, value, _ := strings.Cut(key, ":")
      s[key] = value
   }
   return nil
}

func (s session_state) String() string {
   var (
      b strings.Builder
      sep bool
   )
   for key, value := range s {
      if sep {
         b.WriteByte(';')
      } else {
         sep = true
      }
      b.WriteString(key)
      b.WriteByte(':')
      b.WriteString(value)
   }
   return b.String()
}

type hmac_key struct {
   Id  string
   Key []byte
}

type DefaultLogin struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

const home_market = "amer"

type default_decision struct {
   HmacKeys struct {
      Config struct {
         Android   *hmac_key
         AndroidTv *hmac_key
         FireTv    *hmac_key
         Hwa       *hmac_key
         Ios       *hmac_key
         TvOs      *hmac_key
         Web       *hmac_key
      }
   }
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

type DefaultRoutes struct {
   Data struct {
      Attributes struct {
         Url WebAddress
      }
   }
   Included []route_include
}

func (d DefaultRoutes) video() (*route_include, bool) {
   for _, include := range d.Included {
      if include.Id == d.Data.Attributes.Url.VideoId {
         return &include, true
      }
   }
   return nil, false
}

func (d DefaultRoutes) Season() int {
   if v, ok := d.video(); ok {
      return v.Attributes.SeasonNumber
   }
   return 0
}

func (d DefaultRoutes) Episode() int {
   if v, ok := d.video(); ok {
      return v.Attributes.EpisodeNumber
   }
   return 0
}

func (d DefaultRoutes) Title() string {
   if v, ok := d.video(); ok {
      return v.Attributes.Name
   }
   return ""
}

func (d DefaultRoutes) Year() int {
   if v, ok := d.video(); ok {
      return v.Attributes.AirDate.Year()
   }
   return 0
}

func (d DefaultRoutes) Show() string {
   if v, ok := d.video(); ok {
      if v.Attributes.SeasonNumber >= 1 {
         for _, include := range d.Included {
            if include.Id == v.Relationships.Show.Data.Id {
               return include.Attributes.Name
            }
         }
      }
   }
   return ""
}

type PublicKey struct {
   Token string
}

func (p *PublicKey) New() error {
   resp, err := http.PostForm(
      "https://wbd-api.arkoselabs.com/fc/gt2/public_key/"+arkose_site_key,
      url.Values{
         "public_key": {arkose_site_key},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
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

type route_include struct {
   Attributes struct {
      AirDate       time.Time
      Name          string
      EpisodeNumber int
      SeasonNumber  int
   }
   Id            string
   Relationships *struct {
      Show *struct {
         Data struct {
            Id string
         }
      }
   }
}
