package internal

import (
   "154.pages.dev/dash"
   "154.pages.dev/sofia"
   "bytes"
   "encoding/base64"
   "errors"
   "io"
   "net/http"
)

func (s *Stream) Download(rep *dash.Representation) error {
   if v, ok := rep.Widevine(); ok {
      data, err := base64.StdEncoding.DecodeString(v)
      if err != nil {
         return err
      }
      r := bytes.NewReader(data)
      var pssh sofia.ProtectionSystemSpecificHeader
      err = pssh.BoxHeader.Read(r)
      if err != nil {
         return err
      }
      err = pssh.Read(r)
      if err != nil {
         return err
      }
      s.pssh = pssh.Data
   }
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Representation.Ext")
   }
   base := rep.GetAdaptationSet().GetPeriod().GetMpd().BaseUrl.URL
   if v, ok := rep.GetSegmentTemplate(); ok {
      if v, ok := v.GetInitialization(rep); ok {
         return s.segment_template(rep, base, v, ext)
      }
   }
   return s.segment_base(rep.SegmentBase, base, *rep.BaseUrl, ext)
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

func write_sidx(req *http.Request, r dash.Range) ([]sofia.Reference, error) {
   data, _ := r.MarshalText()
   req.Header.Set("Range", "bytes=" + string(data))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var file sofia.File
   err = file.Read(res.Body)
   if err != nil {
      return nil, err
   }
   return file.SegmentIndex.Reference, nil
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
   if v := file.MovieFragment.TrackFragment.SampleEncryption; v != nil {
      run := file.MovieFragment.TrackFragment.TrackRun
      for i, data := range file.MediaData.Data(run) {
         err := v.Samples[i].DecryptCenc(data, key)
         if err != nil {
            return err
         }
      }
   }
   return file.Write(to)
}
