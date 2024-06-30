package paramount

import (
   "fmt"
   "testing"
)

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
