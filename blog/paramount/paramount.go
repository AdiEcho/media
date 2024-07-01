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
   *at = AppToken(base64.StdEncoding.EncodeToString(dst))
   return nil
}

func (at *AppToken) Default() error {
   return at.New(app_secrets["15.0.26"])
}

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
   // this needs to be encoded
   req.URL.RawQuery = "at=" + url.QueryEscape(string(at))
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

func (at AppToken) Session(content_id string) (*SessionToken, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json"
   req.URL.RawQuery = url.Values{
      // this needs to be encoded
      "at": {string(at)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   session := new(SessionToken)
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   session.URL += content_id
   return session, nil
}

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

type AppToken string

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
// apkmirror.com/apk/cbs-interactive-inc/paramount
var app_secrets = map[string]string{
   "15.0.26": "2b2caa6373626591",
   "12.0.44": "7297a39a244189d6",
   "12.0.40": "2c160dbae70b337f",
   "12.0.36": "a674920042c954d9",
   "12.0.34": "843970cb0df053ba",
   "12.0.33": "f0613d04b9ba4143",
   "12.0.32": "60e1f010c4e7931e",
   "12.0.28": "439ba2e3622c344a",
   "12.0.27": "79b7e56e442e65ed",
   "12.0.26": "f012987182d6f16c",
    "8.1.28": "d0795c0dffebea73",
    "8.1.26": "a75bd3a39bfcbc77",
    "8.1.23": "c0966212aa651e8b",
    "8.1.22": "ddca2f16bfa3d937",
    "8.1.20": "817774cbafb2b797",
    "8.1.18": "1705732089ff4d60",
    "8.1.16": "add603b54be2fc3c",
    "8.1.14": "acacc94f4214ddc1",
    "8.1.12": "3395c01da67a358b",
    "8.1.10": "8150802ffdeffaf0",
    "8.0.54": "6c70b33080758409",
    "8.0.52": "5fcf8c6937dba442",
    "8.0.50": "2e6cd61ba21dd0c4",
    "8.0.48": "00a7ea18c54d674c",
    "8.0.46": "88065c1d30bc15ed",
    "8.0.44": "9c5b3eda87e38402",
    "8.0.42": "c824c27d68eb9fc3",
    "8.0.40": "d08c12908070b2ac",
    "8.0.38": "423187842fdd7eac",
    "8.0.36": "6dfcc58b09fca975",
    "8.0.34": "0f84a8e9f62594ad",
    "8.0.32": "262d30953b16032b",
    "8.0.30": "90946a66385ceeb5",
    "8.0.28": "1fc4f2e07173b30c",
    "8.0.26": "860c7062bb69759d",
    "8.0.24": "2b7feb264967d94f",
    "8.0.22": "36a841291cfecc4e",
    "8.0.20": "003ff1f049feb54a",
    "8.0.16": "79e71194ad8b97d4",
    "8.0.14": "f3577b860abfa76d",
    "8.0.12": "20021bb2eda91db4",
    "8.0.10": "685c401ff9a4a2d9",
    "8.0.00": "5d1d865f929d3daa",
    "7.3.58": "4be3d46aecbcd26d",
    "4.8.06": "a958002817953588",
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

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

type number int

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

type VideoItem struct {
   AirDateIso time.Time `json:"_airDateISO"`
   EpisodeNum number
   Label string
   SeasonNum number
   SeriesTitle string
   MediaType string
}

func (v VideoItem) Show() string {
   if v.MediaType == "Full Episode" {
      return v.SeriesTitle
   }
   return ""
}
func DashCenc(content_id string) (string, error) {
   query := url.Values{
      "formats": {"MPEG-DASH"},
      "assetTypes": {"DASH_CENC"},
      // "assetTypes": {"DASH_CENC_PRECON"},
   }
   return location(content_id, query)
}

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

type SessionToken struct {
   URL string
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
   return s.URL, true
}

func (SessionToken) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
