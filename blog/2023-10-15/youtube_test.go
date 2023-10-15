package youtube

import (
   "154.pages.dev/http"
   "encoding/json"
   "os"
   "testing"
   "time"
)

var ids = []string{
   "2ZcDwdXEVyI", // episode
   "HPkDFc8hq5c", // film
}

func Test_Watch(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   http.No_Location()
   http.Verbose()
   for _, id := range ids {
      c, err := contents(id)
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(c)
      time.Sleep(time.Second)
   }
}
