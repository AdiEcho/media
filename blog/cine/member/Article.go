package member

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func article() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.audienceplayer.com"
   req.URL.Path = "/graphql/2/user"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   body := fmt.Sprintf(`
   {
      "query": %q,
      "variables": {
         "articleUrlSlug": "films/american-hustle"
      }
   }
   `, query_article)
   req.Body = io.NopCloser(strings.NewReader(body))
   return http.DefaultClient.Do(&req)
}

const query_article = `
query Article($articleId: Int, $articleUrlSlug: String) {
   Article(id: $articleId, full_url_slug: $articleUrlSlug) {
      ... on Article {
         id
         url_slug
         full_url_slug
         canonical_url
         canonical_title
         name
         type
         is_auth_required
         parent_id
         assets {
            ... on Asset {
               id
               duration
               linked_type
               accessibility
            }
         }
         metas(output: html) {
            ... on ArticleMeta {
               key
               value
            }
         }
         published_at
         upsell_product_call_to_action_tag
         is_downloadable
         ribbon_title
         ribbon_settings
      }
   }
}
`
