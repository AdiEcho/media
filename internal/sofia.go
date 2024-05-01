package internal

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "io"
   "net/http"
)

func write_segment(dst io.Writer, src io.Reader, key []byte) error {
   if key == nil {
      _, err := io.Copy(dst, src)
      if err != nil {
         return err
      }
      return nil
   }
   var file sofia.File
   err := file.Read(src)
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
   return file.Write(dst)
}

func write_sidx(base_url string, bytes dash.Range) ([]sofia.Reference, error) {
   req, err := http.NewRequest("GET", base_url, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Range", "bytes=" + string(bytes))
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

func read_init(r io.Reader) (widevine.Data, error) {
   var file sofia.File
   err := file.Read(r)
   if err != nil {
      return nil, err
   }
   data := func() widevine.Data {
      movie, ok := file.GetMovie()
      if !ok {
         return nil
      }
      for _, protect := range movie.Protection {
         if protect.Widevine() {
            return widevine.PSSH(protect.Data)
         }
      }
      sample := movie.Track.Media.MediaInformation.SampleTable
      if protect, ok := sample.SampleDecription.Protection(); ok {
         key_id := protect.SchemeInformation.TrackEncryption.DefaultKid[:]
         return widevine.KeyId(key_id)
      }
   }()
   return data, nil
}

func write_init(w io.Writer) error {
   var file sofia.File
   for _, protect := range file.Movie.Protection {
      copy(protect.BoxHeader.Type[:], "free") // Firefox
   }
   description := file.SampleDecription()
   if protect, ok := description.Protection(); ok {
      // Firefox
      copy(protect.BoxHeader.Type[:], "free")
      if sample, ok := description.SampleEntry(); ok {
         // Firefox
         copy(sample.BoxHeader.Type[:], protect.OriginalFormat.DataFormat[:])
      }
   }
   return file.Write(w)
}
