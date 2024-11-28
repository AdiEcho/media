package kanopy

import (
   "net/http"
   "net/url"
)

func (web_token) RequestUrl() (string, bool) {
   var u url.URL
   u.Scheme = "https"
   u.Host = "www.kanopy.com"
   u.Path = "/kapi/licenses/widevine/1732823808506000167-0"
   return u.String(), true
}

func (web_token) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (web_token) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (w *web_token) RequestHeader() (http.Header, error) {
   h := http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {"!/!/!/!"},
   }
   return h, nil
}
