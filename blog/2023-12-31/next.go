package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

// /youtubei/v1/player is missing the name of the series
func make_contents(videoId string) (contents, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         VideoId string `json:"videoId"`
         Context struct {
            Client struct {
               ClientName string `json:"clientName"`
               ClientVersion string `json:"clientVersion"`
            } `json:"client"`
         } `json:"context"`
      }
      s.VideoId = videoId
      s.Context.Client.ClientName = "WEB"
      s.Context.Client.ClientVersion = "2.20231219.04.00"
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://www.youtube.com/youtubei/v1/next?prettyPrint=false" ,
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
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

func (c contents) title() (string, bool) {
   for _, v := range c {
      if v := v.Video_Primary_Info_Renderer; v != nil {
         return v.Title.String(), true
      }
   }
   return "", false
}

func (c contents) season() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Season")
   }
   return "", false
}

func (c contents) episode() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Episode")
   }
   return "", false
}

func (c contents) owner() (string, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return v.Owner.Video_Owner_Renderer.Title.String(), true
      }
   }
   return "", false
}

func (c contents) show() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Show")
   }
   return "", false
}

func (v values) String() string {
   var b strings.Builder
   for _, val := range v {
      if b.Len() >= 1 {
         b.WriteString(", ")
      }
      b.WriteString(val.String())
   }
   return b.String()
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

func (m metadata_row_container) get(s string) (string, bool) {
   for _, v := range m.Metadata_Row_Container_Renderer.Rows {
      if v := v.Metadata_Row_Renderer; v.Title.String() == s {
         return v.Contents.String(), true
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

func (c contents) metadata_row_container() (*metadata_row_container, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return &v.Metadata_Row_Container, true
      }
   }
   return nil, false
}

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

type values []value

type value struct {
   Runs []struct {
      Text string
   }
   Simple_Text string `json:"simpleText"`
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

