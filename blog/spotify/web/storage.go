package web

import (
   "net/http"
   "net/url"
)

func storage(file_id string) (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/storage-resolve/v2/files/audio/interactive/10/" + file_id
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer BQA9WA-ZkKdTitbqkrs9XWYCkfJwDsCR80eW1LIVD6vEnua0V2g60hLaWb1d-ycakaRskAEboHQ2kS4xsh00BoGk9P-t4Ji_EBwiiWzl2Q18_WAl-USydjKNQqQNC1jd87m2ZvmmbDjJbJWZ05HnnTRQVgzmCUJyHXeS4Slk1yx0K_p-LoWnRw4Stvio0xHHLYJTU_4wxsNmpjdg5B-hnYGRckhJh566YLS4Zc1uBWHWUgc8cRIHNeeILIlm9cE84CcZ37ZhLhLSHiO9GKu69H_3LOkufVUMt5-omiiDlrF1xr-u_Wpv5mDuM7MeI16Gpnxr8vWDM-cAhxEvXDgZ"}
   val := make(url.Values)
   val["alt"] = []string{"json"}
   req.URL.RawQuery = val.Encode()
   return http.DefaultClient.Do(&req)
}
