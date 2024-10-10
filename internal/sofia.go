package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/container"
   "41.neocities.org/sofia/pssh"
   "41.neocities.org/sofia/sidx"
   "errors"
   "io"
   "net/http"
)

func (s *Stream) init_protect(buf []byte) ([]byte, error) {
   var file container.File
   err := file.Read(buf)
   if err != nil {
      return nil, err
   }
   if moov, ok := file.GetMoov(); ok {
      for _, value := range moov.Pssh {
         if value.Widevine() {
            s.pssh = value.Data
         }
         copy(value.BoxHeader.Type[:], "free") // Firefox
      }
      description := moov.Trak.Mdia.Minf.Stbl.Stsd
      if sinf, ok := description.Sinf(); ok {
         s.key_id = sinf.Schi.Tenc.DefaultKid[:]
         // Firefox
         copy(sinf.BoxHeader.Type[:], "free")
         if sample, ok := description.SampleEntry(); ok {
            // Firefox
            copy(sample.BoxHeader.Type[:], sinf.Frma.DataFormat[:])
         }
      }
   }
   return file.Append(nil)
}

func (s *Stream) Download(rep dash.Representation) error {
   if buf, ok := rep.Widevine(); ok {
      var box pssh.Box
      n, err := box.BoxHeader.Decode(buf)
      if err != nil {
         return err
      }
      err = box.Read(buf[n:])
      if err != nil {
         return err
      }
      s.pssh = box.Data
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

func write_segment(buf, key []byte) ([]byte, error) {
   if key == nil {
      return buf, nil
   }
   var file container.File
   err := file.Read(buf)
   if err != nil {
      return nil, err
   }
   track := file.Moof.Traf
   if senc := track.Senc; senc != nil {
      for i, text := range file.Mdat.Data(&track) {
         err = senc.Sample[i].DecryptCenc(text, key)
         if err != nil {
            return nil, err
         }
      }
   }
   return file.Append(nil)
}

func write_sidx(req *http.Request, index dash.Range) ([]sidx.Reference, error) {
   buf, _ := index.MarshalText()
   req.Header.Set("range", "bytes=" + string(buf))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   buf, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var file container.File
   err = file.Read(buf)
   if err != nil {
      return nil, err
   }
   return file.Sidx.Reference, nil
}
