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

// YOU CANNOT USE ANONYMOUS QUERY!
const query_axis = `
query axisContent($id: ID!) {
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

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
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

func (a *Address) String() string {
   return a.Path
}

func (a Address) Resolve() (*ResolvePath, error) {
   var body struct {
      Query         string `json:"query"`
      Variables     struct {
         Path string `json:"path"`
      } `json:"variables"`
   }
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

// hard geo block
func (a *AxisContent) Manifest(media *MediaContent) (string, error) {
   req, err := http.NewRequest("", "https://capi.9c9media.com", nil)
   if err != nil {
      return "", err
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
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return "", err
   }
   return strings.Replace(string(data), "/best/", "/ultimate/", 1), nil
}

type AxisContent struct {
   AxisId                int64
   AxisPlaybackLanguages []struct {
      DestinationCode string
   }
}

type Date struct {
   Time time.Time
}

func (d *Date) UnmarshalText(text []byte) error {
   var err error
   d.Time, err = time.Parse(time.DateOnly, string(text))
   if err != nil {
      return err
   }
   return nil
}

func (d *Date) MarshalText() ([]byte, error) {
   return d.Time.AppendFormat(nil, time.DateOnly), nil
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
}

func (m *MediaContent) Unmarshal(data []byte) error {
   return json.Unmarshal(data, m)
}

func (*MediaContent) Marshal(axis *AxisContent) ([]byte, error) {
   req, err := http.NewRequest("", "https://capi.9c9media.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/destinations/")
      b = append(b, axis.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, axis.AxisId, 10)
      return string(b)
   }()
   req.URL.RawQuery = "$include=[ContentPackages,Media,Season]"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Namer struct {
   Media MediaContent
}

func (n *Namer) Season() int {
   return n.Media.Season.Number
}

func (n *Namer) Show() string {
   if v := n.Media.Media; v.Type == "series" {
      return v.Name
   }
   return ""
}

func (n *Namer) Year() int {
   return n.Media.BroadcastDate.Time.Year()
}

func (n *Namer) Episode() int {
   return n.Media.Episode
}

func (n *Namer) Title() string {
   if strings.HasSuffix(n.Media.Name, ")") {
      return n.Media.Name[:len(n.Media.Name)-len(" (9999)")]
   }
   return n.Media.Name
}

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

func (r *ResolvePath) Axis() (*AxisContent, error) {
   var body struct {
      Query         string `json:"query"`
      Variables     struct {
         Id string `json:"id"`
      } `json:"variables"`
   }
   body.Query = graphql_compact(query_axis)
   body.Variables.Id = r.id()
   data, err := json.Marshal(body)
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
   var resp_body struct {
      Data struct {
         AxisContent AxisContent
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return nil, err
   }
   if v := resp_body.Errors; len(v) >= 1 {
      return nil, errors.New(v[0].Message)
   }
   return &resp_body.Data.AxisContent, nil
}

type Wrapper struct{}

func (Wrapper) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      "https://license.9c9media.ca/widevine", "application/x-protobuf",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
