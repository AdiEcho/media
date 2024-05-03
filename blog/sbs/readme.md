# sbs

- https://www.sbs.com.au/ondemand/movie/closer/2229616195516
- https://www.sbs.com.au/ondemand/watch/2229616195516

## android

https://play.google.com/store/apps/details?id=com.sbs.ondemand.android

## 1

~~~
POST https://www.sbs.com.au/api/v3/janrain/auth_native_traditional?context=odwebsite HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/json, text/plain, */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-type: multipart/form-data; boundary=---------------------------3672811325816669204907760291
content-length: 530
origin: https://www.sbs.com.au
referer: https://www.sbs.com.au/ondemand/watch/2229616195516
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-origin
te: trailers

-----------------------------3672811325816669204907760291
Content-Disposition: form-data; name="user"

USER
-----------------------------3672811325816669204907760291
Content-Disposition: form-data; name="pass"

PASS
-----------------------------3672811325816669204907760291
Content-Disposition: form-data; name="express"

1
-----------------------------3672811325816669204907760291
Content-Disposition: form-data; name="lang"

en
-----------------------------3672811325816669204907760291--
~~~

## 2

~~~
GET https://www.sbs.com.au/api/v3/video_stream?context=odwebsite&id=2229616195516&audio=demuxed HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/json, text/plain, */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
x-newrelic-id: Vg4DV1RbCxADU1VXBAUGX1w=
newrelic: eyJ2IjpbMCwxXSwiZCI6eyJ0eSI6IkJyb3dzZXIiLCJhYyI6IjI4NDYyODMiLCJhcCI6IjM3NDUzNDM2NCIsImlkIjoiZTVmNDQzN2E0N2EwNjg2NiIsInRyIjoiYzMwMDA2MGZjMGIyMGY1NjBlZTIwMjNkNTNhNDM0ZjMiLCJ0aSI6MTcxNDc3MjAwOTU3NywidGsiOiIxMjU5NjM3In19
traceparent: 00-c300060fc0b20f560ee2023d53a434f3-e5f4437a47a06866-01
tracestate: 1259637@nr=0-1-2846283-374534364-e5f4437a47a06866----1714772009577
authorization: Bearer odwebsite9d30c7144aac68fbdf0caa716037371b17b9e149
referer: https://www.sbs.com.au/ondemand/watch/2229616195516
cookie: janrainCaptureTokenRefresh_session=session
cookie: janrainFailedLogins_session=session
cookie: janrainSSO_session=session
cookie: core_t=tfrpje9fgh77wqb8
cookie: sbs_session_checktime=1746394388663
cookie: sbs_session=odwebsite9d30c7144aac68fbdf0caa716037371b17b9e149
cookie: auth.refresh-token=ew6M5PPVC-lUR6XFyyHRg_TpHVv0dgHEdfGiIe5lGdTTvNqlQoKqNkyOxtw-lmX0
cookie: auth.authenticated=true
cookie: bitmovin_analytics_uuid=c93384d0-57b8-4aba-a37c-63df2aa241f4
cookie: kndctr_5BD3248D541C319B0A4C98C6_AdobeOrg_cluster=aus3
cookie: kndctr_5BD3248D541C319B0A4C98C6_AdobeOrg_identity=CiY0NjA1Mjk3MDMyMDE0MDIyMDM1MDkwODA4OTE1MzIzOTE2Nzk5OFITCOCDgYP0MRABGAEqBEFVUzMwAKAB6IOBg%5FQxqAHOyNfzi%5FTU3BiwAQDwAeCDgYP0MQ%3D%3D
cookie: AMCV_5BD3248D541C319B0A4C98C6%40AdobeOrg=MCMID|46052970320140220350908089153239167998
cookie: _gcl_au=1.1.1783752359.1714772005
cookie: nol_fpid=ubp3emz6k3ccs2xpoyqffhl9lft0s1714772006|1714772006385|1714772006385|1714772006385
cookie: s_ecid=MCMID%7C46052970320140220350908089153239167998
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-origin
te: trailers
content-length: 0
~~~

## 3

~~~
POST https://pubads.g.doubleclick.net/ondemand/hls/content/2488267/vid/2229616195516A/streams HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-type: application/x-www-form-urlencoded;charset=utf-8
content-length: 1198
origin: https://www.sbs.com.au
referer: https://www.sbs.com.au/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers

cust_params=cxsegments%3D%26cxprnd%3D25460999%26cxid%3D077f2646-7b68-485a-9112-846f1a46344d%26cxsiteid%3D1133921284657819058%26cxurl%3Dhttp%3A%2F%2Fwww.sbs.com.au%2Fondemand%2Fvideo%2Fsingle%2F2229616195516%26device%3Dweb%26ipaddress%3D103.136.147.29%26partner%3Dnone%26programname%3Dcloser%26genre%3DFilm%2CDrama%26season%3D%26ratings%3Dma15%2B%26vid%3D2229616195516%26scor%3D2229616195516%26uid%3D077f2646-7b68-485a-9112-846f1a46344d%26ifa%3D%26ifaoptout%3D0%26tvid%3D%26oztam%3Dcf92a103-2330-c1dc-ac74-ce46fd2638e4&ppid=077F26467B68485A9112846F1A46344D&iu=%2F4117%2Fvideo.web.sbs.com.au&npa=0&description_url=https%3A%2F%2Fwww.sbs.com.au%2Fondemand%2Fvideo%2F2229616195516&ipaddress=103.136.147.29&ctv=0&correlator=2484860642866773&ptt=20&osd=2&sdr=1&sdki=41&sdkv=h.3.639.0&uach=null&ua=Mozilla%2F5.0%20(Windows%20NT%2010.0%3B%20Win64%3B%20x64%3B%20rv%3A109.0)%20Gecko%2F20100101%20Firefox%2F111.0&eid=44777649%2C44781409%2C95321947%2C95322027%2C95323893%2C95324128%2C95324210%2C95326337%2C95329629%2C95331589&frm=0&omid_p=Google1%2Fh.3.639.0&sdk_apis=7&wta=0&sid=29936BB7-54C3-436F-A8D7-7DA72E65AED1&ssss=gima&url=https%3A%2F%2Fwww.sbs.com.au%2Fondemand%2Fwatch%2F2229616195516&cookie_enabled=1
~~~
