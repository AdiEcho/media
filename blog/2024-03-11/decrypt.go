package main

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/hex"
   "os"
)

// go.dev/src/crypto/cipher/example_test.go
func main() {
   bytes, err := os.ReadFile("OGG_VORBIS")
   if err != nil {
      panic(err)
   }
   file_id, err := hex.DecodeString("f682d2a95d0e14eeef4f40b60fddde56bc6721c7")
   if err != nil {
      panic(err)
   }
   obfuscatedKey, err := hex.DecodeString("17232673e8d1ab865346d0f14ef79b97")
   if err != nil {
      panic(err)
   }
   if err := decrypt(bytes, file_id, obfuscatedKey); err != nil {
      panic(err)
   }
   os.WriteFile("dec.ogg", bytes, 0666)
}

var aesIV = []byte{
   0x72, 0xE0, 0x67, 0xFB, 0xDD, 0xCB, 0xCF, 0x77,
   0xEB, 0xE8, 0xBC, 0x64, 0x3F, 0x63, 0x0D, 0x93,
}

func decrypt(data, dataId, keyBasis []byte) error {
   key := processKeyBasis(keyBasis, dataId)
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   stream := cipher.NewCTR(block, aesIV)
   stream.XORKeyStream(data, data)
}
