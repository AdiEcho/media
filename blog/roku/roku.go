package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "googletv.web.roku.com"
   req.URL.Path = "/api/v3/playback"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"}
   req.Header["X-Roku-Content-Token"] = []string{"ZizZafDEidXT//m06oXf0aepPFcxQ+HTY9iyMtVN7JbWpq6XcHp3hrkdZO6PKbCL3WVNoPkEl95XWEOMGVJEZzRtqFKvfvzYpMyNsNU0+WXJgVrQi8NlHcXKx227Ptx3/UyOwUx1bwP9TKVIQATnPuyW5VKsT0RAXuT+Qb2jAFs+x0ozsOyWzmQr3DAtrUNG5qR/7aGFi/0vasW6MsmD49AXxurDMoMW3N6w749+kUN/Mn+vVVmhGUIurxVpR30y7GYpMo4M6pV2TVn6NYC5HgqRP6l+e1/KugHi8L5cDIBSNngQVO+ASNPQj8rOv11H7E/qVNzprrwaFb8aOyjgjjob9N6WLQOU8I8gLGmkwDixppYtuCGoNY2ooh141NOnlqUvzWj9yfHFkeOWhosHVQ=="}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`
{
   "mediaFormat": "DASH",
   "rokuId": "597a64a4a25c5bf6af4a8c7053049a6f"
}
`)
