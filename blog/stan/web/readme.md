# web

~~~
POST https://www.google.com/recaptcha/api2/reload?k=6LdhP7QUAAAAADQ9h... HTTP/2.0
accept-encoding: gzip, deflate, br
accept-language: en-US,en;q=0.5
accept: */*
content-length: 7011
content-type: application/x-protobuffer
origin: https://www.google.com
referer: https://www.google.com/recaptcha/api2/bframe?hl=en&v=Hq4JZivTyQ7GP8Kt...
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-origin
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~

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
cookie: _abck=363E815A356CA40FEABD7C7B6FDC26D5~0~YAAQJ6gtF5/ZhTiOAQAAl0yjdwu1ieZHfLTgpP9NvgZUWR7BvwaZPN7FjXNt+Zrs4YdbeyookoUFT6VcfWPaLieoen+lnRHjEBwmcpIb5hdGTtgSVkwd3WtAlGfJvGKOD1AW3XpwQe+toiyyeXntffZZHoVU2rECMnpSW8SheUeAQMcPr8vbFrcwO0vj8rdFFoZyTt3c6Vs08Ih7ybc9Sa8AGcDi9slFP5XFHSTT8uMVSO5VGF/WXLzR94SuMT4UiLj5n3Dpv7toOy8kEPS6dG72J6X2PaIrOrnR0S2LFcoGJeCWXXiej23TYQaG52MvTYnL9gjls005H8CLP7nMlN9sHJBW1j5x6z4xSNKMK/wDlzjNvHjCRw2lL/cEubgJdkaFmkyFWst5Gp6B8FBEo1TLHUCfOs0/6bg=~-1~-1~-1
cookie: _sp_id.94d0=1fc774d2-4337-476f-beae-9f311d38cfd9.1711404127.1.1711404176..cc9c64fa-cc5a-4f85-a381-b477ad5476d0..51272df0-0f98-4ef7-a929-0aaaf892cc16.1711404126944.3
cookie: _sp_ses.94d0=*
cookie: ak_bmsc=8FD13EFE7AF2573DE8CE7E675D3B060D~000000000000000000000000000000~YAAQJ6gtF9DbhTiOAQAAvXCjdxdK5YZL7YGI2D3XXuNyLCC9oVfQgiokRAIUFYtq3qfABsm3q7mPl9qfoclWgGCuySCcUdTvghnXR8LlIDRh+sM8yX9OKK1ozTZZknll2C9f9XGiXupiSZbfC1oE5202goo+sfgzjuPTVkWvBCSnZS214Q2emBW1MR5E+80YBJKNRuFCPFObN2/27Kkvnj7P/NlXCObZroeoVXSIKt7bFR05m9dx7KEayCMG3R9myAM2JiHtQsrTKCTqZaoQ281/jHO8m6YWyDIS6DyMzJvbUQD/06EYd0kM2IuAcaUqSR+EugNypfIdXSY4QgC1pIKX7UANxIEqBSCo8kY3e1KY0BdKndovuJjUWs9JGNryAm84mmoFI39EcmjdD3eDStUU
cookie: bm_sv=61D6A7DFA4D9C204EB23E08FA63A4CEF~YAAQJ6gtF9HbhTiOAQAAvXCjdxc77/mMynsvs1ne4zWOvXktP1t8fmLrLlMyg/LuBVJtvRdwr13KAUmFDLssyIqDPXX1TkNzy5gjM3KJy7laKAe6HdALT7+p5OF3E4xbKp8/31uqLL/QhRHhnSPcUJik0C5ydsOUrn1WKkJTMmhFvg6VpRBfiBOwbbj8nwpG5ByRNFCAJSEuY6xSrsJOQ2ww9sovg1zJTv8pxK7z8Rtwr6K3tUhJ8oUxt1RB1nF4yG8=~1
cookie: bm_sz=029DF08316F6BB523E69BBD833D51EE2~YAAQJ6gtF6TLhTiOAQAAd0Widxd0lLV5x+9eoSJXbdLF4FexXa8YvcbCTyPU1WIBBjpZ2lpQU7IYZt2Oq5wdTWZQX/M4Wr9jFrQwFSsL0LqM3PxQW/MC9C8RMjYBEChlmKlnFhbIaSBfZb8Z7KGB/7G1q0iXIFeSBws4XxqA3ikSa0aYBYNkLov2+Xw/rS4Xjouf2fzEfApVapL2C4a39f9yJr/Tnm2av7VVl8WzwWOWoNFtFRLm18sGBsmewxk5RKpu7bPQxQEeuiiCbeAvYKK4hs20XZhVFvqosK05w/yZoZVJ/QrmmIPgCEBGZWnrzw0OcNwNk5uqAcTHabYScr5menHX/rq4zeNSyFCExxH8eL6zRGDPKw==~3748932~3682361
cookie: streamco_device_id=10ebcc94-3373-475f-a8a9-dc848b637caa
cookie: streamco_profile_expiry=1711389749
cookie: streamco_profileid=e07528fd3b444b148a246ffb393be652
cookie: streamco_token=eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxODMsImp0aSI6IjE4N2RhNDdhMTM4ZjQ3OTY5NjMzNTUwYTcyOWIwODY0IiwiaWF0IjoxNzExNDA0MTgzLCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsInR6IjoiQW1lcmljYS9DaGljYWdvIiwiYXBwIjoiU3Rhbi1XZWIiLCJ2ZXIiOiJiZWFkMDk2IiwiZmVhdCI6MzM1NjM2MTk4NH0.ZUARuf7IvBBDRsbNJxjhGN4AnVTnarF0t-ZGTv_TFjI
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
