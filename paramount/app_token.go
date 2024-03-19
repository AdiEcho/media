package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (at AppToken) Item(content_id string) (chan Item, error) {
   req, err := http.NewRequest("GET", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v2.0/androidphone/video/cid/")
      b.WriteString(content_id)
      b.WriteString(".json")
      return b.String()
   }()
   // this needs to be encoded
   req.URL.RawQuery = "at=" + url.QueryEscape(string(at))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var video struct {
      ItemList []Item
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   channel := make(chan Item)
   go func() {
      for _, item := range video.ItemList {
         channel <- item
      }
      close(channel)
   }()
   return channel, nil
}

type AppToken string

func (at *AppToken) with(app_secret string) error {
   key, err := hex.DecodeString(secret_key)
   if err != nil {
      return err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return err
   }
   var src []byte
   src = append(src, '|')
   src = append(src, app_secret...)
   src = pad(src)
   var iv [aes.BlockSize]byte
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   var dst []byte
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   *at = AppToken(base64.StdEncoding.EncodeToString(dst))
   return nil
}

func (at *AppToken) New() error {
   app := app_details{"12.0.44", 211204450}
   return at.with(app_secrets[app])
}

func (at AppToken) Session(content_id string) (*SessionToken, error) {
   req, err := http.NewRequest("GET", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json"
   req.URL.RawQuery = url.Values{
      // this needs to be encoded
      "at": {string(at)},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   session := new(SessionToken)
   if err := json.NewDecoder(res.Body).Decode(session); err != nil {
      return nil, err
   }
   session.URL += content_id
   return session, nil
}
