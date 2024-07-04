# Paramount+

## Android client

create Android 6 device. install user certificate. start video. after the
commercial you might get an error, try again.

US:

https://play.google.com/store/apps/details?id=com.cbs.app

INTL:

https://play.google.com/store/apps/details?id=com.cbs.ca

## try paramount+

1. paramountplus.com
2. try it free
3. continue
4. make sure monthly is selected, then under essential click select plan
5. if you see a bundle screen, click maybe later
6. continue
7. uncheck yes, i would like to receive marketing
8. continue
9. start paramount+

## How to get app\_secret?

~~~
sources\com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("a624d7b175f5626b");
~~~

## How to get secret\_key?

~~~
com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
SecretKeySpec secretKeySpec = new SecretKeySpec(b("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"), "AES");
~~~

## link.theplatform.com

why do we need link.theplatform.com? because its the only anonymous option.
logged out the web client has this request:

https://www.paramountplus.com/shows/mayor-of-kingstown/video/xhr/episodes/page/0/size/18/xs/0/season/3

logged in the web client embeds MPD in HTML:

https://www.paramountplus.com/movies/video/dR90xyXDa88Qi5wq7GITtmVm0gcniSar

Android client needs cookie for INTL requests:

~~~
GET https://www.paramountplus.com/apps-api/v3.1/androidphone/irdeto-control/anonymous-session-token.json?contentId=fmzz3juKlg4QZt9qKRI_C1GjVh6ysEZT&model=AOSP%20on%20IA%20Emulator&firmwareVersion=9&version=15.0.28&platform=PP_AndroidApp&locale=en-us&locale=en-us&at=ABBoPFHuygkRnnCKELRhypuq5uEAJvSiVATsY9xOASH88ibse11WuoLrFnSDf0Bv7EY%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...
~~~

and:

~~~
GET https://www.intl.paramountplus.com/apps-api/v3.1/androidtv/irdeto-control/session-token.json?contentId=fmzz3juKlg4QZt9qKRI_C1GjVh6ysEZT&model=sdk_google_atv_x86&firmwareVersion=9&version=15.0.28&platform=PPINTL_AndroidTV&locale=en-us&at=ABBoPFHuygkRnnCKELRhypuq5uEAJvSiVATsY9xOASH88ibse11WuoLrFnSDf0Bv7EY%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...
~~~

and:

~~~
GET https://www.intl.paramountplus.com/apps-api/v3.0/androidtv/movies/fmzz3juKlg4QZt9qKRI_C1GjVh6ysEZT.json?includeTrailerInfo=true&includeContentInfo=true&locale=en-us&at=ABDSbrWqqlbSWOrrXk8u9NaNdokPC88YiXcPvIFhPobM3a%2FJWNOSwiCMklwJDDJq4c0%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...
~~~

and:

~~~
GET https://www.intl.paramountplus.com/apps-api/v2.0/androidtv/video/cid/fmzz3juKlg4QZt9qKRI_C1GjVh6ysEZT.json?locale=en-us&at=ABA3WXXZwgC0rQPN9WtWEUmpHsGCFJb6NP4tGjIFVLTuScgId9WA3LdC44hdHUJysQ0%3D HTTP/2.0
cookie: CBS_COM=N0Q5NkYyRTE2QTRGQUJEQTE1QkYzQTQwREVEMjQxNUY3RjYyQkM0MkVDMzM2OD...
~~~

## aid

If we take a video like this:

<https://paramountplus.com/shows/video/vLJLNTktnWmP_KzDtTm9X7ki0SRvpZ2w>

we can get the info like this:

<https://paramountplus.com/apps-api/v2.0/androidphone/video/cid/vLJLNTktnWmP_KzDtTm9X7ki0SRvpZ2w.json?at=ABDSJZgxTScBBBO4aTPohedBu%2BEYBA1PYEHSGCtyFepqrLX%2BARQN5S9eZ8VWQ2zwNk8%3D>

which includes everything we need:

~~~
"cmsAccountId": "dJ5BDC",
"pid": "i_iqDzG1odu1",
~~~

but other videos:

<https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

are missing the `pid`:

<https://paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABDSJZgxTScBBBO4aTPohedBu%2BEYBA1PYEHSGCtyFepqrLX%2BARQN5S9eZ8VWQ2zwNk8%3D>

unless we request with cookie:

~~~
GET https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABAFzPkMhOzTnRSPIU8SHvxC3JgwGZ8eTCYf3NSJx3wm6fhNd5vz%2FfLw3TlCcaKYC%2Bc%3D HTTP/1.1
cookie: CBS_COM=N0EwMjY0MDVENTU3MzJCNzJBMEQzMkIyMDQ0MjQyQUU6MTcxMjA5OTQwMTg2OT...
~~~

which we can get like this:

~~~
POST https://www.paramountplus.com/apps-api/v2.0/androidphone/auth/login.json?at=ABDFhCKlU... HTTP/1.1
content-type: application/x-www-form-urlencoded

j_username=EMAIL&j_password=PASSWORD
~~~

we can instead use the `aid`:

<http://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

without authentication. Its a little messy though, because the current web client
and Android client no longer use the `aid`, so who knows how long it will be
around. Its in the Android source like this:

~~~
sources\com\cbs\downloader\impl\a.java
String uri = new Uri.Builder().scheme(ProxyConfig.MATCH_HTTP).
authority("link.theplatform.com").
appendPath(Constants.APPBOY_PUSH_SUMMARY_TEXT_KEY).appendPath("dJ5BDC").
appendPath(CommonUtil.Directory.MEDIA_ROOT).
appendPath(DistributedTracing.NR_GUID_ATTRIBUTE).appendPath("2198311517").
appendPath(contentId).appendQueryParameter("assetTypes", "DASH_CENC").
appendQueryParameter("formats", "MPEG-DASH").
appendQueryParameter("format", "smil").build().toString();
~~~

the last version using it is 12.0.28 (211202876). Or, if you already have a
`pid`, you can get it like this:

<http://link.theplatform.com/s/dJ5BDC/i_iqDzG1odu1?format=SMIL&Tracking=true>
