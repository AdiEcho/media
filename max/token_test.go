package max

import (
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   var login DefaultLogin
   login.Credentials.Username = os.Getenv("max_username")
   if login.Credentials.Username == "" {
      t.Fatal("Getenv")
   }
   login.Credentials.Password = os.Getenv("max_password")
   var key PublicKey
   err := key.New()
   if err != nil {
      t.Fatal(err)
   }
   var token DefaultToken
   err = token.New()
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   err = token.Login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("session.txt", token.Session.Raw, os.ModePerm)
   os.WriteFile("token.txt", token.Token.Raw, os.ModePerm)
}
