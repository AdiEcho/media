package internal

import "net/http"

type ClientRequest interface {
   DashClient() http.Client
   DashRequest() (*http.Request, error)
}
