package joyn

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (a anonymous) entitlement(m movie_detail) (*http.Response, error) {
   body, err := json.Marshal(map[string]string{
      "content_id": m.Data.Page.Movie.Video.ID,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://entitlement.p7s1.io/api/user/entitlement-token",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + a.Access_Token)
   return http.DefaultClient.Do(req)
}
