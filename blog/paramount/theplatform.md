# link.theplatform.com

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
