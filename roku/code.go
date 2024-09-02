package roku

import (
   "encoding/json"
   "strings"
)

func (a AccountCode) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(a.v.Code)
   return b.String()
}

type AccountCode struct {
   Data []byte
   v *struct {
      Code string
   }
}

func (a *AccountCode) Unmarshal() error {
   a.v = pointer(a.v)
   return json.Unmarshal(a.Data, a.v)
}
