package draken

import "net/http"

type Client struct {
   Login AuthLogin
   Play *Playback
}

func (c *Client) RequestHeader() (http.Header, error) {
   head := http.Header{}
   magine_accesstoken.set(head)
   head.Set("authorization", "Bearer "+c.Login.Token)
   for key, value := range c.Play.Headers {
      head.Set(key, value)
   }
   return head, nil
}

func (*Client) RequestUrl() (string, bool) {
   return "https://client-api.magine.com/api/playback/v1/widevine/license", true
}

func (*Client) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (*Client) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}
