# peacock

<https://peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR>

this is it:

~~~
POST https://play.ovp.peacocktv.com/drm/widevine/acquirelicense?bt=8-XpyykNa... HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: */*
accept-language: en-US,en;q=0.5
referer: https://www.peacocktv.com/
content-type: text/plain
x-sky-signature: SkyOTT client="NBCU-WEB-v8",signature="cYBRu9SEy2nS2rHzRvHtlKnucF4=",timestamp="1708737168",version="1.0"
origin: https://www.peacocktv.com
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
~~~

what is used here?:

https://github.com/Paco8/plugin.video.skyott/blob/main/resources/lib/signature.py

~~~py
headers['x-skyott-usertoken'] = self.account['user_token']
~~~

then:

~~~py
self.account['user_token'] = data['userToken']
~~~

then:

~~~py
data = self.get_tokens()
~~~

then:

~~~py
def get_tokens(self):
   url = self.endpoints['tokens']
   headers = {}
   headers['cookie'] = self.account['cookie']
~~~

- https://github.com/Paco8/plugin.video.skyott/issues/39
- https://github.com/larsenv/YouTube-TV-Downloader/issues/1

~~~
> play -a com.peacocktv.peacockandroid
downloads: 34.82 million
files: APK APK APK APK
name: Peacock TV: Stream TV & Movies
offered by: Peacock TV LLC
price: 0 USD
requires: 7.0 and up
size: 67.11 megabyte
updated on: Feb 7, 2024
version code: 124050214
version name: 5.2.14
~~~

https://play.google.com/store/apps/details?id=com.peacocktv.peacockandroid

If you start the app and Sign In, this request:

~~~
POST https://rango.id.peacocktv.com/signin/service/international HTTP/2.0
content-type: application/x-www-form-urlencoded
x-skyott-device: MOBILE
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US

userIdentifier=MY_EMAIL&password=MY_PASSWORD
~~~

will fail:

~~~
HTTP/2.0 429
~~~

You can fix this problem by removing this request header before starting the
app:

~~~
set modify_headers '/~u signin.service.international/x-skyott-device/'
~~~

Header needs to be removed from that request only, since other requests need the
header. you can get `x-skyott-usertoken` with web client via `/auth/tokens`,
but it need `idsession` cookie. Looks like Android is the same.
