package ctv

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
   "time"
)

type Date struct {
   Time time.Time
}

func (d Date) MarshalText() ([]byte, error) {
   return d.Time.AppendFormat(nil, time.DateOnly), nil
}

func (d *Date) UnmarshalText(text []byte) error {
   var err error
   d.Time, err = time.Parse(time.DateOnly, string(text))
   if err != nil {
      return err
   }
   return nil
}

type Namer struct {
   Media *MediaContent
}

func (n Namer) Episode() int {
   return n.Media.Episode
}

func (n Namer) Season() int {
   return n.Media.Season.Number
}

func (n Namer) Show() string {
   if v := n.Media.Media; v.Type == "series" {
      return v.Name
   }
   return ""
}

func (n Namer) Year() int {
   return n.Media.BroadcastDate.Time.Year()
}

func (n Namer) Title() string {
   if strings.HasSuffix(n.Media.Name, ")") {
      return n.Media.Name[:len(n.Media.Name)-len(" (9999)")]
   }
   return n.Media.Name
}

type Poster struct{}

func (Poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (Poster) RequestUrl() (string, bool) {
   return "https://license.9c9media.ca/widevine", true
}

func (Poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (r ResolvePath) Axis() (*AxisContent, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         OperationName string `json:"operationName"`
         Query         string `json:"query"`
         Variables     struct {
            Id string `json:"id"`
         } `json:"variables"`
      }
      s.OperationName = "axisContent"
      s.Query = graphql_compact(query_axis)
      s.Variables.Id = r.id()
      return json.MarshalIndent(s, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var data struct {
      Data struct {
         AxisContent AxisContent
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &data.Data.AxisContent, nil
}

const query_axis = `
query($id: ID!) {
   axisContent(id: $id) {
      axisId
      axisPlaybackLanguages {
         ... on AxisPlayback {
            destinationCode
         }
      }
   }
}
`

type AxisContent struct {
   AxisId                int64
   AxisPlaybackLanguages []struct {
      DestinationCode string
   }
}

func (a *AxisContent) Media() (*MediaContent, error) {
   req, err := http.NewRequest("", "https://capi.9c9media.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/destinations/")
      b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, a.AxisId, 10)
      return string(b)
   }()
   req.URL.RawQuery = "$include=[ContentPackages,Media,Season]"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var media MediaContent
   media.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &media, nil
}

type MediaContent struct {
   BroadcastDate   Date
   ContentPackages []struct {
      Id int64
   }
   Episode int
   Media   struct {
      Name string
      Type string
   }
   Name   string
   Season struct {
      Number int
   }
   Raw []byte `json:"-"`
}

func (m *MediaContent) Unmarshal() error {
   return json.Unmarshal(m.Raw, m)
}

// wikipedia.org/wiki/Geo-blocking
func (a *AxisContent) Manifest(media *MediaContent) (string, error) {
   req, err := http.NewRequest("", "https://capi.9c9media.com", nil)
   if err != nil {
      panic(err)
   }
   req.URL.Path = func() string {
      b := []byte("/destinations/")
      b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/playback/contents/"...)
      b = strconv.AppendInt(b, a.AxisId, 10)
      b = append(b, "/contentPackages/"...)
      b = strconv.AppendInt(b, media.ContentPackages[0].Id, 10)
      b = append(b, "/manifest.mpd"...)
      return string(b)
   }()
   req.URL.RawQuery = "action=reference"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return "", errors.New(b.String())
   }
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return "", err
   }
   return strings.Replace(string(text), "/best/", "/ultimate/", 1), nil
}
