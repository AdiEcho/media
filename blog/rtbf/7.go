package rtbf

import (
   "net/http"
   "net/url"
)

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
