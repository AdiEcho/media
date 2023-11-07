# Hulu

https://hulu.com/watch/a75f279c-4156-4fb1-8371-d0ba15061009

If you cancel you lose access immediately

## Sign up

1. https://signup.hulu.com/plans
2. Hulu (With Ads)
3. EMAIL
4. PASSWORD
5. NAME
6. BIRTHDATE January 1 2000
7. GENDER
8. CONTINUE
9. CARD NUMBER
10. EXPIRATION
11. CVC
12. ZIP CODE
13. SUBMIT
14. Account
15. Cancel Your Subscription
16. CONTINUE TO CANCEL
17. CONTINUE TO CANCEL [again]
18. CANCEL SUBSCRIPTION

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

note this version need at least Oreo:

https://github.com/httptoolkit/frida-interception-and-unpinning/issues/52

fail:

https://github.com/httptoolkit/frida-interception-and-unpinning/issues/55
