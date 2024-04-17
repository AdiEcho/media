package internal

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "io"
   "net/http"
)

func write_init(w io.Writer, r io.Reader) ([]byte, error) {
   var file sofia.File
   err := file.Read(r)
   if err != nil {
      return nil, err
   }
   for _, box := range file.Movie.Boxes {
      if box.BoxHeader.Type.String() == "pssh" {
         copy(box.BoxHeader.Type[:], "free") // Firefox
      }
   }
   description := file.
      Movie.
      Track.
      Media.
      MediaInformation.
      SampleTable.
      SampleDescription
   var key_id []byte
   if protect, ok := description.Protection(); ok {
      key_id = protect.SchemeInformation.TrackEncryption.DefaultKid[:]
      // Firefox
      copy(protect.BoxHeader.Type[:], "free")
      if sample, ok := description.SampleEntry(); ok {
         // Firefox
         copy(sample.BoxHeader.Type[:], protect.OriginalFormat.DataFormat[:])
      }
   }
   err = file.Write(w)
   if err != nil {
      return nil, err
   }
   return key_id, nil
}

func write_segment(w io.Writer, r io.Reader, key []byte) error {
   var file sofia.File
   err := file.Read(r)
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
   return file.Write(w)
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
