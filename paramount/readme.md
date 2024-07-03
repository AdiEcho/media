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
