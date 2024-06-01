package main

import (
   "154.pages.dev/protobuf"
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   var s struct {
      Widevine2License struct {
         License []byte
      }
   }
   err := json.Unmarshal(b, &s)
   if err != nil {
      return nil, err
   }
   return s.Widevine2License.License, nil
}

type poster struct{}

func (poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "atv-ps.amazon.com"
   u.Path = "/cdp/catalog/GetPlaybackResources"
   u.Scheme = "https"
   val := make(url.Values)
   val["asin"] = []string{"B0CV72X1BL"}
   val["consumptionType"] = []string{"Streaming"}
   val["desiredResources"] = []string{"Widevine2License"}
   val["deviceID"] = []string{"f1114fdb-4541-47c5-acfe-478978a579f4"}
   val["deviceTypeID"] = []string{"AOAGZA014O5RE"}
   val["firmware"] = []string{"1"}
   val["resourceUsage"] = []string{"ImmediateConsumption"}
   val["videoMaterialType"] = []string{"Feature"}
   u.RawQuery = val.Encode()
   return u.String(), true
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   challenge := url.Values{
      "widevine2Challenge": {base64.StdEncoding.EncodeToString(b)},
   }.Encode()
   return []byte(challenge), nil
}

func (poster) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h["Cookie"] = []string{strings.Join([]string{
      "at-main=Atza|IwEBIOLn2rzz6WCjvtwNVR7bB4A2tkXCFsu8EZG83VGuqQ_xJYvphdavRzuhxscV3dR4Jq9F2BoGHQlUv9ccXisCn9RcArGQMRwz65gYPtX4sviCAb60lAtELtfX5v0TiR4Pec6NQw2bHvjlrb_ZVLMoAtl0X1CbAi59hXBqewh0GUbkwqnFkpptYoA-QYrVdeMLUR647pVkhQo9H1e6YyDmnFYwRopG-k17dWWkeRtyWQq_ig",
      "ubid-main=132-2945600-9834461",
      "x-main=\"KBsT7ivk647R9mWeQeTtNsBr?AQSoRjgZbMX8MjwqELlBpjIEk0?yYNrLxVC7bZp\"",
   }, ";")}
   return h, nil
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      panic(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      panic(err)
   }
   pssh := protobuf.Message{
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("F\xfeqp\x8dwM\b\x92\x93R\x968\xe2Ss")},
      // BAD
      //protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("cid:Rv5xcI13TQiSk1KWOOJTcw==,Jjp5sClWRHiky3k7BqxMpA==")},
   }.Encode()
   var module widevine.CDM
   err = module.New(private_key, client_id, pssh)
   if err != nil {
      panic(err)
   }
   key, err := module.Key(poster{}, nil)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%x\n", key)
}
