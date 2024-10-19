package ctv

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

// server requires GET instead of POST, and
// `extensions.persistedQuery` instead of `query`
func (r *ResolvePath) Axis() (*AxisContent, error) {
   req, err := http.NewRequest(
      "", "https://www.ctv.ca/space-graphql/apq/graphql", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "extensions": {`{"persistedQuery":{"version":1,"sha256Hash":"d6e75de9b5836cd6305c98c8d2411e336f59eb12f095a61f71d454f3fae2ecda"}}`},
      "operationName": {"axisContent"},
      // fmt.Sprintf(`{"id": %q}`, r.id()),
      "variables": {`{"id":"contentid/axis-content-2968346","subscriptions":["CTV","CTV_DRAMA","CTV_COMEDY","CTV_LIFE","CTV_SCIFI","CTV_THROWBACK","CTV_MOVIES","CTV_MTV","CTV_MUCH","DISCOVERY","DISCOVERY_SCIENCE","DISCOVERY_VELOCITY","INVESTIGATION_DISCOVERY","ANIMAL_PLANET","E_NOW"],"maturity":"ADULT","language":"ENGLISH","authenticationState":"UNAUTH","playbackLanguage":"ENGLISH"}`},
   }.Encode()
   // you need this for the first request, then can omit
   req.Header["Graphql-Client-Platform"] = []string{"entpay_web"}
   
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Newrelic"] = []string{"eyJ2IjpbMCwxXSwiZCI6eyJ0eSI6IkJyb3dzZXIiLCJhYyI6IjI2MTQzMjciLCJhcCI6IjE4MzUwMDQwMjEiLCJpZCI6ImM1ZWViZTk5ODhjNDEwMWYiLCJ0ciI6ImMwOWQwZDk2ZGQ5ZjI2NzM2NjE4OTJmNTE0NzNmMDAwIiwidGkiOjE3MjkzMTkyOTY4NTIsInRrIjoiMjM3ODU3NSJ9fQ=="}
   req.Header["Referer"] = []string{"https://www.ctv.ca/movies/ingrid-goes-west/ingrid-goes-west"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["Traceparent"] = []string{"00-c09d0d96dd9f2673661892f51473f000-c5eebe9988c4101f-01"}
   req.Header["Tracestate"] = []string{"2378575@nr=0-1-2614327-1835004021-c5eebe9988c4101f----1729319296852"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0"}
   
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var resp_body struct {
      Data struct {
         AxisContent AxisContent
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return nil, err
   }
   return &resp_body.Data.AxisContent, nil
}

func (a Address) Resolve() (*ResolvePath, error) {
   var body struct {
      OperationName string `json:"operationName"`
      Query         string `json:"query"`
      Variables     struct {
         Path string `json:"path"`
      } `json:"variables"`
   }
   body.OperationName = "resolvePath"
   body.Query = graphql_compact(query_resolve)
   body.Variables.Path = a.Path
   data, err := json.MarshalIndent(body, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(data),
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
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var resp_body struct {
      Data struct {
         ResolvedPath *struct {
            LastSegment struct {
               Content ResolvePath
            }
         }
      }
   }
   err = json.Unmarshal(data, &resp_body)
   if err != nil {
      return nil, err
   }
   if v := resp_body.Data.ResolvedPath; v != nil {
      return &v.LastSegment.Content, nil
   }
   return nil, errors.New(string(data))
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

func (d *Date) MarshalText() ([]byte, error) {
   return d.Time.AppendFormat(nil, time.DateOnly), nil
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

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               id
               ... on AxisMedia {
                  firstPlayableContent {
                     id
                  }
               }
            }
         }
      }
   }
}
`

func (r *ResolvePath) id() string {
   if r.FirstPlayableContent != nil {
      return r.FirstPlayableContent.Id
   }
   return r.Id
}

type ResolvePath struct {
   Id                   string
   FirstPlayableContent *struct {
      Id string
   }
}

type Address struct {
   Path string
}

// https://www.ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   a.Path = strings.TrimPrefix(s, "ctv.ca")
   return nil
}

type Date struct {
   Time time.Time
}

type Poster struct{}

func (a *Address) String() string {
   return a.Path
}
