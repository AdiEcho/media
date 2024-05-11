# stream.sooner.nl

Netherlands

https://stream.sooner.nl/m/american-hustle

this is it:

<https://cfvod.kaltura.com/hls/p/2031841/sp/203184100/serveFlavor/entryId/1_ml2r8fll/v/2/pv/1/ev/8/flavorId/1_2clfqtfq/name/a.mp4/seg-3-v1-a1.ts>

from:

<https://cfvod.kaltura.com/hls/p/2031841/sp/203184100/serveFlavor/entryId/1_ml2r8fll/v/2/pv/1/ev/8/flavorId/1_2clfqtfq/name/a.mp4/index.m3u8>

from:

<https://cdnapisec.kaltura.com/p/2031841/sp/203184100/playManifest/entryId/1_ml2r8fll/protocol/https/format/applehttp/flavorIds/1_yhadq7wm,1_opo0wipl,1_q0tfn2ww,1_iblrs4u1,1_5rx1kjzm,1_2clfqtfq/ks/djJ8MjAzMTg0MXwoBmCwb9cwtz9vh2OP3vK7OOu80f01GGuxeKWUkoMjTaXuD12b7SZkxoxUgJ6gipLlBRP8MRrlDAGqFnTZRrQXB9-1rqWXY_BzCgYCCGYNkA==/a.m3u8?playSessionId=70573445-a5f9-1824-d989-9e66dc027f0a:546d43d1-7f26-297d-192a-73fe0f8875bc&referrer=aHR0cHM6Ly9zdHJlYW0uc29vbmVyLm5sL25sL3BsYXllcj90eXBlPWZ1bGwmY2lkPTE5NzIwOTQwMDAwMjM2ODY0NjI=&clientTag=html5:v3.17.2>

from:

~~~
POST https://cdnapisec.kaltura.com/api_v3/service/multirequest HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-type: application/json
content-length: 914
origin: https://stream.sooner.nl
referer: https://stream.sooner.nl/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers

{"1":{"service":"session","action":"startWidgetSession","widgetId":"_2031841"},"2":{"service":"baseEntry","action":"list","ks":"{1:result:ks}","filter":{"redirectFromEntryId":"1_kqvyiof1"},"responseProfile":{"type":1,"fields":"id,referenceId,name,description,thumbnailUrl,dataUrl,duration,msDuration,flavorParamsIds,mediaType,type,tags,dvrStatus,externalSourceType,status,createdAt,updatedAt,endDate,plays,views,downloadUrl,creatorId"}},"3":{"service":"baseEntry","action":"getPlaybackContext","entryId":"{2:result:objects:0:id}","ks":"{1:result:ks}","contextDataParams":{"objectType":"KalturaContextDataParams","flavorTags":"all"}},"4":{"service":"metadata_metadata","action":"list","filter":{"objectType":"KalturaMetadataFilter","objectIdEqual":"{2:result:objects:0:id}","metadataObjectTypeEqual":"1"},"ks":"{1:result:ks}"},"apiVersion":"3.3.0","format":1,"ks":"","clientTag":"html5:v3.17.2","partnerId":2031841}
~~~
