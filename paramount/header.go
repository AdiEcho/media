package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

type Header struct {
   Header http.Header
}

func (h Header) Location() string {
   return h.Header.Get("location")
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

// must use IP address for correct location
func (h *Header) New(content_id string) error {
   req, err := http.NewRequest("", "https://link.theplatform.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = func() string {
      b := []byte("/s/")
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_id...)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "assetTypes": {"DASH_CENC_PRECON"},
      "formats": {"MPEG-DASH"},
   }.Encode()
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusFound {
      var s struct {
         Description string
      }
      json.NewDecoder(resp.Body).Decode(&s)
      return errors.New(s.Description)
   }
   h.Header = resp.Header
   return nil
}

func (h Header) JsonMarshal() ([]byte, error) {
   return json.MarshalIndent(h, "", " ")
}

func (h *Header) Json(text []byte) error {
   return json.Unmarshal(text, h)
}
