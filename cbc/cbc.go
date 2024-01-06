package cbc

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func New_Token(username, password string) (*Token, error) {
   address := func() string {
      var b strings.Builder
      b.WriteString("https://login.cbc.radio-canada.ca")
      b.WriteString("/bef1b538-1950-4283-9b27-b096cbc18070")
      b.WriteString("/B2C_1A_ExternalClient_ROPC_Auth/oauth2/v2.0/token")
      return b.String()
   }()
   res, err := http.PostForm(address, url.Values{
      "client_id": {"7f44c935-6542-4ce7-ae05-eb887809741c"},
      "grant_type": {"password"},
      "password": {password},
      "scope": {strings.Join(scope, " ")},
      "username": {username},
   })
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   t := new(Token)
   if err := json.NewDecoder(res.Body).Decode(t); err != nil {
      return nil, err
   }
   return t, nil
}

func New_Catalog_Gem(address string) (*Catalog_Gem, error) {
   // you can also use `phone_android`, but it returns combined number and name:
   // 3. Beauty Hath Strange Power
   req, err := http.NewRequest("GET", "https://services.radio-canada.ca", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "device=web"
   req.URL.Path, err = func() (string, error) {
      p, err := url.Parse(address)
      if err != nil {
         return "", err
      }
      return "/ott/catalog/v2/gem/show" + p.Path, nil
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   gem := new(Catalog_Gem)
   if err := json.NewDecoder(res.Body).Decode(gem); err != nil {
      return nil, err
   }
   return gem, nil
}

type Lineup_Item struct {
   URL string
   Formatted_ID_Media string `json:"formattedIdMedia"`
}

func (c Catalog_Gem) Item() *Lineup_Item {
   for _, content := range c.Content {
      for _, lineup := range content.Lineups {
         for _, item := range lineup.Items {
            if item.URL == c.Selected_URL {
               return &item
            }
         }
      }
   }
   return nil
}

const forwarded_for = "99.224.0.0"

func (p Profile) Write_File(name string) error {
   text, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return err
   }
   return os.WriteFile(name, text, 0666)
}

func Read_Profile(name string) (*Profile, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   pro := new(Profile)
   if err := json.Unmarshal(text, pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Profile struct {
   Claims_Token string `json:"claimsToken"`
}

func (t Token) Profile() (*Profile, error) {
   req, err := http.NewRequest("GET", "https://services.radio-canada.ca", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/ott/subscription/v2/gem/Subscriber/profile"
   req.URL.RawQuery = "device=phone_android"
   req.Header.Set("Authorization", "Bearer " + t.Access_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   pro := new(Profile)
   if err := json.NewDecoder(res.Body).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Token struct {
   Access_Token string
}

const manifest_type = "desktop"

func (p Profile) Media(item *Lineup_Item) (*Media, error) {
   req, err := http.NewRequest("GET", "https://services.radio-canada.ca", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/media/validation/v2"
   req.URL.RawQuery = url.Values{
      "appCode": {"gem"},
      "idMedia": {item.Formatted_ID_Media},
      "manifestType": {manifest_type},
      "output": {"json"},
      // you need this one the first request for a video, but can omit after
      // that. we will just send it every time.
      "tech": {"hls"},
   }.Encode()
   req.Header = http.Header{
      "X-Claims-Token": {p.Claims_Token},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   m := new(Media)
   if err := json.NewDecoder(res.Body).Decode(m); err != nil {
      return nil, err
   }
   if m.Message != "" {
      return nil, errors.New(m.Message)
   }
   m.URL = strings.Replace(m.URL, "[manifestType]", manifest_type, 1)
   return m, nil
}

type Media struct {
   Message string
   URL string
}

var scope = []string{
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/subscriptions.write",
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/toutv-profiling",
   "openid",
}
