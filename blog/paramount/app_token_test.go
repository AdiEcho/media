package paramount

import (
   "154.pages.dev/media/paramount"
   "fmt"
   "net/url"
   "testing"
)

var secrets = []struct{
   key string
   value string
}{
   {"google_mobile", "8c4edb1155a410e4"},
   {"intl", "710d40b8f1ec1746"},
   {"com.cbs.ca_15.0.28", "c0b1d5d6ed27a3f6"},
}

func TestEncode(t *testing.T) {
   for _, secret := range secrets {
      var token paramount.AppToken
      err := token.New(secret.value)
      if err != nil {
         t.Fatal(err)
      }
      at := url.Values{
         "at": {string(token)},
      }.Encode()
      fmt.Print(secret.key, "\n", at, "\n\n")
   }
}
