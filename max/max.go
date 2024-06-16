package max

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (w WebAddress) MarshalText() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("https://play.max.com/video/watch/")
	b.WriteString(w.video_id)
	b.WriteByte('/')
	b.WriteString(w.edit_id)
	return b.Bytes(), nil
}

type WebAddress struct {
	video_id string
	edit_id  string
}

func (w *WebAddress) UnmarshalText(text []byte) error {
	s := string(text)
	if !strings.Contains(s, "/video/watch/") {
		return errors.New("/video/watch/ not found")
	}
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "play.max.com")
	s = strings.TrimPrefix(s, "/video/watch/")
	var found bool
	w.video_id, w.edit_id, found = strings.Cut(s, "/")
	if !found {
		return errors.New("/ not found")
	}
	return nil
}

type default_routes struct {
	Data struct {
		Attributes struct {
			Url WebAddress
		}
	}
	Included []route_include
}

func (d default_token) routes(path string) (*default_routes, error) {
	address := func() string {
		var b strings.Builder
		b.WriteString("https://default.any-")
		b.WriteString(home_market)
		b.WriteString(".prd.api.discomax.com/cms/routes")
		b.WriteString(path)
		return b.String()
	}()
	req, err := http.NewRequest("", address, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = url.Values{
		"include": {"default"},
		// this is not required, but results in a smaller response
		"page[items.size]": {"1"},
	}.Encode()
	req.Header.Set("authorization", "Bearer "+d.Data.Attributes.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	route := new(default_routes)
	err = json.NewDecoder(resp.Body).Decode(route)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (d default_routes) video() (*route_include, bool) {
	for _, include := range d.Included {
		if include.Id == d.Data.Attributes.Url.video_id {
			return &include, true
		}
	}
	return nil, false
}

func (d default_routes) Show() string {
	if v, ok := d.video(); ok {
		if v.Attributes.SeasonNumber >= 1 {
			for _, include := range d.Included {
				if include.Id == v.Relationships.Show.Data.Id {
					return include.Attributes.Name
				}
			}
		}
	}
	return ""
}

func (d default_routes) Season() int {
	if v, ok := d.video(); ok {
		return v.Attributes.SeasonNumber
	}
	return 0
}

func (d default_routes) Episode() int {
	if v, ok := d.video(); ok {
		return v.Attributes.EpisodeNumber
	}
	return 0
}

func (d default_routes) Title() string {
	if v, ok := d.video(); ok {
		return v.Attributes.Name
	}
	return ""
}

func (d default_routes) Year() int {
	if v, ok := d.video(); ok {
		return v.Attributes.AirDate.Year()
	}
	return 0
}

func (d default_token) playback(web WebAddress) (*playback, error) {
	body, err := func() ([]byte, error) {
		var p playback_request
		p.ConsumptionType = "streaming"
		p.EditId = web.edit_id
		return json.Marshal(p)
	}()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", "https://default.any-any.prd.api.discomax.com",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.URL.Path = func() string {
		var b bytes.Buffer
		b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
		b.WriteString("/playbackInfo")
		return b.String()
	}()
	req.Header.Set("authorization", "Bearer "+d.Data.Attributes.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var b bytes.Buffer
		resp.Write(&b)
		return nil, errors.New(b.String())
	}
	play := new(playback)
	err = json.NewDecoder(resp.Body).Decode(play)
	if err != nil {
		return nil, err
	}
	return play, nil
}

type playback_request struct {
	AppBundle            string `json:"appBundle"`            // required
	ApplicationSessionId string `json:"applicationSessionId"` // required
	Capabilities         struct {
		Manifests struct {
			Formats struct {
				Dash struct{} `json:"dash"` // required
			} `json:"formats"` // required
		} `json:"manifests"` // required
	} `json:"capabilities"` // required
	ConsumptionType string `json:"consumptionType"`
	DeviceInfo      struct {
		Player struct {
			MediaEngine struct {
				Name    string `json:"name"`    // required
				Version string `json:"version"` // required
			} `json:"mediaEngine"` // required
			PlayerView struct {
				Height int `json:"height"` // required
				Width  int `json:"width"`  // required
			} `json:"playerView"` // required
			Sdk struct {
				Name    string `json:"name"`    // required
				Version string `json:"version"` // required
			} `json:"sdk"` // required
		} `json:"player"` // required
	} `json:"deviceInfo"` // required
	EditId            string   `json:"editId"`
	FirstPlay         bool     `json:"firstPlay"`         // required
	Gdpr              bool     `json:"gdpr"`              // required
	PlaybackSessionId string   `json:"playbackSessionId"` // required
	UserPreferences   struct{} `json:"userPreferences"`   // required
}

type public_key struct {
	Token string
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

func (p *public_key) New() error {
	resp, err := http.PostForm(
		"https://wbd-api.arkoselabs.com/fc/gt2/public_key/"+arkose_site_key,
		url.Values{
			"public_key": {arkose_site_key},
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(p)
}

func (playback) WrapRequest(b []byte) ([]byte, error) {
	return b, nil
}

func (playback) UnwrapResponse(b []byte) ([]byte, error) {
	return b, nil
}

func (playback) RequestHeader() (http.Header, error) {
	return http.Header{}, nil
}

type playback struct {
	Drm struct {
		Schemes struct {
			Widevine struct {
				LicenseUrl string
			}
		}
	}
	Manifest struct {
		Url string
	}
}

func (p playback) RequestUrl() (string, bool) {
	return p.Drm.Schemes.Widevine.LicenseUrl, true
}

type route_include struct {
	Attributes struct {
		AirDate       time.Time
		Name          string
		EpisodeNumber int
		SeasonNumber  int
	}
	Id            string
	Relationships *struct {
		Show *struct {
			Data struct {
				Id string
			}
		}
	}
}
