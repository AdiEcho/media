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

type Namer struct {
	M *MediaContent
}

func (n Namer) Episode() int {
	return n.M.Episode
}

func (n Namer) Season() int {
	return n.M.Season.Number
}

func (n Namer) Show() string {
	if v := n.M.Media; v.Type == "series" {
		return v.Name
	}
	return ""
}

func (n Namer) Title() string {
	if n.M.Media.Type == "movie" {
		return n.M.Name[:len(n.M.Name)-len(" (9999)")]
	}
	return n.M.Name
}

func (n Namer) Year() int {
	return n.M.BroadcastDate.T.Year()
}

type MediaManifest struct {
	Content *MediaContent
	URL     string
}

func (d date) MarshalText() ([]byte, error) {
	return d.T.AppendFormat(nil, time.DateOnly), nil
}

func (d *date) UnmarshalText(text []byte) error {
	var err error
	d.T, err = time.Parse(time.DateOnly, string(text))
	if err != nil {
		return err
	}
	return nil
}

func (a AxisContent) Media() (*MediaContent, error) {
	address := func() string {
		b := []byte("https://capi.9c9media.com/destinations/")
		b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
		b = append(b, "/platforms/desktop/contents/"...)
		b = strconv.AppendInt(b, a.AxisId, 10)
		b = append(b, "?$include=[ContentPackages,Media,Season]"...)
		return string(b)
	}()
	res, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	media := new(MediaContent)
	err = json.NewDecoder(res.Body).Decode(media)
	if err != nil {
		return nil, err
	}
	return media, nil
}

type AxisContent struct {
	AxisId                int64
	AxisPlaybackLanguages []struct {
		DestinationCode string
	}
}

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

func (r ResolvePath) Axis() (*AxisContent, error) {
	body, err := func() ([]byte, error) {
		var s struct {
			OperationName string `json:"operationName"`
			Query         string `json:"query"`
			Variables     struct {
				ID string `json:"id"`
			} `json:"variables"`
		}
		s.OperationName = "axisContent"
		s.Query = graphql_compact(query_axis)
		s.Variables.ID = r.id()
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
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var s struct {
		Data struct {
			AxisContent AxisContent
		}
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.AxisContent, nil
}

type date struct {
	T time.Time
}

type MediaContent struct {
	BroadcastDate   date
	ContentPackages []struct {
		ID int64
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

func (m MediaManifest) marshal() ([]byte, error) {
	return json.MarshalIndent(m, "", " ")
}

func (m *MediaManifest) Unmarshal(text []byte) error {
	return json.Unmarshal(text, m)
}

// wikipedia.org/wiki/Geo-blocking
func (a AxisContent) Manifest(m *MediaContent) (*MediaManifest, error) {
	address := func() string {
		b := []byte("https://capi.9c9media.com/destinations/")
		b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
		b = append(b, "/platforms/desktop/playback/contents/"...)
		b = strconv.AppendInt(b, a.AxisId, 10)
		b = append(b, "/contentPackages/"...)
		b = strconv.AppendInt(b, m.ContentPackages[0].ID, 10)
		b = append(b, "/manifest.mpd?action=reference"...)
		return string(b)
	}()
	res, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var b strings.Builder
		res.Write(&b)
		return nil, errors.New(b.String())
	}
	text, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &MediaManifest{m, string(text)}, nil
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
