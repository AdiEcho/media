package youtube

import (
   "errors"
   "mime"
   "net/url"
   "path"
   "strconv"
   "strings"
)

const (
   android_version = "18.43.39"
   web_version = "2.20231219.04.00"
)

func (r *Request) Set(s string) error {
   ref, err := url.Parse(s)
   if err != nil {
      return err
   }
   r.Video_ID = ref.Query().Get("v")
   if r.Video_ID == "" {
      r.Video_ID = path.Base(ref.Path)
   }
   return nil
}

func (r Request) String() string {
   return r.Video_ID
}

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

type Request struct {
   Content_Check_OK bool `json:"contentCheckOk,omitempty"`
   Context struct {
      Client struct {
         Android_SDK_Version int `json:"androidSdkVersion"`
         Client_Name string `json:"clientName"`
         Client_Version string `json:"clientVersion"`
         // need this to get the correct:
         // This video requires payment to watch
         // instead of the invalid:
         // This video can only be played on newer versions of Android or other
         // supported devices.
         OS_Version string `json:"osVersion"`
      } `json:"client"`
   } `json:"context"`
   Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
   Video_ID string `json:"videoId"`
}

func (r *Request) Web(video_id string) {
   r.Context.Client.Client_Name = "WEB"
   r.Context.Client.Client_Version = web_version
   r.Video_ID = video_id
}

func (r *Request) Android_Embed(video_id string) {
   r.Context.Client.Client_Name = "ANDROID_EMBEDDED_PLAYER"
   r.Context.Client.Client_Version = android_version
   r.Video_ID = video_id
}

func (r *Request) Android(video_id string) {
   r.Content_Check_OK = true
   r.Context.Client.Client_Name = "ANDROID"
   r.Context.Client.Client_Version = android_version
   r.Video_ID = video_id
}

func (r *Request) Android_Check(video_id string) {
   r.Content_Check_OK = true
   r.Context.Client.Client_Name = "ANDROID"
   r.Context.Client.Client_Version = android_version
   r.Racy_Check_OK = true
   r.Video_ID = video_id
}
