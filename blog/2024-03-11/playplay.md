# playplay

how to get key for these:

~~~
OGG_VORBIS_320
https://audio4-fa.scdn.co/audio/98b53c239db400b0a016145700de431f68d28f54?1710211899_20vgNGuVlb5sbsnhQDH9mpL0qOkuSTSO5_o77EL-35c=
encrypted

OGG_VORBIS_160
https://audio4-fa.scdn.co/audio/6a5f12fa51f2c1e284af707a99f3ca8696f7f62f?1710211900_-rvlzvWBXOezPzjNZ6FeiyU-J_f-H24uHaQustNyWIU=
encrypted

OGG_VORBIS_96
https://audio4-fa.scdn.co/audio/f682d2a95d0e14eeef4f40b60fddde56bc6721c7?1710211901_QonRJA6_F_2JvlyPor2NAUCq5xoaIbgFQZf0Zi3Twlg=
encrypted

AAC_24
https://audio4-fa.scdn.co/audio/064aff0bf727385bcb79e0a3d695d694d8c02cfe?1710211906_h0PYn4oGB8nKzD2b35v_Klnhn4cV-9j5tN8eCyRz2UA=
Encrypted, 1 substream, aes 
~~~

- https://github.com/JustYuuto/spotify-downloader/issues/2
- https://github.com/aditya76-git/spotiscrape-spotify-api/issues/1
- https://github.com/klassenserver7b/Klassenserver7bBot/issues/152

interesting results:

~~~
files/audio/interactive
~~~

searching this:

~~~
playplay/v1/key
~~~

we find this:

https://github.com/mIwr/SwiftySpot

post issue:

https://github.com/mIwr/SwiftySpot/issues/1

share here:

https://github.com/librespot-org/librespot-java/discussions/421

OK according to this:

~~~swift
let decoded = YAD.decrypt(bytes, safeFile.id, safeIntent.obfuscatedKey)
~~~

https://github.com/mIwr/SwiftySpot/blob/94a57db8/Example/RhythmRider/Controller/PlaybackController.swift#L499

the `obfuscatedKey` appears to just be used as is. and the decryption seems to
be normal AES-128-CTR:

https://github.com/p0rterB/YAD/blob/cbaa733f/Sources/YAD.swift#L15

check magic number:

~~~
OggS
~~~

download:

~~~swift
self._apiClient.downloadAsOnePiece(cdnLink: safeDirectLink, decryptHandler: decoder)
~~~

decrypt:

~~~swift
data = safeDecrytptHandler(data)
~~~

> Although each spotify track has many codec variants for download (ogg, mp3,
> mp4, flac and others), only OGG files can be downloaded (Free - 96, 160 kbps,
> Premium - 320 kbps)

https://github.com/mIwr/SwiftySpot#track-download-info

so lets try with:

~~~
OGG_VORBIS_160
~~~
