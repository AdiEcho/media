# criterionchannel


## android

https://play.google.com/store/apps/details?id=com.criterionchannel

~~~
> play -i com.criterionchannel
details[6] = The Criterion Collection
details[8] = 0 USD
details[13][1][4] = 8.701.1
details[13][1][16] = Apr 8, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
downloads = 192.95 thousand
name = The Criterion Channel
size = 31.98 megabyte
version code = 11271
~~~

Create Android 6 device. Install user certificate

~~~xml
<intent-filter>
   <action android:name="android.intent.action.VIEW"/>
   <category android:name="android.intent.category.DEFAULT"/>
   <category android:name="android.intent.category.BROWSABLE"/>
   <data android:scheme="@string/scheme"/>
</intent-filter>
~~~

then:

~~~
res\values\strings.xml
797:    <string name="scheme">vhxcriterionchannel</string>
~~~

login:

~~~
POST /v1/oauth/token HTTP/1.1
Accept: application/json
X-OTT-Agent: android site/59054 android-app/8.701.1
User-Agent: Criterion/8.701.1(Unknown Android SDK built for x86, Android 6.0 (API 23))
OTT-Client-Version: 8.701.1
X-OTT-Language: en_US
Content-Type: application/x-www-form-urlencoded
Content-Length: 227
Host: auth.vhx.com
Connection: Keep-Alive
Accept-Encoding: gzip

username=USERNAME&password=PASSWORD&grant_type=password&client_id=9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78&client_secret=145e2d9e209ae44a263069bb818a82ae39ceae64e7bc435d83028c094dcdae5f
~~~

## web

this is it:

https://drm.vimeocdn.com/1716012602-0xd8dc2aae8f6068ef3453aacf8f1b88750c884748/01/115/13/325579370/1273684310,1273684307,1273684308,1273684313/1273684313-4-t1.m4s?assetId=325579370

from:

https://drm.vimeocdn.com/1716012602-0xd8dc2aae8f6068ef3453aacf8f1b88750c884748/01/115/13/325579370/1273684310,1273684307,1273684308,1273684313/master.mpd?assetId=325579370&pssh=0&subtitles=7433271-English-en-cc

from (only good for one minute):

<https://player.vimeo.com/video/325579370/config?autoplay=1&color=b9bcc7&controls=1&speed=1&token=eyJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MTU5ODU4MDQsImV4cCI6MTcxNTk4NTg2NCwiY2xpcF9pZCI6MzI1NTc5MzcwLCJkcm0iOnsiaGRjcCI6dHJ1ZSwicHNzaCI6IjAiLCJzZCI6eyJhc3NldF9pZCI6IjMyNTU3OTM3MC1zZCIsImhkY3AiOmZhbHNlLCJtYXhfaGVpZ2h0Ijo3MjB9fX0.dQaoX_mbtY8ojcJJqHUFRMvzraXXuHqU13uu33XIr0E&trick_play=1>

from:

~~~
GET https://embed.criterionchannel.com/videos/455774?api=1&auth-user-token=eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo2NTc0NjE0NiwiZXhwIjoxNzE2MDA1ODAwfQ.w-9aX0vfcblyXo7Evya5gfFAvzIsV-7o36Mhsf20gao&autoplay=1&collection_id=95574&color=b9bcc7&context=https%3A%2F%2Fwww.criterionchannel.com%2Fmy-dinner-with-andre&is_trailer=false&live=0&locale=en&playsinline=1&product_id=39621&referrer=https%3A%2F%2Fwww.criterionchannel.com%2Fmy-dinner-with-andre&vimeo=1 HTTP/2.0
referer: https://www.criterionchannel.com/
~~~

from:

https://www.criterionchannel.com/my-dinner-with-andre

login:

~~~
POST https://www.criterionchannel.com/login HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.criterionchannel.com/login
content-type: application/x-www-form-urlencoded
content-length: 181
origin: https://www.criterionchannel.com
cookie: locale_det=en
cookie: _session=TEpHYVpYZ0M5dFVlQ0MzSVVtdVRJb1o3RTBFa2tOVWd3TUJaajA0THFUQ2svS09EUS82MDE1dnhacFVGQmpLYzUzZzFaWStoSnVaajJacnIwTGFucjdwZHU5bGpjU2EzK1REaVpNbFB6VGNCeEwyd3orenliejB2VVpyZXRSK3dSWE40NnNGS3lnR1p6eFBaQ2FwWHBTSm1ZR0gwK08reHZHNGZkaVpLUnNJNlQ2UUl2eHl3K0FsWkVBREdPZ1FSLS11endNQWlsY1Vqeld2VVFjSm55L2RRPT0%3D--2a64dbf1e5a4b48d3de0c3483692434fb59c9814
cookie: __cf_bm=OKkRV.zgaf6ZgSq.fI00JT0bGDujOvnrA2ACOggJqZU-1715985763-1.0.1.1-NdxPIngecU15kKF54nQmqc9wbdZQvuZD4c8zysJcxhWymGewBpij4gYIkWvRacuswaR.XWmHlQ7f1TmX6wxNNQ
cookie: tracker=%7B%22country%22%3A%22us%22%2C%22platform%22%3A%22windows%22%2C%22uid%22%3A6433832978435%2C%22site_id%22%3A%2259054%22%7D
cookie: referrer_url=https%3A%2F%2Fwww.criterionchannel.com%2Fmy-dinner-with-andre
upgrade-insecure-requests: 1
sec-fetch-dest: document
sec-fetch-mode: navigate
sec-fetch-site: same-origin
sec-fetch-user: ?1
te: trailers

email=EMAIL&authenticity_token=evo3lO59MN7AGAbYZMzCPYHn7Cvmr8Gw%2FKNge8wLgZ5KYD3DC6jKGdbgHuxRYvLWuHy%2BzAXOR6HzcsHaTvgnsQ%3D%3D&utf8=%E2%9C%93&password=PASSWORD
~~~
