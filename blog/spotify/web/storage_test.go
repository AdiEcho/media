package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "encoding/json"
   "fmt"
   "os"
   "slices"
   "testing"
   "time"
)

/*
MP4_256_DUAL
https://audio4-fa.scdn.co/audio/173df68b5097d407deff841f41576c81c30e1342?1710211902_B30PYrTn8rEZso3x_oiLx3RTnUyp5N60K6klLAr_mIc=
pssh AAAAU3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADMIARIQFz32i1CX1Afe/4QfQVdsgRoHc3BvdGlmeSIUFz32i1CX1Afe/4QfQVdsgcMOE0I=
encrypted

MP4_256
https://audio4-fa.scdn.co/audio/6c9e7830a3f4bc77d0384bc3f5d27df5205e5130?1710211903_ME0xucThfSapIBK1othiOcW2sI4tRgo64dG0tZomFbo=
clear
pssh AAAAU3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADMIARIQbJ54MKP0vHfQOEvD9dJ99RoHc3BvdGlmeSIUbJ54MKP0vHfQOEvD9dJ99SBeUTA=

MP4_128_DUAL
https://audio4-fa.scdn.co/audio/93d1765cfdc5e84a41600051731a7affcaa2e912?1710211904_F9oWDP2EkFJ26_qho78eZE9CAHFKqdL-EAwtVT6rnTA=
encrypted
pssh AAAAU3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADMIARIQk9F2XP3F6EpBYABRcxp6/xoHc3BvdGlmeSIUk9F2XP3F6EpBYABRcxp6/8qi6RI=

MP4_128
https://audio4-fa.scdn.co/audio/392482fe9bed7372d1657d7e22f32b792902f3bd?1710211905_5ZBcdOvXZCMPQ-nNxcD630OLwNkLfDDTmtMWtHQHqHo=
clear
pssh AAAAU3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADMIARIQOSSC/pvtc3LRZX1+IvMreRoHc3BvdGlmeSIUOSSC/pvtc3LRZX1+IvMreSkC870=

OGG_VORBIS_320
https://audio4-fa.scdn.co/audio/98b53c239db400b0a016145700de431f68d28f54?1710211899_20vgNGuVlb5sbsnhQDH9mpL0qOkuSTSO5_o77EL-35c=
encrypted

OGG_VORBIS_160
https://audio4-fa.scdn.co/audio/6a5f12fa51f2c1e284af707a99f3ca8696f7f62f?1710211900_-rvlzvWBXOezPzjNZ6FeiyU-J_f-H24uHaQustNyWIU=
encrypted

OGG_VORBIS_96
https://audio4-fa.scdn.co/audio/f682d2a95d0e14eeef4f40b60fddde56bc6721c7?1710211901_QonRJA6_F_2JvlyPor2NAUCq5xoaIbgFQZf0Zi3Twlg=
encrypted

AAC_24
https://audio4-fa.scdn.co/audio/064aff0bf727385bcb79e0a3d695d694d8c02cfe?1710211906_h0PYn4oGB8nKzD2b35v_Klnhn4cV-9j5tN8eCyRz2UA=
encrypted
*/
func TestStorage(t *testing.T) {
   text, err := os.ReadFile("metadata.json")
   if err != nil {
      t.Fatal(err)
   }
   var meta metadata
   if err := json.Unmarshal(text, &meta); err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var login android.LoginOk
   login.Data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   for _, file := range meta.File {
      fmt.Println(file.Format)
      var storage storage_resolve
      if err := storage.New(login, file.File_ID); err != nil {
         t.Fatal(err)
      }
      slices.SortFunc(storage.CDNURL, func(a, b string) int {
         return len(a) - len(b)
      })
      fmt.Println(storage.CDNURL[0])
      time.Sleep(time.Second)
   }
}
