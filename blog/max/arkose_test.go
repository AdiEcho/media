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
   var st st_cookie
   err = st.New()
   if err != nil {
      t.Fatal(err)
   }
   err = st.login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   text, err := st.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.json", text, 0666)
}

func TestConfig(t *testing.T) {
   var st st_cookie
   err := st.New()
   if err != nil {
      t.Fatal(err)
   }
   config, err := st.config()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%s\n", config)
}
