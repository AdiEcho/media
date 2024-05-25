# pluto

- https://play.google.com/store/apps/details?id=tv.pluto.android
- https://pluto.tv/on-demand/movies/ex-machina-2015-1-1-ptv1

## clips

pass:

- http://api.pluto.tv/v2/episodes/60d9fd1c89632c0013eb2155/clips.json
- https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155

fail:

- http://api.pluto.tv/v2/episodes/la-confidential-1997-1-1/clips.json
- https://pluto.tv/on-demand/movies/la-confidential-1997-1-1

## items

pass:

~~~
GET /v4/vod/items?ids=la-confidential-1997-1-1 HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjBhZjFiZmZmLTIzODEtNDY3NC1i...
~~~

https://pluto.tv/on-demand/movies/la-confidential-1997-1-1

pass:

~~~
GET /v4/vod/items?ids=60d9fd1c89632c0013eb2155 HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjBhZjFiZmZmLTIzODEtNDY3NC1i...
~~~

https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155

fail:

~~~
GET /v4/vod/items?ids=king-of-queens HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjBhZjFiZmZmLTIzODEtNDY3NC1i...
~~~

https://pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2

## seasons

pass:

~~~
GET /v4/vod/series/65ce5c60a3a8580013c4b64a/seasons HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjQzZDZjM2IxLTk2YTQtNDVlOS04...
~~~

https://pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02

pass:

~~~
GET /v4/vod/series/king-of-queens/seasons HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjQzZDZjM2IxLTk2YTQtNDVlOS04...
~~~

https://pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2

fail:

~~~
GET /v4/vod/series/60d9fd1c89632c0013eb2155/seasons HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjQzZDZjM2IxLTk2YTQtNDVlOS04...
~~~

https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155

## slugs

pass:

~~~
GET /v4/vod/slugs?slugs=la-confidential-1997-1-1 HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjBhZjFiZmZmLTIzODEtNDY3NC1i...
~~~

https://pluto.tv/on-demand/movies/la-confidential-1997-1-1

fail:

~~~
GET /v4/vod/slugs?slugs=60d9fd1c89632c0013eb2155 HTTP/1.1
Host: service-vod.clusters.pluto.tv
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjBhZjFiZmZmLTIzODEtNDY3NC1i...
~~~

https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155

## start episodeSlugs

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&episodeSlugs=60d9fd1c89632c0013eb2155
- https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&episodeSlugs=la-confidential-1997-1-1
- https://pluto.tv/on-demand/movies/la-confidential-1997-1-1

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&episodeSlugs=king-of-queens
- https://pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2

fail:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&episodeSlugs=65ce5c60a3a8580013c4b64a
- https://pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02

## start seriesIDs

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&seriesIDs=65ce5c60a3a8580013c4b64a
- https://pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&seriesIDs=king-of-queens
- https://pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&seriesIDs=la-confidential-1997-1-1
- https://pluto.tv/on-demand/movies/la-confidential-1997-1-1

pass:

- https://boot.pluto.tv/v4/start?appName=web&appVersion=9&clientID=9&clientModelNumber=9&drmCapabilities=widevine:L3&seriesIDs=60d9fd1c89632c0013eb2155
- https://pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155
