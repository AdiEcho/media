package rtbf

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (g GigyaLogin) Entitlement(page *AuvioPage) (*Entitlement, error) {
	req, err := http.NewRequest("", "https://exposure.api.redbee.live", nil)
	if err != nil {
		return nil, err
	}
	req.URL.Path = func() string {
		var b strings.Builder
		b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement/")
		b.WriteString(page.Content.AssetId)
		b.WriteString("/play")
		return b.String()
	}()
	req.Header = http.Header{
		"x-forwarded-for": {"91.90.123.17"},
		"authorization":   {"Bearer " + g.SessionToken},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var b strings.Builder
		res.Write(&b)
		return nil, errors.New(b.String())
	}
	title := new(Entitlement)
	err = json.NewDecoder(res.Body).Decode(title)
	if err != nil {
		return nil, err
	}
	return title, nil
}

type Entitlement struct {
	AssetId   string
	PlayToken string
	Formats   []struct {
		Format       string
		MediaLocator string
	}
}

func (e Entitlement) dash() (string, bool) {
	for _, format := range e.Formats {
		if format.Format == "DASH" {
			return format.MediaLocator, true
		}
	}
	return "", false
}

func (e Entitlement) RequestUrl() (string, bool) {
	var u url.URL
	u.Host = "rbm-rtbf.live.ott.irdeto.com"
	u.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
	u.Scheme = "https"
	u.RawQuery = url.Values{
		"contentId":  {e.AssetId},
		"ls_session": {e.PlayToken},
	}.Encode()
	return u.String(), true
}

func (Entitlement) RequestHeader() (http.Header, error) {
	h := make(http.Header)
	h.Set("content-type", "application/x-protobuf")
	return h, nil
}

func (Entitlement) WrapRequest(b []byte) ([]byte, error) {
	return b, nil
}

func (Entitlement) UnwrapResponse(b []byte) ([]byte, error) {
	return b, nil
}

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (a *AccountLogin) New(id, password string) error {
	body := url.Values{
		"APIKey":   {api_key},
		"loginID":  {id},
		"password": {password},
	}.Encode()
	req, err := http.NewRequest(
		"POST", "https://login.auvio.rtbf.be/accounts.login",
		strings.NewReader(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(a)
	if err != nil {
		return err
	}
	if v := a.ErrorMessage; v != "" {
		return errors.New(v)
	}
	return nil
}

type AccountLogin struct {
	ErrorMessage string
	SessionInfo  struct {
		CookieValue string
	}
}

func (a *AccountLogin) Unmarshal(text []byte) error {
	return json.Unmarshal(text, a)
}

func (a AccountLogin) Marshal() ([]byte, error) {
	return json.Marshal(a)
}
func (a AccountLogin) Token() (*WebToken, error) {
	body := url.Values{
		"APIKey": {api_key},
		// from /accounts.login
		"login_token": {a.SessionInfo.CookieValue},
	}.Encode()
	req, err := http.NewRequest(
		"POST", "https://login.auvio.rtbf.be/accounts.getJWT",
		strings.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var web WebToken
	err = json.NewDecoder(res.Body).Decode(&web)
	if err != nil {
		return nil, err
	}
	if v := web.ErrorMessage; v != "" {
		return nil, errors.New(v)
	}
	return &web, nil
}

type GigyaLogin struct {
	SessionToken string
}

type WebToken struct {
	ErrorMessage string
	IdToken      string `json:"id_token"`
}

func (w WebToken) Login() (*GigyaLogin, error) {
	body, err := func() ([]byte, error) {
		var s struct {
			Device struct {
				DeviceId string `json:"deviceId"`
				Type     string `json:"type"`
			} `json:"device"`
			JWT string `json:"jwt"`
		}
		s.Device.Type = "WEB"
		s.JWT = w.IdToken
		return json.Marshal(s)
	}()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", "https://exposure.api.redbee.live", bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin"
	req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	login := new(GigyaLogin)
	err = json.NewDecoder(res.Body).Decode(login)
	if err != nil {
		return nil, err
	}
	return login, nil
}
