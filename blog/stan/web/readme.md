# web

this is it:

https://gec.stan.video/09/dash/live/1540676B-1/sd/sdr/video/avc/200/seg-1.m4f

from:

<https://api.stan.com.au/manifest/v1/dash/web.mpd?url=https%3A%2F%2Fgec.stan.video%2F09%2Fdash%2Flive%2F1540676B-1%2Fsd%2Fsdr%2Fmedium_h264-7ddb6702.mpd%3Fcdns%3Dakamai%252Camazon%252Cgoogle%26clearaudio%3Dtrue%26maxQuality%3Dsd&audioType=all&livecaptions=true&captionsUrl=https%3A%2F%2Fapi.stan.com.au%2Fmanifest%2Fv1%2Fvtt%2Fweb.vtt%3Ftype%3Dfirefox%26url%3Dhttps%253A%252F%252Fstreamcott-a.akamaihd.net%252F4d23a7729757b3c3111a84e95ffb2044cd7d1b326409a6f5%252F1540676B-1%252Feng.vtt&captionsLang=eng&version=111>

from (geo block):

~~~
POST https://api.stan.com.au/concurrency/v1/streams?programId=1540676&jwToken=eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxODUsImp0aSI6IjJiZThmYTBkYTg2NTQ0Njk4ZWUwMjg3YzdiZDc3YTIyIiwiaWF0IjoxNzExNDA0MTg1LCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsInR6IjoiQW1lcmljYS9DaGljYWdvIiwiYXBwIjoiU3Rhbi1XZWIiLCJ2ZXIiOiJiZWFkMDk2IiwiZmVhdCI6MzM1NjM2MTk4NH0.4A1MOC17P7bIA_hQhCowqhj1QSU-FJ5xJcyAktoASu4&clientId=10ebcc94-3373-475f-a8a9-dc848b637caa&format=dash&cid=afe301fc-5762-4d62-b917-539b753766e2&capabilities.drm=widevine&player=html5&quality=sd&stanName=Stan-Web HTTP/2.0
accept-encoding: gzip, deflate, br
accept-language: en-US,en;q=0.5
accept: */*
content-length: 0
content-type: application/x-www-form-urlencoded
origin: https://play.stan.com.au
referer: https://play.stan.com.au/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~

from:

~~~
POST https://api.stan.com.au/login/v1/sessions/web/app?manufacturer=Mozilla&os=Win32&sdk=Windows%20NT%2010.0%3B%20Win64%3B%20x64%3B%20rv%3A109.0&type=web&model=109.0&stanName=Stan-Web&stanVersion=bead096&clientId=10ebcc94-3373-475f-a8a9-dc848b637caa&tz=America%2FChicago HTTP/2.0
accept-language: en-US,en;q=0.5
accept: */*
content-type: application/x-www-form-urlencoded
cookie: _abck=363E815A356CA40FEABD7C7B6FDC26D5~0~YAAQJ6gtF5/ZhTiOAQAAl0yjdwu1i...
cookie: _sp_id.94d0=1fc774d2-4337-476f-beae-9f311d38cfd9.1711404127.1.17114041...
cookie: _sp_ses.94d0=*
cookie: ak_bmsc=8FD13EFE7AF2573DE8CE7E675D3B060D~00000000000000000000000000000...
cookie: bm_sv=61D6A7DFA4D9C204EB23E08FA63A4CEF~YAAQJ6gtF9HbhTiOAQAAvXCjdxc77/m...
cookie: bm_sz=029DF08316F6BB523E69BBD833D51EE2~YAAQJ6gtF6TLhTiOAQAAd0Widxd0lLV...
cookie: streamco_device_id=10ebcc94-3373-475f-a8a9-dc848b637caa
cookie: streamco_profile_expiry=1711389749
cookie: streamco_profileid=e07528fd3b444b148a246ffb393be652
cookie: streamco_token=eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1Q...
cookie: streamco_uid=e07528fd3b444b148a246ffb393be652
origin: https://play.stan.com.au
referer: https://play.stan.com.au/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0

jwToken=eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxODMsImp0aSI6IjE4N2RhNDdhMTM4ZjQ3OTY5NjMzNTUwYTcyOWIwODY0IiwiaWF0IjoxNzExNDA0MTgzLCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsInR6IjoiQW1lcmljYS9DaGljYWdvIiwiYXBwIjoiU3Rhbi1XZWIiLCJ2ZXIiOiJiZWFkMDk2IiwiZmVhdCI6MzM1NjM2MTk4NH0.ZUARuf7IvBBDRsbNJxjhGN4AnVTnarF0t-ZGTv_TFjI&profileId=e07528fd3b444b148a246ffb393be652&pin=&source=play
~~~

from:

~~~
POST https://api.stan.com.au/login/v1/sessions/web/account/recapture HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.stan.com.au/
content-type: application/x-www-form-urlencoded; charset=utf-8
content-length: 854
origin: https://www.stan.com.au
cookie: _abck=363E815A356CA40FEABD7C7B6FDC26D5~0~YAAQJ6gtF9fQhTiOAQAAwKmidwvOQ...
cookie: ak_bmsc=16F975DB410EFED225E95901B079C55B~00000000000000000000000000000...
cookie: bm_sz=029DF08316F6BB523E69BBD833D51EE2~YAAQJ6gtF6TLhTiOAQAAd0Widxd0lLV...
cookie: bm_sv=61D6A7DFA4D9C204EB23E08FA63A4CEF~YAAQJ6gtF9nQhTiOAQAAwKmidxfKRbU...
cookie: _sp_ses.94d0=*
cookie: _sp_id.94d0=1fc774d2-4337-476f-beae-9f311d38cfd9.1711404127.1.17114041...
cookie: streamco_device_id=10ebcc94-3373-475f-a8a9-dc848b637caa
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers

email=EMAIL&password=PASSWORD&source=www&clientId=10ebcc94-3373-475f-a8a9-dc848b637caa&stanName=Stan-Web&recaptureToken=03AFcWeA7cTFBKeqtwbnwagP5TZZJA2i_2TyVHvjZ4obYNCYs0i9sr7iI0atqlSlohsdCTdN1ynFj4O4Y4HKh9S_UiiDOjnZ4h-h24whgFW2t_2nBei2bZIuvHvjAJ1xa50L33LXizpYS8i4SluB3gz7_Hkzr0r2eAw8LlP4L-9ESdxHtl1GlDBGwKIAzylKwcXj6b8dQeJ8vQxUgFWPsLab3ep8N4nBeuLJTGkRK6TZoJ3QcYp9jqYyYi9F6JWjLvTQysi6Pclz1jB74Uycg2Xx1sqCebvNiuupChni2cZLPvDYqBaD8BalDYiPIY8d3VMxldFmf0cndx-HRsawmLeaqIJB7wMt0SZHW0yZP7l4Psl9Ny_BGU8Tn1-lPlCIytxPNRn_jRDvTQzntyudz1IaBirLWHYf7wvc9LP-g5E94ZzNx9q4KPqHuVLbrevNJyncUS3QRWyN0BnHtxHNlGJ80_eFTbnSUKlZTig7oxYyx4vNVRxa_udDb2AJMgOAXPkns8TYmOuRGpxipVVi3w2NwjKlj8aINlOsXHH8wul9Zq5Yc_VcS-2ubwEP_fkP4HmSPXqCLP0mV0pKBqkuTTPhMu7aXsbYANCU9s1OlxAK8KBjK8eLgOjOSgxOFOSAfN4cyLaNHbXF6mMgtD7jcydSR7wuyCJRnxIRLAWXUxbJdzzlELHTb1QQN3TLhxmgkH4-KhWVNLU1e5
~~~
