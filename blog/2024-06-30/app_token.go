package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
)

/*
|
2b2caa6373626591
\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f

#\xca_\xc1\x0e\xae\x99\x1e\xd3x\xad\xac
|
ly\xa00d40b8f1ec1746
\x02\x02
*/
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
   data = data[1 + 1 + aes.BlockSize:]
   cipher.NewCBCDecrypter(block, iv[:]).CryptBlocks(data, data)
   return data, nil
}

type AppToken string

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
