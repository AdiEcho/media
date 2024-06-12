package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.Header["Cookie"] = []string{"st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi05ZGQxYjYzOS01ODVmLTQwYzAtYjE5Ny1kNTkxYWMyOWQxYWEiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6Y2YwMWI0ZDItZDIyNS00Njc4LThkOTItOGU0NTg1MDhkN2U4IiwiaWF0IjoxNzE4MTYwMjA5LCJleHAiOjIwMzM1MjAyMDksInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6ZmFsc2UsImRldmljZUlkIjoiISJ9.YQAxSitkbxkUy97tq6SVUwK3alimjBirNUntLTvsT3xxqBf0c7mq7gcB-OrU1cEf5nSjBqoYw54ZF-4fqFVul8zuJqY0bq3Yd9A-lyFiHOHzgP1zK1OIaCVNxkLL7953KNE6UMl_L9H-EVa0-DQHixD10JMZxCxu2PYLuNeMXbe5aMh_RAgmtc0YV_7dj5sIugw_zbYuU5oNg0LEIGppQ1skCV8X2p44ReIX65bX7-ady54a3wKdpkl5xyMoYOB28srSTi77mpUTGtGuxQZMdEPADXzblf9KT46N1HWajIvenlwwPfsU5-XcF6u7Sn5vwXmnUGo1788c3ARqaR-W-Q"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL.Host = "default.any-amer.prd.api.max.com"
   req.URL.Path = "/content/videos/127b00c5-0131-4bac-b2d1-40762deefe09/activeVideoForShow"
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
