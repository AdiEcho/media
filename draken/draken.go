package draken

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

// NO ANONYMOUS QUERY
const get_custom_id = `
query GetCustomIdFullMovie($customId: ID!) {
   viewer {
      viewableCustomId(customId: $customId) {
         ... on Movie {
            defaultPlayable {
               id
            }
            productionYear
            title
         }
      }
   }
}
`

func (a *AuthLogin) New(identity, key string) error {
   body, err := json.Marshal(map[string]string{
      "accessKey": key,
      "identity":  identity,
   })
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://drakenfilm.se/api/auth/login", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AuthLogin) Entitlement(movie *FullMovie) (*Entitlement, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/entitlement/v2/asset/" + movie.DefaultPlayable.Id
   req.Header.Set("authorization", "Bearer "+a.Token)
   magine_accesstoken.set(req.Header)
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
   title := &Entitlement{}
   err = json.NewDecoder(resp.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

func (a *AuthLogin) Playback(
   movie *FullMovie, title *Entitlement,
) (*Playback, error) {
   req, err := http.NewRequest("POST", "https://client-api.magine.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/playback/v1/preflight/asset/" + movie.DefaultPlayable.Id
   magine_accesstoken.set(req.Header)
   magine_play_devicemodel.set(req.Header)
   magine_play_deviceplatform.set(req.Header)
   magine_play_devicetype.set(req.Header)
   magine_play_drm.set(req.Header)
   magine_play_protocol.set(req.Header)
   req.Header.Set("authorization", "Bearer "+a.Token)
   req.Header.Set("magine-play-deviceid", "!")
   req.Header.Set("magine-play-entitlementid", title.Token)
   x_forwarded_for.set(req.Header)
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
   play := &Playback{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type AuthLogin struct {
   Token string
   Raw []byte `json:"-"`
}

func (a *AuthLogin) Unmarshal() error {
   return json.Unmarshal(a.Raw, a)
}

func (f *FullMovie) New(custom_id string) error {
   var req_body struct {
      Query     string `json:"query"`
      Variables struct {
         CustomId string `json:"customId"`
      } `json:"variables"`
   }
   req_body.Variables.CustomId = custom_id
   req_body.Query = graphql_compact(get_custom_id)
   data, err := json.Marshal(req_body)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/apiql/v2",
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   magine_accesstoken.set(req.Header)
   x_forwarded_for.set(req.Header)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var resp_body struct {
      Data struct {
         Viewer struct {
            ViewableCustomId *FullMovie
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return err
   }
   if id := resp_body.Data.Viewer.ViewableCustomId; id != nil {
      *f = *id
      return nil
   }
   return errors.New("ViewableCustomId")
}

func (n Namer) Title() string {
   return n.Movie.Title
}

func (n Namer) Year() int {
   return n.Movie.ProductionYear
}

type Namer struct {
   Movie *FullMovie
}

func (Namer) Episode() int {
   return 0
}

func (Namer) Season() int {
   return 0
}

func (Namer) Show() string {
   return ""
}

type FullMovie struct {
   DefaultPlayable struct {
      Id string
   }
   ProductionYear int `json:",string"`
   Title          string
}

func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

type Entitlement struct {
   Token string
}

type header struct {
   key   string
   value string
}

var magine_accesstoken = header{
   "magine-accesstoken", "22cc71a2-8b77-4819-95b0-8c90f4cf5663",
}

var magine_play_devicemodel = header{
   "magine-play-devicemodel", "firefox 111.0 / windows 10",
}

var magine_play_deviceplatform = header{
   "magine-play-deviceplatform", "firefox",
}

var magine_play_devicetype = header{
   "magine-play-devicetype", "web",
}

var magine_play_drm = header{
   "magine-play-drm", "widevine",
}

var magine_play_protocol = header{
   "magine-play-protocol", "dashs",
}

// this value is important, with the wrong value you get random failures
var x_forwarded_for = header{
   "x-forwarded-for", "95.192.0.0",
}

type Playback struct {
   Headers  map[string]string
   Playlist string
}

type Poster struct {
   Login *AuthLogin
   Play *Playback
}

func (h *header) set(head http.Header) {
   head.Set(h.key, h.value)
}

func (p *Poster) RequestHeader() (http.Header, error) {
   head := http.Header{}
   magine_accesstoken.set(head)
   head.Set("authorization", "Bearer "+p.Login.Token)
   for key, value := range p.Play.Headers {
      head.Set(key, value)
   }
   return head, nil
}

func (*Poster) RequestUrl() (string, bool) {
   return "https://client-api.magine.com/api/playback/v1/widevine/license", true
}

func (*Poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (*Poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}
