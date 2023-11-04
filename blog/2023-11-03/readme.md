# november 3 2023

this is it:

~~~
GET https://0b5f2c65670e498882d9cf4c89322cb4.mediatailor.us-east-2.amazonaws.com/v1/dash/7f34bf1814de6fddce84b1e6c296b7a70243b88f/oneapp-atp-dash-vod-2s-generic/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd?aws.sessionId=1eb208f8-c060-48aa-a1a1-78035d0a2444 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
referer: https://www.nbc.com/
origin: https://www.nbc.com
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers
content-length: 0
~~~

from:

~~~
POST https://mt.ssai-oneapp.nbcuni.com/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd?mt.config=oneapp-atp-dash-vod-2s-generic HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.nbc.com/
content-type: application/json
origin: https://www.nbc.com
content-length: 1575
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers

{"reportingMode":"client","availSuppression":{"mode":"BEHIND_LIVE_EDGE","value":"00:00:00"},"playerParams":{"origin_domain":"vod-lf-oneapp2-prd.akamaized.net"},"adsParams":{"caid":"NBC_VOD_9000283422","mode":"on-demand","pvrn":"394767237","vprn":"572200546","vdur":"4067","vip":"72.181.23.38","sfid":"9244655","csid":"oneapp_desktop_computer_web_ondemand","crtp":"vast3ap","nw":"169843","prof":"169843%3Aoneapp_web_vod_mt_atp","metr":"1031","flag":"%2Bsltp%2Bemcr%2Bslcb%2Bsbid-fbad%2Baeti%2Bslif%2Bvicb%2Bamcb%2Bplay-uapl%2Bdtrd","resp":"vmap1","afid":"200265138","vcid":"-1","am_abrspec":"not_required","am_appv":"1.224.3","am_buildv":"1.224.3","am_cdn":"akamai","am_crmid":"-1","am_crmid_type":"mParticleId","am_pubid":"-1","am_playerv":"v3.0.5-v58.hotfix","am_sdkv":"2.12.1-peacock","_fw_player_height":"1080","_fw_player_width":"1920","am_pub":"nbcu","am_appname":"oneapp","am_programtype":"television","am_nielsen_genre":"","nielsen_device_group":"","nielsen_platform":"","tms_id":"","gc_id":"","_fw_h_user_agent":"Mozilla%2F5.0%20(Windows%20NT%2010.0%3B%20Win64%3B%20x64%3B%20rv%3A101.0)%20Gecko%2F20100101%20Firefox%2F101.0","_fw_app_bundle":"","_fw_nielsen_app_id":"PAD3C6E72-ED61-417F-A865-3AB63FDB6197","_fw_cdn_provider":"nbcu_akamai","_fw_vcid2":"-1","_fw_h_referer":"https%3A%2F%2Fwww.nbc.com%2Fsaturday-night-live%2Fvideo%2Foctober-21-bad-bunny%2F9000283422","_fw_coppa":"0","_fw_ae":"","am_bc":"0","mini":"false","am_sst":"fullEpisodePlayer","pl":"n%2Fa","bl_enabled":"false","am_brand":"nbc","enable_pabi":"false","_fw_content_language":"en","yo.nl":"none"}}
~~~
