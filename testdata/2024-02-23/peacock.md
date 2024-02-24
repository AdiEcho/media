# peacock

<https://peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR>

this is it:

~~~
POST https://play.ovp.peacocktv.com/drm/widevine/acquirelicense?bt=8-XpyykNaikxBhxH0qaxwR63rC1TRtwE2SpWNPGkKXTKk4ZQix049AphKjASY2M_pwaZKRS9F4BORgQyR3bhHN8HPbilUsUt1-tCQ-wrvk7ly62qWLR22nJ-rZo0sHAIfiZcLEnMuw1125ncalBJcqvZqB3ktiw8MqRGJSosYQ4KBHRBuiDW3BjP6h-cxymg4f5pNuizIyvHP-njwa5SSYr-2nUC9gaEAiFl99hIX9GPc2jspsa52nEnzkP3c-igV-r6pSh6GYrtn5I8ikzvchvHufk8eO5ZNJhDpevWNmIVrJ2Gd9lcnSDS6ZuM_sJ24dcmBmYDB_CP5umKUF8_YrwvFqe35fkunpxtjIo2BsN2T5E1vt2IixD6u2p6uCEHZqekmifox12LsHzT8= HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: */*
accept-language: en-US,en;q=0.5
referer: https://www.peacocktv.com/
content-type: text/plain
x-sky-signature: SkyOTT client="NBCU-WEB-v8",signature="cYBRu9SEy2nS2rHzRvHtlKnucF4=",timestamp="1708737168",version="1.0"
origin: https://www.peacocktv.com
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers
~~~

`bt` value is required and expires quickly. it comes from here:

~~~
POST https://play.ovp.peacocktv.com/video/playouts/vod HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0
accept: application/vnd.playvod.v1+json
accept-language: en-US,en;q=0.5
x-skyott-provider: NBCU
x-skyott-territory: US
x-skyott-proposition: NBCUOTT
x-skyott-platform: PC
x-skyott-device: COMPUTER
x-skyott-pinoverride: true
x-skyott-usertoken: 13-CTnvCpv6dF15UMIhDeReOrNgasnSE+cvwqX+u7raWcahCmUim9G1dQJg311l/MwbPhAvF2BVsN57XPf+T+DHJvSb4f4vZ25jdGNdJ/fbW8YwmQInDV0Ury+V1I8/uvXLgqXQCtdQ/i23NC9RuSzTJ0LUa1Y2meoG+Vrlvy8cZSvwOxOMp6GpJB+IhZBG0iLJlYo1idT6fzD80pWPUdNM6ncp9UnlliWIh5VTXj/Fi+N6hWRgmkLshvKr0GbPVKcIY4uIV5NwslcNUAbMeI3fDaBmEfDVP7FGVM7EsayW/VbQmbu4DU5VXw5faJbINP3uDQ39LoyoH2gIcPZn7rMILVrfRgGlXabvvTDQqyTdFThChqpdVwo7rRjS0RhZGNQ3RX2CY63kKBcrJho5R/k3rj2vwIYyL++EQPHXoAnXSlUGV47JAlRq3Pi+7odT0juAtXqHuUt/Qk78RR1dehTxgzGrC5ajfl3sBgcFZD8FcZhBkFj7yvxjxaAcqA9+z5UE8ditDPSakJJxXDvVoCmH0q0yxr+DpbGWEo7JcwElv+mAoHNroezMebiQN5I/Nl3u
x-skyott-activeterritory: US
content-type: application/vnd.playvod.v1+json
x-sky-signature: SkyOTT client="NBCU-WEB-v8",signature="i9ZXGWBOY0IQM7Eehx47Kv9vfqw=",timestamp="1708737167",version="1.0"
origin: https://www.peacocktv.com
referer: https://www.peacocktv.com/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-site
te: trailers

{
  "device": {
    "capabilities": [
      {
        "protection": "WIDEVINE",
        "container": "ISOBMFF",
        "transport": "DASH",
        "acodec": "AAC",
        "vcodec": "H264"
      },
      {
        "protection": "NONE",
        "container": "ISOBMFF",
        "transport": "DASH",
        "acodec": "AAC",
        "vcodec": "H264"
      }
    ],
    "maxVideoFormat": "HD",
    "model": "PC",
    "hdcpEnabled": true
  },
  "client": {
    "thirdParties": [
      "FREEWHEEL"
    ]
  },
  "contentId": "GMO_00000000224510_02_HDSDR",
  "personaParentalControlRating": "9"
}
~~~

332 loc:

https://github.com/sdhtele/VT/blob/main/vinetrimmer/services/peacock.py

357 loc:

https://github.com/kenil261615/VineTrimmeronColab/blob/main/vinetrimmer/services/peacock.py
