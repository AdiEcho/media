# NBC

only requires Android 5:

~~~
> play -a com.nbcuni.nbc
requires: 5.0 and up
version code: 2000004392
version name: 9.4.1
~~~

which means we can install up to Android 6 and install user certificate. If we
do that and start the app, we get:

> Unable to Connect
>
> Please check your internet connection and try again.

If we FORCE STOP the app, set No proxy and try again, it works. next we can try
system certificate. same result. this result also repeats with older versions
back to:

~~~
package: name='com.nbcuni.nbc' versionCode='2000003396' versionName='9.0.0'
compileSdkVersion='31' compileSdkVersionCodename='12'
~~~

with older versions:

~~~
package: name='com.nbcuni.nbc' versionCode='2000003368' versionName='7.35.1'
compileSdkVersion='31' compileSdkVersionCodename='12'
~~~

we get:

> ACTION REQUIRED
>
> To keep watching your favorite shows, movies and more, you'll need to update
> to the latest version of the NBC App.

next we can try:

https://github.com/httptoolkit/frida-interception-and-unpinning/tree/feddf62

~~~
pip install frida-tools
~~~

download and extract server:

https://github.com/frida/frida/releases

for example:

~~~
frida-server-16.1.4-android-x86.xz
~~~

install app, then push server:

~~~
adb push frida-server-16.1.4-android-x86 /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server
~~~

start MITM proxy, then start Frida:

~~~
frida -U -f com.nbcuni.nbc -l frida-script.js
~~~

app gets stuck on loading screen. what if we install user certificate? same
result. what if we install system certificate? same result. next we can try
main branch:

https://github.com/httptoolkit/frida-interception-and-unpinning

~~~
frida -U `
-l config.js `
-l android/android-certificate-unpinning.js `
-f com.nbcuni.nbc
~~~

result:

~~~
Error: Address(): argument types do not match any of:
        .overload('java.lang.String', 'int', 'javax.net.SocketFactory', 'javax.net.ssl.SSLSocketFactory', 'javax.net.ssl.HostnameVerifier', 'com.android.okhttp.CertificatePinner', 'com.android.okhttp.Authenticator', 'java.net.Proxy', 'java.util.List', 'java.util.List', 'java.net.ProxySelector')
    at X (frida/node_modules/frida-java-bridge/lib/class-factory.js:622)
    at value (frida/node_modules/frida-java-bridge/lib/class-factory.js:1141)
    at e (frida/node_modules/frida-java-bridge/lib/class-factory.js:606)
    at <anonymous> (D:\Desktop\frida-interception-and-unpinning-main\android\android-certificate-unpinning.js:170)
    at apply (native)
    at ne (frida/node_modules/frida-java-bridge/lib/class-factory.js:673)
    at <anonymous> (frida/node_modules/frida-java-bridge/lib/class-factory.js:651)
~~~
