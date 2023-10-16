package youtube

import (
   "154.pages.dev/encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
)

func (c contents) show() (*row, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         for _, v := range v.
         Metadata_Row_Container.
         Metadata_Row_Container_Renderer.
         Rows {
            if v.Metadata_Row_Renderer.Title.String() == "Show" {
               return &v, true
            }
         }
      }
   }
   return nil, false
}

type row struct {
   Metadata_Row_Renderer struct {
      Title value // Show
      Contents values // In The Heat Of The Night
   } `json:"metadataRowRenderer"`
}

func (c contents) String() string {
   var b []byte
   v, ok_show := c.show()
   if ok_show {
      b = fmt.Append(b, v.Metadata_Row_Renderer.Contents)
   }
   if v, ok := c.title(); ok {
      if ok_show {
         b = append(b, " • "...)
      }
      b = append(b, v...)
   }
   return string(b)
}

func (c contents) title() (string, bool) {
   for _, v := range c {
      if v := v.Video_Primary_Info_Renderer; v != nil {
         return v.Title.String(), true
      }
   }
   return "", false
}

type contents []struct {
   Video_Primary_Info_Renderer *struct {
      Title value // The Family Secret
   } `json:"videoPrimaryInfoRenderer"`
   Video_Secondary_Info_Renderer *struct {
      Metadata_Row_Container struct {
         Metadata_Row_Container_Renderer struct {
            Rows []row
         } `json:"metadataRowContainerRenderer"`
      } `json:"metadataRowContainer"`
   } `json:"videoSecondaryInfoRenderer"`
}

// /youtubei/v1/player is missing the name of the series. we can do
// /youtubei/v1/next but the web client is smaller response
func make_contents(video_ID string) (contents, error) {
   res, err := http.Get("https://www.youtube.com/watch?v=" + video_ID)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = json.Cut(text, []byte(" ytInitialData ="), nil)
   var s struct {
      Contents struct {
         Two_Column_Watch_Next_Results struct {
            Results struct {
               Results struct {
                  Contents contents
               }
            }
         } `json:"twoColumnWatchNextResults"`
      }
   }
   if err := json.Unmarshal(text, &s); err != nil {
      return nil, err
   }
   return s.Contents.Two_Column_Watch_Next_Results.Results.Results.Contents, nil
}

type value struct {
   Runs []struct {
      Text string
   }
   Simple_Text string `json:"simpleText"`
}

func (v value) String() string {
   if v.Simple_Text != "" {
      return v.Simple_Text
   }
   var b []byte
   for _, run := range v.Runs {
      b = append(b, run.Text...)
   }
   return string(b)
}

type values []value

func (v values) String() string {
   var b []byte
   for _, val := range v {
      if b != nil {
         b = append(b, ", "...)
      }
      b = fmt.Append(b, val)
   }
   return string(b)
}
