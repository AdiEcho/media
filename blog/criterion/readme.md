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

this is it:

~~~
GET /v2/sites/59054/videos/455774/delivery?offline_license=0&model=Android%20SDK%20built%20for%20x86&max_width=1920&max_height=1080&max_fps=60&codecs=hevc%2Cavc&os_version=6.0 HTTP/1.1
Host: api.vhx.com
X-OTT-Agent: android site/59054 android-app/8.701.1
User-Agent: Criterion/8.701.1(Unknown Android SDK built for x86, Android 6.0 (API 23))
OTT-Client-Version: 8.701.1
X-OTT-Language: en_US
Connection: Keep-Alive
Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E2NDUzNDE2ZTlmZmRjNjk4MDAxMDdhZTQ2ZWJhODg0YTU2ZDBjOGQ4NTYzMzgifQ.eyJhcHBfaWQiOjM0NDksImV4cCI6MTcxNTk5NzcxMSwibm9uY2UiOiI0Y2Q4NzQ5OTQ2N2FhMDY4Iiwic2NvcGVzIjpbXSwic2Vzc2lvbl9pZCI6IlZ1STJoSlNXVnBYc0MxR3k4ejR3TGc9PSIsInNpdGVfaWQiOjU5MDU0LCJ1c2VyX2lkIjo2NTc0NjE0Nn0.Jf5JBUT8rV3zDOKeSUYdlZVOcA1_Y6d4sSCxfkNKMUKULfb4woaSWTykLmmi0-zWPKwFDeyhIiBXJCKciw0iXTDpZ2-CpXpgdiAYH9xmWaxPiqNW6feaoHc5L2n3NLS1jy5EP24RtmJQiXu-4VFSRdQh4H6_tE2LB5MEh0UP-N8w3mc8BJzZTgf_jOqTvhAfbge0OiIUodWuwqBAQinDLgQqTOz9RXcCpiG1VsuOimLGNAPsbuWy4-FL1bNPPs_x1VMhuW0ZS2buksWKJIXw8efh84ZXHrqWKC3y9FwNxs_piYCWktbtobjcehMjGQ5wn_0pPprIPKSjkpcD5PH6yP2UkMUKObmmanlRRkftq4jt2gq8mC5Jv84qhugJy1LHaIYDb40q8IVgpZOxS8ovC1Dzh-8ljnzLfQsHN2vHIOfYvcvGzaZ1c3-_bfMGwgz0J3qNBIgp81WR8pHxzSnjEcQ7CavBQdH2fOijnTf9To5t2xaQMXy-PFE25GPljarklIEl-SN3vk82mNAzfto64DG3q6tk2TGRtkI5EDJO2niIGW846unU6XfXCCiaFwhyWNCkhSkWbNho64r-itS9q1HycNmlXYjD7sKlXGdkE2u-PP_qEswwYi-oSZzSUtzGj-ZWUPy3NNiIv2mEXzi_k4X6hamNcAtg5LlfIlY3EFg
~~~

from:

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
