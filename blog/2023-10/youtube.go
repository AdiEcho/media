package youtube

import (
   "154.pages.dev/encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type content struct {
   Video_Primary_Info_Renderer *struct {
      Title value // The Family Secret
   } `json:"videoPrimaryInfoRenderer,omitempty"`
   Video_Secondary_Info_Renderer *struct {
      Metadata_Row_Container struct {
         Metadata_Row_Container_Renderer struct {
            Rows []struct {
               Metadata_Row_Renderer struct {
                  Title value // Show
                  Contents values // In The Heat Of The Night
               } `json:"metadataRowRenderer"`
            } `json:",omitempty"`
         } `json:"metadataRowContainerRenderer"`
      } `json:"metadataRowContainer"`
   } `json:"videoSecondaryInfoRenderer,omitempty"`
}

// /youtubei/v1/player is missing the name of the series. we can do
// /youtubei/v1/next but the web client is smaller response
func contents(video_ID string) ([]content, error) {
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
                  Contents []content
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
   } `json:",omitempty"`
   Simple_Text string `json:"simpleText,omitempty"`
}

func (v value) String() string {
   if v.Simple_Text != "" {
      return v.Simple_Text
   }
   var b strings.Builder
   for _, run := range v.Runs {
      b.WriteString(run.Text)
   }
   return b.String()
}

type values []value

func (v values) String() string {
   var b strings.Builder
   for _, s := range v {
      if b.Len() >= 1 {
         b.WriteString(", ")
      }
      b.WriteString(s.String())
   }
   return b.String()
}
