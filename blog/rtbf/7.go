package rtbf

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

const (
   raw_content_id = "bzFDMzdUdDVTem1ITW1FZ1FWaVVFQT09"
   raw_key_id = "o1C37Tt5SzmHMmEgQViUEA=="
)

func seven() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      panic(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      panic(err)
   }
   content_id, err := base64.StdEncoding.DecodeString(raw_content_id)
   if err != nil {
      panic(err)
   }
   key_id, err := base64.StdEncoding.DecodeString(raw_key_id)
   if err != nil {
      panic(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, content_id))
   if err != nil {
      panic(err)
   }
   key, err := module.Key(poster{}, key_id)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%x\n", key)
}

type poster struct{}

func (poster) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set("content-type", "application/x-protobuf")
   return h, nil
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "rbm-rtbf.live.ott.irdeto.com"
   u.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
   u.Scheme = "https"
   val := make(url.Values)
   val["contentId"] = []string{"3201987_6BA97Bb"}
   val["ls_session"] = []string{"eyJ0eXAiOiJKV1QiLCJraWQiOiIwOGIzODQwZS0wYThhLTQyYTItODNhNC03ZGM0Mzc0ZDJmYmEiLCJhbGciOiJIUzI1NiJ9.eyJhaWQiOiJyYm0tcnRiZiIsInN1YiI6ImFjY184OTgyNjk5MDUwZWQ0MWY5OGZlZDA4NzdiZGE2ZjYxNl82QkE5N0JiIiwiaWF0IjoxNzE3NjM0Mzk2LCJleHAiOjE3MTc2Mzc5OTYsImp0aSI6InBNWm85SndPMGtjODhScjhTYzlNb3FlTDlMMXNQd2c1V0tJcVB1UG41WFU9IiwiZW50IjpbeyJlcGlkIjoiZGVmYXVsdCIsImJpZCI6ImZyZWVfcHJvZHVjdF82QkE5N0JiIn1dLCJpc2UiOnRydWUsImVuY3J5cHRpb25Qcm9maWxlSWQiOiJydGJmIn0.iBb5s-CL9uR1admhT05TRUo09Va72QfesR_c9V7SVhU"}
   u.RawQuery = val.Encode()
   return u.String(), true
}
