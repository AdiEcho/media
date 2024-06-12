package max

import (
   "fmt"
   "os"
   "testing"
)

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
   res, err := key.login(login, token)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

func TestConfig(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   config, err := token.config()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%s\n", config)
}
