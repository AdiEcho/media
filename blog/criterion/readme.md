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
