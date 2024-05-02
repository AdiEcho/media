package roku

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type AccountToken struct {
	AuthToken string
}

type playback struct {
	DRM struct {
		Widevine struct {
			LicenseServer string
		}
	}
	URL string
}

func (a AccountToken) playback(roku_id string) (*playback, error) {
	body, err := func() ([]byte, error) {
		m := map[string]string{
			"mediaFormat": "DASH",
			"providerId":  "rokuavod",
			"rokuId":      roku_id,
		}
		return json.Marshal(m)
	}()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", "https://googletv.web.roku.com/api/v3/playback",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Content-Type":         {"application/json"},
		"User-Agent":           {user_agent},
		"X-Roku-Content-Token": {a.AuthToken},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	play := new(playback)
	err = json.NewDecoder(res.Body).Decode(play)
	if err != nil {
		return nil, err
	}
	return play, nil
}

func (a *AccountToken) New() error {
	req, err := http.NewRequest(
		"GET", "https://googletv.web.roku.com/api/v1/account/token", nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("user-agent", user_agent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(a)
}

const user_agent = "trc-googletv; production; 0"

func (playback) RequestHeader() (http.Header, error) {
	return http.Header{}, nil
}

func (p playback) RequestUrl() (string, bool) {
	return p.DRM.Widevine.LicenseServer, true
}

func (playback) WrapRequest(b []byte) ([]byte, error) {
	return b, nil
}

func (playback) UnwrapResponse(b []byte) ([]byte, error) {
	return b, nil
}
