package main

import (
   "154.pages.dev/protobuf"
   "crypto/sha1"
   "encoding/binary"
   "math/bits"
)

const response = "\x1a\x18\n\x16\n\x14\n\x10\xeb\xec\x1fy\xc5\x1dSEA\x1a\xea\t\xed\xc9\x16K\x10\n*\xe0\x01\x03\x00D\xab1[P\xbb/\x02s\x17ط!7C\xf5J\x86\xe4\xed\xce~\x169f\xc33\xa0T㰟Sw\xea\xa9W\x9d\xdaĻ\xb5\x1b\xf6\x92\x87w\xd4ں\xc4 թ\xb0\xa0\x86\xbb>/\xfb\xad˕\xb3T\xad\x92࢈\a\x83\xdb\xd1Y\x89)\x1b\xbbn\xbaM\x96O8'vvο\x8f\xed\x83\tT\xe2\x96Z\xf6vd\x80\x8aY\x901\xba\\\xa0X\x9b>\x82\x15\xc6\xf9\xb7\xee\xebJ\xd8\x0e\x02\x05\xe1\x9a[ƆB\xddJ\xa5\xe2\x18:\xd1M\xcck\r\x12\xad\xfc\xb5|\x17\xa5\x85Un\x82aC\x9c\xcc\xe6/At(s\"\xaeTB\xb3\x89Z\x8e\xdb\x00?\x1c\xc1(f\xef\xd73Jd\\h/\xdb\xee\xd0\xef\x12v\xa0\xa6\x92է^\xcd-\xc3}\xb5\xa3\xa0Ο\x18\xa0\xb7\x0e\x87\xf2\\{\xf6I\xeb۩)\xb9"

func hashcash(login_context, message []byte) []byte {
   rand := func() uint64 {
      sum := sha1.Sum(login_context)
      return binary.BigEndian.Uint64(sum[len(sum)-8:])
   }()
   var counter uint64
   for {
      var b []byte
      b = binary.BigEndian.AppendUint64(b, rand)
      b = binary.BigEndian.AppendUint64(b, counter)
      zero_bits := func() int {
         begin := string(login_context)
         sum := sha1.Sum(append(message, b...))
         if string(login_context) != begin {
            panic(1)
         }
         x := binary.BigEndian.Uint16(sum[sha1.Size-2:])
         return bits.TrailingZeros16(x)
      }()
      if zero_bits >= 10 {
         return b
      }
      rand++
      counter++
   }
}

func main() {
   var curse protobuf.Message
   err := curse.Consume([]byte(response))
   if err != nil {
      panic(err)
   }
   login_context, ok := curse.GetBytes(5)
   if !ok {
      panic("login_context")
   }
   prefix, ok := func() ([]byte, bool) {
      m, _ := curse.Get(3)
      m, _ = m.Get(1)
      m, _ = m.Get(1)
      return m.GetBytes(1)
   }()
   if !ok {
      panic("prefix")
   }
   hashcash(login_context, prefix)
}

