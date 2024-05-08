package joyn

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (m *movie_detail) New(path string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Path string `json:"path"`
         } `json:"variables"`
      }
      s.Query = page_movie
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
   return json.NewDecoder(res.Body).Decode(m)
}

type movie_detail struct {
   Data struct {
      Page struct {
         Movie struct {
            Video struct {
               ID string
            }
         }
      }
   }
}

const page_movie = `
query PageMovieDetailStatic($path: String!) {
   page(path: $path) {
      ... on MoviePage {
         movie {
            ... on Movie {
               ... on Movie {
                  video {
                     id
                  }
               }
            }
         }
      }
   }
}
`
