package youtube

import (
   "154.pages.dev/encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type contents []struct {
   Video_Primary_Info_Renderer *struct {
      Title value // The Family Secret
   } `json:"videoPrimaryInfoRenderer"`
   Video_Secondary_Info_Renderer *struct {
      Metadata_Row_Container metadata_row_container `json:"metadataRowContainer"`
      Owner struct {
         Video_Owner_Renderer struct {
            Title value
         } `json:"videoOwnerRenderer"`
      }
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

func (c contents) String() string {
   var b strings.Builder
   date, date_ok := c.release_date()
   show, show_ok := c.show()
   if !date_ok {
      if v, ok := c.owner(); ok {
         b.WriteString(v)
      }
   }
   if show_ok {
      b.WriteString(show)
   }
   if v, ok := c.season(); ok {
      b.WriteByte(' ')
      b.WriteString(v)
   }
   if v, ok := c.episode(); ok {
      b.WriteByte(' ')
      b.WriteString(v)
   }
   if v, ok := c.title(); ok {
      if b.Len() >= 1 {
         b.WriteString(" - ")
      }
      b.WriteString(v)
   }
   if !show_ok {
      if date_ok {
         b.WriteString(" - ")
         b.WriteString(date)
      }
   }
   return b.String()
}

func (c contents) episode() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Episode")
   }
   return "", false
}

func (c contents) metadata_row_container() (*metadata_row_container, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return &v.Metadata_Row_Container, true
      }
   }
   return nil, false
}

func (c contents) owner() (string, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return v.Owner.Video_Owner_Renderer.Title.String(), true
      }
   }
   return "", false
}

func (c contents) release_date() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Release date")
   }
   return "", false
}

func (c contents) season() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Season")
   }
   return "", false
}

func (c contents) show() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Show")
   }
   return "", false
}

func (c contents) title() (string, bool) {
   for _, v := range c {
      if v := v.Video_Primary_Info_Renderer; v != nil {
         return v.Title.String(), true
      }
   }
   return "", false
}

type metadata_row_container struct {
   Metadata_Row_Container_Renderer struct {
      Rows []struct {
         Metadata_Row_Renderer struct {
            Title value // Show
            Contents values // In The Heat Of The Night
         } `json:"metadataRowRenderer"`
      }
   } `json:"metadataRowContainerRenderer"`
}

func (m metadata_row_container) get(s string) (string, bool) {
   for _, v := range m.Metadata_Row_Container_Renderer.Rows {
      if v := v.Metadata_Row_Renderer; v.Title.String() == s {
         return v.Contents.String(), true
      }
   }
   return "", false
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
      b = append(b, val.String()...)
   }
   return string(b)
}
