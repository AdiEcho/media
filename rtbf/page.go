package rtbf

import "net/http"

func page(path string) (*http.Response, error) {
   return http.Get("https://bff-service.rtbf.be/auvio/v1.23/pages" + path)
}
