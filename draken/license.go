package draken

import "net/http"

type poster struct {
   auth auth_login
   play *playback
}

func (poster) RequestUrl() (string, bool) {
   return "https://client-api.magine.com/api/playback/v1/widevine/license", true
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (p poster) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   magine_accesstoken.set(head)
   head.Set("authorization", "Bearer " + p.auth.v.Token)
   for key, value := range p.play.Headers {
      head.Set(key, value)
   }
   return head, nil
}
