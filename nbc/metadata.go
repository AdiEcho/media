package nbc

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "time"
)

func (m Metadata) On_Demand() (*On_Demand, error) {
   body, err := func() ([]byte, error) {
      var v video_request
      v.Device = "android"
      v.Device_ID = "android"
      v.External_Advertiser_ID = "NBC"
      v.MPX.Account_ID = m.MPX_Account_ID
      return json.MarshalIndent(v, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "http://access-cloudpath.media.nbcuni.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/access/vod/nbcuniversal/" + m.MPX_GUID
   req.Header = http.Header{
      "Authorization": {authorization(nil)},
      "Content-Type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   on := new(On_Demand)
   if err := json.NewDecoder(res.Body).Decode(on); err != nil {
      return nil, err
   }
   return on, nil
}

func New_Metadata(guid int64) (*Metadata, error) {
   body, err := func() ([]byte, error) {
      var p page_request
      p.Query = graphQL_compact(query)
      p.Variables.App = "nbc"
      p.Variables.Name = fmt.Sprint(guid)
      p.Variables.One_App = true
      p.Variables.Platform = "android"
      p.Variables.Type = "VIDEO"
      return json.MarshalIndent(p, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var page struct {
      Data struct {
         Bonanza_Page struct {
            Metadata *Metadata
         } `json:"bonanzaPage"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   if page.Data.Bonanza_Page.Metadata == nil {
      return nil, fmt.Errorf(".data.bonanzaPage.metadata")
   }
   return page.Data.Bonanza_Page.Metadata, nil
}

type Metadata struct {
   Air_Date string `json:"airDate"`
   Episode_Number *int64 `json:"episodeNumber,string"`
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Season_Number *int64 `json:"seasonNumber,string"`
   Secondary_Title string `json:"secondaryTitle"`
   Series_Short_Title *string `json:"seriesShortTitle"`
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, m.Air_Date)
}

func (m Metadata) Episode() (int64, error) {
   return m.Episode_Number, nil
}

func (m Metadata) Season() (int64, error) {
   return m.Season_Number, nil
}

func (m Metadata) Series() (string, bool) {
   return m.Series_Short_Title, true
}

func (m Metadata) Title() string {
   return m.Secondary_Title
}
