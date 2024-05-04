# draken

- https://drakenfilm.se/film/michael-clayton
- https://justwatch.com/se/leverantör/draken-films

## android

https://play.google.com/store/apps/details?id=com.draken.android

## web

this is it:

<https://media-drakenfilm.partner.maginepro.com/mediaconvert/8149455c-cb3d-4b15-85a8-b95e3d1570b5/hls_video/576p_h264_2M/49.m4s?version=3>

from:

https://media-drakenfilm.partner.maginepro.com/mediaconvert/8149455c-cb3d-4b15-85a8-b95e3d1570b5/AccurateDTManifest.mpd?version=3

from:

~~~
POST /api/playback/v1/preflight/asset/8149455c-cb3d-4b15-85a8-b95e3d1570b5 HTTP/1.1
Host: client-api.magine.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Magine-Play-DeviceId: RB4s4i4ty-z2sFdJJIQuk
Magine-Play-DeviceModel: firefox 111.0 / windows 10
Magine-Play-DevicePlatform: firefox
Magine-Play-DeviceType: web
Magine-Play-DRM: widevine
Magine-Play-Protocol: dashs
Magine-Play-EntitlementId: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJhbGxvd2VkQ291bnRyaWVzIjpbIklFIiwiRlIiLCJOTCIsIlNJIiwiUk8iLCJCRyIsIkNZIiwiSVMiLCJERSIsIkZJIiwiTFYiLCJQTCIsIlBUIiwiTFUiLCJIUiIsIkVTIiwiQVQiLCJNVCIsIlNFIiwiR1IiLCJJVCIsIkhVIiwiRUUiLCJESyIsIkxJIiwiTFQiLCJTSyIsIkNaIiwiQkUiLCJOTyJdLCJtYXJrZXRJZHMiOlsiU0UiXSwiZXhwIjoxNzE0OTA1NTMwLCJhZHMiOmZhbHNlLCJpYXQiOjE3MTQ4NjIzMzAsInN1YiI6IjE1NktFRkpETEM3SE1JS0FVTlNRQkc4TFNVU1IiLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaXNzIjoiZHJha2VuZmlsbSIsIm9mZmxpbmVFeHBpcmF0aW9uIjoxNzE1MzEyNTAyLCJ1c2VyVGFncyI6W10sIm9mZmVySWQiOiI3OTNGTzg4NFFZVUJXRVM0TTEyRk1ESThKT0ZSIiwiYXNzZXRJZCI6IjgxNDk0NTVjLWNiM2QtNGIxNS04NWE4LWI5NWUzZDE1NzBiNSJ9.tNkc_ZsE2j1cZwNLg0wqpucU3HJcmvwWd0sqVlqijH5bhVEVON9Xmf5mBlb_UQLLhUW_3mMCTvpOfyn38yZclQ
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODYyMzI3LCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.sYyyMBA7gf0q7A9na8E-vkgJntedFYn2pk_LX2WYBgdQgLgNs7xrtUgR2ZoZlMhgN6D5rQj2U6WDzvDUHZCqEQ
Magine-AccessToken: 22cc71a2-8b77-4819-95b0-8c90f4cf5663
Origin: https://drakenfilm.se
Connection: keep-alive
Referer: https://drakenfilm.se/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
Content-Length: 0
~~~

- Magine-Accesstoken is hard coded in the JavaScript
- Magine-Play-Entitlementid is hard coded in the HTML
- asset ID is hard coded in the HTML

Authorization comes from:

~~~
POST https://drakenfilm.se/api/auth/login HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/json, text/plain, */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-type: application/json
sentry-trace: 7f93a62bbdb64d01b0ee9a5f3c115b6e-a99f15898000009b-1
baggage: sentry-environment=production,sentry-release=qeHrYU2a8Vg2nUXvtEyhm,sentry-public_key=713cdb451f5f4b5ba5ff8b451868c7c5,sentry-trace_id=7f93a62bbdb64d01b0ee9a5f3c115b6e,sentry-sample_rate=1,sentry-transaction=%2Fmovie%2F%5Bslug%5D,sentry-sampled=true
content-length: 57
origin: https://drakenfilm.se
referer: https://drakenfilm.se/film/michael-clayton
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-origin
te: trailers

{"identity":"EMAIL","accessKey":"PASSWORD"}
~~~
