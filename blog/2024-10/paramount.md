# paramount

first:

~~~go
signed_request, err := c.sign_request()
~~~

then:

~~~go
signed.AddBytes(2, c.license_request)
~~~

then:

~~~go
c.license_request = protobuf.Message{
   1: {protobuf.Bytes(client_id)},
   2: {protobuf.Message{ // content_id
      1: {protobuf.Message{ // widevine_pssh_data
         1: {protobuf.Bytes(pssh)},
      }},
   }},
}.Marshal()
~~~

then:

~~~
[message]    2.2.1.1
[bytes]      2.2.1.1.2     b'\x883\xfc\xa4\xae{Nc\xbd\xa1\x81\xb8\xae<\x03\xe2'
[bytes]      2.2.1.1.2     b'\x80\t\xe9\xdf`\xa6OK\x89\x9c\xdd\x91\xbf\xf89\x85'
[bytes]      2.2.1.1.2     b'\x00}\xf5W\xd2\xd3E\x1d\xbd\xea\x88\xa9Y\rs#'
[bytes]      2.2.1.1.2     b"\xe4\xc3\x9c\x95{\x8fF\x9e\xaf{'Q\x9f\x91\\\xa8"
[uint32]     2.2.1.1.9     1667591779
~~~

then:

~~~
8833fca4ae7b4e63bda181b8ae3c03e2
8009e9df60a64f4b899cdd91bff83985
007df557d2d3451dbdea88a9590d7323
e4c39c957b8f469eaf7b27519f915ca8
~~~

so we are requesting multiple key IDs? which one is the right one? this one:

~~~go
s.key_id = sinf.Schi.Tenc.DefaultKid[:]
~~~

which is:

~~~
defaultKID: 8009e9df-60a6-4f4b-899c-dd91bff83985
~~~

which does match. now lets look at the response:

~~~go
license, ok := message.Get(2)()
~~~

then:

~~~go
containers := license.Get(3) // KeyContainer key
~~~

then:

~~~go
{
   protobuf.Message{
      2: {protobuf.Bytes("\xeeb\x06\xa5\xd7'\x8d\x9eG\xa0\x01\xf1Lr&t")},
      3: {protobuf.Bytes("s%\xd4|\x89wLH\xa2\x89\xb6v\xc8\xe7\x1aD#\x99@\x84ȲFwӞ\x15\xa3\xda\xd1\xeaȑ\x8c\xe3:\xb4\xc0K\xc1x\xe6֏\xde\f\xd7O")},
      4: {protobuf.Varint(1)},
   },
   protobuf.Message{
      1: {protobuf.Bytes("\x00}\xf5W\xd2\xd3E\x1d\xbdꈩY\rs#")},
      2: {protobuf.Bytes("\x04W\x0e\xe5\xacsp\xa5\xe5\xb2u\xca\xe3\xc1\x1a\xcd")},
      3: {protobuf.Bytes("he\x86\xa9\xd2\xe5%\xbeߔ\xe0֘\xaccNm\x9es؛S%\xbf\xd6t\x85\x17\xc5\xd0\xfd\xc8")},
      4: {protobuf.Varint(2)},
      5: {protobuf.Varint(1)},
      6: {protobuf.Unknown{
         protobuf.Bytes("\b\x00\x10*\x18\x00 \x00"),
         protobuf.Message{
            1: {protobuf.Varint(0)},
            2: {protobuf.Varint(42)},
            3: {protobuf.Varint(0)},
            4: {protobuf.Varint(0)},
         },
      }},
      7: {protobuf.Unknown{
         protobuf.Bytes("\b\x00\x10*\x18\x00 \x00"),
         protobuf.Message{
            1: {protobuf.Varint(0)},
            2: {protobuf.Varint(42)},
            3: {protobuf.Varint(0)},
            4: {protobuf.Varint(0)},
         },
      }},
      12: {protobuf.Bytes("SD")},
   },
   protobuf.Message{
      1: {protobuf.Bytes("\x883\xfc\xa4\xae{Nc\xbd\xa1\x81\xb8\xae<\x03\xe2")},
      2: {protobuf.Bytes("\xd0\xc9\x1a\xf4\x00\x91\xc8\xe5\xb5\xc8?g\xfaVm\xeb")},
      3: {protobuf.Bytes("Gw\xb7\\D\xc7\xf9\xf7<ХŠ\x8f\xd0\x15ӌ|\x0e\xe0\u0085\x17\x80\x95\xa3J0[\xb4\xd0")},
      4: {protobuf.Varint(2)},
      5: {protobuf.Varint(1)},
      6: {protobuf.Unknown{
         protobuf.Bytes("\b\x00\x10*\x18\x00 \x00"),
         protobuf.Message{
            1: {protobuf.Varint(0)},
            2: {protobuf.Varint(42)},
            3: {protobuf.Varint(0)},
            4: {protobuf.Varint(0)},
         },
      }},
      7: {protobuf.Unknown{
         protobuf.Bytes("\b\x00\x10*\x18\x00 \x00"),
         protobuf.Message{
            1: {protobuf.Varint(0)},
            2: {protobuf.Varint(42)},
            3: {protobuf.Varint(0)},
            4: {protobuf.Varint(0)},
         },
      }},
      12: {protobuf.Bytes("AUDIO")},
   },
}
~~~
