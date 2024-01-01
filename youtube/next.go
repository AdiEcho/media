package youtube

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type Contents []struct {
   Video_Primary_Info_Renderer *struct {
      Title Value // The Family Secret
   } `json:"videoPrimaryInfoRenderer"`
   Video_Secondary_Info_Renderer *struct {
      Metadata_Row_Container Metadata_Row_Container `json:"metadataRowContainer"`
      Owner struct {
         Video_Owner_Renderer struct {
            Title Value
         } `json:"videoOwnerRenderer"`
      }
   } `json:"videoSecondaryInfoRenderer"`
}

func (c Contents) Episode() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Episode")
   }
   return "", false
}

func (c Contents) Owner() (string, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return v.Owner.Video_Owner_Renderer.Title.String(), true
      }
   }
   return "", false
}

func (c Contents) Release_Date() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Release date")
   }
   return "", false
}

func (c Contents) Season() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Season")
   }
   return "", false
}

func (c Contents) Show() (string, bool) {
   if v, ok := c.metadata_row_container(); ok {
      return v.get("Show")
   }
   return "", false
}

func (c Contents) Title() (string, bool) {
   for _, v := range c {
      if v := v.Video_Primary_Info_Renderer; v != nil {
         return v.Title.String(), true
      }
   }
   return "", false
}

// /youtubei/v1/player is missing the name of the series. use with WEB client.
func (c *Contents) Next(r Request) error {
   body, err := json.Marshal(r)
   if err != nil {
      return err
   }
   res, err := http.Post(
      "https://www.youtube.com/youtubei/v1/next?prettyPrint=false" ,
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   var s struct {
      Contents struct {
         Two_Column_Watch_Next_Results struct {
            Results struct {
               Results struct {
                  Contents Contents
               }
            }
         } `json:"twoColumnWatchNextResults"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return err
   }
   *c = s.Contents.Two_Column_Watch_Next_Results.Results.Results.Contents
   return nil
}

func (c Contents) metadata_row_container() (*Metadata_Row_Container, bool) {
   for _, v := range c {
      if v := v.Video_Secondary_Info_Renderer; v != nil {
         return &v.Metadata_Row_Container, true
      }
   }
   return nil, false
}

type Metadata_Row_Container struct {
   Metadata_Row_Container_Renderer struct {
      Rows []struct {
         Metadata_Row_Renderer struct {
            Title Value // Show
            Contents Values // In The Heat Of The Night
         } `json:"metadataRowRenderer"`
      }
   } `json:"metadataRowContainerRenderer"`
}

func (m Metadata_Row_Container) get(s string) (string, bool) {
   for _, v := range m.Metadata_Row_Container_Renderer.Rows {
      if v := v.Metadata_Row_Renderer; v.Title.String() == s {
         return v.Contents.String(), true
      }
   }
   return "", false
}

type Value struct {
   Runs []struct {
      Text string
   }
   Simple_Text string `json:"simpleText"`
}

func (v Value) String() string {
   if v.Simple_Text != "" {
      return v.Simple_Text
   }
   var b strings.Builder
   for _, run := range v.Runs {
      b.WriteString(run.Text)
   }
   return b.String()
}

type Values []Value

func (v Values) String() string {
   var b strings.Builder
   for _, val := range v {
      if b.Len() >= 1 {
         b.WriteString(", ")
      }
      b.WriteString(val.String())
   }
   return b.String()
}
