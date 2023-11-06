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
- https://github.com/httptoolkit/frida-interception-and-unpinning/issues/52
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

also value is optional:

https://lemonade.nbc.com/v1/vod/2410887629/9000283422?platform=web&programmingType=Full+Episode

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

web seems to be the better option:

~~~diff
--- a/mpeg_cenc
+++ b/mpeg_cenc_2sec
-            <S t="0" d="180180" r="124"/>
-            <S t="22522500" d="60060" r="0"/>
-            <S t="22582560" d="180180" r="164"/>
-            <S t="52312260" d="120120" r="0"/>
-            <S t="52432380" d="180180" r="30"/>
-            <S t="58017960" d="120120" r="0"/>
-            <S t="58138080" d="180180" r="181"/>
-            <S t="90930840" d="120120" r="0"/>
-            <S t="91050960" d="180180" r="99"/>
-            <S t="109068960" d="60060" r="0"/>
-            <S t="109129020" d="180180" r="70"/>
-            <S t="121921800" d="114114" r="0"/>
+            <S t="0" d="60060" r="2030"/>
+            <S t="121981860" d="54054" r="0"/>
~~~

## programmingType

if you provide an invalid value:

https://lemonade.nbc.com/v1/vod/2410887629/9000283422?platform=web&programmingType=Clips

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

these all return the same thing:

- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=mpeg-dash
- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=mpeg-dash&assetTypes=2SEC
- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=mpeg-dash&assetTypes=VBW

here is another option:

http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=m3u

but again its not usable:

~~~
#EXT-X-SESSION-KEY:
   KEYFORMAT="com.apple.streamingkeydelivery",
   KEYFORMATVERSIONS="1",
   METHOD=SAMPLE-AES,IV=0xfcf13caf41cb4ec7bcc918872de873b9,
   URI="skd://fcf13caf41cb4ec7bcc918872de873b9"
~~~

this option:

http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422

fails:

> Invalid URL

regarding locked content:

https://nbc.com/saturday-night-live/video/october-28-nate-bargatze/9000283426

this fails:

http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283426?formats=mpeg-dash

~~~json
{
   "description": "This content requires a valid, unexpired auth token.",
   "exception": "InvalidAuthToken",
   "isException": true,
   "responseCode": "403",
   "title": "Invalid Token"
}
~~~

this works:

https://lemonade.nbc.com/v1/vod/2410887629/9000283426?platform=web&programmingType=Full+Episode

## license

# NBC

~~~
POST https://drmproxy.digitalsvc.apps.nbcuni.com/drm-proxy/license/widevine?time=1699126072315&hash=0df6ff2b81e42c3ec10d1c4946cebce7ddd631a4130841fd74ba5fa0c3d7c02a&device=web HTTP/2.0
content-type: application/octet-stream
~~~

in the code we have this:

~~~js
return ""
   .concat(const174_.drmProxyUrl, "/")
   .concat(const181_, "?time=")
   .concat(const182_, "&hash=")
   .concat(const184_, "&device=web")
   .concat(param178_2 ? "&keyId=".concat(param178_2) : "");
~~~

simplify:

~~~js
var91_ = param74_3(1358),
var92_ = param74_3.n(var91_),
let let179_ = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : 0;
const182_ = new Date().getTime() + let179_,
const const181_ = param178_.toLowerCase(),
const183_ = const182_ + const181_,
const184_ = var92_()(const183_, const174_.drmProxySecret);
"&hash=" + const184_
~~~

we should print these:

~~~js
console.log('DRM', var92_().toString());
console.log('DRM', const183_);
console.log('DRM', const174_.drmProxySecret);
~~~

script:

~~~py
from mitmproxy import ctx, http

def response(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/generetic/generated/chunks/12.ff734ba67f44a707e609.js'):
      f.response.text = open('2.js', 'r').read()
~~~

examples:

~~~
hash=85487b2dbb1020c3585d43a40ee25545dc465ec4cc499cc1bd7420e4858c5ad8 
function(t,n){return new d.HMAC.init(e,n).finalize(t)} 
1699225330576widevine 
Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6
~~~
