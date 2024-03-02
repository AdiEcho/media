# android

https://play.google.com/store/apps/details?id=com.spotify.music

~~~
> play -a com.spotify.music
files: APK APK APK APK
name: Spotify: Music and Podcasts
offered by: Spotify AB
price: 0 USD
requires: 5.0 and up
size: 72.02 megabyte
updated on: Feb 26, 2024
version code: 111414784
version name: 8.9.18.512
~~~

Create Android 6 device. Install user certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ
~~~

- https://github.com/Ahmeth4n/javatify/issues/3
- https://github.com/glomatico/spotify-aac-downloader/issues/21

## 1

~~~
POST https://login5.spotify.com/v3/login HTTP/2.0
cache-control: no-cache, no-store, max-age=0
client-token: AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q==
user-agent: Spotify/8.9.18.512 Android/23 (Android SDK built for x86)
content-type: application/x-protobuf
content-length: 127
accept-encoding: gzip

protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("58cebdd226ac462a")},
   }},
   protobuf.Field{Number: 101, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("EMAIL ADDRESS")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("PASSWORD")},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")},
   }},
}
~~~

## 2

~~~
POST https://login5.spotify.com/v3/login HTTP/2.0
cache-control: no-cache, no-store, max-age=0
client-token: AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q==
user-agent: Spotify/8.9.18.512 Android/23 (Android SDK built for x86)
content-type: application/x-protobuf
content-length: 440
accept-encoding: gzip

protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("58cebdd226ac462a")},
   }},
   protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("\x03\x00\xa4\x12\xf2\xcb9\x1e\xc9\x0e\v z\xfe/\xa9\xf9\x9a=\xa2\x1a\f\xb6\xab\x9e=\xef\xa6q<\xa2恨a\x05\xa6kǪ\xa0\xcc\xd99]+ثHX\xca\xe8h\x85,\x02\xe4I\x05i\xb5\xc9/\xea#yզ\x1e܈UG\x18\xa5\\\xe4\xf2\xde\xea.\xfd\xf3\x1a\xa7\xed\x06N\xea\xb8\x026|\x17\x06\xae<)_R\x1e\xa0\xbebfG\x94\xf5i\xb6\x91\x00\x88ر\x90G{\xf4\xe0d\xb6\x11\x82\x16\xb5\xc0\n\x81HZ\xd6g\xd3K\x96\xb0:\x0f\x8eH^\xba\n\xc7.3UJ\xb6\xc88\x02\xd1.\xf5\x8b\x94ќb\x89\a\xd3DI:Fur\x89\xf6\xa4\xddtL\x1a\xfbso\xe6\x11\xc6μo\xb1\xb7\x99\x8a\x1b\xae\x10[\xf7\xb7\x19=\xacU\xb0\x19\x01\x1b\x05&\xbaZ\x02r\xa6\xab\xff\xea\x1b\x19\xdb\ra\xd8R\xb9'{\x12*]\xe2\xa7(\\\x06x#\x8a@}\xe0\x98_\x03-e\xbe\xec\xbc:\xf1\xc4\x12\x92\xe5[\xe7\xacd\xd6\x10H@춲\xe8\xf5L\xf5\xf4\xeeC\xd1\x02\xa6\xbf\x8bc\xbf\x8b\x8c\xe6\xda")},
   protobuf.Field{Number: 3, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xa84\xe7\xf9\x8aֵ\xbb\x00\x00\x00\x00\x00\x00\x006")},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Message{
               protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(14400)},
            }},
         }},
      }},
   }},
   protobuf.Field{Number: 101, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("EMAIL ADDRESS")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("PASSWORD")},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")},
   }},
}
~~~

## 3

~~~
POST https://login5.spotify.com/v3/login HTTP/2.0
cache-control: no-cache, no-store, max-age=0
client-token: AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q==
user-agent: Spotify/8.9.18.512 Android/23 (Android SDK built for x86)
content-type: application/x-protobuf
content-length: 306
accept-encoding: gzip

protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("58cebdd226ac462a")},
   }},
   protobuf.Field{Number: 100, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("tzjngbdihh9uk2wd4w5016f21")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("AQC16FhLEAnXz3AmrmM8VsWA85EqrwpOiS3HUFM2R-escWtteW_eiNTawRqy18tS9CqHdkIalCp31A0S_gy2sqOLhxGTcloaYX4wN8zNbZUuf6kyXQ5CIoHScC9sGykWauDtBSwy9v5gIG7GfofLjwjW5BQOH5xgm5-ywHiHS-G1DRBjR6Asud75ThdbDpeg8oiAdxjmmYAQ7oAuyyKcgg4pkQ")},
   }},
}
~~~

## 4

~~~
POST https://guc3-spclient.spotify.com/extended-metadata/v0/extended-metadata HTTP/2.0
authorization: Bearer BQCkkXlvEzT-iTS4rlLwOnnAzmyxcuz7yI19Joys5qvLZxwB0XCm8bea7ikhOoioxprBD8jGa0gqnBq1wSIUXbi6Yt9iB-uZYRv5Ogwu6Ccq_59CfHlB6x8dzHeFxuvGVvQCdCQ7RMZfZ3aucXPXNNMnt_Pm8hp1dNLGeb92CKWSIf7f6UziCrBVTfJap2f0j_uHbjZamT3DKve-xhj0ViqHA30WPY6EZFhs6pzAAPmBp4hjNmheQvwMU9GWhKjvxVlJvbRV994gWlg01krDWis4CC7CsEVKOVRYBCIkg3H5vl5ymO2dNFuVvFQSCmUuWYPqx350UmulKbObUvzz
client-token: AADfPTq9lGRU/AhlIKp0BygtbRyID6gkDzjuL7PJcNUvflzFJkXDNfM8KGYi+tMCdTPwDbyiP2EYFydVmcUkkP+R2l6s2+KuV6weSWFi8QyAyXA5MCYyc+p5yNFAxBvaah0tYmoL82LR3z0m/yrXgj1hlEwL4h30BidK6bnF8GK3TAv3aDQHBR09AuSSSOqYtHTRFg2XSl2TI0P86cGgN/w94Ca1j5u9/e2YcW2irkx9woFnvBgKvgCRbLQdWr5Trc1K80FZSqEIsWVJG70pICyfLYmTcciRaaBtGzwwLY8Mi1KqsSJ8Y5Y+zqTP671NI/gotDB52yz/GQJJ+Q==
user-agent: Spotify/8.9.18.512 Android/23 (Android SDK built for x86)
accept: application/protobuf
spotify-app-version: 8.9.18.512
app-platform: Android
accept-language: en-US
content-type: application/protobuf
content-length: 74
accept-encoding: gzip
~~~

## PouleR/spotify-login

OK start here:

~~~php
public function login(string $username, string $password): ?AccessToken
~~~

https://github.com/PouleR/spotify-login/blob/main/src/SpotifyLogin.php

then:

~~~php
$challengeSolutions = $this->solveHashCashChallenge($loginResponse);
~~~

then get field 3.1:

~~~php
$hashCashChallenge = $loginResponse->getChallenges()->getChallenges()[0];
~~~

example `$loginResponse`:

~~~go
protobuf.Message{
   protobuf.Field{Number: 3, Type: 2, Value: protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("6\x90l\b\xb2\xf27\xdal\xff\x9cV=\x00\xf7\xe0")},
            protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(10)},
         }},
      }},
   }},
   protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("\x03\x00\x11\x1dj<\xa6,\xf7\x8d\x0eٞA$\x1dW\x019\x06\xe34/8<\x8fg\xc9Q\f+\xb5Q\xe7^e\n]E\xdbdp\x962\xd5\x17\xd9\xe2\x17\xbf\x1a\xa45E\xa4\x9d\x95ʴ\xc8\xfcw\xac\xc1\xbe\xb8\x1b6#RzM<wha,W\xdf\xfauE\x7f\x83\xe72\xec\x9f_!\xdb\xfe\x05\xb5\xd1.\xa7\x9a\xf3\x01\xf4s\x7f+\x9a\x93\xc5kK/ X\xc2b\xb7\x8e\xa7Q|\xd4\xee\xf1\xf6r\x85\x87y\xf4\xe3\x15\xe6\a+\fV{:\xb1l\xa0&҂\xf7ӧ\xaa\x1f\x7f\xa0O\xb1{\xdb\x1c\xa7\xbd\x16\xa21\xca\rQQ\x15Y ݅\xf7\"\xa4\x8f\x04\xe3\x97j\xad\xa3\x89\xac֓\xa9\xc8\x05\xc6v\xbaz\x03$\x7f\xbd\xa7\x01p\xba\xc5v4L\xbd\x16\xa4\xb8\xda\xd7\xe9\xbb\xf1IT>\xc9\"\xc6-y\xf1\xae\xd3\x165")},
}
~~~

then get field 1:

~~~php
$hashCash = $hashCashChallenge->getHashcash();
~~~

then get field 5:

~~~php
$loginResponse->getLoginContext(),
~~~

then get field 1:

~~~php
$hashCash->getPrefix(),
~~~

then get field 2:

~~~php
$hashCash->getLength()
~~~

then:

~~~php
$seed = substr(sha1($loginContext), -16);
$seed = HexUtils::hex2ByteArray($seed);
$start = hrtime(false);
$solved = $this->solveHashCash($hashCashPrefix, $hashCashLength, $seed);
$end = hrtime(false);
$solved->setDuration($end[1] - $start[1]);
return $solved;
~~~
