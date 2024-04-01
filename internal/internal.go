package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "errors"
   "io"
   "net/http"
   "os"
   "slices"
)

func (h HttpStream) segment_template(
   ext, initial string, rep dash.Representation,
) error {
   req, err := http.NewRequest("GET", initial, nil)
   if err != nil {
      return err
   }
   req.URL = h.base.ResolveReference(req.URL)
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
   pssh, err := write_init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := h.key(pssh)
   if err != nil {
      return err
   }
   media := rep.Media()
   var meter log.ProgressMeter
   meter.Set(len(media))
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, ref := range media {
      // with DASH, initialization and media URLs are relative to the MPD URL
      req.URL, err = h.base.Parse(ref)
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
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (h HttpStream) DASH(rep dash.Representation) error {
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("dash.Representation.Ext")
   }
   if initial, ok := rep.Initialization(); ok {
      return h.segment_template(ext, initial, rep)
   }
   return h.segment_base(ext, rep.BaseURL, rep)
}

func (h *HttpStream) DashMedia(url string) ([]dash.Representation, error) {
   res, err := http.Get(url)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   h.base = res.Request.URL
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   reps, err := dash.Unmarshal(text)
   if err != nil {
      return nil, err
   }
   reps = slices.DeleteFunc(reps, func(r dash.Representation) bool {
      if _, ok := r.Ext(); !ok {
         return true
      }
      if r.Protection() == nil {
         return true
      }
      return false
   })
   slices.SortFunc(reps, func(a, b dash.Representation) int {
      return int(a.Bandwidth - b.Bandwidth)
   })
   return reps, nil
}

func (h HttpStream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(encoding.Name(h.Name) + ".vtt")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func (h HttpStream) segment_base(
   ext, base_url string, rep dash.Representation,
) error {
   sb := rep.SegmentBase
   // Initialization
   req, err := http.NewRequest("GET", base_url, nil)
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
   pssh, err := write_init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := h.key(pssh)
   if err != nil {
      return err
   }
   byte_ranges, err := write_sidx(base_url, sb.IndexRange)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   meter.Set(len(byte_ranges))
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, r := range byte_ranges {
      err := func() error {
         req, err := http.NewRequest("GET", base_url, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", r.String())
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusPartialContent {
            return errors.New(res.Status)
         }
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
