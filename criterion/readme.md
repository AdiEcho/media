# criterion

1. criterionchannel.com/checkout
2. monthly
3. email
   - 2024-6-8@mailsac.com
4. confirm email
5. password
6. name on card
7. card number
   - they check the card number, so must use new card
8. expiration
9. CVC
10. zip
11. start subscription

## android

https://play.google.com/store/apps/details?id=com.criterionchannel

~~~
> play -i com.criterionchannel
details[6] = The Criterion Collection
details[8] = 0 USD
details[13][1][4] = 8.701.1
details[13][1][16] = Apr 8, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
downloads = 192.95 thousand
name = The Criterion Channel
size = 31.98 megabyte
version code = 11271
~~~

Create Android 6 device. Install user certificate

~~~xml
<intent-filter>
   <action android:name="android.intent.action.VIEW"/>
   <category android:name="android.intent.category.DEFAULT"/>
   <category android:name="android.intent.category.BROWSABLE"/>
   <data android:scheme="@string/scheme"/>
</intent-filter>
~~~

then:

~~~
res\values\strings.xml
797:    <string name="scheme">vhxcriterionchannel</string>
~~~
