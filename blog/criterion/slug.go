package criterion

import (
   "fmt"
   "net/http"
   "strings"
   "time"
)

var as = []string{
   "api.vhx.com",
   "api.vhx.tv",
}

var bs = []string{
   "",
   "/v2/sites/59054",
}

var cs = []string{
   "/collections/my-dinner-with-andre",
   "/videos/my-dinner-with-andre",
}

var ds = []string{
   "",
   "/items",
}

var es = []string{
   "",
   "?site_id=59054",
}

func (t auth_token) slug() error {
   for _, a := range as {
      for _, b := range bs {
         for _, c := range cs {
            for _, d := range ds {
               for _, e := range es {
                  address := func() string {
                     var f strings.Builder
                     f.WriteString("https://")
                     f.WriteString(a)
                     f.WriteString(b)
                     f.WriteString(c)
                     f.WriteString(d)
                     f.WriteString(e)
                     return f.String()
                  }()
                  req, err := http.NewRequest("", address, nil)
                  if err != nil {
                     return err
                  }
                  req.Header.Set("authorization", "Bearer " + t.v.AccessToken)
                  err = func() error {
                     res, err := http.DefaultClient.Do(req)
                     if err != nil {
                        return err
                     }
                     defer res.Body.Close()
                     fmt.Println(res.Status, address)
                     return nil
                  }()
                  if err != nil {
                     return err
                  }
                  time.Sleep(99 * time.Millisecond)
               }
            }
         }
      }
   }
   return nil
}
