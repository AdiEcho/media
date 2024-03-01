package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "fmt"
   "errors"
   "io"
   "net/http"
   "net/url"
   "os"
   "slices"
)

func (h HttpStream) DASH(reps []dash.Representation, id string) error {
   i := slices.IndexFunc(reps, func(r dash.Representation) bool {
      return r.ID == id
   })
   if i == -1 {
      slices.SortFunc(reps, func(a, b dash.Representation) int {
         return int(b.Bandwidth - a.Bandwidth)
      })
      for i, rep := range reps {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(rep)
      }
      return nil
   }
   rep := reps[i]
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("dash.Representation.Ext")
   }
   if initial, ok := rep.Initialization(); ok {
      return h.segment_template(ext, initial, rep)
   }
   return h.segment_base(ext, rep.BaseURL, rep)
}

func (h HttpStream) key(rep dash.Representation) ([]byte, error) {
   var protect widevine.PSSH
   data, err := rep.PSSH()
   if err != nil {
      key_id, err := rep.Default_KID()
      if err != nil {
         return nil, err
      }
      protect.Key_ID = key_id
   } else {
      err := protect.New(data)
      if err != nil {
         return nil, err
      }
   }
   private_key, err := os.ReadFile(h.Private_Key)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(h.Client_ID)
   if err != nil {
      return nil, err
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      return nil, err
   }
   license, err := module.License(h.Poster)
   if err != nil {
      return nil, err
   }
   key, ok := module.Key(license)
   if !ok {
      return nil, errors.New("widevine.CDM.Key")
   }
   return key, nil
}

func encode_sidx(base_URL string, index dash.Range) ([][2]uint32, error) {
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Range", "bytes=" + string(index))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var f sofia.File
   if err := f.Decode(res.Body); err != nil {
      return nil, err
   }
   _, sidx, err := index.Scan()
   if err != nil {
      return nil, err
   }
   return f.SegmentIndex.ByteRanges(uint32(sidx)+1), nil
}

func encode_init(dst io.Writer, src io.Reader) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for _, b := range f.Movie.Boxes {
      if b.BoxHeader.BoxType() == "pssh" {
         copy(b.BoxHeader.Type[:], "free") // Firefox
      }
   }
   sd := &f.Movie.Track.Media.MediaInformation.SampleTable.SampleDescription
   if as := sd.AudioSample; as != nil {
      copy(as.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         as.Entry.BoxHeader.Type[:],
         as.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   if vs := sd.VisualSample; vs != nil {
      copy(vs.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         vs.Entry.BoxHeader.Type[:],
         vs.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   return f.Encode(dst)
}

func encode_segment(dst io.Writer, src io.Reader, key []byte) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for i, data := range f.MediaData.Data {
      sample := f.MovieFragment.TrackFragment.SampleEncryption.Samples[i]
      err := sample.DecryptCenc(data, key)
      if err != nil {
         return err
      }
   }
   return f.Encode(dst)
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
type HttpStream struct {
   Client_ID string
   Name encoding.Namer
   Poster widevine.Poster
   Private_Key string
   base *url.URL
}

func (h *HttpStream) DashMedia(uri string) ([]dash.Representation, error) {
   res, err := http.Get(uri)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   h.base = res.Request.URL
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return dash.Unmarshal(text)
}