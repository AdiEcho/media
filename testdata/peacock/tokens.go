package peacock

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a *auth_tokens) New() error {
   const body = `
   {
     "auth": {
       "authScheme": "MESSO",
       "authIssuer": "NOWTV",
       "personaId": "772bf4eb-4b98-4466-8922-d8c42f900e9c",
       "provider": "NBCU",
       "providerTerritory": "US",
       "proposition": "NBCUOTT"
     },
     "device": {
       "type": "COMPUTER",
       "platform": "PC"
     }
   }
   `
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "ovp.peacocktv.com"
   req.URL.Path = "/auth/tokens"
   req.URL.Scheme = "https"
   req.Header["Cookie"] = []string{
      "idsession=E0-AAAADBaA4I4bW1mAUJgD8HiYL3/d97POed3b0DcxA/VMs87+3JJYZA6V23xO2DTE9hgF1p2EUh6C5RWt0snSpc8+XZnCOetS3GsBlae8zfESQbonJTONmZxa4aypEqgYgWNW",
   }
   req.Header["Content-Type"] = []string{"application/vnd.tokens.v1+json"}
   req.Body = io.NopCloser(strings.NewReader(body))
   req.Header["X-Skyott-Device"] = []string{"COMPUTER"}
   req.Header["X-Skyott-Platform"] = []string{"PC"}
   req.Header["X-Skyott-Proposition"] = []string{"NBCUOTT"}
   req.Header["X-Skyott-Provider"] = []string{"NBCU"}
   req.Header["X-Skyott-Territory"] = []string{"US"}
   req.Header["X-Sky-Signature"] = []string{
      sign(req.Method, req.URL.Path, req.Header, []byte(body)),
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

type auth_tokens struct {
   UserToken string
}
