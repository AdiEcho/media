package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/sofia"
   "bytes"
   "errors"
   "io"
   "net/http"
)

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
