# peacock

how to get SkyOTT key?

## account

1. https://privacy.com
2. New Card
3. Create Card
4. $10
5. Single-Use
6. Set $10 Spend Limit
7. https://peacocktv.com/plans/all-monthly
8. Monthly
9. GET PREMIUM
10. Email
11. Password
12. Re-enter Password
13. First Name
14. Last Name
15. Gender
16. Birth Year
17. Zip Code
18. CREATE ACCOUNT
19. first name
20. last name
21. address
22. city
23. state
24. zip
25. card number
26. expiry date
27. CVC
28. SUBSCRIBE
29. PAY NOW

## android

~~~
> play -a com.peacocktv.peacockandroid
downloads: 34.82 million
files: APK APK APK APK
name: Peacock TV: Stream TV & Movies
offered by: Peacock TV LLC
price: 0 USD
requires: 7.0 and up
size: 67.11 megabyte
updated on: Feb 7, 2024
version code: 124050214
version name: 5.2.14
~~~

https://play.google.com/store/apps/details?id=com.peacocktv.peacockandroid

If you start the app and Sign In, this request:

~~~
POST https://rango.id.peacocktv.com/signin/service/international HTTP/2.0
content-type: application/x-www-form-urlencoded
x-skyott-device: MOBILE
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US

userIdentifier=MY_EMAIL&password=MY_PASSWORD
~~~

will fail:

~~~
HTTP/2.0 429
~~~

You can fix this problem by removing this request header before starting the
app:

~~~
set modify_headers '/~u signin.service.international/x-skyott-device/'
~~~

Header needs to be removed from that request only, since other requests need the
header.

## thanks

https://github.com/Paco8/plugin.video.skyott/blob/main/resources/lib/signature.py

## web

you can get `x-skyott-usertoken` with web client via `/auth/tokens`, but it
need `idsession` cookie. Looks like Android is the same.
