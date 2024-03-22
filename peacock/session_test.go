package peacock

import (
   "fmt"
   "os"
   "testing"
)

func TestSession(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/peacock.json")
   if err != nil {
      t.Fatal(err)
   }
   var session IdSession
   session.Unmarshal(text)
   auth, err := session.Auth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
