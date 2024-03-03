package spotify

import "crypto/sha1"

func solve_hash_cash_challenge(login_context, prefix []byte) []byte {
   seed := func() []byte {
      b := sha1.Sum(login_context)
      return b[len(b)-8:]
   }()
   suffix := append(seed, 0, 0, 0, 0, 0, 0, 0, 0)
   for {
      input := append(prefix, suffix...)
      digest := sha1.Sum(input)
      if check_ten_trailing_bits(digest[:]) {
         return suffix
      }
      increment_ctr(suffix, len(suffix)-1)
      increment_ctr(suffix, 7)
   }
}

func increment_ctr(ctr []byte, index int) {
   ctr[index]++
   if ctr[index] == 0 {
      if index >= 1 {
         increment_ctr(ctr, index-1)
      }
   }
}

func check_ten_trailing_bits(trailing_data []byte) bool {
   length := len(trailing_data)
   if trailing_data[length-1] >= 1 {
      return false
   }
   return count_trailing_zero(trailing_data[length-2]) >= 2
}

func count_trailing_zero(x byte) byte {
   if x == 0 {
      return 32
   }
   var count byte
   for x&1 == 0 {
      x >>= 1
      count++
   }
   return count
}
