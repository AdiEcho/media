# Hulu

https://hulu.com/watch/a75f279c-4156-4fb1-8371-d0ba15061009

## Sign up

If you cancel you lose access immediately

1. https://support.privacy.com/hc/articles/360012285094
2. https://signup.hulu.com/plans
3. Hulu (With Ads)
4. EMAIL
5. PASSWORD
6. NAME
7. BIRTHDATE January 1 2000
8. GENDER
9. CONTINUE
10. CARD NUMBER
11. EXPIRATION
12. CVC
13. ZIP CODE
14. SUBMIT

## Android

~~~
> play -a com.hulu.plus
requires: 5.0 and up
~~~

https://play.google.com/store/apps/details?id=com.hulu.plus

Create Android 6 device. Install user certificate. after entering password, if
you click LOG IN you get this:

> Hmm. Something’s up. Please check your internet settings and try again. If
> all’s fine on your end, visit our Help Center.

system certificate? same result. if we disable proxy? it works. next:

https://github.com/httptoolkit/frida-interception-and-unpinning

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
adb root
adb push frida-server-16.1.4-android-x86 /data/app/frida-server
adb shell chmod +x /data/app/frida-server
adb shell /data/app/frida-server
~~~

then:

~~~
frida -U `
-l config.js `
-l android/android-certificate-unpinning.js `
-f com.hulu.plus
~~~

this worked a couple of times:

~~~diff
+++ b/android/android-certificate-unpinning.js
@@ -223,7 +223,7 @@ const PINNING_FIXES = {

     'okhttp3.CertificatePinner': [
         {
-            methodName: 'check',
+            methodName: 'a',
             overload: ['java.lang.String', 'java.util.List'],
             replacement: () => NO_OP
         },
~~~

but it seems to be a race condition or something, as it only works sometimes.
like it might fail the first time, but then if I restart the app it will work.
not sure.

https://github.com/httptoolkit/frida-interception-and-unpinning/issues/55

> Hulu requires Recaptcha for authentication so just passing account credentials
> is not possible without captcha solving services. To work around this, this
> tool simply takes a Hulu session cookie.

https://github.com/jkmartindale/hulu

is this true on Android? example request:

~~~
POST https://guide.hulu.com/guide/details?user_token=nk77TZQgj1xc245G... HTTP/2.0
x-hulu-user-agent: androidv4/5.3.0+12541-google/b3d7ca343f99384;OS_23,MODEL_Android SDK built for x86
user-agent: Hulu/5.3.0+12541-google (Android 6.0; en_US; Android SDK built for x86; Build/MASTER;)
content-type: application/json; charset=UTF-8
content-length: 76
accept-encoding: gzip

{"eabs":["EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326"]}
~~~

`user_token` comes from here:

~~~
POST https://auth.hulu.com/v1/mobile/mfa/authenticate HTTP/2.0
x-hulu-user-agent: androidv4/5.3.0+12541-google/b3d7ca343f99384;OS_23,MODEL_Android SDK built for x86
user-agent: Hulu/5.3.0+12541-google (Android 6.0; en_US; Android SDK built for x86; Build/MASTER;)
content-type: application/x-www-form-urlencoded
content-length: 166
accept-encoding: gzip

code=941741&
friendly_name=Android%20-%20unknown%20Android%20SDK%20built%20for%20x86%20Android&
serial_number=b3d7ca343f99384&
token=83c42269-296c-47ea-ac62-023d02ef2a47
~~~

code is 2FA.
