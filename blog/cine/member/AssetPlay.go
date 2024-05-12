package member

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "fmt"
)

const query_asset = `
mutation($article_id: Int, $asset_id: Int) {
   ArticleAssetPlay(article_id: $article_id asset_id: $asset_id) {
      entitlements {
         ... on ArticleAssetPlayEntitlement {
            manifest
            protocol
         }
      }
   }
}
`

func asset_play() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.audienceplayer.com"
   req.URL.Path = "/graphql/2/user"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiZDNhZDQzNzRhZTdkMWY5OTkyZmRhZGRkY2NiZTI0YTIwYTFiNjdiODg0YjNjYzJlOTM4MmQwZWU3YzQzNTdiZmQ1NjRmOWEzMGI0OWQzMjAiLCJpYXQiOjE3MTU1MzMzMTAsIm5iZiI6MTcxNTUzMzMxMCwiZXhwIjoyMDMwODkzMzEwLCJzdWIiOiIxMjM4NjMiLCJzY29wZXMiOlsiYXBpLXVzZXItYWNjZXNzIl0sImFwX3BpZCI6MiwiYXBfYWlkIjpudWxsLCJhcF9yaWQiOm51bGwsImFwX2tpZCI6bnVsbH0.tM2GLP7yGtT2hLyPteXJAEahSmMDdTWhi28A_8oLwf7U3aHmmyZrPSfk2Rwceai9jVu8HiDre8_JbXmr6gS7v7M2nur77cSkUAXA0IYfgdhKjO67YWmyCDzN27fh_Gur4je-uNcT8dw0gi4kURoXcnkjB6Er3AV8ktpPaXbRtmdeMBVzkNTAcUftvkfgGoftE6oUuFoSnL5Ra40JICAqHPiqSTtACRRxvJjSPSP9zm1oaH07Bj2oeQX711hhxZWvq1eXkr89VP984xGypOJYWkAA_g6HYH3TVupWpEmNlqov1h20PtHTekhcjh1lhmEr_dIY0n3QHogj9wQY8TRHG49Vl8p7Gi7a885ElEcU6OC9FJnU_lgT6_xbZxuLUZoxridDF6ikvCZA4WS91RiuHc9N8Nfy4SYPk0KYHP60bXC_qhMdYdcCY4u3RHhlVuRdr6YBmAbvWzTDogoKCckatBRuKnZLBOqy2Yvl7y02iM2wW0b4b2iE78aonmcGZcDDOT8iK39v8JQBwfKJfaPbKbUeC3MZXoU-a-DMKK8CcpTbTwUkNJrkY9D14rCjAtM4myHebNCs5rj9z8FgWAd205wfQuX2D5-0PZn0BACRvqpvijM7QJlFTwwA6NWmeOY7b8GVV-A07sG58U0Oal0nO2-VT_DC0xX4ZlwyWmOajKM"}
   req.Header["Content-Type"] = []string{"application/json"}
   body := fmt.Sprintf(`
   {
      "query": %q,
      "variables": {
         "article_id": 768,
         "asset_id": 1415
      }
   }
   `, query_asset)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
