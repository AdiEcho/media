# draken

- https://drakenfilm.se/film/michael-clayton
- https://justwatch.com/se/leverantör/draken-films

241 LOC:

https://github.com/kenil261615/VineTrimmeronColab/blob/main/vinetrimmer/services/flixole.py

554 LOC:

https://github.com/hypnotoad/kodinerds/blob/master/plugin.video.magine/default.py

## web

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

## android

https://play.google.com/store/apps/details?id=com.draken.android

~~~
> play -i com.draken.android
details[6] = Draken Film
details[8] = 0 USD
details[13][1][4] = 4.5.0
details[13][1][16] = Feb 15, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
downloads = 27.68 thousand
name = Draken Film
size = 14.65 megabyte
version code = 1707910466
~~~

https://apkcombo.com/draken-film/com.draken.android

Create Android 6 device. Install user certificate. this is it:

~~~
POST /api/playback/v1/preflight/asset/17139357-ed0b-4a16-8be6-69e418c4ba40 HTTP/1.1
Magine-Play-DRM: widevine
Magine-Play-Protocol: dashs
Magine-Play-EntitlementId: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJhbGxvd2VkQ291bnRyaWVzIjpbIkxWIiwiUEwiLCJOTyIsIkxVIiwiTkwiLCJDWiIsIkVFIiwiREsiLCJGUiIsIlJPIiwiRkkiLCJFUyIsIlNJIiwiQkciLCJDWSIsIlNFIiwiQVQiLCJIUiIsIkhVIiwiTVQiLCJJUyIsIklUIiwiU0siLCJHUiIsIkxUIiwiSUUiLCJQVCIsIkxJIiwiQkUiLCJERSJdLCJtYXJrZXRJZHMiOlsiU0UiXSwiZXhwIjoxNzE0OTEyMDM3LCJhZHMiOmZhbHNlLCJpYXQiOjE3MTQ4Njg4MzcsInN1YiI6IjE1NktFRkpETEM3SE1JS0FVTlNRQkc4TFNVU1IiLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaXNzIjoiZHJha2VuZmlsbSIsIm9mZmxpbmVFeHBpcmF0aW9uIjoxNzE1MzEyNTAyLCJ1c2VyVGFncyI6W10sIm9mZmVySWQiOiI3OTNGTzg4NFFZVUJXRVM0TTEyRk1ESThKT0ZSIiwiYXNzZXRJZCI6IjE3MTM5MzU3LWVkMGItNGExNi04YmU2LTY5ZTQxOGM0YmE0MCJ9.07zrQ-4WuxGsVRcp4xzrYuf-UpCMkRjv7fPwy30nl_bf5MQIWiQg70cT5Hv4A_HlMvYoFqYAZ6TDrEdoJTpbRQ
Magine-Play-DeviceId: b18948c2-3b7f-4745-ac85-60de0dae6008
Magine-Play-DeviceType: smartphone
Magine-Play-DeviceModel: unknown/Android SDK built for x86
Magine-Play-DevicePlatform: android
User-Agent: Draken Film/4.5.0/1707910466;Magine-SDK/3.3.1;os=Android/6.0;device=Android SDK built for x86
Accept-Language: en-US;q=1.0, en-US;q=0.9, *;q=0.1
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODY4NzExLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.5Km1FYvwuj5J3aMepNxwzbjY3w7aXfWcWtRD9q3stwBC3nZgKSMOIJwhotQQF1r4SJqBKRC5sGoz7p77LDoKRA
Magine-AccessToken: 22cc71a2-8b77-4819-95b0-8c90f4cf5663
Content-Length: 0
Host: client-api.magine.com
Connection: Keep-Alive
Accept-Encoding: gzip
~~~

Magine-Play-Entitlementid comes from:

~~~
POST /api/entitlement/v2/asset/17139357-ed0b-4a16-8be6-69e418c4ba40 HTTP/1.1
User-Agent: Draken Film/4.5.0/1707910466;Magine-SDK/3.3.1;os=Android/6.0;device=Android SDK built for x86
Accept-Language: en-US;q=1.0, en-US;q=0.9, *;q=0.1
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODY4NzExLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.5Km1FYvwuj5J3aMepNxwzbjY3w7aXfWcWtRD9q3stwBC3nZgKSMOIJwhotQQF1r4SJqBKRC5sGoz7p77LDoKRA
Magine-AccessToken: 22cc71a2-8b77-4819-95b0-8c90f4cf5663
Content-Length: 0
Host: client-api.magine.com
Connection: Keep-Alive
~~~

Magine-AccessToken in the APK:

~~~java
@Override // wc.b
public String a() {
  return "22cc71a2-8b77-4819-95b0-8c90f4cf5663";
}

@Override // wc.b
public String b() {
  return "ea6fc0bb-8352-4bd6-9c4d-040a2c478fe8";
}
~~~
