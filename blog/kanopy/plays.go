package kanopy

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func plays(video_id int) (*http.Response, error) {
   data, err := json.Marshal(map[string]int{
      "domainId": 2918,
      "userId": 8177465,
      "videoId": video_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/plays", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMyODIzODAzOTUzMDMxNzE5Iiwic2Vzc2lvbl9pZCI6IjE3MzI4MjM4MDM5NTMwODc0MDMiLCJjb25uZWN0aW9uX2lkIjoiMTczMjgyMzgwMzk1MzA4NzQwMyIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzI4MjM4MDMsImV4cCI6MjA0ODE4MzgwMywiaXNzIjoia2FwaSJ9.M6n5KPLzsLE1U8xuWc1tQ_gCAIUCFb4BtJXQDHd07m8"},
      "Content-Type": {"application/json"},
      "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"},
      "X-Version": {"web/prod/4.16.0/2024-11-07-14-23-23"},
   }
   return http.DefaultClient.Do(req)
}
