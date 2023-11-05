# NBC

~~~
POST https://drmproxy.digitalsvc.apps.nbcuni.com/drm-proxy/license/widevine?time=1699126072315&hash=0df6ff2b81e42c3ec10d1c4946cebce7ddd631a4130841fd74ba5fa0c3d7c02a&device=web HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.nbc.com/
content-type: application/octet-stream
origin: https://www.nbc.com
content-length: 5753
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers
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

https://github.com/hkato/mitmproxy-replace-response-body/blob/master/replace.py

can we replace the JavaScript? this is the target:

https://www.nbc.com/generetic/generated/chunks/12.ff734ba67f44a707e609.js

yes we can:

~~~
<link rel="preload" as="script" href="/generetic/generated/chunks/12.ff734ba67f44a707e609.js" />
<script async src=/generetic/generated/chunks/12.ff734ba67f44a707e609.js></script>
~~~

got it:

~~~
hash=e6c0f79f43950524cb36de6d19170f228ac7d75ea10e2574890724b2cdeb8f35 12.ff734ba67f44a707e609.js:1793:21
hash=85487b2dbb1020c3585d43a40ee25545dc465ec4cc499cc1bd7420e4858c5ad8 12.ff734ba67f44a707e609.js:1794:21
function(t,n){return new d.HMAC.init(e,n).finalize(t)} 12.ff734ba67f44a707e609.js:186:18
1699225330576widevine 12.ff734ba67f44a707e609.js:187:18
Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6
~~~
