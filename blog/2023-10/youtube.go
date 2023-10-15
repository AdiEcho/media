package youtube

import (
   "fmt"
   "io"
   "net/http"
   "strings"
   option "154.pages.dev/http"
)

type watch struct {
   Player_Overlays struct {
      Player_Overlay_Renderer struct {
         Autoplay struct {
            Player_Overlay_Autoplay_Renderer struct {
               Video_Title struct {
                  Simple_Text string simpleText
               } videoTitle
            } playerOverlayAutoplayRenderer
         }
         Video_Details struct {
            Player_Overlay_Video_Details_Renderer struct {
               Title struct {
                  Simple_Text string simpleText
               }
            } playerOverlayVideoDetailsRenderer
         } videoDetails
      } playerOverlayRenderer
   } playerOverlays
}

func main() {
   option.No_Location()
   option.Verbose()
   // /youtubei/v1/player is missing the name of the series. we can do
   // /youtubei/v1/next but the web client is smaller response
   res, err := http.Get("https://www.youtube.com/watch?v=2ZcDwdXEVyI")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      panic(res.Status)
   }
   text, err := func() (string, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return "", err
      }
      return string(b), nil
   }()
   if err != nil {
      panic(err)
   }
   fmt.Println(len(text))
   if strings.Contains(text, "In The Heat Of The Night") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
