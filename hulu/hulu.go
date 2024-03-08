package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "path"
   "strconv"
   "strings"
)

func LivingRoom(email, password string) (*Authenticate, error) {
   res, err := http.PostForm(
      "https://auth.hulu.com/v2/livingroom/password/authenticate", url.Values{
         "friendly_name": {"!"},
         "password": {password},
         "serial_number": {"!"},
         "user_email": {email},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var auth Authenticate
   auth.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Raw, &a.Value)
}

type DeepLink struct {
   EAB_ID string
}

type ID struct {
   s string
}

func (i ID) String() string {
   return i.s
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
func (i *ID) Set(s string) error {
   i.s = path.Base(s)
   return nil
}
func (a Authenticate) DeepLink(watch ID) (*DeepLink, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {watch.s},
      "namespace": {"entity"},
   }.Encode()
   req.Header.Set("Authorization", "Bearer " + a.Value.Data.User_Token)
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
   link := new(DeepLink)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}

func (a Authenticate) Details(d *DeepLink) (*Details, error) {
   body, err := func() ([]byte, error) {
      m := map[string][]string{
         "eabs": {d.EAB_ID},
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://guide.hulu.com/guide/details", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + a.Value.Data.User_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      Items []Details
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   if len(s.Items) == 0 {
      return nil, errors.New("items length is zero")
   }
   return &s.Items[0], nil
}

type Details struct {
   Episode_Name string
   Episode_Number int
   Headline string
   Premiere_Date string
   Season_Number int
   Series_Name string
}

func (Details) Owner() (string, bool) {
   return "", false
}

func (d Details) Show() (string, bool) {
   return d.Series_Name, d.Series_Name != ""
}

func (d Details) Season() (string, bool) {
   if d.Season_Number >= 1 {
      return strconv.Itoa(d.Season_Number), true
   }
   return "", false
}

func (d Details) Episode() (string, bool) {
   if d.Episode_Number >= 1 {
      return strconv.Itoa(d.Episode_Number), true
   }
   return "", false
}

func (d Details) Title() (string, bool) {
   if d.Episode_Name != "" {
      return d.Episode_Name, true
   }
   return d.Headline, true
}

func (d Details) Year() (string, bool) {
   if d.Episode_Name != "" {
      return "", false
   }
   year, _, _ := strings.Cut(d.Premiere_Date, "-")
   return year, true
}

type codec_value struct {
   Height int `json:"height,omitempty"`
   Level   string `json:"level,omitempty"`
   Profile string `json:"profile,omitempty"`
   Type    string `json:"type"`
   Width int `json:"width,omitempty"`
}

type drm_value struct {
   Security_Level string `json:"security_level"`
   Type          string `json:"type"`
   Version       string `json:"version"`
}

type playlist_request struct {
   Content_EAB_ID   string `json:"content_eab_id"`
   Deejay_Device_ID int    `json:"deejay_device_id"`
   Unencrypted    bool   `json:"unencrypted"`
   Version        int    `json:"version"`
   Playback       struct {
      Audio struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"audio"`
      Video   struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
      DRM struct {
         Selection_Mode string `json:"selection_mode"`
         Values []drm_value `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         Selection_Mode string `json:"selection_mode"`
         Values []segment_value `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
   } `json:"playback"`
}

type segment_value struct {
   Encryption struct {
      Mode string `json:"mode"`
      Type string `json:"type"`
   } `json:"encryption"`
   Type string `json:"type"`
}

type Authenticate struct {
   Raw []byte
   Value struct {
      Data struct {
         User_Token string
      }
   }
}

