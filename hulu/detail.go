package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
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
