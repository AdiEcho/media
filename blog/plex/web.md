# web

~~~
POST https://play.provider.plex.tv/playQueues?uri=provider%3A%2F%2Ftv.plex.provider.vod%2Flibrary%2Fmetadata%2F5d776d15f617c90020185cc6&type=video&continuous=1 HTTP/2.0
accept-encoding: identity
accept-language: en-US,en;q=0.5
accept: application/json
content-type: application/json
origin: https://watch.plex.tv
referer: https://watch.plex.tv/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
x-plex-client-identifier: ff8a91f5-8f93-4dba-b61b-e0f286101d29
x-plex-drm: widevine
x-plex-language: en
x-plex-product: Plex Mediaverse
x-plex-provider-version: 6.5.0
x-plex-token: fc1WPqnLdmq3J4Axt5pn
~~~

or:

~~~
GET https://vod.provider.plex.tv/library/metadata/movie:cruel-intentions HTTP/2.0
accept-encoding: identity
accept-language: en-US,en;q=0.5
accept: application/json
content-length: 0
content-type: application/json
origin: https://watch.plex.tv
referer: https://watch.plex.tv/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
x-plex-client-identifier: ff8a91f5-8f93-4dba-b61b-e0f286101d29
x-plex-language: en
x-plex-product: Plex Mediaverse
x-plex-provider-version: 6.5.0
~~~
