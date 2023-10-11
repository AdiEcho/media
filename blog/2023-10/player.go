package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "bytes"
   "154.pages.dev/protobuf"
   "fmt"
)

const video_ID = "oCjW6gdEDa4"

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/16.49.39(Linux; U; Android 12; en_US; sdk_gphone64_x86_64 Build/SE1B.220616.007) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"CgtoM3lRMWFvN3NjRSiNiJipBjIICgJVUxICGgA6CiDJ46fP1IGBk2U%3D"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   val := make(url.Values)
   val["id"] = []string{video_ID}
   val["key"] = []string{"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"}
   val["t"] = []string{"k1zUJdkoktKC"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Append(nil)))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if strings.Contains(string(data), "This video requires payment to watch") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var req_body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("Google")},
         protobuf.Field{Number: 13, Type: 2,  Value: protobuf.Bytes("sdk_gphone64_x86_64")},
         protobuf.Field{Number: 16, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 17, Type: 2,  Value: protobuf.Bytes("16.49.39")},
         protobuf.Field{Number: 18, Type: 2,  Value: protobuf.Bytes("Android")},
         protobuf.Field{Number: 19, Type: 2,  Value: protobuf.Bytes("12")},
         protobuf.Field{Number: 21, Type: 2,  Value: protobuf.Bytes("en-US")},
         protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("US")},
         protobuf.Field{Number: 25, Type: 2,  Value: protobuf.Bytes("c005a31e9f64d574")},
         protobuf.Field{Number: 37, Type: 0,  Value: protobuf.Varint(432)},
         protobuf.Field{Number: 38, Type: 0,  Value: protobuf.Varint(848)},
         protobuf.Field{Number: 39, Type: 5,  Value: protobuf.Fixed32(1076677837)},
         protobuf.Field{Number: 40, Type: 5,  Value: protobuf.Fixed32(1084856730)},
         protobuf.Field{Number: 41, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 46, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 50, Type: 0,  Value: protobuf.Varint(225014047)},
         protobuf.Field{Number: 52, Type: 0,  Value: protobuf.Varint(4)},
         protobuf.Field{Number: 55, Type: 0,  Value: protobuf.Varint(432)},
         protobuf.Field{Number: 56, Type: 0,  Value: protobuf.Varint(848)},
         protobuf.Field{Number: 61, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 62, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("CJCImKkGEhM0MjMxOTA1NDYzMjE1NTI4NzkwGJCImKkGMjJBT2pGb3gwd2dqUVI5ZHZSU1diTFFsaWtFNFpnMnMtTUFIa3VZSzdLcHZaZDN2Szh6dzoyQU9qRm94MHdnalFSOWR2UlNXYkxRbGlrRTRaZzJzLU1BSGt1WUs3S3B2WmQzdks4endCQENBTVNLZzBYOHQycEFvZ3h5QmItQmZBcV9RYllJeFVSd1BQV0RNcVhBdHYyQm80ZXNrYmVZcDB2NGVfSUN3PT0%3D")},
            protobuf.Field{Number: 5, Type: 2,  Value: protobuf.Bytes("CJCImKkGEhMyNDkwMzcyMDI3Njc3Mzk3MDA2GJCImKkGKJTk_BIoufX8Eijck_0SKI6i_RIoxrL9EiiqtP0SKJ6R_hIomq3-EijIyv4SKN3O_hIoqOH-Eijv5v4SKPzm_hIowe7-Eijt7_4SKJDx_hIyMkFPakZveDB3Z2pRUjlkdlJTV2JMUWxpa0U0Wmcycy1NQUhrdVlLN0twdlpkM3ZLOHp3OjJBT2pGb3gwd2dqUVI5ZHZSU1diTFFsaWtFNFpnMnMtTUFIa3VZSzdLcHZaZDN2Szh6d0IsQ0FNU0d3ME4ySV81RmNvQXFEbkFFUlVLamVMTkRJdnVBZkhnRHJXQUJBPT0%3D")},
         }},
         protobuf.Field{Number: 64, Type: 0,  Value: protobuf.Varint(32)},
         protobuf.Field{Number: 65, Type: 5,  Value: protobuf.Fixed32(1075838976)},
         protobuf.Field{Number: 67, Type: 0,  Value: protobuf.Varint(18446744073709551316)},
         protobuf.Field{Number: 78, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 80, Type: 2,  Value: protobuf.Bytes("America/Chicago")},
         protobuf.Field{Number: 84, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(18213749727395657192)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(14432795415577216842)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(9980950154268185112)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(12581305622008181298)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(14687699447298249089)},
         }},
         protobuf.Field{Number: 86, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
         }},
         protobuf.Field{Number: 89, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 97, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1)},
         }},
         protobuf.Field{Number: 98, Type: 2,  Value: protobuf.Bytes("google")},
      }},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 7, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1696990254032308)},
               protobuf.Field{Number: 2, Type: 5,  Value: protobuf.Fixed32(29529192)},
               protobuf.Field{Number: 3, Type: 5,  Value: protobuf.Fixed32(4278530031)},
            }},
            protobuf.Field{Number: 6, Type: 2,  Value: protobuf.Bytes("external")},
            protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("com.android.shell")},
         }},
      }},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("ms")},
            protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("CoACQO9c4nI5jUekFP1UdvhsX82ya5ho1GydPlqusmu1cI00aEoXOT7OhogGQUELvAh07a5CEUoHKQ-kfUt5OvE-36yL92O2aq2UGtyDliLzjAZ-aetbCvCAm0cUcvP15ErU3cvli0cGDifUMgblD30eBB_H7sGDJvtXokUzvTHGdIbORORAgnPUsdfscUMV0VcYJgOSzUFoQnrJXarJzXUnA4r-L8Pd5kzBXKNyvayQSCj6JKfmo0GiHPr_Yr90jKSEymSra1VDjO0YA-fc70X3CfLr2arbVTNkyOYaFUHNbpsgstv2CWbYlTPXlzGqabm3lFKb06At7WqddSI4_kbEVgqAAki9nZdZ_Nhi7IMbd5L-cr_aoWSb_cO2OsYAMgItj5dVEnr3BejRuzUC3SubhbJFFEn2XGq3MREZAUI97Kv3CUP-SYqlodQkChgZYnvknABR28BnvvLBl6r9AS_UWGH1TZSju4_ssaJBKPmQNL-ZdSLUHrxJ93ZWAtLSbAiSgqe4BkaiRwMdk_xMYzHTejRzThzF4Y12fw8THfBAa6YsK94ge5AYn7ZVSIuTaWsfubG0_ZPNN4qSV_hZdirXa_QhYUwRjRvc-R_PoneFe1ZmNxkMRPl79v7R91OL977cdoBQ0rES4tFC8dHbyNuIodOiHYnILt1DWLJhcLkALD0EWNISEDR6CefLRxFpBJTEmempugU")},
         }},
      }},
   }},
   protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes(video_ID)},
   protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("mvapp-unknown")},
         protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 5, Type: 0,  Value: protobuf.Varint(254)},
         protobuf.Field{Number: 6, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 7, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 8, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 10, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 11, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("sdkv=a.16.49.39&output=xml_vast2")},
         protobuf.Field{Number: 29, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 31, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 7, Type: 2,  Value: protobuf.Bytes("")},
         }},
         protobuf.Field{Number: 37, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 38, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 41, Type: 0,  Value: protobuf.Varint(0)},
      }},
   }},
   protobuf.Field{Number: 5, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 8, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("6AQB")},
   protobuf.Field{Number: 23, Type: 2,  Value: protobuf.Bytes("gzTkP5l3ckmZtVoy")},
   protobuf.Field{Number: 28, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
   }},
}
