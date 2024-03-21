# apiguard3

https://github.com/Paco8/plugin.video.skyott/issues/42

it seems like some extra headers are required now:

~~~
POST https://rango.id.peacocktv.com/signin/service/international HTTP/2.0
content-type: application/x-www-form-urlencoded
q5vwyrl1ft-a: BNkVZI27NsmyFCLyQ5m96-pt08IWMe2yIzDGsqnyF3RI4NK3W1fEoNWdrjjKxcS8...
q5vwyrl1ft-b: -pcr5ty
q5vwyrl1ft-c: AIAjzVmOAQAA4DUuU7d92onTmzxKOdi03Rxv4LEP7u9BA5sqSkKmoQHIC5FD
q5vwyrl1ft-d: ABaAhIDBCKGFgQGAAYIQgISigaIAwBGAzvpCzi_33wdCpqEByAuRQwAAAAAYtt92...
q5vwyrl1ft-f: Axqi0VmOAQAAmmovVBDjSRGm5OSRPuecqrOuhYYSDZiFfvRw5HDsZUkwQSmgAUi1...
q5vwyrl1ft-z: q
x-skyott-device: COMPUTER
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US

password=PASSWORD&userIdentifier=EMAIL
~~~

otherwise you get this response:

~~~
{
  "class": [
    "error"
  ],
  "properties": {
    "eventType": "error",
    "code": "289"
  }
}
~~~

does it work for you now? also, I think I found something. I just captured the
Android client, and it uses similar headers:

~~~
POST https://rango.id.peacocktv.com/signin/service/international?continuationUrl=https%3A%2F%2Frango.id.peacocktv.com%2Foauth%2Fauthorize%2Fservice%2Finternational%3Fresponse_type%3Dtoken%26client_id%3Dnbcu_android_phone%26redirect_uri%3Dnbcu%3A%2F%2Fauth%26api_id%3Doauth HTTP/2.0
content-type: application/x-www-form-urlencoded
k1qbesdhkq-a: qL3qGeu0_Dlj77510t--kpH7UnXcFUPaf2PDfno2N4UDc-q0LbJINv1LGkT16HdX...
k1qbesdhkq-b: -a9e76m
k1qbesdhkq-c: AKBrEdZ0AQAArVdTjn_SkoEWiTMVcNyrMH-1PCzhOG5UMM4HRYxVVcquZtH8
k1qbesdhkq-d: AAZAkAiQgrIAgACAAoAQ0EKACIGOAIxVVcquZtH8XgpXmJKKud4BuzjLK430oqPu...
k1qbesdhkq-e: b;RCiVUVSqyaa9PAE52r1XUryratGPIVM9jQ2k6aPMjmFoT08aztpmZQF3CgbQVI...
k1qbesdhkq-f: AyD0FdZ0AQAA6uu8Y3xGr49wcwYCDcxUpXCNuPWJLu7fBrzMSn8NEKkh8Xy_AawR...
k1qbesdhkq-z: q
x-skyott-device: MOBILE
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US

password=PASSWORD&userIdentifier=EMAIL
~~~

I searched the other requests for `k1qbesdhkq` and didn't find anything. then I
searched online, and I did find this iOS request:

~~~
GET /signin/service/international?continuationUrl=https%3A%2F%2Frango.id.peacocktv.com%2Foauth%2Fauthorize%2Fservice%2Finternational%3Fresponse_type%3Dtoken%26client_id%3Dnowtv%26redirect_uri%3Dnowtv%3A%2F%2Fauth%26api_id%3Doauth HTTP/1.1
Host: rango.id.peacocktv.com
Content-Type: application/x-www-form-urlencoded
k1QbeSDhKQ-a: fH=KYMy5xzjbTuZtCTnOmqbviRz9XkARWUd8487I8nxr3Sf01vOdcGep1ycjsG8W...
k1QbeSDhKQ-b: wq608u
k1QbeSDhKQ-c: AGAx9sdyAQAAts_Msl1E8io7ggbDZkUnpSqi2YYTrgSb5P_sxabw9c-_wilW
k1QbeSDhKQ-d: AAYAoAqAAKAAgACAAIAQwAKIAAGm8PXPv8IpVqRAge4E21AJAHs-kJfZmS70evyC...
k1QbeSDhKQ-e: b;u2-uVGERbmX6pnB01b2h7qbOh3eHOYDtiAF0A_FvYXXPGNQ8KaGFEImDKXd5E3...
k1QbeSDhKQ-f: A0Wi-cdyAQAANDFsvDPlCy930XRA7ZU0ZSiCHgjmUMEN8kBTla2qoTyXwAeWAawR...
k1QbeSDhKQ-z: q
x-skyott-device: MOBILE
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US
~~~

https://freetexthost.net/gLTgbHQ

notice its the same headers, but they are mixed case. so I searched again for
`k1QbeSDhKQ`, and found this:

https://peacocktv.com/android/native-app/update

which has the same value (note I removed some of the JSON below):

~~~json
{
  "kernelId": "A08VCl6OAQAAq7e8OaQ7eGFmkyhLze1_FbMySJOf9qu7hs9BfluL1Alc2Rc8AUi1B9EAAAAAAABgfTJdYD2ykA==",
  "dl": 0,
  "nativeSignalHeaderPrefix": "k1QbeSDhKQ-",
  "uriBypass": [],
  "enableCVM": false,
  "urlChecksumLength": 1024,
  "bodyChecksumLength": 1024,
  "id": "android_config",
  "updateTimeout": 60000,
  "rtd": 255,
  "maxUpdateInterval": 28800000,
  "kernel": "(function o(E){var DR={},Dq={};var DK=ReferenceError,Dy=TypeErr...",
  "ck": {},
  "ttl": 86400000,
  "uriBypass2": [],
  "version": "1.1",
  "storePID": true,
  "updateInterval": 14400000,
  "updateURL": "https://www.peacocktv.com/android/native-app/update",
  "updateHeaderName": "x-17Dzmhep",
  "updateURLMap": {
    "qa": "https://www.stable-int.peacocktv.com/android/native-app/update",
    "default": "https://www.peacocktv.com/android/native-app/update",
    "prod": "https://www.peacocktv.com/android/native-app/update"
  },
  "support": 255,
  "enableNHC": false
}
~~~

also I found this URL in the web client:

<https://peacocktv.com/assets/peacock_common.js?single>

which also includes the header key, but mixed case again:

~~~js
(function(a) {
    var d = document,
        w = window,
        u = "/assets/peacock_common.js?async&seed=AMA8M16OAQAAW19BzSodn7jeK1y3BUNLpfX2Cmu8iKObSs5vROGLcIYD1HGL&q5VwYrl1FT--z=q",
        v = "OVtKqOTac",
        i = "78caf726a5be1b0d37ab40d28df3e055";
    var s = d.currentScript;
    addEventListener(v, function f(e) {
        e.stopImmediatePropagation();
        removeEventListener(v, f, !0);
        e.detail.init("A1VPNV6OAQAAs2vfONTBMldgR05XQNGAcRcup7wTKJiKu7BeALalgesblKTlAUi1B9GcuNk0wH8AAEB3AAAAAA==", "r7jVpFu61CyaQzAE20ZcgovIXSbDd4=fJ3W_Pmx8w-KLnliqNHhkYs5O9MTReBGUt", [], [1549287101, 1864292991, 1433390943, 1610369413, 1425779306, 350974642, 742174146, 202138393], document.currentScript && document.currentScript.nonce || "kUEu6K03CCWoNS2Gd8Ca5+az", document.currentScript && document.currentScript.nonce || "kUEu6K03CCWoNS2Gd8Ca5+az", [], a)
    }, !0);
    var o = s && s.nonce ? s.nonce : "";
    try {
        s && s.parentNode.removeChild(s)
    } catch (e) {} {
        var n = d.createElement("script");
        n.id = i;
        n.src = u;
        n.async = !0;
        n.nonce = o;
        d.head.appendChild(n)
    }
}(typeof arguments === "undefined" ? void 0 : arguments));
~~~

it seems this protection is called `apiguard3`:

https://github.com/fdciabdul/Dana-Decompile/blob/main/sources/com/apiguard3/domain/Config.java

other sites are using the same security:

- <https://mobile.southwest.com/sw_check/android/init>
- <https://mobile.southwest.com/sw_check/ios/init>
- https://tdousmobile.tdbank.com/v2/android/native-app/initialize

looks like some extended discussion here:

- https://github.com/davidkassa/southwest-checkin-3/issues/95
- https://github.com/pyro2927/SouthwestCheckin/issues/70
- https://github.com/xur17/southwest-alerts/issues/16
