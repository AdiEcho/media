package itv

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "strings"
)

func (Client) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "itvpnp.live.ott.irdeto.com"
   u.Path = "/Widevine/getlicense"
   u.RawQuery = "AccountId=itvpnp"
   u.Scheme = "https"
   return u.String(), true
}

type Client struct{}

func (Client) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Client) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (Client) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

func (p *Playlist) Resolution1080() (string, bool) {
   for _, file := range p.Playlist.Video.MediaFiles {
      if file.Resolution == "1080" {
         return file.Href, true
      }
   }
   return "", false
}

func (p *Playlist) Resolution720() (string, bool) {
   for _, file := range p.Playlist.Video.MediaFiles {
      if file.Resolution == "720" {
         return file.Href, true
      }
   }
   return "", false
}

// hard geo block
func (d *DiscoveryTitle) Playlist() (*Playlist, error) {
   var value struct {
      Client struct {
         Id string `json:"id"`
      } `json:"client"`
      VariantAvailability struct {
         Drm         struct {
            MaxSupported string `json:"maxSupported"`
            System       string `json:"system"`
         } `json:"drm"`
         FeatureSet  []string `json:"featureset"`
         PlatformTag string   `json:"platformTag"`
      } `json:"variantAvailability"`
   }
   value.Client.Id = "browser"
   value.VariantAvailability.Drm.MaxSupported = "L3"
   value.VariantAvailability.Drm.System = "widevine"
   // need all these to get 720:
   value.VariantAvailability.FeatureSet = []string{
      "hd",
      "mpeg-dash",
      "single-track",
      "widevine",
   }
   value.VariantAvailability.PlatformTag = "dotcom"
   // 1080p but need L1/SL3000
   // value.VariantAvailability.PlatformTag = "ctv"
   data, err := json.MarshalIndent(value, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", d.LatestAvailableVersion.PlaylistUrl, bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/vnd.itv.vod.playlist.v4+json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   play := &Playlist{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (n Namer) Show() string {
   if n.Discovery.Brand != nil {
      return n.Discovery.Brand.Title
   }
   return ""
}

func (n Namer) Title() string {
   return n.Discovery.Title
}

func (i LegacyId) Discovery() (*DiscoveryTitle, error) {
   req, err := http.NewRequest(
      "", "https://content-inventory.prd.oasvc.itv.com/discovery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "query": {fmt.Sprintf(graphql_compact(query_discovery), i)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Titles []DiscoveryTitle
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if v := value.Errors; len(v) >= 1 {
      return nil, errors.New(v[0].Message)
   }
   return &value.Data.Titles[0], nil
}

func (n Namer) Episode() int64 {
   return n.Discovery.EpisodeNumber
}

func (n Namer) Season() int64 {
   return n.Discovery.SeriesNumber
}

type Namer struct {
   Discovery *DiscoveryTitle
}

type DiscoveryTitle struct {
   LatestAvailableVersion struct {
      PlaylistUrl string
   }
   Brand *struct {
      Title string
   }
   EpisodeNumber int64
   ProductionYear int64
   SeriesNumber int64
   Title string
}

func (n Namer) Year() int64 {
   return n.Discovery.ProductionYear
}

const query_discovery = `
{
   titles(filter: {
      legacyId: %q
   }) {
      ... on Episode {
         seriesNumber
         episodeNumber
      }
      ... on Film {
         productionYear
      }
      brand {
         title
      }
      latestAvailableVersion {
         playlistUrl
      }
      title
   }
}
`

func (i LegacyId) String() string {
   var b strings.Builder
   for index, value := range i {
      if value != "" {
         if index >= 1 {
            b.WriteByte('/')
         }
         b.WriteString(value)
      }
   }
   return b.String()
}

func (i *LegacyId) Set(text string) error {
   var found bool
   (*i)[0], text, found = strings.Cut(text, "a")
   if !found {
      return errors.New(`"a" not found`)
   }
   (*i)[1], (*i)[2], found = strings.Cut(text, "a")
   if !found {
      (*i)[2] = "0001"
   }
   return nil
}

type LegacyId [3]string

type Playlist struct {
   Playlist struct {
      Video struct {
         MediaFiles []struct {
            Href string
            Resolution string
         }
      }
   }
}
