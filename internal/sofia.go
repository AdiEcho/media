package internal

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "io"
   "net/http"
)

func encode_sidx(base_url string, raw dash.RawRange) ([]sofia.Range, error) {
   req, err := http.NewRequest("GET", base_url, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Range", "bytes=" + string(raw))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var file sofia.File
   if err := file.Read(res.Body); err != nil {
      return nil, err
   }
   index, err := raw.Scan()
   if err != nil {
      return nil, err
   }
   return file.SegmentIndex.Ranges(index.End+1), nil
}

func encode_init(dst io.Writer, src io.Reader) error {
   var file sofia.File
   if err := file.Read(src); err != nil {
      return err
   }
   for _, b := range file.Movie.Boxes {
      if b.BoxHeader.GetType() == "pssh" {
         copy(b.BoxHeader.Type[:], "free") // Firefox
      }
   }
   sd := &file.Movie.Track.Media.MediaInformation.SampleTable.SampleDescription
   if as := sd.AudioSample; as != nil {
      copy(as.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         as.SampleEntry.BoxHeader.Type[:],
         as.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   if vs := sd.VisualSample; vs != nil {
      copy(vs.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         vs.SampleEntry.BoxHeader.Type[:],
         vs.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   return file.Write(dst)
}

func encode_segment(dst io.Writer, src io.Reader, key []byte) error {
   var file sofia.File
   if err := file.Read(src); err != nil {
      return err
   }
   for i, data := range file.MediaData.Data {
      sample := file.MovieFragment.TrackFragment.SampleEncryption.Samples[i]
      err := sample.DecryptCenc(data, key)
      if err != nil {
         return err
      }
   }
   return file.Write(dst)
}
