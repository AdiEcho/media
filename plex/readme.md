# plex

## android

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
