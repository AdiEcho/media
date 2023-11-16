package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "path"
)

type Authenticate struct {
   Raw []byte
   Value struct {
      Data struct {
         User_Token string
      }
   }
}

func Living_Room(email, password string) (*Authenticate, error) {
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
      return nil, errors.New(res.Status)
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
func (a Authenticate) Details(d *Deep_Link) (*Details, error) {
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
   req.URL.RawQuery = "user_token=" + a.Value.Data.User_Token
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

func (d Details) Series() string {
   return d.Series_Name
}

func (d Details) Title() string {
   return d.Episode_Name
}

type Details struct {
   Episode_Name string
   Episode_Number int64
   Season_Number int64
   Series_Name string
}

func (d Details) Season() (int64, error) {
   return d.Season_Number, nil
}

func (d Details) Episode() (int64, error) {
   return d.Episode_Number, nil
}

type Deep_Link struct {
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

func (a Authenticate) Deep_Link(watch ID) (*Deep_Link, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {watch.s},
      "namespace": {"entity"},
      "user_token": {a.Value.Data.User_Token},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   link := new(Deep_Link)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}
