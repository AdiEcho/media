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

// must use IP address for correct location
func (h *Header) New(content_id string) error {
   req, err := http.NewRequest("", "https://link.theplatform.com", nil)
   if err != nil {
      return err
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
   req.URL.RawQuery = "formats=MPEG-DASH"
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusFound {
      var s struct {
         Description string
      }
      json.NewDecoder(resp.Body).Decode(&s)
      return errors.New(s.Description)
   }
   h.Header = resp.Header
   return nil
}

type Header struct {
   Header http.Header
}

func (h Header) Location() string {
   return h.Header.Get("location")
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

func (h Header) JsonMarshal() ([]byte, error) {
   return json.MarshalIndent(h, "", " ")
}

func (h *Header) Json(text []byte) error {
   return json.Unmarshal(text, h)
}
type VideoItem struct {
   SeriesTitle string
   SeasonNum Number
   EpisodeNum Number
   Label string
   AirDateIso time.Time `json:"_airDateISO"`
   MediaType string
}

// must use app token and IP address for correct location
func (at AppToken) Items(content_id string) (VideoItems, error) {
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
      Error string
      ItemList VideoItems
   }
   err = json.NewDecoder(resp.Body).Decode(&video)
   if err != nil {
      return nil, err
   }
   if video.Error != "" {
      return nil, errors.New(video.Error)
   }
   return video.ItemList, nil
}

type SessionToken struct {
   LsSession string `json:"ls_session"`
   Url string
}

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
   at.v.Set("contentId", content_id)
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
   session := &SessionToken{}
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   return session, nil
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

func (v *VideoItem) Json(text []byte) error {
   return json.Unmarshal(text, v)
}

func (v VideoItem) JsonMarshal() ([]byte, error) {
   return json.MarshalIndent(v, "", " ")
}

// 15.0.28
func (at *AppToken) ComCbsApp() error {
   return at.New("a624d7b175f5626b")
}

// 15.0.28
func (at *AppToken) ComCbsCa() error {
   return at.New("c0b1d5d6ed27a3f6")
}

func (v VideoItems) Item() (*VideoItem, bool) {
   if len(v) >= 1 {
      return &v[0], true
   }
   return nil, false
}

func (v VideoItem) Season() int {
   return int(v.SeasonNum)
}

func (v VideoItem) Episode() int {
   return int(v.EpisodeNum)
}

func (v VideoItem) Title() string {
   return v.Label
}

func (v VideoItem) Year() int {
   return v.AirDateIso.Year()
}

func (v VideoItem) Show() string {
   if v.MediaType == "Full Episode" {
      return v.SeriesTitle
   }
   return ""
}

type VideoItems []VideoItem

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
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
   at.v = url.Values{
      "at": {base64.StdEncoding.EncodeToString(dst)},
   }
   return nil
}

type AppToken struct {
   v url.Values
}

func (SessionToken) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (s SessionToken) RequestUrl() (string, bool) {
   return s.Url, true
}

func (SessionToken) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (s SessionToken) RequestHeader() (http.Header, error) {
   head := http.Header{
      "authorization": {"Bearer " + s.LsSession},
      "content-type": {"application/x-protobuf"},
   }
   return head, nil
}
