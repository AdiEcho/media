package max

import (
   "fmt"
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   var request login_request
   request.Credentials.Username = os.Getenv("max_username")
   if request.Credentials.Username == "" {
      t.Fatal("Getenv")
   }
   request.Credentials.Password = os.Getenv("max_password")
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
   login, err := request.login(key, token)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
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
