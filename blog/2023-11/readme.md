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

same error with system certificate. if we disable MITM Proxy it works. install
Frida:

~~~
pip install frida-tools
~~~

download and extract server:

https://github.com/frida/frida/releases

for example:

~~~
frida-server-16.0.0-android-x86.xz
~~~

then push:

~~~
adb root
adb push server /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server
~~~

then start Frida:

~~~
frida -U `
-l android/android-certificate-unpinning-fallback.js `
-l android/android-certificate-unpinning.js `
-l android/android-proxy-override.js `
-l android/android-system-certificate-injection.js `
-l config.js `
-l native-connect-hook.js `
-f com.nbcuni.nbc
~~~

https://github.com/httptoolkit/frida-interception-and-unpinning/issues/51
