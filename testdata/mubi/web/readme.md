# Mubi

~~~
url = https://mubi.com/films/mulholland-drive
monetization = FLATRATE
country = Greece
country = Italy
country = Poland
country = Romania
country = Turkey
~~~

1. https://privacy.com
2. https://mubi.com
3. Email address
4. Yearly
5. NEXT
6. Card Number
7. Expiry Date
8. CVV
9. START FREE TRIAL

authorization header is in the HTML response:

~~~
GET https://mubi.com/en/films/187/player HTTP/2.0
accept-encoding: gzip, deflate, br
accept-language: en-US,en;q=0.5
accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
cache-control: no-cache
content-length: 0
cookie: __stripe_mid=c34e24d7-...
cookie: __stripe_sid=5d3af8e4-...
cookie: __utmz=0.1234567890.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)
cookie: _mubi_session=ZUCrjPG5otHcbpwo7eOO%2B56iq3nrjVsWKrD8P3sv14u%2BXxgZJnm2...
cookie: _sp_id.c006=7a11294a-53df-46d4-93c4-c3dff10d44aa.1708308225.2.17083131...
cookie: _sp_ses.c006=*
cookie: app_startup_session=Sun%20Feb%2018%202024%2021%3A25%3A25%20GMT-0600%20(Central%20Standard%20Time)
cookie: flash_store=%7B%7D
cookie: lt=1abf84406...
cookie: mubi-cookie-consent=reject
cookie: mubi-cookie-local-storage=true
cookie: mubi_speedtest=true
cookie: muxData=mux_viewer_id=d89a0422-...
pragma: no-cache
sec-fetch-dest: document
sec-fetch-mode: navigate
sec-fetch-site: cross-site
te: trailers
upgrade-insecure-requests: 1
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~
