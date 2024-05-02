package paramount

import (
   "crypto/aes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type app_details struct {
   version string
   code int
}

// com.cbs.app
var app_secrets = map[app_details]string{
   { "4.8.06",   1648603}: "a958002817953588",
   { "7.3.58", 210735833}: "4be3d46aecbcd26d",
   { "8.0.00", 210800061}: "5d1d865f929d3daa",
   { "8.0.10", 210801034}: "685c401ff9a4a2d9",
   { "8.0.12", 210801228}: "20021bb2eda91db4",
   { "8.0.14", 210801417}: "f3577b860abfa76d",
   { "8.0.16", 210801619}: "79e71194ad8b97d4",
   { "8.0.20", 210802025}: "003ff1f049feb54a",
   { "8.0.22", 210802235}: "36a841291cfecc4e",
   { "8.0.24", 210802415}: "2b7feb264967d94f",
   { "8.0.26", 210802628}: "860c7062bb69759d",
   { "8.0.28", 210802834}: "1fc4f2e07173b30c",
   { "8.0.30", 210803027}: "90946a66385ceeb5",
   { "8.0.32", 210803227}: "262d30953b16032b",
   { "8.0.34", 210803426}: "0f84a8e9f62594ad",
   { "8.0.36", 210803633}: "6dfcc58b09fca975",
   { "8.0.38", 210803826}: "423187842fdd7eac",
   { "8.0.40", 210804027}: "d08c12908070b2ac",
   { "8.0.42", 210804235}: "c824c27d68eb9fc3",
   { "8.0.44", 210804421}: "9c5b3eda87e38402",
   { "8.0.46", 210804638}: "88065c1d30bc15ed",
   { "8.0.48", 210804826}: "00a7ea18c54d674c",
   { "8.0.50", 210805068}: "2e6cd61ba21dd0c4",
   { "8.0.52", 210805212}: "5fcf8c6937dba442",
   { "8.0.54", 210805434}: "6c70b33080758409",
   { "8.1.10", 210811006}: "8150802ffdeffaf0",
   { "8.1.12", 210811230}: "3395c01da67a358b",
   { "8.1.14", 210811423}: "acacc94f4214ddc1",
   { "8.1.16", 210811625}: "add603b54be2fc3c",
   { "8.1.18", 210811824}: "1705732089ff4d60",
   { "8.1.20", 210812024}: "817774cbafb2b797",
   { "8.1.22", 210812237}: "ddca2f16bfa3d937",
   { "8.1.23", 210812302}: "c0966212aa651e8b",
   { "8.1.26", 210812630}: "a75bd3a39bfcbc77",
   { "8.1.28", 210812860}: "d0795c0dffebea73",
   {"12.0.26", 211202668}: "f012987182d6f16c",
   {"12.0.27", 211202700}: "79b7e56e442e65ed",
   {"12.0.28", 211202876}: "439ba2e3622c344a",
   {"12.0.32", 211203249}: "60e1f010c4e7931e",
   {"12.0.33", 211203304}: "f0613d04b9ba4143",
   {"12.0.34", 211203457}: "843970cb0df053ba",
   {"12.0.36", 211203732}: "a674920042c954d9",
   {"12.0.40", 211204029}: "2c160dbae70b337f",
   {"12.0.44", 211204450}: "7297a39a244189d6",
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

func DashCenc(content_id string) (string, error) {
   query := url.Values{
      "assetTypes": {"DASH_CENC"},
      "formats": {"MPEG-DASH"},
   }
   return location(content_id, query)
}

func use_last_response(*http.Request, []*http.Request) error {
   return http.ErrUseLastResponse
}

func location(content_id string, query url.Values) (string, error) {
   req, err := http.NewRequest("GET", "http://link.theplatform.com", nil)
   if err != nil {
      return "", err
   }
   req.URL.RawQuery = query.Encode()
   req.URL.Path = func() string {
      b := []byte("/s/")
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_id...)
      return string(b)
   }()
   http.DefaultClient.CheckRedirect = use_last_response
   defer func() { http.DefaultClient.CheckRedirect = nil }()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusFound {
      var s struct {
         Description string
      }
      json.NewDecoder(res.Body).Decode(&s)
      return "", errors.New(s.Description)
   }
   return res.Header.Get("Location"), nil
}

type SessionToken struct {
   URL string
   LS_Session string
}

func (SessionToken) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (s SessionToken) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("authorization", "Bearer " + s.LS_Session)
   return head, nil
}

func (s SessionToken) RequestUrl() (string, bool) {
   return s.URL, true
}

func (SessionToken) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (i Item) Show() string {
   return i.SeriesTitle
}

func (i Item) Title() string {
   return i.Label
}

func (i Item) Year() int {
   if v, _, ok := strings.Cut(i.AirDateIso, "-"); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}

type Item struct {
   AirDateIso string `json:"_airDateISO"`
   Label string
   MediaType string
   SeriesTitle string
   // these can be empty string, so we cannot use these:
   // int `json:",string"`
   // json.Number
   SeasonNum string
   EpisodeNum string
}

func (i Item) Season() int {
   if v, err := strconv.Atoi(i.SeasonNum); err == nil {
      return v
   }
   return 0
}

func (i Item) Episode() int {
   if v, err := strconv.Atoi(i.EpisodeNum); err == nil {
      return v
   }
   return 0
}
