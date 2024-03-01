# android

https://play.google.com/store/apps/details?id=com.spotify.music

~~~
> play -a com.spotify.music
files: APK APK APK APK
name: Spotify: Music and Podcasts
offered by: Spotify AB
price: 0 USD
requires: 5.0 and up
size: 72.02 megabyte
updated on: Feb 26, 2024
version code: 111414784
version name: 8.9.18.512
~~~

Create Android 6 device. Install user certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ
~~~

this is it:

<https://audio-ak-spotify-com.akamaized.net/audio/f682d2a95d0e14eeef4f40b60fddde56bc6721c7?__token__=exp=1709349625~hmac=ac8977fadda0917bce483dd40be0fbf77960729e90e3a8df8111959108bb935d>

also:

<https://github.com/glomatico/spotify-aac-downloader/blob/main/spotify_aac_downloader/downloader.py>
