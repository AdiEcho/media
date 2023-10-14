package youtube

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

type config struct {
   Innertube_Client_Name string
   Innertube_Client_Version string
}

func new_config() (*config, error) {
   res, err := http.Get("https://www.youtube.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = json.Cut(text, []byte("\nytcfg.set("), nil)
   con := new(config)
   if err := json.Unmarshal(text, con); err != nil {
      return nil, err
   }
   return con, nil
}
