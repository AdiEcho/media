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

var apps = []struct{
   country string
   token string
}{
   {
      "US",
      "ABAAAAAAAAAAAAAAAAAAAAAAzj7EcNQMRW+T8yy4tGGC1080Sl81f+pj+oSiktWnDEA=",
   },
   {
      "FR",
      "ABAS/G30Pp6tJuNOlZ1OEE6Rf5goS0KjICkGkBVIapVuxemiiASyWVfW4v7SUeAkogc=",
   },
}

func TestDecode(t *testing.T) {
   for _, app := range apps {
      data, err := decode(app.token)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", data)
   }
}
