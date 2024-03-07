package android

import (
   "154.pages.dev/protobuf"
   "fmt"
   "os"
   "testing"
)

const canonical_uri = "spotify:track:1oaaSrDJimABpOdCEbw2DJ"

func TestMetadata(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var login LoginOk
   login.Data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   meta, err := login.metadata(canonical_uri)
   if err != nil {
      t.Fatal(err)
   }
   for _, file := range meta.file() {
      if file.OGG_VORBIS_320() {
         fmt.Println(file.file_id())
      }
   }
}

// open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ
var _ = protobuf.Message{
   protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\xc8\x01\x10\x80\xa3\x05\x18\x80\xb4\xbc\x02")},
      protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(200)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(86400)},
         protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(5184000)},
      }},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(5)},
      // 2.3
      protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\xc8\x01\x12 972a5901e968a4832539d94ffba2a453 \x80\xa3\x05(\x80\xb4\xbc\x02")},
         protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(200)},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("972a5901e968a4832539d94ffba2a453")},
            protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(86400)},
            protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(5184000)},
         }},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("spotify:track:1oaaSrDJimABpOdCEbw2DJ")},
         // 2.3.3
         protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("type.googleapis.com/spotify.extendedmetadata.audiofiles.AudioFilesExtensionResponse")},
            // 2.3.3.2
            protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xf6\x82ҩ]\x0e\x14\xee\xefO@\xb6\x0f\xdd\xdeV\xbcg!\xc7")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("j_\x12\xfaQ\xf2\xc1ℯpz\x99\xf3ʆ\x96\xf7\xf6/")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(2)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x06J\xff\v\xf7'8[\xcby\xe0\xa3֕֔\xd8\xc0,\xfe")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(8)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xfe\xb3\xd0e\x0e\xc8|\xcf\xdf\x18\xb6\x89i\xb8_\xa4\xbb\x10\xfb\xb4")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(16)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("{T\x91\xd2rO\xf2HG\x16\x19.A\xe8\x06v\x824\xf68")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(18)},
                  }},
               }},
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xbd\x03z\xb3\x18a\x053\xfbk\xc5 \xe2:!\x83\x18\xf6C\xa8")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(19)},
                  }},
               }},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("\rqn6\xc1\x15ۻ\x06?")},
               protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 5, Value: protobuf.Fixed32(3241569905)},
                  protobuf.Field{Number: 2, Type: 5, Value: protobuf.Fixed32(1057405915)},
               }},
               protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("\r\xdf33\xc1\x15\a5\x8c?")},
               protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 5, Value: protobuf.Fixed32(3241358303)},
                  protobuf.Field{Number: 2, Type: 5, Value: protobuf.Fixed32(1066153223)},
               }},
               protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("\x18͊\xfdV\x1eC|\xa2HID/2A\x13")},
            }},
         }},
      }},
   }},
   protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\xc8\x01\x10\xf4\xc3\x04\x18\x80\x9a\x9e\x01")},
      protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(200)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(74228)},
         protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(2592000)},
      }},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(10)},
      // 2.3
      protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\xc8\x01\x12\baecfe9d9 \xf4\xc3\x04(\x80\x9a\x9e\x01")},
         protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(200)},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("aecfe9d9")},
            protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(74228)},
            protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(2592000)},
         }},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("spotify:track:1oaaSrDJimABpOdCEbw2DJ")},
         // 2.3.3
         protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("type.googleapis.com/spotify.metadata.Track")},
            // 2.3.3.2
            protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("-\xa9\xa1\x102fD\x13\xb2M\xe1\x81\xc54\xf1W")},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("No Ordinary Love")},
               protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\\\xe0\x0f<\x8a[D\xb0\xbaӀ\xf8\xdbB\xf4:")},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Love Deluxe")},
                  protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("\n\x10\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY\x12\x04Sade")},
                  protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY")},
                     protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Sade")},
                  }},
                  protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("Epic")},
                  protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("\b\x90\x1f\x10\x14\x184")},
                  protobuf.Field{Number: 6, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3984)},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(20)},
                     protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(52)},
                  }},
                  protobuf.Field{Number: 17, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                        protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xabgam\x00\x00\x1e\x02\xeee\xbb\xd5O\x99;_\x01\xd5\xc5\x11")},
                        protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
                        protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(600)},
                        protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(600)},
                     }},
                     protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                        protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xabgam\x00\x00HQ\xeee\xbb\xd5O\x99;_\x01\xd5\xc5\x11")},
                        protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
                        protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(128)},
                        protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(128)},
                     }},
                     protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                        protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xabgam\x00\x00\xb2s\xeee\xbb\xd5O\x99;_\x01\xd5\xc5\x11")},
                        protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(2)},
                        protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1280)},
                        protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(1280)},
                     }},
                  }},
                  protobuf.Field{Number: 25, Type: 2, Value: protobuf.Bytes("\n\x10\xaa\x94hS\x9d\x19@\f\xa4\xc3/\xdc\x15\x99&\xa9")},
                  protobuf.Field{Number: 25, Type: -2, Value: protobuf.Message{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xaa\x94hS\x9d\x19@\f\xa4\xc3/\xdc\x15\x99&\xa9")},
                  }},
               }},
               protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("\n\x10\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY\x12\x04Sade")},
               protobuf.Field{Number: 4, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY")},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Sade")},
               }},
               protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(2)},
               protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(2)},
               protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(880932)},
               protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(142)},
               protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("\n\x04isrc\x12\fGBBBM0002118")},
               protobuf.Field{Number: 10, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("isrc")},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("GBBBM0002118")},
               }},
               // 2.3.3.2.12
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(2)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("j_\x12\xfaQ\xf2\xc1ℯpz\x99\xf3ʆ\x96\xf7\xf6/")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xf6\x82ҩ]\x0e\x14\xee\xefO@\xb6\x0f\xdd\xdeV\xbcg!\xc7")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x17=\xf6\x8bP\x97\xd4\a\xde\xff\x84\x1fAWl\x81\xc3\x0e\x13B")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(13)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("l\x9ex0\xa3\xf4\xbcw\xd08K\xc3\xf5\xd2}\xf5 ^Q0")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(11)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x93\xd1v\\\xfd\xc5\xe8JA`\x00Qs\x1az\xffʢ\xe9\x12")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(12)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9$\x82\xfe\x9b\xedsr\xd1e}~\"\xf3+y)\x02\xf3\xbd")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(10)},
               }},
               protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x06J\xff\v\xf7'8[\xcby\xe0\xa3֕֔\xd8\xc0,\xfe")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(8)},
               }},
               protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("\n\x14\xaf\xd4T\x13\xf09\x88d\xf2\xee\xc0\x17M\xa79\x1bh\x0f\xcf\x19\x10\x06")},
               protobuf.Field{Number: 15, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xaf\xd4T\x13\xf09\x88d\xf2\xee\xc0\x17M\xa79\x1bh\x0f\xcf\x19")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(6)},
               }},
               protobuf.Field{Number: 17, Type: 0, Value: protobuf.Varint(1595430884)},
               protobuf.Field{Number: 18, Type: 0, Value: protobuf.Varint(1)},
               protobuf.Field{Number: 21, Type: 2, Value: protobuf.Bytes("\n\x10\xaa\x94hS\x9d\x19@\f\xa4\xc3/\xdc\x15\x99&\xa9")},
               protobuf.Field{Number: 21, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xaa\x94hS\x9d\x19@\f\xa4\xc3/\xdc\x15\x99&\xa9")},
               }},
               protobuf.Field{Number: 22, Type: 2, Value: protobuf.Bytes("en")},
               protobuf.Field{Number: 24, Type: 2, Value: protobuf.Bytes("\n\x10\x18͊\xfdV\x1eC|\xa2HID/2A\x13")},
               protobuf.Field{Number: 24, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x18͊\xfdV\x1eC|\xa2HID/2A\x13")},
               }},
               protobuf.Field{Number: 27, Type: 2, Value: protobuf.Bytes("No Ordinary Love")},
               protobuf.Field{Number: 32, Type: 2, Value: protobuf.Bytes("\n\x10\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY\x12\x04Sade\x18\x01")},
               protobuf.Field{Number: 32, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x87p\xea?l\xe4B\xe4\x9a'/V\xc5\xff\x8dY")},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Sade")},
                  protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
               }},
               protobuf.Field{Number: 36, Type: 2, Value: protobuf.Bytes("spotify:track:1oaaSrDJimABpOdCEbw2DJ")},
            }},
         }},
      }},
   }},
}
