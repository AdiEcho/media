# spotify

https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ

## web

~~~
HTTP/1.1 400 Bad Request
Connection: close

not a WebSocket handshake request: missing upgrade
~~~

also login is protected by `recaptchaToken`

## android

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

- https://github.com/Ahmeth4n/javatify/issues/3
- https://github.com/glomatico/spotify-aac-downloader/issues/21

## PouleR/spotify-login

https://github.com/PouleR/spotify-login

## hashcash

- https://github.com/CrazyHoneyBadger/pow/issues/1
- https://github.com/LeastAuthority/hashcash/issues/1
- https://github.com/PaulSnow/hashproof/issues/1
- https://github.com/PoW-HC/hashcash/issues/2
- https://github.com/agfy/hashcash/issues/1
- https://github.com/alextanhongpin/go-hashcash/issues/1
- https://github.com/catalinc/hashcash/issues/4
- https://github.com/cleanunicorn/hashcash/issues/1
- https://github.com/denismitr/hashcache/issues/1
- https://github.com/kirvader/wow-using-pow/issues/1
- https://github.com/laonix/pow-word-of-wisdom/issues/1
- https://github.com/maksadbek/pow/issues/1
- https://github.com/maxim-dzh/word-of-wisdom/issues/1
- https://github.com/prestonTao/hashcash/issues/1
- https://github.com/rezamt/hashcash-go/issues/5
- https://github.com/rodzevich/hashcash/issues/1
- https://github.com/rolandshoemaker/hashcash-go/issues/1
- https://github.com/sstelfox/provingwork/issues/2
- https://github.com/timurkash/hashcash/issues/1
- https://github.com/umahmood/hashcash/issues/1
- https://github.com/vxxvvxxv/hashcash/issues/1

17 imports:

https://pkg.go.dev/github.com/robotomize/powwy/pkg/hashcash
