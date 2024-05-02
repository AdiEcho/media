# Android

## com.roku.remote

remote only:

https://play.google.com/store/apps/details?id=com.roku.remote

## com.roku.trc 

broken:

https://apkmirror.com/apk/roku-inc/the-roku-channel

## com.roku.web.trc

The Roku Channel (Android TV):

https://play.google.com/store/apps/details?id=com.roku.web.trc

create Android 9 device. install system certificate.

this is it:

<https://vod.delivery.roku.com/e28e3f8b5f104e1cbc1230d5279338e9/6b5b61c51d63474e9bd0da2afb68f451/9112014b44d444a79a902e7f2cabca40/8b5c1ef7d46c4425801d0fc32deff84a/index_video_1_0_6.mp4>

from:

<https://vod-playlist.sr.roku.com/1.mpd?origin=https%3A%2F%2Fvod.delivery.roku.com%2Fe28e3f8b5f104e1cbc1230d5279338e9%2F6b5b61c51d63474e9bd0da2afb68f451%2F40563e7d535e401c86db08e3c5794a22%2Findex.mpd%3Faws.manifestfilter%3Dsubtitle_language%3Aunused&ovpFilter=descriptiveAudio>

from:

~~~
POST https://googletv.web.roku.com/api/v3/playback HTTP/2.0
content-length: 396
origin: file://
x-roku-reserved-culture-code: en-US
x-roku-reserved-rida: 0438834b-91db-5327-811c-c9b5b64f7e7c
x-roku-reserved-lat: 0
user-agent: Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2
x-roku-reserved-experiment-configs: e30=
content-type: application/json
x-roku-code-version: 2
x-roku-content-token: ZizZafDEidXT//m06oXf0aepPFcxQ+HTY9iyMtVN7JbWpq6XcHp3hrkdZO6PKbCL3WVNoPkEl95XWEOMGVJEZzRtqFKvfvzYpMyNsNU0+WXJgVrQi8NlHcXKx227Ptx3/UyOwUx1bwP9TKVIQATnPuyW5VKsT0RAXuT+Qb2jAFs+x0ozsOyWzmQr3DAtrUNG5qR/7aGFi/0vasW6MsmD49AXxurDMoMW3N6w749+kUN/Mn+vVVmhGUIurxVpR30y7GYpMo4M6pV2TVn6NYC5HgqRP6l+e1/KugHi8L5cDIBSNngQVO+ASNPQj8rOv11H7E/qVNzprrwaFb8aOyjgjjob9N6WLQOU8I8gLGmkwDixppYtuCGoNY2ooh141NOnlqUvzWj9yfHFkeOWhosHVQ==
x-roku-reserved-channel-store-code: us
x-roku-reserved-time-zone-offset: +00:00
x-roku-reserved-experiment-state: W10=
x-roku-reserved-session-id: 108dcd6f-99e7-40ef-a9c5-5739dabfd048
accept: */*
accept-encoding: gzip, deflate
accept-language: en-US
x-requested-with: com.roku.web.trc

{"rokuId":"597a64a4a25c5bf6af4a8c7053049a6f","playId":"s-roku_originals.NGViNmY3MWMtMDM3NC00NDAzLThlNzMtODMzMzRjMWNhNjJi","mediaFormat":"DASH","drmType":"widevine","id":"597a64a4a25c5bf6af4a8c7053049a6f","quality":"fhd","bifUrl":"https://static-delivery.sr.roku.com/4eb6f71c-0374-4403-8e73-83334c1ca62b/images/4eb6f71c-0374-4403-8e73-83334c1ca62b-fhd.bif","adPolicyId":"","providerId":"rokuavod"}
~~~

from:

~~~
GET https://googletv.web.roku.com/api/v1/account/token?includeUserInfo=true&includeAccountLat=true&duid=e352cddd4c9ddc6b&getRida=false&platformName=googletv HTTP/2.0
x-roku-code-version: 2
x-roku-reserved-culture-code: en-US
x-roku-reserved-channel-store-code: us
x-roku-reserved-time-zone-offset: +00:00
user-agent: Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2
x-roku-reserved-session-id: 108dcd6f-99e7-40ef-a9c5-5739dabfd048
accept: */*
accept-language: en-US
x-requested-with: com.roku.web.trc
content-length: 0
~~~
