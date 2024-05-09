package joyn

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

const query_episode = `
query EpisodeDetailPageStatic($path: String!) {
   page(path: $path) {
      ... on EpisodePage {
         episode {
            ... on Episode {
               video {
                  id
               }
               series {
                  title
               }
               season {
                  ... on Season {
                     number
                  }
               }
               number
               title
            }
         }
      }
   }
}
`

func (e *episode_detail) New(path string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.Query = query_episode
      s.Variables.Path = path
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://api.joyn.de/graphql", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "joyn-platform": {"web"},
      "x-api-key": {"4f0fd9f18abbe3cf0e87fdb556bc39c8"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   err = json.Unmarshal(text, e)
   if err != nil {
      return err
   }
   if e.Data.Page.Episode == nil {
      return errors.New(string(text))
   }
   return nil
}

type episode_detail struct {
   Data struct {
      Page struct {
         Episode *struct {
            Video struct {
               ID string
            }
            Series struct {
               Title string
            }
            Season struct {
               Number int
            }
            Number int
            Title string
         }
      }
   }
}
