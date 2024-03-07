# web

## problems

login is protected by `recaptchaToken`. also some requests are protected by
websocket:

~~~
GET /?access_token=BQANivsOWMZT8y5kWv6jsp5Owmnh9hMeau7GZMttM2AteQJ1r7eiJwkySRA-CYffnINF2VwZEkS6gtJI9eLG2HkQkwyzJ8GG3B3ZUdEAb2jIFZo4_wpILoI2oa12X-Oz-1QBFeY0l_lePhuutnWZ2Qcqu0Lp8B2ikKwRkLtPyTiX_ua6FBxVFG2Z9G3LUOLCorUswoftw235L4sQrSFpxprREm2H8mvT7IZH50qzthkYUT9RxAZWF_whc1XsSQNaIb0V_tmUG1kpK_8ku8chIKYi8DVHt0rhkuud4yuQc84AIhYNQkxjcKsHvOUp-9ZN5gqeKWtbeHPjUNljEzBo82yX HTTP/1.1
Host: guc3-dealer.spotify.com
Accept-Language: en-US,en;q=0.5
Accept: */*
Cache-Control: no-cache
Connection: keep-alive, Upgrade
Cookie: sp_t=f183547eec914ed51e2e804ea6b7dec1; sp_dc=AQBBwsLqQcfzzLMI8sWz-rA8_fX8mR-7_f3OU5WuW-XmgI3JFcnhmeC0mNw0zn6695OdPN0hleks12_F7WRA61Z5BnN0uJM1dDskG4-uFKzVdmdgY-KoiAzIXCTKEXOdWiBjNlmc2oxSoDj4GYo5KEdre1jEglI; sp_key=0bb58db5-2e26-46c6-a714-460c0167a3f6; sp_landing=https%3A%2F%2Fopen.spotify.com%2Ftrack%2F1oaaSrDJimABpOdCEbw2DJ%3Fsp_cid%3Df183547eec914ed51e2e804ea6b7dec1%26device%3Ddesktop
Origin: https://open.spotify.com
Pragma: no-cache
Sec-Fetch-Dest: websocket
Sec-Fetch-Mode: websocket
Sec-Fetch-Site: same-site
Sec-WebSocket-Extensions: permessage-deflate
Sec-WebSocket-Key: gyxe1xW8seFU4pmmRdvVcw==
Sec-WebSocket-Version: 13
Upgrade: websocket
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~

## 0

https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ

## 1

web click uses websocket here to convert `canonical_uri` to `gid`. use Android
client instead.

## 2

web client uses below to return `file_id` values. previous Android client request
above also returns the values, so use that instead.

~~~
GET https://spclient.wg.spotify.com/metadata/4/track/2da9a11032664413b24de181c534f157?market=from_token HTTP/2.0
accept-language: en
accept: application/json
app-platform: WebPlayer
authorization: Bearer BQC-YIT2Pkn3Nwe1i9142WtH0y1St-oDQ0WI_62yQk7nYZSa0fTx5e4i...
client-token: AABsMhk3m2tbne420W9XI9pFlyLkZbZVGXXMs7+NkVhNot/ZfSss59Td670Svntn...
content-length: 0
origin: https://open.spotify.com
referer: https://open.spotify.com/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
spotify-app-version: 1.2.33.0-unknown
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~

## 3

this is next

~~~
GET https://guc3-spclient.spotify.com/storage-resolve/v2/files/audio/interactive/10/392482fe9bed7372d1657d7e22f32b792902f3bd?version=10000000&product=9&platform=39&alt=json HTTP/2.0
accept-language: en-US,en;q=0.5
accept: */*
authorization: Bearer BQC-YIT2Pkn3Nwe1i9142WtH0y1St-oDQ0WI_62yQk7nYZSa0fTx5e4i...
client-token: AABsMhk3m2tbne420W9XI9pFlyLkZbZVGXXMs7+NkVhNot/ZfSss59Td670Svntn...
content-length: 0
origin: https://open.spotify.com
referer: https://open.spotify.com/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~

## 4

<https://audio-ak-spotify-com.akamaized.net/audio/392482fe9bed7372d1657d7e22f32b792902f3bd?__token__=exp=1709865199~hmac=044acd1c179870511191c49932f31887d0d722de0ed912ce380cb444b1dd1e3d>
