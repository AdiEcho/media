# NBC

## how to get `drmProxySecret` value?

if you visit a page such as this:

https://nbc.com/saturday-night-live/video/january-20-jacob-elordi/9000283424

you should see a request like this:

https://www.nbc.com/generetic/generated/generetic.971bc5df5bddfb45624e.js

in the response body, you should see something like this:

~~~json
{
  "coreVideo": {
    "drmProxyUrl": "https://drmproxy.digitalsvc.apps.nbcuni.com/drm-proxy/license",
    "drmProxySecret": "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"
  }
}
~~~

## how to get `hash` value?

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
      f.response.text = open('hello.js', 'r').read()
~~~
