package arkose

import (
   "encoding/base64"
   "encoding/json"
)

func get_bda() (string, error) {
   data, err := json.Marshal(map[string][]byte{
      "ct": make([]byte, 3504),
   })
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(data), nil
}
