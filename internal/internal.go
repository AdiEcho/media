package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "bytes"
   "encoding/hex"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "os"
   "strings"
)

func (s Stream) segment_template(
   ext, initial string, base *dash.BaseUrl, media []string,
) error {
   file, err := s.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("", initial, nil)
   if err != nil {
      return err
   }
   if initial != "" {
      req.URL = base.Url.ResolveReference(req.URL)
      resp, err := http.DefaultClient.Do(req)
      if err != nil {
         return err
      }
      defer resp.Body.Close()
      if resp.StatusCode != http.StatusOK {
         return errors.New(resp.Status)
      }
      err = s.init_protect(file, resp.Body)
      if err != nil {
         return err
      }
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(media))
   var transport text.Transport
   transport.Set(false)
   defer transport.Set(true)
   for _, medium := range media {
      req.URL, err = base.Url.Parse(medium)
      if err != nil {
         return err
      }
      err := func() error {
         resp, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusOK {
            var b strings.Builder
            resp.Write(&b)
            return errors.New(b.String())
         }
         return write_segment(file, meter.Reader(resp), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func write_segment(to io.Writer, from io.Reader, key []byte) error {
   if key == nil {
      _, err := io.Copy(to, from)
      if err != nil {
         return err
      }
      return nil
   }
   var file sofia.File
   err := file.Read(from)
   if err != nil {
      return err
   }
   track := file.MovieFragment.TrackFragment
   if encrypt := track.SampleEncryption; encrypt != nil {
      for i, data := range file.MediaData.Data(track) {
         err := encrypt.Samples[i].DecryptCenc(data, key)
         if err != nil {
            return err
         }
      }
   }
   return file.Write(to)
}

func (s Stream) segment_base(
   ext string, base *dash.BaseUrl, segment *dash.SegmentBase,
) error {
   data, _ := segment.Initialization.Range.MarshalText()
   var req http.Request
   req.URL = base.Url
   req.Header = make(http.Header)
   // need to use Set for lower case
   req.Header.Set("range", "bytes=" + string(data))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusPartialContent {
      return errors.New(resp.Status)
   }
   file, err := s.Create(ext)
   if err != nil {
      return err
   }
   defer file.Close()
   err = s.init_protect(file, resp.Body)
   if err != nil {
      return err
   }
   key, err := s.key()
   if err != nil {
      return err
   }
   references, err := write_sidx(&req, segment.IndexRange)
   if err != nil {
      return err
   }
   var meter text.ProgressMeter
   meter.Set(len(references))
   var transport text.Transport
   transport.Set(false)
   defer transport.Set(true)
   for _, reference := range references {
      segment.IndexRange.Start = segment.IndexRange.End + 1
      segment.IndexRange.End += uint64(reference.ReferencedSize())
      data, _ := segment.IndexRange.MarshalText()
      err := func() error {
         req.Header.Set("range", "bytes=" + string(data))
         resp, err := http.DefaultClient.Do(&req)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusPartialContent {
            return errors.New(resp.Status)
         }
         return write_segment(file, meter.Reader(resp), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
type ForwardedFor []struct {
   Country string
   IP string
}

func (f ForwardedFor) String() string {
   var b strings.Builder
   for _, each := range f {
      if b.Len() >= 1 {
         b.WriteByte('\n')
      }
      b.WriteString(each.Country)
      b.WriteByte(' ')
      b.WriteString(each.IP)
   }
   return b.String()
}

var Forward = ForwardedFor{
{"Argentina", "186.128.0.0"},
{"Australia", "1.128.0.0"},
{"Bolivia", "179.58.0.0"},
{"Brazil", "179.192.0.0"},
{"Canada", "99.224.0.0"},
{"Chile", "191.112.0.0"},
{"Colombia", "181.128.0.0"},
{"Costa Rica", "201.192.0.0"},
{"Denmark", "2.104.0.0"},
{"Ecuador", "186.68.0.0"},
{"Egypt", "197.32.0.0"},
{"Germany", "53.0.0.0"},
{"Guatemala", "190.56.0.0"},
{"India", "106.192.0.0"},
{"Indonesia", "39.192.0.0"},
{"Ireland", "87.32.0.0"},
{"Italy", "79.0.0.0"},
{"Latvia", "78.84.0.0"},
{"Malaysia", "175.136.0.0"},
{"Mexico", "189.128.0.0"},
{"Netherlands", "145.160.0.0"},
{"New Zealand", "49.224.0.0"},
{"Norway", "88.88.0.0"},
{"Peru", "190.232.0.0"},
{"Russia", "95.24.0.0"},
{"South Africa", "105.0.0.0"},
{"South Korea", "175.192.0.0"},
{"Spain", "88.0.0.0"},
{"Sweden", "78.64.0.0"},
{"Taiwan", "120.96.0.0"},
{"United Kingdom", "25.0.0.0"},
{"Venezuela", "190.72.0.0"},
}

func write_sidx(req *http.Request, index dash.Range) ([]sofia.Reference, error) {
   data, _ := index.MarshalText()
   req.Header.Set("range", "bytes=" + string(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var file sofia.File
   err = file.Read(resp.Body)
   if err != nil {
      return nil, err
   }
   return file.SegmentIndex.Reference, nil
}

func (s Stream) key() ([]byte, error) {
   if s.key_id == nil {
      return nil, nil
   }
   private_key, err := os.ReadFile(s.PrivateKey)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(s.ClientId)
   if err != nil {
      return nil, err
   }
   if s.pssh == nil {
      s.pssh = widevine.Pssh{KeyId: s.key_id}.Encode()
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, s.pssh)
   if err != nil {
      return nil, err
   }
   key, err := module.Key(s.Poster, s.key_id)
   if err != nil {
      return nil, err
   }
   slog.Info(
      "CDM", "ID", hex.EncodeToString(s.key_id), "key", hex.EncodeToString(key),
   )
   return key, nil
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type Stream struct {
   ClientId string
   PrivateKey string
   Name text.Namer
   Poster widevine.Poster
   pssh []byte
   key_id []byte
}

func (s *Stream) init_protect(to io.Writer, from io.Reader) error {
   var file sofia.File
   err := file.Read(from)
   if err != nil {
      return err
   }
   if movie, ok := file.GetMovie(); ok {
      for _, protect := range movie.Protection {
         if protect.Widevine() {
            s.pssh = protect.Data
         }
         copy(protect.BoxHeader.Type[:], "free") // Firefox
      }
      description := movie.
         Track.
         Media.
         MediaInformation.
         SampleTable.
         SampleDescription
      if protect, ok := description.Protection(); ok {
         s.key_id = protect.SchemeInformation.TrackEncryption.DefaultKid[:]
         // Firefox
         copy(protect.BoxHeader.Type[:], "free")
         if sample, ok := description.SampleEntry(); ok {
            // Firefox
            copy(sample.BoxHeader.Type[:], protect.OriginalFormat.DataFormat[:])
         }
      }
   }
   return file.Write(to)
}

func Dash(req *http.Request) ([]dash.Representation, error) {
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   switch resp.Status {
   case "200 OK", "403 OK":
   default:
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return dash.Unmarshal(data, resp.Request.URL)
}
func (s Stream) Create(ext string) (*os.File, error) {
   name, err := text.Name(s.Name)
   if err != nil {
      return nil, err
   }
   return os.Create(text.Clean(name) + ext)
}

func (s *Stream) Download(rep dash.Representation) error {
   if data, ok := rep.Widevine(); ok {
      read := bytes.NewReader(data)
      var pssh sofia.ProtectionSystemSpecificHeader
      err := pssh.BoxHeader.Read(read)
      if err != nil {
         return err
      }
      err = pssh.Read(read)
      if err != nil {
         return err
      }
      s.pssh = pssh.Data
   }
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Representation.Ext")
   }
   base, ok := rep.GetBaseUrl()
   if !ok {
      return errors.New("Representation.GetBaseUrl")
   }
   if rep.SegmentBase != nil {
      return s.segment_base(ext, base, rep.SegmentBase)
   }
   initial, _ := rep.Initialization()
   return s.segment_template(ext, initial, base, rep.Media())
}
