# Mubi

1. mubi.com
2. email address
   - email.ml
3. get started
4. mubi
5. next
6. cardholder name
7. card number
8. expiry date
9. CVV
10. zip code
11. start free trial
12. mubi.com/subscription/cancel
13. other
14. continue
15. extend my free trial
16. mubi.com/subscription/cancel
17. other
18. continue
19. no thanks
20. cancel subscription

~~~
Date: Sat, 23 Nov 2024 20:02:49 -0600
Subject: mubi.com
To: Privacy Support <support@privacy.com>

I am not able to use my Privacy Mubi card on their site, it says

Please check your card has sufficient funds to complete the purchase.

even though the card has enough funds
~~~

also I cannot add Privacy card to PayPal:

~~~json
{
  "error": {
    "message": "We're sorry. We're not able to process your request right now. Please try again later.",
    "name": "ISSUER_DECLINE"
  }
}
~~~

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
