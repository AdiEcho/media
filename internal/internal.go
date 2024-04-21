package internal

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "errors"
   "io"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strconv"
   "strings"
   "text/template"
)

func (h *HttpStream) DashMedia(url string) ([]*dash.Representation, error) {
   res, err := http.Get(url)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   switch res.Status {
   case "200 OK", "403 OK":
   default:
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var media dash.MPD
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   err = media.Unmarshal(text)
   if err != nil {
      return nil, err
   }
   if media.BaseURL == "" {
      media.BaseURL = res.Request.URL.String()
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

func (h HttpStream) DASH(rep *dash.Representation) error {
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


func (h HttpStream) TimedText(url string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := func() (*os.File, error) {
      s, err := Name(h.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(CleanName(s) + ".vtt")
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   _, err = file.ReadFrom(res.Body)
   if err != nil {
      return err
   }
   return nil
}

var NameFormat = 
   "{{if .Show}}" +
      "{{.Show}} - {{.Season}} {{.Episode}} - {{.Title}}" +
   "{{else}}" +
      "{{.Title}} - {{.Year}}" +
   "{{end}}"

func CleanName(s string) string {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return '-'
      }
      return r
   }
   return strings.Map(mapping, s)
}

func Name(n Namer) (string, error) {
   text, err := new(template.Template).Parse(NameFormat)
   if err != nil {
      return "", err
   }
   var b strings.Builder
   err = text.Execute(&b, n)
   if err != nil {
      return "", err
   }
   return b.String(), nil
}

type Namer interface {
   Show() string
   Season() int
   Episode() int
   Title() string
   Year() int
}

func (h HttpStream) segment_base(
   ext, base_url string, rep *dash.Representation,
) error {
   sb := rep.SegmentBase
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
   file, err := func() (*os.File, error) {
      s, err := Name(h.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(CleanName(s) + ext)
   }()
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
   references, err := write_sidx(base_url, sb.IndexRange)
   if err != nil {
      return err
   }
   var meter log.ProgressMeter
   meter.Set(len(references))
   var start uint64
   end, err := func() (uint64, error) {
      _, s, _ := sb.IndexRange.Cut()
      return strconv.ParseUint(s, 10, 64)
   }()
   if err != nil {
      return err
   }
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, reference := range references {
      start = end + 1
      end += uint64(reference.ReferencedSize())
      bytes := func() string {
         b := []byte("bytes=")
         b = strconv.AppendUint(b, start, 10)
         b = append(b, '-')
         b = strconv.AppendUint(b, end, 10)
         return string(b)
      }()
      err := func() error {
         req, err := http.NewRequest("GET", base_url, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", bytes)
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

func (h HttpStream) segment_template(
   ext, initial string, rep *dash.Representation,
) error {
   base, err := url.Parse(rep.GetAdaptationSet().GetPeriod().GetMpd().BaseURL)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("GET", initial, nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   file, err := func() (*os.File, error) {
      s, err := Name(h.Name)
      if err != nil {
         return nil, err
      }
      return os.Create(CleanName(s) + ext)
   }()
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
   for _, medium := range media {
      req.URL, err = base.Parse(medium)
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
