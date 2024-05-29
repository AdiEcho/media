package roku

import (
   "net/http"
   "net/url"
   "os"
)

func four() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "googletv.web.roku.com"
   req.URL.Path = "/api/v1/account/token"
   req.URL.Scheme = "https"
   req.Header["X-Roku-Content-Token"] = []string{"gGYqeHG716gvJZBVV45adR9dWXXumkiNI4DQLgDE3oImJ0gX0SkbjQkByHGgWJyvDHho169GtwtkhyTLS/0Apdzq6kU0VWhwraXjL0Tm7Z72ndpRif1R5WJ94SO2cPrlcnCGhgoQlAmMJvhFTdM0qXbc+TqzsIEUdEQ4BkHb2BvdhUe7Y9cVtGMB6qvJAljX18KiTL4A2pD/tjf0NFo34E0zlJiwUj+f3VBnRDk68XbMDoZrIDWSX+NRm0V4eTUTA3XX+DxHYRl0mV+ZD9e402LyGNhJOcqA2EyI3Ev7ihmQKKIb8kzMq9SvANlVttmGDPX9Hx7OpnTePgeSiMXiR04de6laYyOtwfSN4/G0hmeu3c9Fb5Z/D/sXUjidUKcMsy+X0HK9a3myQ77QhqLb2cCVyl4buS2N1m2zcZuoZF/axG4ueO6b913u3LgBBq0s+NwdXap+FJGLYZtOawL2WjioEYLdqeywFJMvLlXrNXw9dufjnxLxyCDux5JA7Wy+b08JeymIn0Qnhtf17imf0x78dMUzyYN8iSFnFvuublo="}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
