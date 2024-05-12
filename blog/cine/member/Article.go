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
    ...Article
  }
}

fragment Article on Article {
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
    ...Asset
  }
  images {
    ...ArticleFile
  }
  posters {
    ...ArticleFile
  }
  metas(output: html) {
    ...ArticleMeta
  }
  categories {
    ...ArticleCategory
  }
  children {
    ...ChildArticle
  }
  recommended_articles {
    ...ArticleRelated
  }
  published_at
  products {
    ...ProductListed
  }
  upsell_products {
    ...ProductListed
  }
  upsell_product_call_to_action_tag
  is_downloadable
  ribbon_title
  ribbon_settings
}

fragment ArticleMeta on ArticleMeta {
  key
  value
}

fragment Asset on Asset {
  id
  duration
  linked_type
  accessibility
  screenshots {
    ...File
  }
}

fragment File on File {
  type
  url
  title
  base_url
  file_name
  width
  height
}

fragment ArticleFile on File {
  type
  url
  base_url
  file_name
  title
  width
  height
}

fragment ArticleCategory on CategoryListedAsRelation {
  parent_id
  id
  name
  type
  metas {
    key
    value
  }
}

fragment ChildArticle on Article {
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
    ...Asset
  }
  images {
    ...ArticleFile
  }
  posters {
    ...ArticleFile
  }
  metas(output: html) {
    ...ArticleMeta
  }
  categories {
    ...ArticleCategory
  }
  published_at
  products {
    ...ProductListed
  }
  upsell_products {
    ...ProductListed
  }
  recommended_articles {
    ...ArticleRelated
  }
  upsell_product_call_to_action_tag
  is_downloadable
  ribbon_title
  ribbon_settings
}

fragment ProductListed on ProductListedAsRelation {
  id
  remote_product_id
  name
  title
  description(output: html)
  description_short(output: html)
  price
  credit_price
  currency
  currency_symbol
  expires_in
  call_to_action_tag
  articles {
    ...ArticleProductListed
  }
  type
}

fragment ArticleProductListed on ArticleListedAsRelation {
  id
  name
  type
}

fragment ArticleRelated on ArticleRelated {
  id
  type
  name
  ribbon_title
  ribbon_settings
  canonical_url
  canonical_title
  url_slug
  full_url_slug
  metas {
    ...ArticleMeta
  }
  images {
    ...ArticleFile
  }
  posters {
    ...ArticleFile
  }
  assets {
    ...Asset
  }
  categories {
    ...ArticleCategory
  }
  published_at
  parent_id
  is_downloadable
  is_auth_required
}
`
