# NBC

~~~
> play -a com.nbcuni.nbc
requires: 5.0 and up
~~~

Create device with Android API 23. Install user certificate. If you start the
app, you get:

> Unable to Connect
>
> Please check your internet connection and try again.

same error with system certificate. if we disable MITM Proxy it works.

## Frida install

~~~
pip install frida-tools
~~~

download and extract server:

https://github.com/frida/frida/releases

for example:

~~~
frida-server-16.1.4-android-x86.xz
~~~

## Frida run

install app, then push server:

~~~
adb push frida-server-16.1.4-android-x86 /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server
~~~

start MITM proxy, then start Frida:

~~~
frida -U `
-f com.nbcuni.nbc `
-l config.js `
-l android/android-certificate-unpinning.js
~~~

https://github.com/httptoolkit/frida-interception-and-unpinning/issues/51

~~~
9.0.0 fail
7.35.1 forced update
~~~
