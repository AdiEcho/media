# cine member

~~~
url = https://www.cinemember.nl/nl/films/american-hustle
monetization = FLATRATE
country = Belgium
country = Netherlands
~~~

## android

https://play.google.com/store/apps/details?id=nl.peoplesplayground.audienceplayer.cinemember

## web

this is it:

<https://u2.cdn.jetstre.am/D1RyXtu6VAzm7V8YXoM54w/1715619755/unified/audp2-prod/americanhustle-25fps-ov-lt-rt.ism/dash/americanhustle-25fps-ov-lt-rt-audio_eng=127999-480000.dash>

from:

https://u2.cdn.jetstre.am/D1RyXtu6VAzm7V8YXoM54w/1715619755/unified/audp2-prod/americanhustle-25fps-ov-lt-rt.ism/.mpd

from:

https://takeoff.jetstre.am/?account=audp2-prod&type=streaming&service=unified&protocol=https&file=americanhustle-25fps-ov-lt-rt.ism&output=index.mpd&token=66412e6a|cfa164eb03c137c163cea607a59c9416

from (VPN):

~~~
POST /graphql/2/user HTTP/1.1
Host: api.audienceplayer.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
Accept: application/json, text/plain, */*
Accept-Language: nl
Accept-Encoding: gzip, deflate, br
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiZDNhZDQzNzRhZTdkMWY5OTkyZmRhZGRkY2NiZTI0YTIwYTFiNjdiODg0YjNjYzJlOTM4MmQwZWU3YzQzNTdiZmQ1NjRmOWEzMGI0OWQzMjAiLCJpYXQiOjE3MTU1MzMzMTAsIm5iZiI6MTcxNTUzMzMxMCwiZXhwIjoyMDMwODkzMzEwLCJzdWIiOiIxMjM4NjMiLCJzY29wZXMiOlsiYXBpLXVzZXItYWNjZXNzIl0sImFwX3BpZCI6MiwiYXBfYWlkIjpudWxsLCJhcF9yaWQiOm51bGwsImFwX2tpZCI6bnVsbH0.tM2GLP7yGtT2hLyPteXJAEahSmMDdTWhi28A_8oLwf7U3aHmmyZrPSfk2Rwceai9jVu8HiDre8_JbXmr6gS7v7M2nur77cSkUAXA0IYfgdhKjO67YWmyCDzN27fh_Gur4je-uNcT8dw0gi4kURoXcnkjB6Er3AV8ktpPaXbRtmdeMBVzkNTAcUftvkfgGoftE6oUuFoSnL5Ra40JICAqHPiqSTtACRRxvJjSPSP9zm1oaH07Bj2oeQX711hhxZWvq1eXkr89VP984xGypOJYWkAA_g6HYH3TVupWpEmNlqov1h20PtHTekhcjh1lhmEr_dIY0n3QHogj9wQY8TRHG49Vl8p7Gi7a885ElEcU6OC9FJnU_lgT6_xbZxuLUZoxridDF6ikvCZA4WS91RiuHc9N8Nfy4SYPk0KYHP60bXC_qhMdYdcCY4u3RHhlVuRdr6YBmAbvWzTDogoKCckatBRuKnZLBOqy2Yvl7y02iM2wW0b4b2iE78aonmcGZcDDOT8iK39v8JQBwfKJfaPbKbUeC3MZXoU-a-DMKK8CcpTbTwUkNJrkY9D14rCjAtM4myHebNCs5rj9z8FgWAd205wfQuX2D5-0PZn0BACRvqpvijM7QJlFTwwA6NWmeOY7b8GVV-A07sG58U0Oal0nO2-VT_DC0xX4ZlwyWmOajKM
Content-Type: application/json
Content-Length: 1235
Origin: https://www.cinemember.nl
Connection: keep-alive
Referer: https://www.cinemember.nl/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site

{"operationName":"ArticleAssetPlay","variables":{"article_id":768,"asset_id":1415,"protocols":["dash","hls"],"is_offline_download":false},"query":"mutation ArticleAssetPlay($article_id: Int, $asset_id: Int, $protocols: [ArticlePlayProtocolEnum], $is_offline_download: Boolean, $resolution: ArticlePlayResolutionEnum, $device_model_context: DeviceModelContextEnum) {\n  ArticleAssetPlay(\n    article_id: $article_id\n    asset_id: $asset_id\n    protocols: $protocols\n    is_offline_download: $is_offline_download\n    resolution: $resolution\n    device_model_context: $device_model_context\n  ) {\n    article_id\n    asset_id\n    entitlements {\n      ...Entitlement\n    }\n    subtitles {\n      ...SubtitleFile\n    }\n    pulse_token\n    appa\n    appr\n    time_marker_end\n    user_subtitle_locale\n    user_audio_locale\n    aspect_ratio\n    issued_at(is_format_with_milliseconds: true)\n    fairplay_certificate_url\n  }\n}\n\nfragment Entitlement on ArticleAssetPlayEntitlement {\n  mime_type\n  protocol\n  manifest\n  token\n  encryption_type\n  key_delivery_url\n  download_expires_in\n  encryption_provider\n  hls_key_uri\n  media_provider\n}\n\nfragment SubtitleFile on File {\n  url\n  locale\n  locale_label\n}"}
~~~
