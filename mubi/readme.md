# Mubi

1. https://privacy.com
2. https://mubi.com
3. Email address
4. Yearly
5. NEXT
6. Card Number
7. Expiry Date
8. CVV
9. START FREE TRIAL

first commit Feb 19 2024:

https://github.com/3052/media/commit/9e98200605d6450ef826a4204cd8dad305ac4ed1

## Android

https://play.google.com/store/apps/details?id=com.mubi

~~~
downloads: 4.27 million
files: APK APK APK APK
name: MUBI: Curated Cinema
offered by: MUBI
price: 0 USD
requires: 5.1 and up
size: 24.37 megabyte
updated on: Feb 3, 2024
version code: 29145256
version name: 41.2
~~~

Create Android 6 device. Install user certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://mubi.com/en/us/films/passages-2022
~~~

## x-forwarded-for

doesnt seem to work
