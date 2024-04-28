package ctv

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const query_axis = `
query axisContent($id: ID!) {
   axisContent(id: $id) {
      axisId
      axisPlaybackLanguages {
         ... on AxisPlayback {
            destinationCode
         }
      }
   }
}
`

type axis_content struct {
   Data struct {
      AxisContent struct {
         AxisId int
         AxisPlaybackLanguages []struct {
            DestinationCode string
         }
      }
   }
}

func (a *axis_content) New(id string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         OperationName string `json:"operationName"`
         Query string `json:"query"`
         Variables struct {
            ID string `json:"id"`
         } `json:"variables"`
      }
      s.OperationName = "axisContent"
      s.Variables.ID = id
      s.Query = query_axis
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
