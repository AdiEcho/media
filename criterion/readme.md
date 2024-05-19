# criterion

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

widevine:

~~~
POST /v2/widevine?token=eyJhbGciOiJIUzI1NiJ9.eyJhc3NldF9pZCI6MzI1NTc5MzcwLCJ1c2VyX2lkIjo2NTc0NjE0NiwidmlkZW9faWQiOjQ1NTc3NCwiY2xpcF9pZCI6MzI1NTc5MzcwLCJleHAiOjE3MTYwMTEyNTMsIm9mZmxpbmVfbGljZW5zZSI6ZmFsc2UsImhkY3AiOmZhbHNlfQ.bVifzIrmAxtYw1iXKKrGe_oe3kBeLelE8CYz4yM8IGI HTTP/1.1
Content-Type: application/octet-stream
Accept-Encoding: gzip
Content-Length: 2135
User-Agent: Dalvik/2.1.0 (Linux; U; Android 6.0; Android SDK built for x86 Build/MASTER)
Host: drm.vhx.com
Connection: Keep-Alive
~~~
