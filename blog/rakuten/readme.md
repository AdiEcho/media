# rakuten.tv

## android

https://play.google.com/store/apps/details?id=tv.wuaki

~~~
> play -i tv.wuaki
details[6] = Rakuten TV
details[8] = 0 USD
details[13][1][4] = 3.29.3
details[13][1][16] = Apr 30, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 6.0 and up
downloads = 5.01 million
name = Rakuten TV -Movies & TV Series
size = 30.30 megabyte
version code = 3290300
~~~

https://apkmirror.com/apk/rakuten-tv/rakuten-tv-movies-tv-series-android-tv

create Android 9 device

## france

~~~
url = https://www.rakuten.tv/fr/movies/jerry-maguire
monetization = ADS
country = France
~~~

1. S'inscrire (sign up)
2. Votre adresse électronique (your e-mail address)
   - fr-2024-5-20@mailsac.com
3. password
4. J’accepte les conditions d’utilisation (I accept the terms of use)
5. Créez un compte (create an account)

## sweden

~~~
url = https://www.rakuten.tv/se/movies/i-heart-huckabees
monetization = ADS
country = Sweden
~~~

smart proxy:

~~~
mitmproxy `
--mode upstream:http://se.smartproxy.com:20001 `
--upstream-auth USER:PASS
~~~

1. Se gratis med annonser (watch for free with ads)
2. Registrera dig (register yourself)
3. Din e-postadress (your e-mail address)
4. password
5. Jag accepterar användningsvillkoren (I accept the terms of use)
6. Skapa konto (create an account)

## web

this is it:

https://prod-kami.wuaki.tv/v1/delivery/dash/avod/683de23c-fbb1-4541-9c68-f1f7e63dcd8e.mpd

from:

~~~
POST https://gizmo.rakuten.tv/v3/avod/streamings?classification_id=23&device_identifier=web&device_stream_audio_quality=2.0&device_stream_hdr_type=NONE&device_stream_video_quality=FHD&disable_dash_legacy_packages=false&locale=fr&market_code=fr HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/json, text/plain, */*
accept-language: en-US,en;q=0.5
content-type: application/json
origin: https://www.rakuten.tv
referer: https://www.rakuten.tv/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers

{"audio_language":"FRA","audio_quality":"2.0","classification_id":"23","content_id":"jerry-maguire","content_type":"movies","device_make":"firefox","device_model":"GENERIC","device_serial":"not implemented","device_stream_video_quality":"HD","device_uid":"4505abd8-0dc1-4552-af62-809fe3b07413","device_year":1970,"gam_correlator":1501924200588823,"gdpr_consent_opt_out":"0","gdpr_consent":"","hdr_type":"NONE","ifa_subscriber_id":null,"ifa_type":"ppid","player_height":1080,"player_width":1920,"player":"web:DASH-CENC:WVM","publisher_provided_id":"fc07d4f2-53e8-4a85-8aea-1e822bb9ea12","strict_video_quality":false,"subtitle_formats":["vtt"],"subtitle_language":"MIS","support_closed_captions":true,"video_type":"stream","support_thumbnails":true,"app_bundle":"com.rakutentv.web","app_name":"RakutenTV","url":"rakuten.tv","requestedContent":{"live":false,"duration":8324,"genres":["Comédie","Drame","Romantique"],"id":50928,"rating":"Tous","title":"Jerry maguire","adOptions":{"channelID":3273,"gdprAccepted":false,"gdprConsentString":"","limitTargetedAds":1,"playerOrigin":"playerpage","gdprConsentOptOut":"0","publisherProvidedId":"fc07d4f2-53e8-4a85-8aea-1e822bb9ea12"}}}
~~~
