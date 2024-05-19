package criterion

import (
   "os"
   "testing"
)

func TestSlug(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   err = token.slug()
   if err != nil {
      t.Fatal(err)
   }
}
