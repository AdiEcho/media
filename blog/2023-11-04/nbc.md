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

if we enable, we get this:

https://lemonade.nbc.com/v1/vod/2410887629/9000283422?platform=web&browser=other&programmingType=Full+Episode

## browser

if you change user agent:

~~~
general.useragent.override
~~~

to:

~~~
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Safari/605.1.15
~~~

you get this instead:

https://lemonade.nbc.com/v1/vod/2410887629/9000283422?platform=web&browser=safari&programmingType=Full+Episode

but it doesnt help:

~~~
#EXT-X-SESSION-KEY:
   IV=0xfcf13caf41cb4ec7bcc918872de873b9,
   KEYFORMAT="com.apple.streamingkeydelivery",
   KEYFORMATVERSIONS="1",
   METHOD=SAMPLE-AES,
   URI="skd://fcf13caf41cb4ec7bcc918872de873b9"
~~~

## platform

web:

~~~json
{
  "playbackUrl": "https://vod-lf-oneapp2-prd.akamaized.net/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd",
  "type": "DASH"
}
~~~

android:

~~~json
{
  "playbackUrl": "https://vod-lf-oneapp2-prd.akamaized.net/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc/master_cmaf.mpd",
  "type": "DASH"
}
~~~

## programmingType

if you provide an invalid value:

https://lemonade.nbc.com/v1/vod/2410887629/9000283422?platform=web&browser=other&programmingType=Clips

you get:

~~~json
{
  "code": 400,
  "error": "No AssetType/ProtectionScheme/Format Matches",
  "message": "Bad Request",
  "meta": {
    "mpxUrl": "https://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=mpeg-dash&assetTypes=NONE,2SEC,VBW&restriction=108697384&sig=006546b8badb29ec17d8fb9f393733900635f73326b600a1a1736563726574",
    "message": {
      "title": "No AssetType/ProtectionScheme/Format Matches",
      "description": "None of the available releases match the specified AssetType, ProtectionScheme, and/or Format preferences",
      "isException": true,
      "exception": "NoAssetTypeFormatMatches",
      "responseCode": "412"
    }
  }
}
~~~
