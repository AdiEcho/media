package roku

import "net/http"

const user_agent = "Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"

func one() (*http.Response, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/token"
   req.Header["User-Agent"] = []string{user_agent}
   return http.DefaultClient.Do(req)
}
