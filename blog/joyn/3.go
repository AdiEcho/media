package joyn

import (
   "bytes"
   "crypto/sha1"
   "encoding/hex"
   "encoding/json"
   "net/http"
)

const signature_key = "5C7838365C7864665C786638265C783064595C783935245C7865395C7838323F5C7866333D3B5C78386635"

func (a anonymous) entitlement(content_id string) (*entitlement, error) {
   body, err := json.Marshal(map[string]string{"content_id": content_id})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://entitlement.p7s1.io/api/user/entitlement-token",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + a.Access_Token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   title := new(entitlement)
   err = json.NewDecoder(res.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

type entitlement struct {
   Entitlement_Token string
}

func (e entitlement) signature(text []byte) string {
   text = append(text, ',')
   text = append(text, e.Entitlement_Token...)
   text = hex.AppendEncode(text, []byte(signature_key))
   sum := sha1.Sum(text)
   return hex.EncodeToString(sum[:])
}
