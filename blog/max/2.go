package max

import (
   "net/http"
   "net/url"
   "os"
)

func Two() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "default.prd.api.discomax.com"
   req.URL.Path = "/authentication/linkDevice/initiate"
   req.URL.Scheme = "https"
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi0yYjhjODBkZi03OGYzLTRlN2ItOWE1Ny04NjcwOWJlMThlMmEiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6M2NiZGZlY2UtZWJlMi00ZDYzLWEzNzYtODRiZjNmNDgyNDlkIiwiaWF0IjoxNzI5NDY2MjQ4LCJleHAiOjIwNDQ4MjYyNDgsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6dHJ1ZSwiZGV2aWNlSWQiOiIwNDEzN2FhMi0xZTFlLTZmNTItN2EwOC0xMjI0OWM4NjQ2OTAifQ.GY145CIXxN2WBac2xL18nvZmSfQyxKukBfAfjRCt_L3N5g5TzBTFgHS2V6ppCK9E6i0rgKIeYiEf2UgWJj8saZYyuNubTyu-JO326dbV6KmKA48rMDnO0dHVC5cwPM0JJmU55IoKUDyg4qV06kR7Z5_R-QRmyXIBXSK5kYEqmp_LtZsOcXeQi7UuHhoWpkr4p41veIEvfrhrf_yD5SMadbfDZI61yDl0xeaBidjqODvfeZiHhJ6PvVfK_dEhvvJobxotEni76s9zSPUtU0c7zT4YLc_B414N2PA9AFDYa5Ro5E5LaFyj4SSAx9mJoP9qxy4TIMQeQE8FS1SYjZzzkg",
   }
   req.Header["X-Device-Info"] = []string{
      "beam/1.1.2.2 (LG/OLED55C9PVA; webOS/4.9.0-05.00.03; 04137aa2-1e1e-6f52-7a08-12249c864690/9e1b83a7-ddde-42c9-b335-a54232bd2a9f)",
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

