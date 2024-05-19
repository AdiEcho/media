package criterion

import (
   "net/http"
   "net/url"
)

func videos() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E2NDUzNDE2ZTlmZmRjNjk4MDAxMDdhZTQ2ZWJhODg0YTU2ZDBjOGQ4NTYzMzgifQ.eyJhcHBfaWQiOjM0NDksImV4cCI6MTcxNjE0NzQyNiwibm9uY2UiOiI1MWRlYzhmYTE4NjZiYjI5Iiwic2NvcGVzIjpbXSwic2Vzc2lvbl9pZCI6InorL2I5aC9laUVCNGkxZmR2VEw5a2c9PSIsInNpdGVfaWQiOjU5MDU0LCJ1c2VyX2lkIjo2NTc0NjE0Nn0.cIAPnjuaxgRj7rHJcVaFlZsiPkLt-6e6AdKCutEKp2z9JFaSPXkbOn6VS_z6tgcNtlzPmrK08G-A_nvLrdAPurojQMGzJYvY4853ni8tRsvdqvzHyGK9m6Zarcyg9gPJa-Sj5b4x1VtUuWeo4aZ7WrF8pXriptULPL-VZj9XnpX4BRK4xAq6AAVlFVjjOv8MD0_ry3KVKNTxgDgtvFumjO1qgjK3JcLRSmHmxUyjzp3_xtulFBFdYNcTs0ivNnLGbaIEtVfujyDnmrFsetiQyNMxEwAR-CLkdZxMMRZdLTiAgAOd5lPJAmTFIY3EHuB3u6y3IuSoI1UsTY7gXduG78jxFu6045XlSSj-fWAwlYzqwhI-Hi7JJsPMPlh23q4cS3e6KAG_1Jg2EcFCuCr0aXDJjm43UUrt9zXwdaMs7O8VpzqbxF1YDp-lZwNLku8kUxvgvD_krdISWa0c-tN_NIon44sbIJBZj0EfrWAb_oGswPKu4_Ng8Urgi26g9HCt8HYalfuSgW_QZzZ5Wq0czK-Gpz7H0V1vmcYJAweU7tc2pws84fC8GZSSbSzzNcUD-fhKwHh_qeWxj_L8RfGG7O65lGBNoeI1CsM6luGcS-ioc0_fSMnxU4Xd8_RQw-127-zT1aTlCArPQZFM7Mb7xooX2tCjsUAma691G6lP7A4"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.vhx.com"
   req.URL.Scheme = "http"
   req.URL.Path = "/videos/my-dinner-with-andre"
   req.URL.RawQuery = "url=my-dinner-with-andre"
   return http.DefaultClient.Do(&req)
}
