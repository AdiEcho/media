package ctv

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               id
               ... on AxisMedia {
                  firstPlayableContent {
                     id
                  }
               }
            }
         }
      }
   }
}
`

func new_resolve(path string) (*ResolvePath, error) {
	body, err := func() ([]byte, error) {
		var s struct {
			OperationName string `json:"operationName"`
			Query         string `json:"query"`
			Variables     struct {
				Path string `json:"path"`
			} `json:"variables"`
		}
		s.OperationName = "resolvePath"
		s.Variables.Path = path
		s.Query = query_resolve
		return json.Marshal(s)
	}()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST", "https://www.ctv.ca/space-graphql/apq/graphql",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	// you need this for the first request, then can omit
	req.Header.Set("graphql-client-platform", "entpay_web")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var s struct {
		Data struct {
			ResolvedPath struct {
				LastSegment struct {
					Content ResolvePath
				}
			}
		}
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s.Data.ResolvedPath.LastSegment.Content, nil
}

type ResolvePath struct {
	ID                   string
	FirstPlayableContent *struct {
		ID string
	}
}

func (r ResolvePath) id() string {
	if v := r.FirstPlayableContent; v != nil {
		return v.ID
	}
	return r.ID
}
