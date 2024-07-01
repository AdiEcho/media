package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
)

func decode(s string) ([]byte, error) {
   data, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
      return nil, err
   }
   key, err := hex.DecodeString(secret_key)
   if err != nil {
      return nil, err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   var iv [aes.BlockSize]byte
   data = data[2:]
   cipher.NewCBCDecrypter(block, iv[:]).CryptBlocks(data, data)
   return data, nil
}
