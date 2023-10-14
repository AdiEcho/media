package youtube

import (
   "154.pages.dev/http"
   "encoding/json"
   "os"
   "testing"
)

func Test_Config(t *testing.T) {
   http.No_Location()
   http.Verbose()
   con, err := new_config()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(con)
}
