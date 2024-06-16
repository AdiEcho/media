package max

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func TestConfig(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   decision, err := token.decision()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(decision)
}

func TestToken(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestLogin(t *testing.T) {
   var login default_login
   login.Credentials.Username = os.Getenv("max_username")
   if login.Credentials.Username == "" {
      t.Fatal("Getenv")
   }
   login.Credentials.Password = os.Getenv("max_password")
   var key public_key
   err := key.New()
   if err != nil {
      t.Fatal(err)
   }
   var token default_token
   err = token.New()
   if err != nil {
      t.Fatal(err)
   }
   err = token.login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   text, err := token.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", text, 0666)
}
