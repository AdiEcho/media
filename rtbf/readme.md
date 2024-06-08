# RTBF

~~~
url = https://www.rtbf.be/auvio/detail_i-care-a-lot?id=3201987
monetization = FREE
country = Belgium
~~~

https://justwatch.com/be/plateforme/rtbf-auvio

## android

https://play.google.com/store/apps/details?id=be.rtbf.auvio

~~~
> play -i be.rtbf.auvio -s
details[6] = RTBF
details[8] = 0 USD
details[13][1][4] = 3.1.35
details[13][1][16] = May 15, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 8.0 and up
downloads = 1.58 million
name = RTBF Auvio : direct et replay
size = 28.57 megabyte
version code = 1301035
~~~

create Android 8 device. install system certificate

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://auvio.rtbf.be/emission/i-care-a-lot-27462
~~~
