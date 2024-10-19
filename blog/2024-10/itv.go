package itv

import (
   "net/http"
   "net/url"
)

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "itvpnp.live.ott.irdeto.com"
   u.Path = "/Widevine/getlicense"
   u.RawQuery = "AccountId=itvpnp"
   u.Scheme = "https"
   return u.String(), true
}

type poster struct{}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}
