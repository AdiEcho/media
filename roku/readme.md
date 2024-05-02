# Roku

## Web premium

https://therokuchannel.roku.com/watch/32c95b576307502b98f7fe32c4aa0a22

We can create free account:

~~~
POST https://my.roku.com/api/auth/1/login HTTP/1.1
csrf-token: sVzM79JV-kKC0kN2Jlz_PlI6vHLZ3NhqOqFk
content-type: application/json
cookie: _csrf=LHBM-wxg8GRExB8JboGxeJCC

{
  "email": "EMAIL",
  "password": "PASSWORD",
  "rememberMe": false
}
~~~

but login uses CAPTCHA:

~~~
{"error": "captcha"}
~~~

## Android

com.roku.remote is remote only:

https://play.google.com/store/apps/details?id=com.roku.remote

com.roku.trc is broken:

https://apkmirror.com/apk/roku-inc/the-roku-channel

com.roku.web.trc is The Roku Channel (Android TV):

https://play.google.com/store/apps/details?id=com.roku.web.trc

create Android 9 device. install system certificate.
