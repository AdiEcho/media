package joyn

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func entitlement() {
   // from PageMovieDetailStatic
   var body = strings.NewReader(`
   {
      "content_id": "a_p4svn4a28fq"
   }
   `)
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "entitlement.p7s1.io"
   req.URL.Path = "/api/user/entitlement-token"
   req.URL.Scheme = "https"
   // from /auth/anonymous
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImEwZDQwYjkxZTA2OGEzY2ZhODQ1ZjRkZTViNmY3NjA2NmEzMzc3NTEifQ.eyJkaXN0cmlidXRpb25fdGVuYW50IjoiSk9ZTiIsImNvdW50cnkiOiJERSIsImNhdEN0eSI6IkRFIiwibG9jQ3R5IjoiREUiLCJwcmZMbmciOiJkZSIsImF2b2RDdHkiOiJERSIsImF2b2RFbmwiOnRydWUsInN2b2RDdHkiOiJERSIsInN2b2RFbmwiOnRydWUsImpJZCI6IjcxZGI5YWIyM2ZlMWY1OTA1MTk0ZWNlNWY5NjYyNTBiOjU3Yzk0ODNhMTdlODkxZThhYjQ0MGZkNzcyNTkwODVmZTk0ODRiOTJmMGMwN2RhODkzYTlhYWE0NmQ4MzM0Yjg3YTY1N2EwMGUxODNiMzg3NTJjYmNjMzMyNzBlODZiMyIsInBJZCI6ImExMzcyMjRjNzE0MmM5NjEzMjllZjExMWNkODNmNjlmOmNkOTU5NzViOTQzMGJiNzgxYTA4ZGEyOWRiZDYyNmQ0MDEyODBiMjMzODZkZWNiYzZjOWI2YjUyYmU1ZWJmNTE1YWQ3MDIwMTY1MTU0NDkzMmIyMDk0NzlkOGNjMDJmOCIsImpJZEMiOiJKTkFBLWFlM2Y4MGNmLTU3NDMtNTIxZS05MDkyLTA5MmE4Nzc5MmM4ZCIsInBJZEMiOiJKTkFBLWFlM2Y4MGNmLTU3NDMtNTIxZS05MDkyLTA5MmE4Nzc5MmM4ZCIsImNJZCI6IjZiZDE2YjQ4LWE0OGEtNDAwZS05Y2RiLTI3ZjMzOTFlMmU0ZSIsImNOIjoid2ViIiwiZW50IjoiZWM6MCxmOjAsZmw6MCIsImludGVybmFsIjpmYWxzZSwiZWxpZ2libGVGb3JFbXBsb3llZVN1YnNjcmlwdGlvbiI6ZmFsc2UsImlhdCI6MTcxNTEyMjE0NywiZXhwIjoxNzE1MjA4NTQ3LCJhdWQiOlsid2ViIl0sImlzcyI6Imh0dHBzOi8vam95bi5kZS8iLCJzdWIiOiIyMTI5ZmY4OS1hOGQzLTQxOGMtOTkyNy00ZGY1ZDU1ZDUwNTgifQ.HxWomuc00pn2o4tacXAMdZGrLwlonw_fPfoAlMy2dKuMZ9Lp3v0CsEC1E6HuyLZ38UBbeotwKKSymDG8Rlyn3xiEjyYwqO7pEjLwKedDIgEb1m7AD-KchPjj1xAXZMDOw2IBPBsDqgQ3LdC33lNgWpng9rp-6irNcZFHM_KnfjFW0M_RL3jPlJzHWgbi1PamxsxroEPwBSZjup_xBgRQ9_phatgBezjlbJMJEk1-6eZ6WWn-8yJLyDEp1qD2ObyVE6mKtcHa6agooUlsr8YZ72wzSlV5yRIWKta_qikAxxVy0oI5OqsBU46aQPyoNdpenXY3OllLq3ZYWAWCT3Gxnho5UEUaFW7MtIuvwtwia4eRg2_piZg1dSLPRt0NRtBe_x4j0KX1UDxnGZzU6REIQvE2CxpTNeAXvz7gVqzdN0ZY_nar_BwuY9Gq2PKAU9wyTm9Dyy1YdydFd3pucbPtKM8iLZdE_wDNYFgG-Yi1vpsCfO0_FHokY2l-9L_r198Htk5z-gnE-3_sDxuP7wQlYpIYGdMz6vGsAcVNdbWE8VXTYgiT-zyYtuzUJFFQItI_4AvcizzVaLJf-7G1N6PM2f9FqDV1rQwFeZQ808jQdTJGYoF0qT_FS0ONgI1URSNE7vLvmSFVjoJPsAeMQBf1iOnM4qGH2c4kQcCowY-N3pc"}
   req.Body = io.NopCloser(body)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
