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
   "strconv"
   "strings"
   "time"
)

// hard geo block
func (at AppToken) Session(content_id string) (*SessionToken, error) {
   req, err := http.NewRequest("", "https://www.intl.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = at.v.Encode()
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v3.0/androidphone/irdeto-control")
      b.WriteString("/anonymous-session-token.json")
      return b.String()
   }()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   session := new(SessionToken)
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   session.Url += content_id
   return session, nil
}

type SessionToken struct {
   Url string
   LsSession string `json:"ls_session"`
}

func (SessionToken) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (s SessionToken) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("authorization", "Bearer " + s.LsSession)
   return head, nil
}

func (s SessionToken) RequestUrl() (string, bool) {
   return s.Url, true
}

func (SessionToken) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

// hard geo block
func location(content_id string, query url.Values) (string, error) {
   req, err := http.NewRequest("", "https://link.theplatform.com", nil)
   if err != nil {
      return "", err
   }
   req.URL.Path = func() string {
      b := []byte("/s/")
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_id...)
      return string(b)
   }()
   req.URL.RawQuery = query.Encode()
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusFound {
      var s struct {
         Description string
      }
      json.NewDecoder(resp.Body).Decode(&s)
      return "", errors.New(s.Description)
   }
   return resp.Header.Get("Location"), nil
}

type AppToken struct {
   v url.Values
}

func DashCenc(content_id string) (string, error) {
   query := url.Values{
      "assetTypes": {"DASH_CENC_PRECON"},
      "formats": {"MPEG-DASH"},
   }
   return location(content_id, query)
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

type number int

type VideoItem struct {
   AirDateIso time.Time `json:"_airDateISO"`
   EpisodeNum number
   Label string
   SeasonNum number
   SeriesTitle string
   MediaType string
}

func (n *number) UnmarshalText(text []byte) error {
   if len(text) >= 1 {
      i, err := strconv.Atoi(string(text))
      if err != nil {
         return err
      }
      *n = number(i)
   }
   return nil
}

func (at *AppToken) New() error {
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
   at.v = url.Values{
      "at": {base64.StdEncoding.EncodeToString(dst)},
   }
   return nil
}

// com.cbs.ca_15.0.28
const app_secret = "c0b1d5d6ed27a3f6"

// hard geo block
func (at AppToken) Item(content_id string) (*VideoItem, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
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
   req.URL.RawQuery = at.v.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var video struct {
      ItemList []VideoItem
   }
   err = json.NewDecoder(resp.Body).Decode(&video)
   if err != nil {
      return nil, err
   }
   return &video.ItemList[0], nil
}
