# web

from this:

https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ

what is this called:

~~~
1oaaSrDJimABpOdCEbw2DJ
~~~

its here:

~~~
"canonical_uri": "spotify:track:1oaaSrDJimABpOdCEbw2D"
~~~

## problems

login is protected by `recaptchaToken`. also some requests are protected by
websocket:

~~~
GET /?access_token=BQANivsOWMZT8y5kWv6jsp5Owmnh9hMeau7GZMttM2AteQJ1r7eiJwkySRA-CYffnINF2VwZEkS6gtJI9eLG2HkQkwyzJ8GG3B3ZUdEAb2jIFZo4_wpILoI2oa12X-Oz-1QBFeY0l_lePhuutnWZ2Qcqu0Lp8B2ikKwRkLtPyTiX_ua6FBxVFG2Z9G3LUOLCorUswoftw235L4sQrSFpxprREm2H8mvT7IZH50qzthkYUT9RxAZWF_whc1XsSQNaIb0V_tmUG1kpK_8ku8chIKYi8DVHt0rhkuud4yuQc84AIhYNQkxjcKsHvOUp-9ZN5gqeKWtbeHPjUNljEzBo82yX HTTP/1.1
Host: guc3-dealer.spotify.com
Accept-Language: en-US,en;q=0.5
Accept: */*
Cache-Control: no-cache
Connection: keep-alive, Upgrade
Cookie: sp_t=f183547eec914ed51e2e804ea6b7dec1; sp_dc=AQBBwsLqQcfzzLMI8sWz-rA8_fX8mR-7_f3OU5WuW-XmgI3JFcnhmeC0mNw0zn6695OdPN0hleks12_F7WRA61Z5BnN0uJM1dDskG4-uFKzVdmdgY-KoiAzIXCTKEXOdWiBjNlmc2oxSoDj4GYo5KEdre1jEglI; sp_key=0bb58db5-2e26-46c6-a714-460c0167a3f6; sp_landing=https%3A%2F%2Fopen.spotify.com%2Ftrack%2F1oaaSrDJimABpOdCEbw2DJ%3Fsp_cid%3Df183547eec914ed51e2e804ea6b7dec1%26device%3Ddesktop
Origin: https://open.spotify.com
Pragma: no-cache
Sec-Fetch-Dest: websocket
Sec-Fetch-Mode: websocket
Sec-Fetch-Site: same-site
Sec-WebSocket-Extensions: permessage-deflate
Sec-WebSocket-Key: gyxe1xW8seFU4pmmRdvVcw==
Sec-WebSocket-Version: 13
Upgrade: websocket
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
~~~
