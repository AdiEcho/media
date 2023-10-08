package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "strings"
   "time"
)

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      oneApp: $oneApp
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            airDate
            episodeNumber
            mpxAccountId
            mpxGuid
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphQL_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

type On_Demand struct {
   // this is only valid for one minute
   Manifest_Path string `json:"manifestPath"`
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      One_App bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      User_ID string `json:"userId"`
   } `json:"variables"`
}

type video_request struct {
   Device string `json:"device"`
   Device_ID string `json:"deviceId"`
   External_Advertiser_ID string `json:"externalAdvertiserId"`
   MPX struct {
      Account_ID string `json:"accountId"`
   } `json:"mpx"`
}

var secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

func authorization(b []byte) string {
   now := time.Now().UnixMilli()
   hash := hmac.New(sha256.New, secret_key)
   fmt.Fprint(hash, now)
   b = append(b, "NBC-Security key=android_nbcuniversal,version=2.4"...)
   b = append(b, ",time="...)
   b = fmt.Append(b, now)
   b = append(b, ",hash="...)
   b = fmt.Appendf(b, "%x", hash.Sum(nil))
   return string(b)
}
