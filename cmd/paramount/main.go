package main

import (
   "154.pages.dev/http"
   "154.pages.dev/media/paramount"
   "154.pages.dev/stream"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int
   codec string
   content_ID string
   dash_cenc bool
   height int
   lang string
   s stream.Stream
}

func main() {
   home, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      return filepath.ToSlash(s) + "/widevine/", nil
   }()
   if err != nil {
      panic(err)
   }
   var f flags
   flag.StringVar(&f.content_ID, "b", "", "content ID")
   flag.IntVar(&f.bandwidth, "bandwidth", 5_000_000, "maximum bandwidth")
   flag.StringVar(&f.codec, "c", "mp4a", "audio codec")
   flag.StringVar(&f.s.Client_ID, "client", home+"client_id.bin", "client ID")
   flag.BoolVar(&f.dash_cenc, "d", false, "DASH_CENC")
   flag.IntVar(&f.height, "h", 720, "maximum height")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(&f.s.Private_Key, "key", home+"private_key.pem", "private key")
   flag.StringVar(&f.lang, "language", "en", "audio language")
   flag.Parse()
   if f.content_ID != "" {
      http.No_Location()
      http.Verbose()
      token, err := paramount.New_App_Token()
      if err != nil {
         panic(err)
      }
      if f.dash_cenc {
         err = f.dash(token)
      } else {
         err = f.downloadable(token)
      }
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
