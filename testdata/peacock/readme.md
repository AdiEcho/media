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

`bt` value is required and expires quickly. it comes from here:

~~~
POST https://play.ovp.peacocktv.com/video/playouts/vod HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/vnd.playvod.v1+json
accept-language: en-US,en;q=0.5
x-skyott-provider: NBCU
x-skyott-territory: US
x-skyott-proposition: NBCUOTT
x-skyott-platform: PC
x-skyott-device: COMPUTER
x-skyott-pinoverride: true
x-skyott-usertoken: 13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJ...
x-skyott-activeterritory: US
content-type: application/vnd.playvod.v1+json
x-sky-signature: SkyOTT client="NBCU-WEB-v8",signature="i9ZXGWBOY0IQM7Eehx47Kv9vfqw=",timestamp="1708737167",version="1.0"
origin: https://www.peacocktv.com
referer: https://www.peacocktv.com/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers

{
  "device": {
    "capabilities": [
      {
        "protection": "WIDEVINE",
        "container": "ISOBMFF",
        "transport": "DASH",
        "acodec": "AAC",
        "vcodec": "H264"
      },
      {
        "protection": "NONE",
        "container": "ISOBMFF",
        "transport": "DASH",
        "acodec": "AAC",
        "vcodec": "H264"
      }
    ],
    "maxVideoFormat": "HD",
    "model": "PC",
    "hdcpEnabled": true
  },
  "client": {
    "thirdParties": [
      "FREEWHEEL"
    ]
  },
  "contentId": "GMO_00000000224510_02_HDSDR",
  "personaParentalControlRating": "9"
}
~~~

https://play.google.com/store/apps/details?id=com.peacocktv.peacockandroid

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

you can get `x-skyott-usertoken` with web client via `/auth/tokens`, but it need
`idsession` cookie. what is used here?:

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

https://github.com/Paco8/plugin.video.skyott/issues/39
