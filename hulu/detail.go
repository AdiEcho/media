package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

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
