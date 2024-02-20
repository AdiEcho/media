package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "encoding/hex"
   "errors"
   "fmt"
   "log/slog"
   "net/http"
   "os"
)

func (h HttpStream) segment_base(
   ext, base_URL string, point dash.Pointer,
) error {
   key, err := h.key(point)
   if err != nil {
      return err
   }
   slog.Debug("hex", "key", hex.EncodeToString(key))
   sb := point.Representation.SegmentBase
   // Initialization
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return err
   }
   req.Header.Set("Range", "bytes=" + string(sb.Initialization.Range))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(encoding.Name(h.Name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   byte_ranges, err := encode_sidx(base_URL, sb.IndexRange)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   meter.Set(len(byte_ranges))
   log.TransportDebug()
   defer log.TransportInfo()
   for _, r := range byte_ranges {
      err := func() error {
         req, err := http.NewRequest("GET", base_URL, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r[0], r[1]))
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         return encode_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (h HttpStream) segment_template(
   ext, initialization string, point dash.Pointer,
) error {
   key, err := h.key(point)
   if err != nil {
      return err
   }
   slog.Debug("hex", "key", hex.EncodeToString(key))
   req, err := http.NewRequest("GET", initialization, nil)
   if err != nil {
      return err
   }
   req.URL = h.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(encoding.Name(h.Name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   media := point.Media()
   var meter log.ProgressMeter
   meter.Set(len(media))
   log.TransportDebug()
   defer log.TransportInfo()
   for _, ref := range media {
      // with DASH, initialization and media URLs are relative to the MPD URL
      req.URL, err = h.Base.Parse(ref)
      if err != nil {
         return err
      }
      err := func() error {
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            return errors.New(res.Status)
         }
         return encode_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
