# app secret

start here:

~~~
sources\com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("a624d7b175f5626b");
~~~

then:

~~~
sources\com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
private final String m8522a(String appSecret) {
   String str;
   String str2 = System.currentTimeMillis() + "|" + appSecret;
~~~

call m8522a:

~~~
public final String m8529g(String oldAppSecret) {
   AbstractC13105t.m39573i(oldAppSecret, "oldAppSecret");
   return m8522a(oldAppSecret);
}
~~~

call to m8529g:

~~~
sources\com\cbs\app\androiddata\retrofit\datasource\RetrofitDataSource.java
356:            m8529g = RetrofitUtil.INSTANCE.m8529g(m34348b);
~~~

String m34348b:

~~~
sources\com\cbs\app\androiddata\retrofit\datasource\RetrofitDataSource.java
352:        String m34348b = AbstractC11536l.m34348b(this$0.context);
~~~

function m34348b:

~~~
sources\com\viacbs\android\pplus\util\AbstractC11536l.java
public static String m34348b(Context context) {
  return PreferenceManager.getDefaultSharedPreferences(context).getString("cookie_migration_secret", null);
}
~~~

`cookie_migration_secret`:

~~~
sources\com\paramount\android\pplus\migrations\internal\device\DeviceMigrationImpl.java
this.f20599b.mo33308d("cookie_migration_secret", m25713f());
~~~

function m25713f:

~~~
sources\com\paramount\android\pplus\migrations\internal\device\DeviceMigrationImpl.java
private final String m25713f() {
  String str;
  if (this.f20600c.getIsAmazonBuild()) {
      int i10 = C9628b.f20603a[this.f20601d.getDeviceType().ordinal()];
      if (i10 != 1) {
          if (i10 == 2) {
              str = "amazon_tablet";
          }
          str = "";
      } else {
          str = "amazon_mobile";
      }
  } else {
      if (this.f20601d.getDeviceType() == DeviceType.PHONE) {
          str = "google_mobile";
      }
      str = "";
  }
  JSONObject mo4469a = this.f20602e.mo4469a();
  if (!mo4469a.has(str)) {
      return "";
  }
  String string = mo4469a.getString(str);
  AbstractC13105t.m39572h(string, "getString(...)");
  return string;
}
~~~

function mo4469a:

~~~
sources\p117df\C11757a.java
public JSONObject mo4469a() {
  String str = this.f29373a;
  if (str != null && str.length() != 0) {
      return new JSONObject(this.f29373a);
  }
  return new JSONObject();
}
~~~

String f29373a:

~~~
public C11757a(String str) {
  this.f29373a = str;
}
~~~

function C11757a:

~~~
sources\com\cbs\app\dagger\SharedComponentModule.java
return new C11757a("{\"amazon_tablet\":\"c4abf90e3aa8131f\",\"amazon_mobile\":\"c1353af7ed0252d8\",\"google_mobile\":\"8c4edb1155a410e4\"}");
~~~
