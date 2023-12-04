package youtube

import (
   "encoding/json"
   "errors"
   "mime"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
)

func (f Format) Ranges() []string {
   const bytes = 10_000_000
   var byte_ranges []string
   var pos int64
   for pos < f.Content_Length {
      byte_range := func() string {
         b := []byte("&range=")
         b = strconv.AppendInt(b, pos, 10)
         b = append(b, '-')
         b = strconv.AppendInt(b, pos+bytes-1, 10)
         return string(b)
      }()
      byte_ranges = append(byte_ranges, byte_range)
      pos += bytes
   }
   return byte_ranges
}

// YouTube on TV
const (
   client_ID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func New_Device_Code() (*Device_Code, error) {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/device/code",
      url.Values{
         "client_id": {client_ID},
         "scope": {"https://www.googleapis.com/auth/youtube"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   code := new(Device_Code)
   if err := json.NewDecoder(res.Body).Decode(code); err != nil {
      return nil, err
   }
   return code, nil
}

func (t *Token) Refresh() error {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_ID},
         "client_secret": {client_secret},
         "grant_type": {"refresh_token"},
         "refresh_token": {t.Refresh_Token},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}

func (d Device_Code) Token() (*Token, error) {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_ID},
         "client_secret": {client_secret},
         "device_code": {d.Device_Code},
         "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tok := new(Token)
   if err := json.NewDecoder(res.Body).Decode(tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func (d Device_Code) String() string {
   var b strings.Builder
   b.WriteString("1. Go to\n")
   b.WriteString(d.Verification_URL)
   b.WriteString("\n\n2. Enter this code\n")
   b.WriteString(d.User_Code)
   b.WriteString("\n\n3. Press Enter to continue")
   return b.String()
}

type Device_Code struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func Read_Token(name string) (*Token, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   tok := new(Token)
   if err := json.Unmarshal(text, tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func (t Token) Write_File(name string) error {
   text, err := json.Marshal(t)
   if err != nil {
      return err
   }
   return os.WriteFile(name, text, 0666)
}

type Token struct {
   Access_Token string
   Error string
   Refresh_Token string
}

const user_agent = "com.google.android.youtube/"

type Format struct {
   Quality_Label string `json:"qualityLabel"`
   Audio_Quality string `json:"audioQuality"`
   Bitrate int64
   Content_Length int64 `json:"contentLength,string"`
   MIME_Type string `json:"mimeType"`
   URL string
}

func (f Format) String() string {
   var b []byte
   b = append(b, "quality: "...)
   if f.Quality_Label != "" {
      b = append(b, f.Quality_Label...)
   } else {
      b = append(b, f.Audio_Quality...)
   }
   b = append(b, "\nbitrate: "...)
   b = strconv.AppendInt(b, f.Bitrate, 10)
   b = append(b, "\ntype: "...)
   b = append(b, f.MIME_Type...)
   return string(b)
}

func (f Format) Ext() (string, error) {
   media, _, err := mime.ParseMediaType(f.MIME_Type)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", errors.New(f.MIME_Type)
}

func (i Image) URL(id string) *url.URL {
   return &url.URL{
      Scheme: "http",
      Host: "i.ytimg.com",
      Path: func() string {
         var b strings.Builder
         b.WriteString("/vi")
         if strings.HasSuffix(i.Name, ".webp") {
            b.WriteString("_webp")
         }
         b.WriteByte('/')
         b.WriteString(id)
         b.WriteByte('/')
         b.WriteString(i.Name)
         return b.String()
      }(),
   }
}

type Image struct {
   Crop bool
   Height int
   Name string
   Width int
}

var Images = []Image{
   {Width:120, Height:90, Name:"default.jpg"},
   {Width:120, Height:90, Name:"1.jpg"},
   {Width:120, Height:90, Name:"2.jpg"},
   {Width:120, Height:90, Name:"3.jpg"},
   {Width:120, Height:90, Name:"default.webp"},
   {Width:120, Height:90, Name:"1.webp"},
   {Width:120, Height:90, Name:"2.webp"},
   {Width:120, Height:90, Name:"3.webp"},
   {Width:320, Height:180, Name:"mq1.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq2.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq3.jpg", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.jpg"},
   {Width:320, Height:180, Name:"mq1.webp", Crop:true},
   {Width:320, Height:180, Name:"mq2.webp", Crop:true},
   {Width:320, Height:180, Name:"mq3.webp", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.webp"},
   {Width:480, Height:360, Name:"0.jpg"},
   {Width:480, Height:360, Name:"hqdefault.jpg"},
   {Width:480, Height:360, Name:"hq1.jpg"},
   {Width:480, Height:360, Name:"hq2.jpg"},
   {Width:480, Height:360, Name:"hq3.jpg"},
   {Width:480, Height:360, Name:"0.webp"},
   {Width:480, Height:360, Name:"hqdefault.webp"},
   {Width:480, Height:360, Name:"hq1.webp"},
   {Width:480, Height:360, Name:"hq2.webp"},
   {Width:480, Height:360, Name:"hq3.webp"},
   {Width:640, Height:480, Name:"sddefault.jpg"},
   {Width:640, Height:480, Name:"sd1.jpg"},
   {Width:640, Height:480, Name:"sd2.jpg"},
   {Width:640, Height:480, Name:"sd3.jpg"},
   {Width:640, Height:480, Name:"sddefault.webp"},
   {Width:640, Height:480, Name:"sd1.webp"},
   {Width:640, Height:480, Name:"sd2.webp"},
   {Width:640, Height:480, Name:"sd3.webp"},
   {Width:1280, Height:720, Name:"hq720.jpg"},
   {Width:1280, Height:720, Name:"maxresdefault.jpg"},
   {Width:1280, Height:720, Name:"maxres1.jpg"},
   {Width:1280, Height:720, Name:"maxres2.jpg"},
   {Width:1280, Height:720, Name:"maxres3.jpg"},
   {Width:1280, Height:720, Name:"hq720.webp"},
   {Width:1280, Height:720, Name:"maxresdefault.webp"},
   {Width:1280, Height:720, Name:"maxres1.webp"},
   {Width:1280, Height:720, Name:"maxres2.webp"},
   {Width:1280, Height:720, Name:"maxres3.webp"},
}
