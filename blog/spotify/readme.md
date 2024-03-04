# spotify

https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ

## web

~~~
HTTP/1.1 400 Bad Request
Connection: close

not a WebSocket handshake request: missing upgrade
~~~

also login is protected by `recaptchaToken`

## android

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
~~~

https://github.com/PouleR/spotify-login/blob/main/src/ChallengeSolver.php

~~~php
$suffix = array_merge($random, [0, 0, 0, 0, 0, 0, 0, 0]);
$i = 0;
while (true) {
   $input = (bin2hex($prefix) . HexUtils::byteArray2Hex($suffix));
   $digest = sha1(hex2bin($input));
   if ($this->checkTenTrailingBits(HexUtils::hex2ByteArray($digest))) {
       return new SolvedLoginChallenge($suffix, $i);
   }
   $this->incrementCtr($suffix, count($suffix) - 1);
   $this->incrementCtr($suffix, 7);
   $i++;
}
~~~

then:

~~~php
private function checkTenTrailingBits(array $trailingData) : bool {
  if ($trailingData[count($trailingData) - 1] != 0) {
      return false;
  }
  return $this->countTrailingZero($trailingData[count($trailingData) - 2]) >= 2;
}
~~~

then:

~~~php
private function countTrailingZero(int $x): int {
  if ($x === 0) {
      return 32;
  }
  $count = 0;
  while (($x & 1) == 0) {
      $x = $x >> 1;
      $count++;
  }
  return $count;
}
~~~

then:

~~~php
private function incrementCtr(array &$ctr, int $index): void {
  $ctr[$index]++;
  if ($ctr[$index] > 0xFF && $index != 0) {
      $ctr[$index] = 0;
      $this->incrementCtr($ctr, $index - 1);
  }
}
~~~

- https://github.com/CrazyHoneyBadger/pow/issues/1
- https://github.com/LeastAuthority/hashcash/issues/1
- https://github.com/PaulSnow/hashproof/issues/1
- https://github.com/PoW-HC/hashcash/issues/2
- https://github.com/agfy/hashcash/issues/1
- https://github.com/alextanhongpin/go-hashcash/issues/1
- https://github.com/catalinc/hashcash/issues/4
- https://github.com/cleanunicorn/hashcash/issues/1
- https://github.com/denismitr/hashcache/issues/1
- https://github.com/kirvader/wow-using-pow/issues/1
- https://github.com/laonix/pow-word-of-wisdom/issues/1
- https://github.com/maksadbek/pow/issues/1
- https://github.com/maxim-dzh/word-of-wisdom/issues/1
- https://github.com/prestonTao/hashcash/issues/1
- https://github.com/rezamt/hashcash-go/issues/5
- https://github.com/rodzevich/hashcash/issues/1
- https://github.com/rolandshoemaker/hashcash-go/issues/1
- https://github.com/sstelfox/provingwork/issues/2
- https://github.com/timurkash/hashcash/issues/1
- https://github.com/umahmood/hashcash/issues/1
- https://github.com/vxxvvxxv/hashcash/issues/1

17 imports:

https://pkg.go.dev/github.com/robotomize/powwy/pkg/hashcash
