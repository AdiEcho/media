package paramount

import (
   "strconv"
   "strings"
)

// hard geo block
func (v *VideoItem) Mpd() string {
   b := []byte("https://link.theplatform.com/s/")
   b = append(b, v.CmsAccountId...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, decode_id(v.CmsAccountId), 10)
   b = append(b, '/')
   b = append(b, v.ContentId...)
   b = append(b, "?assetTypes="...)
   b = append(b, v.asset_type()...)
   b = append(b, "&formats=MPEG-DASH"...)
   return string(b)
}

const encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func decode_id(s string) int64 {
   var (
      i = 0
      j = 1
   )
   for _, value := range s {
      i += strings.IndexRune(encoding, value) * j
      j *= len(encoding)
   }
   return int64(i)
}
