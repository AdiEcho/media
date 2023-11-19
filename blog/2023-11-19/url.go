package url

import "strings"

func path(s string) string {
   _, after, found := strings.Cut(s, "://")
   if found {
      s = after // remove scheme
   }
   i := strings.IndexByte(s, '/')
   if i >= 1 {
      s = s[i:] // remove host
   }
   return s
}
