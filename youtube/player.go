package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "time"
)

func (r Request) Player(tok *Token) (*Player, error) {
   r.Context.Client.Android_SDK_Version = 32
   r.Context.Client.OS_Version = "12"
   body, err := json.Marshal(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", user_agent + r.Context.Client.Client_Version)
   if tok != nil {
      req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Player struct {
   Microformat struct {
      Player_Microformat_Renderer struct {
         Publish_Date string `json:"publishDate"`
      } `json:"playerMicroformatRenderer"`
   }
   Playability struct {
      Status string
      Reason string
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

// stream.Video
func (p Player) Author() string {
   return p.Video_Details.Author
}

func (p Player) Time() (time.Time, error) {
   return time.Parse(
      time.RFC3339, p.Microformat.Player_Microformat_Renderer.Publish_Date,
   )
}

// stream.Video
func (p Player) Title() string {
   return p.Video_Details.Title
}
