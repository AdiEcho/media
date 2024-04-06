# Android

https://play.google.com/store/apps/details?id=com.plexapp.android

~~~
> play -a com.plexapp.android
downloads = 55.08 million
files = APK APK APK APK
name = Plex: Stream Movies & TV
offered by = Plex, Inc.
price = 0 USD
requires = 5.0 and up
size = 87.43 megabyte (87430993)
updated on = Apr 3, 2024
version code = 952112929
version name = 10.12.1.370
~~~

1. create Android 6 device
2. install app
3. install user certificate

~~~
adb root
adb push frida-server-16.2.1-android-x86 /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server

frida -U `
-l config.js `
-l android/android-certificate-unpinning.js `
-f com.plexapp.android
~~~

## license

~~~
GET https://vod.provider.plex.tv/library/metadata/5d7768b8594b2b001e693c52?checkFiles=1&includeReviews=1&includeExtras=1&asyncAugmentMetadata=1&X-Plex-Language=en-us HTTP/2.0
accept-language: en-us
accept: application/json
user-agent: Mozilla/5.0 (Linux; Android 6.0; Android SDK built for x86 Build/MASTER; wv)PlexMobile/10.12.1.370
x-plex-advertising-donottrack: 0
x-plex-advertising-identifier: 621bba5a-646d-41d4-9e4a-202479d59127
x-plex-client-identifier: 1adb7ab9dee363ed-com-plexapp-android
x-plex-client-platform: Android
x-plex-device-screen-density: 420
x-plex-device-screen-resolution: 1920x1080 (Mobile)
x-plex-device-vendor: unknown
x-plex-device: Android SDK built for x86
x-plex-drm: widevine:video
x-plex-features: external-media,indirect-media,hub-style-list
x-plex-marketplace: googlePlay
x-plex-model: generic_x86
x-plex-platform-version: 6.0
x-plex-platform: Android
x-plex-product: Plex for Android (Mobile)
x-plex-provider-version: 6.5.0
x-plex-provides: controller,sync-target
x-plex-session-id: ffdd0a71-3725-4d2d-ba63-5989080912df
x-plex-token: V3KoARMyL631Hfjst8iP
x-plex-version: 10.12.1.370
~~~
