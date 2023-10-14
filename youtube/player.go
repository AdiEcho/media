package youtube

import "time"

type Player struct {
   Microformat struct {
      Player_Microformat_Renderer struct {
         Publish_Date string `json:"publishDate"`
      } `json:"playerMicroformatRenderer"`
   }
   Playability_Status struct {
      Reason string
      Status string
   } `json:"playabilityStatus"`
   Streaming_Data struct {
      Adaptive_Formats []Format `json:"adaptiveFormats"`
   } `json:"streamingData"`
   Video_Details struct {
      Author string
      Length_Seconds int64 `json:"lengthSeconds,string"`
      Short_Description string `json:"shortDescription"`
      Title string
      Video_ID string `json:"videoId"`
      View_Count int64 `json:"viewCount,string"`
   } `json:"videoDetails"`
}

func (p Player) Author() string {
   return p.Video_Details.Author
}

func (p Player) Duration() time.Duration {
   return time.Duration(p.Video_Details.Length_Seconds) * time.Second
}

func (p Player) Time() (time.Time, error) {
   return time.Parse(
      time.RFC3339, p.Microformat.Player_Microformat_Renderer.Publish_Date,
   )
}

func (p Player) Title() string {
   return p.Video_Details.Title
}
