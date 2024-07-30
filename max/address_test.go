package max

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

func TestRoutes(t *testing.T) {
   var token DefaultToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web WebAddress
      web.UnmarshalText([]byte(test.url))
      routes, err := token.Routes(web)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", routes)
      name, err := text.Name(routes)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
