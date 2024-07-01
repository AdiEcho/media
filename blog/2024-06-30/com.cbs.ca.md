# com.cbs.ca

first:

~~~
sources\p606vo\C13089a.java
public static byte[] m19774b() {
  byte[] bArr = new byte[32];
  for (int i10 = 0; i10 < 64; i10 += 2) {
      bArr[i10 / 2] = (byte) (Character.digit("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5".charAt(i10 + 1), 16) + (Character.digit("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5".charAt(i10), 16) << 4));
  }
  return bArr;
}
~~~

m19774b called:

~~~
public static String m19773a(String str) {
  String str2;
  String str3 = System.currentTimeMillis() + C3322f.f4170c + str;
  try {
      SecretKeySpec secretKeySpec = new SecretKeySpec(m19774b(), "AES");
~~~

m19773a called:

~~~
sources\p606vo\C13090b.java
String m13005d = ((C8044h) this.f34563a).m13005d("cookie_migration_secret", null);
C13089a c13089a = this.f34565c;
if (m13005d != null && m13005d.length() != 0) {
   c13089a.getClass();
   m19773a = C13089a.m19773a(m13005d);
} else {
   m19773a = C13089a.m19773a(c13089a.f34562a.f31218b);
}
~~~

`cookie_migration_secret`

~~~
sources\com\paramount\android\pplus\migrations\internal\device\C6698a.java
if (((C7902e) this.f14801d).m12879a() != DeviceType.PHONE) {
   str = "";
} else {
   str = "google_mobile";
}
((C11400a) this.f14802e).getClass();
JSONObject jSONObject = new JSONObject();
if (jSONObject.has(str)) {
   str2 = jSONObject.getString(str);
   AbstractC13072h.m19658h(str2, "getString(...)");
}
((C8044h) this.f14799b).m13010i("cookie_migration_secret", str2);
~~~

`google_mobile` not found for all versions. define `f31218b`:

~~~
sources\p462po\C11984b.java
56:        this.f31218b = "c0b1d5d6ed27a3f6";
~~~
