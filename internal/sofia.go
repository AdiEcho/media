package internal

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "io"
   "net/http"
   "strconv"
)

func write_init(w io.Writer, r io.Reader) ([]byte, error) {
   var file sofia.File
   if err := file.Read(r); err != nil {
      return nil, err
   }
   for _, box := range file.Movie.Boxes {
      if box.BoxHeader.Type.String() == "pssh" {
         copy(box.BoxHeader.Type[:], "free") // Firefox
      }
   }
   sample, protect := file.
      Movie.
      Track.
      Media.
      MediaInformation.
      SampleTable.
      SampleDescription.
      SampleEntry()
   // Firefox enca encv sinf
   copy(protect.BoxHeader.Type[:], "free")
   // Firefox stsd enca encv
   copy(sample.BoxHeader.Type[:], protect.OriginalFormat.DataFormat[:])
   if err := file.Write(w); err != nil {
      return nil, err
   }
   return protect.SchemeInformation.TrackEncryption.DefaultKid[:], nil
}

func write_sidx(base_url string, bytes dash.Range) ([]sofia.Range, error) {
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
   if err := file.Read(res.Body); err != nil {
      return nil, err
   }
   _, raw_end, _ := bytes.Cut()
   end, err := strconv.ParseUint(raw_end, 10, 64)
   if err != nil {
      return nil, err
   }
   return file.SegmentIndex.Ranges(end + 1), nil
}

func write_segment(w io.Writer, r io.Reader, key []byte) error {
   var file sofia.File
   if err := file.Read(r); err != nil {
      return err
   }
   fragment := file.MovieFragment.TrackFragment
   for i, data := range file.MediaData.Data(fragment.TrackRun) {
      err := fragment.SampleEncryption.Samples[i].DecryptCenc(data, key)
      if err != nil {
         return err
      }
   }
   return file.Write(w)
}
