package main

import (
   "fmt"
   "io"
   "net/http"
   "strings"
)

func main() {
   s := `
   {
      "context": {
         "client": {
            "clientName": "ANDROID",
            "clientVersion": "18.39.41",
            "osVersion": "12",
            "androidSdkVersion": 32
         }
      },
      "videoId": "oCjW6gdEDa4"
   }
   `
   req, err := http.NewRequest(
      "POST", "https://youtubei.googleapis.com/youtubei/v1/player",
      strings.NewReader(s),
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "com.google.android.youtube/16.49.39")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if strings.Contains(string(data), "This video requires payment to watch") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
