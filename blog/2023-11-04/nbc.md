# NBC

only requires Android 5:

~~~
> play -a com.nbcuni.nbc
requires: 5.0 and up
version code: 2000004392
version name: 9.4.1
~~~

which means we can install up to Android 6 and install user certificate. If we
do that and start the app, we get:

> Unable to Connect
>
> Please check your internet connection and try again.

If we FORCE STOP the app, set No proxy and try again, it works. next we can try
system certificate. same result. this result also repeats with older versions
back to:

~~~
package: name='com.nbcuni.nbc' versionCode='2000003396' versionName='9.0.0'
compileSdkVersion='31' compileSdkVersionCodename='12'
~~~

with older versions:

~~~
package: name='com.nbcuni.nbc' versionCode='2000003368' versionName='7.35.1'
compileSdkVersion='31' compileSdkVersionCodename='12'
~~~

we get:

> ACTION REQUIRED
>
> To keep watching your favorite shows, movies and more, you'll need to update
> to the latest version of the NBC App.

next we can try:

~~~
pip install frida-tools
~~~

download and extract server:

https://github.com/frida/frida/releases

for example:

~~~
frida-server-16.1.4-android-x86.xz
~~~

then:

https://github.com/httptoolkit/frida-interception-and-unpinning/tree/feddf62

~~~
frida -U -f com.nbcuni.nbc -l frida-script.js
~~~

app gets stuck on loading screen. what if we install user certificate? same
result. what if we install system certificate? same result. next we can try
main branch:

https://github.com/httptoolkit/frida-interception-and-unpinning

~~~
frida -U `
-l config.js `
-l android/android-certificate-unpinning.js `
-f com.nbcuni.nbc
~~~

result:

~~~
Error: Address(): argument types do not match any of:
.overload('java.lang.String', 'int', 'javax.net.SocketFactory', 'javax.net.ssl.SSLSocketFactory', 'javax.net.ssl.HostnameVerifier', 'com.android.okhttp.CertificatePinner', 'com.android.okhttp.Authenticator', 'java.net.Proxy', 'java.util.List', 'java.util.List', 'java.net.ProxySelector')
at X (frida/node_modules/frida-java-bridge/lib/class-factory.js:622)
at value (frida/node_modules/frida-java-bridge/lib/class-factory.js:1141)
at e (frida/node_modules/frida-java-bridge/lib/class-factory.js:606)
at <anonymous> (D:\Desktop\frida-interception-and-unpinning-main\android\android-certificate-unpinning.js:170)
at apply (native)
at ne (frida/node_modules/frida-java-bridge/lib/class-factory.js:673)
at <anonymous> (frida/node_modules/frida-java-bridge/lib/class-factory.js:651)
~~~

- <https://github.com/Frida-Modules-Repo/frida_multiple_unpinning.js/issues/1>
- <https://github.com/Frida-Modules-Repo/ssl_unpinning.js/issues/1>
- https://github.com/0xXyc/SSL-Certpinning-Bypass/issues/1
- https://github.com/54m4r174n/SSL-Pinning-Bypass-Automation/issues/2
- https://github.com/Benson306/SSL-Bypass-Script/issues/2
- https://github.com/fdciabdul/Frida-Multiple-Bypass/issues/1
- https://github.com/hunterxxx/Frida-Bypass-SSL-Pinning/issues/2
- https://github.com/hyugogirubato/Frida-CodeShare/issues/1
- https://github.com/vicsanjinez/ANTI-ROOT-AND-SSL-PINNING2/issues/1

install app, then push server:

~~~
adb push frida-server-16.1.4-android-x86 /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server
~~~

https://httptoolkit.com/blog/android-reverse-engineering

if you disable Widevine, and visit:

https://nbc.com/saturday-night-live/video/october-21-bad-bunny/9000283422

you get:

> SORRY! WE'RE HAVING SOME TROUBLE.
>
> If the problem persists, please contact us and send us a note.

here:

~~~
GET https://0b5f2c65670e498882d9cf4c89322cb4.mediatailor.us-east-2.amazonaws.com/v1/dash/7f34bf1814de6fddce84b1e6c296b7a70243b88f/oneapp-atp-dash-vod-2s-generic/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd?aws.sessionId=bc6f0b49-2149-4e45-871a-859f0bbdd483 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.nbc.com/
origin: https://www.nbc.com
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers
content-length: 0
~~~

from:

~~~
POST https://mt.ssai-oneapp.nbcuni.com/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd?mt.config=oneapp-atp-dash-vod-2s-generic HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.nbc.com/
content-type: application/json
origin: https://www.nbc.com
content-length: 1575
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers

{"reportingMode":"client","availSuppression":{"mode":"BEHIND_LIVE_EDGE","value":"00:00:00"},"playerParams":{"origin_domain":"vod-lf-oneapp2-prd.akamaized.net"},"adsParams":{"caid":"NBC_VOD_9000283422","mode":"on-demand","pvrn":"517362333","vprn":"940555295","vdur":"4067","vip":"72.181.23.38","sfid":"9244655","csid":"oneapp_desktop_computer_web_ondemand","crtp":"vast3ap","nw":"169843","prof":"169843%3Aoneapp_web_vod_mt_atp","metr":"1031","flag":"%2Bsltp%2Bemcr%2Bslcb%2Bsbid-fbad%2Baeti%2Bslif%2Bvicb%2Bamcb%2Bplay-uapl%2Bdtrd","resp":"vmap1","afid":"200265138","vcid":"-1","am_abrspec":"not_required","am_appv":"1.224.3","am_buildv":"1.224.3","am_cdn":"akamai","am_crmid":"-1","am_crmid_type":"mParticleId","am_pubid":"-1","am_playerv":"v3.0.5-v58.hotfix","am_sdkv":"2.12.1-peacock","_fw_player_height":"1080","_fw_player_width":"1920","am_pub":"nbcu","am_appname":"oneapp","am_programtype":"television","am_nielsen_genre":"","nielsen_device_group":"","nielsen_platform":"","tms_id":"","gc_id":"","_fw_h_user_agent":"Mozilla%2F5.0%20(Windows%20NT%2010.0%3B%20Win64%3B%20x64%3B%20rv%3A101.0)%20Gecko%2F20100101%20Firefox%2F101.0","_fw_app_bundle":"","_fw_nielsen_app_id":"PAD3C6E72-ED61-417F-A865-3AB63FDB6197","_fw_cdn_provider":"nbcu_akamai","_fw_vcid2":"-1","_fw_h_referer":"https%3A%2F%2Fwww.nbc.com%2Fsaturday-night-live%2Fvideo%2Foctober-21-bad-bunny%2F9000283422","_fw_coppa":"0","_fw_ae":"","am_bc":"0","mini":"false","am_sst":"fullEpisodePlayer","pl":"n%2Fa","bl_enabled":"false","am_brand":"nbc","enable_pabi":"false","_fw_content_language":"en","yo.nl":"none"}}
~~~

from:

~~~
GET /v1/vod/2410887629/9000283422?platform=web&browser=other&programmingType=Full%20Episode HTTP/1.1
Host: lemonade.nbc.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
Accept: application/json
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Origin: https://www.nbc.com
Connection: keep-alive
Referer: https://www.nbc.com/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site
content-length: 0
~~~
