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
)

func (SessionToken) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (SessionToken) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

type AppToken struct {
   values url.Values
}

func (n Number) MarshalText() ([]byte, error) {
   return strconv.AppendInt(nil, int64(n), 10), nil
}

func (n *Number) UnmarshalText(text []byte) error {
   if len(text) >= 1 {
      i, err := strconv.ParseInt(string(text), 10, 64)
      if err != nil {
         return err
      }
      *n = Number(i)
   }
   return nil
}

type Number int64

// 15.0.28
func (at *AppToken) ComCbsApp() error {
   return at.New("a624d7b175f5626b")
}

// 15.0.28
func (at *AppToken) ComCbsCa() error {
   return at.New("c0b1d5d6ed27a3f6")
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

func (s *SessionToken) RequestUrl() (string, bool) {
   return s.Url, true
}

type SessionToken struct {
   LsSession string `json:"ls_session"`
   Url string
}

func (s *SessionToken) RequestHeader() (http.Header, error) {
   head := http.Header{
      "authorization": {"Bearer " + s.LsSession},
      "content-type": {"application/x-protobuf"},
   }
   return head, nil
}

///

// must use app token and IP address for US
func (at AppToken) Session(content_id string) (*SessionToken, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v3.1/androidphone/irdeto-control")
      b.WriteString("/anonymous-session-token.json")
      return b.String()
   }()
   at.values.Set("contentId", content_id)
   req.URL.RawQuery = at.values.Encode()
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
   session := &SessionToken{}
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   return session, nil
}

func (at *AppToken) New(app_secret string) error {
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
   at.values = url.Values{
      "at": {base64.StdEncoding.EncodeToString(dst)},
   }
   return nil
}
