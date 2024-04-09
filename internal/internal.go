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
   "strings"
)

func (h HttpStream) DASH(rep dash.Representation) error {
   ext, ok := rep.Ext()
   if !ok {
      return errors.New("Ext")
   }
   if v, ok := rep.GetSegmentTemplate(); ok {
      if v, ok := v.GetInitialization(rep); ok {
         return h.segment_template(ext, v, rep)
      }
   }
   return h.segment_base(ext, *rep.BaseURL, rep)
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
   name, err := encoding.Name(h.Name)
   if err != nil {
      return err
   }
   file, err := os.Create(encoding.Clean(name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   key_id, err := write_init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := h.key(key_id)
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

func (h HttpStream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name, err := encoding.Name(h.Name)
   if err != nil {
      return err
   }
   file, err := os.Create(encoding.Clean(name) + ".vtt")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
func (h *HttpStream) DashMedia(url string) ([]*dash.Representation, error) {
   res, err := http.Get(url)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var media dash.MPD
   if err := media.Unmarshal(text); err != nil {
      return nil, err
   }
   if media.BaseURL == "" {
      media.BaseURL = url
   }
   var reps []*dash.Representation
   for _, v := range media.Period {
      seconds, err := v.Seconds()
      if err != nil {
         return nil, err
      }
      for _, v := range v.AdaptationSet {
         for _, v := range v.Representation {
            if seconds > 9 {
               if _, ok := v.Ext(); ok {
                  reps = append(reps, v)
               }
            }
         }
      }
   }
   slices.SortFunc(reps, func(a, b *dash.Representation) int {
      return int(a.Bandwidth - b.Bandwidth)
   })
   return reps, nil
}

func (h HttpStream) segment_template(
   ext, initial string, rep *dash.Representation,
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
   name, err := encoding.Name(h.Name)
   if err != nil {
      return err
   }
   file, err := os.Create(encoding.Clean(name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   key_id, err := write_init(file, res.Body)
   if err != nil {
      return err
   }
   key, err := h.key(key_id)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   template, ok := rep.GetSegmentTemplate()
   if !ok {
      return errors.New("GetSegmentTemplate")
   }
   media, err := template.GetMedia(rep)
   if err != nil {
      return err
   }
   meter.Set(len(media))
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
            var b strings.Builder
            res.Write(&b)
            return errors.New(b.String())
         }
         return write_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

